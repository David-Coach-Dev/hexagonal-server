package drivers

import (
    "context"
    "logs-api/models"
)

type LogsDriverPort interface {
    SaveLog(ctx context.Context, log models.ErrorLog) error
    ListLogs(ctx context.Context) ([]string, error)
    ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error)
    DeleteLog(ctx context.Context, filename string) error
    DeleteAllLogs(ctx context.Context) error
}
