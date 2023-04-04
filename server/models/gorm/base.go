/**
 * @Author: smono
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2022/9/24 14:36
 */

package gorm

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        uint32 `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
