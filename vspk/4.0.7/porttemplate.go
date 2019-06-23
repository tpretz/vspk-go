/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PortTemplateIdentity represents the Identity of the object
var PortTemplateIdentity = bambou.Identity{
	Name:     "porttemplate",
	Category: "porttemplates",
}

// PortTemplatesList represents a list of PortTemplates
type PortTemplatesList []*PortTemplate

// PortTemplatesAncestor is the interface that an ancestor of a PortTemplate must implement.
// An Ancestor is defined as an entity that has PortTemplate as a descendant.
// An Ancestor can get a list of its child PortTemplates, but not necessarily create one.
type PortTemplatesAncestor interface {
	PortTemplates(*bambou.FetchingInfo) (PortTemplatesList, *bambou.Error)
}

// PortTemplatesParent is the interface that a parent of a PortTemplate must implement.
// A Parent is defined as an entity that has PortTemplate as a child.
// A Parent is an Ancestor which can create a PortTemplate.
type PortTemplatesParent interface {
	PortTemplatesAncestor
	CreatePortTemplate(*PortTemplate) *bambou.Error
}

// PortTemplate represents the model of a porttemplate
type PortTemplate struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewPortTemplate returns a new *PortTemplate
func NewPortTemplate() *PortTemplate {

	return &PortTemplate{}
}

// Identity returns the Identity of the object.
func (o *PortTemplate) Identity() bambou.Identity {

	return PortTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PortTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PortTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PortTemplate from the server
func (o *PortTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PortTemplate into the server
func (o *PortTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PortTemplate from the server
func (o *PortTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PortTemplate
func (o *PortTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PortTemplate
func (o *PortTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VLANTemplates retrieves the list of child VLANTemplates of the PortTemplate
func (o *PortTemplate) VLANTemplates(info *bambou.FetchingInfo) (VLANTemplatesList, *bambou.Error) {

	var list VLANTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, VLANTemplateIdentity, &list, info)
	return list, err
}

// CreateVLANTemplate creates a new child VLANTemplate under the PortTemplate
func (o *PortTemplate) CreateVLANTemplate(child *VLANTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PortTemplate
func (o *PortTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PortTemplate
func (o *PortTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
