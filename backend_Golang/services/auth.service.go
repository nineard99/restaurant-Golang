package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/nineard99/restaurant-Golang/models"
	"github.com/nineard99/restaurant-Golang/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username, password, email, role string) (string, *models.User, error) {
	if strings.TrimSpace(username) == "" {
		return "", nil, errors.New("username is required")
	}
	if strings.TrimSpace(password) == "" {
		return "", nil, errors.New("password is required")
	}

	// regex check
	validPattern := regexp.MustCompile(`^[a-zA-Z0-9_$%!@#]+$`)
	if !validPattern.MatchString(username) {
		return "", nil, errors.New("invalid username format")
	}
	if !validPattern.MatchString(password) {
		return "", nil, errors.New("invalid password format")
	}

	role = strings.ToUpper(role)
	allowedRoles := []string{"DEV", "SUPERADMIN", "CUSTOMER"}
	found := false
	for _, r := range allowedRoles {
		if r == role {
			found = true
			break
		}
	}
	if !found {
		role = "CUSTOMER" // default
	}

	// Check if username or email exists
	var count int64
	db.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return "", nil, errors.New("username is already taken")
	}
	if email != "" {
		db.Model(&models.User{}).Where("email = ?", email).Count(&count)
		if count > 0 {
			return "", nil, errors.New("email is already taken")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	userID := uuid.New().String()

	user := models.User{
		ID:       userID,
		Username: username,
		Email:    utils.NullableString(email),
		Password: string(hashedPassword),
		Role:     utils.ParseGlobalRole(role),
	}

	if err := db.Create(&user).Error; err != nil {
		return "", nil, err
	}

	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}

func LoginUser(db *gorm.DB, username, password string) (string, *models.User, error) {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return "", nil, errors.New("username and password required")
	}

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}
