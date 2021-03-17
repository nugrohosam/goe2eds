package grpc

import (
	"context"
	"log"
	"testing"

	pb "github.com/nugrohosam/goe2eds/services/grpc/pb"
	utilities "github.com/nugrohosam/goe2eds/tests/utilities"
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
}