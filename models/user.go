package models

import "time"

type User struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	Email          string    `bson:"email" json:"email"`
	FirstName      string    `bson:"firstName" json:"firstName"`
	LastName       string    `bson:"lastName" json:"lastName"`
	PassHash       string    `bson:"passHash" json:"passHash"`
	Slug           string    `bson:"slug" json:"slug"`
	Bio            string    `bson:"bio" json:"bio"`
	IPath          string    `bson:"iPath" json:"iPath"`
	IsActive       bool      `bson:"isActive" json:"isActive"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	Communities    []string  `bson:"communities" json:"communities"`
	CommunityCount int       `bson:"communityCount" json:"communityCount"`
	IsDeleted      bool      `bson:"isDeleted" json:"isDeleted"`
}
