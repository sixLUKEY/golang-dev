package lsp

type InitialiseRequest struct {
	Request
	Params InitialiseRequestParams `json:"params"`
}

type InitialiseRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... lots more can go here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
