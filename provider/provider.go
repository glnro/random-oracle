package provider

import (
	"context"
	"encoding/hex"
	"github.com/drand/drand/client"
	"github.com/drand/drand/client/http"
)

func NewRandomnessProvider(urls []string, chainHash string) *RandProvider {
	return &RandProvider{
		urls,
		chainHash,
	}
}

func (p *RandProvider) GetLatestRandomness() (*LatestRandomRound, error) {
	var chainHash, _ = hex.DecodeString(p.chainHash)

	c, err := client.New(
		client.From(http.ForURLs(p.clientUrls, chainHash)...),
		client.WithChainHash(chainHash),
	)

	if err != nil {
		return nil, err
	}

	randRes, err := c.Get(context.Background(), 0)
	if err != nil {
		return nil, err
	}

	return &LatestRandomRound{
		Round:      randRes.Round(),
		Randomness: string(randRes.Randomness()),
		Signature:  string(randRes.Signature()),
	}, nil
}
