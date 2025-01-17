/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SiteInfoIdentity represents the Identity of the object
var SiteInfoIdentity = bambou.Identity{
	Name:     "site",
	Category: "sites",
}

// SiteInfosList represents a list of SiteInfos
type SiteInfosList []*SiteInfo

// SiteInfosAncestor is the interface that an ancestor of a SiteInfo must implement.
// An Ancestor is defined as an entity that has SiteInfo as a descendant.
// An Ancestor can get a list of its child SiteInfos, but not necessarily create one.
type SiteInfosAncestor interface {
	SiteInfos(*bambou.FetchingInfo) (SiteInfosList, *bambou.Error)
}

// SiteInfosParent is the interface that a parent of a SiteInfo must implement.
// A Parent is defined as an entity that has SiteInfo as a child.
// A Parent is an Ancestor which can create a SiteInfo.
type SiteInfosParent interface {
	SiteInfosAncestor
	CreateSiteInfo(*SiteInfo) *bambou.Error
}

// SiteInfo represents the model of a site
type SiteInfo struct {
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	Name           string `json:"name,omitempty"`
	LastUpdatedBy  string `json:"lastUpdatedBy,omitempty"`
	Address        string `json:"address,omitempty"`
	Description    string `json:"description,omitempty"`
	SiteIdentifier string `json:"siteIdentifier,omitempty"`
	XmppDomain     string `json:"xmppDomain,omitempty"`
	EntityScope    string `json:"entityScope,omitempty"`
	ExternalID     string `json:"externalID,omitempty"`
}

// NewSiteInfo returns a new *SiteInfo
func NewSiteInfo() *SiteInfo {

	return &SiteInfo{}
}

// Identity returns the Identity of the object.
func (o *SiteInfo) Identity() bambou.Identity {

	return SiteInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SiteInfo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SiteInfo) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SiteInfo from the server
func (o *SiteInfo) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SiteInfo into the server
func (o *SiteInfo) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SiteInfo from the server
func (o *SiteInfo) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the SiteInfo
func (o *SiteInfo) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SiteInfo
func (o *SiteInfo) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SiteInfo
func (o *SiteInfo) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SiteInfo
func (o *SiteInfo) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
