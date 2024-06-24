package main

import (
    "github.com/gin-gonic/gin"
    controlpanel "hexagonal-server/control-panel"
)

func setupRoutes(router *gin.Engine, controlPanelService *controlpanel.ControlPanelService) {
    router.GET("/", controlPanelService.HandleHome)

    // Authentication routes
    router.POST("/authenticate", controlPanelService.HandleAuthenticate)
    router.GET("/authorize", controlPanelService.HandleAuthorize)

    // YouTube routes
    router.GET("/youtube/search", controlPanelService.HandleYouTubeSearching)
    router.GET("/youtube/random", controlPanelService.HandleYouTubeRandomVideo)
    router.GET("/youtube/playlists", controlPanelService.HandleYouTubePlaylists)
    router.GET("/youtube/playlist-items", controlPanelService.HandleYouTubePlaylistItems)
    router.GET("/youtube/live", controlPanelService.HandleYouTubeLive)

    // Monitor routes
    router.POST("/internal/monitor/log", controlPanelService.HandleMonitorLog)
    router.GET("/internal/monitor/logs", controlPanelService.HandleListMonitorLogs)
    router.GET("/internal/monitor/logs/:filename", controlPanelService.HandleReadMonitorLog)
    router.DELETE("/internal/monitor/logs/:filename", controlPanelService.HandleDeleteMonitorLog)
    router.DELETE("/internal/monitor/logs", controlPanelService.HandleDeleteAllMonitorLogs)

    // Logs routes
    router.POST("/internal/log/error", controlPanelService.HandleErrorLog)
    router.GET("/internal/logs", controlPanelService.HandleListLogs)
    router.GET("/internal/logs/:filename", controlPanelService.HandleReadLog)
    router.DELETE("/internal/logs/:filename", controlPanelService.HandleDeleteLog)
    router.DELETE("/internal/logs", controlPanelService.HandleDeleteAllLogs)
}
