package utils

func checkErr(err error) {
	if err != nil {
		err.Error()
	}
}
