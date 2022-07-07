package dto

type PanelGroupDataOutPut struct {
	TaskNum       int `json:"task_num"`
	HistoryNum    int `json:"history_num"`
	RunningProNum int `json:"running_pro_num"`
	HostNum       int `json:"host_num"`
}

type DashServiceStatItemOutput struct {
	HostID int    `json:"host_id"`
	Name   string `json:"name"`
	Value  int64  `json:"value"`
}

type DashServiceStatOutput struct {
	Legend []string                    `json:"legend"`
	Data   []DashServiceStatItemOutput `json:"data"`
}
