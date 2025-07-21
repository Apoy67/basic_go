package halper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
