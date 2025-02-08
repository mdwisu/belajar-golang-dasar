package main

import "fmt"

type validationError struct {
	msg string
}

func (e *validationError) Error() string {
	return e.msg
}

type notFoundError struct {
	msg string
}

func (e *notFoundError) Error() string {
	return e.msg
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID Tidak Boleh Kosong"}
	}
	if id != "dwi" {
		return &notFoundError{"Data Tidak Ditemukan"}
	}
	// ok
	return nil
}

func main() {
	err := saveData("dwi", nil)
	if err != nil {
		// tejadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error", validationErr.Error())
		// 	fmt.Println(validationErr.msg)
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error", notFoundErr.Error())
		// 	fmt.Println(notFoundErr.msg)
		// } else {
		// 	fmt.Println(err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Error())
			fmt.Println(finalError.msg)
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
			fmt.Println(finalError.msg)
		default:
			fmt.Println(err.Error())
		}
	} else {
		// success
		fmt.Println("Success")
	}
}