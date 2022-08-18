package agentdto

type BarChartOutPut struct {
	ServiceName []string `json:"service_name"`
	TaskNum     []int64  `json:"task_num"`
}
