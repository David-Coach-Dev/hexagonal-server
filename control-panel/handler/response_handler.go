package handler

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "hexagonal-server/control-panel/models"
    "context"
)

type ResponseHandler struct {
    monitorAdapter MonitorPort
    logsAdapter    LogsPort
}

func NewResponseHandler(monitorAdapter MonitorPort, logsAdapter LogsPort) *ResponseHandler {
    return &ResponseHandler{
        monitorAdapter: monitorAdapter,
        logsAdapter:    logsAdapter,
    }
}

func (rh *ResponseHandler) HandleResponse(c *gin.Context, data interface{}, err error, endpoint, usuario string) {
    ctx := context.Background()
    if err != nil {
        log := models.ErrorLog{
            ID:       "0",
            Fecha:    time.Now().Format("2006-01-02"),
            Hora:     time.Now().Format("15:04:05"),
            Endpoint: endpoint,
            Usuario:  usuario,
            Error:    err.Error(),
        }
        if saveErr := rh.logsAdapter.SaveLog(ctx, log); saveErr != nil {
            log.Println("Error saving log: ", saveErr)
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    } else {
        response := models.MonitorLog{
            ID:        "0",
            Fecha:     time.Now().Format("2006-01-02"),
            Hora:      time.Now().Format("15:04:05"),
            Endpoint:  endpoint,
            Usuario:   usuario,
            Respuesta: data,
        }
        if saveErr := rh.monitorAdapter.SaveLog(ctx, response); saveErr != nil {
            log.Println("Error saving monitor log: ", saveErr)
        }
        c.JSON(http.StatusOK, data)
    }
}
