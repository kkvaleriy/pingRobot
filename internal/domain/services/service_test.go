package services

import (
	"reflect"
	"testing"
	"time"
)

func TestService_Status(t *testing.T) {
	tests := []struct {
		name string
		s    *Service
		want string
	}{
		{
			name: "check correct status",
			s: &Service{
				status:    "200 OK",
				lastCheck: time.Now(),
			},
			want: "200 OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Status(); got != tt.want {
				t.Errorf("Service.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_TimeOfLastCheck(t *testing.T) {
	timeOfTest := time.Now()
	tests := []struct {
		name string
		s    *Service
		want time.Time
	}{
		{
			name: "check correct lastCheck time",
			s: &Service{
				status:    "200 OK",
				lastCheck: timeOfTest,
			},
			want: timeOfTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.TimeOfLastCheck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.TimeOfLastCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
