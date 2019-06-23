/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ZoneTemplateIdentity represents the Identity of the object
var ZoneTemplateIdentity = bambou.Identity{
	Name:     "zonetemplate",
	Category: "zonetemplates",
}

// ZoneTemplatesList represents a list of ZoneTemplates
type ZoneTemplatesList []*ZoneTemplate

// ZoneTemplatesAncestor is the interface that an ancestor of a ZoneTemplate must implement.
// An Ancestor is defined as an entity that has ZoneTemplate as a descendant.
// An Ancestor can get a list of its child ZoneTemplates, but not necessarily create one.
type ZoneTemplatesAncestor interface {
	ZoneTemplates(*bambou.FetchingInfo) (ZoneTemplatesList, *bambou.Error)
}

// ZoneTemplatesParent is the interface that a parent of a ZoneTemplate must implement.
// A Parent is defined as an entity that has ZoneTemplate as a child.
// A Parent is an Ancestor which can create a ZoneTemplate.
type ZoneTemplatesParent interface {
	ZoneTemplatesAncestor
	CreateZoneTemplate(*ZoneTemplate) *bambou.Error
}

// ZoneTemplate represents the model of a zonetemplate
type ZoneTemplate struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewZoneTemplate returns a new *ZoneTemplate
func NewZoneTemplate() *ZoneTemplate {

	return &ZoneTemplate{}
}

// Identity returns the Identity of the object.
func (o *ZoneTemplate) Identity() bambou.Identity {

	return ZoneTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ZoneTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ZoneTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ZoneTemplate from the server
func (o *ZoneTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ZoneTemplate into the server
func (o *ZoneTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ZoneTemplate from the server
func (o *ZoneTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ZoneTemplate
func (o *ZoneTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ZoneTemplate
func (o *ZoneTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ZoneTemplate
func (o *ZoneTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ZoneTemplate
func (o *ZoneTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// QOSs retrieves the list of child QOSs of the ZoneTemplate
func (o *ZoneTemplate) QOSs(info *bambou.FetchingInfo) (QOSsList, *bambou.Error) {

	var list QOSsList
	err := bambou.CurrentSession().FetchChildren(o, QOSIdentity, &list, info)
	return list, err
}

// CreateQOS creates a new child QOS under the ZoneTemplate
func (o *ZoneTemplate) CreateQOS(child *QOS) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// SubnetTemplates retrieves the list of child SubnetTemplates of the ZoneTemplate
func (o *ZoneTemplate) SubnetTemplates(info *bambou.FetchingInfo) (SubnetTemplatesList, *bambou.Error) {

	var list SubnetTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, SubnetTemplateIdentity, &list, info)
	return list, err
}

// CreateSubnetTemplate creates a new child SubnetTemplate under the ZoneTemplate
func (o *ZoneTemplate) CreateSubnetTemplate(child *SubnetTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the ZoneTemplate
func (o *ZoneTemplate) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
