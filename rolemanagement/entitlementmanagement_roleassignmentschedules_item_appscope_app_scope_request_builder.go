package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
type EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetQueryParameters read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
type EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal instantiates a new EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    m := &EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentSchedules/{unifiedRoleAssignmentSchedule%2Did}/appScope{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder instantiates a new EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
// returns a AppScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppScopeable), nil
}
// ToGetRequestInformation read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
