package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
type ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/webParts/{webPart%2Did}/getPositionOfWebPart", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder instantiates a new ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPositionOfWebPart
// returns a WebPartPositionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartPositionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartPositionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartPositionable), nil
}
// ToPostRequestInformation invoke action getPositionOfWebPart
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
