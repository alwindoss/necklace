package chain

type Block struct {
	id      string `json:"id"`
	data    string `json:"data"`
	key     string `json:"key"`
	prevKey string `json:"prev_key"`
}
