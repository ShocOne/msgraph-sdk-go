package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemChangeScreenSharingRolePostRequestBody 
type CallsItemChangeScreenSharingRolePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The role property
    role *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole
}
// NewCallsItemChangeScreenSharingRolePostRequestBody instantiates a new CallsItemChangeScreenSharingRolePostRequestBody and sets the default values.
func NewCallsItemChangeScreenSharingRolePostRequestBody()(*CallsItemChangeScreenSharingRolePostRequestBody) {
    m := &CallsItemChangeScreenSharingRolePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateCallsItemChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemChangeScreenSharingRolePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemChangeScreenSharingRolePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemChangeScreenSharingRolePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemChangeScreenSharingRolePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseScreenSharingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
func (m *CallsItemChangeScreenSharingRolePostRequestBody) GetRole()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole) {
    return m.role
}
// Serialize serializes information the current object
func (m *CallsItemChangeScreenSharingRolePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemChangeScreenSharingRolePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRole sets the role property value. The role property
func (m *CallsItemChangeScreenSharingRolePostRequestBody) SetRole(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScreenSharingRole)() {
    m.role = value
}