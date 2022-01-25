package repository

import "gorm.io/gorm"

type profileRepository struct {
	conn *gorm.DB
}
