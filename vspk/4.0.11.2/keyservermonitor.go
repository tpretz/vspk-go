/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// KeyServerMonitorIdentity represents the Identity of the object
var KeyServerMonitorIdentity = bambou.Identity{
	Name:     "keyservermonitor",
	Category: "keyservermonitors",
}

// KeyServerMonitorsList represents a list of KeyServerMonitors
type KeyServerMonitorsList []*KeyServerMonitor

// KeyServerMonitorsAncestor is the interface that an ancestor of a KeyServerMonitor must implement.
// An Ancestor is defined as an entity that has KeyServerMonitor as a descendant.
// An Ancestor can get a list of its child KeyServerMonitors, but not necessarily create one.
type KeyServerMonitorsAncestor interface {
	KeyServerMonitors(*bambou.FetchingInfo) (KeyServerMonitorsList, *bambou.Error)
}

// KeyServerMonitorsParent is the interface that a parent of a KeyServerMonitor must implement.
// A Parent is defined as an entity that has KeyServerMonitor as a child.
// A Parent is an Ancestor which can create a KeyServerMonitor.
type KeyServerMonitorsParent interface {
	KeyServerMonitorsAncestor
	CreateKeyServerMonitor(*KeyServerMonitor) *bambou.Error
}

// KeyServerMonitor represents the model of a keyservermonitor
type KeyServerMonitor struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	LastUpdateTime                     int    `json:"lastUpdateTime,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	GatewaySecuredDataRecordCount      int    `json:"gatewaySecuredDataRecordCount,omitempty"`
	KeyserverMonitorEncryptedSEKCount  int    `json:"keyserverMonitorEncryptedSEKCount,omitempty"`
	KeyserverMonitorEncryptedSeedCount int    `json:"keyserverMonitorEncryptedSeedCount,omitempty"`
	KeyserverMonitorSEKCount           int    `json:"keyserverMonitorSEKCount,omitempty"`
	KeyserverMonitorSeedCount          int    `json:"keyserverMonitorSeedCount,omitempty"`
	EnterpriseSecuredDataRecordCount   int    `json:"enterpriseSecuredDataRecordCount,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
}

// NewKeyServerMonitor returns a new *KeyServerMonitor
func NewKeyServerMonitor() *KeyServerMonitor {

	return &KeyServerMonitor{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitor) Identity() bambou.Identity {

	return KeyServerMonitorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitor from the server
func (o *KeyServerMonitor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitor into the server
func (o *KeyServerMonitor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitor from the server
func (o *KeyServerMonitor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitor
func (o *KeyServerMonitor) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitor
func (o *KeyServerMonitor) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// KeyServerMonitorEncryptedSeeds retrieves the list of child KeyServerMonitorEncryptedSeeds of the KeyServerMonitor
func (o *KeyServerMonitor) KeyServerMonitorEncryptedSeeds(info *bambou.FetchingInfo) (KeyServerMonitorEncryptedSeedsList, *bambou.Error) {

	var list KeyServerMonitorEncryptedSeedsList
	err := bambou.CurrentSession().FetchChildren(o, KeyServerMonitorEncryptedSeedIdentity, &list, info)
	return list, err
}

// CreateKeyServerMonitorEncryptedSeed creates a new child KeyServerMonitorEncryptedSeed under the KeyServerMonitor
func (o *KeyServerMonitor) CreateKeyServerMonitorEncryptedSeed(child *KeyServerMonitorEncryptedSeed) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// KeyServerMonitorSeeds retrieves the list of child KeyServerMonitorSeeds of the KeyServerMonitor
func (o *KeyServerMonitor) KeyServerMonitorSeeds(info *bambou.FetchingInfo) (KeyServerMonitorSeedsList, *bambou.Error) {

	var list KeyServerMonitorSeedsList
	err := bambou.CurrentSession().FetchChildren(o, KeyServerMonitorSeedIdentity, &list, info)
	return list, err
}

// CreateKeyServerMonitorSeed creates a new child KeyServerMonitorSeed under the KeyServerMonitor
func (o *KeyServerMonitor) CreateKeyServerMonitorSeed(child *KeyServerMonitorSeed) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// KeyServerMonitorSEKs retrieves the list of child KeyServerMonitorSEKs of the KeyServerMonitor
func (o *KeyServerMonitor) KeyServerMonitorSEKs(info *bambou.FetchingInfo) (KeyServerMonitorSEKsList, *bambou.Error) {

	var list KeyServerMonitorSEKsList
	err := bambou.CurrentSession().FetchChildren(o, KeyServerMonitorSEKIdentity, &list, info)
	return list, err
}

// CreateKeyServerMonitorSEK creates a new child KeyServerMonitorSEK under the KeyServerMonitor
func (o *KeyServerMonitor) CreateKeyServerMonitorSEK(child *KeyServerMonitorSEK) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitor
func (o *KeyServerMonitor) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitor
func (o *KeyServerMonitor) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
