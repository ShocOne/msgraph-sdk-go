package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCreatedobjectsCreatedObjectsRequestBuilder provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
type ItemCreatedobjectsCreatedObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCreatedobjectsCreatedObjectsRequestBuilderGetQueryParameters directory objects created by this service principal. Read-only. Nullable.
type ItemCreatedobjectsCreatedObjectsRequestBuilderGetQueryParameters struct {
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
// ItemCreatedobjectsCreatedObjectsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCreatedobjectsCreatedObjectsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCreatedobjectsCreatedObjectsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the createdObjects property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemCreatedobjectsDirectoryObjectItemRequestBuilder when successful
func (m *ItemCreatedobjectsCreatedObjectsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemCreatedobjectsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemCreatedobjectsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCreatedobjectsCreatedObjectsRequestBuilderInternal instantiates a new ItemCreatedobjectsCreatedObjectsRequestBuilder and sets the default values.
func NewItemCreatedobjectsCreatedObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCreatedobjectsCreatedObjectsRequestBuilder) {
    m := &ItemCreatedobjectsCreatedObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/createdObjects{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCreatedobjectsCreatedObjectsRequestBuilder instantiates a new ItemCreatedobjectsCreatedObjectsRequestBuilder and sets the default values.
func NewItemCreatedobjectsCreatedObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCreatedobjectsCreatedObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCreatedobjectsCreatedObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCreatedobjectsCountRequestBuilder when successful
func (m *ItemCreatedobjectsCreatedObjectsRequestBuilder) Count()(*ItemCreatedobjectsCountRequestBuilder) {
    return NewItemCreatedobjectsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get directory objects created by this service principal. Read-only. Nullable.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCreatedobjectsCreatedObjectsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCreatedobjectsCreatedObjectsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemCreatedobjectsGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemCreatedobjectsCreatedObjectsRequestBuilder) GraphServicePrincipal()(*ItemCreatedobjectsGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemCreatedobjectsGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation directory objects created by this service principal. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCreatedobjectsCreatedObjectsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCreatedobjectsCreatedObjectsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCreatedobjectsCreatedObjectsRequestBuilder when successful
func (m *ItemCreatedobjectsCreatedObjectsRequestBuilder) WithUrl(rawUrl string)(*ItemCreatedobjectsCreatedObjectsRequestBuilder) {
    return NewItemCreatedobjectsCreatedObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
