package user

import (
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"
	//"github.com/go-openapi/swag"

	"github.com/jaingaurav001/swagger_sample/models"
)

func Search(params GetSearchParams) middleware.Responder {
	payload := &models.User{
		UserID: &params.UserID,
		Name:   fmt.Sprintf("name_%d", params.UserID),
	}

	return NewGetSearchOK().WithPayload(payload)
}
