package service

import (
	"backend/biz/dal"
	"backend/biz/dal/model"
	"backend/biz/model/task"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type TaskService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTaskService(ctx context.Context, c *app.RequestContext) *TaskService {
	return &TaskService{ctx: ctx, c: c}
}

func (task *TaskService) AllMedals(ctx context.Context, req *task.AllMedalsReq) (*model.OlympicsData, error) {
	return dal.GetAllMedals()
}
