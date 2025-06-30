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
	for range MaxRetries {
		resp, err := http.Get(serviceEndpoint)
		if err != nil {
			<-time.After(RetriesInterval)
			continue
		}
		defer resp.Body.Close()

		return resp.Status
	}

	return "can't connect to service"
}
