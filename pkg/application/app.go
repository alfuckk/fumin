package application

type appService struct{}

func NewService() Service {
	return &appService{}
}
