package fonction_go

import (
	"net/http"
	"strconv"
)

type queryParams struct {
	pagination  int
	page        int
	searchQuery string
	searchType  string
	artistName  string
	reset       string
	sortOrder   string
}

func extractQueryParams(r *http.Request) queryParams {
	params := queryParams{
		searchQuery: r.URL.Query().Get("Search"),
		searchType:  r.URL.Query().Get("SearchType"),
		artistName:  r.URL.Query().Get("Is_artist"),
		reset:       r.URL.Query().Get("Reset"),
		sortOrder:   r.URL.Query().Get("sort"),
	}

	if p := r.URL.Query().Get("pagination"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			params.pagination = val
		}
	}

	if pa := r.URL.Query().Get("page"); pa != "" {
		if val, err := strconv.Atoi(pa); err == nil && val >= 0 {
			params.page = val
		}
	}

	return params
}
