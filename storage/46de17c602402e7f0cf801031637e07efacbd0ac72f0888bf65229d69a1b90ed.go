package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse struct {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse{
        FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse: *NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseable interface {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}