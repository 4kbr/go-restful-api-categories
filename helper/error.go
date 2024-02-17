package helper

// untuk memanggil panic jika error
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
