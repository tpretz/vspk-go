/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IngressExternalServiceTemplateIdentity represents the Identity of the object
var IngressExternalServiceTemplateIdentity = bambou.Identity{
	Name:     "ingressexternalservicetemplate",
	Category: "ingressexternalservicetemplates",
}

// IngressExternalServiceTemplatesList represents a list of IngressExternalServiceTemplates
type IngressExternalServiceTemplatesList []*IngressExternalServiceTemplate

// IngressExternalServiceTemplatesAncestor is the interface that an ancestor of a IngressExternalServiceTemplate must implement.
// An Ancestor is defined as an entity that has IngressExternalServiceTemplate as a descendant.
// An Ancestor can get a list of its child IngressExternalServiceTemplates, but not necessarily create one.
type IngressExternalServiceTemplatesAncestor interface {
	IngressExternalServiceTemplates(*bambou.FetchingInfo) (IngressExternalServiceTemplatesList, *bambou.Error)
}

// IngressExternalServiceTemplatesParent is the interface that a parent of a IngressExternalServiceTemplate must implement.
// A Parent is defined as an entity that has IngressExternalServiceTemplate as a child.
// A Parent is an Ancestor which can create a IngressExternalServiceTemplate.
type IngressExternalServiceTemplatesParent interface {
	IngressExternalServiceTemplatesAncestor
	CreateIngressExternalServiceTemplate(*IngressExternalServiceTemplate) *bambou.Error
}

// IngressExternalServiceTemplate represents the model of a ingressexternalservicetemplate
type IngressExternalServiceTemplate struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	Name                   string `json:"name,omitempty"`
	Active                 bool   `json:"active"`
	Description            string `json:"description,omitempty"`
	EntityScope            string `json:"entityScope,omitempty"`
	PolicyState            string `json:"policyState,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	PriorityType           string `json:"priorityType,omitempty"`
	AssociatedLiveEntityID string `json:"associatedLiveEntityID,omitempty"`
	ExternalID             string `json:"externalID,omitempty"`
}

// NewIngressExternalServiceTemplate returns a new *IngressExternalServiceTemplate
func NewIngressExternalServiceTemplate() *IngressExternalServiceTemplate {

	return &IngressExternalServiceTemplate{}
}

// Identity returns the Identity of the object.
func (o *IngressExternalServiceTemplate) Identity() bambou.Identity {

	return IngressExternalServiceTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressExternalServiceTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressExternalServiceTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressExternalServiceTemplate from the server
func (o *IngressExternalServiceTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressExternalServiceTemplate into the server
func (o *IngressExternalServiceTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressExternalServiceTemplate from the server
func (o *IngressExternalServiceTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IngressExternalServiceTemplate
func (o *IngressExternalServiceTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IngressExternalServiceTemplate
func (o *IngressExternalServiceTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressExternalServiceTemplate
func (o *IngressExternalServiceTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressExternalServiceTemplate
func (o *IngressExternalServiceTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IngressExternalServiceTemplateEntries retrieves the list of child IngressExternalServiceTemplateEntries of the IngressExternalServiceTemplate
func (o *IngressExternalServiceTemplate) IngressExternalServiceTemplateEntries(info *bambou.FetchingInfo) (IngressExternalServiceTemplateEntriesList, *bambou.Error) {

	var list IngressExternalServiceTemplateEntriesList
	err := bambou.CurrentSession().FetchChildren(o, IngressExternalServiceTemplateEntryIdentity, &list, info)
	return list, err
}

// CreateIngressExternalServiceTemplateEntry creates a new child IngressExternalServiceTemplateEntry under the IngressExternalServiceTemplate
func (o *IngressExternalServiceTemplate) CreateIngressExternalServiceTemplateEntry(child *IngressExternalServiceTemplateEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
