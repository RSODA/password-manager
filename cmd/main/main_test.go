package main

import (
	"testing"
	"time"
)

func TestIsAccessAllowed(t *testing.T) {
	tests := []struct {
		name string
		hour int
		want bool
	}{
		{name: "before window", hour: 8, want: false},
		{name: "window start", hour: 9, want: true},
		{name: "inside window", hour: 13, want: true},
		{name: "window end", hour: 18, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Date(2026, 6, 14, tt.hour, 0, 0, 0, time.Local)
			if got := isAccessAllowed(now); got != tt.want {
				t.Fatalf("isAccessAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}
