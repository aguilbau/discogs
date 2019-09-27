package discogs

type Release struct {
	Artists           []Artist      `json:"artist"`
	Community         Community     `json:"community"`
	Companies         []Company     `json:"companies"`
	Country           string        `json:"country"`
	DataQuality       string        `json:"data_quality"`
	DateAdded         string        `json:"date_added"`
	DateChanged       string        `json:"date_changed"`
	EstimatedWeight   int           `json:"estimated_weight"`
	ExtraArtists      []Artist      `json:"extraartists"`
	FormatQuantity    int           `json:"format_quantity"`
	Formats           []Format      `json:"formats"`
	Genres            []string      `json:"genres"`
	ID                int64         `json:"id"`
	Identifiers       []Identifier  `json:"identifiers"`
	Images            []Image       `json:"images"`
	Labels            []Label       `json:"labels"`
	LowestPrice       float64       `json:"lowest_price"`
	MasterID          int64         `json:"master_id"`
	MasterURL         string        `json:"master_url"`
	Notes             string        `json:"notes"`
	NumForSale        int           `json:"num_for_sale"`
	Released          string        `json:"released"`
	ReleasedFormatted string        `json:"released_formatted"`
	ResourceURL       string        `json:"resource_url"`
	Series            []interface{} `json:"series"`
	Status            string        `json:"status"`
	Styles            []string      `json:"styles"`
	Thumb             string        `json:"thumb"`
	Title             string        `json:"title"`
	Tracklist         []Tracklist   `json:"tracklist"`
	URI               string        `json:"uri"`
	Videos            []Video       `json:"videos"`
	Year              int           `json:"year"`
}

type Currency string

const (
	CurrencyAUD Currency = "AUD"
	CurrencyBRL Currency = "BRL"
	CurrencyCAD Currency = "CAD"
	CurrencyCHF Currency = "CHF"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyJPY Currency = "JPY"
	CurrencyMXN Currency = "MXN"
	CurrencyNZD Currency = "NZD"
	CurrencySEK Currency = "SEK"
	CurrencyUSD Currency = "USD"
	CurrencyZAR Currency = "ZAR"
)

type Artist struct {
	Anv         string `json:"anv"`
	ID          int    `json:"id"`
	Join        string `json:"join"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
}

type Contributors struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

type Rating struct {
	Average float64 `json:"average"`
	Count   int     `json:"count"`
}

type Submitter struct {
	ResourceURL string `json:"resource_url"`
	Username    string `json:"username"`
}

type Community struct {
	Contributors []Contributors `json:"contributors"`
	DataQuality  string         `json:"data_quality"`
	Have         int            `json:"have"`
	Rating       Rating         `json:"rating"`
	Status       string         `json:"status"`
	Submitter    Submitter      `json:"submitter"`
	Want         int            `json:"want"`
}

type Company struct {
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ResourceURL    string `json:"resource_url"`
}

type Format struct {
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
}

type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Image struct {
	Height      int    `json:"height"`
	ResourceURL string `json:"resource_url"`
	Type        string `json:"type"`
	URI         string `json:"uri"`
	URI150      string `json:"uri150"`
	Width       int    `json:"width"`
}

type Label struct {
	Catno       string `json:"catno"`
	EntityType  string `json:"entity_type"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

type Tracklist struct {
	Duration string `json:"duration"`
	Position string `json:"position"`
	Title    string `json:"title"`
	Type     string `json:"type_"`
}

type Video struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
}
