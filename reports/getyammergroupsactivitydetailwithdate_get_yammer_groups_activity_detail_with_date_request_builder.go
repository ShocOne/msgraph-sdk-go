package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder provides operations to call the getYammerGroupsActivityDetail method.
type GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderInternal instantiates a new GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    m := &GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getYammerGroupsActivityDetail(date={date})", pathParameters),
    }
    if date != nil {
        m.BaseRequestBuilder.PathParameters["date"] = (*date).String()
    }
    return m
}
// NewGetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder instantiates a new GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get details about Yammer groups activity by group.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getyammergroupsactivitydetail?view=graph-rest-1.0
func (m *GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder) Get(ctx context.Context, requestConfiguration *GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get details about Yammer groups activity by group.
// returns a *RequestInformation when successful
func (m *GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder when successful
func (m *GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder) WithUrl(rawUrl string)(*GetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewGetyammergroupsactivitydetailwithdateGetYammerGroupsActivityDetailWithDateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
