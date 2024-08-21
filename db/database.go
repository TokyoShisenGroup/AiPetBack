package db

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
	dbErr      error
)

// GetDatabaseInstance 获取数据库实例 (SQLite)
func GetDatabaseInstance() (*gorm.DB, error) {
	dbOnce.Do(func() {
		// 使用 SQLite 打开数据库连接
		dbInstance, dbErr = gorm.Open(sqlite.Open("databasd.db"), &gorm.Config{})
	})
	return dbInstance, dbErr
}
