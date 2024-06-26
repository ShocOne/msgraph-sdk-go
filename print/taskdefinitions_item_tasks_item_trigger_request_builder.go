package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TaskdefinitionsItemTasksItemTriggerRequestBuilder provides operations to manage the trigger property of the microsoft.graph.printTask entity.
type TaskdefinitionsItemTasksItemTriggerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TaskdefinitionsItemTasksItemTriggerRequestBuilderGetQueryParameters the printTaskTrigger that triggered this task's execution. Read-only.
type TaskdefinitionsItemTasksItemTriggerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TaskdefinitionsItemTasksItemTriggerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TaskdefinitionsItemTasksItemTriggerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TaskdefinitionsItemTasksItemTriggerRequestBuilderGetQueryParameters
}
// NewTaskdefinitionsItemTasksItemTriggerRequestBuilderInternal instantiates a new TaskdefinitionsItemTasksItemTriggerRequestBuilder and sets the default values.
func NewTaskdefinitionsItemTasksItemTriggerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TaskdefinitionsItemTasksItemTriggerRequestBuilder) {
    m := &TaskdefinitionsItemTasksItemTriggerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/print/taskDefinitions/{printTaskDefinition%2Did}/tasks/{printTask%2Did}/trigger{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTaskdefinitionsItemTasksItemTriggerRequestBuilder instantiates a new TaskdefinitionsItemTasksItemTriggerRequestBuilder and sets the default values.
func NewTaskdefinitionsItemTasksItemTriggerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TaskdefinitionsItemTasksItemTriggerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTaskdefinitionsItemTasksItemTriggerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the printTaskTrigger that triggered this task's execution. Read-only.
// returns a PrintTaskTriggerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TaskdefinitionsItemTasksItemTriggerRequestBuilder) Get(ctx context.Context, requestConfiguration *TaskdefinitionsItemTasksItemTriggerRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintTaskTriggerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintTaskTriggerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintTaskTriggerable), nil
}
// ToGetRequestInformation the printTaskTrigger that triggered this task's execution. Read-only.
// returns a *RequestInformation when successful
func (m *TaskdefinitionsItemTasksItemTriggerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TaskdefinitionsItemTasksItemTriggerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TaskdefinitionsItemTasksItemTriggerRequestBuilder when successful
func (m *TaskdefinitionsItemTasksItemTriggerRequestBuilder) WithUrl(rawUrl string)(*TaskdefinitionsItemTasksItemTriggerRequestBuilder) {
    return NewTaskdefinitionsItemTasksItemTriggerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
