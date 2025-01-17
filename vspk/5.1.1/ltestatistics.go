/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LtestatisticsIdentity represents the Identity of the object
var LtestatisticsIdentity = bambou.Identity{
	Name:     "ltestatistics",
	Category: "ltestatistics",
}

// LtestatisticsList represents a list of Ltestatistics
type LtestatisticsList []*Ltestatistics

// LtestatisticsAncestor is the interface that an ancestor of a Ltestatistics must implement.
// An Ancestor is defined as an entity that has Ltestatistics as a descendant.
// An Ancestor can get a list of its child Ltestatistics, but not necessarily create one.
type LtestatisticsAncestor interface {
	Ltestatistics(*bambou.FetchingInfo) (LtestatisticsList, *bambou.Error)
}

// LtestatisticsParent is the interface that a parent of a Ltestatistics must implement.
// A Parent is defined as an entity that has Ltestatistics as a child.
// A Parent is an Ancestor which can create a Ltestatistics.
type LtestatisticsParent interface {
	LtestatisticsAncestor
	CreateLtestatistics(*Ltestatistics) *bambou.Error
}

// Ltestatistics represents the model of a ltestatistics
type Ltestatistics struct {
	ID         string        `json:"ID,omitempty"`
	ParentID   string        `json:"parentID,omitempty"`
	ParentType string        `json:"parentType,omitempty"`
	Owner      string        `json:"owner,omitempty"`
	Version    int           `json:"version,omitempty"`
	EndTime    int           `json:"endTime,omitempty"`
	StartTime  int           `json:"startTime,omitempty"`
	StatsData  []interface{} `json:"statsData,omitempty"`
}

// NewLtestatistics returns a new *Ltestatistics
func NewLtestatistics() *Ltestatistics {

	return &Ltestatistics{}
}

// Identity returns the Identity of the object.
func (o *Ltestatistics) Identity() bambou.Identity {

	return LtestatisticsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Ltestatistics) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Ltestatistics) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Ltestatistics from the server
func (o *Ltestatistics) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Ltestatistics into the server
func (o *Ltestatistics) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Ltestatistics from the server
func (o *Ltestatistics) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Ltestatistics
func (o *Ltestatistics) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Ltestatistics
func (o *Ltestatistics) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Ltestatistics
func (o *Ltestatistics) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Ltestatistics
func (o *Ltestatistics) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
