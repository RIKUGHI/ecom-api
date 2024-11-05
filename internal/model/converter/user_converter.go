package converter

import (
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Token:     user.Token,
	}
}
