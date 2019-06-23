/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// FirewallRuleIdentity represents the Identity of the object
var FirewallRuleIdentity = bambou.Identity{
	Name:     "firewallrule",
	Category: "firewallrules",
}

// FirewallRulesList represents a list of FirewallRules
type FirewallRulesList []*FirewallRule

// FirewallRulesAncestor is the interface that an ancestor of a FirewallRule must implement.
// An Ancestor is defined as an entity that has FirewallRule as a descendant.
// An Ancestor can get a list of its child FirewallRules, but not necessarily create one.
type FirewallRulesAncestor interface {
	FirewallRules(*bambou.FetchingInfo) (FirewallRulesList, *bambou.Error)
}

// FirewallRulesParent is the interface that a parent of a FirewallRule must implement.
// A Parent is defined as an entity that has FirewallRule as a child.
// A Parent is an Ancestor which can create a FirewallRule.
type FirewallRulesParent interface {
	FirewallRulesAncestor
	CreateFirewallRule(*FirewallRule) *bambou.Error
}

// FirewallRule represents the model of a firewallrule
type FirewallRule struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewFirewallRule returns a new *FirewallRule
func NewFirewallRule() *FirewallRule {

	return &FirewallRule{}
}

// Identity returns the Identity of the object.
func (o *FirewallRule) Identity() bambou.Identity {

	return FirewallRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FirewallRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FirewallRule) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FirewallRule from the server
func (o *FirewallRule) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FirewallRule into the server
func (o *FirewallRule) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FirewallRule from the server
func (o *FirewallRule) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
