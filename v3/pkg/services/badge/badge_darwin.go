//go:build darwin

package badge

/*
#cgo CFLAGS: -mmacosx-version-min=10.13 -x objective-c
#cgo LDFLAGS: -framework Cocoa -mmacosx-version-min=10.13
#import "./badge_darwin.h"
*/
import "C"
import (
	"context"
	"unsafe"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

// ServiceName returns the name of the service
func (ns *Service) ServiceName() string {
	return "github.com/wailsapp/wails/v3/services/badge"
}

// ServiceStartup is called when the service is loaded
func (ns *Service) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

// ServiceShutdown is called when the service is unloaded
func (ns *Service) ServiceShutdown() error {
	return nil
}

func (b *Service) SetBadge(label string) {
	var cLabel *C.char
	if label != "" {
		cLabel = C.CString(label)
		defer C.free(unsafe.Pointer(cLabel))
	}
	C.setBadgeLabel(cLabel)
}
