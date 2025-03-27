package badge

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type platformBadge interface {
	Startup(ctx context.Context, options application.ServiceOptions) error
	Shutdown() error

	SetBadge(value string) error
}

type Service struct {
	impl platformBadge
}

// ServiceName returns the name of the service.
func (bs *Service) ServiceName() string {
	return "github.com/wailsapp/wails/v3/services/badge"
}

// ServiceStartup is called when the service is loaded.
func (bs *Service) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return bs.impl.Startup(ctx, options)
}

// ServiceShutdown is called when the service is unloaded.
func (bs *Service) ServiceShutdown() error {
	return bs.impl.Shutdown()
}

func (bs *Service) SetBadge(label string) error {
	return bs.impl.SetBadge(label)
}
