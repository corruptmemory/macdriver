// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewItem] class.
var CollectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}

type _CollectionViewItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewItem] class.
type ICollectionViewItem interface {
	IViewController
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	IsSelected() bool
	SetSelected(value bool)
	TextField() TextField
	SetTextField(value ITextField)
	CollectionView() CollectionView
	DraggingImageComponents() []DraggingImageComponent
	ImageView() ImageView
	SetImageView(value IImageView)
}

// The visual representation for a single data element in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem?language=objc
type CollectionViewItem struct {
	ViewController
}

func CollectionViewItemFrom(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ViewController: ViewControllerFrom(ptr),
	}
}

func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := objc.Call[CollectionViewItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewItem_Alloc() CollectionViewItem {
	return CollectionViewItemClass.Alloc()
}

func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := objc.Call[CollectionViewItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewItem() CollectionViewItem {
	return CollectionViewItemClass.New()
}

func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := objc.Call[CollectionViewItem](c_, objc.Sel("init"))
	return rv
}

func (c_ CollectionViewItem) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) CollectionViewItem {
	rv := objc.Call[CollectionViewItem](c_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func CollectionViewItem_InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) CollectionViewItem {
	return CollectionViewItemClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
}

// The highlight state currently applied to the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1527689-highlightstate?language=objc
func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	rv := objc.Call[CollectionViewItemHighlightState](c_, objc.Sel("highlightState"))
	return rv
}

// The highlight state currently applied to the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1527689-highlightstate?language=objc
func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	objc.Call[objc.Void](c_, objc.Sel("setHighlightState:"), value)
}

// A Boolean indicating whether the item is currently selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1528214-selected?language=objc
func (c_ CollectionViewItem) IsSelected() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSelected"))
	return rv
}

// A Boolean indicating whether the item is currently selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1528214-selected?language=objc
func (c_ CollectionViewItem) SetSelected(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSelected:"), value)
}

// A text field outlet that you can use to display a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1527126-textfield?language=objc
func (c_ CollectionViewItem) TextField() TextField {
	rv := objc.Call[TextField](c_, objc.Sel("textField"))
	return rv
}

// A text field outlet that you can use to display a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1527126-textfield?language=objc
func (c_ CollectionViewItem) SetTextField(value ITextField) {
	objc.Call[objc.Void](c_, objc.Sel("setTextField:"), objc.Ptr(value))
}

// The collection view that owns the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1528221-collectionview?language=objc
func (c_ CollectionViewItem) CollectionView() CollectionView {
	rv := objc.Call[CollectionView](c_, objc.Sel("collectionView"))
	return rv
}

// Dragging images for multi-image drag and drop support. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1528303-draggingimagecomponents?language=objc
func (c_ CollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.Call[[]DraggingImageComponent](c_, objc.Sel("draggingImageComponents"))
	return rv
}

// An image view outlet that you can use to display images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1525366-imageview?language=objc
func (c_ CollectionViewItem) ImageView() ImageView {
	rv := objc.Call[ImageView](c_, objc.Sel("imageView"))
	return rv
}

// An image view outlet that you can use to display images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewitem/1525366-imageview?language=objc
func (c_ CollectionViewItem) SetImageView(value IImageView) {
	objc.Call[objc.Void](c_, objc.Sel("setImageView:"), objc.Ptr(value))
}