package data

import (
	"errors"
	"github.com/PepperTalk/api/database/model"
	"github.com/google/uuid"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/database"
)

func ExtractUUIDFromRequest(request *goyave.Request, field string) (uuid.UUID, error) {
	rawId, found := request.Params[field]
	if !found {
		return uuid.UUID{}, errors.New("no " + field + " provided")
	}
	return uuid.Parse(rawId)
}

func LoadUserById(uuid string) (*model.User, bool) {
	user := model.User{}
	res := database.Conn().First(&model.User{}, "id = ?", uuid).Scan(&user)
	_ = res
	return &user, true
}

func ExtractUserFromRequest(request *goyave.Request, idField string) (*model.User, error) {
	extractedUuid, err := ExtractUUIDFromRequest(request, idField)
	if err != nil {
		return nil, err
	}
	user := model.User{}
	database.Conn().First(&model.User{}, "id = ?", extractedUuid.String()).Scan(&user)
	return &user, nil
}
