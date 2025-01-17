/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EgressAdvFwdTemplateIdentity represents the Identity of the object
var EgressAdvFwdTemplateIdentity = bambou.Identity{
	Name:     "egressadvfwdtemplate",
	Category: "egressadvfwdtemplates",
}

// EgressAdvFwdTemplatesList represents a list of EgressAdvFwdTemplates
type EgressAdvFwdTemplatesList []*EgressAdvFwdTemplate

// EgressAdvFwdTemplatesAncestor is the interface that an ancestor of a EgressAdvFwdTemplate must implement.
// An Ancestor is defined as an entity that has EgressAdvFwdTemplate as a descendant.
// An Ancestor can get a list of its child EgressAdvFwdTemplates, but not necessarily create one.
type EgressAdvFwdTemplatesAncestor interface {
	EgressAdvFwdTemplates(*bambou.FetchingInfo) (EgressAdvFwdTemplatesList, *bambou.Error)
}

// EgressAdvFwdTemplatesParent is the interface that a parent of a EgressAdvFwdTemplate must implement.
// A Parent is defined as an entity that has EgressAdvFwdTemplate as a child.
// A Parent is an Ancestor which can create a EgressAdvFwdTemplate.
type EgressAdvFwdTemplatesParent interface {
	EgressAdvFwdTemplatesAncestor
	CreateEgressAdvFwdTemplate(*EgressAdvFwdTemplate) *bambou.Error
}

// EgressAdvFwdTemplate represents the model of a egressadvfwdtemplate
type EgressAdvFwdTemplate struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	Name                   string `json:"name,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty"`
	Active                 bool   `json:"active"`
	Description            string `json:"description,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	PolicyState            string `json:"policyState,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	PriorityType           string `json:"priorityType,omitempty"`
	AssociatedLiveEntityID string `json:"associatedLiveEntityID,omitempty"`
	AutoGeneratePriority   bool   `json:"autoGeneratePriority"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewEgressAdvFwdTemplate returns a new *EgressAdvFwdTemplate
func NewEgressAdvFwdTemplate() *EgressAdvFwdTemplate {

	return &EgressAdvFwdTemplate{}
}

// Identity returns the Identity of the object.
func (o *EgressAdvFwdTemplate) Identity() bambou.Identity {

	return EgressAdvFwdTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressAdvFwdTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressAdvFwdTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressAdvFwdTemplate from the server
func (o *EgressAdvFwdTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressAdvFwdTemplate into the server
func (o *EgressAdvFwdTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressAdvFwdTemplate from the server
func (o *EgressAdvFwdTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EgressAdvFwdTemplate
func (o *EgressAdvFwdTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressAdvFwdTemplate
func (o *EgressAdvFwdTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EgressAdvFwdEntryTemplates retrieves the list of child EgressAdvFwdEntryTemplates of the EgressAdvFwdTemplate
func (o *EgressAdvFwdTemplate) EgressAdvFwdEntryTemplates(info *bambou.FetchingInfo) (EgressAdvFwdEntryTemplatesList, *bambou.Error) {

	var list EgressAdvFwdEntryTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, EgressAdvFwdEntryTemplateIdentity, &list, info)
	return list, err
}

// CreateEgressAdvFwdEntryTemplate creates a new child EgressAdvFwdEntryTemplate under the EgressAdvFwdTemplate
func (o *EgressAdvFwdTemplate) CreateEgressAdvFwdEntryTemplate(child *EgressAdvFwdEntryTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressAdvFwdTemplate
func (o *EgressAdvFwdTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressAdvFwdTemplate
func (o *EgressAdvFwdTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
