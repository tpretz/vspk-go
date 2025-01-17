/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VSPIdentity represents the Identity of the object
var VSPIdentity = bambou.Identity{
	Name:     "vsp",
	Category: "vsps",
}

// VSPsList represents a list of VSPs
type VSPsList []*VSP

// VSPsAncestor is the interface that an ancestor of a VSP must implement.
// An Ancestor is defined as an entity that has VSP as a descendant.
// An Ancestor can get a list of its child VSPs, but not necessarily create one.
type VSPsAncestor interface {
	VSPs(*bambou.FetchingInfo) (VSPsList, *bambou.Error)
}

// VSPsParent is the interface that a parent of a VSP must implement.
// A Parent is defined as an entity that has VSP as a child.
// A Parent is an Ancestor which can create a VSP.
type VSPsParent interface {
	VSPsAncestor
	CreateVSP(*VSP) *bambou.Error
}

// VSP represents the model of a vsp
type VSP struct {
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	Name           string `json:"name,omitempty"`
	LastUpdatedBy  string `json:"lastUpdatedBy,omitempty"`
	Description    string `json:"description,omitempty"`
	EntityScope    string `json:"entityScope,omitempty"`
	Location       string `json:"location,omitempty"`
	ProductVersion string `json:"productVersion,omitempty"`
	ExternalID     string `json:"externalID,omitempty"`
}

// NewVSP returns a new *VSP
func NewVSP() *VSP {

	return &VSP{}
}

// Identity returns the Identity of the object.
func (o *VSP) Identity() bambou.Identity {

	return VSPIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VSP) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VSP) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VSP from the server
func (o *VSP) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VSP into the server
func (o *VSP) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VSP from the server
func (o *VSP) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VSP
func (o *VSP) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VSP
func (o *VSP) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VSP
func (o *VSP) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VSP
func (o *VSP) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// HSCs retrieves the list of child HSCs of the VSP
func (o *VSP) HSCs(info *bambou.FetchingInfo) (HSCsList, *bambou.Error) {

	var list HSCsList
	err := bambou.CurrentSession().FetchChildren(o, HSCIdentity, &list, info)
	return list, err
}

// VSCs retrieves the list of child VSCs of the VSP
func (o *VSP) VSCs(info *bambou.FetchingInfo) (VSCsList, *bambou.Error) {

	var list VSCsList
	err := bambou.CurrentSession().FetchChildren(o, VSCIdentity, &list, info)
	return list, err
}

// VSDs retrieves the list of child VSDs of the VSP
func (o *VSP) VSDs(info *bambou.FetchingInfo) (VSDsList, *bambou.Error) {

	var list VSDsList
	err := bambou.CurrentSession().FetchChildren(o, VSDIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the VSP
func (o *VSP) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
