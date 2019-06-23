/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ExternalServiceIdentity represents the Identity of the object
var ExternalServiceIdentity = bambou.Identity{
	Name:     "externalservice",
	Category: "externalservices",
}

// ExternalServicesList represents a list of ExternalServices
type ExternalServicesList []*ExternalService

// ExternalServicesAncestor is the interface that an ancestor of a ExternalService must implement.
// An Ancestor is defined as an entity that has ExternalService as a descendant.
// An Ancestor can get a list of its child ExternalServices, but not necessarily create one.
type ExternalServicesAncestor interface {
	ExternalServices(*bambou.FetchingInfo) (ExternalServicesList, *bambou.Error)
}

// ExternalServicesParent is the interface that a parent of a ExternalService must implement.
// A Parent is defined as an entity that has ExternalService as a child.
// A Parent is an Ancestor which can create a ExternalService.
type ExternalServicesParent interface {
	ExternalServicesAncestor
	CreateExternalService(*ExternalService) *bambou.Error
}

// ExternalService represents the model of a externalservice
type ExternalService struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewExternalService returns a new *ExternalService
func NewExternalService() *ExternalService {

	return &ExternalService{}
}

// Identity returns the Identity of the object.
func (o *ExternalService) Identity() bambou.Identity {

	return ExternalServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ExternalService) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ExternalService) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ExternalService from the server
func (o *ExternalService) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ExternalService into the server
func (o *ExternalService) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ExternalService from the server
func (o *ExternalService) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ExternalService
func (o *ExternalService) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ExternalService
func (o *ExternalService) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// MetadataTags retrieves the list of child MetadataTags of the ExternalService
func (o *ExternalService) MetadataTags(info *bambou.FetchingInfo) (MetadataTagsList, *bambou.Error) {

	var list MetadataTagsList
	err := bambou.CurrentSession().FetchChildren(o, MetadataTagIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ExternalService
func (o *ExternalService) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ExternalService
func (o *ExternalService) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EndPoints retrieves the list of child EndPoints of the ExternalService
func (o *ExternalService) EndPoints(info *bambou.FetchingInfo) (EndPointsList, *bambou.Error) {

	var list EndPointsList
	err := bambou.CurrentSession().FetchChildren(o, EndPointIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the ExternalService
func (o *ExternalService) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
