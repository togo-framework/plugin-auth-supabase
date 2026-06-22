# plugin-auth-supabase

A [togo](https://github.com/togo-framework/togo) plugin adding **Supabase (GoTrue)
JWT authentication**: a `/auth/me` endpoint and reusable bearer-token middleware.

## Install

```bash
togo install togo-framework/plugin-auth-supabase
```

Then blank-import it so the kernel auto-discovers it (the CLI wires this for you):

```go
import _ "github.com/togo-framework/plugin-auth-supabase"
```

Set `SUPABASE_JWT_SECRET` in your `.env`. The plugin:

- mounts `GET /api/auth/me` — validates the `Authorization: Bearer <jwt>` token and returns the claims;
- exports `Middleware` to protect any route, and `Claims(ctx)` to read the user.

```go
r.With(authPlugin.Middleware).Get("/api/secret", handler)
```

## License

MIT
