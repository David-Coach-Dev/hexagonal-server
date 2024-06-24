package drivers

import (
    "context"
    "authorization-api/models"
    "authorization-api/ports/drivens"
)

type AuthorizationDriverAdapter struct {
    drivenAdapter drivens.AuthorizationDrivenPort
}

func NewAuthorizationDriverAdapter(drivenAdapter drivens.AuthorizationDrivenPort) *AuthorizationDriverAdapter {
    return &AuthorizationDriverAdapter{
        drivenAdapter: drivenAdapter,
    }
}

func (ada *AuthorizationDriverAdapter) Authorize(ctx context.Context, token string, resource string, action string) (bool, error) {
    return ada.drivenAdapter.Authorize(ctx, token, resource, action)
}
