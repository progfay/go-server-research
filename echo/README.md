# Echo

- Guide: https://echo.labstack.com/guide/
- Repository: https://github.com/labstack/echo

## Import

```go
import (
  "github.com/labstack/echo/v4"
)
```

## Create Instance

```go
e := echo.New()
```

## Register Endpoint

```go
e.GET(path, handler)
e.POST(path, handler)
```

## Start Server

```go
e.Start(":1323")
```

## Path Pattern

Ref. [Routing | Echo - High performance, minimalist Go web framework](https://echo.labstack.com/guide/routing/)

### Static

- e.g.: `/users/new`
- Match:
  - `/users/new`
- Not Match:
  - others

### Param

- e.g.: `/users/:id`
- Match:
  - `/users/1`
  - `/users/progfay`
- Not Match:
  - `/users`
  - `/users/1/2`

### Match Any

- e.g.: `/users/files/*`
- Match:
  - `/users/files/go.mod`
  - `/users/files/.git/CONFIG`
- Not Match:
  - `/users/attachment`
- `/users/files/*/password.txt` is works, but match `/users/files/some_path`

### Path Matching Order

- Static
- Param
- Match any
