package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder provides operations to call the offsetRange method.
type ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, columnOffset *int32, rowOffset *int32)(*ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/headerRowRange()/offsetRange(rowOffset={rowOffset},columnOffset={columnOffset})", pathParameters),
    }
    if columnOffset != nil {
        m.BaseRequestBuilder.PathParameters["columnOffset"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*columnOffset), 10)
    }
    if rowOffset != nil {
        m.BaseRequestBuilder.PathParameters["rowOffset"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*rowOffset), 10)
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function offsetRange
// returns a WorkbookRangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookRangeable, error) {
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
// ToGetRequestInformation invoke function offsetRange
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemHeaderRowRangeOffsetRangeWithRowOffsetWithColumnOffsetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}