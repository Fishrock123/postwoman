package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"

    // "postwoman/controllers"
)

func TemplateHandler() *echo.Echo {

    // initLoadData := map[string][]map[string]interface{} {
    //     "title": [{"Postwoman": "hi"}],
    //     "description": 1,
    //     controllers.GetAllRequests(), controllers.ValidateCreds, controllers.GetUser,
    // }

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index", "Hello Sandra!")
    })
    
    e.GET("/login", func(c echo.Context) error {
        return c.Render(http.StatusOK, "login", "login")
    })

    e.GET("/signup", func(c echo.Context) error {
        return c.Render(http.StatusOK, "signup", "signup")
    })

    return e
}