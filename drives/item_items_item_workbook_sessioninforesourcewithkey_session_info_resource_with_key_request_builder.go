package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder provides operations to call the sessionInfoResource method.
type ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderInternal instantiates a new ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, key *string)(*ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) {
    m := &ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/sessionInfoResource(key='{key}')", pathParameters),
    }
    if key != nil {
        m.BaseRequestBuilder.PathParameters["key"] = *key
    }
    return m
}
// NewItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder instantiates a new ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function sessionInfoResource
// returns a WorkbookSessionInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSessionInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookSessionInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookSessionInfoable), nil
}
// ToGetRequestInformation invoke function sessionInfoResource
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder when successful
func (m *ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder) {
    return NewItemItemsItemWorkbookSessioninforesourcewithkeySessionInfoResourceWithKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
