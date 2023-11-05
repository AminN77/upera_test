package internal

type mockRepository struct {
}

func (*mockRepository) Add(p *Product) (*Product, error) {
	return p, nil
}

func (*mockRepository) Update(up *Product, id int) (*Product, *Product, []string, error) {
	return &Product{}, &Product{}, []string{"Color"}, nil
}

func (*mockRepository) Fetch(id int) (*Product, error) {
	return &Product{}, nil
}

type errMockRepository struct {
}

func (*errMockRepository) Add(p *Product) (*Product, error) {
	return nil, ErrRepo
}

func (*errMockRepository) Update(up *Product, id int) (*Product, *Product, []string, error) {
	return nil, nil, nil, ErrRepo
}

func (*errMockRepository) Fetch(id int) (*Product, error) {
	return nil, ErrRepo
}