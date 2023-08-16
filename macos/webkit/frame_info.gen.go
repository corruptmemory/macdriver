// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FrameInfo] class.
var FrameInfoClass = _FrameInfoClass{objc.GetClass("WKFrameInfo")}

type _FrameInfoClass struct {
	objc.Class
}

// An interface definition for the [FrameInfo] class.
type IFrameInfo interface {
	objc.IObject
	Request() foundation.URLRequest
	IsMainFrame() bool
	SecurityOrigin() SecurityOrigin
}

// An object that contains information about a frame on a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo?language=objc
type FrameInfo struct {
	objc.Object
}

func FrameInfoFrom(ptr unsafe.Pointer) FrameInfo {
	return FrameInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FrameInfoClass) Alloc() FrameInfo {
	rv := objc.Call[FrameInfo](fc, objc.Sel("alloc"))
	return rv
}

func FrameInfo_Alloc() FrameInfo {
	return FrameInfoClass.Alloc()
}

func (fc _FrameInfoClass) New() FrameInfo {
	rv := objc.Call[FrameInfo](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFrameInfo() FrameInfo {
	return FrameInfoClass.New()
}

func (f_ FrameInfo) Init() FrameInfo {
	rv := objc.Call[FrameInfo](f_, objc.Sel("init"))
	return rv
}

// The frame’s current request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/1503091-request?language=objc
func (f_ FrameInfo) Request() foundation.URLRequest {
	rv := objc.Call[foundation.URLRequest](f_, objc.Sel("request"))
	return rv
}

// A Boolean value indicating whether the frame is the web site's main frame or a subframe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/1503096-mainframe?language=objc
func (f_ FrameInfo) IsMainFrame() bool {
	rv := objc.Call[bool](f_, objc.Sel("isMainFrame"))
	return rv
}

// The frame’s security origin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkframeinfo/1503089-securityorigin?language=objc
func (f_ FrameInfo) SecurityOrigin() SecurityOrigin {
	rv := objc.Call[SecurityOrigin](f_, objc.Sel("securityOrigin"))
	return rv
}