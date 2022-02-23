package userService

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"microservice/internal/domain"
	drivenPorts "microservice/internal/driven_ports"
	"strconv"
)

type UserService struct {
	UserRepo    drivenPorts.UserPort
	SessionRepo drivenPorts.SessionPort
	Logger      *log.Logger
}

func (service UserService) IsLoggedIn(username string) (int, error) {
	user, err := service.UserRepo.GetUserFromName(username)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: User %v does not exist", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	session, err := service.SessionRepo.GetSession(user.ID)
	if err != nil {

		errorMessage := fmt.Sprintf("Error: User %v is not logged in", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	// TODO: Check if session expired

	service.Logger.Printf("User %v is logged in\n", username)

	return session.ID, nil
}

func (service UserService) Login(username string, password string) (int, error) {
	var passwordHash []byte
	var salt []byte

	user, err := service.UserRepo.GetUserFromName(username)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: User %v does not exist", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	salt = []byte(user.Salt)
	passwordHash = []byte(user.PwHash)

	inputHash := hashPw([]byte(password), salt)
	if bytes.Compare(inputHash, passwordHash) != 0 {
		errorMessage := fmt.Sprintf("Error: Invalid username and password combination for user %v", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	sessionID, err := service.SessionRepo.CreateSession(user.ID)
	if err != nil {
		session, _ := service.SessionRepo.GetSession(user.ID)
		sessionID = session.ID
	}

	service.Logger.Printf("Successfully logged in user %v\n", username)

	return sessionID, nil
}

func (service UserService) Logout(sessionID int) error {

	err := service.SessionRepo.DeleteSession(sessionID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Failed to end session with id %v", sessionID)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return errors.New(errorMessage)
	}

	service.Logger.Printf("Successfully logged out user with session %v\n", sessionID)

	return nil
}

func (service UserService) Register(username string, password string) (int, error) {
	salt := generateSalt(saltLength)

	passwordHash := hashPw([]byte(password), salt)

	_, err := service.UserRepo.GetUserFromName(username)
	if err == nil {
		errorMessage := fmt.Sprintf("Error: username %v is already taken", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	userID, err := service.UserRepo.Register(username, string(passwordHash), string(salt))
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Failed to register user %v", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	service.Logger.Printf("Successfully registered user %v\n", username)

	sessionID, err := service.SessionRepo.CreateSession(userID)
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Failed to create session for user %v", username)
		service.Logger.Println(errorMessage)
		service.Logger.Println(err)

		return 0, errors.New(errorMessage)
	}

	return 0, errors.New(strconv.Itoa(sessionID))
}

func (service UserService) CreateSession(userID int) (domain.Session, error) {

	session, err := service.SessionRepo.GetSession(userID)
	if err != nil {
		return domain.Session{}, nil
	}

	service.Logger.Printf("Successfully created new session for user with id %v\n", userID)

	return session, nil
}

func (service *UserService) GetSession(userID int) (domain.Session, error) {
	return service.SessionRepo.GetSession(userID)
}
