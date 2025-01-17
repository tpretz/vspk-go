/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FloatingIPACLTemplateIdentity represents the Identity of the object
var FloatingIPACLTemplateIdentity = bambou.Identity{
	Name:     "egressfloatingipacltemplate",
	Category: "egressfloatingipacltemplates",
}

// FloatingIPACLTemplatesList represents a list of FloatingIPACLTemplates
type FloatingIPACLTemplatesList []*FloatingIPACLTemplate

// FloatingIPACLTemplatesAncestor is the interface that an ancestor of a FloatingIPACLTemplate must implement.
// An Ancestor is defined as an entity that has FloatingIPACLTemplate as a descendant.
// An Ancestor can get a list of its child FloatingIPACLTemplates, but not necessarily create one.
type FloatingIPACLTemplatesAncestor interface {
	FloatingIPACLTemplates(*bambou.FetchingInfo) (FloatingIPACLTemplatesList, *bambou.Error)
}

// FloatingIPACLTemplatesParent is the interface that a parent of a FloatingIPACLTemplate must implement.
// A Parent is defined as an entity that has FloatingIPACLTemplate as a child.
// A Parent is an Ancestor which can create a FloatingIPACLTemplate.
type FloatingIPACLTemplatesParent interface {
	FloatingIPACLTemplatesAncestor
	CreateFloatingIPACLTemplate(*FloatingIPACLTemplate) *bambou.Error
}

// FloatingIPACLTemplate represents the model of a egressfloatingipacltemplate
type FloatingIPACLTemplate struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	Name                   string `json:"name,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	Active                 bool   `json:"active"`
	DefaultAllowIP         bool   `json:"defaultAllowIP"`
	DefaultAllowNonIP      bool   `json:"defaultAllowNonIP"`
	Description            string `json:"description,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	PolicyState            string `json:"policyState,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	PriorityType           string `json:"priorityType,omitempty"`
	AssociatedLiveEntityID string `json:"associatedLiveEntityID,omitempty"`
	AutoGeneratePriority   string `json:"autoGeneratePriority,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewFloatingIPACLTemplate returns a new *FloatingIPACLTemplate
func NewFloatingIPACLTemplate() *FloatingIPACLTemplate {

	return &FloatingIPACLTemplate{}
}

// Identity returns the Identity of the object.
func (o *FloatingIPACLTemplate) Identity() bambou.Identity {

	return FloatingIPACLTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FloatingIPACLTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FloatingIPACLTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FloatingIPACLTemplate from the server
func (o *FloatingIPACLTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FloatingIPACLTemplate into the server
func (o *FloatingIPACLTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FloatingIPACLTemplate from the server
func (o *FloatingIPACLTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the FloatingIPACLTemplate
func (o *FloatingIPACLTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the FloatingIPACLTemplate
func (o *FloatingIPACLTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// FloatingIPACLTemplateEntries retrieves the list of child FloatingIPACLTemplateEntries of the FloatingIPACLTemplate
func (o *FloatingIPACLTemplate) FloatingIPACLTemplateEntries(info *bambou.FetchingInfo) (FloatingIPACLTemplateEntriesList, *bambou.Error) {

	var list FloatingIPACLTemplateEntriesList
	err := bambou.CurrentSession().FetchChildren(o, FloatingIPACLTemplateEntryIdentity, &list, info)
	return list, err
}

// CreateFloatingIPACLTemplateEntry creates a new child FloatingIPACLTemplateEntry under the FloatingIPACLTemplate
func (o *FloatingIPACLTemplate) CreateFloatingIPACLTemplateEntry(child *FloatingIPACLTemplateEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the FloatingIPACLTemplate
func (o *FloatingIPACLTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the FloatingIPACLTemplate
func (o *FloatingIPACLTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
