// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse struct {
    DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponse
}
// NewDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse instantiates a new DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse and sets the default values.
func NewDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse()(*DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse) {
    m := &DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse{
        DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponse: *NewDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnResponseable interface {
    DirectoryRoleAssignmentScheduleInstancesFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
