// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/macdriver/objc"
)

// Support for decorating items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemdecorating?language=objc
type PFileProviderItemDecorating interface {
	// optional
	Decorations() []FileProviderItemDecorationIdentifier
	HasDecorations() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderItemDecorating = (*FileProviderItemDecoratingObject)(nil)

// A concrete type for the [PFileProviderItemDecorating] protocol.
type FileProviderItemDecoratingObject struct {
	objc.Object
}

func (f_ FileProviderItemDecoratingObject) HasDecorations() bool {
	return f_.RespondsToSelector(objc.Sel("decorations"))
}

// Asks the item for an array of decorations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemdecorating/3172722-decorations?language=objc
func (f_ FileProviderItemDecoratingObject) Decorations() []FileProviderItemDecorationIdentifier {
	rv := objc.Call[[]FileProviderItemDecorationIdentifier](f_, objc.Sel("decorations"))
	return rv
}