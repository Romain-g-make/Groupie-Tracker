package fonction_go

func Pagination(nb int, artistes []Artist, page int) []Artist {
    if nb <= 0 || page < 0 || len(artistes) == 0 {
        return artistes
    }
    var result []Artist
	for i := 0; i < nb; i++ {
		result = append(result,artistes[i+(page*nb)])
	}
	return result
}

func GetTotalPages(nb int, artistes []Artist) int {
    if nb <= 0 {
        return 1
    }
    total := len(artistes) / nb
    if len(artistes) % nb != 0 {
        total++
    }
    return total
}