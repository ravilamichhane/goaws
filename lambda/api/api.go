package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) *ApiHandler {

	return &ApiHandler{
		dbStore: dbStore,
	}
}

func (a *ApiHandler) RegisterUserHandler(event types.RegisterUserRequest) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("username or password is empty")
	}

	userExists, err := a.dbStore.DoesUserExist(event.Username)

	if err != nil {
		return fmt.Errorf("error checking if user exists %w", err)
	}

	if userExists {
		return fmt.Errorf("user already exists")
	}

	err = a.dbStore.InsertUser(event)

	if err != nil {
		return fmt.Errorf("error registering user %w", err)
	}

	return nil

}
