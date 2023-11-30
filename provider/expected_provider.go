package provider

type RandomnessProvider interface {
	GetLatestRandomness() (*LatestRandomRound, error)
}
