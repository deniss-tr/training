package day

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Day {
	return allDays
}
