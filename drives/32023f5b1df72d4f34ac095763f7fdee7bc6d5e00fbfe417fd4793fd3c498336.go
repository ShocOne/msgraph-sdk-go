package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponseable instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemImageWithWidthWithHeightWithFittingModeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}