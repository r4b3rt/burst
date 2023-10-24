package api

type PortMapping struct {
	ClientPort int64  `json:"clientPort"`
	ServerPort int64  `json:"serverPort"`
	Protocol   string `json:"protocol"`
}

type ExportRequest struct {
	ClientWsPort int64          `json:"clientWsPort"`
	PortMapping  []*PortMapping `json:"portMapping"`
}

type PortMappingResp struct {
	Mapping      *PortMapping `json:"mapping"`
	Error        string       `json:"error"`
	ConnectionId string       `json:"connectionId"`
}

type ExportResponse struct {
	Items []*PortMappingResp `json:"items"`
}

type TransFromData struct {
	ConnectionId     string `json:"connectionId"`
	UserConnectionId string `json:"userConnectionId"`
	Data             []byte `json:"data"`
}
