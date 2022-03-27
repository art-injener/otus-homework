package response

type AccountResponse struct {
	ID      uint64 `json:"id"`
	LoginID uint64 `json:"-"`
}
