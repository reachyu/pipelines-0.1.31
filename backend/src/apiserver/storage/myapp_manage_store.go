package storage

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
)

type MyAppManageStoreInterface interface {
	GetMyAppManage(appId string) (*model.AppManage, error)
	CreateMyAppManage(*model.AppManage) (*model.AppManage, error)
}

type MyAppManageStore struct {
	db   *DB
	time util.TimeInterface
	uuid util.UUIDGeneratorInterface
}

func (s *MyAppManageStore) scanRows(rows *sql.Rows) ([]*model.AppManage, error) {
	var apps []*model.AppManage
	for rows.Next() {
		var ApplyName, ApplyType, ApplyFrame, ApplyEnvironment string
		var ApplyBrief, CreatedBy, UpdatedBy sql.NullString
		var id int64
		var CreatedAt, UpdateAt sql.NullInt64
		if err := rows.Scan(&id, &ApplyName, &ApplyType, &ApplyFrame, &ApplyEnvironment, &ApplyBrief, &CreatedAt, &UpdateAt, &CreatedBy, &UpdatedBy); err != nil {
			return nil, err
		}
		apps = append(apps, &model.AppManage{
			ID:               id,
			ApplyName:        ApplyName,
			ApplyType:        ApplyType,
			ApplyFrame:       ApplyFrame,
			ApplyEnvironment: ApplyEnvironment,
			ApplyBrief:       ApplyBrief,
			CreatedAt:        CreatedAt,
			UpdateAt:         UpdateAt,
			CreatedBy:        CreatedBy,
			UpdatedBy:        UpdatedBy})
	}
	return apps, nil
}

func (s *MyAppManageStore) GetMyAppManage(id string) (*model.AppManage, error) {
	sql, args, err := sq.
		Select("*").
		From("app_manages").
		Where(sq.Eq{"id": id}).
		Limit(1).ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to get app_manages: %v", err.Error())
	}
	r, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to get app_manages: %v", err.Error())
	}
	defer r.Close()
	apps, err := s.scanRows(r)

	if err != nil || len(apps) > 1 {
		return nil, util.NewInternalServerError(err, "Failed to get app_manages: %v", err.Error())
	}
	if len(apps) == 0 {
		return nil, util.NewResourceNotFoundError("AppManage", fmt.Sprint(id))
	}
	return apps[0], nil
}

func (s *MyAppManageStore) CreateMyAppManage(p *model.AppManage) (*model.AppManage, error) {
	newAppManage := *p
	now := s.time.Now().Unix()
	sql, args, err := sq.
		Insert("app_manages").
		SetMap(
			sq.Eq{
				"apply_name":        newAppManage.ApplyName,
				"apply_type":        newAppManage.ApplyType,
				"apply_frame":       newAppManage.ApplyFrame,
				"apply_environment": newAppManage.ApplyEnvironment,
				"apply_brief":       newAppManage.ApplyBrief,
				"created_at":        now}).
		ToSql()
	if err != nil {
		return nil, util.NewInternalServerError(err, "Failed to create query to insert app to app_manages table: %v",
			err.Error())
	}
	_, err = s.db.Exec(sql, args...)
	if err != nil {
		if s.db.IsDuplicateError(err) {
			return nil, util.NewInvalidInputError(
				"Failed to create a new app. The name %v already exist. Please specify a new name.", p.ApplyName)
		}
		return nil, util.NewInternalServerError(err, "Failed to add app to app_manages table: %v",
			err.Error())
	}
	return &newAppManage, nil
}

func NewMyAppManageStore(db *DB, time util.TimeInterface, uuid util.UUIDGeneratorInterface) *MyAppManageStore {
	return &MyAppManageStore{db: db, time: time, uuid: uuid}
}