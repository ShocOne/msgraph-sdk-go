package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder provides operations to call the percentRank_Exc method.
type ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/percentRank_Exc", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action percentRank_Exc
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action percentRank_Exc
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsPercentrank_excPercentRank_ExcRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
