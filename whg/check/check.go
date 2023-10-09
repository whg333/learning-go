package check

import "log"

func CheckAndLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogPrintln(err error) {
	if err != nil {
		log.Println(err)
	}
}
