package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperienceUser entity.
type ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
type ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters
}
// NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal instantiates a new ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder and sets the default values.
func NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, externalcourseActivityId *string)(*ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    m := &ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/employeeExperience/learningCourseActivities(externalcourseActivityId='{externalcourseActivityId}'){?%24expand,%24select}", pathParameters),
    }
    if externalcourseActivityId != nil {
        m.BaseRequestBuilder.PathParameters["externalcourseActivityId"] = *externalcourseActivityId
    }
    return m
}
// NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder instantiates a new ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder and sets the default values.
func NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
// returns a LearningCourseActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcourseactivity-get?view=graph-rest-1.0
func (m *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningCourseActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable), nil
}
// ToGetRequestInformation get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
// returns a *RequestInformation when successful
func (m *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) WithUrl(rawUrl string)(*ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
