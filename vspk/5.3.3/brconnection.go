/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// BRConnectionIdentity represents the Identity of the object
var BRConnectionIdentity = bambou.Identity{
	Name:     "brconnection",
	Category: "brconnections",
}

// BRConnectionsList represents a list of BRConnections
type BRConnectionsList []*BRConnection

// BRConnectionsAncestor is the interface that an ancestor of a BRConnection must implement.
// An Ancestor is defined as an entity that has BRConnection as a descendant.
// An Ancestor can get a list of its child BRConnections, but not necessarily create one.
type BRConnectionsAncestor interface {
	BRConnections(*bambou.FetchingInfo) (BRConnectionsList, *bambou.Error)
}

// BRConnectionsParent is the interface that a parent of a BRConnection must implement.
// A Parent is defined as an entity that has BRConnection as a child.
// A Parent is an Ancestor which can create a BRConnection.
type BRConnectionsParent interface {
	BRConnectionsAncestor
	CreateBRConnection(*BRConnection) *bambou.Error
}

// BRConnection represents the model of a brconnection
type BRConnection struct {
	ID                    string `json:"ID,omitempty"`
	ParentID              string `json:"parentID,omitempty"`
	ParentType            string `json:"parentType,omitempty"`
	Owner                 string `json:"owner,omitempty"`
	DNSAddress            string `json:"DNSAddress,omitempty"`
	DNSAddressV6          string `json:"DNSAddressV6,omitempty"`
	LastUpdatedBy         string `json:"lastUpdatedBy,omitempty"`
	Gateway               string `json:"gateway,omitempty"`
	GatewayV6             string `json:"gatewayV6,omitempty"`
	Address               string `json:"address,omitempty"`
	AddressFamily         string `json:"addressFamily,omitempty"`
	AddressV6             string `json:"addressV6,omitempty"`
	AdvertisementCriteria string `json:"advertisementCriteria,omitempty"`
	Netmask               string `json:"netmask,omitempty"`
	Inherited             bool   `json:"inherited"`
	EntityScope           string `json:"entityScope,omitempty"`
	Mode                  string `json:"mode,omitempty"`
	UplinkID              int    `json:"uplinkID,omitempty"`
	ExternalID            string `json:"externalID,omitempty"`
}

// NewBRConnection returns a new *BRConnection
func NewBRConnection() *BRConnection {

	return &BRConnection{
		AddressFamily:         "IPV4",
		AdvertisementCriteria: "OPERATIONAL_LINK",
		Inherited:             false,
	}
}

// Identity returns the Identity of the object.
func (o *BRConnection) Identity() bambou.Identity {

	return BRConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BRConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BRConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BRConnection from the server
func (o *BRConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BRConnection into the server
func (o *BRConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BRConnection from the server
func (o *BRConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BRConnection
func (o *BRConnection) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BRConnection
func (o *BRConnection) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// BFDSessions retrieves the list of child BFDSessions of the BRConnection
func (o *BRConnection) BFDSessions(info *bambou.FetchingInfo) (BFDSessionsList, *bambou.Error) {

	var list BFDSessionsList
	err := bambou.CurrentSession().FetchChildren(o, BFDSessionIdentity, &list, info)
	return list, err
}

// CreateBFDSession creates a new child BFDSession under the BRConnection
func (o *BRConnection) CreateBFDSession(child *BFDSession) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BRConnection
func (o *BRConnection) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BRConnection
func (o *BRConnection) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
