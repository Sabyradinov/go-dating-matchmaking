package repo

import (
	"context"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/entity"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/repository"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/storage"
)

// user info about store goods
type user struct {
	db storage.IDB
}

// NewUser constructor to create matching repository instance
func NewUser(db storage.IDB) repository.IUser {
	return &user{db: db}
}

// GetById function to get matching by id with context
func (r *user) GetById(ctx context.Context, userID string) (res entity.User, err error) {
	err = r.db.GetFirst(ctx, &res, "user_id = ?", userID)
	return
}

// GetPotentialMatches retrieves potential matches for a given user ID
func (r *user) GetPotentialMatches(ctx context.Context, currentUser entity.User, offset, limit int) ([]entity.User, error) {

	var potentialMatches []entity.User

	sql := `SELECT * FROM users WHERE gender = ? AND age BETWEEN ? AND ? AND ST_DWithin(location, ?, ?)`
	values := []interface{}{currentUser.Preferences.Gender, currentUser.Preferences.AgeRangeMin, currentUser.Preferences.AgeRangeMax, currentUser.Location, currentUser.Preferences.MaxDistance}

	err := r.db.GetByScript(ctx, &potentialMatches, sql, values, limit, offset)
	if err != nil {
		return nil, err
	}

	return potentialMatches, nil
}
