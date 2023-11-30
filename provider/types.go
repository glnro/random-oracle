package provider

type LatestRandomRound struct {
	Round      uint64
	Randomness string
	Signature  string
}

type RandProvider struct {
	clientUrls []string
	chainHash  string
}
