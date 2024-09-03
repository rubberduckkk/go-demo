package strings

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code uint32

const (
	OK       Code = 0
	Canceled Code = 1
	Unknown  Code = 2
)

func TestCode(t *testing.T) {
	st := status.New(codes.Unknown, "test err msg")
	fmt.Printf("%d\n", st.Code())
}
