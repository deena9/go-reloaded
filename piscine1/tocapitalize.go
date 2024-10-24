package piscine

func ToCapitalize(s string) string {
	result := ""
	former := ' '
	for i, latter := range s {

		if latter >= 'a' && latter <= 'z' && (i == 0 || (!(former >= 'a' && former <= 'z') && !(former >= 'A' && former <= 'Z') || (former >= '0' && former <= '9'))) {
			latter -= 32
		}
		if (latter >= 'A' && latter <= 'Z') && ((former >= 'a' && former <= 'z') || (former >= 'A' && former <= 'Z') || (former >= '0' && former <= '9')) {
			latter += 32
		}
		former = latter
		result = result + string(latter)
	}
	return result
}
