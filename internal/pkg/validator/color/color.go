package color

type HEXColor struct {
	hex string
}

func validateHex(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}
func Validate(s string) bool {
	return s[0] == '#' && len(s) == 7 && validateHex(s[1]) && validateHex(s[2]) && validateHex(s[3]) && validateHex(s[4]) && validateHex(s[5]) && validateHex(s[6])
}
