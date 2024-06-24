package handler

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
    log.Println("Error: ", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
