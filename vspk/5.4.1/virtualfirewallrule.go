/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VirtualFirewallRuleIdentity represents the Identity of the object
var VirtualFirewallRuleIdentity = bambou.Identity{
	Name:     "virtualfirewallrule",
	Category: "virtualfirewallrules",
}

// VirtualFirewallRulesList represents a list of VirtualFirewallRules
type VirtualFirewallRulesList []*VirtualFirewallRule

// VirtualFirewallRulesAncestor is the interface that an ancestor of a VirtualFirewallRule must implement.
// An Ancestor is defined as an entity that has VirtualFirewallRule as a descendant.
// An Ancestor can get a list of its child VirtualFirewallRules, but not necessarily create one.
type VirtualFirewallRulesAncestor interface {
	VirtualFirewallRules(*bambou.FetchingInfo) (VirtualFirewallRulesList, *bambou.Error)
}

// VirtualFirewallRulesParent is the interface that a parent of a VirtualFirewallRule must implement.
// A Parent is defined as an entity that has VirtualFirewallRule as a child.
// A Parent is an Ancestor which can create a VirtualFirewallRule.
type VirtualFirewallRulesParent interface {
	VirtualFirewallRulesAncestor
	CreateVirtualFirewallRule(*VirtualFirewallRule) *bambou.Error
}

// VirtualFirewallRule represents the model of a virtualfirewallrule
type VirtualFirewallRule struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	ACLTemplateName                    string `json:"ACLTemplateName,omitempty"`
	ICMPCode                           string `json:"ICMPCode,omitempty"`
	ICMPType                           string `json:"ICMPType,omitempty"`
	IPv6AddressOverride                string `json:"IPv6AddressOverride,omitempty"`
	DSCP                               string `json:"DSCP,omitempty"`
	LastUpdatedBy                      string `json:"lastUpdatedBy,omitempty"`
	Action                             string `json:"action,omitempty"`
	AddressOverride                    string `json:"addressOverride,omitempty"`
	WebFilterID                        string `json:"webFilterID,omitempty"`
	WebFilterType                      string `json:"webFilterType,omitempty"`
	Description                        string `json:"description,omitempty"`
	DestinationPort                    string `json:"destinationPort,omitempty"`
	NetworkID                          string `json:"networkID,omitempty"`
	NetworkType                        string `json:"networkType,omitempty"`
	MirrorDestinationID                string `json:"mirrorDestinationID,omitempty"`
	FlowLoggingEnabled                 bool   `json:"flowLoggingEnabled"`
	EnterpriseName                     string `json:"enterpriseName,omitempty"`
	EntityScope                        string `json:"entityScope,omitempty"`
	LocationID                         string `json:"locationID,omitempty"`
	LocationType                       string `json:"locationType,omitempty"`
	PolicyState                        string `json:"policyState,omitempty"`
	DomainName                         string `json:"domainName,omitempty"`
	SourcePort                         string `json:"sourcePort,omitempty"`
	Priority                           int    `json:"priority,omitempty"`
	Protocol                           string `json:"protocol,omitempty"`
	AssociatedEgressEntryID            string `json:"associatedEgressEntryID,omitempty"`
	AssociatedIngressEntryID           string `json:"associatedIngressEntryID,omitempty"`
	AssociatedL7ApplicationSignatureID string `json:"associatedL7ApplicationSignatureID,omitempty"`
	AssociatedLiveEntityID             string `json:"associatedLiveEntityID,omitempty"`
	AssociatedLiveTemplateID           string `json:"associatedLiveTemplateID,omitempty"`
	AssociatedTrafficType              string `json:"associatedTrafficType,omitempty"`
	AssociatedTrafficTypeID            string `json:"associatedTrafficTypeID,omitempty"`
	Stateful                           bool   `json:"stateful"`
	StatsID                            string `json:"statsID,omitempty"`
	StatsLoggingEnabled                bool   `json:"statsLoggingEnabled"`
	EtherType                          string `json:"etherType,omitempty"`
	OverlayMirrorDestinationID         string `json:"overlayMirrorDestinationID,omitempty"`
	ExternalID                         string `json:"externalID,omitempty"`
}

// NewVirtualFirewallRule returns a new *VirtualFirewallRule
func NewVirtualFirewallRule() *VirtualFirewallRule {

	return &VirtualFirewallRule{
		Action:              "FORWARD",
		NetworkType:         "ANY",
		FlowLoggingEnabled:  false,
		Stateful:            false,
		StatsLoggingEnabled: false,
	}
}

// Identity returns the Identity of the object.
func (o *VirtualFirewallRule) Identity() bambou.Identity {

	return VirtualFirewallRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VirtualFirewallRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VirtualFirewallRule) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VirtualFirewallRule from the server
func (o *VirtualFirewallRule) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VirtualFirewallRule into the server
func (o *VirtualFirewallRule) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VirtualFirewallRule from the server
func (o *VirtualFirewallRule) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VirtualFirewallRule
func (o *VirtualFirewallRule) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VirtualFirewallRule
func (o *VirtualFirewallRule) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VirtualFirewallRule
func (o *VirtualFirewallRule) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VirtualFirewallRule
func (o *VirtualFirewallRule) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
