// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Responder] class.
var ResponderClass = _ResponderClass{objc.GetClass("NSResponder")}

type _ResponderClass struct {
	objc.Class
}

// An interface definition for the [Responder] class.
type IResponder interface {
	objc.IObject
	TouchesMovedWithEvent(event IEvent)
	MouseEntered(event IEvent)
	MouseUp(event IEvent)
	NewWindowForTab(sender objc.IObject) objc.Object
	WillPresentError(error foundation.IError) foundation.Error
	PressureChangeWithEvent(event IEvent)
	WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool
	QuickLookWithEvent(event IEvent)
	OtherMouseDragged(event IEvent)
	InvalidateRestorableState()
	RightMouseDragged(event IEvent)
	ShouldBeTreatedAsInkEvent(event IEvent) bool
	ShowContextHelp(sender objc.IObject)
	NoResponderFor(eventSelector objc.Selector)
	MouseMoved(event IEvent)
	TabletProximity(event IEvent)
	BecomeFirstResponder() bool
	BeginGestureWithEvent(event IEvent)
	FlagsChanged(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	ScrollWheel(event IEvent)
	CursorUpdate(event IEvent)
	ValidateProposedFirstResponderForEvent(responder IResponder, event IEvent) bool
	RestoreStateWithCoder(coder foundation.ICoder)
	EndGestureWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	MouseExited(event IEvent)
	UpdateUserActivityState(userActivity foundation.IUserActivity)
	SmartMagnifyWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	PerformKeyEquivalent(event IEvent) bool
	FlushBufferedKeyEvents()
	KeyDown(event IEvent)
	SupplementalTargetForActionSender(action objc.Selector, sender objc.IObject) objc.Object
	TabletPoint(event IEvent)
	ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object
	SwipeWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseUp(event IEvent)
	MagnifyWithEvent(event IEvent)
	ResignFirstResponder() bool
	MakeTouchBar() TouchBar
	TryToPerformWith(action objc.Selector, object objc.IObject) bool
	MouseDragged(event IEvent)
	HelpRequested(eventPtr IEvent)
	WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool
	InterpretKeyEvents(eventArray []IEvent)
	PerformTextFinderAction(sender objc.IObject)
	MouseDown(event IEvent)
	KeyUp(event IEvent)
	ChangeModeWithEvent(event IEvent)
	RightMouseDown(event IEvent)
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer)
	RightMouseUp(event IEvent)
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue)
	UndoManager() foundation.UndoManager
	TouchBar() TouchBar
	SetTouchBar(value ITouchBar)
	NextResponder() Responder
	SetNextResponder(value IResponder)
	Menu() Menu
	SetMenu(value IMenu)
	AcceptsFirstResponder() bool
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.IUserActivity)
}

// An abstract class that forms the basis of event and command processing in AppKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder?language=objc
type Responder struct {
	objc.Object
}

func ResponderFrom(ptr unsafe.Pointer) Responder {
	return Responder{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ Responder) Init() Responder {
	rv := objc.Call[Responder](r_, objc.Sel("init"))
	return rv
}

func (rc _ResponderClass) Alloc() Responder {
	rv := objc.Call[Responder](rc, objc.Sel("alloc"))
	return rv
}

func Responder_Alloc() Responder {
	return ResponderClass.Alloc()
}

func (rc _ResponderClass) New() Responder {
	rv := objc.Call[Responder](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewResponder() Responder {
	return ResponderClass.New()
}

// Informs the receiver that one or more touches has moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524501-touchesmovedwithevent?language=objc
func (r_ Responder) TouchesMovedWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("touchesMovedWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the cursor has entered a tracking rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1529306-mouseentered?language=objc
func (r_ Responder) MouseEntered(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("mouseEntered:"), objc.Ptr(event))
}

// Returns the classes that support secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/3762523-allowedclassesforrestorablestate?language=objc
func (rc _ResponderClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.Call[[]objc.Class](rc, objc.Sel("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}

// Returns the classes that support secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/3762523-allowedclassesforrestorablestate?language=objc
func Responder_AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	return ResponderClass.AllowedClassesForRestorableStateKeyPath(keyPath)
}

// Informs the receiver that the user has released the left mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1535349-mouseup?language=objc
func (r_ Responder) MouseUp(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("mouseUp:"), objc.Ptr(event))
}

// Creates a new window to show as a tab in a tabbed window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1644675-newwindowfortab?language=objc
func (r_ Responder) NewWindowForTab(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](r_, objc.Sel("newWindowForTab:"), sender)
	rv.Autorelease()
	return rv
}

// Returns a custom version of the supplied error object that’s more suitable for presentation in alert sheets and dialogs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525188-willpresenterror?language=objc
func (r_ Responder) WillPresentError(error foundation.IError) foundation.Error {
	rv := objc.Call[foundation.Error](r_, objc.Sel("willPresentError:"), objc.Ptr(error))
	return rv
}

// Indicates a pressure change as the result of a user input event on a system that supports pressure sensitivity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534071-pressurechangewithevent?language=objc
func (r_ Responder) PressureChangeWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("pressureChangeWithEvent:"), objc.Ptr(event))
}

// Returns whether to forward elastic scrolling gesture events up the responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534209-wantsforwardedscrolleventsforaxi?language=objc
func (r_ Responder) WantsForwardedScrollEventsForAxis(axis EventGestureAxis) bool {
	rv := objc.Call[bool](r_, objc.Sel("wantsForwardedScrollEventsForAxis:"), axis)
	return rv
}

// Performs a Quick Look on the content at the location specified by the supplied event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1535080-quicklookwithevent?language=objc
func (r_ Responder) QuickLookWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("quickLookWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the user has moved the mouse with a button other than the left or right button pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1529804-othermousedragged?language=objc
func (r_ Responder) OtherMouseDragged(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("otherMouseDragged:"), objc.Ptr(event))
}

// Marks the responder’s interface-related state as dirty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526243-invalidaterestorablestate?language=objc
func (r_ Responder) InvalidateRestorableState() {
	objc.Call[objc.Void](r_, objc.Sel("invalidateRestorableState"))
}

// Informs the receiver that the user has moved the mouse with the right button pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1529135-rightmousedragged?language=objc
func (r_ Responder) RightMouseDragged(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("rightMouseDragged:"), objc.Ptr(event))
}

// Indicates whether a pen-down event should be treated as an ink event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534105-shouldbetreatedasinkevent?language=objc
func (r_ Responder) ShouldBeTreatedAsInkEvent(event IEvent) bool {
	rv := objc.Call[bool](r_, objc.Sel("shouldBeTreatedAsInkEvent:"), objc.Ptr(event))
	return rv
}

// Implemented by subclasses to invoke the help system, displaying information relevant to the receiver and its current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1529686-showcontexthelp?language=objc
func (r_ Responder) ShowContextHelp(sender objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("showContextHelp:"), sender)
}

// Handles the case where an event or action message falls off the end of the responder chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534197-noresponderfor?language=objc
func (r_ Responder) NoResponderFor(eventSelector objc.Selector) {
	objc.Call[objc.Void](r_, objc.Sel("noResponderFor:"), eventSelector)
}

// Informs the receiver that the mouse has moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525114-mousemoved?language=objc
func (r_ Responder) MouseMoved(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("mouseMoved:"), objc.Ptr(event))
}

// Informs the receiver that a tablet-proximity event has occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527018-tabletproximity?language=objc
func (r_ Responder) TabletProximity(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("tabletProximity:"), objc.Ptr(event))
}

// Notifies the receiver that it’s about to become first responder in its NSWindow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526750-becomefirstresponder?language=objc
func (r_ Responder) BecomeFirstResponder() bool {
	rv := objc.Call[bool](r_, objc.Sel("becomeFirstResponder"))
	return rv
}

// Informs the receiver that the user has begun a touch gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526368-begingesturewithevent?language=objc
func (r_ Responder) BeginGestureWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("beginGestureWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the user has pressed or released a modifier key (Shift, Control, and so on). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527647-flagschanged?language=objc
func (r_ Responder) FlagsChanged(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("flagsChanged:"), objc.Ptr(event))
}

// Informs the receiver that new set of touches has been recognized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1531151-touchesbeganwithevent?language=objc
func (r_ Responder) TouchesBeganWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("touchesBeganWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the mouse’s scroll wheel has moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534192-scrollwheel?language=objc
func (r_ Responder) ScrollWheel(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("scrollWheel:"), objc.Ptr(event))
}

// Informs the receiver that the mouse cursor has moved into a cursor rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525066-cursorupdate?language=objc
func (r_ Responder) CursorUpdate(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("cursorUpdate:"), objc.Ptr(event))
}

// Allows controls to determine when they should become first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527066-validateproposedfirstresponder?language=objc
func (r_ Responder) ValidateProposedFirstResponderForEvent(responder IResponder, event IEvent) bool {
	rv := objc.Call[bool](r_, objc.Sel("validateProposedFirstResponder:forEvent:"), objc.Ptr(responder), objc.Ptr(event))
	return rv
}

// Restores the interface-related state of the responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526253-restorestatewithcoder?language=objc
func (r_ Responder) RestoreStateWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](r_, objc.Sel("restoreStateWithCoder:"), objc.Ptr(coder))
}

// Informs the receiver that the user has ended a touch gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1531251-endgesturewithevent?language=objc
func (r_ Responder) EndGestureWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("endGestureWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that tracking of touches has been cancelled for any reason. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1530614-touchescancelledwithevent?language=objc
func (r_ Responder) TouchesCancelledWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("touchesCancelledWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the cursor has exited a tracking rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527561-mouseexited?language=objc
func (r_ Responder) MouseExited(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("mouseExited:"), objc.Ptr(event))
}

// Updates the state of the given user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534884-updateuseractivitystate?language=objc
func (r_ Responder) UpdateUserActivityState(userActivity foundation.IUserActivity) {
	objc.Call[objc.Void](r_, objc.Sel("updateUserActivityState:"), objc.Ptr(userActivity))
}

// Informs the receiver that the user performed a smart zoom gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1532984-smartmagnifywithevent?language=objc
func (r_ Responder) SmartMagnifyWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("smartMagnifyWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the user has begun a rotation gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525572-rotatewithevent?language=objc
func (r_ Responder) RotateWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("rotateWithEvent:"), objc.Ptr(event))
}

// Handle a key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524690-performkeyequivalent?language=objc
func (r_ Responder) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Call[bool](r_, objc.Sel("performKeyEquivalent:"), objc.Ptr(event))
	return rv
}

// Clears any unprocessed key events when overridden by subclasses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527264-flushbufferedkeyevents?language=objc
func (r_ Responder) FlushBufferedKeyEvents() {
	objc.Call[objc.Void](r_, objc.Sel("flushBufferedKeyEvents"))
}

// Informs the receiver that the user has pressed a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525805-keydown?language=objc
func (r_ Responder) KeyDown(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("keyDown:"), objc.Ptr(event))
}

// Finds a target for an action method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1535269-supplementaltargetforaction?language=objc
func (r_ Responder) SupplementalTargetForActionSender(action objc.Selector, sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](r_, objc.Sel("supplementalTargetForAction:sender:"), action, sender)
	return rv
}

// Informs the receiver that a tablet-point event has occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1530905-tabletpoint?language=objc
func (r_ Responder) TabletPoint(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("tabletPoint:"), objc.Ptr(event))
}

// Overridden by subclasses to determine what services are available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524638-validrequestorforsendtype?language=objc
func (r_ Responder) ValidRequestorForSendTypeReturnType(sendType PasteboardType, returnType PasteboardType) objc.Object {
	rv := objc.Call[objc.Object](r_, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}

// Informs the receiver that the user has begun a swipe gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524275-swipewithevent?language=objc
func (r_ Responder) SwipeWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("swipeWithEvent:"), objc.Ptr(event))
}

// Returns that a set of touches have been removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525779-touchesendedwithevent?language=objc
func (r_ Responder) TouchesEndedWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("touchesEndedWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the user has pressed a mouse button other than the left or right one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525719-othermousedown?language=objc
func (r_ Responder) OtherMouseDown(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("otherMouseDown:"), objc.Ptr(event))
}

// Informs the receiver that the user has released a mouse button other than the left or right button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1531343-othermouseup?language=objc
func (r_ Responder) OtherMouseUp(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("otherMouseUp:"), objc.Ptr(event))
}

// Informs the receiver that the user has begun a pinch gesture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525862-magnifywithevent?language=objc
func (r_ Responder) MagnifyWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("magnifyWithEvent:"), objc.Ptr(event))
}

// Notifies the receiver that it’s been asked to relinquish its status as first responder in its window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1532115-resignfirstresponder?language=objc
func (r_ Responder) ResignFirstResponder() bool {
	rv := objc.Call[bool](r_, objc.Sel("resignFirstResponder"))
	return rv
}

// Your custom subclass of the NSResponder class should override this method to create and configure your subclass’s default NSTouchBar object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/2544690-maketouchbar?language=objc
func (r_ Responder) MakeTouchBar() TouchBar {
	rv := objc.Call[TouchBar](r_, objc.Sel("makeTouchBar"))
	return rv
}

// Attempts to perform the method indicated by an action with a specified argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524516-trytoperform?language=objc
func (r_ Responder) TryToPerformWith(action objc.Selector, object objc.IObject) bool {
	rv := objc.Call[bool](r_, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}

// Informs the receiver that the user has moved the mouse with the left button pressed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527420-mousedragged?language=objc
func (r_ Responder) MouseDragged(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("mouseDragged:"), objc.Ptr(event))
}

// Displays context-sensitive help for the receiver if help has been registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525123-helprequested?language=objc
func (r_ Responder) HelpRequested(eventPtr IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("helpRequested:"), objc.Ptr(eventPtr))
}

// Implement this method to track gesture scroll events such as a swipe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527456-wantsscrolleventsforswipetrackin?language=objc
func (r_ Responder) WantsScrollEventsForSwipeTrackingOnAxis(axis EventGestureAxis) bool {
	rv := objc.Call[bool](r_, objc.Sel("wantsScrollEventsForSwipeTrackingOnAxis:"), axis)
	return rv
}

// Handles a series of key events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1531599-interpretkeyevents?language=objc
func (r_ Responder) InterpretKeyEvents(eventArray []IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("interpretKeyEvents:"), eventArray)
}

// Performs all find oriented actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1525967-performtextfinderaction?language=objc
func (r_ Responder) PerformTextFinderAction(sender objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("performTextFinderAction:"), sender)
}

// Informs the receiver that the user has pressed the left mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524634-mousedown?language=objc
func (r_ Responder) MouseDown(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("mouseDown:"), objc.Ptr(event))
}

// Informs the receiver that the user has released a key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1527436-keyup?language=objc
func (r_ Responder) KeyUp(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("keyUp:"), objc.Ptr(event))
}

// Informs the responder that performed a double-tap on the side of an Apple Pencil. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/3237219-changemodewithevent?language=objc
func (r_ Responder) ChangeModeWithEvent(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("changeModeWithEvent:"), objc.Ptr(event))
}

// Informs the receiver that the user has pressed the right mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1524727-rightmousedown?language=objc
func (r_ Responder) RightMouseDown(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("rightMouseDown:"), objc.Ptr(event))
}

// Presents an error alert to the user as a document-modal sheet attached to document window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534705-presenterror?language=objc
func (r_ Responder) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error foundation.IError, window IWindow, delegate objc.IObject, didPresentSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](r_, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), objc.Ptr(error), objc.Ptr(window), delegate, didPresentSelector, contextInfo)
}

// Informs the receiver that the user has released the right mouse button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526309-rightmouseup?language=objc
func (r_ Responder) RightMouseUp(event IEvent) {
	objc.Call[objc.Void](r_, objc.Sel("rightMouseUp:"), objc.Ptr(event))
}

// Saves the interface-related state of the responder to a keyed archiver either synchronously or asynchronously on the given operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/2876293-encoderestorablestatewithcoder?language=objc
func (r_ Responder) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.ICoder, queue foundation.IOperationQueue) {
	objc.Call[objc.Void](r_, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), objc.Ptr(coder), objc.Ptr(queue))
}

// The undo manager for this responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1535376-undomanager?language=objc
func (r_ Responder) UndoManager() foundation.UndoManager {
	rv := objc.Call[foundation.UndoManager](r_, objc.Sel("undoManager"))
	return rv
}

// The NSTouchBar object associated with the responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/2544731-touchbar?language=objc
func (r_ Responder) TouchBar() TouchBar {
	rv := objc.Call[TouchBar](r_, objc.Sel("touchBar"))
	return rv
}

// The NSTouchBar object associated with the responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/2544731-touchbar?language=objc
func (r_ Responder) SetTouchBar(value ITouchBar) {
	objc.Call[objc.Void](r_, objc.Sel("setTouchBar:"), objc.Ptr(value))
}

// Returns an array of key paths representing the restorable attributes of the responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526242-restorablestatekeypaths?language=objc
func (rc _ResponderClass) RestorableStateKeyPaths() []string {
	rv := objc.Call[[]string](rc, objc.Sel("restorableStateKeyPaths"))
	return rv
}

// Returns an array of key paths representing the restorable attributes of the responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1526242-restorablestatekeypaths?language=objc
func Responder_RestorableStateKeyPaths() []string {
	return ResponderClass.RestorableStateKeyPaths()
}

// The next responder after this one, or nil if it has none. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1528245-nextresponder?language=objc
func (r_ Responder) NextResponder() Responder {
	rv := objc.Call[Responder](r_, objc.Sel("nextResponder"))
	return rv
}

// The next responder after this one, or nil if it has none. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1528245-nextresponder?language=objc
func (r_ Responder) SetNextResponder(value IResponder) {
	objc.Call[objc.Void](r_, objc.Sel("setNextResponder:"), objc.Ptr(value))
}

// Returns the responder’s menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1533094-menu?language=objc
func (r_ Responder) Menu() Menu {
	rv := objc.Call[Menu](r_, objc.Sel("menu"))
	return rv
}

// Returns the responder’s menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1533094-menu?language=objc
func (r_ Responder) SetMenu(value IMenu) {
	objc.Call[objc.Void](r_, objc.Sel("setMenu:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the responder accepts first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1528708-acceptsfirstresponder?language=objc
func (r_ Responder) AcceptsFirstResponder() bool {
	rv := objc.Call[bool](r_, objc.Sel("acceptsFirstResponder"))
	return rv
}

// An object encapsulating a user activity supported by this responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534108-useractivity?language=objc
func (r_ Responder) UserActivity() foundation.UserActivity {
	rv := objc.Call[foundation.UserActivity](r_, objc.Sel("userActivity"))
	return rv
}

// An object encapsulating a user activity supported by this responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsresponder/1534108-useractivity?language=objc
func (r_ Responder) SetUserActivity(value foundation.IUserActivity) {
	objc.Call[objc.Void](r_, objc.Sel("setUserActivity:"), objc.Ptr(value))
}