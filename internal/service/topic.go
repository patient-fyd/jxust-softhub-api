package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// ITopic 话题服务接口
type ITopic interface {
	// List 获取话题列表
	List(ctx context.Context, in model.TopicListInput) (*model.TopicListOutput, error)
}

// 获取话题服务
func Topic() ITopic {
	return localTopic
}

// 注册话题服务
func RegisterTopic(i ITopic) {
	localTopic = i
}

var localTopic ITopic
