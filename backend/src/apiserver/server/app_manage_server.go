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

package server

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	api "github.com/kubeflow/pipelines/backend/api/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/apiserver/resource"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"net/http"
)

type AppManageServer struct {
	resourceManager *resource.ResourceManager
}

func (s *AppManageServer) CreateApp(w http.ResponseWriter, r *http.Request) {
	glog.Infof("CreateApp called")

	r.ParseForm()
	ApplyName := r.Form.Get("ApplyName")
	ApplyType := r.Form.Get("ApplyType")
	ApplyFrame := r.Form.Get("ApplyFrame")
	ApplyEnvironment := r.Form.Get("ApplyEnvironment")
	ApplyBrief := r.Form.Get("ApplyBrief")

	app := model.AppManage{}
	app.ApplyName = ApplyName
	app.ApplyType = ApplyType
	app.ApplyFrame = ApplyFrame
	app.ApplyEnvironment = ApplyEnvironment
	app.ApplyBrief.String = ApplyBrief

	AppManage, err := s.resourceManager.CreateAppManage(&app)
	if err != nil {
		s.writeErrorToResponse(w, http.StatusBadRequest, util.Wrap(err, "Failed to create app."))
		return
	}

	pipelineJson, err := json.Marshal(AppManage)
	if err != nil {
		s.writeErrorToResponse(w, http.StatusInternalServerError, util.Wrap(err, "Error creating pipeline"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(pipelineJson)

}



func (s *AppManageServer) GetApp(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")

	appManage, err := s.resourceManager.GetAppManage(id)
	if err != nil {
		s.writeErrorToResponse(w, http.StatusBadRequest, util.Wrap(err, "Error get app ."))
		return
	}

	appManageJson, err := json.Marshal(appManage)
	if err != nil {
		s.writeErrorToResponse(w, http.StatusInternalServerError, util.Wrap(err, "Error get app ."))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(appManageJson)

}



func (s *AppManageServer) writeErrorToResponse(w http.ResponseWriter, code int, err error) {
	glog.Errorf("Failed. Error: %+v", err)
	w.WriteHeader(code)
	errorResponse := api.Error{ErrorMessage: err.Error(), ErrorDetails: fmt.Sprintf("%+v", err)}
	errBytes, err := json.Marshal(errorResponse)
	if err != nil {
		w.Write([]byte("处理错误"))
	}
	w.Write(errBytes)
}

func NewAppManageServer(resourceManager *resource.ResourceManager) *AppManageServer {
	return &AppManageServer{resourceManager: resourceManager}
}
