/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSPortStaticConfigurationIdentity represents the Identity of the object
var NSPortStaticConfigurationIdentity = bambou.Identity{
	Name:     "nsportstaticconfiguration",
	Category: "nsportstaticconfigurations",
}

// NSPortStaticConfigurationsList represents a list of NSPortStaticConfigurations
type NSPortStaticConfigurationsList []*NSPortStaticConfiguration

// NSPortStaticConfigurationsAncestor is the interface that an ancestor of a NSPortStaticConfiguration must implement.
// An Ancestor is defined as an entity that has NSPortStaticConfiguration as a descendant.
// An Ancestor can get a list of its child NSPortStaticConfigurations, but not necessarily create one.
type NSPortStaticConfigurationsAncestor interface {
	NSPortStaticConfigurations(*bambou.FetchingInfo) (NSPortStaticConfigurationsList, *bambou.Error)
}

// NSPortStaticConfigurationsParent is the interface that a parent of a NSPortStaticConfiguration must implement.
// A Parent is defined as an entity that has NSPortStaticConfiguration as a child.
// A Parent is an Ancestor which can create a NSPortStaticConfiguration.
type NSPortStaticConfigurationsParent interface {
	NSPortStaticConfigurationsAncestor
	CreateNSPortStaticConfiguration(*NSPortStaticConfiguration) *bambou.Error
}

// NSPortStaticConfiguration represents the model of a nsportstaticconfiguration
type NSPortStaticConfiguration struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	DNSAddress    string `json:"DNSAddress,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Gateway       string `json:"gateway,omitempty"`
	Address       string `json:"address,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
	Enabled       bool   `json:"enabled"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewNSPortStaticConfiguration returns a new *NSPortStaticConfiguration
func NewNSPortStaticConfiguration() *NSPortStaticConfiguration {

	return &NSPortStaticConfiguration{}
}

// Identity returns the Identity of the object.
func (o *NSPortStaticConfiguration) Identity() bambou.Identity {

	return NSPortStaticConfigurationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSPortStaticConfiguration) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSPortStaticConfiguration) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSPortStaticConfiguration from the server
func (o *NSPortStaticConfiguration) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSPortStaticConfiguration into the server
func (o *NSPortStaticConfiguration) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSPortStaticConfiguration from the server
func (o *NSPortStaticConfiguration) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NSPortStaticConfiguration
func (o *NSPortStaticConfiguration) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NSPortStaticConfiguration
func (o *NSPortStaticConfiguration) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NSPortStaticConfiguration
func (o *NSPortStaticConfiguration) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NSPortStaticConfiguration
func (o *NSPortStaticConfiguration) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
