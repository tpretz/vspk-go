/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// AddressRangeIdentity represents the Identity of the object
var AddressRangeIdentity = bambou.Identity{
	Name:     "addressrange",
	Category: "addressranges",
}

// AddressRangesList represents a list of AddressRanges
type AddressRangesList []*AddressRange

// AddressRangesAncestor is the interface that an ancestor of a AddressRange must implement.
// An Ancestor is defined as an entity that has AddressRange as a descendant.
// An Ancestor can get a list of its child AddressRanges, but not necessarily create one.
type AddressRangesAncestor interface {
	AddressRanges(*bambou.FetchingInfo) (AddressRangesList, *bambou.Error)
}

// AddressRangesParent is the interface that a parent of a AddressRange must implement.
// A Parent is defined as an entity that has AddressRange as a child.
// A Parent is an Ancestor which can create a AddressRange.
type AddressRangesParent interface {
	AddressRangesAncestor
	CreateAddressRange(*AddressRange) *bambou.Error
}

// AddressRange represents the model of a addressrange
type AddressRange struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	DHCPPoolType  string `json:"DHCPPoolType,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	MaxAddress    string `json:"maxAddress,omitempty"`
	MinAddress    string `json:"minAddress,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewAddressRange returns a new *AddressRange
func NewAddressRange() *AddressRange {

	return &AddressRange{}
}

// Identity returns the Identity of the object.
func (o *AddressRange) Identity() bambou.Identity {

	return AddressRangeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AddressRange) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AddressRange) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AddressRange from the server
func (o *AddressRange) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AddressRange into the server
func (o *AddressRange) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AddressRange from the server
func (o *AddressRange) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the AddressRange
func (o *AddressRange) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the AddressRange
func (o *AddressRange) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the AddressRange
func (o *AddressRange) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the AddressRange
func (o *AddressRange) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the AddressRange
func (o *AddressRange) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
