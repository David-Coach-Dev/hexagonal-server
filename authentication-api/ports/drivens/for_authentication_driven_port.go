package drivens

import (
    "context"
    "authentication-api/models"
)

type AuthenticationDrivenPort interface {
    Authenticate(ctx context.Context, credentials models.Credentials) (string, error)
}
