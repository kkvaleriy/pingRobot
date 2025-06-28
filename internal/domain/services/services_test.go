package services

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	timeOfTest := time.Now()
	type args struct {
		serviceName string
		status      string
		timeOfCheck time.Time
	}
	tests := []struct {
		name string
		args args
		want map[string]Service
	}{
		{
			name: "init services test",
			args: args{
				serviceName: "testServiceName",
				status:      "200 OK",
				timeOfCheck: timeOfTest,
			},
			want: map[string]Service{"testServiceName": Service{status: "200 OK", lastCheck: timeOfTest}},
		},
		{
			name: "add services test",
			args: args{
				serviceName: "secondTestServiceName",
				status:      "200 OK",
				timeOfCheck: timeOfTest,
			},
			want: map[string]Service{"testServiceName": {status: "200 OK", lastCheck: timeOfTest},
				"secondTestServiceName": {status: "200 OK", lastCheck: timeOfTest}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.serviceName, tt.args.status, tt.args.timeOfCheck)

			if !reflect.DeepEqual(s.store, tt.want) {
				t.Errorf("services.Set() = %v, want %v", s.store, tt.want)
			}
		})
	}
}

func TestStatus(t *testing.T) {
	timeOfTest := time.Now()
	tests := []struct {
		name string
		want map[string]Service
	}{
		{name: "get status",
			want: map[string]Service{"testServiceName": {status: "200 OK", lastCheck: timeOfTest},
				"secondTestServiceName": {status: "200 OK", lastCheck: timeOfTest}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.mu = &sync.Mutex{}
			s.store = tt.want
			if got := Status(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}
