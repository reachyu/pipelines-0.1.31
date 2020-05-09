// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"database/sql"
	"fmt"
)

/*
CREATE TABLE `application` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `apply_name` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '应用名称',
  `apply_type` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '应用类型',
  `apply_frame` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '预加载框架',
  `apply_environment` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '应用环境',
  `apply_brief` varchar(255) COLLATE utf8_bin NOT NULL COMMENT '应用简介',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `created_by` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

*/

type AppManageStatus string

const (
	AppCreating AppManageStatus = "CREATING"
	AppReady    AppManageStatus = "READY"
	AppDeleting AppManageStatus = "DELETING"
)

type AppManage struct {
	ID               int64          `gorm:"column:id; not null; primary_key"`
	ApplyName        string         `gorm:"column:apply_name; not null; unique"`
	ApplyType        string         `gorm:"column:apply_type; not null;"`
	ApplyFrame       string         `gorm:"column:apply_frame; not null;"`
	ApplyEnvironment string         `gorm:"column:apply_environment; not null;"`
	ApplyBrief       sql.NullString `gorm:"column:apply_brief;"`
	CreatedAt        sql.NullInt64  `gorm:"column:created_at;"`
	UpdateAt         sql.NullInt64  `gorm:"column:update_at;"`
	CreatedBy        sql.NullString `gorm:"column:created_by;"`
	UpdatedBy        sql.NullString `gorm:"column:updated_by; "`
}

func (p AppManage) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.ID)
}

func GetAppManageTablePrimaryKeyColumn() string {
	return "ID"
}

// PrimaryKeyColumnName returns the primary key for model .
func (p *AppManage) PrimaryKeyColumnName() string {
	return "ID"
}

// DefaultSortField returns the default sorting field for model .
func (p *AppManage) DefaultSortField() string {
	return "CreatedAt"
}

var AppManageAPIToModelFieldMap = map[string]string{
	"id":          "ID",
	"name":        "ApplyName",
	"created_at":  "CreatedAt",
	"description": "ApplyBrief",
}

// APIToModelFieldMap returns a map from API names to field names for model
// .
func (p *AppManage) APIToModelFieldMap() map[string]string {
	return AppManageAPIToModelFieldMap
}
