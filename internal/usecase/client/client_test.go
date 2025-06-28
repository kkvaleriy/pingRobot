package client

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	timeOfTest := time.Now()
	type args struct {
		serviceInfo Dto
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "check building client use case",
			args: args{serviceInfo: Dto{
				Name:        "testName",
				Status:      "200 OK",
				TimeOfCheck: timeOfTest,
			}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.serviceInfo)
		})
	}
}
