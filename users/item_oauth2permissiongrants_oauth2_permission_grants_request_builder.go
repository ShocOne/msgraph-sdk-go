package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
type ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetQueryParameters retrieve a list of oAuth2PermissionGrant entities, which represent delegated permissions granted to enable a client application to access an API on behalf of the user.
type ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetQueryParameters struct {
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
// ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetQueryParameters
}
// ByOAuth2PermissionGrantId provides operations to manage the oauth2PermissionGrants property of the microsoft.graph.user entity.
// returns a *ItemOauth2permissiongrantsOAuth2PermissionGrantItemRequestBuilder when successful
func (m *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) ByOAuth2PermissionGrantId(oAuth2PermissionGrantId string)(*ItemOauth2permissiongrantsOAuth2PermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if oAuth2PermissionGrantId != "" {
        urlTplParams["oAuth2PermissionGrant%2Did"] = oAuth2PermissionGrantId
    }
    return NewItemOauth2permissiongrantsOAuth2PermissionGrantItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderInternal instantiates a new ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder and sets the default values.
func NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) {
    m := &ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/oauth2PermissionGrants{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder instantiates a new ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder and sets the default values.
func NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOauth2permissiongrantsCountRequestBuilder when successful
func (m *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) Count()(*ItemOauth2permissiongrantsCountRequestBuilder) {
    return NewItemOauth2permissiongrantsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of oAuth2PermissionGrant entities, which represent delegated permissions granted to enable a client application to access an API on behalf of the user.
// returns a OAuth2PermissionGrantCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-list-oauth2permissiongrants?view=graph-rest-1.0
func (m *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OAuth2PermissionGrantCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOAuth2PermissionGrantCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OAuth2PermissionGrantCollectionResponseable), nil
}
// ToGetRequestInformation retrieve a list of oAuth2PermissionGrant entities, which represent delegated permissions granted to enable a client application to access an API on behalf of the user.
// returns a *RequestInformation when successful
func (m *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder when successful
func (m *ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) WithUrl(rawUrl string)(*ItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder) {
    return NewItemOauth2permissiongrantsOauth2PermissionGrantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
