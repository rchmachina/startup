package helper


func Formaterword(format string) string{
	changerword :=""
	for _, formats := range format{
		s := string(formats)
		if s ==" "{
			s ="-"
		}
		changerword += s		
	}
	return changerword
}