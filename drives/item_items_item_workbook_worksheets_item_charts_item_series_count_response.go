// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
