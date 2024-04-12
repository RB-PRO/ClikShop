package bitrixupdate

import "regexp"

func naaktstring(str string) string {
	re := regexp.MustCompile("[A-Za-z0-9.-]+")
	result := re.FindString(str)
	return result
}
