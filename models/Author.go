package models 

import "gorm.io/gorm"

type Author struct {
	gorm.Model 
	name string 
	country string 
}