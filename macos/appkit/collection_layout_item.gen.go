// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionLayoutItem] class.
var CollectionLayoutItemClass = _CollectionLayoutItemClass{objc.GetClass("NSCollectionLayoutItem")}

type _CollectionLayoutItemClass struct {
	objc.Class
}

// An interface definition for the [CollectionLayoutItem] class.
type ICollectionLayoutItem interface {
	objc.IObject
	LayoutSize() CollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	EdgeSpacing() CollectionLayoutEdgeSpacing
	SetEdgeSpacing(value ICollectionLayoutEdgeSpacing)
}

// The most basic component of a collection view’s layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem?language=objc
type CollectionLayoutItem struct {
	objc.Object
}

func CollectionLayoutItemFrom(ptr unsafe.Pointer) CollectionLayoutItem {
	return CollectionLayoutItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionLayoutItemClass) ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	rv := objc.Call[CollectionLayoutItem](cc, objc.Sel("itemWithLayoutSize:"), objc.Ptr(layoutSize))
	return rv
}

// Creates an item of the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213871-itemwithlayoutsize?language=objc
func CollectionLayoutItem_ItemWithLayoutSize(layoutSize ICollectionLayoutSize) CollectionLayoutItem {
	return CollectionLayoutItemClass.ItemWithLayoutSize(layoutSize)
}

func (cc _CollectionLayoutItemClass) Alloc() CollectionLayoutItem {
	rv := objc.Call[CollectionLayoutItem](cc, objc.Sel("alloc"))
	return rv
}

func CollectionLayoutItem_Alloc() CollectionLayoutItem {
	return CollectionLayoutItemClass.Alloc()
}

func (cc _CollectionLayoutItemClass) New() CollectionLayoutItem {
	rv := objc.Call[CollectionLayoutItem](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutItem() CollectionLayoutItem {
	return CollectionLayoutItemClass.New()
}

func (c_ CollectionLayoutItem) Init() CollectionLayoutItem {
	rv := objc.Call[CollectionLayoutItem](c_, objc.Sel("init"))
	return rv
}

// The item's size expressed in width and height dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213873-layoutsize?language=objc
func (c_ CollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	rv := objc.Call[CollectionLayoutSize](c_, objc.Sel("layoutSize"))
	return rv
}

// An array of the supplementary items attached to the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3213874-supplementaryitems?language=objc
func (c_ CollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	rv := objc.Call[[]CollectionLayoutSupplementaryItem](c_, objc.Sel("supplementaryItems"))
	return rv
}

// The amount of space added around the content of the item to adjust its final size after its position is computed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3199084-contentinsets?language=objc
func (c_ CollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	rv := objc.Call[DirectionalEdgeInsets](c_, objc.Sel("contentInsets"))
	return rv
}

// The amount of space added around the content of the item to adjust its final size after its position is computed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3199084-contentinsets?language=objc
func (c_ CollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	objc.Call[objc.Void](c_, objc.Sel("setContentInsets:"), value)
}

// The amount of space added around the boundaries of the item between other items and this item's container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3199085-edgespacing?language=objc
func (c_ CollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	rv := objc.Call[CollectionLayoutEdgeSpacing](c_, objc.Sel("edgeSpacing"))
	return rv
}

// The amount of space added around the boundaries of the item between other items and this item's container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nscollectionlayoutitem/3199085-edgespacing?language=objc
func (c_ CollectionLayoutItem) SetEdgeSpacing(value ICollectionLayoutEdgeSpacing) {
	objc.Call[objc.Void](c_, objc.Sel("setEdgeSpacing:"), objc.Ptr(value))
}