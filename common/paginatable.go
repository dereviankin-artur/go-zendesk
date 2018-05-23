package common

// Paginatable is base struct for resource pagination
type Paginatable struct {
	PreviousPage int64 `json:"previous_page"`
	NextPage     int64 `json:"next_page"`
	Count        int64 `json:"count"`
}