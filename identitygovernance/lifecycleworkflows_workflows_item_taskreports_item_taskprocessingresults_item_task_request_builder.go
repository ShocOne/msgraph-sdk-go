package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
type LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetQueryParameters the related workflow task
type LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/task{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the related workflow task
// returns a Taskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable, error) {
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
func (m *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskprocessingresultsItemTaskRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
