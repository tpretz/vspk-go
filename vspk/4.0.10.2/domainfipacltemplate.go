/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DomainFIPAclTemplateIdentity represents the Identity of the object
var DomainFIPAclTemplateIdentity = bambou.Identity{
	Name:     "egressdomainfloatingipacltemplate",
	Category: "egressdomainfloatingipacltemplates",
}

// DomainFIPAclTemplatesList represents a list of DomainFIPAclTemplates
type DomainFIPAclTemplatesList []*DomainFIPAclTemplate

// DomainFIPAclTemplatesAncestor is the interface that an ancestor of a DomainFIPAclTemplate must implement.
// An Ancestor is defined as an entity that has DomainFIPAclTemplate as a descendant.
// An Ancestor can get a list of its child DomainFIPAclTemplates, but not necessarily create one.
type DomainFIPAclTemplatesAncestor interface {
	DomainFIPAclTemplates(*bambou.FetchingInfo) (DomainFIPAclTemplatesList, *bambou.Error)
}

// DomainFIPAclTemplatesParent is the interface that a parent of a DomainFIPAclTemplate must implement.
// A Parent is defined as an entity that has DomainFIPAclTemplate as a child.
// A Parent is an Ancestor which can create a DomainFIPAclTemplate.
type DomainFIPAclTemplatesParent interface {
	DomainFIPAclTemplatesAncestor
	CreateDomainFIPAclTemplate(*DomainFIPAclTemplate) *bambou.Error
}

// DomainFIPAclTemplate represents the model of a egressdomainfloatingipacltemplate
type DomainFIPAclTemplate struct {
	ID                     string        `json:"ID,omitempty"`
	ParentID               string        `json:"parentID,omitempty"`
	ParentType             string        `json:"parentType,omitempty"`
	Owner                  string        `json:"owner,omitempty"`
	Name                   string        `json:"name,omitempty"`
	LastUpdatedBy          string        `json:"lastUpdatedBy,omitempty"`
	Active                 bool          `json:"active"`
	DefaultAllowIP         bool          `json:"defaultAllowIP"`
	DefaultAllowNonIP      bool          `json:"defaultAllowNonIP"`
	Description            string        `json:"description,omitempty"`
	EntityScope            string        `json:"entityScope,omitempty"`
	Entries                []interface{} `json:"entries,omitempty"`
	PolicyState            string        `json:"policyState,omitempty"`
	Priority               int           `json:"priority,omitempty"`
	PriorityType           string        `json:"priorityType,omitempty"`
	AssociatedLiveEntityID string        `json:"associatedLiveEntityID,omitempty"`
	ExternalID             string        `json:"externalID,omitempty"`
}

// NewDomainFIPAclTemplate returns a new *DomainFIPAclTemplate
func NewDomainFIPAclTemplate() *DomainFIPAclTemplate {

	return &DomainFIPAclTemplate{}
}

// Identity returns the Identity of the object.
func (o *DomainFIPAclTemplate) Identity() bambou.Identity {

	return DomainFIPAclTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DomainFIPAclTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DomainFIPAclTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DomainFIPAclTemplate from the server
func (o *DomainFIPAclTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DomainFIPAclTemplate into the server
func (o *DomainFIPAclTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DomainFIPAclTemplate from the server
func (o *DomainFIPAclTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DomainFIPAclTemplate
func (o *DomainFIPAclTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DomainFIPAclTemplate
func (o *DomainFIPAclTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// DomainFIPAclTemplateEntries retrieves the list of child DomainFIPAclTemplateEntries of the DomainFIPAclTemplate
func (o *DomainFIPAclTemplate) DomainFIPAclTemplateEntries(info *bambou.FetchingInfo) (DomainFIPAclTemplateEntriesList, *bambou.Error) {

	var list DomainFIPAclTemplateEntriesList
	err := bambou.CurrentSession().FetchChildren(o, DomainFIPAclTemplateEntryIdentity, &list, info)
	return list, err
}

// CreateDomainFIPAclTemplateEntry creates a new child DomainFIPAclTemplateEntry under the DomainFIPAclTemplate
func (o *DomainFIPAclTemplate) CreateDomainFIPAclTemplateEntry(child *DomainFIPAclTemplateEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DomainFIPAclTemplate
func (o *DomainFIPAclTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DomainFIPAclTemplate
func (o *DomainFIPAclTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
