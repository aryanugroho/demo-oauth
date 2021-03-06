package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aryanugroho/demo-oauth/backend/pkg/template"
	"github.com/aryanugroho/demo-oauth/backend/pkg/tracer"
	"github.com/aryanugroho/demo-oauth/backend/svc/partner/views"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/oauth2"
)

// Endpoint is OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "http://localhost:4444/oauth2/auth",
	TokenURL: "http://localhost:4444/oauth2/token",
}

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var OAuthConf = &oauth2.Config{
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	ClientID:     os.Getenv("CLIENT_ID"),     // TODO from hydra
	ClientSecret: os.Getenv("CLIENT_SECRET"), // TODO from hydra

	// https://github.com/coreos/go-oidc/blob/v3/oidc/oidc.go#L23-L36
	// offline scope for requesting Refresh Token
	// openid for Open ID Connect
	Scopes:   []string{"linkage", "deposit", "withdrawal", "payment", "transfer"},
	Endpoint: Endpoint,
}

var stateStore = map[string]bool{}

func main() {

	// Prepare Opentracing
	var (
		tracerServiceName     = "Front End"
		tracerURL             = "localhost:6831"
		tracerService, closer = tracer.New(true, tracerServiceName, tracerURL, 1)
	)

	defer func() {
		if closer == nil {
			_, _ = fmt.Fprintf(os.Stderr, "tracer closer is nil\n")
			return
		}

		if err := closer.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "closing tracer error: %s\n", err.Error())
			return
		}
	}()

	// set global tracer of this application
	opentracing.SetGlobalTracer(tracerService)

	t := &Template{}

	e := echo.New()
	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(jaegertracing.TraceWithConfig(jaegertracing.TraceConfig{
		Tracer: tracerService,
	}))

	e.GET("/", home)
	e.GET("/hi", linkage)
	e.GET("/callbacks", callback)

	if err := e.Start(":1234"); err != nil {
		log.Fatal(err)
	}
}

func linkage(c echo.Context) error {
	return c.Render(http.StatusOK, "hi.html", nil)
}

func home(c echo.Context) error {
	ctx := c.Request().Context()

	span, ctx := opentracing.StartSpanFromContext(ctx, "Homepage")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	state := base64.StdEncoding.EncodeToString(b)

	stateStore[state] = true

	// Will return loginURL,
	// for example: http://localhost:4444/oauth2/auth?client_id=myclient&prompt=consent&redirect_uri=http%3A%2F%2Fexample.com&response_type=code&scope=users.write+users.read+users.edit&state=XfFcFf7KL7ajzA2nBY%2F8%2FX3lVzZ6VZ0q7a8rM3kOfMM%3D
	loginURL := OAuthConf.AuthCodeURL(state)

	return c.Render(http.StatusOK, "homepage.html", map[string]interface{}{
		"LoginURL": loginURL,
	})
}

func callback(c echo.Context) error {
	ctx := c.Request().Context()

	span, ctx := opentracing.StartSpanFromContext(ctx, "Callback")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	code := c.QueryParam("code")
	state := c.QueryParam("state")
	scopes := c.QueryParam("scope")

	if code == "" {
		return c.String(http.StatusOK, "authorization code is empty")
	}

	// If state is exist
	if _, exist := stateStore[state]; !exist {
		return c.String(http.StatusOK, "state is generated by this Client")
	}

	delete(stateStore, state)

	// Exchange code for access token
	accessToken, err := OAuthConf.Exchange(ctx, code)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	// Print anything usefull
	fmt.Println("--------------------")
	fmt.Println("code", code)
	fmt.Println("state", state)
	fmt.Println("scopes", scopes)
	fmt.Println("accessToken", accessToken.AccessToken)
	fmt.Println("refreshToken", accessToken.RefreshToken)
	fmt.Println("idToken", accessToken.Extra("id_token"))
	fmt.Println("--------------------")

	return c.Redirect(http.StatusFound, "/hi")
}

type Template struct{}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	tmpl := template.New("", &template.BinData{
		Asset:      views.Asset,
		AssetDir:   views.AssetDir,
		AssetNames: views.AssetNames,
	})

	tpl, err := tmpl.Parse(fmt.Sprintf("../../svc/partner/views/%s", name))
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}
