# Summon
[![GoDoc](https://godoc.org/github.com/go-redis/redis?status.svg)](https://godoc.org/github.com/fenixwu/summon)

Summon makes using envirenvironment variables safely.

## Quickstart
```go
sm := summon.New(os.Getenv("ENV"))
// Inject will valid variable is exist first.
// If not, it should have panic, for example, run go like this "ENV=release go run main.go".
sm.Inject("AUTH_KEY", "token_for_local")
token , err := sm.Get("AUTH_KEY").ToString()

if err != nil {
  return
}

// token will be "token_for_local" when using `go run main.go` or `ENV=local go run main.go`.
// token will be "aaa" when using `ENV=release AUTH_KEY="aaa" go run main.go`.
fmt.Println(token)
```
