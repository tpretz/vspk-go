/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IngressExternalServiceTemplateEntryIdentity represents the Identity of the object
var IngressExternalServiceTemplateEntryIdentity = bambou.Identity{
	Name:     "ingressexternalserviceentrytemplate",
	Category: "ingressexternalserviceentrytemplates",
}

// IngressExternalServiceTemplateEntriesList represents a list of IngressExternalServiceTemplateEntries
type IngressExternalServiceTemplateEntriesList []*IngressExternalServiceTemplateEntry

// IngressExternalServiceTemplateEntriesAncestor is the interface that an ancestor of a IngressExternalServiceTemplateEntry must implement.
// An Ancestor is defined as an entity that has IngressExternalServiceTemplateEntry as a descendant.
// An Ancestor can get a list of its child IngressExternalServiceTemplateEntries, but not necessarily create one.
type IngressExternalServiceTemplateEntriesAncestor interface {
	IngressExternalServiceTemplateEntries(*bambou.FetchingInfo) (IngressExternalServiceTemplateEntriesList, *bambou.Error)
}

// IngressExternalServiceTemplateEntriesParent is the interface that a parent of a IngressExternalServiceTemplateEntry must implement.
// A Parent is defined as an entity that has IngressExternalServiceTemplateEntry as a child.
// A Parent is an Ancestor which can create a IngressExternalServiceTemplateEntry.
type IngressExternalServiceTemplateEntriesParent interface {
	IngressExternalServiceTemplateEntriesAncestor
	CreateIngressExternalServiceTemplateEntry(*IngressExternalServiceTemplateEntry) *bambou.Error
}

// IngressExternalServiceTemplateEntry represents the model of a ingressexternalserviceentrytemplate
type IngressExternalServiceTemplateEntry struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	ICMPCode                          string `json:"ICMPCode,omitempty"`
	ICMPType                          string `json:"ICMPType,omitempty"`
	DSCP                              string `json:"DSCP,omitempty"`
	Name                              string `json:"name,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	Action                            string `json:"action,omitempty"`
	AddressOverride                   string `json:"addressOverride,omitempty"`
	RedirectExternalServiceEndPointID string `json:"redirectExternalServiceEndPointID,omitempty"`
	Description                       string `json:"description,omitempty"`
	DestinationPort                   string `json:"destinationPort,omitempty"`
	NetworkID                         string `json:"networkID,omitempty"`
	NetworkType                       string `json:"networkType,omitempty"`
	MirrorDestinationID               string `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled                bool   `json:"flowLoggingEnabled"`
	EntityScope                       string `json:"entityScope,omitempty"`
	LocationID                        string `json:"locationID,omitempty"`
	LocationType                      string `json:"locationType,omitempty"`
	PolicyState                       string `json:"policyState,omitempty"`
	SourcePort                        string `json:"sourcePort,omitempty"`
	Priority                          int    `json:"priority,omitempty"`
	Protocol                          string `json:"protocol,omitempty"`
	AssociatedApplicationID           string `json:"associatedApplicationID,omitempty"`
	AssociatedApplicationObjectID     string `json:"associatedApplicationObjectID,omitempty"`
	AssociatedApplicationObjectType   string `json:"associatedApplicationObjectType,omitempty"`
	AssociatedLiveEntityID            string `json:"associatedLiveEntityID,omitempty"`
	StatsID                           string `json:"statsID,omitempty"`
	StatsLoggingEnabled               bool   `json:"statsLoggingEnabled"`
	EtherType                         string `json:"etherType,omitempty"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewIngressExternalServiceTemplateEntry returns a new *IngressExternalServiceTemplateEntry
func NewIngressExternalServiceTemplateEntry() *IngressExternalServiceTemplateEntry {

	return &IngressExternalServiceTemplateEntry{}
}

// Identity returns the Identity of the object.
func (o *IngressExternalServiceTemplateEntry) Identity() bambou.Identity {

	return IngressExternalServiceTemplateEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressExternalServiceTemplateEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressExternalServiceTemplateEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressExternalServiceTemplateEntry from the server
func (o *IngressExternalServiceTemplateEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressExternalServiceTemplateEntry into the server
func (o *IngressExternalServiceTemplateEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressExternalServiceTemplateEntry from the server
func (o *IngressExternalServiceTemplateEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IngressExternalServiceTemplateEntry
func (o *IngressExternalServiceTemplateEntry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IngressExternalServiceTemplateEntry
func (o *IngressExternalServiceTemplateEntry) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressExternalServiceTemplateEntry
func (o *IngressExternalServiceTemplateEntry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressExternalServiceTemplateEntry
func (o *IngressExternalServiceTemplateEntry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the IngressExternalServiceTemplateEntry
func (o *IngressExternalServiceTemplateEntry) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the IngressExternalServiceTemplateEntry
func (o *IngressExternalServiceTemplateEntry) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}
