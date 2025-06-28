package ping

import (
	"net/http"
	"time"

	"github.com/kkvaleriy/pingRobot/internal/usecase/client"
)

var MaxRetries = 5
var RetriesInterval = time.Second * 3

func Do(serviceEndpoint, serviceName string) {

	client.Set(client.Dto{
		Name:        serviceName,
		Status:      getWithRetry(serviceEndpoint),
		TimeOfCheck: time.Now(),
	})

}

func getWithRetry(serviceEndpoint string) string {
	for attempt := 0; attempt < MaxRetries; attempt++ {
		resp, err := http.Get(serviceEndpoint)
		defer resp.Body.Close()
		if err != nil {
			<-time.After(RetriesInterval)
			continue
		}

		return resp.Status
	}

	return "can't connect to service"
}
