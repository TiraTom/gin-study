package domain_obj

type ImportanceID struct {
	Id int64
}

func (i *ImportanceID) IsValid() bool {
	return i.Id > 0
}
