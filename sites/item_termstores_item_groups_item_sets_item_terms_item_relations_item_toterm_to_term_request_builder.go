package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder provides operations to manage the toTerm property of the microsoft.graph.termStore.relation entity.
type ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetQueryParameters the to [term] of the relation. The term to which the relationship is defined.
type ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetQueryParameters
}
// NewItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderInternal instantiates a new ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder and sets the default values.
func NewItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder) {
    m := &ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStores/{store%2Did}/groups/{group%2Did}/sets/{set%2Did}/terms/{term%2Did}/relations/{relation%2Did}/toTerm{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder instantiates a new ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder and sets the default values.
func NewItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the to [term] of the relation. The term to which the relationship is defined.
// returns a Termable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Termable), nil
}
// ToGetRequestInformation the to [term] of the relation. The term to which the relationship is defined.
// returns a *RequestInformation when successful
func (m *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder when successful
func (m *ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder) WithUrl(rawUrl string)(*ItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder) {
    return NewItemTermstoresItemGroupsItemSetsItemTermsItemRelationsItemTotermToTermRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
