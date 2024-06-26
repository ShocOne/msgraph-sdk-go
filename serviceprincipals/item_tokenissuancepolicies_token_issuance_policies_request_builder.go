package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
type ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetQueryParameters the tokenIssuancePolicies assigned to this service principal.
type ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetQueryParameters
}
// ByTokenIssuancePolicyId provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder when successful
func (m *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) ByTokenIssuancePolicyId(tokenIssuancePolicyId string)(*ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tokenIssuancePolicyId != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = tokenIssuancePolicyId
    }
    return NewItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal instantiates a new ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder and sets the default values.
func NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    m := &ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/tokenIssuancePolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder instantiates a new ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder and sets the default values.
func NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTokenissuancepoliciesCountRequestBuilder when successful
func (m *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) Count()(*ItemTokenissuancepoliciesCountRequestBuilder) {
    return NewItemTokenissuancepoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the tokenIssuancePolicies assigned to this service principal.
// returns a TokenIssuancePolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTokenIssuancePolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyCollectionResponseable), nil
}
// ToGetRequestInformation the tokenIssuancePolicies assigned to this service principal.
// returns a *RequestInformation when successful
func (m *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder when successful
func (m *ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    return NewItemTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
