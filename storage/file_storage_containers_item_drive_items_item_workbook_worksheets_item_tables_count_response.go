package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse struct {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountGetResponse
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse{
        FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountGetResponse: *NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountResponseable interface {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}