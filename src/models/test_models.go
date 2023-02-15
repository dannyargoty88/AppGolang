package models

type Pokemon struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
