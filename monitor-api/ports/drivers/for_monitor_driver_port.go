package drivers

import (
    "context"
    "monitor-api/models"
)

type MonitorDriverPort interface {
    SaveLog(ctx context.Context, log models.MonitorLog) error
    ListLogs(ctx context.Context) ([]string, error)
    ReadLog(ctx context.Context, filename string) (*models.MonitorLog, error)
    DeleteLog(ctx context.Context, filename string) error
    DeleteAllLogs(ctx context.Context) error
}
