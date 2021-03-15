package infrastructure

import (
	viper "github.com/spf13/viper"
	"context"

	"google.golang.org/grpc"
	pb "github.com/nugrohosam/goe2eds/services/grpc/pb"
)

func ValidateToken(token string) (bool, error) {
	host := viper.GetString("authorization.grpc.host")
	port := viper.GetString("authorization.grpc.port")

	grpcUri := host+":"+port
	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		return false, err
	}
	defer conn.Close()
	client := pb.NewValidationServiceClient(conn)

	ctx := context.Background()
	req := &pb.GetAuthRequest{Token: token}
	res, err := client.Validate(ctx, req)
	if err != nil {
		return false, err
	}

	return res.Valid, nil
}

func GetDataAuth(token string) (map[string]interface{}, error) {
	host := viper.GetString("authorization.grpc.host")
	port := viper.GetString("authorization.grpc.port")

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewGetAuthServiceClient(conn)

	ctx := context.Background()
	req := &pb.GetAuthRequest{Token: token}
	res, err := client.GetAuth(ctx, req)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"id" : res.Id,
		"name" : res.Name,
		"username" : res.Username,
		"email" : res.Email,
	}

	return data, nil
}