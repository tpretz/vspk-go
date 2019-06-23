/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// AppIdentity represents the Identity of the object
var AppIdentity = bambou.Identity{
	Name:     "application",
	Category: "applications",
}

// AppsList represents a list of Apps
type AppsList []*App

// AppsAncestor is the interface that an ancestor of a App must implement.
// An Ancestor is defined as an entity that has App as a descendant.
// An Ancestor can get a list of its child Apps, but not necessarily create one.
type AppsAncestor interface {
	Apps(*bambou.FetchingInfo) (AppsList, *bambou.Error)
}

// AppsParent is the interface that a parent of a App must implement.
// A Parent is defined as an entity that has App as a child.
// A Parent is an Ancestor which can create a App.
type AppsParent interface {
	AppsAncestor
	CreateApp(*App) *bambou.Error
}

// App represents the model of a application
type App struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	Description                 string `json:"description,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	AssocEgressACLTemplateId    string `json:"assocEgressACLTemplateId,omitempty"`
	AssocIngressACLTemplateId   string `json:"assocIngressACLTemplateId,omitempty"`
	AssociatedDomainID          string `json:"associatedDomainID,omitempty"`
	AssociatedDomainType        string `json:"associatedDomainType,omitempty"`
	AssociatedNetworkObjectID   string `json:"associatedNetworkObjectID,omitempty"`
	AssociatedNetworkObjectType string `json:"associatedNetworkObjectType,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewApp returns a new *App
func NewApp() *App {

	return &App{}
}

// Identity returns the Identity of the object.
func (o *App) Identity() bambou.Identity {

	return AppIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *App) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *App) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the App from the server
func (o *App) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the App into the server
func (o *App) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the App from the server
func (o *App) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the App
func (o *App) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the App
func (o *App) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Tiers retrieves the list of child Tiers of the App
func (o *App) Tiers(info *bambou.FetchingInfo) (TiersList, *bambou.Error) {

	var list TiersList
	err := bambou.CurrentSession().FetchChildren(o, TierIdentity, &list, info)
	return list, err
}

// CreateTier creates a new child Tier under the App
func (o *App) CreateTier(child *Tier) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the App
func (o *App) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the App
func (o *App) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Flows retrieves the list of child Flows of the App
func (o *App) Flows(info *bambou.FetchingInfo) (FlowsList, *bambou.Error) {

	var list FlowsList
	err := bambou.CurrentSession().FetchChildren(o, FlowIdentity, &list, info)
	return list, err
}

// CreateFlow creates a new child Flow under the App
func (o *App) CreateFlow(child *Flow) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the App
func (o *App) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the App
func (o *App) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
