package userService

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"microservice/internal/core/driven_ports/repository"
	types "microservice/internal/core/types"
)

type UserService struct {
	UserRepository    repository.UserRepositoryPort
	SessionRepository repository.SessionRepositoryPort
	Logger            *log.Logger
}

func (service *UserService) GetUserRepository() repository.UserRepositoryPort {
	return service.UserRepository
}

func (service *UserService) IsLoggedIn(username string) (string, error) {
	user, err := service.UserRepository.GetUserFromName(username)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: User %v does not exist", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	session, err := service.SessionRepository.GetSession(user.ID)
	if err != nil {

		errorMessage := fmt.Sprintf("Error: User %v is not logged in", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	// TODO: Check if session expired

	service.Logger.Printf("User %v is logged in\n", username)

	return session.ID, nil
}

func (service *UserService) Login(username string, password string) (string, error) {
	var passwordHash []byte
	var salt []byte

	user, err := service.UserRepository.GetUserFromName(username)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: User %v does not exist", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	salt = []byte(user.Salt)
	passwordHash = []byte(user.PwHash)

	inputHash := hashPw([]byte(password), salt)
	if bytes.Compare(inputHash, passwordHash) != 0 {
		errorMessage := fmt.Sprintf("Error: Invalid username and password combination for user %v", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	sessionID, err := service.SessionRepository.CreateSession(user.ID)
	if err != nil {
		session, _ := service.SessionRepository.GetSession(user.ID)
		sessionID = session.ID
		service.Logger.Printf("Found existing session with id %v, it will be reused", sessionID)
	}

	service.Logger.Printf("Successfully logged in user %v\n", username)

	return sessionID, nil
}

func (service *UserService) Logout(sessionID string) error {
	err := service.SessionRepository.DeleteSession(sessionID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Failed to end session with id %v", sessionID)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return errors.New(errorMessage)
	}

	service.Logger.Printf("Successfully logged out user with session %v\n", sessionID)

	return nil
}

func (service *UserService) Register(username string, password string) (string, error) {
	salt := generateSalt(saltLength)

	passwordHash := hashPw([]byte(password), salt)

	_, err := service.UserRepository.GetUserFromName(username)
	if err == nil {
		errorMessage := fmt.Sprintf("Error: username %v is already taken", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	userID, err := service.UserRepository.Register(username, string(passwordHash), string(salt))
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Failed to register user %v", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	service.Logger.Printf("Successfully registered user %v\n", username)

	sessionID, err := service.SessionRepository.CreateSession(userID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Failed to create session for user %v", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return "", errors.New(errorMessage)
	}

	return sessionID, nil
}

func (service *UserService) CreateSession(userID string) (types.Session, error) {

	session, err := service.SessionRepository.GetSession(userID)
	if err != nil {
		return types.Session{}, nil
	}

	service.Logger.Printf("Successfully created new session for user with id %v\n", userID)

	return session, nil
}

func (service *UserService) GetSession(userID string) (types.Session, error) {
	return service.SessionRepository.GetSession(userID)
}
