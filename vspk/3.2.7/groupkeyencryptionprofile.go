/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GroupKeyEncryptionProfileIdentity represents the Identity of the object
var GroupKeyEncryptionProfileIdentity = bambou.Identity{
	Name:     "groupkeyencryptionprofile",
	Category: "groupkeyencryptionprofiles",
}

// GroupKeyEncryptionProfilesList represents a list of GroupKeyEncryptionProfiles
type GroupKeyEncryptionProfilesList []*GroupKeyEncryptionProfile

// GroupKeyEncryptionProfilesAncestor is the interface that an ancestor of a GroupKeyEncryptionProfile must implement.
// An Ancestor is defined as an entity that has GroupKeyEncryptionProfile as a descendant.
// An Ancestor can get a list of its child GroupKeyEncryptionProfiles, but not necessarily create one.
type GroupKeyEncryptionProfilesAncestor interface {
	GroupKeyEncryptionProfiles(*bambou.FetchingInfo) (GroupKeyEncryptionProfilesList, *bambou.Error)
}

// GroupKeyEncryptionProfilesParent is the interface that a parent of a GroupKeyEncryptionProfile must implement.
// A Parent is defined as an entity that has GroupKeyEncryptionProfile as a child.
// A Parent is an Ancestor which can create a GroupKeyEncryptionProfile.
type GroupKeyEncryptionProfilesParent interface {
	GroupKeyEncryptionProfilesAncestor
	CreateGroupKeyEncryptionProfile(*GroupKeyEncryptionProfile) *bambou.Error
}

// GroupKeyEncryptionProfile represents the model of a groupkeyencryptionprofile
type GroupKeyEncryptionProfile struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewGroupKeyEncryptionProfile returns a new *GroupKeyEncryptionProfile
func NewGroupKeyEncryptionProfile() *GroupKeyEncryptionProfile {

	return &GroupKeyEncryptionProfile{}
}

// Identity returns the Identity of the object.
func (o *GroupKeyEncryptionProfile) Identity() bambou.Identity {

	return GroupKeyEncryptionProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GroupKeyEncryptionProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GroupKeyEncryptionProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GroupKeyEncryptionProfile from the server
func (o *GroupKeyEncryptionProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GroupKeyEncryptionProfile into the server
func (o *GroupKeyEncryptionProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GroupKeyEncryptionProfile from the server
func (o *GroupKeyEncryptionProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the GroupKeyEncryptionProfile
func (o *GroupKeyEncryptionProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the GroupKeyEncryptionProfile
func (o *GroupKeyEncryptionProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the GroupKeyEncryptionProfile
func (o *GroupKeyEncryptionProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the GroupKeyEncryptionProfile
func (o *GroupKeyEncryptionProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
