package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalGetResponseable instead.
type FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse struct {
    FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalGetResponse
}
// NewFileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse instantiates a new FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse and sets the default values.
func NewFileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse()(*FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse) {
    m := &FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse{
        FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalGetResponse: *NewFileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalGetResponseable instead.
type FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalResponseable interface {
    FileStorageContainersItemDriveListItemsItemGetActivitiesByIntervalGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
