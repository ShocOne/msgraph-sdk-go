package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthImageWithWidthGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
