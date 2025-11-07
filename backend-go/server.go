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
// Hash 8820
// Hash 2126
// Hash 1451
// Hash 1476
// Hash 8868
// Hash 4717
// Hash 9542
// Hash 9064
// Hash 9844
// Hash 9219
// Hash 8172
// Hash 2270
// Hash 5180
// Hash 4455
// Hash 1291
// Hash 9815
// Hash 1318
// Hash 3281
// Hash 6098
// Hash 3837
// Hash 1269
// Hash 6897
// Hash 9321
// Hash 9787
// Hash 9643
// Hash 7099
// Hash 4592
// Hash 2743
// Hash 2655
// Hash 9895
// Hash 1154
// Hash 7052
// Hash 7973
// Hash 8328
// Hash 1074
// Hash 4099
// Hash 9853
// Hash 1194
// Hash 5672
// Hash 6569
// Hash 4758
// Hash 7783
// Hash 3054
// Hash 9059
// Hash 1883
// Hash 6164
// Hash 2966
// Hash 2287
// Hash 3167
// Hash 6901
// Hash 2204
// Hash 2283
// Hash 7687
// Hash 1784
// Hash 2477
// Hash 9342
// Hash 8693
// Hash 9609
// Hash 6264
// Hash 1784
// Hash 9771
// Hash 8475
// Hash 3381
// Hash 6497
// Hash 2995
// Hash 6667
// Hash 3758
// Hash 4435
// Hash 9802
// Hash 3083
// Hash 2626
// Hash 1601
// Hash 8840
// Hash 2676
// Hash 5261
// Hash 5407
// Hash 2947
// Hash 8465
// Hash 7394
// Hash 7399
// Hash 3591
// Hash 9459
// Hash 8017
// Hash 4293
// Hash 9209
// Hash 1401
// Hash 5265
// Hash 3013
// Hash 3578
// Hash 6388
// Hash 8189
// Hash 6783
// Hash 8733
// Hash 4714
// Hash 9350
// Hash 5804
// Hash 3415
// Hash 2641
// Hash 3106
// Hash 7463
// Hash 8223
// Hash 2885
// Hash 8082
// Hash 3227
// Hash 7637
// Hash 2131
// Hash 5692
// Hash 4590
// Hash 7350
// Hash 8420
// Hash 3448