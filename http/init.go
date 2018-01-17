package http

var (
	//Requester is the http requester
	Requester Requestable
	//Interceptors holds external request interceptors
	Interceptors []Interceptor
)

func init() {
	Requester = new(requester)
}
