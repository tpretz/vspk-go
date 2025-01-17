/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// KeyServerMonitorSeedIdentity represents the Identity of the object
var KeyServerMonitorSeedIdentity = bambou.Identity{
	Name:     "keyservermonitorseed",
	Category: "keyservermonitorseeds",
}

// KeyServerMonitorSeedsList represents a list of KeyServerMonitorSeeds
type KeyServerMonitorSeedsList []*KeyServerMonitorSeed

// KeyServerMonitorSeedsAncestor is the interface that an ancestor of a KeyServerMonitorSeed must implement.
// An Ancestor is defined as an entity that has KeyServerMonitorSeed as a descendant.
// An Ancestor can get a list of its child KeyServerMonitorSeeds, but not necessarily create one.
type KeyServerMonitorSeedsAncestor interface {
	KeyServerMonitorSeeds(*bambou.FetchingInfo) (KeyServerMonitorSeedsList, *bambou.Error)
}

// KeyServerMonitorSeedsParent is the interface that a parent of a KeyServerMonitorSeed must implement.
// A Parent is defined as an entity that has KeyServerMonitorSeed as a child.
// A Parent is an Ancestor which can create a KeyServerMonitorSeed.
type KeyServerMonitorSeedsParent interface {
	KeyServerMonitorSeedsAncestor
	CreateKeyServerMonitorSeed(*KeyServerMonitorSeed) *bambou.Error
}

// KeyServerMonitorSeed represents the model of a keyservermonitorseed
type KeyServerMonitorSeed struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	SeedTrafficAuthenticationAlgorithm string `json:"seedTrafficAuthenticationAlgorithm,omitempty"`
	SeedTrafficEncryptionAlgorithm     string `json:"seedTrafficEncryptionAlgorithm,omitempty"`
	SeedTrafficEncryptionKeyLifetime   int    `json:"seedTrafficEncryptionKeyLifetime,omitempty"`
	Lifetime                           int    `json:"lifetime,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	CreationTime                       int    `json:"creationTime,omitempty"`
	StartTime                          int    `json:"startTime,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
}

// NewKeyServerMonitorSeed returns a new *KeyServerMonitorSeed
func NewKeyServerMonitorSeed() *KeyServerMonitorSeed {

	return &KeyServerMonitorSeed{}
}

// Identity returns the Identity of the object.
func (o *KeyServerMonitorSeed) Identity() bambou.Identity {

	return KeyServerMonitorSeedIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerMonitorSeed) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerMonitorSeed) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerMonitorSeed from the server
func (o *KeyServerMonitorSeed) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerMonitorSeed into the server
func (o *KeyServerMonitorSeed) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerMonitorSeed from the server
func (o *KeyServerMonitorSeed) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerMonitorSeed
func (o *KeyServerMonitorSeed) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerMonitorSeed
func (o *KeyServerMonitorSeed) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// KeyServerMonitorEncryptedSeeds retrieves the list of child KeyServerMonitorEncryptedSeeds of the KeyServerMonitorSeed
func (o *KeyServerMonitorSeed) KeyServerMonitorEncryptedSeeds(info *bambou.FetchingInfo) (KeyServerMonitorEncryptedSeedsList, *bambou.Error) {

	var list KeyServerMonitorEncryptedSeedsList
	err := bambou.CurrentSession().FetchChildren(o, KeyServerMonitorEncryptedSeedIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerMonitorSeed
func (o *KeyServerMonitorSeed) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerMonitorSeed
func (o *KeyServerMonitorSeed) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
