package utils

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

type JWTUtil struct {
    secretKey string
}

func NewJWTUtil(secretKey string) *JWTUtil {
    return &JWTUtil{
        secretKey: secretKey,
    }
}

func (ju *JWTUtil) GenerateJWT(username string, role string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "rol":      role,
        "exp":      time.Now().Add(time.Hour * 72).Unix(), // El token expira en 72 horas
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(ju.secretKey))
}

func (ju *JWTUtil) ValidateJWT(tokenString string) (map[string]interface{}, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(ju.secretKey), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil
}
