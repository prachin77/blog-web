package client_handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func PostLogin(ctx *gin.Context) {
    email := ctx.PostForm("email")
    password := ctx.PostForm("password")

    if email == "" || password == "" {
        ctx.String(http.StatusBadRequest, "Please provide both email and password.")
        return
    }

    // TODO: Integrate with gRPC AuthService.Login

    cookie := &http.Cookie{
        Name:     "session_token",
        Value:    "demo-session-token",
        Path:     "/",
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode,
        Expires:  time.Now().Add(24 * time.Hour),
    }
    http.SetCookie(ctx.Writer, cookie)

    // Redirect to main route via HTMX
    ctx.Header("HX-Redirect", "/blogger")
    ctx.Status(http.StatusOK)
}

func PostRegister(ctx *gin.Context) {
    name := ctx.PostForm("name")
    email := ctx.PostForm("email")
    password := ctx.PostForm("password")

    if name == "" || email == "" || password == "" {
        ctx.String(http.StatusBadRequest, "All fields are required.")
        return
    }

    // TODO: Integrate with gRPC registration when available

    ctx.String(http.StatusOK, "Registration successful. You can now sign in.")
}



