package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder provides operations to call the completeMigration method.
type JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewJoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderInternal instantiates a new MicrosoftGraphCompleteMigrationRequestBuilder and sets the default values.
func NewJoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder) {
    m := &JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/channels/{channel%2Did}/microsoft.graph.completeMigration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewJoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder instantiates a new MicrosoftGraphCompleteMigrationRequestBuilder and sets the default values.
func NewJoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete the message migration process by removing `migration mode` from a channel in a team. `Migration mode` is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/channel-completemigration?view=graph-rest-1.0
func (m *JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder) Post(ctx context.Context, requestConfiguration *JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation complete the message migration process by removing `migration mode` from a channel in a team. `Migration mode` is a special state that prevents certain operations, like sending messages and adding members, during the data migration process. After a **completeMigration** request is made, you cannot import additional messages into the team. You can add members to the team after the request returns a successful response.
func (m *JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *JoinedTeamsItemChannelsItemMicrosoftGraphCompleteMigrationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}