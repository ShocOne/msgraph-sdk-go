package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder provides operations to call the isPublished method.
type ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal instantiates a new ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    m := &ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}/isPublished()", pathParameters),
    }
    return m
}
// NewItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder instantiates a new ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder and sets the default values.
func NewItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get check the publishing status of a contentType in a content type hub site.
// Deprecated: This method is obsolete. Use GetAsIsPublishedGetResponse instead.
// returns a ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemContenttypesItemIspublishedIsPublishedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedResponseable), nil
}
// GetAsIsPublishedGetResponse check the publishing status of a contentType in a content type hub site.
// returns a ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-ispublished?view=graph-rest-1.0
func (m *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) GetAsIsPublishedGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemListsItemContenttypesItemIspublishedIsPublishedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedGetResponseable), nil
}
// ToGetRequestInformation check the publishing status of a contentType in a content type hub site.
// returns a *RequestInformation when successful
func (m *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewItemSitesItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
