/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSPortTemplateIdentity represents the Identity of the object
var NSPortTemplateIdentity = bambou.Identity{
	Name:     "nsporttemplate",
	Category: "nsporttemplates",
}

// NSPortTemplatesList represents a list of NSPortTemplates
type NSPortTemplatesList []*NSPortTemplate

// NSPortTemplatesAncestor is the interface that an ancestor of a NSPortTemplate must implement.
// An Ancestor is defined as an entity that has NSPortTemplate as a descendant.
// An Ancestor can get a list of its child NSPortTemplates, but not necessarily create one.
type NSPortTemplatesAncestor interface {
	NSPortTemplates(*bambou.FetchingInfo) (NSPortTemplatesList, *bambou.Error)
}

// NSPortTemplatesParent is the interface that a parent of a NSPortTemplate must implement.
// A Parent is defined as an entity that has NSPortTemplate as a child.
// A Parent is an Ancestor which can create a NSPortTemplate.
type NSPortTemplatesParent interface {
	NSPortTemplatesAncestor
	CreateNSPortTemplate(*NSPortTemplate) *bambou.Error
}

// NSPortTemplate represents the model of a nsporttemplate
type NSPortTemplate struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	InfrastructureProfileID     string `json:"infrastructureProfileID,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	Speed                       string `json:"speed,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	Mtu                         int    `json:"mtu,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewNSPortTemplate returns a new *NSPortTemplate
func NewNSPortTemplate() *NSPortTemplate {

	return &NSPortTemplate{
		Mtu: 1500,
	}
}

// Identity returns the Identity of the object.
func (o *NSPortTemplate) Identity() bambou.Identity {

	return NSPortTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSPortTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSPortTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSPortTemplate from the server
func (o *NSPortTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSPortTemplate into the server
func (o *NSPortTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSPortTemplate from the server
func (o *NSPortTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NSPortTemplate
func (o *NSPortTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSPortTemplate
func (o *NSPortTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VLANTemplates retrieves the list of child VLANTemplates of the NSPortTemplate
func (o *NSPortTemplate) VLANTemplates(info *bambou.FetchingInfo) (VLANTemplatesList, *bambou.Error) {

	var list VLANTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, VLANTemplateIdentity, &list, info)
	return list, err
}

// CreateVLANTemplate creates a new child VLANTemplate under the NSPortTemplate
func (o *NSPortTemplate) CreateVLANTemplate(child *VLANTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSPortTemplate
func (o *NSPortTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSPortTemplate
func (o *NSPortTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
