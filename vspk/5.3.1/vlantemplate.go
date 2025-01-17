/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VLANTemplateIdentity represents the Identity of the object
var VLANTemplateIdentity = bambou.Identity{
	Name:     "vlantemplate",
	Category: "vlantemplates",
}

// VLANTemplatesList represents a list of VLANTemplates
type VLANTemplatesList []*VLANTemplate

// VLANTemplatesAncestor is the interface that an ancestor of a VLANTemplate must implement.
// An Ancestor is defined as an entity that has VLANTemplate as a descendant.
// An Ancestor can get a list of its child VLANTemplates, but not necessarily create one.
type VLANTemplatesAncestor interface {
	VLANTemplates(*bambou.FetchingInfo) (VLANTemplatesList, *bambou.Error)
}

// VLANTemplatesParent is the interface that a parent of a VLANTemplate must implement.
// A Parent is defined as an entity that has VLANTemplate as a child.
// A Parent is an Ancestor which can create a VLANTemplate.
type VLANTemplatesParent interface {
	VLANTemplatesAncestor
	CreateVLANTemplate(*VLANTemplate) *bambou.Error
}

// VLANTemplate represents the model of a vlantemplate
type VLANTemplate struct {
	ID                           string `json:"ID,omitempty"`
	ParentID                     string `json:"parentID,omitempty"`
	ParentType                   string `json:"parentType,omitempty"`
	Owner                        string `json:"owner,omitempty"`
	Value                        int    `json:"value,omitempty"`
	LastUpdatedBy                string `json:"lastUpdatedBy,omitempty"`
	Description                  string `json:"description,omitempty"`
	EntityScope                  string `json:"entityScope,omitempty"`
	IsUplink                     bool   `json:"isUplink"`
	AssociatedConnectionType     string `json:"associatedConnectionType,omitempty"`
	AssociatedEgressQOSPolicyID  string `json:"associatedEgressQOSPolicyID,omitempty"`
	AssociatedIngressQOSPolicyID string `json:"associatedIngressQOSPolicyID,omitempty"`
	AssociatedUplinkConnectionID string `json:"associatedUplinkConnectionID,omitempty"`
	AssociatedVSCProfileID       string `json:"associatedVSCProfileID,omitempty"`
	DucVlan                      bool   `json:"ducVlan"`
	ExternalID                   string `json:"externalID,omitempty"`
	Type                         string `json:"type,omitempty"`
}

// NewVLANTemplate returns a new *VLANTemplate
func NewVLANTemplate() *VLANTemplate {

	return &VLANTemplate{
		IsUplink: false,
	}
}

// Identity returns the Identity of the object.
func (o *VLANTemplate) Identity() bambou.Identity {

	return VLANTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VLANTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VLANTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VLANTemplate from the server
func (o *VLANTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VLANTemplate into the server
func (o *VLANTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VLANTemplate from the server
func (o *VLANTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VLANTemplate
func (o *VLANTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VLANTemplate
func (o *VLANTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VLANTemplate
func (o *VLANTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VLANTemplate
func (o *VLANTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// UplinkConnections retrieves the list of child UplinkConnections of the VLANTemplate
func (o *VLANTemplate) UplinkConnections(info *bambou.FetchingInfo) (UplinkConnectionsList, *bambou.Error) {

	var list UplinkConnectionsList
	err := bambou.CurrentSession().FetchChildren(o, UplinkConnectionIdentity, &list, info)
	return list, err
}

// CreateUplinkConnection creates a new child UplinkConnection under the VLANTemplate
func (o *VLANTemplate) CreateUplinkConnection(child *UplinkConnection) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// BRConnections retrieves the list of child BRConnections of the VLANTemplate
func (o *VLANTemplate) BRConnections(info *bambou.FetchingInfo) (BRConnectionsList, *bambou.Error) {

	var list BRConnectionsList
	err := bambou.CurrentSession().FetchChildren(o, BRConnectionIdentity, &list, info)
	return list, err
}

// CreateBRConnection creates a new child BRConnection under the VLANTemplate
func (o *VLANTemplate) CreateBRConnection(child *BRConnection) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
