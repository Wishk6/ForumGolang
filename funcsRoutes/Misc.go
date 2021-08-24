package funcsRoutes

import "math/rand"

var idList []string

func GenerateNumber() (Id string) {
retry:
	var charKind int
	var NextChar rune
	for len(Id) < 20 {
		charKind = rand.Intn(3)
		switch charKind {
		case 0: // Uppercase letter
			NextChar = rune(rand.Intn(26) + 65)
			break
		case 1: // Number
			NextChar = rune(rand.Intn(26) + 48)
			break
		case 2: // LowerCase Letter
			NextChar = rune(rand.Intn(26) + 97)
			break
		}
		Id += string(NextChar)
	}

	if isValid(Id) == false {
		goto retry
	}

	return Id
}

func isValid(Id string) bool {
	CreateDb([]string{"create table ids(id text unique)"}, "sql/ids.db")
	_, exist := IdExist(Id)

	return !exist
}
