package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder provides operations to call the cancel method.
type PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/{privilegedAccessGroupAssignmentScheduleRequest%2Did}/cancel", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel an access assignment request to a group whose membership and ownership are governed by PIM.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupassignmentschedulerequest-cancel?view=graph-rest-1.0
func (m *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation cancel an access assignment request to a group whose membership and ownership are governed by PIM.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
