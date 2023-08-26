// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitybutton?language=objc
type PAccessibilityButton interface {
	// optional
	AccessibilityPerformPress() bool
	HasAccessibilityPerformPress() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool
}

// ensure impl type implements protocol interface
var _ PAccessibilityButton = (*AccessibilityButtonObject)(nil)

// A concrete type for the [PAccessibilityButton] protocol.
type AccessibilityButtonObject struct {
	objc.Object
}

func (a_ AccessibilityButtonObject) HasAccessibilityPerformPress() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformPress"))
}

// Simulates clicking the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitybutton/1525542-accessibilityperformpress?language=objc
func (a_ AccessibilityButtonObject) AccessibilityPerformPress() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformPress"))
	return rv
}

func (a_ AccessibilityButtonObject) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the button. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitybutton/1524910-accessibilitylabel?language=objc
func (a_ AccessibilityButtonObject) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}