package grpc

import (
	"context"
	"log"
	"testing"

	pb "github.com/nugrohosam/gocashier/services/grpc/pb"
	utilities "github.com/nugrohosam/gocashier/tests/utilities"
	viper "github.com/spf13/viper"
	assert "github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const bufSize = 1024 * 1024

// BookResponse ...
type BookResponse struct{}

// GetBook ...
func TestRun(t *testing.T) {
	InitialTest(t)
	defer utilities.DbCleaner(t)

	test1(t)
}

func test1(t *testing.T) {
	// Set up a connection to the server.
	port := viper.GetString("grpc.port")

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGetAuthServiceClient(conn)

	// Contact the server and prnt out its riesponse.
	ctx := context.Background()
	token := "Anjeng"

	req := &pb.GetAuthRequest{Token: token}
	res, err := client.GetAuthID(ctx, req)
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEmpty(t, res.Id)
}
