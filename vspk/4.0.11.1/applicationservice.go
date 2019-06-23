/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// ApplicationServiceIdentity represents the Identity of the object
var ApplicationServiceIdentity = bambou.Identity{
	Name:     "applicationservice",
	Category: "applicationservices",
}

// ApplicationServicesList represents a list of ApplicationServices
type ApplicationServicesList []*ApplicationService

// ApplicationServicesAncestor is the interface that an ancestor of a ApplicationService must implement.
// An Ancestor is defined as an entity that has ApplicationService as a descendant.
// An Ancestor can get a list of its child ApplicationServices, but not necessarily create one.
type ApplicationServicesAncestor interface {
	ApplicationServices(*bambou.FetchingInfo) (ApplicationServicesList, *bambou.Error)
}

// ApplicationServicesParent is the interface that a parent of a ApplicationService must implement.
// A Parent is defined as an entity that has ApplicationService as a child.
// A Parent is an Ancestor which can create a ApplicationService.
type ApplicationServicesParent interface {
	ApplicationServicesAncestor
	CreateApplicationService(*ApplicationService) *bambou.Error
}

// ApplicationService represents the model of a applicationservice
type ApplicationService struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	DSCP            string `json:"DSCP,omitempty"`
	Name            string `json:"name,omitempty"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	Description     string `json:"description,omitempty"`
	DestinationPort string `json:"destinationPort,omitempty"`
	Direction       string `json:"direction,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	SourcePort      string `json:"sourcePort,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	EtherType       string `json:"etherType,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewApplicationService returns a new *ApplicationService
func NewApplicationService() *ApplicationService {

	return &ApplicationService{}
}

// Identity returns the Identity of the object.
func (o *ApplicationService) Identity() bambou.Identity {

	return ApplicationServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ApplicationService) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ApplicationService) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ApplicationService from the server
func (o *ApplicationService) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ApplicationService into the server
func (o *ApplicationService) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ApplicationService from the server
func (o *ApplicationService) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ApplicationService
func (o *ApplicationService) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ApplicationService
func (o *ApplicationService) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ApplicationService
func (o *ApplicationService) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ApplicationService
func (o *ApplicationService) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the ApplicationService
func (o *ApplicationService) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
