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

type InitialiseResponse struct {
	Response
	Result InitialiseResult `json:"result"`
}

type InitialiseResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct{}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitialiseResponse(id int) InitialiseResponse {
	return InitialiseResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitialiseResult{
			Capabilities: ServerCapabilities{},
			ServerInfo:   ServerInfo{},
		},
	}
}
