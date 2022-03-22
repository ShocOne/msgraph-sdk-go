package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Eventable 
type Eventable interface {
    OutlookItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowNewTimeProposals()(*bool)
    GetAttachments()([]Attachmentable)
    GetAttendees()([]Attendeeable)
    GetBody()(ItemBodyable)
    GetBodyPreview()(*string)
    GetCalendar()(Calendarable)
    GetEnd()(DateTimeTimeZoneable)
    GetExtensions()([]Extensionable)
    GetHasAttachments()(*bool)
    GetHideAttendees()(*bool)
    GetICalUId()(*string)
    GetImportance()(*Importance)
    GetInstances()([]Eventable)
    GetIsAllDay()(*bool)
    GetIsCancelled()(*bool)
    GetIsDraft()(*bool)
    GetIsOnlineMeeting()(*bool)
    GetIsOrganizer()(*bool)
    GetIsReminderOn()(*bool)
    GetLocation()(Locationable)
    GetLocations()([]Locationable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetOnlineMeeting()(OnlineMeetingInfoable)
    GetOnlineMeetingProvider()(*OnlineMeetingProviderType)
    GetOnlineMeetingUrl()(*string)
    GetOrganizer()(Recipientable)
    GetOriginalEndTimeZone()(*string)
    GetOriginalStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOriginalStartTimeZone()(*string)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderMinutesBeforeStart()(*int32)
    GetResponseRequested()(*bool)
    GetResponseStatus()(ResponseStatusable)
    GetSensitivity()(*Sensitivity)
    GetSeriesMasterId()(*string)
    GetShowAs()(*FreeBusyStatus)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetStart()(DateTimeTimeZoneable)
    GetSubject()(*string)
    GetTransactionId()(*string)
    GetType()(*EventType)
    GetWebLink()(*string)
    SetAllowNewTimeProposals(value *bool)()
    SetAttachments(value []Attachmentable)()
    SetAttendees(value []Attendeeable)()
    SetBody(value ItemBodyable)()
    SetBodyPreview(value *string)()
    SetCalendar(value Calendarable)()
    SetEnd(value DateTimeTimeZoneable)()
    SetExtensions(value []Extensionable)()
    SetHasAttachments(value *bool)()
    SetHideAttendees(value *bool)()
    SetICalUId(value *string)()
    SetImportance(value *Importance)()
    SetInstances(value []Eventable)()
    SetIsAllDay(value *bool)()
    SetIsCancelled(value *bool)()
    SetIsDraft(value *bool)()
    SetIsOnlineMeeting(value *bool)()
    SetIsOrganizer(value *bool)()
    SetIsReminderOn(value *bool)()
    SetLocation(value Locationable)()
    SetLocations(value []Locationable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetOnlineMeeting(value OnlineMeetingInfoable)()
    SetOnlineMeetingProvider(value *OnlineMeetingProviderType)()
    SetOnlineMeetingUrl(value *string)()
    SetOrganizer(value Recipientable)()
    SetOriginalEndTimeZone(value *string)()
    SetOriginalStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOriginalStartTimeZone(value *string)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderMinutesBeforeStart(value *int32)()
    SetResponseRequested(value *bool)()
    SetResponseStatus(value ResponseStatusable)()
    SetSensitivity(value *Sensitivity)()
    SetSeriesMasterId(value *string)()
    SetShowAs(value *FreeBusyStatus)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetStart(value DateTimeTimeZoneable)()
    SetSubject(value *string)()
    SetTransactionId(value *string)()
    SetType(value *EventType)()
    SetWebLink(value *string)()
}