package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder provides operations to call the gammaLn_Precise method.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/gammaLn_Precise", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action gammaLn_Precise
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PrecisePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action gammaLn_Precise
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PrecisePostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookFunctionsGammaLn_PreciseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
