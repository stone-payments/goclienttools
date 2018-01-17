package http

//RequestMethod define function signature to make http request
type RequestMethod func(url string, request RequestOptions) (Response, error)
