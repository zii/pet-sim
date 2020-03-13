package base

func Raise(err error) {
	if err != nil {
		panic(err)
	}
}
