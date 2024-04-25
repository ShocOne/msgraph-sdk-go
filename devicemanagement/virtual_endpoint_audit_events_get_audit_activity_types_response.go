package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use VirtualEndpointAuditEventsGetAuditActivityTypesGetResponseable instead.
type VirtualEndpointAuditEventsGetAuditActivityTypesResponse struct {
    VirtualEndpointAuditEventsGetAuditActivityTypesGetResponse
}
// NewVirtualEndpointAuditEventsGetAuditActivityTypesResponse instantiates a new VirtualEndpointAuditEventsGetAuditActivityTypesResponse and sets the default values.
func NewVirtualEndpointAuditEventsGetAuditActivityTypesResponse()(*VirtualEndpointAuditEventsGetAuditActivityTypesResponse) {
    m := &VirtualEndpointAuditEventsGetAuditActivityTypesResponse{
        VirtualEndpointAuditEventsGetAuditActivityTypesGetResponse: *NewVirtualEndpointAuditEventsGetAuditActivityTypesGetResponse(),
    }
    return m
}
// CreateVirtualEndpointAuditEventsGetAuditActivityTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEndpointAuditEventsGetAuditActivityTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointAuditEventsGetAuditActivityTypesResponse(), nil
}
// Deprecated: This class is obsolete. Use VirtualEndpointAuditEventsGetAuditActivityTypesGetResponseable instead.
type VirtualEndpointAuditEventsGetAuditActivityTypesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointAuditEventsGetAuditActivityTypesGetResponseable
}