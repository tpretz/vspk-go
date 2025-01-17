/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IngressACLEntryTemplateIdentity represents the Identity of the object
var IngressACLEntryTemplateIdentity = bambou.Identity{
	Name:     "ingressaclentrytemplate",
	Category: "ingressaclentrytemplates",
}

// IngressACLEntryTemplatesList represents a list of IngressACLEntryTemplates
type IngressACLEntryTemplatesList []*IngressACLEntryTemplate

// IngressACLEntryTemplatesAncestor is the interface that an ancestor of a IngressACLEntryTemplate must implement.
// An Ancestor is defined as an entity that has IngressACLEntryTemplate as a descendant.
// An Ancestor can get a list of its child IngressACLEntryTemplates, but not necessarily create one.
type IngressACLEntryTemplatesAncestor interface {
	IngressACLEntryTemplates(*bambou.FetchingInfo) (IngressACLEntryTemplatesList, *bambou.Error)
}

// IngressACLEntryTemplatesParent is the interface that a parent of a IngressACLEntryTemplate must implement.
// A Parent is defined as an entity that has IngressACLEntryTemplate as a child.
// A Parent is an Ancestor which can create a IngressACLEntryTemplate.
type IngressACLEntryTemplatesParent interface {
	IngressACLEntryTemplatesAncestor
	CreateIngressACLEntryTemplate(*IngressACLEntryTemplate) *bambou.Error
}

// IngressACLEntryTemplate represents the model of a ingressaclentrytemplate
type IngressACLEntryTemplate struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	ACLTemplateName                 string `json:"ACLTemplateName,omitempty"`
	ICMPCode                        string `json:"ICMPCode,omitempty"`
	ICMPType                        string `json:"ICMPType,omitempty"`
	IPv6AddressOverride             string `json:"IPv6AddressOverride,omitempty"`
	DSCP                            string `json:"DSCP,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	Action                          string `json:"action,omitempty"`
	AddressOverride                 string `json:"addressOverride,omitempty"`
	Description                     string `json:"description,omitempty"`
	DestinationPort                 string `json:"destinationPort,omitempty"`
	NetworkID                       string `json:"networkID,omitempty"`
	NetworkType                     string `json:"networkType,omitempty"`
	MirrorDestinationID             string `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled              bool   `json:"flowLoggingEnabled"`
	EnterpriseName                  string `json:"enterpriseName,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	LocationID                      string `json:"locationID,omitempty"`
	LocationType                    string `json:"locationType,omitempty"`
	PolicyState                     string `json:"policyState,omitempty"`
	DomainName                      string `json:"domainName,omitempty"`
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

// NewIngressACLEntryTemplate returns a new *IngressACLEntryTemplate
func NewIngressACLEntryTemplate() *IngressACLEntryTemplate {

	return &IngressACLEntryTemplate{}
}

// Identity returns the Identity of the object.
func (o *IngressACLEntryTemplate) Identity() bambou.Identity {

	return IngressACLEntryTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressACLEntryTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressACLEntryTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressACLEntryTemplate from the server
func (o *IngressACLEntryTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressACLEntryTemplate into the server
func (o *IngressACLEntryTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressACLEntryTemplate from the server
func (o *IngressACLEntryTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the IngressACLEntryTemplate
func (o *IngressACLEntryTemplate) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}
