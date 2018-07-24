package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/nemesisesq/vaux_server/models"
	"github.com/pkg/errors"

	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/uuid"
	log "github.com/sirupsen/logrus"
)

var ErrNoRefreshToken = errors.New("no refresh token")
var ErrNoToken = errors.New("not token present")

var ErrTokenInvalid = errors.New("token in valid")

// ValidateTokensFromHeader is a middleware that handles validating the token on each request
func ValidateTokensFromHeader(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// Validate token
		token, err := getJwtTokenFromHeader(c.Request().Header.Get("Authorization"))

		tx := c.Value("tx").(*pop.Connection)
		if err != nil {
			return c.Error(http.StatusUnauthorized, err)
		}

		if err = validateJwtToken([]byte(token)); err != nil {
			// look up the refresh token and retrieve the user
			u := models.User{}
			refreshToken := c.Request().Header.Get("Refresh-Token")
			if refreshToken == "" {
				return c.Error(http.StatusUnauthorized, ErrNoRefreshToken)
			}
			err = tx.Where("refresh_token = ?", refreshToken).First(&u)
			if err != nil {
				if errors.Cause(err) == sql.ErrNoRows {
					return c.Error(http.StatusUnauthorized, errors.New("You must login again"))
				}
			}
			// ensure the refresh token is still a valid JWT -- odds of this happening at this point is slim
			if err = validateJwtToken([]byte(refreshToken)); err != nil {
				return c.Error(http.StatusUnauthorized, err)
			}
			// since it's valid -- now refresh the JWT and return it
			twoHours := time.Now().Add(2 * time.Hour)
			claims := getJWSClaims(u, twoHours)
			token, err := generateJwtToken(claims)
			if err != nil {
				return c.Error(http.StatusUnauthorized, err)
			}
			// send the new token in the header
			c.Response().Header().Set("Access-Token", fmt.Sprintf("%s", token))
		}
		u := &models.User{}
		userID := getUIDFromJWT(token)
		tx.Eager("Profile.Addresses").Where("id = ?", userID).First(u)
		c.Set("user", u)
		return next(c)
	}
}
func getUIDFromJWT(token string) string {
	jwt, err := jws.ParseJWT([]byte(token))
	if err != nil {
		log.Fatal(err)
	}
	uid, _ := jwt.Claims().Subject()
	return uid

}
func validateJwtToken(token []byte) error {
	box := packr.NewBox("../keys")
	publicKey := box.Bytes("key.pub")
	rsaPublicKey, err := crypto.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return err
	}
	jwt, err := jws.ParseJWT([]byte(token))
	if err != nil {
		return err
	}
	if err = jwt.Validate(rsaPublicKey, crypto.SigningMethodRS512); err != nil {
		return err
	}
	return nil
}

// getJwtTokenFromHeader gets the token from the Authorization header
// removes the Bearer part from the authorization header value.
// returns No token error if Token is not found
// returns Token Invalid error if the token value cannot be obtained by removing `Bearer `
func getJwtTokenFromHeader(authString string) (string, error) {
	if authString == "" {
		return "", ErrNoToken
	}
	splitToken := strings.Split(authString, "Bearer ")
	if len(splitToken) != 2 {
		return "", ErrTokenInvalid
	}
	tokenString := splitToken[1]
	return tokenString, nil
}

// getJWSClaims is simply a helper method to DRY up things
// it takes in a User and an expiry
func getJWSClaims(user models.User, expiry time.Time) jws.Claims {
	claims := jws.Claims{}
	if !expiry.IsZero() {
		claims.SetExpiration(expiry)
	}
	claims.SetIssuedAt(time.Now())
	claims.SetJWTID(fmt.Sprintf("%v", uuid.Must(uuid.NewV4())))
	claims.SetSubject(fmt.Sprintf("%v", user.ID))
	return claims
}

// generateJwtToken will generate and return the JWT token with the claims
func generateJwtToken(claims jws.Claims) ([]byte, error) {
	box := packr.NewBox("../keys")
	privateKey := box.Bytes("key.pem")
	rsaPrivateKey, err := crypto.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	jwt := jws.NewJWT(claims, crypto.SigningMethodRS512)
	token, err := jwt.Serialize(rsaPrivateKey)
	if err != nil {
		return nil, err
	}
	return token, nil
}
