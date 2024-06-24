package drivers

import (
    "context"
    "authentication-api/models"
)

type AuthenticationDriverPort interface {
    Authenticate(ctx context.Context, credentials models.Credentials) (string, error)
}
