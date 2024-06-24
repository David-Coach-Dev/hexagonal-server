package drivens

import (
    "context"
    "logs-api/models"
)

type LogsDrivenPort interface {
    SaveLog(ctx context.Context, log models.ErrorLog) error
    ListLogs(ctx context.Context) ([]string, error)
    ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error)
    DeleteLog(ctx context.Context, filename string) error
    DeleteAllLogs(ctx context.Context) error
}
