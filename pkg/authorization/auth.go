package authorization

import "context"

type authService struct{}

func NewService() Service {
	return &authService{}
}
func (a *authService) Get(ctx context.Context) {

}

func (a *authService) ServiceStatus(ctx context.Context) (result int, err error) {
	return
}
