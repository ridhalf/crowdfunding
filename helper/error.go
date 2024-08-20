package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
func ResultOrError[T any](result T, err error) (T, error) {
	if err != nil {
		return result, err
	}
	return result, nil
}
