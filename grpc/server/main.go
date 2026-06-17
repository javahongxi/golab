package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/javahongxi/golab/grpc/pb"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	mu     sync.Mutex
	users  map[uint64]*pb.User
	nextID uint64
}

func newUserService() *userService {
	return &userService{
		users:  make(map[uint64]*pb.User),
		nextID: 1,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().Format(time.RFC3339)
	user := &pb.User{
		Id:         s.nextID,
		Username:   req.Username,
		Nickname:   req.Nickname,
		Gender:     req.Gender,
		Age:        req.Age,
		CreateDate: now,
		UpdateDate: now,
	}

	s.users[s.nextID] = user
	s.nextID++

	log.Printf("Created user: ID=%d, Username=%s", user.Id, user.Username)
	return &pb.CreateUserResponse{User: user}, nil
}

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &pb.GetUserResponse{User: user}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Age > 0 {
		user.Age = req.Age
	}
	user.UpdateDate = time.Now().Format(time.RFC3339)

	log.Printf("Updated user: ID=%d, Nickname=%s, Age=%d", user.Id, user.Nickname, user.Age)
	return &pb.UpdateUserResponse{User: user}, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[req.Id]; !ok {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	delete(s.users, req.Id)
	log.Printf("Deleted user: ID=%d", req.Id)
	return &pb.DeleteUserResponse{Success: true}, nil
}

func (s *userService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]*pb.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return &pb.ListUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, newUserService())

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
