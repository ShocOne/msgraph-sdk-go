package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder provides operations to call the boundingRect method.
type ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, anotherRange *string)(*ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/dataBodyRange()/boundingRect(anotherRange='{anotherRange}')", pathParameters),
    }
    if anotherRange != nil {
        m.BaseRequestBuilder.PathParameters["anotherRange"] = *anotherRange
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function boundingRect
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookRangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable), nil
}
// ToGetRequestInformation invoke function boundingRect
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemDataBodyRangeBoundingRectWithAnotherRangeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}