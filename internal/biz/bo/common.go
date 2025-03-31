package bo

type Pagination struct {
	Limit  int32
	Offset int32
}

func FromPageAndSize(page, size int32) *Pagination {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	return &Pagination{
		Limit:  size,
		Offset: page * size,
	}
}

func (p *Pagination) GetOffset4Db() int {
	if p == nil || p.Offset < 0 {
		return 0
	}
	return int(p.Offset)
}

func (p *Pagination) GetLimit4Db() int {
	if p == nil || p.Limit <= 0 {
		return 10
	}
	return int(p.Limit)
}
