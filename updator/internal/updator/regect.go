package updator

import "regexp"

type naakt string

var rgNaakt = regexp.MustCompile("[A-Za-z0-9.-]+")

func (n naakt) String() string {
	return rgNaakt.FindString(string(n))
}
