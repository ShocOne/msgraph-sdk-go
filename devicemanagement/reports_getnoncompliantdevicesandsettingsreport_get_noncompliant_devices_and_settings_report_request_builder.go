package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder provides operations to call the getNoncompliantDevicesAndSettingsReport method.
type ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal instantiates a new ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder and sets the default values.
func NewReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    m := &ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getNoncompliantDevicesAndSettingsReport", pathParameters),
    }
    return m
}
// NewReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder instantiates a new ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder and sets the default values.
func NewReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-reporting-devicemanagementreports-getnoncompliantdevicesandsettingsreport?view=graph-rest-1.0
func (m *ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder) Post(ctx context.Context, body ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportPostRequestBodyable, requestConfiguration *ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation not yet documented
// returns a *RequestInformation when successful
func (m *ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportPostRequestBodyable, requestConfiguration *ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder when successful
func (m *ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewReportsGetnoncompliantdevicesandsettingsreportGetNoncompliantDevicesAndSettingsReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
