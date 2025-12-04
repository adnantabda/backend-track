package data

import (
	"task_manager/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(username, password string, role models.Role) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
		Role:     role,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	result := s.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := s.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *UserService) PromoteToAdmin(userID uint) error {
	return s.db.Model(&models.User{}).Where("id = ?", userID).Update("role", models.AdminRole).Error
}

// CountUsers returns the total number of users in the database
func (s *UserService) CountUsers() (int64, error) {
	var count int64
	err := s.db.Model(&models.User{}).Count(&count).Error
	return count, err
}
