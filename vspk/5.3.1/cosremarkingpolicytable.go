/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// COSRemarkingPolicyTableIdentity represents the Identity of the object
var COSRemarkingPolicyTableIdentity = bambou.Identity{
	Name:     "cosremarkingpolicytable",
	Category: "cosremarkingpolicytables",
}

// COSRemarkingPolicyTablesList represents a list of COSRemarkingPolicyTables
type COSRemarkingPolicyTablesList []*COSRemarkingPolicyTable

// COSRemarkingPolicyTablesAncestor is the interface that an ancestor of a COSRemarkingPolicyTable must implement.
// An Ancestor is defined as an entity that has COSRemarkingPolicyTable as a descendant.
// An Ancestor can get a list of its child COSRemarkingPolicyTables, but not necessarily create one.
type COSRemarkingPolicyTablesAncestor interface {
	COSRemarkingPolicyTables(*bambou.FetchingInfo) (COSRemarkingPolicyTablesList, *bambou.Error)
}

// COSRemarkingPolicyTablesParent is the interface that a parent of a COSRemarkingPolicyTable must implement.
// A Parent is defined as an entity that has COSRemarkingPolicyTable as a child.
// A Parent is an Ancestor which can create a COSRemarkingPolicyTable.
type COSRemarkingPolicyTablesParent interface {
	COSRemarkingPolicyTablesAncestor
	CreateCOSRemarkingPolicyTable(*COSRemarkingPolicyTable) *bambou.Error
}

// COSRemarkingPolicyTable represents the model of a cosremarkingpolicytable
type COSRemarkingPolicyTable struct {
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

// NewCOSRemarkingPolicyTable returns a new *COSRemarkingPolicyTable
func NewCOSRemarkingPolicyTable() *COSRemarkingPolicyTable {

	return &COSRemarkingPolicyTable{}
}

// Identity returns the Identity of the object.
func (o *COSRemarkingPolicyTable) Identity() bambou.Identity {

	return COSRemarkingPolicyTableIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *COSRemarkingPolicyTable) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *COSRemarkingPolicyTable) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the COSRemarkingPolicyTable from the server
func (o *COSRemarkingPolicyTable) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the COSRemarkingPolicyTable into the server
func (o *COSRemarkingPolicyTable) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the COSRemarkingPolicyTable from the server
func (o *COSRemarkingPolicyTable) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the COSRemarkingPolicyTable
func (o *COSRemarkingPolicyTable) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the COSRemarkingPolicyTable
func (o *COSRemarkingPolicyTable) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the COSRemarkingPolicyTable
func (o *COSRemarkingPolicyTable) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the COSRemarkingPolicyTable
func (o *COSRemarkingPolicyTable) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// COSRemarkingPolicies retrieves the list of child COSRemarkingPolicies of the COSRemarkingPolicyTable
func (o *COSRemarkingPolicyTable) COSRemarkingPolicies(info *bambou.FetchingInfo) (COSRemarkingPoliciesList, *bambou.Error) {

	var list COSRemarkingPoliciesList
	err := bambou.CurrentSession().FetchChildren(o, COSRemarkingPolicyIdentity, &list, info)
	return list, err
}

// CreateCOSRemarkingPolicy creates a new child COSRemarkingPolicy under the COSRemarkingPolicyTable
func (o *COSRemarkingPolicyTable) CreateCOSRemarkingPolicy(child *COSRemarkingPolicy) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
