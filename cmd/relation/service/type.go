package service

import "context"

const (
	FOLLOW    int32 = 1
	UNFOLLOW  int32 = 2
	FOLLOWING int32 = 1
	FOLLOWER  int32 = 2
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}
