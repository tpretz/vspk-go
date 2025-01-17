/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DSCPForwardingClassTableIdentity represents the Identity of the object
var DSCPForwardingClassTableIdentity = bambou.Identity{
	Name:     "dscpforwardingclasstable",
	Category: "dscpforwardingclasstables",
}

// DSCPForwardingClassTablesList represents a list of DSCPForwardingClassTables
type DSCPForwardingClassTablesList []*DSCPForwardingClassTable

// DSCPForwardingClassTablesAncestor is the interface that an ancestor of a DSCPForwardingClassTable must implement.
// An Ancestor is defined as an entity that has DSCPForwardingClassTable as a descendant.
// An Ancestor can get a list of its child DSCPForwardingClassTables, but not necessarily create one.
type DSCPForwardingClassTablesAncestor interface {
	DSCPForwardingClassTables(*bambou.FetchingInfo) (DSCPForwardingClassTablesList, *bambou.Error)
}

// DSCPForwardingClassTablesParent is the interface that a parent of a DSCPForwardingClassTable must implement.
// A Parent is defined as an entity that has DSCPForwardingClassTable as a child.
// A Parent is an Ancestor which can create a DSCPForwardingClassTable.
type DSCPForwardingClassTablesParent interface {
	DSCPForwardingClassTablesAncestor
	CreateDSCPForwardingClassTable(*DSCPForwardingClassTable) *bambou.Error
}

// DSCPForwardingClassTable represents the model of a dscpforwardingclasstable
type DSCPForwardingClassTable struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewDSCPForwardingClassTable returns a new *DSCPForwardingClassTable
func NewDSCPForwardingClassTable() *DSCPForwardingClassTable {

	return &DSCPForwardingClassTable{}
}

// Identity returns the Identity of the object.
func (o *DSCPForwardingClassTable) Identity() bambou.Identity {

	return DSCPForwardingClassTableIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DSCPForwardingClassTable) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DSCPForwardingClassTable) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DSCPForwardingClassTable from the server
func (o *DSCPForwardingClassTable) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DSCPForwardingClassTable into the server
func (o *DSCPForwardingClassTable) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DSCPForwardingClassTable from the server
func (o *DSCPForwardingClassTable) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DSCPForwardingClassTable
func (o *DSCPForwardingClassTable) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DSCPForwardingClassTable
func (o *DSCPForwardingClassTable) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DSCPForwardingClassTable
func (o *DSCPForwardingClassTable) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DSCPForwardingClassTable
func (o *DSCPForwardingClassTable) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// DSCPForwardingClassMappings retrieves the list of child DSCPForwardingClassMappings of the DSCPForwardingClassTable
func (o *DSCPForwardingClassTable) DSCPForwardingClassMappings(info *bambou.FetchingInfo) (DSCPForwardingClassMappingsList, *bambou.Error) {

	var list DSCPForwardingClassMappingsList
	err := bambou.CurrentSession().FetchChildren(o, DSCPForwardingClassMappingIdentity, &list, info)
	return list, err
}

// CreateDSCPForwardingClassMapping creates a new child DSCPForwardingClassMapping under the DSCPForwardingClassTable
func (o *DSCPForwardingClassTable) CreateDSCPForwardingClassMapping(child *DSCPForwardingClassMapping) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
