package simple

type FooBarService struct {
	*FooService
	*BarService
}

func NweFooBarService(fooservice *FooService, barservice *BarService) *FooBarService {
	return &FooBarService{FooService: fooservice, BarService: barservice}
}
