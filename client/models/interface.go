package models

type RecordX interface {
	Equals(RecordX) bool
	GetID() (uint, bool)
	GetZoneID() uint
	Refs() map[string]string
	Serialise() map[string]string
	Type() string
}
