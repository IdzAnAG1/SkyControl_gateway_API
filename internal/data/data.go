package data

import (
	auth "sc_gateway/api/skycontrol/generated/proto/auth/v1"
	"sc_gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	// TODO wrapped database client
	AuthClient auth.AuthClient
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	client, err := grpc.NewClient(
		c.Auth.GetAddr(),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, nil, err
	}
	authClient := auth.NewAuthClient(client)

	cleanup := func(client *grpc.ClientConn) {
		err := client.Close()
		if err != nil {
			log.Error("failed to close grpc client: %v", err)
		}

		log.Info("closing the data resources")
	}

	return &Data{
		AuthClient: authClient,
	}, func() { cleanup(client) }, nil
}
