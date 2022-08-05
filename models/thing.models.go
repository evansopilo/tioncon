package models

import "time"

type IThing interface {
	SetID(id string)
	SetDeviceID(id string)
	SetLimit(limit float32)
	SetLoc(loc Loc)
	SetTime(time time.Time)
	GetThing() *Thing
}

type Thing struct {
	ID       string    `json:"id,omitempty" bson:"_id,omitempty"`
	DeviceID string    `json:"device_id,omitempty" bson:"device_id,omitempty"`
	Speed    float32   `json:"speed,omitempty" bson:"speed,omitempty"`
	Limit    float32   `json:"limit,omitempty" bson:"limit,omitempty"`
	Loc      Loc       `json:"loc,omitempty" bson:"loc,omitempty"`
	Time     time.Time `json:"time,omitempty" bson:"time,omitempty"`
}

type Loc struct {
	Type        string    `json:"type,omitempty" bson:"type,omitempty"`
	Coordinates []float32 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
}

func NewThing() *Thing { return &Thing{} }

func (t *Thing) SetID(id string) { t.ID = id }

func (t *Thing) SetDeviceID(id string) { t.DeviceID = id }

func (t *Thing) SetLimit(limit float32) { t.Limit = limit }

func (t *Thing) SetLoc(loc Loc) { t.Loc = loc }

func (t *Thing) SetTime(time time.Time) { t.Time = time }

func (t *Thing) GetThing() *Thing {
	return &Thing{
		ID:       t.ID,
		DeviceID: t.DeviceID,
		Speed:    t.Speed,
		Limit:    t.Limit,
		Loc:      t.Loc,
		Time:     t.Time,
	}
}
