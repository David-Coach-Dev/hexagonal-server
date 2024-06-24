package drivens

import (
    "context"
)

type AuthorizationDrivenPort interface {
    Authorize(ctx context.Context, token string, resource string, action string) (bool, error)
}
