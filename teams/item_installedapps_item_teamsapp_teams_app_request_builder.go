package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemInstalledappsItemTeamsappTeamsAppRequestBuilder provides operations to manage the teamsApp property of the microsoft.graph.teamsAppInstallation entity.
type ItemInstalledappsItemTeamsappTeamsAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetQueryParameters the app that is installed.
type ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetQueryParameters
}
// NewItemInstalledappsItemTeamsappTeamsAppRequestBuilderInternal instantiates a new ItemInstalledappsItemTeamsappTeamsAppRequestBuilder and sets the default values.
func NewItemInstalledappsItemTeamsappTeamsAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemTeamsappTeamsAppRequestBuilder) {
    m := &ItemInstalledappsItemTeamsappTeamsAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/installedApps/{teamsAppInstallation%2Did}/teamsApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInstalledappsItemTeamsappTeamsAppRequestBuilder instantiates a new ItemInstalledappsItemTeamsappTeamsAppRequestBuilder and sets the default values.
func NewItemInstalledappsItemTeamsappTeamsAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemTeamsappTeamsAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstalledappsItemTeamsappTeamsAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the app that is installed.
// returns a TeamsAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInstalledappsItemTeamsappTeamsAppRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppable), nil
}
// ToGetRequestInformation the app that is installed.
// returns a *RequestInformation when successful
func (m *ItemInstalledappsItemTeamsappTeamsAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInstalledappsItemTeamsappTeamsAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInstalledappsItemTeamsappTeamsAppRequestBuilder when successful
func (m *ItemInstalledappsItemTeamsappTeamsAppRequestBuilder) WithUrl(rawUrl string)(*ItemInstalledappsItemTeamsappTeamsAppRequestBuilder) {
    return NewItemInstalledappsItemTeamsappTeamsAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
