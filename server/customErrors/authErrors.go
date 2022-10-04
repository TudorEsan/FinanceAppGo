package customErrors

type Error interface {
	Error() string
}

type ExpiredToken struct {
}

func (e ExpiredToken) Error() string {
	return "Token was Expired"
}

type InvalidToken struct {
	E error
}

func (e InvalidToken) Error() string {
	return e.E.Error()
}

func GetJsonError(e Error) string {
	return e.Error()
}
