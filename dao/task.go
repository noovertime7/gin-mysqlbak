package dao

type TaskDetail struct {
	Info *TaskInfo     `json:"task_info"`
	Ding *DingDatabase `json:"ding"`
	Oss  *OssDatabase  `json:"oss"`
}
