package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder provides operations to call the resizedRange method.
type ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deltaColumns *int32, deltaRows *int32)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/cell(row={row},column={column})/resizedRange(deltaRows={deltaRows},deltaColumns={deltaColumns})", pathParameters),
    }
    if deltaColumns != nil {
        m.BaseRequestBuilder.PathParameters["deltaColumns"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*deltaColumns), 10)
    }
    if deltaRows != nil {
        m.BaseRequestBuilder.PathParameters["deltaRows"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*deltaRows), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function resizedRange
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable, error) {
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
// ToGetRequestInformation invoke function resizedRange
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemCellWithRowWithColumnResizedRangeWithDeltaRowsWithDeltaColumnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}