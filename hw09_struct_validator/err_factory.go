package hw09structvalidator

func getTypeError(valueType string, conditionType string) error {
	switch valueType + conditionType {
	case "intin":
		return ErrIntIn
	case "intmin":
		return ErrIntMin
	case "intmax":
		return ErrIntMax
	case "stringin":
		return ErrStrIn
	case "stringlen":
		return ErrStrLen
	case "stringregexp":
		return ErrStrReg
	}
	return nil
}
