package piscine

func StrRev(s string) string {
	var nur string
	for _, g := range s {
		nur = string(g) + nur
	}
	return nur
}
