package middleware

import (
	"context"
	"food_delivery/config"
	"food_delivery/response"
	"food_delivery/service"
	"net/http"
)

func ValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		cfg := c.Value("cfg").(*config.Config)

		AuthHeader := r.Header.Get("Authorization")

		tokenService := service.NewTokenService(cfg)
		accessTokenString := tokenService.GetTokenFromBearerString(AuthHeader)

		claims, err := tokenService.ValidateAccessToken(accessTokenString)
		if err != nil {
			response.SendStatusUnauthorizedError(w, err)
			return
		}

		c = context.WithValue(r.Context(), "claims", claims)
		req := r.WithContext(c)

		next.ServeHTTP(w, req)
	})
}

func ValidateRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		cfg := c.Value("cfg").(*config.Config)

		AuthHeader := r.Header.Get("Authorization")

		tokenService := service.NewTokenService(cfg)
		refreshTokenString := tokenService.GetTokenFromBearerString(AuthHeader)

		claims, err := tokenService.ValidateRefreshToken(refreshTokenString)
		if err != nil {
			response.SendStatusUnauthorizedError(w, err)
			return
		}

		c = context.WithValue(r.Context(), "claims", claims)
		req := r.WithContext(c)

		next.ServeHTTP(w, req)
	})
}
