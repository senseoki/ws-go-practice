package service

// OkService ...
type OkService interface {
	SayOK(to string) string
}

type MyOkService struct {
	Prefix string
}

func (s *MyOkService) SayOK(to string) string {
	return s.Prefix + " " + to
}
