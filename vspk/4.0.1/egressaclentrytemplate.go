/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EgressACLEntryTemplateIdentity represents the Identity of the object
var EgressACLEntryTemplateIdentity = bambou.Identity{
	Name:     "egressaclentrytemplate",
	Category: "egressaclentrytemplates",
}

// EgressACLEntryTemplatesList represents a list of EgressACLEntryTemplates
type EgressACLEntryTemplatesList []*EgressACLEntryTemplate

// EgressACLEntryTemplatesAncestor is the interface that an ancestor of a EgressACLEntryTemplate must implement.
// An Ancestor is defined as an entity that has EgressACLEntryTemplate as a descendant.
// An Ancestor can get a list of its child EgressACLEntryTemplates, but not necessarily create one.
type EgressACLEntryTemplatesAncestor interface {
	EgressACLEntryTemplates(*bambou.FetchingInfo) (EgressACLEntryTemplatesList, *bambou.Error)
}

// EgressACLEntryTemplatesParent is the interface that a parent of a EgressACLEntryTemplate must implement.
// A Parent is defined as an entity that has EgressACLEntryTemplate as a child.
// A Parent is an Ancestor which can create a EgressACLEntryTemplate.
type EgressACLEntryTemplatesParent interface {
	EgressACLEntryTemplatesAncestor
	CreateEgressACLEntryTemplate(*EgressACLEntryTemplate) *bambou.Error
}

// EgressACLEntryTemplate represents the model of a egressaclentrytemplate
type EgressACLEntryTemplate struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	ICMPCode                        string `json:"ICMPCode,omitempty"`
	ICMPType                        string `json:"ICMPType,omitempty"`
	DSCP                            string `json:"DSCP,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	Action                          string `json:"action,omitempty"`
	AddressOverride                 string `json:"addressOverride,omitempty"`
	Reflexive                       bool   `json:"reflexive"`
	Description                     string `json:"description,omitempty"`
	DestinationPort                 string `json:"destinationPort,omitempty"`
	NetworkID                       string `json:"networkID,omitempty"`
	NetworkType                     string `json:"networkType,omitempty"`
	MirrorDestinationID             string `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled              bool   `json:"flowLoggingEnabled"`
	EntityScope                     string `json:"entityScope,omitempty"`
	LocationID                      string `json:"locationID,omitempty"`
	LocationType                    string `json:"locationType,omitempty"`
	PolicyState                     string `json:"policyState,omitempty"`
	SourcePort                      string `json:"sourcePort,omitempty"`
	Priority                        int    `json:"priority,omitempty"`
	Protocol                        string `json:"protocol,omitempty"`
	AssociatedApplicationID         string `json:"associatedApplicationID,omitempty"`
	AssociatedApplicationObjectID   string `json:"associatedApplicationObjectID,omitempty"`
	AssociatedApplicationObjectType string `json:"associatedApplicationObjectType,omitempty"`
	AssociatedLiveEntityID          string `json:"associatedLiveEntityID,omitempty"`
	Stateful                        bool   `json:"stateful"`
	StatsID                         string `json:"statsID,omitempty"`
	StatsLoggingEnabled             bool   `json:"statsLoggingEnabled"`
	EtherType                       string `json:"etherType,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
}

// NewEgressACLEntryTemplate returns a new *EgressACLEntryTemplate
func NewEgressACLEntryTemplate() *EgressACLEntryTemplate {

	return &EgressACLEntryTemplate{}
}

// Identity returns the Identity of the object.
func (o *EgressACLEntryTemplate) Identity() bambou.Identity {

	return EgressACLEntryTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressACLEntryTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressACLEntryTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressACLEntryTemplate from the server
func (o *EgressACLEntryTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressACLEntryTemplate into the server
func (o *EgressACLEntryTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressACLEntryTemplate from the server
func (o *EgressACLEntryTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EgressACLEntryTemplate
func (o *EgressACLEntryTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressACLEntryTemplate
func (o *EgressACLEntryTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressACLEntryTemplate
func (o *EgressACLEntryTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressACLEntryTemplate
func (o *EgressACLEntryTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the EgressACLEntryTemplate
func (o *EgressACLEntryTemplate) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the EgressACLEntryTemplate
func (o *EgressACLEntryTemplate) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}
