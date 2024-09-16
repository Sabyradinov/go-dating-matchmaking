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
	if err != nil {
		return
	}

	var preference entity.Preference
	err = r.db.GetFirst(ctx, &preference, "user_id = ?", userID)
	if err != nil {
		return
	}

	res.Preference = &preference

	return
}

// GetPotentialMatches retrieves potential matches for a given user ID
func (r *user) GetPotentialMatches(ctx context.Context, currentUser entity.User, offset, limit int) (potentialMatches []entity.User, err error) {

	sql := `SELECT u.* FROM public.user u
         inner join public.preference p on u.user_id = p.user_id
         WHERE u.gender = ? AND u.age BETWEEN ? AND ? AND ST_DWithin(u.location, ?, ?)
         AND p.gender = ? AND ? BETWEEN p.age_range_min AND p.age_range_max
         AND p.max_distance >= ST_Distance(u.location, ?)`
	values := []interface{}{currentUser.Preference.Gender, currentUser.Preference.AgeRangeMin, currentUser.Preference.AgeRangeMax, currentUser.Location, currentUser.Preference.MaxDistance,
		currentUser.Gender, currentUser.Age, currentUser.Location}

	err = r.db.GetByScript(ctx, &potentialMatches, limit, offset, sql, values)

	return
}
