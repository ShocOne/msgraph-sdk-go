package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetQueryParameters the related workflow task
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/task{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the related workflow task
// returns a Taskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable), nil
}
// ToGetRequestInformation the related workflow task
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
