package dto

type PanelGroupDataOutPut struct {
	TaskNum       int `json:"task_num"`
	HistoryNum    int `json:"history_num"`
	RunningProNum int `json:"running_pro_num"`
	StopProNum    int `json:"stop_pro_num"`
}
