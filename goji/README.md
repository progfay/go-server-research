# Echo

- GoDoc: https://pkg.go.dev/goji.io
- Repository: https://github.com/goji/goji

## Import

```go
import (
  "goji.io"
  "goji.io/pat"
)
```

## Create Instance

```go
mux := goji.NewMux()
```

## Register Endpoint

```go
mux.HandleFunc(pat.Get(path), handler)
mux.HandleFunc(pat.Post(path), handler)
```

## Start Server

```go
http.ListenAndServe("localhost:1323", mux)
```

## Path Pattern

Ref. [pat・pkg.go.dev](https://pkg.go.dev/goji.io@v2.0.2+incompatible/pat#hdr-Quick_Reference)

### Static Paths

- e.g.: `/users/new`
- Match:
  - `/users/new`
- Not Match:
  - `/users/new/`
  - others

### Named Matches

- e.g.: `/users/:id`
- Match:
  - `/users/1`
  - `/users/progfay`
- Not Match:
  - `/users`
  - `/users/1/2`
- `/users/:file.:ext`
  - Match `/users/icon.png`
    - `:file` = `icon`
    - `:ext` = `png`
  - Match `/users/profile.tar.gz`
    - `:file` = `profile`
    - `:ext` = `tar.gz`
    - Like Lazy Matching

### Prefix Matches

- e.g.: `/users/files/*`
- Match:
  - `/users/files/go.mod`
  - `/users/files/.git/CONFIG`
- Not Match:
  - `/users/attachment`

### Path Matching Order

- Static Paths
- Named Matches
- Prefix Matches
