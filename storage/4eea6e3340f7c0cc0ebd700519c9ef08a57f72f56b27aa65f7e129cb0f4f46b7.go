package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse struct {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageGetResponse
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse{
        FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageGetResponse: *NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageResponseable interface {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWithNameImageGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
