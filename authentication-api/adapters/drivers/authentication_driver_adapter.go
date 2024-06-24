package drivers

import (
    "context"
    "authentication-api/models"
    "authentication-api/ports/drivens"
)

type AuthenticationDriverAdapter struct {
    drivenAdapter drivens.AuthenticationDrivenPort
}

func NewAuthenticationDriverAdapter(drivenAdapter drivens.AuthenticationDrivenPort) *AuthenticationDriverAdapter {
    return &AuthenticationDriverAdapter{
        drivenAdapter: drivenAdapter,
    }
}

func (ada *AuthenticationDriverAdapter) Authenticate(ctx context.Context, credentials models.Credentials) (string, error) {
    return ada.drivenAdapter.Authenticate(ctx, credentials)
}
