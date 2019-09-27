package discogs

import "fmt"

type Pagination struct {
	Page    uint `json:"page"`
	Pages   uint `json:"pages"`
	Items   uint `json:"items"`
	PerPage uint `json:"per_page"`
	URLs    URLs `json:"urls"`
}

type URLs struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type PaginationParams struct {
	Page    uint
	PerPage uint
}

func (p *PaginationParams) toQuery() string {
	return fmt.Sprintf(
		"page=%d&per_page=%d",
		p.Page,
		p.PerPage)
}
