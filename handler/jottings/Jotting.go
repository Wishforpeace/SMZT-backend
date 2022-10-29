package jottings

type JottingRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type JottingResponse struct {
	Publisher string `json:"publisher"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Date      string `json:"date"`
}
