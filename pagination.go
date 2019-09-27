package discogs

type Pagination struct {
	Page    int  `json:"page"`
	Pages   int  `json:"pages"`
	Items   int  `json:"items"`
	PerPage int  `json:"per_page"`
	URLs    URLs `json:"urls"`
}

type URLs struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}
