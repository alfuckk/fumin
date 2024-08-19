package authorization

type Service struct {
	Login(ctx context.Context)
	ServiceStatus(ctx context.Context) (int, error)
}
