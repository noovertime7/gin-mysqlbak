package agentdto

type TaskOverViewListInput struct{}

type TaskOverViewOut struct {
	ID          int64  `json:"id" form:"db_name"`
	ServiceName string `json:"service_name" form:"db_name"`
	HostID      int64  `json:"host_id" form:"db_name"`
	Host        string `json:"host" form:"db_name"`
	TaskID      int64  `json:"task_id" form:"db_name"`
	DBName      string `json:"db_name" form:"db_name"`
	BackupCycle string `json:"backup_cycle" form:"backup_cycle"`
	KeepNumber  int64  `json:"keep_number" form:"db_name"`
	Status      int64  `json:"status" form:"db_name"`
}
