// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OutputSettingsAssistant] class.
var OutputSettingsAssistantClass = _OutputSettingsAssistantClass{objc.GetClass("AVOutputSettingsAssistant")}

type _OutputSettingsAssistantClass struct {
	objc.Class
}

// An interface definition for the [OutputSettingsAssistant] class.
type IOutputSettingsAssistant interface {
	objc.IObject
	SourceVideoFormat() coremedia.VideoFormatDescriptionRef
	SetSourceVideoFormat(value coremedia.VideoFormatDescriptionRef)
	SourceVideoMinFrameDuration() coremedia.Time
	SetSourceVideoMinFrameDuration(value coremedia.Time)
	SourceVideoAverageFrameDuration() coremedia.Time
	SetSourceVideoAverageFrameDuration(value coremedia.Time)
	VideoSettings() map[string]objc.Object
	AudioSettings() map[string]objc.Object
	OutputFileType() FileType
	SourceAudioFormat() coremedia.AudioFormatDescriptionRef
	SetSourceAudioFormat(value coremedia.AudioFormatDescriptionRef)
}

// An object that builds audio and video output settings dictionaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant?language=objc
type OutputSettingsAssistant struct {
	objc.Object
}

func OutputSettingsAssistantFrom(ptr unsafe.Pointer) OutputSettingsAssistant {
	return OutputSettingsAssistant{
		Object: objc.ObjectFrom(ptr),
	}
}

func (oc _OutputSettingsAssistantClass) OutputSettingsAssistantWithPreset(presetIdentifier OutputSettingsPreset) OutputSettingsAssistant {
	rv := objc.Call[OutputSettingsAssistant](oc, objc.Sel("outputSettingsAssistantWithPreset:"), presetIdentifier)
	return rv
}

// Returns a new output settings assistant for a preset configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1387909-outputsettingsassistantwithprese?language=objc
func OutputSettingsAssistant_OutputSettingsAssistantWithPreset(presetIdentifier OutputSettingsPreset) OutputSettingsAssistant {
	return OutputSettingsAssistantClass.OutputSettingsAssistantWithPreset(presetIdentifier)
}

func (oc _OutputSettingsAssistantClass) Alloc() OutputSettingsAssistant {
	rv := objc.Call[OutputSettingsAssistant](oc, objc.Sel("alloc"))
	return rv
}

func OutputSettingsAssistant_Alloc() OutputSettingsAssistant {
	return OutputSettingsAssistantClass.Alloc()
}

func (oc _OutputSettingsAssistantClass) New() OutputSettingsAssistant {
	rv := objc.Call[OutputSettingsAssistant](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOutputSettingsAssistant() OutputSettingsAssistant {
	return OutputSettingsAssistantClass.New()
}

func (o_ OutputSettingsAssistant) Init() OutputSettingsAssistant {
	rv := objc.Call[OutputSettingsAssistant](o_, objc.Sel("init"))
	return rv
}

// Returns an array of preset values to use to initialize an output settings assistant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1388118-availableoutputsettingspresets?language=objc
func (oc _OutputSettingsAssistantClass) AvailableOutputSettingsPresets() []OutputSettingsPreset {
	rv := objc.Call[[]OutputSettingsPreset](oc, objc.Sel("availableOutputSettingsPresets"))
	return rv
}

// Returns an array of preset values to use to initialize an output settings assistant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1388118-availableoutputsettingspresets?language=objc
func OutputSettingsAssistant_AvailableOutputSettingsPresets() []OutputSettingsPreset {
	return OutputSettingsAssistantClass.AvailableOutputSettingsPresets()
}

// The format of the source video data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1387885-sourcevideoformat?language=objc
func (o_ OutputSettingsAssistant) SourceVideoFormat() coremedia.VideoFormatDescriptionRef {
	rv := objc.Call[coremedia.VideoFormatDescriptionRef](o_, objc.Sel("sourceVideoFormat"))
	return rv
}

// The format of the source video data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1387885-sourcevideoformat?language=objc
func (o_ OutputSettingsAssistant) SetSourceVideoFormat(value coremedia.VideoFormatDescriptionRef) {
	objc.Call[objc.Void](o_, objc.Sel("setSourceVideoFormat:"), value)
}

// A time value that describes the minimum frame duration of the video data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1386812-sourcevideominframeduration?language=objc
func (o_ OutputSettingsAssistant) SourceVideoMinFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](o_, objc.Sel("sourceVideoMinFrameDuration"))
	return rv
}

// A time value that describes the minimum frame duration of the video data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1386812-sourcevideominframeduration?language=objc
func (o_ OutputSettingsAssistant) SetSourceVideoMinFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](o_, objc.Sel("setSourceVideoMinFrameDuration:"), value)
}

// A time value that describes the average frame duration of the video data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1387414-sourcevideoaverageframeduration?language=objc
func (o_ OutputSettingsAssistant) SourceVideoAverageFrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](o_, objc.Sel("sourceVideoAverageFrameDuration"))
	return rv
}

// A time value that describes the average frame duration of the video data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1387414-sourcevideoaverageframeduration?language=objc
func (o_ OutputSettingsAssistant) SetSourceVideoAverageFrameDuration(value coremedia.Time) {
	objc.Call[objc.Void](o_, objc.Sel("setSourceVideoAverageFrameDuration:"), value)
}

// A video settings dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1386880-videosettings?language=objc
func (o_ OutputSettingsAssistant) VideoSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](o_, objc.Sel("videoSettings"))
	return rv
}

// An audio settings dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1386233-audiosettings?language=objc
func (o_ OutputSettingsAssistant) AudioSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](o_, objc.Sel("audioSettings"))
	return rv
}

// A uniform type identifier (UTI) that indicates the type of file to write. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1390842-outputfiletype?language=objc
func (o_ OutputSettingsAssistant) OutputFileType() FileType {
	rv := objc.Call[FileType](o_, objc.Sel("outputFileType"))
	return rv
}

// The format of the source audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1390673-sourceaudioformat?language=objc
func (o_ OutputSettingsAssistant) SourceAudioFormat() coremedia.AudioFormatDescriptionRef {
	rv := objc.Call[coremedia.AudioFormatDescriptionRef](o_, objc.Sel("sourceAudioFormat"))
	return rv
}

// The format of the source audio data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avoutputsettingsassistant/1390673-sourceaudioformat?language=objc
func (o_ OutputSettingsAssistant) SetSourceAudioFormat(value coremedia.AudioFormatDescriptionRef) {
	objc.Call[objc.Void](o_, objc.Sel("setSourceAudioFormat:"), value)
}