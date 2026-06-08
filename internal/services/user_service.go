package services

import (
	"errors"
	"fmt"

	"github.com/pablomillaquen/speckit_golang_api/internal/auth"
	"github.com/pablomillaquen/speckit_golang_api/internal/domain/user"
	"github.com/pablomillaquen/speckit_golang_api/internal/dto"
	"github.com/pablomillaquen/speckit_golang_api/internal/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo  *repositories.UserRepository
	jwtService *auth.JWTService
}

func NewUserService(userRepo *repositories.UserRepository, jwtService *auth.JWTService) *UserService {
	return &UserService{userRepo: userRepo, jwtService: jwtService}
}

func (s *UserService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	u, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, err
	}
	if !u.Active {
		return nil, fmt.Errorf("user is deactivated")
	}
	if !auth.CheckPassword(req.Password, u.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}
	pair, err := s.jwtService.GenerateTokenPair(u.ID, u.Email, string(u.Role))
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
		ExpiresIn:    pair.ExpiresIn,
		User: dto.UserBrief{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  string(u.Role),
		},
	}, nil
}

func (s *UserService) RefreshToken(tokenString string) (*dto.TokenResponse, error) {
	claims, err := s.jwtService.ValidateToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired refresh token")
	}
	pair, err := s.jwtService.GenerateTokenPair(claims.UserID, claims.Email, claims.Role)
	if err != nil {
		return nil, err
	}
	return &dto.TokenResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
		ExpiresIn:    pair.ExpiresIn,
	}, nil
}

func (s *UserService) ChangePassword(userID uint64, currentPassword, newPassword string) error {
	u, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}
	if !auth.CheckPassword(currentPassword, u.Password) {
		return fmt.Errorf("current password is incorrect")
	}
	if err := auth.ValidatePasswordPolicy(newPassword); err != nil {
		return err
	}
	hashed, err := auth.HashPassword(newPassword)
	if err != nil {
		return err
	}
	u.Password = hashed
	return s.userRepo.Update(u)
}

func (s *UserService) Create(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	if !user.Role(req.Role).IsValid() {
		return nil, fmt.Errorf("invalid role: must be administrator, technician, or viewer")
	}
	if err := auth.ValidatePasswordPolicy(req.Password); err != nil {
		return nil, err
	}
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, fmt.Errorf("email already registered")
	}
	hashed, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Role:     user.Role(req.Role),
		Active:   true,
	}
	if err := s.userRepo.Create(u); err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Role:   string(u.Role),
		Active: u.Active,
	}, nil
}

func (s *UserService) FindByID(id uint64) (*dto.UserResponse, error) {
	u, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Role:   string(u.Role),
		Active: u.Active,
	}, nil
}

func (s *UserService) FindAll(page, perPage int) ([]dto.UserResponse, int64, error) {
	users, total, err := s.userRepo.FindAll(page, perPage)
	if err != nil {
		return nil, 0, err
	}
	var responses []dto.UserResponse
	for _, u := range users {
		responses = append(responses, dto.UserResponse{
			ID:     u.ID,
			Name:   u.Name,
			Email:  u.Email,
			Role:   string(u.Role),
			Active: u.Active,
		})
	}
	return responses, total, nil
}

func (s *UserService) Update(id uint64, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	if !user.Role(req.Role).IsValid() {
		return nil, fmt.Errorf("invalid role: must be administrator, technician, or viewer")
	}
	existing, _ := s.userRepo.FindByEmailExcludingID(req.Email, id)
	if existing != nil {
		return nil, fmt.Errorf("email already registered")
	}
	u, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	u.Name = req.Name
	u.Email = req.Email
	u.Role = user.Role(req.Role)
	if req.Active != nil {
		u.Active = *req.Active
	}
	if err := s.userRepo.Update(u); err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Role:   string(u.Role),
		Active: u.Active,
	}, nil
}

func (s *UserService) Delete(id uint64) error {
	return s.userRepo.Delete(id)
}
