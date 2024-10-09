package err

func Gen_error() string {

	var er MyError = "Error!!!"
	return er.Error()
}
