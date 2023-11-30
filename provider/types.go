package provider

type LatestRandomRound struct {
	Round      uint64
	Randomness []byte
	Signature  []byte
}

type RandProvider struct {
	clientUrls []string
	chainHash  string
}
