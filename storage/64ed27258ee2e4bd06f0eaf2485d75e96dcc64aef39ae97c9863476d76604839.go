package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse struct {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponse
}
// NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse instantiates a new FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse()(*FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse{
        FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponse: *NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponseable instead.
type FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountResponseable interface {
    FileStorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemRowsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}