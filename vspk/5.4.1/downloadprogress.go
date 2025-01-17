/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// DownloadProgressIdentity represents the Identity of the object
var DownloadProgressIdentity = bambou.Identity{
	Name:     "None",
	Category: "None",
}

// DownloadProgressList represents a list of DownloadProgress
type DownloadProgressList []*DownloadProgress

// DownloadProgressAncestor is the interface that an ancestor of a DownloadProgress must implement.
// An Ancestor is defined as an entity that has DownloadProgress as a descendant.
// An Ancestor can get a list of its child DownloadProgress, but not necessarily create one.
type DownloadProgressAncestor interface {
	DownloadProgress(*bambou.FetchingInfo) (DownloadProgressList, *bambou.Error)
}

// DownloadProgressParent is the interface that a parent of a DownloadProgress must implement.
// A Parent is defined as an entity that has DownloadProgress as a child.
// A Parent is an Ancestor which can create a DownloadProgress.
type DownloadProgressParent interface {
	DownloadProgressAncestor
	CreateDownloadProgress(*DownloadProgress) *bambou.Error
}

// DownloadProgress represents the model of a None
type DownloadProgress struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Percentage    int    `json:"percentage,omitempty"`
	TimeLeft      string `json:"timeLeft,omitempty"`
	TimeSpent     string `json:"timeSpent,omitempty"`
	ImageFileName string `json:"imageFileName,omitempty"`
	ImageVersion  string `json:"imageVersion,omitempty"`
	StartTime     int    `json:"startTime,omitempty"`
	CurrentSpeed  int    `json:"currentSpeed,omitempty"`
	AverageSpeed  int    `json:"averageSpeed,omitempty"`
}

// NewDownloadProgress returns a new *DownloadProgress
func NewDownloadProgress() *DownloadProgress {

	return &DownloadProgress{}
}

// Identity returns the Identity of the object.
func (o *DownloadProgress) Identity() bambou.Identity {

	return DownloadProgressIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DownloadProgress) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DownloadProgress) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the DownloadProgress from the server
func (o *DownloadProgress) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the DownloadProgress into the server
func (o *DownloadProgress) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the DownloadProgress from the server
func (o *DownloadProgress) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
