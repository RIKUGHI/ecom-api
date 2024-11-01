package service

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (h *HelloService) Hello() string {
	return "Hello World"
}
