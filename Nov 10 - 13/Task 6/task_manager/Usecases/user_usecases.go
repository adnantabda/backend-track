package usecases

import (
	"task_manager/Domain"
	"task_manager/Infrastructure"
)

// UserUseCaseImpl implements the UserUseCase interface
type UserUseCaseImpl struct {
	userRepo     domain.UserRepository
	passwordSvc *infrastructure.PasswordService
	jwtService  *infrastructure.JWTService
}

// NewUserUseCase creates a new UserUseCaseImpl instance
func NewUserUseCase(
	userRepo domain.UserRepository,
	passwordSvc *infrastructure.PasswordService,
	jwtService *infrastructure.JWTService,
) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		userRepo:     userRepo,
		passwordSvc: passwordSvc,
		jwtService:  jwtService,
	}
}

// Register creates a new user account
func (uc *UserUseCaseImpl) Register(user domain.User) (domain.User, error) {
	// Validate input
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return domain.User{}, domain.ErrInvalidInput
	}

	// Hash the password
	hashedPassword, err := uc.passwordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	user.Password = hashedPassword

	// Create the user
	createdUser, err := uc.userRepo.Create(user)
	if err != nil {
		return domain.User{}, err
	}

	// Don't return the hashed password
	createdUser.Password = ""

	return createdUser, nil
}

// Login authenticates a user and returns a JWT token
func (uc *UserUseCaseImpl) Login(email, password string) (string, error) {
	// Get user by email
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Verify password
	err = uc.passwordSvc.ComparePasswords(user.Password, password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := uc.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserProfile retrieves a user's profile by ID
func (uc *UserUseCaseImpl) GetUserProfile(id string) (domain.User, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}

	// Don't return the hashed password
	user.Password = ""

	return user, nil
}
