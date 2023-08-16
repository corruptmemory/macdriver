// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentCloudKitContainerEventRequest] class.
var PersistentCloudKitContainerEventRequestClass = _PersistentCloudKitContainerEventRequestClass{objc.GetClass("NSPersistentCloudKitContainerEventRequest")}

type _PersistentCloudKitContainerEventRequestClass struct {
	objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEventRequest] class.
type IPersistentCloudKitContainerEventRequest interface {
	IPersistentStoreRequest
	ResultType() PersistentCloudKitContainerEventResultType
	SetResultType(value PersistentCloudKitContainerEventResultType)
}

// A request to fetch setup, import, or export events in a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest?language=objc
type PersistentCloudKitContainerEventRequest struct {
	PersistentStoreRequest
}

func PersistentCloudKitContainerEventRequestFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsAfterDate(date foundation.IDate) PersistentCloudKitContainerEventRequest {
	rv := objc.Call[PersistentCloudKitContainerEventRequest](pc, objc.Sel("fetchEventsAfterDate:"), objc.Ptr(date))
	return rv
}

// Creates a fetch request for events after a specified date from a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618815-fetcheventsafterdate?language=objc
func PersistentCloudKitContainerEventRequest_FetchEventsAfterDate(date foundation.IDate) PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequestClass.FetchEventsAfterDate(date)
}

func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsMatchingFetchRequest(fetchRequest IFetchRequest) PersistentCloudKitContainerEventRequest {
	rv := objc.Call[PersistentCloudKitContainerEventRequest](pc, objc.Sel("fetchEventsMatchingFetchRequest:"), objc.Ptr(fetchRequest))
	return rv
}

// Creates a fetch request for events that match a specified fetch request from a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618817-fetcheventsmatchingfetchrequest?language=objc
func PersistentCloudKitContainerEventRequest_FetchEventsMatchingFetchRequest(fetchRequest IFetchRequest) PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequestClass.FetchEventsMatchingFetchRequest(fetchRequest)
}

func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsAfterEvent(event IPersistentCloudKitContainerEvent) PersistentCloudKitContainerEventRequest {
	rv := objc.Call[PersistentCloudKitContainerEventRequest](pc, objc.Sel("fetchEventsAfterEvent:"), objc.Ptr(event))
	return rv
}

// Creates a fetch request for events that occur after a specified event from a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618816-fetcheventsafterevent?language=objc
func PersistentCloudKitContainerEventRequest_FetchEventsAfterEvent(event IPersistentCloudKitContainerEvent) PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequestClass.FetchEventsAfterEvent(event)
}

func (pc _PersistentCloudKitContainerEventRequestClass) Alloc() PersistentCloudKitContainerEventRequest {
	rv := objc.Call[PersistentCloudKitContainerEventRequest](pc, objc.Sel("alloc"))
	return rv
}

func PersistentCloudKitContainerEventRequest_Alloc() PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequestClass.Alloc()
}

func (pc _PersistentCloudKitContainerEventRequestClass) New() PersistentCloudKitContainerEventRequest {
	rv := objc.Call[PersistentCloudKitContainerEventRequest](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentCloudKitContainerEventRequest() PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequestClass.New()
}

func (p_ PersistentCloudKitContainerEventRequest) Init() PersistentCloudKitContainerEventRequest {
	rv := objc.Call[PersistentCloudKitContainerEventRequest](p_, objc.Sel("init"))
	return rv
}

// Creates a fetch request for all events in a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618818-fetchrequestforevents?language=objc
func (pc _PersistentCloudKitContainerEventRequestClass) FetchRequestForEvents() FetchRequest {
	rv := objc.Call[FetchRequest](pc, objc.Sel("fetchRequestForEvents"))
	return rv
}

// Creates a fetch request for all events in a persistent CloudKit container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618818-fetchrequestforevents?language=objc
func PersistentCloudKitContainerEventRequest_FetchRequestForEvents() FetchRequest {
	return PersistentCloudKitContainerEventRequestClass.FetchRequestForEvents()
}

// The type of result that the request returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618819-resulttype?language=objc
func (p_ PersistentCloudKitContainerEventRequest) ResultType() PersistentCloudKitContainerEventResultType {
	rv := objc.Call[PersistentCloudKitContainerEventResultType](p_, objc.Sel("resultType"))
	return rv
}

// The type of result that the request returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainereventrequest/3618819-resulttype?language=objc
func (p_ PersistentCloudKitContainerEventRequest) SetResultType(value PersistentCloudKitContainerEventResultType) {
	objc.Call[objc.Void](p_, objc.Sel("setResultType:"), value)
}