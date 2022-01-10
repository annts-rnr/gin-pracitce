package input

type ItemCreateInput struct {
	Item struct {
		Title    string `json:"title" binding:"required"`
		Contents string `json:"contents" binding:"required"`
		Price    int    `json:"price" binding:"required"`
	} `json:"item"`
}
