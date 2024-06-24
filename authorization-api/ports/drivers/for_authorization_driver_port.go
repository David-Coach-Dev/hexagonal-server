package drivers

import (
    "context"
)

type AuthorizationDriverPort interface {
    Authorize(ctx context.Context, token string, resource string, action string) (bool, error)
}
