package main

import (
    "context"
    "monitor-api/models"
    "github.com/gin-gonic/gin"
    "monitor-api/ports/drivens"
)

type MonitorService struct {
    drivenAdapter drivens.MonitorDrivenPort
}

func NewMonitorService(drivenAdapter drivens.MonitorDrivenPort) *MonitorService {
    return &MonitorService{
        drivenAdapter: drivenAdapter,
    }
}

func (ms *MonitorService) HandleSaveLog(c *gin.Context) {
    var log models.MonitorLog
    if err := c.BindJSON(&log); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    err := ms.drivenAdapter.SaveLog(context.Background(), log)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Log saved successfully"})
}

func (ms *MonitorService) HandleListLogs(c *gin.Context) {
    logs, err := ms.drivenAdapter.ListLogs(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"logs": logs})
}

func (ms *MonitorService) HandleReadLog(c *gin.Context) {
    filename := c.Param("filename")
    log, err := ms.drivenAdapter.ReadLog(context.Background(), filename)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"log": log})
}

func (ms *MonitorService) HandleDeleteLog(c *gin.Context) {
    filename := c.Param("filename")
    err := ms.drivenAdapter.DeleteLog(context.Background(), filename)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Log deleted successfully"})
}

func (ms *MonitorService) HandleDeleteAllLogs(c *gin.Context) {
    err := ms.drivenAdapter.DeleteAllLogs(context.Background())
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "All logs deleted successfully"})
}
