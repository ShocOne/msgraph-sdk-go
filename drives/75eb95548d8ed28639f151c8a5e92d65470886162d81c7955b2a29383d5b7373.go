package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable instead.
type ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse struct {
    ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse
}
// NewItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse instantiates a new ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse and sets the default values.
func NewItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse()(*ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse) {
    m := &ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse{
        ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse: *NewItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(),
    }
    return m
}
// CreateItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable instead.
type ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable interface {
    ItemItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}