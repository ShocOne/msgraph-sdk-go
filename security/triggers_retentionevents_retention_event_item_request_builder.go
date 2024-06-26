package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// TriggersRetentioneventsRetentionEventItemRequestBuilder provides operations to manage the retentionEvents property of the microsoft.graph.security.triggersRoot entity.
type TriggersRetentioneventsRetentionEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TriggersRetentioneventsRetentionEventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggersRetentioneventsRetentionEventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TriggersRetentioneventsRetentionEventItemRequestBuilderGetQueryParameters read the properties and relationships of a retentionEvent object.
type TriggersRetentioneventsRetentionEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TriggersRetentioneventsRetentionEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggersRetentioneventsRetentionEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TriggersRetentioneventsRetentionEventItemRequestBuilderGetQueryParameters
}
// TriggersRetentioneventsRetentionEventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggersRetentioneventsRetentionEventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTriggersRetentioneventsRetentionEventItemRequestBuilderInternal instantiates a new TriggersRetentioneventsRetentionEventItemRequestBuilder and sets the default values.
func NewTriggersRetentioneventsRetentionEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggersRetentioneventsRetentionEventItemRequestBuilder) {
    m := &TriggersRetentioneventsRetentionEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/triggers/retentionEvents/{retentionEvent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTriggersRetentioneventsRetentionEventItemRequestBuilder instantiates a new TriggersRetentioneventsRetentionEventItemRequestBuilder and sets the default values.
func NewTriggersRetentioneventsRetentionEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggersRetentioneventsRetentionEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggersRetentioneventsRetentionEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a retentionEvent object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentionevent-delete?view=graph-rest-1.0
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TriggersRetentioneventsRetentionEventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get read the properties and relationships of a retentionEvent object.
// returns a RetentionEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentionevent-get?view=graph-rest-1.0
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TriggersRetentioneventsRetentionEventItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateRetentionEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventable), nil
}
// Patch update the navigation property retentionEvents in security
// returns a RetentionEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventable, requestConfiguration *TriggersRetentioneventsRetentionEventItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateRetentionEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventable), nil
}
// RetentionEventType provides operations to manage the retentionEventType property of the microsoft.graph.security.retentionEvent entity.
// returns a *TriggersRetentioneventsItemRetentioneventtypeRetentionEventTypeRequestBuilder when successful
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) RetentionEventType()(*TriggersRetentioneventsItemRetentioneventtypeRetentionEventTypeRequestBuilder) {
    return NewTriggersRetentioneventsItemRetentioneventtypeRetentionEventTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a retentionEvent object.
// returns a *RequestInformation when successful
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TriggersRetentioneventsRetentionEventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a retentionEvent object.
// returns a *RequestInformation when successful
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TriggersRetentioneventsRetentionEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property retentionEvents in security
// returns a *RequestInformation when successful
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionEventable, requestConfiguration *TriggersRetentioneventsRetentionEventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TriggersRetentioneventsRetentionEventItemRequestBuilder when successful
func (m *TriggersRetentioneventsRetentionEventItemRequestBuilder) WithUrl(rawUrl string)(*TriggersRetentioneventsRetentionEventItemRequestBuilder) {
    return NewTriggersRetentioneventsRetentionEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
