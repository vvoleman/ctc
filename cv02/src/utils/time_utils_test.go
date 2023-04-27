package utils

import (
	"testing"
	"time"
)

func TestGetRandomDuration(t *testing.T) {
	type args struct {
		lowerBound float64
		upperBound float64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomDuration(tt.args.lowerBound, tt.args.upperBound); got != tt.want {
				t.Errorf("GetRandomDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}