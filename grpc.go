package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

type ServiceName string

const (
	GAMESERVICENAME   ServiceName = "GAME"
	AUTHSERVICENAME   ServiceName = "AUTH"
	SOCIALSERVICENAME ServiceName = "SOCIAL"
	RANKSERVICENAME   ServiceName = "RANK"
	MULTISERVICENAME  ServiceName = "MULTI"

	grpcHostEnvKey = "GRPC_HOST_"
	grpcPortEnvKey = "GRPC_PORT_"
)

var (
	errBuildingGrpcServiceConfig = func(name ServiceName, err error) error {
		return fmt.Errorf("error when building %s grpc service config: %w", name, err)
	}
)

type GrpcService interface {
	ConnectionString() string
}

type grpcService struct {
	host string
	port string
}

func (g grpcService) ConnectionString() string {
	return fmt.Sprintf("%s:%s", g.host, g.port)
}

func CreateGrpcServiceConfig(name ServiceName) (GrpcService, error) {

	host, err := env.MustGetEnvironmentVariable(grpcHostEnvKey + string(name))
	if err != nil {
		return nil, errBuildingGrpcServiceConfig(name, err)
	}

	port, err := env.MustGetEnvironmentVariable(grpcPortEnvKey + string(name))
	if err != nil {
		return nil, errBuildingGrpcServiceConfig(name, err)
	}

	return grpcService{
		host: host,
		port: port,
	}, nil

}
