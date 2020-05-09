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

type MyAppManageServer struct {
	resourceManager *resource.ResourceManager
}

func (s *MyAppManageServer) CreateMyApp(w http.ResponseWriter, r *http.Request) {
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

	AppManage, err := s.resourceManager.CreateMyAppManage(&app)
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

func (s *MyAppManageServer) GetMyApp(w http.ResponseWriter, r *http.Request) {
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

func (s *MyAppManageServer) writeErrorToResponse(w http.ResponseWriter, code int, err error) {
	glog.Errorf("Failed. Error: %+v", err)
	w.WriteHeader(code)
	errorResponse := api.Error{ErrorMessage: err.Error(), ErrorDetails: fmt.Sprintf("%+v", err)}
	errBytes, err := json.Marshal(errorResponse)
	if err != nil {
		w.Write([]byte("处理错误"))
	}
	w.Write(errBytes)
}

func NewMyAppManageServer(resourceManager *resource.ResourceManager) *MyAppManageServer {
	return &MyAppManageServer{resourceManager: resourceManager}
}
