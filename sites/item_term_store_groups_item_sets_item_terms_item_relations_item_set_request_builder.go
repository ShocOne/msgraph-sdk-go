// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3 "github.com/microsoftgraph/msgraph-sdk-go/models/termstore"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder provides operations to manage the set property of the microsoft.graph.termStore.relation entity.
type ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetQueryParameters the [set] in which the relation is relevant.
type ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetQueryParameters
}
// NewItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderInternal instantiates a new ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) {
    m := &ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/termStore/groups/{group%2Did}/sets/{set%2Did}/terms/{term%2Did}/relations/{relation%2Did}/set{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder instantiates a new ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder and sets the default values.
func NewItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the [set] in which the relation is relevant.
// returns a Setable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration)(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.CreateSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia3c27b33aa3d3ed80f9de797c48fbb8ed73f13887e301daf51f08450e9a634a3.Setable), nil
}
// ToGetRequestInformation the [set] in which the relation is relevant.
// returns a *RequestInformation when successful
func (m *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder when successful
func (m *ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) WithUrl(rawUrl string)(*ItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder) {
    return NewItemTermStoreGroupsItemSetsItemTermsItemRelationsItemSetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
