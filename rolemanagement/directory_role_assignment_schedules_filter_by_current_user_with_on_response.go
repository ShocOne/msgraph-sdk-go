// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse struct {
    DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnGetResponse
}
// NewDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse instantiates a new DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse and sets the default values.
func NewDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse()(*DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse) {
    m := &DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse{
        DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnGetResponse: *NewDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse(), nil
}
// Deprecated: This class is obsolete. Use DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnGetResponseable instead.
type DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponseable interface {
    DirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
