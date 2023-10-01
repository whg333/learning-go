package check

import "log"

func CheckAndLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
