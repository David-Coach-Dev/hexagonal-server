package main

import (
    "context"
    "github.com/gin-gonic/gin"
    "youtube-api/models"
    "youtube-api/ports/drivens"
)

type YouTubeService struct {
    drivenAdapter drivens.YouTubeDrivenPort
}

func NewYouTubeService(drivenAdapter drivens.YouTubeDrivenPort) *YouTubeService {
    return &YouTubeService{
        drivenAdapter: drivenAdapter,
    }
}

func (ys *YouTubeService) HandleGetSearching(c *gin.Context) {
    query := c.Query("search")
    data, err := ys.drivenAdapter.GetSearching(context.Background(), query)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": data})
}

func (ys *YouTubeService) HandleGetRandomVideo(c *gin.Context) {
    data, err := ys.drivenAdapter.GetRandomVideo(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": data})
}

func (ys *YouTubeService) HandleGetSearchPlaylists(c *gin.Context) {
    data, err := ys.drivenAdapter.GetSearchPlaylists(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": data})
}

func (ys *YouTubeService) HandleGetSearchPlaylistsItems(c *gin.Context) {
    playlistID := c.Query("playlistID")
    data, err := ys.drivenAdapter.GetSearchPlaylistsItems(context.Background(), playlistID)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": data})
}

func (ys *YouTubeService) HandleGetCheckLive(c *gin.Context) {
    data, err := ys.drivenAdapter.GetCheckLive(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"data": data})
}
