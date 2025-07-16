package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/nineard99/restaurant-Golang/models"
	"github.com/nineard99/restaurant-Golang/utils"
	"github.com/nineard99/restaurant-Golang/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/nineard99/restaurant-Golang/types"
)

func RegisterUser(input *types.RegisterInput) (string, *models.User, error) {

	if strings.TrimSpace(input.Username) == "" {
		return "", nil, errors.New("username is required")
	}
	if strings.TrimSpace(input.Password) == "" {
		return "", nil, errors.New("password is required")
	}

	// regex check
	validPattern := regexp.MustCompile(`^[a-zA-Z0-9_$%!@#]+$`)
	if !validPattern.MatchString(input.Username) {
		return "", nil, errors.New("invalid username format")
	}
	if !validPattern.MatchString(input.Password) {
		return "", nil, errors.New("invalid password format")
	}

	input.Role = strings.ToUpper(input.Role)
	allowedRoles := []string{"DEV", "SUPERADMIN", "CUSTOMER"}
	found := false
	for _, r := range allowedRoles {
		if r == input.Role {
			found = true
			break
		}
	}
	if !found {
		input.Role = "CUSTOMER" // default
	}

	// Check if username or email exists
	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", input.Username).Count(&count)
	if count > 0 {
		return "", nil, errors.New("username is already taken")
	}
	if input.Email != "" {
		config.DB.Model(&models.User{}).Where("email = ?", input.Email).Count(&count)
		if count > 0 {
			return "", nil, errors.New("email is already taken")
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}


	user := models.User{
		ID:       utils.GenerateUUID(),
		Username: input.Username,
		Email:    utils.NullableString(input.Email),
		Password: string(hashedPassword),
		Role:     utils.ParseGlobalRole(input.Role),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return "", nil, err
	}

	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}

func LoginUser( username, password string) (string, *models.User, error) {
	
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return "", nil, errors.New("username and password required")
	}

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
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
