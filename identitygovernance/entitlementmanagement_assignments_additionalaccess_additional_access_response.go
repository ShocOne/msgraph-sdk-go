package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseable instead.
type EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse struct {
    EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponse
}
// NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse instantiates a new EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse and sets the default values.
func NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse()(*EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse) {
    m := &EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse{
        EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponse: *NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponse(),
    }
    return m
}
// CreateEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponse(), nil
}
// Deprecated: This class is obsolete. Use EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseable instead.
type EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessResponseable interface {
    EntitlementmanagementAssignmentsAdditionalaccessAdditionalAccessGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
