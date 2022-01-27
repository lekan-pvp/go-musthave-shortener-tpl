package models

type OriginLink struct {
	Link string
	Deleted bool
}

func (l OriginLink) IsDeleted() bool {
	return l.Deleted
}
