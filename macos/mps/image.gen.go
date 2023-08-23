// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Image] class.
var ImageClass = _ImageClass{objc.GetClass("MPSImage")}

type _ImageClass struct {
	objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	objc.IObject
	BatchRepresentationWithSubRange(subRange foundation.Range) *foundation.Array
	ResourceSize() uint
	BatchRepresentation() *foundation.Array
	WriteBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout DataLayout, imageIndex uint)
	ReadBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout DataLayout, imageIndex uint)
	SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer)
	SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject)
	SubImageWithFeatureChannelRange(range_ foundation.Range) Image
	SetPurgeableState(state PurgeableState) PurgeableState
	Width() uint
	Usage() metal.TextureUsage
	Parent() Image
	Device() metal.DeviceWrapper
	NumberOfImages() uint
	Height() uint
	FeatureChannelFormat() ImageFeatureChannelFormat
	TextureType() metal.TextureType
	PixelFormat() metal.PixelFormat
	FeatureChannels() uint
	Label() string
	SetLabel(value string)
	Texture() metal.TextureWrapper
	Precision() uint
	PixelSize() uint
}

// A texture that may have more than four channels for use in convolutional neural networks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage?language=objc
type Image struct {
	objc.Object
}

func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ Image) InitWithParentImageSliceRangeFeatureChannels(parent IImage, sliceRange foundation.Range, featureChannels uint) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithParentImage:sliceRange:featureChannels:"), objc.Ptr(parent), sliceRange, featureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942493-initwithparentimage?language=objc
func NewImageWithParentImageSliceRangeFeatureChannels(parent IImage, sliceRange foundation.Range, featureChannels uint) Image {
	instance := ImageClass.Alloc().InitWithParentImageSliceRangeFeatureChannels(parent, sliceRange, featureChannels)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithTextureFeatureChannels(texture metal.PTexture, featureChannels uint) Image {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	rv := objc.Call[Image](i_, objc.Sel("initWithTexture:featureChannels:"), po0, featureChannels)
	return rv
}

// Initializes an image from a texture. The user-allocated texture has been created for a specific number of feature channels and number of images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2097547-initwithtexture?language=objc
func NewImageWithTextureFeatureChannels(texture metal.PTexture, featureChannels uint) Image {
	instance := ImageClass.Alloc().InitWithTextureFeatureChannels(texture, featureChannels)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithDeviceImageDescriptor(device metal.PDevice, imageDescriptor IImageDescriptor) Image {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Image](i_, objc.Sel("initWithDevice:imageDescriptor:"), po0, objc.Ptr(imageDescriptor))
	return rv
}

// Initializes an empty image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648920-initwithdevice?language=objc
func NewImageWithDeviceImageDescriptor(device metal.PDevice, imageDescriptor IImageDescriptor) Image {
	instance := ImageClass.Alloc().InitWithDeviceImageDescriptor(device, imageDescriptor)
	instance.Autorelease()
	return instance
}

func (ic _ImageClass) Alloc() Image {
	rv := objc.Call[Image](ic, objc.Sel("alloc"))
	return rv
}

func Image_Alloc() Image {
	return ImageClass.Alloc()
}

func (ic _ImageClass) New() Image {
	rv := objc.Call[Image](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImage() Image {
	return ImageClass.New()
}

func (i_ Image) Init() Image {
	rv := objc.Call[Image](i_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942492-batchrepresentationwithsubrange?language=objc
func (i_ Image) BatchRepresentationWithSubRange(subRange foundation.Range) *foundation.Array {
	rv := objc.Call[*foundation.Array](i_, objc.Sel("batchRepresentationWithSubRange:"), subRange)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942494-resourcesize?language=objc
func (i_ Image) ResourceSize() uint {
	rv := objc.Call[uint](i_, objc.Sel("resourceSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942495-batchrepresentation?language=objc
func (i_ Image) BatchRepresentation() *foundation.Array {
	rv := objc.Call[*foundation.Array](i_, objc.Sel("batchRepresentation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867189-writebytes?language=objc
func (i_ Image) WriteBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout DataLayout, imageIndex uint) {
	objc.Call[objc.Void](i_, objc.Sel("writeBytes:dataLayout:imageIndex:"), dataBytes, dataLayout, imageIndex)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867188-readbytes?language=objc
func (i_ Image) ReadBytesDataLayoutImageIndex(dataBytes unsafe.Pointer, dataLayout DataLayout, imageIndex uint) {
	objc.Call[objc.Void](i_, objc.Sel("readBytes:dataLayout:imageIndex:"), dataBytes, dataLayout, imageIndex)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942491-synchronizeoncommandbuffer?language=objc
func (i_ Image) SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](i_, objc.Sel("synchronizeOnCommandBuffer:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942491-synchronizeoncommandbuffer?language=objc
func (i_ Image) SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("synchronizeOnCommandBuffer:"), objc.Ptr(commandBufferObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867148-defaultallocator?language=objc
func (ic _ImageClass) DefaultAllocator() ImageAllocatorWrapper {
	rv := objc.Call[ImageAllocatorWrapper](ic, objc.Sel("defaultAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2867148-defaultallocator?language=objc
func Image_DefaultAllocator() ImageAllocatorWrapper {
	return ImageClass.DefaultAllocator()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942488-subimagewithfeaturechannelrange?language=objc
func (i_ Image) SubImageWithFeatureChannelRange(range_ foundation.Range) Image {
	rv := objc.Call[Image](i_, objc.Sel("subImageWithFeatureChannelRange:"), range_)
	return rv
}

// Set (or query) the purgeable state of the image’s underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648820-setpurgeablestate?language=objc
func (i_ Image) SetPurgeableState(state PurgeableState) PurgeableState {
	rv := objc.Call[PurgeableState](i_, objc.Sel("setPurgeableState:"), state)
	return rv
}

// The formal width of the image, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648884-width?language=objc
func (i_ Image) Width() uint {
	rv := objc.Call[uint](i_, objc.Sel("width"))
	return rv
}

// The intended usage of the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648828-usage?language=objc
func (i_ Image) Usage() metal.TextureUsage {
	rv := objc.Call[metal.TextureUsage](i_, objc.Sel("usage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942490-parent?language=objc
func (i_ Image) Parent() Image {
	rv := objc.Call[Image](i_, objc.Sel("parent"))
	return rv
}

// The device on which the image will be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648857-device?language=objc
func (i_ Image) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](i_, objc.Sel("device"))
	return rv
}

// The number of images for batch processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648900-numberofimages?language=objc
func (i_ Image) NumberOfImages() uint {
	rv := objc.Call[uint](i_, objc.Sel("numberOfImages"))
	return rv
}

// The formal height of the image, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648952-height?language=objc
func (i_ Image) Height() uint {
	rv := objc.Call[uint](i_, objc.Sel("height"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/3131715-featurechannelformat?language=objc
func (i_ Image) FeatureChannelFormat() ImageFeatureChannelFormat {
	rv := objc.Call[ImageFeatureChannelFormat](i_, objc.Sel("featureChannelFormat"))
	return rv
}

// The type of the underlying texture, typically [metal/mtltexturetype/mtltexturetype2d] or [metal/mtltexturetype/mtltexturetype2darray]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648948-texturetype?language=objc
func (i_ Image) TextureType() metal.TextureType {
	rv := objc.Call[metal.TextureType](i_, objc.Sel("textureType"))
	return rv
}

// The pixel format of the underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648844-pixelformat?language=objc
func (i_ Image) PixelFormat() metal.PixelFormat {
	rv := objc.Call[metal.PixelFormat](i_, objc.Sel("pixelFormat"))
	return rv
}

// The number of feature channels per pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648901-featurechannels?language=objc
func (i_ Image) FeatureChannels() uint {
	rv := objc.Call[uint](i_, objc.Sel("featureChannels"))
	return rv
}

// A string to help identify this object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648899-label?language=objc
func (i_ Image) Label() string {
	rv := objc.Call[string](i_, objc.Sel("label"))
	return rv
}

// A string to help identify this object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648899-label?language=objc
func (i_ Image) SetLabel(value string) {
	objc.Call[objc.Void](i_, objc.Sel("setLabel:"), value)
}

// The underlying texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648903-texture?language=objc
func (i_ Image) Texture() metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](i_, objc.Sel("texture"))
	return rv
}

// The number of bits of numeric precision available for each feature channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648880-precision?language=objc
func (i_ Image) Precision() uint {
	rv := objc.Call[uint](i_, objc.Sel("precision"))
	return rv
}

// The number of bytes from the first byte of one pixel to the first byte of the next pixel, in storage order. (Includes padding.) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648854-pixelsize?language=objc
func (i_ Image) PixelSize() uint {
	rv := objc.Call[uint](i_, objc.Sel("pixelSize"))
	return rv
}