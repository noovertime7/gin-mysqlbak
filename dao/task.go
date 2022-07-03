package dao

type TaskDetail struct {
	Host *HostDatabase `json:"host_info"`
	Info *TaskInfo     `json:"task_info"`
	Ding *DingDatabase `json:"ding"`
	Oss  *OssDatabase  `json:"oss"`
}
