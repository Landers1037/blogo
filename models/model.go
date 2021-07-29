/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package models

// Model 公共表模型
type Model struct {
	PrimaryID int `gorm:"primary_key" json:"primary_id"`
	//CreatedOn int `json:"created_on"`
	//ModifiedOn int `json:"modified_on"`
}
