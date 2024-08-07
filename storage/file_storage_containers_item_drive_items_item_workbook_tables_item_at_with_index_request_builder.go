package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder provides operations to call the itemAt method.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearFilters provides operations to call the clearFilters method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexClearFiltersRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ClearFilters()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexClearFiltersRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexClearFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexColumnsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Columns()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexColumnsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, index *int32)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/itemAt(index={index})", pathParameters),
    }
    if index != nil {
        m.BaseRequestBuilder.PathParameters["index"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*index), 10)
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ConvertToRange provides operations to call the convertToRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexConvertToRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ConvertToRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexConvertToRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexConvertToRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataBodyRange provides operations to call the dataBodyRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexDataBodyRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) DataBodyRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexDataBodyRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexDataBodyRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function itemAt
// returns a WorkbookTableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableable), nil
}
// HeaderRowRange provides operations to call the headerRowRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexHeaderRowRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) HeaderRowRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexHeaderRowRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexHeaderRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RangeEscaped provides operations to call the range method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) RangeEscaped()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReapplyFilters provides operations to call the reapplyFilters method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexReapplyFiltersRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ReapplyFilters()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexReapplyFiltersRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexReapplyFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rows provides operations to manage the rows property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRowsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Rows()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRowsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sort provides operations to manage the sort property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexSortRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Sort()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexSortRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexSortRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function itemAt
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// TotalRowRange provides operations to call the totalRowRange method.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexTotalRowRangeRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) TotalRowRange()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexTotalRowRangeRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexTotalRowRangeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Worksheet provides operations to manage the worksheet property of the microsoft.graph.workbookTable entity.
// returns a *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexWorksheetRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexRequestBuilder) Worksheet()(*FileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexWorksheetRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemWorkbookTablesItemAtWithIndexWorksheetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
