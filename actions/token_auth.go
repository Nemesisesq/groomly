package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/nemesisesq/vaux_server/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"github.com/gobuffalo/pop"
)

// TokenAuthLogin default implementation.
func TokenAuthLogin(c buffalo.Context) error {
	userForm := &models.UserForm{}

	if err := c.Bind(userForm); err != nil {
		return errors.WithStack(err)
	}

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	u := &models.User{}
	err := tx.Where("email =  ?", userForm.Email).First(u)

	if err != nil {
		return c.Error(404, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(userForm.Password))
	if err != nil {
		return errors.WithStack(err)
	}

	twoWeeks := time.Now().Add(2 * 7 * 24 * time.Hour)
	claims := getJWSClaims(*u, twoWeeks)

	refreshToken, err := generateJwtToken(claims)


	u.RefreshToken.String = string(refreshToken)

	tx.Save(u)

	twoHours := time.Now().Add(2 * time.Hour)
	claims = getJWSClaims(*u, twoHours)

	jwt, err := generateJwtToken(claims)

	if err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.Auto(c, map[string]interface{}{
		"status": "successfully logged in",
		"jwt":    string(jwt),
	}))
}
