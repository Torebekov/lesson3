package runeindex

func RuneByIndex(s *string, i *int) (r rune, err error) {
	defer func() {
		if r,ok := recover().(error); ok{
			err = r
		}
	}()
	return rune((*s)[*i]), nil
}
