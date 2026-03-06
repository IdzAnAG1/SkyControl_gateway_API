package server

import (
	auth "sc_gateway/api/skycontrol/generated/proto/auth/v1"
	v1 "sc_gateway/api/skycontrol/viability"
	"sc_gateway/internal/conf"
	"sc_gateway/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	healthChecker *service.HealthService,
	authService *service.AuthService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.ResponseEncoder(func(writer http.ResponseWriter, request *http.Request, a any) error {
			if m, ok := a.(interface{ proto.Message }); ok {
				mo := protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				}
				b, err := mo.Marshal(m)
				if err != nil {
					return err
				}
				writer.Header().Set("Content-Type", "application/json")
				_, err = writer.Write(b)
				return err
			}
			return http.DefaultResponseEncoder(writer, request, a)
		}),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterViabilityHTTPServer(srv, healthChecker)
	auth.RegisterAuthHTTPServer(srv, authService)
	return srv
}
