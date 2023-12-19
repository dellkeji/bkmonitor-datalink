// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/config"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

// DBSession session of databases
type DBSession struct {
	DB *gorm.DB
}

var dbSession *DBSession

var once sync.Once

func GetDBSession() *DBSession {
	if dbSession != nil {
		return dbSession
	}
	once.Do(func() {
		dbSession = &DBSession{}
		err := dbSession.Open()
		if err != nil {
			logger.Errorf("connection mysql error, %v", err)
			panic(err)
		}
	})
	return dbSession
}

// Open connect the mysql db
func (db *DBSession) Open() error {
	var err error
	mysqlConfig := config.GlobalConfig.Store.MysqlConfig
	dbhost := fmt.Sprintf("tcp(%s:%d)", mysqlConfig.Host, mysqlConfig.Port)
	db.DB, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		dbhost,
		mysqlConfig.DbName,
		mysqlConfig.Charset,
	))
	if err != nil {
		logger.Errorf("new a mysql connection error, %v", err)
		return err
	}
	sqldb := db.DB.DB()
	sqldb.SetMaxIdleConns(mysqlConfig.MaxIdleConnections)
	sqldb.SetMaxOpenConns(mysqlConfig.MaxOpenConnections)

	// 判断连通性
	if err := sqldb.Ping(); err != nil {
		return err
	}

	// 是否开启 debug 模式
	if mysqlConfig.Debug {
		logger.Info("Debug mode of mysql is enabled. Check if it is in the test environment!")
		db.DB.LogMode(true)
	}

	return nil
}

// Close closes connection
func (db *DBSession) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
