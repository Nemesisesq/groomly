package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"fmt"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/nemesisesq/groomly/models"
	"github.com/rs/cors"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_groomly_session",
		})
		// Automatically redirect to SSL
		//app.Use(Cors())
		app.Use(forceSSL())

		// Set the request content type to JSON
		//app.Use(middleware.SetContentType("application/json"))

		app.Use(middleware.ParameterLogger)
		app.PreWares = []buffalo.PreWare{cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}).Handler}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)
		//API endpoints
		api := app.Group("/api")
		api.Resource("/opportunities", OpportunitiesResource{})
		api.Resource("/metrics", MetricsResource{})
		api.Resource("/values", ValuesResource{})
		api.Resource("/fatal_attributes", FatalAttributesResource{})
		api.Resource("/project_reports", ProjectReportsResource{})
		api.Resource("/project_reports", ProjectReportsResource{})

		//RestEndpoints
		app.Resource("/opportunities", OpportunitiesResource{})
		app.Resource("/metrics", MetricsResource{})
		app.Resource("/values", ValuesResource{})
		app.Resource("/fatal_attributes", FatalAttributesResource{})
		app.Resource("/project_reports", ProjectReportsResource{})
		app.Resource("/project_reports", ProjectReportsResource{})
		
		app.Resource("/metric_values", MetricValuesResource{})
		app.Resource("/opportunity_fatal_attributes", OpportunityFatalAttributesResource{})
		app.Resource("/metric_values", MetricValuesResource{})
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func Cors() buffalo.MiddlewareFunc {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			next(c)
			c.Response().Header().Add("Access-Control-Allow-Origin", "*")
			fmt.Println("Set Header")
			return nil
		}
	}
}
