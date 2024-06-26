package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesArchiveRequestBuilder provides operations to call the archive method.
type ServiceannouncementMessagesArchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesArchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesArchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesArchiveRequestBuilderInternal instantiates a new ServiceannouncementMessagesArchiveRequestBuilder and sets the default values.
func NewServiceannouncementMessagesArchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesArchiveRequestBuilder) {
    m := &ServiceannouncementMessagesArchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/archive", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesArchiveRequestBuilder instantiates a new ServiceannouncementMessagesArchiveRequestBuilder and sets the default values.
func NewServiceannouncementMessagesArchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesArchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesArchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post archive a list of serviceUpdateMessages for the signed in user.
// Deprecated: This method is obsolete. Use PostAsArchivePostResponse instead.
// returns a ServiceannouncementMessagesArchiveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-archive?view=graph-rest-1.0
func (m *ServiceannouncementMessagesArchiveRequestBuilder) Post(ctx context.Context, body ServiceannouncementMessagesArchivePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesArchiveRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesArchiveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesArchiveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesArchiveResponseable), nil
}
// PostAsArchivePostResponse archive a list of serviceUpdateMessages for the signed in user.
// returns a ServiceannouncementMessagesArchivePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-archive?view=graph-rest-1.0
func (m *ServiceannouncementMessagesArchiveRequestBuilder) PostAsArchivePostResponse(ctx context.Context, body ServiceannouncementMessagesArchivePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesArchiveRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesArchivePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesArchivePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesArchivePostResponseable), nil
}
// ToPostRequestInformation archive a list of serviceUpdateMessages for the signed in user.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesArchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceannouncementMessagesArchivePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesArchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesArchiveRequestBuilder when successful
func (m *ServiceannouncementMessagesArchiveRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesArchiveRequestBuilder) {
    return NewServiceannouncementMessagesArchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
