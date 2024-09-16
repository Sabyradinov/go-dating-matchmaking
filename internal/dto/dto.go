package dto

import (
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/entity"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/model"
)

func ToUserData(users []entity.User) (res []model.UserData) {
	for _, user := range users {
		res = append(res, model.UserData{
			UserID:    user.UserID,
			Name:      user.Name,
			Age:       user.Age,
			Gender:    user.Gender,
			Interests: user.Interests,
		})
	}

	return res
}
