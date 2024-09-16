package matching

import (
	"context"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/adapter/storage/repo"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/repository"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/domain/port/service"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/dto"
	"github.com/Sabyradinov/go-dating-matchmaking/internal/model"
	"github.com/lib/pq"
	"sort"
)

type matching struct {
	userRepo repository.IUser
}

func New(repo *repo.Builder) service.IMatching {
	return &matching{userRepo: repo.User}
}

func (s *matching) GetPotentialMatches(ctx context.Context, userID string, offset, limit int) (res model.UserResponse, err error) {
	currentUser, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return
	}

	potentialMatches, err := s.userRepo.GetPotentialMatches(ctx, currentUser, offset, limit)

	// Rank users based on mutual interests
	for i := range potentialMatches {
		potentialMatches[i].Rank = calculateMutualInterests(currentUser.Interests, potentialMatches[i].Interests)
	}

	// Sort by rank
	sort.Slice(potentialMatches, func(i, j int) bool {
		return potentialMatches[i].Rank > potentialMatches[j].Rank
	})

	res.PotentialMatches = dto.ToUserData(potentialMatches)

	return
}

// calculateMutualInterests calculates the number of mutual interests between two users
func calculateMutualInterests(interests1, interests2 pq.StringArray) int {
	interestSet := make(map[string]struct{})
	for _, interest := range interests1 {
		interestSet[interest] = struct{}{}
	}

	mutualCount := 0
	for _, interest := range interests2 {
		if _, exists := interestSet[interest]; exists {
			mutualCount++
		}
	}

	return mutualCount
}

func (s *matching) GetUserById(ctx context.Context, userId string) (res interface{}, err error) {
	res, err = s.userRepo.GetById(ctx, userId)
	return
}
