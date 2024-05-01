package models

import "time"

type CommunityContent struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	Type       int8      `bson:"type" json:"type"` // 0: Announcement, 1:Post
	Author     string    `bson:"author" json:"author"`
	Community  string    `bson:"community" json:"community"`
	Title      string    `bson:"title,omitempty" json:"title"` // Title of the Announcement, Post does not have a title
	Content    string    `bson:"content" json:"content"`
	IPaths     []string  `bson:"iPaths" json:"iPaths"`
	Permalink  string    `bson:"permalink" json:"permalink"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	Likes      []string  `bson:"likes" json:"likes"`           // User IDs who liked the post
	LikesCount int       `bson:"likesCount" json:"likesCount"` // Like count
	Tags       []string  `bson:"tags" json:"tags"`
	IsDeleted  bool      `bson:"isDeleted" json:"isDeleted"`
}
