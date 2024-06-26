package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetQueryParameters
}
// NewItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal instantiates a new ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    m := &ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/{conversationMember%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder instantiates a new ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a collection of team members who have access to the shared channel.
// returns a ConversationMemberable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConversationMemberFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConversationMemberable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
// returns a *RequestInformation when successful
func (m *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder when successful
func (m *ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    return NewItemTeamPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
