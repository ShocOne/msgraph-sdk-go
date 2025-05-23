// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MailboxProtectionUnit struct {
    ProtectionUnitBase
}
// NewMailboxProtectionUnit instantiates a new MailboxProtectionUnit and sets the default values.
func NewMailboxProtectionUnit()(*MailboxProtectionUnit) {
    m := &MailboxProtectionUnit{
        ProtectionUnitBase: *NewProtectionUnitBase(),
    }
    odataTypeValue := "#microsoft.graph.mailboxProtectionUnit"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMailboxProtectionUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMailboxProtectionUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxProtectionUnit(), nil
}
// GetDirectoryObjectId gets the directoryObjectId property value. The ID of the directory object.
// returns a *string when successful
func (m *MailboxProtectionUnit) GetDirectoryObjectId()(*string) {
    val, err := m.GetBackingStore().Get("directoryObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name of the directory object.
// returns a *string when successful
func (m *MailboxProtectionUnit) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmail gets the email property value. Email address associated with the directory object.
// returns a *string when successful
func (m *MailboxProtectionUnit) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MailboxProtectionUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProtectionUnitBase.GetFieldDeserializers()
    res["directoryObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryObjectId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MailboxProtectionUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProtectionUnitBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("directoryObjectId", m.GetDirectoryObjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectoryObjectId sets the directoryObjectId property value. The ID of the directory object.
func (m *MailboxProtectionUnit) SetDirectoryObjectId(value *string)() {
    err := m.GetBackingStore().Set("directoryObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name of the directory object.
func (m *MailboxProtectionUnit) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmail sets the email property value. Email address associated with the directory object.
func (m *MailboxProtectionUnit) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
type MailboxProtectionUnitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProtectionUnitBaseable
    GetDirectoryObjectId()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    SetDirectoryObjectId(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
}
