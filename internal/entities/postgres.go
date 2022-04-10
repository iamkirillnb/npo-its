package entities

type Item struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Metric int    `json:"metric"`
}
