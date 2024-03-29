package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BaseModel struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      *time.Time         `json:"updated_at" bson:"updated_at"`
	DeletedAt      *time.Time         `json:"deleted_at" bson:"deleted_at"`
	Version        int                `json:"_version" bson:"_version"`
	usedProjection bool
}

type Model interface {
	Initialize(id primitive.ObjectID, now time.Time)
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	SetUsedProjection(flag bool)
	DidUseProjection() bool
	SetUpdatedAt()
}

func NewBaseModel() BaseModel {
	return BaseModel{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		Version:   1,
	}
}

func (s *BaseModel) GetId() string {
	return s.ID.Hex()
}

func (s *BaseModel) SetUpdatedAt() {
	t := time.Now().UTC()
	s.UpdatedAt = &t
}

func (s *BaseModel) SetID(id primitive.ObjectID) {
	s.ID = id
}

func (s *BaseModel) Initialize(id primitive.ObjectID, now time.Time) {
	s.ID = id
	s.CreatedAt = now.UTC()
	s.Version = 1
}

func (s *BaseModel) SetUsedProjection(flag bool) {
	s.usedProjection = flag
}
func (s *BaseModel) DidUseProjection() bool {
	return s.usedProjection
}
