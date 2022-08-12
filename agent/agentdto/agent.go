package agentdto

type AgentOutPut struct {
	Total           int               `json:"total"`
	AgentOutPutItem []AgentOutPutItem `json:"agent_list"`
}

type AgentOutPutItem struct {
	ServiceID string `json:"service_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}
