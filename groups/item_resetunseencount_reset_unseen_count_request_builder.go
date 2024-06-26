package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemResetunseencountResetUnseenCountRequestBuilder provides operations to call the resetUnseenCount method.
type ItemResetunseencountResetUnseenCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemResetunseencountResetUnseenCountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemResetunseencountResetUnseenCountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemResetunseencountResetUnseenCountRequestBuilderInternal instantiates a new ItemResetunseencountResetUnseenCountRequestBuilder and sets the default values.
func NewItemResetunseencountResetUnseenCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetunseencountResetUnseenCountRequestBuilder) {
    m := &ItemResetunseencountResetUnseenCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/resetUnseenCount", pathParameters),
    }
    return m
}
// NewItemResetunseencountResetUnseenCountRequestBuilder instantiates a new ItemResetunseencountResetUnseenCountRequestBuilder and sets the default values.
func NewItemResetunseencountResetUnseenCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemResetunseencountResetUnseenCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemResetunseencountResetUnseenCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset the unseenCount of all the posts that the current user hasn't seen since their last visit. Supported for Microsoft 365 groups only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-resetunseencount?view=graph-rest-1.0
func (m *ItemResetunseencountResetUnseenCountRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemResetunseencountResetUnseenCountRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation reset the unseenCount of all the posts that the current user hasn't seen since their last visit. Supported for Microsoft 365 groups only.
// returns a *RequestInformation when successful
func (m *ItemResetunseencountResetUnseenCountRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemResetunseencountResetUnseenCountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemResetunseencountResetUnseenCountRequestBuilder when successful
func (m *ItemResetunseencountResetUnseenCountRequestBuilder) WithUrl(rawUrl string)(*ItemResetunseencountResetUnseenCountRequestBuilder) {
    return NewItemResetunseencountResetUnseenCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
