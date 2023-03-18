package models

type RecordX interface {
	GetId() (uint, bool)
	GetParentId() uint
	Refs() map[string]string
	Serialise() map[string]string
}
