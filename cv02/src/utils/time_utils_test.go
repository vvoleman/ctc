package utils

import (
	"testing"
	"time"
)

func TestGetRandomDuration(t *testing.T) {
	testCases := []struct {
		lowerBound float64
		upperBound float64
	}{
		{1.0, 10.0},
		{0.0, 5.0},
		{10.0, 100.0},
	}

	for _, tc := range testCases {
		duration := GetRandomDuration(tc.lowerBound, tc.upperBound)

		if duration < time.Duration(tc.lowerBound)*time.Second || duration > time.Duration(tc.upperBound)*time.Second {
			t.Errorf("Expected duration to be between %f and %f seconds, but got %f seconds", tc.lowerBound, tc.upperBound, duration.Seconds())
		}
	}
}
