package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder provides operations to manage the deploymentSummary property of the microsoft.graph.iosManagedAppProtection entity.
type IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetQueryParameters read properties and relationships of the managedAppPolicyDeploymentSummary object.
type IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetQueryParameters
}
// IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal instantiates a new IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder and sets the default values.
func NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    m := &IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection%2Did}/deploymentSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder instantiates a new IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder and sets the default values.
func NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deploymentSummary for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the managedAppPolicyDeploymentSummary object.
// returns a ManagedAppPolicyDeploymentSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-managedapppolicydeploymentsummary-get?view=graph-rest-1.0
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyDeploymentSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyDeploymentSummaryable), nil
}
// Patch update the properties of a managedAppPolicyDeploymentSummary object.
// returns a ManagedAppPolicyDeploymentSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-managedapppolicydeploymentsummary-update?view=graph-rest-1.0
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyDeploymentSummaryable, requestConfiguration *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyDeploymentSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyDeploymentSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deploymentSummary for deviceAppManagement
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the managedAppPolicyDeploymentSummary object.
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a managedAppPolicyDeploymentSummary object.
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyDeploymentSummaryable, requestConfiguration *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder when successful
func (m *IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) WithUrl(rawUrl string)(*IosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    return NewIosmanagedappprotectionsItemDeploymentsummaryDeploymentSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
