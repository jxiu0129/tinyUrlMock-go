package edb

/*
const (
	// NoticeWait wait
	NoticeWait NoticeStatus = iota
	// NoticeRunning running
	NoticeRunning
	// NoticeDone done
	NoticeDone
) */

// UnusedKeys
type UsedKeys struct {
	UniqueKey string `gorm:"column:UniqueKey;primary_key"`
}

// TableName sets the insert table name for this struct type
func (n *UsedKeys) TableName() string {
	return "UsedKeys"
}
