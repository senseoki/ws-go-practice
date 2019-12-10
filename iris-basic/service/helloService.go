package service

// HelloService ...
type HelloService interface {
	SayHello(to string) string
}

type MyHelloService struct {
	Prefix string
}

func (s *MyHelloService) SayHello(to string) string {
	return s.Prefix + " " + to
}

type YourHelloService struct {
	Prefix string
	name   string
}

func (s *YourHelloService) SayHello(to string) string {
	return s.Prefix + " " + s.name + to
}
