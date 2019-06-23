/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PSPATMapIdentity represents the Identity of the object
var PSPATMapIdentity = bambou.Identity{
	Name:     "pspatmap",
	Category: "pspatmaps",
}

// PSPATMapsList represents a list of PSPATMaps
type PSPATMapsList []*PSPATMap

// PSPATMapsAncestor is the interface that an ancestor of a PSPATMap must implement.
// An Ancestor is defined as an entity that has PSPATMap as a descendant.
// An Ancestor can get a list of its child PSPATMaps, but not necessarily create one.
type PSPATMapsAncestor interface {
	PSPATMaps(*bambou.FetchingInfo) (PSPATMapsList, *bambou.Error)
}

// PSPATMapsParent is the interface that a parent of a PSPATMap must implement.
// A Parent is defined as an entity that has PSPATMap as a child.
// A Parent is an Ancestor which can create a PSPATMap.
type PSPATMapsParent interface {
	PSPATMapsAncestor
	CreatePSPATMap(*PSPATMap) *bambou.Error
}

// PSPATMap represents the model of a pspatmap
type PSPATMap struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewPSPATMap returns a new *PSPATMap
func NewPSPATMap() *PSPATMap {

	return &PSPATMap{}
}

// Identity returns the Identity of the object.
func (o *PSPATMap) Identity() bambou.Identity {

	return PSPATMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PSPATMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PSPATMap) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PSPATMap from the server
func (o *PSPATMap) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PSPATMap into the server
func (o *PSPATMap) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PSPATMap from the server
func (o *PSPATMap) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PSPATMap
func (o *PSPATMap) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PSPATMap
func (o *PSPATMap) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PSPATMap
func (o *PSPATMap) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PSPATMap
func (o *PSPATMap) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
