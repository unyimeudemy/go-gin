package models

// import (
//     "time"
// )

import "gorm.io/gorm"

// type Post struct {
// 	gorm.Model
// 	Title string
// 	Body string
// }

type Post struct {
	gorm.Model
    Title     string    `json:"title"`
    Body      string    `json:"body"`
}