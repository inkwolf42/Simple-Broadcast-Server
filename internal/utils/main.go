package utils

import "os"

func ExitIfErr(err error) {
	if err != nil {
		print("Error : ")
		println(err)
		os.Exit(1)
	}
}
func ExitIfErrCM(err error, msg string) {
	if err != nil {
		println(msg, err)
		os.Exit(1)
	}
}
