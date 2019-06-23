/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// KeyServerNotificationIdentity represents the Identity of the object
var KeyServerNotificationIdentity = bambou.Identity{
	Name:     "keyservernotification",
	Category: "keyservernotifications",
}

// KeyServerNotificationsList represents a list of KeyServerNotifications
type KeyServerNotificationsList []*KeyServerNotification

// KeyServerNotificationsAncestor is the interface that an ancestor of a KeyServerNotification must implement.
// An Ancestor is defined as an entity that has KeyServerNotification as a descendant.
// An Ancestor can get a list of its child KeyServerNotifications, but not necessarily create one.
type KeyServerNotificationsAncestor interface {
	KeyServerNotifications(*bambou.FetchingInfo) (KeyServerNotificationsList, *bambou.Error)
}

// KeyServerNotificationsParent is the interface that a parent of a KeyServerNotification must implement.
// A Parent is defined as an entity that has KeyServerNotification as a child.
// A Parent is an Ancestor which can create a KeyServerNotification.
type KeyServerNotificationsParent interface {
	KeyServerNotificationsAncestor
	CreateKeyServerNotification(*KeyServerNotification) *bambou.Error
}

// KeyServerNotification represents the model of a keyservernotification
type KeyServerNotification struct {
	ID               string      `json:"ID,omitempty"`
	ParentID         string      `json:"parentID,omitempty"`
	ParentType       string      `json:"parentType,omitempty"`
	Owner            string      `json:"owner,omitempty"`
	Base64JSONString string      `json:"base64JSONString,omitempty"`
	Message          interface{} `json:"message,omitempty"`
	EntityScope      string      `json:"entityScope,omitempty"`
	NotificationType string      `json:"notificationType,omitempty"`
	ExternalID       string      `json:"externalID,omitempty"`
}

// NewKeyServerNotification returns a new *KeyServerNotification
func NewKeyServerNotification() *KeyServerNotification {

	return &KeyServerNotification{}
}

// Identity returns the Identity of the object.
func (o *KeyServerNotification) Identity() bambou.Identity {

	return KeyServerNotificationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *KeyServerNotification) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *KeyServerNotification) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the KeyServerNotification from the server
func (o *KeyServerNotification) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the KeyServerNotification into the server
func (o *KeyServerNotification) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the KeyServerNotification from the server
func (o *KeyServerNotification) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the KeyServerNotification
func (o *KeyServerNotification) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the KeyServerNotification
func (o *KeyServerNotification) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the KeyServerNotification
func (o *KeyServerNotification) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the KeyServerNotification
func (o *KeyServerNotification) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
