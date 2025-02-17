package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 00.0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result is ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0") // string error thì ko nên viết hoa
	}
	result = x / y
	return result, nil
}
