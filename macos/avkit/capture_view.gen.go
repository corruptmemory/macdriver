// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/avfoundation"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureView] class.
var CaptureViewClass = _CaptureViewClass{objc.GetClass("AVCaptureView")}

type _CaptureViewClass struct {
	objc.Class
}

// An interface definition for the [CaptureView] class.
type ICaptureView interface {
	appkit.IView
	SetSessionShowVideoPreviewShowAudioPreview(session avfoundation.ICaptureSession, showVideoPreview bool, showAudioPreview bool)
	VideoGravity() avfoundation.LayerVideoGravity
	SetVideoGravity(value avfoundation.LayerVideoGravity)
	Session() avfoundation.CaptureSession
	FileOutput() avfoundation.CaptureFileOutput
	Delegate() CaptureViewDelegateWrapper
	SetDelegate(value PCaptureViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ControlsStyle() CaptureViewControlsStyle
	SetControlsStyle(value CaptureViewControlsStyle)
}

// A view that displays standard user interface controls for capturing media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview?language=objc
type CaptureView struct {
	appkit.View
}

func CaptureViewFrom(ptr unsafe.Pointer) CaptureView {
	return CaptureView{
		View: appkit.ViewFrom(ptr),
	}
}

func (cc _CaptureViewClass) Alloc() CaptureView {
	rv := objc.Call[CaptureView](cc, objc.Sel("alloc"))
	return rv
}

func CaptureView_Alloc() CaptureView {
	return CaptureViewClass.Alloc()
}

func (cc _CaptureViewClass) New() CaptureView {
	rv := objc.Call[CaptureView](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureView() CaptureView {
	return CaptureViewClass.New()
}

func (c_ CaptureView) Init() CaptureView {
	rv := objc.Call[CaptureView](c_, objc.Sel("init"))
	return rv
}

func (c_ CaptureView) InitWithFrame(frameRect foundation.Rect) CaptureView {
	rv := objc.Call[CaptureView](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewCaptureViewWithFrame(frameRect foundation.Rect) CaptureView {
	instance := CaptureViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Sets the view’s capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519163-setsession?language=objc
func (c_ CaptureView) SetSessionShowVideoPreviewShowAudioPreview(session avfoundation.ICaptureSession, showVideoPreview bool, showAudioPreview bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSession:showVideoPreview:showAudioPreview:"), objc.Ptr(session), showVideoPreview, showAudioPreview)
}

// A string value that defines how the capture view displays video within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519134-videogravity?language=objc
func (c_ CaptureView) VideoGravity() avfoundation.LayerVideoGravity {
	rv := objc.Call[avfoundation.LayerVideoGravity](c_, objc.Sel("videoGravity"))
	return rv
}

// A string value that defines how the capture view displays video within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519134-videogravity?language=objc
func (c_ CaptureView) SetVideoGravity(value avfoundation.LayerVideoGravity) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoGravity:"), value)
}

// The view’s associated capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519186-session?language=objc
func (c_ CaptureView) Session() avfoundation.CaptureSession {
	rv := objc.Call[avfoundation.CaptureSession](c_, objc.Sel("session"))
	return rv
}

// The capture file output used to record media data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519149-fileoutput?language=objc
func (c_ CaptureView) FileOutput() avfoundation.CaptureFileOutput {
	rv := objc.Call[avfoundation.CaptureFileOutput](c_, objc.Sel("fileOutput"))
	return rv
}

// The capture view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519144-delegate?language=objc
func (c_ CaptureView) Delegate() CaptureViewDelegateWrapper {
	rv := objc.Call[CaptureViewDelegateWrapper](c_, objc.Sel("delegate"))
	return rv
}

// The capture view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519144-delegate?language=objc
func (c_ CaptureView) SetDelegate(value PCaptureViewDelegate) {
	po0 := objc.WrapAsProtocol("AVCaptureViewDelegate", value)
	objc.SetAssociatedObject(c_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), po0)
}

// The capture view’s delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519144-delegate?language=objc
func (c_ CaptureView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The style of the capture controls presented by the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519147-controlsstyle?language=objc
func (c_ CaptureView) ControlsStyle() CaptureViewControlsStyle {
	rv := objc.Call[CaptureViewControlsStyle](c_, objc.Sel("controlsStyle"))
	return rv
}

// The style of the capture controls presented by the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureview/1519147-controlsstyle?language=objc
func (c_ CaptureView) SetControlsStyle(value CaptureViewControlsStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setControlsStyle:"), value)
}