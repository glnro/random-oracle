package provider

import "context"

type RandomnessProvider interface {
	GetLatestRandomness(ctx context.Context) (*LatestRandomRound, error)
}
