package service

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	userv1 "github.com/javahongxi/golab/kratos/proto/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserService 用户服务实现
type UserService struct {
	userv1.UnimplementedUserServiceServer

	mu      sync.RWMutex
	users   map[int64]*userv1.User
	nextID  int64
	log     *log.Helper
}

// NewUserService 创建用户服务
func NewUserService(logger log.Logger) *UserService {
	return &UserService{
		users:  make(map[int64]*userv1.User),
		nextID: 1,
		log:    log.NewHelper(logger),
	}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &userv1.User{
		Id:       s.nextID,
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Age:      req.Age,
	}
	s.users[s.nextID] = user
	s.nextID++

	s.log.WithContext(ctx).Infof("created user: id=%d, username=%s", user.Id, user.Username)
	return user, nil
}

// GetUser 获取用户
func (s *UserService) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[req.Id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(ctx context.Context, req *userv1.UpdateUserRequest) (*userv1.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[req.Id]
	if !ok {
		return nil, nil
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Age > 0 {
		user.Age = req.Age
	}

	s.log.WithContext(ctx).Infof("updated user: id=%d", user.Id)
	return user, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *userv1.DeleteUserRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.users, req.Id)
	s.log.WithContext(ctx).Infof("deleted user: id=%d", req.Id)
	return &emptypb.Empty{}, nil
}

// ListUsers 列出用户
func (s *UserService) ListUsers(ctx context.Context, req *userv1.ListUsersRequest) (*userv1.ListUsersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	page := int(req.Page)
	pageSize := int(req.PageSize)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	var users []*userv1.User
	for _, user := range s.users {
		users = append(users, user)
	}

	total := len(users)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	return &userv1.ListUsersResponse{
		Users: users[start:end],
		Total: int32(total),
	}, nil
}
