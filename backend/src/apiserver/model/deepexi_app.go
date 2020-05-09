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

import "fmt"
// 20200304 yuhongbo MLOps create application table
// AppStatus a label for the status of the Pipeline.
// This is intend to make application(应用) creation and deletion atomic.
type AppStatus string

const (
	AppCreated AppStatus = "CREATED"
	AppDeleted AppStatus = "DELETED"
)

type DeepexiApp struct {
	AppId           string `gorm:"column:AppId; not null; primary_key"`
	AppName         string `gorm:"column:AppName; not null;"`
	AppDesc         string `gorm:"column:AppDesc;"`
	AppType         string `gorm:"column:AppType;"`
	AppFrame        string `gorm:"column:AppFrame;"`
	AppEnvironment  string `gorm:"column:AppEnvironment;"`
	CreatedAt       int64  `gorm:"column:CreatedAt; not null;"`
	UpdateAt        int64  `gorm:"column:UpdateAt; not null;"`
	CreatedBy       string `gorm:"column:CreatedBy;"`
	UpdatedBy       string `gorm:"column:UpdatedBy;"`
	AppStatus       AppStatus `gorm:"column:AppStatus; not null"`
}

func (p DeepexiApp) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.AppId)
}

func GetDeepexiappTablePrimaryKeyColumn() string {
	return "AppId"
}

// PrimaryKeyColumnName returns the primary key for model Pipeline.
func (p *DeepexiApp) PrimaryKeyColumnName() string {
	return "AppId"
}

// DefaultSortField returns the default sorting field for model Pipeline.
func (p *DeepexiApp) DefaultSortField() string {
	return "CreatedAt"
}

var deepexiappAPIToModelFieldMap = map[string]string{
	"id":          "AppId",
	"name":        "AppName",
	"created_at":  "CreatedAt",
	"description": "AppDesc",
}

// APIToModelFieldMap returns a map from API names to field names for model
// 应用.
func (p *DeepexiApp) APIToModelFieldMap() map[string]string {
	return deepexiappAPIToModelFieldMap
}
