/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IngressACLEntryTemplateIdentity represents the Identity of the object
var IngressACLEntryTemplateIdentity = bambou.Identity{
	Name:     "ingressaclentrytemplate",
	Category: "ingressaclentrytemplates",
}

// IngressACLEntryTemplatesList represents a list of IngressACLEntryTemplates
type IngressACLEntryTemplatesList []*IngressACLEntryTemplate

// IngressACLEntryTemplatesAncestor is the interface that an ancestor of a IngressACLEntryTemplate must implement.
// An Ancestor is defined as an entity that has IngressACLEntryTemplate as a descendant.
// An Ancestor can get a list of its child IngressACLEntryTemplates, but not necessarily create one.
type IngressACLEntryTemplatesAncestor interface {
	IngressACLEntryTemplates(*bambou.FetchingInfo) (IngressACLEntryTemplatesList, *bambou.Error)
}

// IngressACLEntryTemplatesParent is the interface that a parent of a IngressACLEntryTemplate must implement.
// A Parent is defined as an entity that has IngressACLEntryTemplate as a child.
// A Parent is an Ancestor which can create a IngressACLEntryTemplate.
type IngressACLEntryTemplatesParent interface {
	IngressACLEntryTemplatesAncestor
	CreateIngressACLEntryTemplate(*IngressACLEntryTemplate) *bambou.Error
}

// IngressACLEntryTemplate represents the model of a ingressaclentrytemplate
type IngressACLEntryTemplate struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewIngressACLEntryTemplate returns a new *IngressACLEntryTemplate
func NewIngressACLEntryTemplate() *IngressACLEntryTemplate {

	return &IngressACLEntryTemplate{}
}

// Identity returns the Identity of the object.
func (o *IngressACLEntryTemplate) Identity() bambou.Identity {

	return IngressACLEntryTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressACLEntryTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressACLEntryTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressACLEntryTemplate from the server
func (o *IngressACLEntryTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressACLEntryTemplate into the server
func (o *IngressACLEntryTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressACLEntryTemplate from the server
func (o *IngressACLEntryTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}
