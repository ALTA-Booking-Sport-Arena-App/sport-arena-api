package helper

import "net/mail"

func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
