package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse struct {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponse
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse{
        FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponse: *NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponseable interface {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
