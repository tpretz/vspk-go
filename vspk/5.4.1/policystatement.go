/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PolicyStatementIdentity represents the Identity of the object
var PolicyStatementIdentity = bambou.Identity{
	Name:     "policystatement",
	Category: "policystatements",
}

// PolicyStatementsList represents a list of PolicyStatements
type PolicyStatementsList []*PolicyStatement

// PolicyStatementsAncestor is the interface that an ancestor of a PolicyStatement must implement.
// An Ancestor is defined as an entity that has PolicyStatement as a descendant.
// An Ancestor can get a list of its child PolicyStatements, but not necessarily create one.
type PolicyStatementsAncestor interface {
	PolicyStatements(*bambou.FetchingInfo) (PolicyStatementsList, *bambou.Error)
}

// PolicyStatementsParent is the interface that a parent of a PolicyStatement must implement.
// A Parent is defined as an entity that has PolicyStatement as a child.
// A Parent is an Ancestor which can create a PolicyStatement.
type PolicyStatementsParent interface {
	PolicyStatementsAncestor
	CreatePolicyStatement(*PolicyStatement) *bambou.Error
}

// PolicyStatement represents the model of a policystatement
type PolicyStatement struct {
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

// NewPolicyStatement returns a new *PolicyStatement
func NewPolicyStatement() *PolicyStatement {

	return &PolicyStatement{}
}

// Identity returns the Identity of the object.
func (o *PolicyStatement) Identity() bambou.Identity {

	return PolicyStatementIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyStatement) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyStatement) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyStatement from the server
func (o *PolicyStatement) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyStatement into the server
func (o *PolicyStatement) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyStatement from the server
func (o *PolicyStatement) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PolicyStatement
func (o *PolicyStatement) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyStatement
func (o *PolicyStatement) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyStatement
func (o *PolicyStatement) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyStatement
func (o *PolicyStatement) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PolicyEntries retrieves the list of child PolicyEntries of the PolicyStatement
func (o *PolicyStatement) PolicyEntries(info *bambou.FetchingInfo) (PolicyEntriesList, *bambou.Error) {

	var list PolicyEntriesList
	err := bambou.CurrentSession().FetchChildren(o, PolicyEntryIdentity, &list, info)
	return list, err
}

// CreatePolicyEntry creates a new child PolicyEntry under the PolicyStatement
func (o *PolicyStatement) CreatePolicyEntry(child *PolicyEntry) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
