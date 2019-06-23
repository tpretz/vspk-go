/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DomainFIPAclTemplateEntryIdentity represents the Identity of the object
var DomainFIPAclTemplateEntryIdentity = bambou.Identity{
	Name:     "egressdomainfloatingipaclentrytemplate",
	Category: "egressdomainfloatingipaclentrytemplates",
}

// DomainFIPAclTemplateEntriesList represents a list of DomainFIPAclTemplateEntries
type DomainFIPAclTemplateEntriesList []*DomainFIPAclTemplateEntry

// DomainFIPAclTemplateEntriesAncestor is the interface that an ancestor of a DomainFIPAclTemplateEntry must implement.
// An Ancestor is defined as an entity that has DomainFIPAclTemplateEntry as a descendant.
// An Ancestor can get a list of its child DomainFIPAclTemplateEntries, but not necessarily create one.
type DomainFIPAclTemplateEntriesAncestor interface {
	DomainFIPAclTemplateEntries(*bambou.FetchingInfo) (DomainFIPAclTemplateEntriesList, *bambou.Error)
}

// DomainFIPAclTemplateEntriesParent is the interface that a parent of a DomainFIPAclTemplateEntry must implement.
// A Parent is defined as an entity that has DomainFIPAclTemplateEntry as a child.
// A Parent is an Ancestor which can create a DomainFIPAclTemplateEntry.
type DomainFIPAclTemplateEntriesParent interface {
	DomainFIPAclTemplateEntriesAncestor
	CreateDomainFIPAclTemplateEntry(*DomainFIPAclTemplateEntry) *bambou.Error
}

// DomainFIPAclTemplateEntry represents the model of a egressdomainfloatingipaclentrytemplate
type DomainFIPAclTemplateEntry struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewDomainFIPAclTemplateEntry returns a new *DomainFIPAclTemplateEntry
func NewDomainFIPAclTemplateEntry() *DomainFIPAclTemplateEntry {

	return &DomainFIPAclTemplateEntry{}
}

// Identity returns the Identity of the object.
func (o *DomainFIPAclTemplateEntry) Identity() bambou.Identity {

	return DomainFIPAclTemplateEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DomainFIPAclTemplateEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DomainFIPAclTemplateEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DomainFIPAclTemplateEntry from the server
func (o *DomainFIPAclTemplateEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DomainFIPAclTemplateEntry into the server
func (o *DomainFIPAclTemplateEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DomainFIPAclTemplateEntry from the server
func (o *DomainFIPAclTemplateEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DomainFIPAclTemplateEntry
func (o *DomainFIPAclTemplateEntry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DomainFIPAclTemplateEntry
func (o *DomainFIPAclTemplateEntry) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DomainFIPAclTemplateEntry
func (o *DomainFIPAclTemplateEntry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DomainFIPAclTemplateEntry
func (o *DomainFIPAclTemplateEntry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
