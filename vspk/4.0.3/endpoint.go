/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EndPointIdentity represents the Identity of the object
var EndPointIdentity = bambou.Identity{
	Name:     "endpoint",
	Category: "endpoints",
}

// EndPointsList represents a list of EndPoints
type EndPointsList []*EndPoint

// EndPointsAncestor is the interface that an ancestor of a EndPoint must implement.
// An Ancestor is defined as an entity that has EndPoint as a descendant.
// An Ancestor can get a list of its child EndPoints, but not necessarily create one.
type EndPointsAncestor interface {
	EndPoints(*bambou.FetchingInfo) (EndPointsList, *bambou.Error)
}

// EndPointsParent is the interface that a parent of a EndPoint must implement.
// A Parent is defined as an entity that has EndPoint as a child.
// A Parent is an Ancestor which can create a EndPoint.
type EndPointsParent interface {
	EndPointsAncestor
	CreateEndPoint(*EndPoint) *bambou.Error
}

// EndPoint represents the model of a endpoint
type EndPoint struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewEndPoint returns a new *EndPoint
func NewEndPoint() *EndPoint {

	return &EndPoint{}
}

// Identity returns the Identity of the object.
func (o *EndPoint) Identity() bambou.Identity {

	return EndPointIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EndPoint) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EndPoint) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EndPoint from the server
func (o *EndPoint) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EndPoint into the server
func (o *EndPoint) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EndPoint from the server
func (o *EndPoint) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EndPoint
func (o *EndPoint) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EndPoint
func (o *EndPoint) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EndPoint
func (o *EndPoint) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EndPoint
func (o *EndPoint) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the EndPoint
func (o *EndPoint) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
