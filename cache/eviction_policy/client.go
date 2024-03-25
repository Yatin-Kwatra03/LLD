package eviction_policy

import "errors"

// GetEvictionPolicy : decides the eviction policy based
// on some parameters. For our implementation we have taken
// just one variable argument, actual logic can be very complex
func GetEvictionPolicy(arg string) (IEviction, error) {
	evictionClient := newEviction()

	switch arg {
	// - we need to have the latest vendor data stored on priority
	//   and might need to evict the old data if required
	// - lru caching is a suitable option for this use case
	case "cache vendor data":
		return evictionClient.lru, nil
	default:
		return nil, errors.New("useCase not supported")
	}
}
