/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// QOSIdentity represents the Identity of the object
var QOSIdentity = bambou.Identity{
	Name:     "qos",
	Category: "qos",
}

// QOSsList represents a list of QOSs
type QOSsList []*QOS

// QOSsAncestor is the interface that an ancestor of a QOS must implement.
// An Ancestor is defined as an entity that has QOS as a descendant.
// An Ancestor can get a list of its child QOSs, but not necessarily create one.
type QOSsAncestor interface {
	QOSs(*bambou.FetchingInfo) (QOSsList, *bambou.Error)
}

// QOSsParent is the interface that a parent of a QOS must implement.
// A Parent is defined as an entity that has QOS as a child.
// A Parent is an Ancestor which can create a QOS.
type QOSsParent interface {
	QOSsAncestor
	CreateQOS(*QOS) *bambou.Error
}

// QOS represents the model of a qos
type QOS struct {
	ID         string `json:"ID,omitempty"`
	ParentID   string `json:"parentID,omitempty"`
	ParentType string `json:"parentType,omitempty"`
	Owner      string `json:"owner,omitempty"`
}

// NewQOS returns a new *QOS
func NewQOS() *QOS {

	return &QOS{}
}

// Identity returns the Identity of the object.
func (o *QOS) Identity() bambou.Identity {

	return QOSIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *QOS) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *QOS) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the QOS from the server
func (o *QOS) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the QOS into the server
func (o *QOS) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the QOS from the server
func (o *QOS) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the QOS
func (o *QOS) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the QOS
func (o *QOS) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the QOS
func (o *QOS) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the QOS
func (o *QOS) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VMs retrieves the list of child VMs of the QOS
func (o *QOS) VMs(info *bambou.FetchingInfo) (VMsList, *bambou.Error) {

	var list VMsList
	err := bambou.CurrentSession().FetchChildren(o, VMIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the QOS
func (o *QOS) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
