package schema

import (
    "inapp_payment_service/helpers"
    "inapp_payment_service/model"

    "context"
    "fmt"
    "net/http"
    "os"
    "time"
    "crypto/sha256"

    "github.com/graphql-go/graphql"
)

func AuthMiddleware(next func(p graphql.ResolveParams) *model.GenericInatePaymentResponse) func(p graphql.ResolveParams) *model.GenericInatePaymentResponse {
    return func(p graphql.ResolveParams) *model.GenericInatePaymentResponse {
        ctx := p.Context
        userInterface := ctx.Value("user")

        var user *model.User
        if userInterface == nil {
            if req, ok := ctx.Value(model.RequestKey).(*http.Request); ok {
                authHeader := req.Header.Get("Authorization")
                u, err := helpers.ValidateToken(authHeader)
                if err != nil {
                    return helpers.FormatError(err)
                }
                if u == nil {
                    return helpers.FormatError(helpers.NewUnauthorizedError("invalid token"))
                }

                ctx = context.WithValue(ctx, "user", u)
                p.Context = ctx
                user = u
            } else {
                return helpers.FormatError(helpers.NewUnauthorizedError("invalid token"))
            }
        } else {
            user, _ = userInterface.(*model.User)
        }

        if user == nil {
            return helpers.FormatError(helpers.NewUnauthorizedError("invalid token"))
        }
        return next(p)
    }
}

func PublicAuthMiddleware(next func(p graphql.ResolveParams) *model.GenericInatePaymentResponse) func(p graphql.ResolveParams) *model.GenericInatePaymentResponse {
    return func(p graphql.ResolveParams) *model.GenericInatePaymentResponse {
        ctx := p.Context
        req, ok := ctx.Value(model.RequestKey).(*http.Request)
        if !ok {
            return helpers.FormatError(helpers.NewUnauthorizedError("invalid request"))
        }

        authHeader := req.Header.Get("Authorization")
        if authHeader == "" {
            return helpers.FormatError(helpers.NewUnauthorizedError("unauthorized"))
        }

        envPublicToken := os.Getenv("PUBLIC_ACCESS_TOKEN")
        dailyToken := generateYearlyPublicToken() // optional rotating token

        if authHeader == envPublicToken || authHeader == fmt.Sprintf("Bearer %s", envPublicToken) ||
            authHeader == dailyToken || authHeader == fmt.Sprintf("Bearer %s", dailyToken) {

            ctx = context.WithValue(ctx, "isPublic", true)
            p.Context = ctx
            return next(p)
        }

        u, err := helpers.ValidateToken(authHeader)
        if err != nil {
            return helpers.FormatError(err)
        }
        if u == nil {
            return helpers.FormatError(helpers.NewUnauthorizedError("invalid token"))
        }

        ctx = context.WithValue(ctx, model.UserKey, u)
        p.Context = ctx

        return next(p)
    }
}

func generateYearlyPublicToken() string {
    secret := os.Getenv("PUBLIC_TOKEN_SECRET")

    currentYear := time.Now().Format("2006")

    hash := sha256.Sum256([]byte(secret + currentYear))
    return fmt.Sprintf("%x", hash)
}
