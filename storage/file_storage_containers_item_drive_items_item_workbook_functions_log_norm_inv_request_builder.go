package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder provides operations to call the logNorm_Inv method.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/logNorm_Inv", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action logNorm_Inv
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action logNorm_Inv
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsLogNorm_InvRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
