// Package authsupabase is a togo plugin that adds Supabase (GoTrue) JWT auth:
// a /auth/me endpoint and reusable middleware that validates the bearer token.
//
// Blank-import it to auto-register with the kernel:
//
//	import _ "github.com/togo-framework/plugin-auth-supabase"
package authsupabase

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/togo-framework/togo"
)

// ctxKey is the context key under which validated claims are stored.
type ctxKey struct{}

// Plugin implements togo.Plugin.
type Plugin struct {
	jwtSecret string
}

func init() { togo.Register(&Plugin{}) }

// Name identifies the plugin.
func (*Plugin) Name() string { return "auth-supabase" }

// Priority boots auth early (after infra, before feature plugins).
func (*Plugin) Priority() int { return 20 }

// Register reads the Supabase JWT secret from the environment.
func (p *Plugin) Register(k *togo.Kernel) error {
	p.jwtSecret = os.Getenv("SUPABASE_JWT_SECRET")
	return nil
}

// Boot mounts the /auth/me endpoint and fires an auth.ready hook.
func (p *Plugin) Boot(ctx context.Context, k *togo.Kernel) error {
	k.Router.Get(k.Config.RESTPath+"/auth/me", func(w http.ResponseWriter, r *http.Request) {
		claims, err := p.parse(bearer(r))
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}
		_ = json.NewEncoder(w).Encode(claims)
	})
	return k.Hooks.Fire(ctx, "auth.ready", p.Name())
}

// Middleware protects handlers: it 401s unless a valid Supabase JWT is present,
// and stores the claims in the request context (use Claims(ctx) to read them).
func (p *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := p.parse(bearer(r))
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKey{}, claims)))
	})
}

// Claims returns the validated JWT claims stored by Middleware, if any.
func Claims(ctx context.Context) (jwt.MapClaims, bool) {
	c, ok := ctx.Value(ctxKey{}).(jwt.MapClaims)
	return c, ok
}

func (p *Plugin) parse(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return []byte(p.jwtSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func bearer(r *http.Request) string {
	h := r.Header.Get("Authorization")
	return strings.TrimPrefix(h, "Bearer ")
}
