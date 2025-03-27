//go:build windows

package badge

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/wailsapp/wails/v3/pkg/application"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type windowsBadge struct {
	taskbarList  *ole.IUnknown
	taskbarList3 *ole.IDispatch
	hwnd         uintptr
	color        color.RGBA
}

// Windows COM GUIDs for taskbar
var (
	CLSID_TaskbarList = ole.NewGUID("{56FDF344-FD6D-11d0-958A-006097C9A090}")
	IID_ITaskbarList3 = ole.NewGUID("{EA1AFB91-9E28-4B86-90E9-9E9F8A5EEFAF}")
)

// Windows API constants for icon creation
const (
	LR_DEFAULTCOLOR = 0x0000
	IMAGE_ICON      = 1
	DI_NORMAL       = 0x0003
)

// Windows structures needed for icon creation
type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type ICONINFO struct {
	FIcon    int32
	XHotspot uint32
	YHotspot uint32
	HbmMask  windows.Handle
	HbmColor windows.Handle
}

// Windows API function declarations
var (
	user32             = windows.NewLazySystemDLL("user32.dll")
	gdi32              = windows.NewLazySystemDLL("gdi32.dll")
	createIconIndirect = user32.NewProc("CreateIconIndirect")
	createCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
	deleteObject       = gdi32.NewProc("DeleteObject")
	createDIBSection   = gdi32.NewProc("CreateDIBSection")
	selectObject       = gdi32.NewProc("SelectObject")
	deleteDC           = gdi32.NewProc("DeleteDC")
	setDIBits          = gdi32.NewProc("SetDIBits")
)

func New() *Service {
	return &Service{
		impl: &windowsBadge{},
	}
}

func (wb *windowsBadge) Startup(ctx context.Context, options application.ServiceOptions) error {
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return fmt.Errorf("failed to initialize COM: %w", err)
	}

	// Create TaskbarList instance
	unknown, err := ole.CreateInstance(CLSID_TaskbarList, ole.IID_IUnknown)
	if err != nil {
		return fmt.Errorf("failed to create TaskbarList instance: %w", err)
	}
	wb.taskbarList = unknown

	// Get ITaskbarList3 interface
	disp, err := unknown.QueryInterface(IID_ITaskbarList3)
	if err != nil {
		wb.taskbarList.Release()
		return fmt.Errorf("failed to get ITaskbarList3 interface: %w", err)
	}
	wb.taskbarList3 = disp

	// Initialize the taskbar
	_, err = oleutil.CallMethod(wb.taskbarList3, "HrInit")
	if err != nil {
		wb.taskbarList3.Release()
		wb.taskbarList.Release()
		return fmt.Errorf("failed to initialize taskbar: %w", err)
	}

	wb.color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	return nil
}

func (wb *windowsBadge) Shutdown() error {
	return nil
}

func (wb *windowsBadge) SetBadge(label string) error {
	if label == "" {
		return wb.clearBadge()
	}

	// Create a badge icon
	icon := wb.createBadgeIcon(label)
	defer deleteObject.Call(icon)

	// Set the overlay icon
	_, err := oleutil.CallMethod(wb.taskbarList3, "SetOverlayIcon",
		wb.hwnd, uintptr(icon), label)

	return err
}

func (wb *windowsBadge) clearBadge() error {
	_, err := oleutil.CallMethod(wb.taskbarList3, "SetOverlayIcon",
		wb.hwnd, uintptr(0), "")

	return err
}

func (s *windowsBadge) createBadgeIcon(text string) uintptr {
	// Create a small transparent bitmap
	const size = 16
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	// Draw a circle with the badge color
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx, dy := x-size/2, y-size/2
			if dx*dx+dy*dy <= size*size/4 {
				img.SetRGBA(x, y, s.color)
			}
		}
	}

	// Draw text
	d := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: basicfont.Face7x13,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(3 * 64), Y: fixed.Int26_6(12 * 64)},
	}
	d.DrawString(text)

	return createIconFromImage(img)
}

func createIconFromImage(img *image.RGBA) uintptr {
	// Get image dimensions
	width, height := img.Bounds().Dx(), img.Bounds().Dy()

	// Set up bitmap info header
	bmiHeader := BITMAPINFOHEADER{
		BiSize:        uint32(unsafe.Sizeof(BITMAPINFOHEADER{})),
		BiWidth:       int32(width),
		BiHeight:      -int32(height), // Negative height for top-down
		BiPlanes:      1,
		BiBitCount:    32,
		BiCompression: 0, // BI_RGB
	}

	// Create device context
	hdcScreen, _, _ := createCompatibleDC.Call(0)
	if hdcScreen == 0 {
		return 0
	}
	defer deleteDC.Call(hdcScreen)

	// Create DIB section for color bitmap
	var bits unsafe.Pointer
	hbmColor, _, _ := createDIBSection.Call(
		hdcScreen,
		uintptr(unsafe.Pointer(&bmiHeader)),
		0,
		uintptr(unsafe.Pointer(&bits)),
		0,
		0,
	)
	if hbmColor == 0 {
		return 0
	}
	defer deleteObject.Call(hbmColor)

	// Select bitmap into DC
	oldObj, _, _ := selectObject.Call(hdcScreen, hbmColor)
	if oldObj == 0 {
		return 0
	}
	defer selectObject.Call(hdcScreen, oldObj)

	// Copy image data to DIB
	pixelData := make([]byte, width*height*4)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get pixel color - RGBA from Go image
			r, g, b, a := img.At(x, y).RGBA()

			// Windows DIBs are in BGRA format, with 8 bits per channel
			i := (y*width + x) * 4
			pixelData[i+0] = byte(b >> 8) // Blue
			pixelData[i+1] = byte(g >> 8) // Green
			pixelData[i+2] = byte(r >> 8) // Red
			pixelData[i+3] = byte(a >> 8) // Alpha
		}
	}

	// Copy pixel data to bitmap
	memmove(bits, unsafe.Pointer(&pixelData[0]), uintptr(len(pixelData)))

	// Create monochrome mask bitmap (required for icons)
	hbmMask, _, _ := createCompatibleDC.Call(0)
	if hbmMask == 0 {
		return 0
	}
	defer deleteObject.Call(hbmMask)

	// Fill icon info structure
	iconInfo := ICONINFO{
		FIcon:    1, // TRUE for icon (vs. cursor)
		XHotspot: 0,
		YHotspot: 0,
		HbmMask:  windows.Handle(hbmMask),
		HbmColor: windows.Handle(hbmColor),
	}

	// Create icon
	hIcon, _, _ := createIconIndirect.Call(uintptr(unsafe.Pointer(&iconInfo)))

	return hIcon
}

// memmove copies data from src to dest for count bytes
func memmove(dest, src unsafe.Pointer, count uintptr) {
	// This is a simplified implementation
	// In a real app, use runtime.memmove or a proper unsafe copy
	destSlice := unsafe.Slice((*byte)(dest), count)
	srcSlice := unsafe.Slice((*byte)(src), count)
	copy(destSlice, srcSlice)
}
