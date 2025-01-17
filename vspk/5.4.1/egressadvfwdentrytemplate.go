/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EgressAdvFwdEntryTemplateIdentity represents the Identity of the object
var EgressAdvFwdEntryTemplateIdentity = bambou.Identity{
	Name:     "egressadvfwdentrytemplate",
	Category: "egressadvfwdentrytemplates",
}

// EgressAdvFwdEntryTemplatesList represents a list of EgressAdvFwdEntryTemplates
type EgressAdvFwdEntryTemplatesList []*EgressAdvFwdEntryTemplate

// EgressAdvFwdEntryTemplatesAncestor is the interface that an ancestor of a EgressAdvFwdEntryTemplate must implement.
// An Ancestor is defined as an entity that has EgressAdvFwdEntryTemplate as a descendant.
// An Ancestor can get a list of its child EgressAdvFwdEntryTemplates, but not necessarily create one.
type EgressAdvFwdEntryTemplatesAncestor interface {
	EgressAdvFwdEntryTemplates(*bambou.FetchingInfo) (EgressAdvFwdEntryTemplatesList, *bambou.Error)
}

// EgressAdvFwdEntryTemplatesParent is the interface that a parent of a EgressAdvFwdEntryTemplate must implement.
// A Parent is defined as an entity that has EgressAdvFwdEntryTemplate as a child.
// A Parent is an Ancestor which can create a EgressAdvFwdEntryTemplate.
type EgressAdvFwdEntryTemplatesParent interface {
	EgressAdvFwdEntryTemplatesAncestor
	CreateEgressAdvFwdEntryTemplate(*EgressAdvFwdEntryTemplate) *bambou.Error
}

// EgressAdvFwdEntryTemplate represents the model of a egressadvfwdentrytemplate
type EgressAdvFwdEntryTemplate struct {
	ID                       string `json:"ID,omitempty"`
	ParentID                 string `json:"parentID,omitempty"`
	ParentType               string `json:"parentType,omitempty"`
	Owner                    string `json:"owner,omitempty"`
	ACLTemplateName          string `json:"ACLTemplateName,omitempty"`
	ICMPCode                 string `json:"ICMPCode,omitempty"`
	ICMPType                 string `json:"ICMPType,omitempty"`
	FCOverride               string `json:"FCOverride,omitempty"`
	IPv6AddressOverride      string `json:"IPv6AddressOverride,omitempty"`
	DSCP                     string `json:"DSCP,omitempty"`
	FailsafeDatapath         string `json:"failsafeDatapath,omitempty"`
	LastUpdatedBy            string `json:"lastUpdatedBy,omitempty"`
	Action                   string `json:"action,omitempty"`
	AddressOverride          string `json:"addressOverride,omitempty"`
	WebFilterID              string `json:"webFilterID,omitempty"`
	WebFilterType            string `json:"webFilterType,omitempty"`
	RedirectVPortTagID       string `json:"redirectVPortTagID,omitempty"`
	Description              string `json:"description,omitempty"`
	DestinationPort          string `json:"destinationPort,omitempty"`
	NetworkID                string `json:"networkID,omitempty"`
	NetworkType              string `json:"networkType,omitempty"`
	MirrorDestinationID      string `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled       bool   `json:"flowLoggingEnabled"`
	EnterpriseName           string `json:"enterpriseName,omitempty"`
	EntityScope              string `json:"entityScope,omitempty"`
	LocationID               string `json:"locationID,omitempty"`
	LocationType             string `json:"locationType,omitempty"`
	PolicyState              string `json:"policyState,omitempty"`
	DomainName               string `json:"domainName,omitempty"`
	SourcePort               string `json:"sourcePort,omitempty"`
	UplinkPreference         string `json:"uplinkPreference,omitempty"`
	Priority                 int    `json:"priority,omitempty"`
	Protocol                 string `json:"protocol,omitempty"`
	AssociatedLiveEntityID   string `json:"associatedLiveEntityID,omitempty"`
	AssociatedLiveTemplateID string `json:"associatedLiveTemplateID,omitempty"`
	AssociatedTrafficType    string `json:"associatedTrafficType,omitempty"`
	AssociatedTrafficTypeID  string `json:"associatedTrafficTypeID,omitempty"`
	StatsID                  string `json:"statsID,omitempty"`
	StatsLoggingEnabled      bool   `json:"statsLoggingEnabled"`
	EtherType                string `json:"etherType,omitempty"`
	ExternalID               string `json:"externalID,omitempty"`
}

// NewEgressAdvFwdEntryTemplate returns a new *EgressAdvFwdEntryTemplate
func NewEgressAdvFwdEntryTemplate() *EgressAdvFwdEntryTemplate {

	return &EgressAdvFwdEntryTemplate{
		FailsafeDatapath: "FAIL_TO_BLOCK",
	}
}

// Identity returns the Identity of the object.
func (o *EgressAdvFwdEntryTemplate) Identity() bambou.Identity {

	return EgressAdvFwdEntryTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressAdvFwdEntryTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressAdvFwdEntryTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressAdvFwdEntryTemplate from the server
func (o *EgressAdvFwdEntryTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressAdvFwdEntryTemplate into the server
func (o *EgressAdvFwdEntryTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressAdvFwdEntryTemplate from the server
func (o *EgressAdvFwdEntryTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EgressAdvFwdEntryTemplate
func (o *EgressAdvFwdEntryTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressAdvFwdEntryTemplate
func (o *EgressAdvFwdEntryTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressAdvFwdEntryTemplate
func (o *EgressAdvFwdEntryTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressAdvFwdEntryTemplate
func (o *EgressAdvFwdEntryTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
