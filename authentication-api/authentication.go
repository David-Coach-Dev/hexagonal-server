package main

import (
    "context"
    "authentication-api/models"
    "github.com/gin-gonic/gin"
    "authentication-api/ports/drivens"
    "authentication-api/utils"
)

type AuthenticationService struct {
    drivenAdapter drivens.AuthenticationDrivenPort
    jwtUtil       *utils.JWTUtil
}

func NewAuthenticationService(drivenAdapter drivens.AuthenticationDrivenPort, jwtUtil *utils.JWTUtil) *AuthenticationService {
    return &AuthenticationService{
        drivenAdapter: drivenAdapter,
        jwtUtil:       jwtUtil,
    }
}

func (as *AuthenticationService) HandleAuthenticate(c *gin.Context) {
    var credentials models.Credentials
    if err := c.BindJSON(&credentials); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    token, err := as.drivenAdapter.Authenticate(context.Background(), credentials)
    if err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"token": token})
}
