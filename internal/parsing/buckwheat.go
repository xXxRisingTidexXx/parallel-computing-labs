package parsing

type Buckwheat struct {
	URL      string  `json:"url"`
	ImageURL string  `json:"imageURL"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Mass     float64 `json:"mass"`
}
