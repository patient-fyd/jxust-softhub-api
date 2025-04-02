// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ContentTags is the golang structure for table content_tags.
type ContentTags struct {
	ContentType string `json:"content_type" description:"内容类型，如news、project、resource等"`
	ContentId   uint   `json:"content_id"   description:"内容ID"`
	TagId       uint   `json:"tag_id"       description:"标签ID，关联tags表"`
}
