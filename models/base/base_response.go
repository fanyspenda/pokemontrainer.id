package base

//Response used as response format
type Response struct {
	Code    int
	Status  string
	Message string
	Data    interface{}
}
