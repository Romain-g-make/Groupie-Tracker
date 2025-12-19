package fonction_go

func Get_alphabet(artistes []Artist) []string {
	var alphabet []string
	for i := 0; i < len(artistes); i++ {
		first_letter := string(artistes[i].Nom[0])
		found := false
		for j := 0; j < len(alphabet); j++ {
			if alphabet[j] == first_letter {
				found = true
				j = len(alphabet)
			}
		}
		if !found {
			alphabet = append(alphabet, first_letter)
		}
	}
	return alphabet
}