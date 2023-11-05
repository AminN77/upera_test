package internal

import "time"

type Revision struct {
	RevisionNumber string    `json:"revisionNumber" bson:"revisionNumber"`
	ProductID      uint      `json:"productID" bson:"productID"`
	ChangedAttr    []string  `json:"changedAttr" bson:"changedAttr"`
	PrevValue      *Product  `json:"prevValue" bson:"prevValue"`
	NewValue       *Product  `json:"newValue" bson:"newValue"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
}
