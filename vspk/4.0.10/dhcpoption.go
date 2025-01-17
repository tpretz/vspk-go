/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DHCPOptionIdentity represents the Identity of the object
var DHCPOptionIdentity = bambou.Identity{
	Name:     "dhcpoption",
	Category: "dhcpoptions",
}

// DHCPOptionsList represents a list of DHCPOptions
type DHCPOptionsList []*DHCPOption

// DHCPOptionsAncestor is the interface that an ancestor of a DHCPOption must implement.
// An Ancestor is defined as an entity that has DHCPOption as a descendant.
// An Ancestor can get a list of its child DHCPOptions, but not necessarily create one.
type DHCPOptionsAncestor interface {
	DHCPOptions(*bambou.FetchingInfo) (DHCPOptionsList, *bambou.Error)
}

// DHCPOptionsParent is the interface that a parent of a DHCPOption must implement.
// A Parent is defined as an entity that has DHCPOption as a child.
// A Parent is an Ancestor which can create a DHCPOption.
type DHCPOptionsParent interface {
	DHCPOptionsAncestor
	CreateDHCPOption(*DHCPOption) *bambou.Error
}

// DHCPOption represents the model of a dhcpoption
type DHCPOption struct {
	ID            string        `json:"ID,omitempty"`
	ParentID      string        `json:"parentID,omitempty"`
	ParentType    string        `json:"parentType,omitempty"`
	Owner         string        `json:"owner,omitempty"`
	Value         string        `json:"value,omitempty"`
	LastUpdatedBy string        `json:"lastUpdatedBy,omitempty"`
	ActualType    int           `json:"actualType,omitempty"`
	ActualValues  []interface{} `json:"actualValues,omitempty"`
	Length        string        `json:"length,omitempty"`
	EntityScope   string        `json:"entityScope,omitempty"`
	ExternalID    string        `json:"externalID,omitempty"`
	Type          string        `json:"type,omitempty"`
}

// NewDHCPOption returns a new *DHCPOption
func NewDHCPOption() *DHCPOption {

	return &DHCPOption{}
}

// Identity returns the Identity of the object.
func (o *DHCPOption) Identity() bambou.Identity {

	return DHCPOptionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DHCPOption) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DHCPOption) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DHCPOption from the server
func (o *DHCPOption) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DHCPOption into the server
func (o *DHCPOption) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DHCPOption from the server
func (o *DHCPOption) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the DHCPOption
func (o *DHCPOption) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the DHCPOption
func (o *DHCPOption) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the DHCPOption
func (o *DHCPOption) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the DHCPOption
func (o *DHCPOption) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the DHCPOption
func (o *DHCPOption) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
