package repo

import (
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/repository"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/storage"
)

type Builder struct {
	User repository.IUser
}

func Init(db storage.IDB) *Builder {
	return &Builder{
		User: NewUser(db),
	}
}
