package main

import (
    "context"
    "authorization-api/models"
    "github.com/gin-gonic/gin"
    "authorization-api/ports/drivens"
)

type AuthorizationService struct {
    drivenAdapter drivens.AuthorizationDrivenPort
}

func NewAuthorizationService(drivenAdapter drivens.AuthorizationDrivenPort) *AuthorizationService {
    return &AuthorizationService{
        drivenAdapter: drivenAdapter,
    }
}

func (as *AuthorizationService) HandleAuthorize(c *gin.Context) {
    var req models.AuthorizationRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    authorized, err := as.drivenAdapter.Authorize(context.Background(), req.Token, req.Resource, req.Action)
    if err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"authorized": authorized})
}
