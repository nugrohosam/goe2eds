package infrastructure

import (
	viper "github.com/spf13/viper"

	"google.golang.org/grpc"
	pb "github.com/nugrohosam/goe2eds/services/grpc/pb"
)

func ValidateToken(token string) bool {
	host := viper.GetString("authorization.grpc.host")
	port := viper.GetString("authorization.grpc.port")

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGetAuthServiceClient(conn)

	ctx := context.Background()
	req := &pb.GetAuthRequest{Token: token}
	res, err := client.GetAuthID(ctx, req)
	if err != nil {
		return false
	}

	return res.id != 0
}