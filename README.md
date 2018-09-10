# Summon
[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://godoc.org/github.com/fenixwu/summon)

Summon makes using environment variables safely.

## Quickstart
```go
sm := summon.New(os.Getenv("ENV"))
sm.Inject("AUTH_KEY", "token_for_local")
token , err := sm.Get("AUTH_KEY").ToString()
```

## Simulation
### ENV=local go run main.go/go run main.go
```go
sm := summon.New(os.Getenv("ENV"))
sm.Inject("AUTH_KEY", "token_for_local") // Insert "ENV":"token_for_local" into map.
token , err := sm.Get("AUTH_KEY").ToString() // "token_for_local", nil
```

### ENV=release go run main.go
```go
sm := summon.New(os.Getenv("ENV"))
sm.Inject("AUTH_KEY", "token_for_local") // panic
// Try getting a null key
token , err := sm.Get("TOKEN").ToString() // return "", error:`"TOKEN" s not injected yet`
```

### ENV=release AUTH_KEY=aaa go run main.go
```go
sm := summon.New(os.Getenv("ENV"))
sm.Inject("AUTH_KEY", "token_for_local") // Insert "ENV":"aaa" into map.
token , err := sm.Get("AUTH_KEY").ToString() // "aaa", nil
```
