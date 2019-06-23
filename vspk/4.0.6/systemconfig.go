/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SystemConfigIdentity represents the Identity of the object
var SystemConfigIdentity = bambou.Identity{
	Name:     "systemconfig",
	Category: "systemconfigs",
}

// SystemConfigsList represents a list of SystemConfigs
type SystemConfigsList []*SystemConfig

// SystemConfigsAncestor is the interface that an ancestor of a SystemConfig must implement.
// An Ancestor is defined as an entity that has SystemConfig as a descendant.
// An Ancestor can get a list of its child SystemConfigs, but not necessarily create one.
type SystemConfigsAncestor interface {
	SystemConfigs(*bambou.FetchingInfo) (SystemConfigsList, *bambou.Error)
}

// SystemConfigsParent is the interface that a parent of a SystemConfig must implement.
// A Parent is defined as an entity that has SystemConfig as a child.
// A Parent is an Ancestor which can create a SystemConfig.
type SystemConfigsParent interface {
	SystemConfigsAncestor
	CreateSystemConfig(*SystemConfig) *bambou.Error
}

// SystemConfig represents the model of a systemconfig
type SystemConfig struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewSystemConfig returns a new *SystemConfig
func NewSystemConfig() *SystemConfig {

	return &SystemConfig{}
}

// Identity returns the Identity of the object.
func (o *SystemConfig) Identity() bambou.Identity {

	return SystemConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SystemConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SystemConfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SystemConfig from the server
func (o *SystemConfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SystemConfig into the server
func (o *SystemConfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SystemConfig from the server
func (o *SystemConfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the SystemConfig
func (o *SystemConfig) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SystemConfig
func (o *SystemConfig) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SystemConfig
func (o *SystemConfig) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SystemConfig
func (o *SystemConfig) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
