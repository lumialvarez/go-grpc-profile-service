package error

type InvalidData struct {
	message string
}

func NewInvalidData(message string) InvalidData {
	return InvalidData{message: message}
}

func (err InvalidData) Error() string {
	return err.message
}
