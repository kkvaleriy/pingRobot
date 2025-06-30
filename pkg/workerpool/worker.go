package workerpool

type Job struct {
	ServiceEndpoint string
	ServiceName     string
}

type worker struct {
	f func(serviceEndpoint, serviceName string)
}