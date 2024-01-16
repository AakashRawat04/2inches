package types

type Link struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	LongUrl   string `json:"long_url"`
}
