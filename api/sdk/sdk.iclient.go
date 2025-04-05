// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sdk

import (
	"github.com/patient-fyd/jxust-softhub-api/api/auth"
	"github.com/patient-fyd/jxust-softhub-api/api/config"
	"github.com/patient-fyd/jxust-softhub-api/api/file"
	"github.com/patient-fyd/jxust-softhub-api/api/join"
	"github.com/patient-fyd/jxust-softhub-api/api/member"
	"github.com/patient-fyd/jxust-softhub-api/api/stat"
	"github.com/patient-fyd/jxust-softhub-api/api/tag"
	"github.com/patient-fyd/jxust-softhub-api/api/user"
	"github.com/patient-fyd/jxust-softhub-api/api/news"
	"github.com/patient-fyd/jxust-softhub-api/api/circle"
	"github.com/patient-fyd/jxust-softhub-api/api/comment"
	"github.com/patient-fyd/jxust-softhub-api/api/follow"
	"github.com/patient-fyd/jxust-softhub-api/api/like"
	"github.com/patient-fyd/jxust-softhub-api/api/post"
	"github.com/patient-fyd/jxust-softhub-api/api/topic"
)

type IClient interface {
	// 在此处添加新的服务接口
	AuthV1() auth.IAuthV1
	ConfigV1() config.IConfigV1
	FileV1() file.IFileV1
	JoinV1() join.IJoinV1
	MemberV1() member.IMemberV1
	StatV1() stat.IStatV1
	TagV1() tag.ITagV1
	UserV1() user.IUserV1
	NewsV1() news.INewsV1
	CircleV1() circle.ICircleV1
	CommentV1() comment.ICommentV1
	FollowV1() follow.IFollowV1
	LikeV1() like.ILikeV1
	PostV1() post.IPostV1
	TopicV1() topic.ITopicV1
}
