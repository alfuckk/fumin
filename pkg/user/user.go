package user

type userService struct{}

func NewService() Service {
	return &userService{}
}
sealos run labring/kubernetes:v1.29.7 labring/helm:v3.9.4 labring/calico:v3.26.1 --masters 96.9.214.165


