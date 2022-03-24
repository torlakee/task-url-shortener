package server

import (
	"chaos-url-shortener/deployment"
	"chaos-url-shortener/models"
	"chaos-url-shortener/templates"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ShortenUrl(app *deployment.Application) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		address := string(ctx.Request.Body())

		_, err := url.ParseRequestURI(strings.TrimSpace(address))

		if err != nil {
			errorResponse := models.ErrorResponse{
				Message:  "Невалиден адрес",
				Solution: "Моля, проверете въведения адрес",
			}

			result, _ := json.Marshal(errorResponse)

			_, _ = fmt.Fprint(ctx, string(result))

			ctx.Response.Header.Set("Content-Type", "application/json")
		} else {

			algorithm := sha256.New()
			algorithm.Write([]byte(address))

			encoded := base64.StdEncoding.EncodeToString(algorithm.Sum(nil))

			key := encoded[:5]

			scope := app.Storage.Scope("_default")
			collection := scope.Collection("urls")
			_, err := collection.Upsert(key, encoded[:5], &gocb.UpsertOptions{})

			if err != nil {
				fmt.Println(err)
				errorResponse := models.ErrorResponse{
					Message:  "Сървърен проблем",
					Solution: "Свържете се с администратора",
				}

				result, _ := json.Marshal(errorResponse)

				_, _ = fmt.Fprint(ctx, string(result))

				ctx.Response.Header.Set("Content-Type", "application/json")
			}

			_, _ = fmt.Fprint(ctx, key)

			ctx.Response.Header.Set("Content-Type", "text/html")
		}

		elapsed := time.Since(start)
		ctx.Response.Header.Set("b-elapsed", fmt.Sprintf("duration: %s", elapsed))
	}

}

// Пренасочва към оригиналния адрес
func Redirect(app *deployment.Application) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		//validate shortcode
		hash := ctx.UserValue("hash").(string)

		var isValid = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(hash)

		if isValid {

			scope := app.Storage.Scope("default")
			collection := scope.Collection("urls")

			addressResult, err := collection.Get(ctx.UserValue("hash").(string), nil)

			if err != nil {
				p := &templates.ErrorPage{
					CTX:      ctx,
					Message:  "Адресът не е открит",
					Solution: "Проверете дали адресът, който сте въвели е коректен.",
					Js:       "error.js",
					Css:      "base.css",
					Host: app.Configuration.External.Host,
				}

				templates.WritePageTemplate(ctx, p)

				ctx.Response.Header.Set("Content-Type", "text/html")
			} else {

				var address string
				err = addressResult.Content(&address)
				if err != nil {
					panic(err)
				}

				ctx.Redirect(address, 301)
			}
		} else {
			p := &templates.ErrorPage{
				CTX:      ctx,
				Message:  "Адресът не е валиден",
				Solution: "Проверете дали адресът, който сте въвели е коректен.",
				Js:       "error.js",
				Css:      "base.css",
				Host: app.Configuration.External.Host,
			}

			templates.WritePageTemplate(ctx, p)

			ctx.Response.Header.Set("Content-Type", "text/html")
		}
	}
}
func Task(app *deployment.Application) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		p := &templates.TaskPage{
			CTX: ctx,
			Js:  "docs.js",
			Css: "base.css",
			Host: app.Configuration.External.Host,
		}

		templates.WritePageTemplate(ctx, p)

		ctx.Response.Header.Set("Content-Type", "text/html")
	}
}
func TechDocs(app *deployment.Application) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		p := &templates.DocsPage{
			CTX: ctx,
			Js:  "docs.js",
			Css: "base.css",
			Host: app.Configuration.External.Host,
		}

		templates.WritePageTemplate(ctx, p)

		ctx.Response.Header.Set("Content-Type", "text/html")
	}
}
func Docs(app *deployment.Application) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		p := &templates.DocsPage{
			CTX: ctx,
			Js:  "docs.js",
			Css: "base.css",
			Host: app.Configuration.External.Host,
		}

		templates.WritePageTemplate(ctx, p)

		ctx.Response.Header.Set("Content-Type", "text/html")
	}
}

func Stylesheet(ctx *fasthttp.RequestCtx) {
	b, err := ioutil.ReadFile("res/css/base.css")
	if err != nil {
		fmt.Print(err)
	}
	ctx.Response.Header.Set("Content-Type", "text/css")
	_, _ = fmt.Fprint(ctx, string(b))
}

func Scripts(ctx *fasthttp.RequestCtx) {
	b, err := ioutil.ReadFile("scripts/core.js")
	if err != nil {
		fmt.Print(err)
	}
	ctx.Response.Header.Set("Content-Type", "application/javascript")
	_, _ = fmt.Fprint(ctx, string(b))
}

func Script(ctx *fasthttp.RequestCtx) {
	b, err := ioutil.ReadFile("scripts/" + ctx.UserValue("script").(string) + ".js")
	if err != nil {
		fmt.Print(err)
	}
	ctx.Response.Header.Set("Content-Type", "application/javascript")
	fmt.Fprint(ctx, string(b))
}

func Start(app *deployment.Application) {
	r := router.New()
	r.GET("/scripts/{script}.js", Script)
	r.GET("/{hash}", Redirect(app))
	r.POST("/", ShortenUrl(app))
	r.GET("/tech-docs", TechDocs(app))
	r.GET("/task", Task(app))
	r.GET("/", Docs(app))
	r.GET("/base.css", Stylesheet)
	r.ServeFiles("/gfx/{filepath:*}", "res/img/")
	r.ServeFiles("/pdf/{filepath:*}", "res/pdf/")

	s := &fasthttp.Server{
		Handler: r.Handler,
	}

	if err := s.ListenAndServe(":" + strconv.Itoa(app.Configuration.Internal.Server.Port)); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
