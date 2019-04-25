package models

type LSFJobReq struct {
    JobID       string     `gorm:"primary_key"`
    JobName     string  `gorm:"type:varchar(256);not null;"`
    WorkDir     string  `gorm:"type:varchar(1024);not null;"`
    Command     string  `gorm:"type:varchar(1024);not null;"`
//    ProjectName string  `gorm:"type:varchar(256);not null;"`
//    JobQueue    string  `gorm:"type:varchar(256);not null;"`
//    UserGroup   string  `gorm:"type:varchar(256);not null;"`
}
