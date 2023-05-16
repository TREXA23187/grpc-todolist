package handler

import "errors"

func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("user service--" + err.Error())
		panic(err)
	}

}
