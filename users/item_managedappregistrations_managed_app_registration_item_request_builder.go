package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder provides operations to manage the managedAppRegistrations property of the microsoft.graph.user entity.
type ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters zero or more managed app registrations that belong to the user.
type ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetQueryParameters
}
// NewItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderInternal instantiates a new ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder) {
    m := &ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedAppRegistrations/{managedAppRegistration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder instantiates a new ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder and sets the default values.
func NewItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get zero or more managed app registrations that belong to the user.
// returns a ManagedAppRegistrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppRegistrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppRegistrationable), nil
}
// ToGetRequestInformation zero or more managed app registrations that belong to the user.
// returns a *RequestInformation when successful
func (m *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder when successful
func (m *ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder) WithUrl(rawUrl string)(*ItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder) {
    return NewItemManagedappregistrationsManagedAppRegistrationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
