package core

func Is(err any) bool {
	return err != nil
}
func Some[T comparable](data T) bool {
	var temp T
	return data != temp
}
