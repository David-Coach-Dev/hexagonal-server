package drivens

import (
    "context"
    "monitor-api/models"
)

type MonitorDrivenPort interface {
    SaveLog(ctx context.Context, log models.MonitorLog) error
    ListLogs(ctx context.Context) ([]string, error)
    ReadLog(ctx context.Context, filename string) (*models.MonitorLog, error)
    DeleteLog(ctx context.Context, filename string) error
    DeleteAllLogs(ctx context.Context) error
}
