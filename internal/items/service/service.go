package service

import (
	"context"
	"log/slog"

	pb "auth-service/genproto/auth"

	"auth-service/internal/items/repository"
)

type (
	Service struct {
		pb.UnimplementedAuthServiceServer
		storage repository.IAuthRepo
		logger  *slog.Logger
	}
)

func New(storage repository.IAuthRepo, logger *slog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.logger.Info("Register function was invoked", slog.String("request", in.String()))
	return s.storage.Register(ctx, in)
}

func (s *Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.logger.Info("Login function was invoked", slog.String("request", in.String()))
	return s.storage.Login(ctx, in)
}

func (s *Service) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	s.logger.Info("Logout function was invoked", slog.String("request", in.String()))
	return s.storage.Logout(ctx, in)
}

func (s *Service) CreateAdmin(ctx context.Context, in *pb.CreateAdminRequest) (*pb.CreateAdminResponse, error) {
	s.logger.Info("CreateAdmin function was invoked", slog.String("request", in.String()))
	return s.storage.CreateAdmin(ctx, in)
}

func (s *Service) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	s.logger.Info("UpdateUser function was invoked", slog.String("request", in.String()))
	return s.storage.UpdateUser(ctx, in)
}

func (s *Service) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	s.logger.Info("DeleteUser function was invoked", slog.String("request", in.String()))
	return s.storage.DeleteUser(ctx, in)
}

func (s *Service) GetUserByEmail(ctx context.Context, in *pb.GetUserByEmailRequest) (*pb.RegisterResponse, error) {
	s.logger.Info("GetUserByEmail function was invoked", slog.String("request", in.String()))
	return s.storage.GetUserByEmail(ctx, in)
}
