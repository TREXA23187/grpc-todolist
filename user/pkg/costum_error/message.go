package costum_error

var MessageFlags = map[uint]string{
	Success:       "Ok",
	Error:         "Fail",
	InvalidParams: "Invalid Parameters",
}

func GetMessage(code uint) string {
	msg, ok := MessageFlags[code]
	if ok {
		return msg
	}

	return MessageFlags[Error]
}
