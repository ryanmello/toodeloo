package store

import (
	"context"
	"database/sql"
)

type Follower struct {
	UserId     int64  `json:"user_id"`
	FollowerId int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowerStore struct {
	sb *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followerId int64, userId int64) error {
	return nil
}

func (s *FollowerStore) Unfollow(ctx context.Context, followerId int64, userId int64) error {
	return nil
}
