package repository

import (
    "time"
    "gorm.io/gorm"
	"github.com/gavril-s/erpico-test-assignemnt/models"
)

type Repository interface {
	CreateOrganization(org *models.Organization) error
	GetOrganizationByID(id uint) (*models.Organization, error)
	UpdateOrganization(org *models.Organization) error
	DeleteOrganization(id uint) error

	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error

	CreateActivity(activity *models.Activity) error
	GetActivityByID(id uint) (*models.Activity, error)
	UpdateActivity(activity *models.Activity) error
	DeleteActivity(id uint) error

	CreateSchedule(schedule *models.Schedule) error
	GetScheduleByID(id uint) (*models.Schedule, error)
    GetScheduleForDate(gymID string, date time.Time) (*models.Schedule, error)
	GetScheduledUsers(gymID string, scheduleID string) ([]*models.User, error)
	UpdateSchedule(schedule *models.Schedule) error
	DeleteSchedule(id uint) error

	CreateRecord(record *models.Record) error
	GetRecordByID(id uint) (*models.Record, error)
	UpdateRecord(record *models.Record) error
	DeleteRecord(id uint) error
}

type GormRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{DB: db}
}

func (r *GormRepository) CreateOrganization(org *models.Organization) error {
	return r.DB.Create(org).Error
}

func (r *GormRepository) GetOrganizationByID(id uint) (*models.Organization, error) {
	org := &models.Organization{}
	if err := r.DB.First(org, id).Error; err != nil {
		return nil, err
	}
	return org, nil
}

func (r *GormRepository) UpdateOrganization(org *models.Organization) error {
	return r.DB.Save(org).Error
}

func (r *GormRepository) DeleteOrganization(id uint) error {
	return r.DB.Delete(&models.Organization{}, id).Error
}

func (r *GormRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *GormRepository) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	if err := r.DB.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *GormRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *GormRepository) DeleteUser(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}

func (r *GormRepository) CreateActivity(activity *models.Activity) error {
	return r.DB.Create(activity).Error
}

func (r *GormRepository) GetActivityByID(id uint) (*models.Activity, error) {
	activity := &models.Activity{}
	if err := r.DB.First(activity, id).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

func (r *GormRepository) UpdateActivity(activity *models.Activity) error {
	return r.DB.Save(activity).Error
}

func (r *GormRepository) DeleteActivity(id uint) error {
	return r.DB.Delete(&models.Activity{}, id).Error
}

func (r *GormRepository) CreateSchedule(schedule *models.Schedule) error {
	return r.DB.Create(schedule).Error
}

func (r *GormRepository) GetScheduleByID(id uint) (*models.Schedule, error) {
	schedule := &models.Schedule{}
	if err := r.DB.First(schedule, id).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}

func (r *GormRepository) GetScheduleForDate(gymID string, date time.Time) (*models.Schedule, error) {
    var schedule models.Schedule
    if err := r.DB.Table("schedule").
        Joins("JOIN schedule_date_activation_deactivation ON schedule.id = schedule_date_activation_deactivation.schedule_id").
        Where("schedule.org_id = ? AND schedule_date_activation_deactivation.date_activate <= ? AND schedule_date_activation_deactivation.date_deactivate > ?", gymID, date, date).
        First(&schedule).Error; err != nil {
        return nil, err
    }
    return &schedule, nil
}

func (r *GormRepository) GetScheduledUsers(gymID string, scheduleID string) ([]*models.User, error) {
    var users []*models.User
    if err := r.DB.Table("record").
        Select("user.id, user.name, user.last_name, user.phone, user.email, user.creator_id, user.trainer_id, user.number, user.pin").
        Joins("JOIN schedule ON record.schedule_id = schedule.id").
        Joins("JOIN `user` ON record.user_id = `user`.id").
        Where("schedule.org_id = ? AND schedule.id = ?", gymID, scheduleID).
        Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *GormRepository) UpdateSchedule(schedule *models.Schedule) error {
	return r.DB.Save(schedule).Error
}

func (r *GormRepository) DeleteSchedule(id uint) error {
	return r.DB.Delete(&models.Schedule{}, id).Error
}

func (r *GormRepository) CreateRecord(record *models.Record) error {
	return r.DB.Create(record).Error
}

func (r *GormRepository) GetRecordByID(id uint) (*models.Record, error) {
	record := &models.Record{}
	if err := r.DB.First(record, id).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (r *GormRepository) UpdateRecord(record *models.Record) error {
	return r.DB.Save(record).Error
}

func (r *GormRepository) DeleteRecord(id uint) error {
	return r.DB.Delete(&models.Record{}, id).Error
}
