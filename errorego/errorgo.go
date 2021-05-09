package errorego

import "errors"
import "fmt"

func BasicErrorHandling(number int16) (int16, error) {

	if number == 4 {
		return -1, errors.New("found 4 Replace with Other than 4")
	}
	return number, nil
}

func RaiseError(arg int8) (int8, error) {

	if arg == 5 {
		return -1, &ArgError{arg, "DSA1234", "Template Not Found"}
	}
	return 6, nil
}

type ArgError struct {
	Arg     int8
	DsaCode string
	Reason  string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s => %s", e.Arg, e.DsaCode, e.Reason)
}
