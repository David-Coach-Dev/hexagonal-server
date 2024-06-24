package drivens

import (
    "context"
    "encoding/json"
    "errors"
    "io/ioutil"
    "logs-api/models"
    "os"
    "path/filepath"
    "time"
)

type LogsAdapter struct {
    logDir string
}

func NewLogsAdapter() *LogsAdapter {
    logDir := "logs"
    if _, err := os.Stat(logDir); os.IsNotExist(err) {
        os.Mkdir(logDir, os.ModePerm)
    }
    return &LogsAdapter{logDir: logDir}
}

func (la *LogsAdapter) SaveLog(ctx context.Context, log models.ErrorLog) error {
    filename := filepath.Join(la.logDir, "log-"+time.Now().Format("2006-01-02-15-04-05")+".json")
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(log); err != nil {
        return err
    }

    return nil
}

func (la *LogsAdapter) ListLogs(ctx context.Context) ([]string, error) {
    files, err := ioutil.ReadDir(la.logDir)
    if err != nil {
        return nil, err
    }

    var filenames []string
    for _, file := range files {
        if !file.IsDir() {
            filenames = append(filenames, file.Name())
        }
    }
    return filenames, nil
}

func (la *LogsAdapter) ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error) {
    filepath := filepath.Join(la.logDir, filename)
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var log models.ErrorLog
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&log); err != nil {
        return nil, err
    }

    return &log, nil
}

func (la *LogsAdapter) DeleteLog(ctx context.Context, filename string) error {
    filepath := filepath.Join(la.logDir, filename)
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        return errors.New("file not found")
    }

    return os.Remove(filepath)
}

func (la *LogsAdapter) DeleteAllLogs(ctx context.Context) error {
    files, err := ioutil.ReadDir(la.logDir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            filepath := filepath.Join(la.logDir, file.Name())
            if err := os.Remove(filepath); err != nil {
                return err
            }
        }
    }
    return nil
}
