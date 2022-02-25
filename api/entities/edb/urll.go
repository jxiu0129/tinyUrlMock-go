package edb

import "time"

/*
const (
	// NoticeWait wait
	NoticeWait NoticeStatus = iota
	// NoticeRunning running
	NoticeRunning
	// NoticeDone done
	NoticeDone
) */

// Url
type Url struct {
	ID          uint64    `gorm:"column:ID;primary_key"`
	ShortenUrl  string    `gorm:"column:ShortenUrl"`
	OriginalUrl string    `gorm:"column:OriginalUrl"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
}

// TableName sets the insert table name for this struct type
func (n *Url) TableName() string {
	return "Url"
}
