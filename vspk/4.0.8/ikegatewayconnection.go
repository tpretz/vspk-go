/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKEGatewayConnectionIdentity represents the Identity of the object
var IKEGatewayConnectionIdentity = bambou.Identity{
	Name:     "ikegatewayconnection",
	Category: "ikegatewayconnections",
}

// IKEGatewayConnectionsList represents a list of IKEGatewayConnections
type IKEGatewayConnectionsList []*IKEGatewayConnection

// IKEGatewayConnectionsAncestor is the interface that an ancestor of a IKEGatewayConnection must implement.
// An Ancestor is defined as an entity that has IKEGatewayConnection as a descendant.
// An Ancestor can get a list of its child IKEGatewayConnections, but not necessarily create one.
type IKEGatewayConnectionsAncestor interface {
	IKEGatewayConnections(*bambou.FetchingInfo) (IKEGatewayConnectionsList, *bambou.Error)
}

// IKEGatewayConnectionsParent is the interface that a parent of a IKEGatewayConnection must implement.
// A Parent is defined as an entity that has IKEGatewayConnection as a child.
// A Parent is an Ancestor which can create a IKEGatewayConnection.
type IKEGatewayConnectionsParent interface {
	IKEGatewayConnectionsAncestor
	CreateIKEGatewayConnection(*IKEGatewayConnection) *bambou.Error
}

// IKEGatewayConnection represents the model of a ikegatewayconnection
type IKEGatewayConnection struct {
	ID                               string `json:"ID,omitempty"`
	ParentID                         string `json:"parentID,omitempty"`
	ParentType                       string `json:"parentType,omitempty"`
	Owner                            string `json:"owner,omitempty"`
	NSGIdentifier                    string `json:"NSGIdentifier,omitempty"`
	NSGIdentifierType                string `json:"NSGIdentifierType,omitempty"`
	NSGRole                          string `json:"NSGRole,omitempty"`
	Name                             string `json:"name,omitempty"`
	LastUpdatedBy                    string `json:"lastUpdatedBy,omitempty"`
	Sequence                         int    `json:"sequence,omitempty"`
	AllowAnySubnet                   bool   `json:"allowAnySubnet"`
	UnencryptedPSK                   string `json:"unencryptedPSK,omitempty"`
	EntityScope                      string `json:"entityScope,omitempty"`
	PortVLANName                     string `json:"portVLANName,omitempty"`
	Priority                         int    `json:"priority,omitempty"`
	AssociatedIKEAuthenticationID    string `json:"associatedIKEAuthenticationID,omitempty"`
	AssociatedIKEAuthenticationType  string `json:"associatedIKEAuthenticationType,omitempty"`
	AssociatedIKEEncryptionProfileID string `json:"associatedIKEEncryptionProfileID,omitempty"`
	AssociatedIKEGatewayProfileID    string `json:"associatedIKEGatewayProfileID,omitempty"`
	AssociatedVLANID                 string `json:"associatedVLANID,omitempty"`
	ExternalID                       string `json:"externalID,omitempty"`
}

// NewIKEGatewayConnection returns a new *IKEGatewayConnection
func NewIKEGatewayConnection() *IKEGatewayConnection {

	return &IKEGatewayConnection{
		NSGIdentifierType: "ID_KEY_ID",
	}
}

// Identity returns the Identity of the object.
func (o *IKEGatewayConnection) Identity() bambou.Identity {

	return IKEGatewayConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEGatewayConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEGatewayConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEGatewayConnection from the server
func (o *IKEGatewayConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEGatewayConnection into the server
func (o *IKEGatewayConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEGatewayConnection from the server
func (o *IKEGatewayConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEGatewayConnection
func (o *IKEGatewayConnection) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEGatewayConnection
func (o *IKEGatewayConnection) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEGatewayConnection
func (o *IKEGatewayConnection) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEGatewayConnection
func (o *IKEGatewayConnection) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Subnets retrieves the list of child Subnets of the IKEGatewayConnection
func (o *IKEGatewayConnection) Subnets(info *bambou.FetchingInfo) (SubnetsList, *bambou.Error) {

	var list SubnetsList
	err := bambou.CurrentSession().FetchChildren(o, SubnetIdentity, &list, info)
	return list, err
}
