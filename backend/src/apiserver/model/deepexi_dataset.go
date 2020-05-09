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

type DeepexiDataset struct {
	DatasetId            string `gorm:"column:DatasetId; not null; primary_key"`
	DatasetName          string `gorm:"column:DatasetName; not null;"`
	DatasetDesc          string `gorm:"column:DatasetDesc;"`
	DatasetRootpath      string `gorm:"column:DatasetRootpath;"`
	AppId           string `gorm:"column:AppId;"`
	CreatedAt       int64  `gorm:"column:CreatedAt; not null;"`
	UpdateAt        int64  `gorm:"column:UpdateAt; not null;"`
	CreatedBy       string `gorm:"column:CreatedBy;"`
	UpdatedBy       string `gorm:"column:UpdatedBy"`
}

func (p DeepexiDataset) GetValueOfPrimaryKey() string {
	return fmt.Sprint(p.DatasetId)
}

func GetDeepexidatasetTablePrimaryKeyColumn() string {
	return "DatasetId"
}

// PrimaryKeyColumnName returns the primary key for model Pipeline.
func (p *DeepexiDataset) PrimaryKeyColumnName() string {
	return "DatasetId"
}

// DefaultSortField returns the default sorting field for model Pipeline.
func (p *DeepexiDataset) DefaultSortField() string {
	return "CreatedAt"
}

var deepexidatasetAPIToModelFieldMap = map[string]string{
	"id":          "DatasetId",
	"name":        "DatasetName",
	"created_at":  "CreatedAt",
	"description": "DatasetDesc",
}

// APIToModelFieldMap returns a map from API names to field names for model
// 应用.
func (p *DeepexiDataset) APIToModelFieldMap() map[string]string {
	return deepexidatasetAPIToModelFieldMap
}
