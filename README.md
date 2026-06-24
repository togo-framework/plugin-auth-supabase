<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/plugin-auth-supabase</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/plugin-auth-supabase"><img src="https://pkg.go.dev/badge/github.com/togo-framework/plugin-auth-supabase.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/plugin-auth-supabase
```

<!-- /togo-header -->

<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">plugin-auth-supabase</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

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


---

## 💎 Premium sponsors

togo is proudly sponsored by **ID8 Media** and **One Studio**.

<p align="center">
  <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
</p>

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
