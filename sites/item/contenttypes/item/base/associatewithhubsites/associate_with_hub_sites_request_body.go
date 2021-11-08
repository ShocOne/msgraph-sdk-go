package associatewithhubsites

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssociateWithHubSitesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    hubSiteUrls []string;
    // 
    propagateToExistingLists *bool;
}
// Instantiates a new associateWithHubSitesRequestBody and sets the default values.
func NewAssociateWithHubSitesRequestBody()(*AssociateWithHubSitesRequestBody) {
    m := &AssociateWithHubSitesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssociateWithHubSitesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the hubSiteUrls property value. 
func (m *AssociateWithHubSitesRequestBody) GetHubSiteUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.hubSiteUrls
    }
}
// Gets the propagateToExistingLists property value. 
func (m *AssociateWithHubSitesRequestBody) GetPropagateToExistingLists()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateToExistingLists
    }
}
// The deserialization information for the current model
func (m *AssociateWithHubSitesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hubSiteUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetHubSiteUrls(res)
        return nil
    }
    res["propagateToExistingLists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPropagateToExistingLists(val)
        return nil
    }
    return res
}
func (m *AssociateWithHubSitesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssociateWithHubSitesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("hubSiteUrls", m.GetHubSiteUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateToExistingLists", m.GetPropagateToExistingLists())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AssociateWithHubSitesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the hubSiteUrls property value. 
// Parameters:
//  - value : Value to set for the hubSiteUrls property.
func (m *AssociateWithHubSitesRequestBody) SetHubSiteUrls(value []string)() {
    m.hubSiteUrls = value
}
// Sets the propagateToExistingLists property value. 
// Parameters:
//  - value : Value to set for the propagateToExistingLists property.
func (m *AssociateWithHubSitesRequestBody) SetPropagateToExistingLists(value *bool)() {
    m.propagateToExistingLists = value
}