package utils

import "errors"

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func UserNotFoundError() error {
	err := errors.New("user not found")
	return err
}

func InvalidCredentialsError() error {
	err := errors.New("credentials are invalid")
	return err
}
