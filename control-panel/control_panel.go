package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "hexagonal-server/control-panel/models"
    "hexagonal-server/control-panel/ports/drivens"
    "hexagonal-server/control-panel/handler"
)

type ControlPanelService struct {
    drivenAdapter    drivens.ControlPanelDrivenPort
    responseHandler  *handler.ResponseHandler
}

func NewControlPanelService(drivenAdapter drivens.ControlPanelDrivenPort, responseHandler *handler.ResponseHandler) *ControlPanelService {
    return &ControlPanelService{
        drivenAdapter:   drivenAdapter,
        responseHandler: responseHandler,
    }
}

func (cps *ControlPanelService) HandleAuthenticate(c *gin.Context) {
    var credentials models.Credentials
    if err := c.BindJSON(&credentials); err != nil {
        handler.HandleError(c, err)
        return
    }
    token, err := cps.drivenAdapter.Authenticate(context.Background(), credentials)
    cps.responseHandler.HandleResponse(c, token, err, "/authenticate", credentials.Username)
}

func (cps *ControlPanelService) HandleAuthorize(c *gin.Context) {
    token := c.GetHeader("Authorization")
    resource := c.Query("resource")
    action := c.Query("action")
    authorized, err := cps.drivenAdapter.Authorize(context.Background(), token, resource, action)
    cps.responseHandler.HandleResponse(c, authorized, err, "/authorize", "unknown")
}

// Funciones de YouTube
func (cps *ControlPanelService) HandleYouTubeSearching(c *gin.Context) {
    query := c.Query("search")
    data, err := cps.drivenAdapter.GetYouTubeSearching(context.Background(), query)
    cps.responseHandler.HandleResponse(c, data, err, "/youtube/search", "unknown")
}

func (cps *ControlPanelService) HandleYouTubeRandomVideo(c *gin.Context) {
    data, err := cps.drivenAdapter.GetYouTubeRandomVideo(context.Background())
    cps.responseHandler.HandleResponse(c, data, err, "/youtube/random", "unknown")
}

func (cps *ControlPanelService) HandleYouTubePlaylists(c *gin.Context) {
    data, err := cps.drivenAdapter.GetYouTubePlaylists(context.Background())
    cps.responseHandler.HandleResponse(c, data, err, "/youtube/playlists", "unknown")
}

func (cps *ControlPanelService) HandleYouTubePlaylistItems(c *gin.Context) {
    playlistID := c.Query("playlistID")
    data, err := cps.drivenAdapter.GetYouTubePlaylistItems(context.Background(), playlistID)
    cps.responseHandler.HandleResponse(c, data, err, "/youtube/playlist-items", "unknown")
}

func (cps *ControlPanelService) HandleYouTubeLive(c *gin.Context) {
    data, err := cps.drivenAdapter.GetYouTubeLive(context.Background())
    cps.responseHandler.HandleResponse(c, data, err, "/youtube/live", "unknown")
}

// Funciones de monitorización
func (cps *ControlPanelService) HandleMonitorLog(c *gin.Context) {
    var log models.MonitorLog
    if err := c.BindJSON(&log); err != nil {
        handler.HandleError(c, err)
        return
    }
    err := cps.drivenAdapter.SaveMonitorLog(context.Background(), log)
    cps.responseHandler.HandleResponse(c, nil, err, "/internal/monitor/log", log.Usuario)
}

func (cps *ControlPanelService) HandleListMonitorLogs(c *gin.Context) {
    logs, err := cps.drivenAdapter.ListMonitorLogs(context.Background())
    cps.responseHandler.HandleResponse(c, logs, err, "/internal/monitor/logs", "unknown")
}

func (cps *ControlPanelService) HandleReadMonitorLog(c *gin.Context) {
    filename := c.Param("filename")
    log, err := cps.drivenAdapter.ReadMonitorLog(context.Background(), filename)
    cps.responseHandler.HandleResponse(c, log, err, "/internal/monitor/logs/"+filename, "unknown")
}

func (cps *ControlPanelService) HandleDeleteMonitorLog(c *gin.Context) {
    filename := c.Param("filename")
    err := cps.drivenAdapter.DeleteMonitorLog(context.Background(), filename)
    cps.responseHandler.HandleResponse(c, nil, err, "/internal/monitor/logs/"+filename, "unknown")
}

func (cps *ControlPanelService) HandleDeleteAllMonitorLogs(c *gin.Context) {
    err := cps.drivenAdapter.DeleteAllMonitorLogs(context.Background())
    cps.responseHandler.HandleResponse(c, nil, err, "/internal/monitor/logs", "unknown")
}

// Funciones de logs
func (cps *ControlPanelService) HandleErrorLog(c *gin.Context) {
    var log models.ErrorLog
    if err := c.BindJSON(&log); err != nil {
        handler.HandleError(c, err)
        return
    }
    err := cps.drivenAdapter.SaveErrorLog(context.Background(), log)
    cps.responseHandler.HandleResponse(c, nil, err, "/internal/log/error", log.Usuario)
}

func (cps *ControlPanelService) HandleListLogs(c *gin.Context) {
    logs, err := cps.drivenAdapter.ListLogs(context.Background())
    cps.responseHandler.HandleResponse(c, logs, err, "/internal/logs", "unknown")
}

func (cps *ControlPanelService) HandleReadLog(c *gin.Context) {
    filename := c.Param("filename")
    log, err := cps.drivenAdapter.ReadLog(context.Background(), filename)
    cps.responseHandler.HandleResponse(c, log, err, "/internal/logs/"+filename, "unknown")
}

func (cps *ControlPanelService) HandleDeleteLog(c *gin.Context) {
    filename := c.Param("filename")
    err := cps.drivenAdapter.DeleteLog(context.Background(), filename)
    cps.responseHandler.HandleResponse(c, nil, err, "/internal/logs/"+filename, "unknown")
}

func (cps *ControlPanelService) HandleDeleteAllLogs(c *gin.Context) {
    err := cps.drivenAdapter.DeleteAllLogs(context.Background())
    cps.responseHandler.HandleResponse(c, nil, err, "/internal/logs", "unknown")
}

func (cps *ControlPanelService) HandleHome(c *gin.Context) {
    message := "Welcome to the API"
    c.JSON(200, gin.H{"message": message})
}
