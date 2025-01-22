package responses

type CreateResponse struct {
	Id *string `json:"id"`
}

type StatusResponse struct {
	Status bool `json:"status"`
}
