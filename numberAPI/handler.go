package numberAPI

import (
	"context"
	"log"
	"strings"
)

type Server struct {
}

func (s *Server) Get(ctx context.Context, in *Request) (*Response, error) {
	log.Printf("Received message: %s", in.Name)
	num := strings.TrimSpace(strings.ToLower(in.Name))
	val, _ := NumbersDict[num]
	return &Response{Value: val}, nil
}
