package main

import (
    "context"
    "logs-api/models"
    "github.com/gin-gonic/gin"
    "logs-api/ports/drivens"
)

type LogsService struct {
    drivenAdapter drivens.LogsDrivenPort
}

func NewLogsService(drivenAdapter drivens.LogsDrivenPort) *LogsService {
    return &LogsService{
        drivenAdapter: drivenAdapter,
    }
}

func (ls *LogsService) HandleSaveLog(c *gin.Context) {
    var log models.ErrorLog
    if err := c.BindJSON(&log); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    err := ls.drivenAdapter.SaveLog(context.Background(), log)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Log saved successfully"})
}

func (ls *LogsService) HandleListLogs(c *gin.Context) {
    logs, err := ls.drivenAdapter.ListLogs(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"logs": logs})
}

func (ls *LogsService) HandleReadLog(c *gin.Context) {
    filename := c.Param("filename")
    log, err := ls.drivenAdapter.ReadLog(context.Background(), filename)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"log": log})
}

func (ls *LogsService) HandleDeleteLog(c *gin.Context) {
    filename := c.Param("filename")
    err := ls.drivenAdapter.DeleteLog(context.Background(), filename)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Log deleted successfully"})
}

func (ls *LogsService) HandleDeleteAllLogs(c *gin.Context) {
    err := ls.drivenAdapter.DeleteAllLogs(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "All logs deleted successfully"})
}
