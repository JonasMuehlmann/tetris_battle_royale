package userService

import (
	"bytes"
	"errors"
	"log"
	"microservice/internal/domain"
	drivenPorts "microservice/internal/driven_ports"
	"strconv"
)

type UserService struct {
	UserRepo    drivenPorts.UserPort
	SessionRepo drivenPorts.SessionPort
}

func (service UserService) IsLoggedIn(username string) (int, error) {
	user, err := service.UserRepo.GetUserFromName(username)
	if err != nil {
		return 0, errors.New("User does not exist")
	}

	session, err := service.SessionRepo.GetSession(user.ID)
	if err != nil {
		return 0, errors.New("User not logged in")
	}

	// TODO: Check if session expired
	return session.ID, nil
}

func (service UserService) Login(username string, password string) (int, error) {
	var passwordHash []byte
	var salt []byte

	user, err := service.UserRepo.GetUserFromName(username)
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, errors.New("User does not exist")
	}

	salt = []byte(user.Salt)
	passwordHash = []byte(user.PwHash)

	inputHash := hashPw([]byte(password), salt)
	if bytes.Compare(inputHash, passwordHash) != 0 {
		return 0, errors.New("InvalID username or password")
	}

	sessionID, err := service.SessionRepo.CreateSession(user.ID)
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, errors.New("User already logged in")
	}

	return sessionID, nil
}

func (service UserService) Logout(sessionID int) error {

	err := service.SessionRepo.DeleteSession(sessionID)
	if err != nil {
		log.Printf("Error: %v", err)
		return errors.New("Failed to end session")
	}

	return nil
}

func (service UserService) Register(username string, password string) (int, error) {
	salt := generateSalt(saltLength)

	passwordHash := hashPw([]byte(password), salt)

	_, err := service.UserRepo.GetUserFromName(username)
	if err == nil {
		log.Printf("Error: %v", err)
		return 0, errors.New("Username is already in use")
	}

	log.Println("Created new password salt")

	userID, err := service.UserRepo.Register(username, string(passwordHash), string(salt))
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, errors.New("Failed to create account")
	}

	log.Printf("Created new user")

	sessionID, err := service.SessionRepo.CreateSession(userID)
	if err != nil {
		log.Printf("Error: %v", err)
		return 0, errors.New("Failed to create account")
	}

	return 0, errors.New(strconv.Itoa(sessionID))
}

func (service UserService) CreateSession(userID int) (domain.Session, error) {

	session, err := service.SessionRepo.GetSession(userID)
	if err != nil {
		return domain.Session{}, nil
	}

	log.Println("Created new session")

	return session, nil
}

func (service *UserService) GetSession(userID int) (domain.Session, error) {
	return service.SessionRepo.GetSession(userID)
}
