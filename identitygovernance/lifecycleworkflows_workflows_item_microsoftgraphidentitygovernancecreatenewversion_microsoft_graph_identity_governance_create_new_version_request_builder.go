package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder provides operations to call the createNewVersion method.
type LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/microsoft.graph.identityGovernance.createNewVersion", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a new version of the workflow object.
// returns a Workflowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-createnewversion?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) Post(ctx context.Context, body LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionCreateNewVersionPostRequestBodyable, requestConfiguration *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderPostRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable), nil
}
// ToPostRequestInformation create a new version of the workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, body LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionCreateNewVersionPostRequestBodyable, requestConfiguration *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
