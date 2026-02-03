# ADR-001: Huma API Framework

## Status
Accepted

## Context
We need an HTTP API framework for the Nuki BLE Bridge that:
- Generates OpenAPI specs automatically for client code generation
- Provides request validation
- Supports streaming (SSE) for real-time lock status updates
- Is lightweight enough for embedded deployment (Raspberry Pi)

## Decision
Use [Huma v2](https://huma.rocks) as the HTTP API framework.

## Rationale

### Why Huma
- **OpenAPI 3.1 native**: Automatic spec generation from Go types, enabling typed client generation (Go, TypeScript)
- **JSON Schema validation**: Built-in request validation via struct tags
- **Bring your own router**: Works with standard `net/http`, Chi, Fiber, etc.
- **SSE support**: Built-in Server-Sent Events for real-time streaming
- **Lightweight**: Minimal dependencies, suitable for embedded systems
- **RFC 9457 errors**: Standard problem+json error responses

### Key Features Used

#### Operations
```go
huma.Get(api, "/status", getStatus)
huma.Post(api, "/lock", postLock)
huma.Post(api, "/unlock", postUnlock)
```

#### Request/Response Types
```go
type StatusOutput struct {
    Body struct {
        LockState       string `json:"lock_state" example:"locked" doc:"Current lock state"`
        BatteryPercent  int    `json:"battery_percent" example:"85" doc:"Battery percentage"`
    }
}
```

#### Validation Tags
| Tag | Purpose |
|-----|---------|
| `doc` | Field description |
| `example` | Example value for docs |
| `minimum`/`maximum` | Numeric bounds |
| `pattern` | Regex validation |
| `enum` | Allowed values |
| `required` | Mark as required |

#### Custom Validation (Resolvers)
```go
func (i *MyInput) Resolve(ctx huma.Context) []error {
    if i.Value == "invalid" {
        return []error{&huma.ErrorDetail{
            Location: "body.value",
            Message:  "Invalid value",
            Value:    i.Value,
        }}
    }
    return nil
}
```

#### SSE Streaming
```go
sse.Register(api, huma.Operation{
    OperationID: "lock-events",
    Method:      http.MethodGet,
    Path:        "/events",
}, map[string]any{
    "status": StatusEvent{},
    "error":  ErrorEvent{},
}, func(ctx context.Context, input *struct{}, send sse.Sender) {
    send.Data(StatusEvent{State: "locked"})
})
```

#### Middleware
```go
func AuthMiddleware(ctx huma.Context, next func(huma.Context)) {
    token := ctx.Header("Authorization")
    if !validateToken(token) {
        huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
        return
    }
    next(ctx)
}
```

### OpenAPI Spec Generation
Huma generates OpenAPI 3.1 specs. For compatibility with code generators (ogen, oapi-codegen), we use the `Downgrade()` method:

```go
spec, err := api.OpenAPI().Downgrade() // Returns OpenAPI 3.0.3
```

## Consequences

### Positive
- Type-safe API with automatic documentation
- Generated clients stay in sync with server
- Built-in validation reduces boilerplate
- SSE enables real-time lock status updates

### Negative
- OpenAPI 3.1 requires downgrade for some code generators
- Learning curve for struct tag conventions

## References

### Documentation
- [Huma Homepage](https://huma.rocks)
- [Features Overview](https://huma.rocks/features/)
- [Operations](https://huma.rocks/features/operations/)
- [Request Inputs](https://huma.rocks/features/request-inputs/)
- [Response Outputs](https://huma.rocks/features/response-outputs/)
- [Request Validation](https://huma.rocks/features/request-validation/)
- [Middleware](https://huma.rocks/features/middleware/)
- [Server-Sent Events](https://huma.rocks/features/server-sent-events-sse/)

### How-To Guides
- [Custom Validation](https://huma.rocks/how-to/custom-validation/)
- [Conditional Fields](https://huma.rocks/how-to/conditional-fields/)
- [Graceful Shutdown](https://huma.rocks/how-to/graceful-shutdown/)

### API Reference
- [pkg.go.dev/huma/v2](https://pkg.go.dev/github.com/danielgtaylor/huma/v2)
- [pkg.go.dev/huma/v2/sse](https://pkg.go.dev/github.com/danielgtaylor/huma/v2/sse)

### Source
- [GitHub](https://github.com/danielgtaylor/huma)
