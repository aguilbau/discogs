package discogs

type LabelReleases struct {
	Pagination Pagination     `json:"pagination"`
	Releases   []LabelRelease `json:"releases"`
}

type LabelRelease struct {
	Artist      string `json:"artist"`
	Catno       string `json:"catno"`
	Format      string `json:"format"`
	ID          int    `json:"id"`
	ResourceURL string `json:"resource_url"`
	Status      string `json:"status"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Year        int    `json:"year"`
}
