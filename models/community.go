package models

import "time"

type Community struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Email       string    `bson:"email" json:"email"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Slug        string    `bson:"slug" json:"slug"`
	IPath       string    `bson:"iPath" json:"iPath"`
	Members     []string  `bson:"members" json:"members"`
	MemberCount int       `bson:"memberCount" json:"memberCount"`
	Categories  []uint16  `bson:"categories" json:"categories"`
	Posts       []string  `bson:"posts" json:"posts"`
	PostCount   int       `bson:"postCount" json:"postCount"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	IsDeleted   bool      `bson:"isDeleted" json:"isDeleted"`
}
