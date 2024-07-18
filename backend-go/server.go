package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 1423
// Hash 1719
// Hash 4773
// Hash 9481
// Hash 2098
// Hash 4869
// Hash 4957
// Hash 8457
// Hash 1974
// Hash 1573
// Hash 2952
// Hash 4349
// Hash 2376
// Hash 5557
// Hash 8230
// Hash 1265
// Hash 7583
// Hash 4862
// Hash 9211
// Hash 5511
// Hash 1514
// Hash 3724
// Hash 8358
// Hash 5920
// Hash 9497
// Hash 9313
// Hash 2182
// Hash 1950
// Hash 8908
// Hash 4311
// Hash 3857
// Hash 6640
// Hash 4283
// Hash 3482
// Hash 9184
// Hash 3247
// Hash 7993
// Hash 6641
// Hash 3898
// Hash 1745
// Hash 9475
// Hash 9547
// Hash 6444
// Hash 7099
// Hash 9241
// Hash 1040
// Hash 7925
// Hash 5382
// Hash 9048