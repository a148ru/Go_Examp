package err

type MyError string

func (e MyError) Error() string {
	return string(e)
}
