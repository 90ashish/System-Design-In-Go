package main

import (
	"errors"
	"fmt"
)

// ------------------------
// Interfaces & Data Types
// ------------------------

// User represents a simple user entity.
type User struct {
	Username string
	Email    string
	Password string
}

// AuthServiceInterface defines methods for user authentication.
type AuthServiceInterface interface {
	Authenticate(username, password string) (User, error)
}

// EmailServiceInterface defines methods for sending email notifications.
type EmailServiceInterface interface {
	SendWelcomeEmail(user User) error
}

// DBServiceInterface defines methods for user database operations.
type DBServiceInterface interface {
	SaveUser(user User) error
	GetUser(username string) (User, error)
}

// ------------------------
// Implementation of Services
// ------------------------

// AuthService is responsible for user authentication.
type AuthService struct {
	db DBServiceInterface // Dependency injection for database operations.
}

func NewAuthService(db DBServiceInterface) *AuthService {
	return &AuthService{db: db}
}

// Authenticate validates user credentials.
func (a *AuthService) Authenticate(username, password string) (User, error) {
	user, err := a.db.GetUser(username)
	if err != nil {
		return User{}, err
	}
	if user.Password != password {
		return User{}, errors.New("invalid credentials")
	}
	return user, nil
}

// EmailService is responsible for sending email notifications.
type EmailService struct{}

func NewEmailService() *EmailService {
	return &EmailService{}
}

// SendWelcomeEmail sends a welcome email to the user.
func (e *EmailService) SendWelcomeEmail(user User) error {
	// Simulated email sending (in real world, integrate with an SMTP server or email API)
	fmt.Printf("Sending welcome email to %s at %s\n", user.Username, user.Email)
	return nil
}

// DBService is responsible for database operations.
type DBService struct {
	users map[string]User
}

func NewDBService() *DBService {
	return &DBService{users: make(map[string]User)}
}

// SaveUser stores the user in the database.
func (db *DBService) SaveUser(user User) error {
	db.users[user.Username] = user
	return nil
}

// GetUser retrieves the user from the database.
func (db *DBService) GetUser(username string) (User, error) {
	user, exists := db.users[username]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// ------------------------
// UserManagement System
// ------------------------

// UserManagement coordinates between the authentication, email, and database services.
type UserManagement struct {
	authService  AuthServiceInterface
	emailService EmailServiceInterface
	dbService    DBServiceInterface
}

func NewUserManagement(auth AuthServiceInterface, email EmailServiceInterface, db DBServiceInterface) *UserManagement {
	return &UserManagement{
		authService:  auth,
		emailService: email,
		dbService:    db,
	}
}

// RegisterUser handles user registration.
func (um *UserManagement) RegisterUser(username, email, password string) error {
	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}
	// Save user in database.
	if err := um.dbService.SaveUser(user); err != nil {
		return err
	}
	// Send welcome email.
	if err := um.emailService.SendWelcomeEmail(user); err != nil {
		return err
	}
	fmt.Printf("User %s registered successfully.\n", username)
	return nil
}

// LoginUser handles user login.
func (um *UserManagement) LoginUser(username, password string) error {
	user, err := um.authService.Authenticate(username, password)
	if err != nil {
		return err
	}
	fmt.Printf("User %s logged in successfully.\n", user.Username)
	return nil
}

// ------------------------
// Main Function: Demonstration
// ------------------------

func main() {
	// Initialize services
	dbService := NewDBService()
	emailService := NewEmailService()
	authService := NewAuthService(dbService)

	// Create User Management System using dependency injection
	userManagement := NewUserManagement(authService, emailService, dbService)

	// Simulate user registration
	err := userManagement.RegisterUser("johndoe", "john@example.com", "secret123")
	if err != nil {
		fmt.Println("Registration error:", err)
		return
	}

	// Simulate user login
	err = userManagement.LoginUser("johndoe", "secret123")
	if err != nil {
		fmt.Println("Login error:", err)
		return
	}
}
