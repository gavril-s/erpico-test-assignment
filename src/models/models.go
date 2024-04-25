package models

import "time"

type Organization struct {
	ID uint `gorm:"primaryKey"`
}

type User struct {
	ID         uint
	Name       string
	LastName   string
	Phone      string
	Email      string
	CreatorID  uint
	TrainerID  uint
	Number     string
	PIN        string
}

func (User) TableName() string {
    return "user"
}

type Activity struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Title       string
	ExternalID  string
}

type ActivityDetail struct {
	ID                uint `gorm:"primaryKey"`
	ActivityID        uint
	Comment           string
	Duration          int
	MaxParticipants   int
	MinParticipants   int
	Quota             int
	WaitlistActive    bool
	WaitlistLimitMins int
}

type Schedule struct {
	ID         uint `gorm:"primaryKey"`
	OrgID      uint
	ActivityID uint
	TrainerID  uint
	ExternalID string
	Name       string
	Description string
}

func (Schedule) TableName() string {
    return "schedule"
}

type ScheduleDetail struct {
	ID              uint `gorm:"primaryKey"`
	ScheduleID      uint
	StartTime       time.Time
	EndTime         time.Time
	ActivityDate    time.Time
	StartDate       time.Time
	ActivityDuration int
	CycleDay        int
}

type Record struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint
	ScheduleID        uint
	ActivityID        uint
	ActivityDate      time.Time
	StartTime         time.Time
	Comment           string
	State             string
	ParentID          uint
	OverlappedRecordID uint
}

type RecordEquipment struct {
	RecordID    uint `gorm:"primaryKey"`
	EquipmentID uint `gorm:"primaryKey"`
	Deleted     bool
}
