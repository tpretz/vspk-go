/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CTranslationMapIdentity represents the Identity of the object
var CTranslationMapIdentity = bambou.Identity{
	Name:     "ctranslationmap",
	Category: "ctranslationmaps",
}

// CTranslationMapsList represents a list of CTranslationMaps
type CTranslationMapsList []*CTranslationMap

// CTranslationMapsAncestor is the interface that an ancestor of a CTranslationMap must implement.
// An Ancestor is defined as an entity that has CTranslationMap as a descendant.
// An Ancestor can get a list of its child CTranslationMaps, but not necessarily create one.
type CTranslationMapsAncestor interface {
	CTranslationMaps(*bambou.FetchingInfo) (CTranslationMapsList, *bambou.Error)
}

// CTranslationMapsParent is the interface that a parent of a CTranslationMap must implement.
// A Parent is defined as an entity that has CTranslationMap as a child.
// A Parent is an Ancestor which can create a CTranslationMap.
type CTranslationMapsParent interface {
	CTranslationMapsAncestor
	CreateCTranslationMap(*CTranslationMap) *bambou.Error
}

// CTranslationMap represents the model of a ctranslationmap
type CTranslationMap struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	MappingType        string `json:"mappingType,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	AssociatedDomainID string `json:"associatedDomainID,omitempty"`
	CustomerAliasIP    string `json:"customerAliasIP,omitempty"`
	CustomerIP         string `json:"customerIP,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewCTranslationMap returns a new *CTranslationMap
func NewCTranslationMap() *CTranslationMap {

	return &CTranslationMap{}
}

// Identity returns the Identity of the object.
func (o *CTranslationMap) Identity() bambou.Identity {

	return CTranslationMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CTranslationMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CTranslationMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CTranslationMap from the server
func (o *CTranslationMap) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CTranslationMap into the server
func (o *CTranslationMap) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CTranslationMap from the server
func (o *CTranslationMap) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the CTranslationMap
func (o *CTranslationMap) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the CTranslationMap
func (o *CTranslationMap) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the CTranslationMap
func (o *CTranslationMap) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the CTranslationMap
func (o *CTranslationMap) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
