package service

import (
	"errors"
	"user-api/internal/domain"
	"user-api/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo   domain.UserRepository
	jwtService *auth.JWTService
}

func NewUserService(userRepo domain.UserRepository, jwtService *auth.JWTService) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *UserService) RegisterUser(name, email, plainPassword string) (*domain.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user with hashed password
	user := domain.NewUser(name, email, string(hashedPassword))

	// Save to database
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) AuthenticateUser(email, plainPassword string) (*domain.User, error) {
	// Find user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password with hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *UserService) LoginUser(email, plainPassword string) (string, *domain.User, error) {
	// Authenticate user
	user, err := s.AuthenticateUser(email, plainPassword)
	if err != nil {
		return "", nil, err
	}

	// Generate JWT token
	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) UpdateUser(user *domain.User) error {
	user.UpdateTimestamp()
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) GetUserCount() (int64, error) {
	return s.userRepo.Count()
}
