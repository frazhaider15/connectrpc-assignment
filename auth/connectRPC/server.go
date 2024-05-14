package connectrpc

import (
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/auth/gen/authproto/v1/authprotov1connect" // generated by protoc-gen-connect-go
)

type AuthService struct{}

func StartServer() {
	url := os.Getenv("RPC_URL")
	authService := &AuthService{}
	mux := http.NewServeMux()
	path, handler := authprotov1connect.NewAuthServiceHandler(authService)
	mux.Handle(path, handler)
	http.ListenAndServe(
		url,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
