/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKEEncryptionprofileIdentity represents the Identity of the object
var IKEEncryptionprofileIdentity = bambou.Identity{
	Name:     "ikeencryptionprofile",
	Category: "ikeencryptionprofiles",
}

// IKEEncryptionprofilesList represents a list of IKEEncryptionprofiles
type IKEEncryptionprofilesList []*IKEEncryptionprofile

// IKEEncryptionprofilesAncestor is the interface that an ancestor of a IKEEncryptionprofile must implement.
// An Ancestor is defined as an entity that has IKEEncryptionprofile as a descendant.
// An Ancestor can get a list of its child IKEEncryptionprofiles, but not necessarily create one.
type IKEEncryptionprofilesAncestor interface {
	IKEEncryptionprofiles(*bambou.FetchingInfo) (IKEEncryptionprofilesList, *bambou.Error)
}

// IKEEncryptionprofilesParent is the interface that a parent of a IKEEncryptionprofile must implement.
// A Parent is defined as an entity that has IKEEncryptionprofile as a child.
// A Parent is an Ancestor which can create a IKEEncryptionprofile.
type IKEEncryptionprofilesParent interface {
	IKEEncryptionprofilesAncestor
	CreateIKEEncryptionprofile(*IKEEncryptionprofile) *bambou.Error
}

// IKEEncryptionprofile represents the model of a ikeencryptionprofile
type IKEEncryptionprofile struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	DPDInterval                       int    `json:"DPDInterval,omitempty"`
	DPDMode                           string `json:"DPDMode,omitempty"`
	DPDTimeout                        int    `json:"DPDTimeout,omitempty"`
	IPsecAuthenticationAlgorithm      string `json:"IPsecAuthenticationAlgorithm,omitempty"`
	IPsecDontFragment                 bool   `json:"IPsecDontFragment"`
	IPsecEnablePFS                    bool   `json:"IPsecEnablePFS"`
	IPsecEncryptionAlgorithm          string `json:"IPsecEncryptionAlgorithm,omitempty"`
	IPsecPreFragment                  bool   `json:"IPsecPreFragment"`
	IPsecSALifetime                   int    `json:"IPsecSALifetime,omitempty"`
	IPsecSAReplayWindowSize           string `json:"IPsecSAReplayWindowSize,omitempty"`
	ISAKMPAuthenticationMode          string `json:"ISAKMPAuthenticationMode,omitempty"`
	ISAKMPDiffieHelmanGroupIdentifier string `json:"ISAKMPDiffieHelmanGroupIdentifier,omitempty"`
	ISAKMPEncryptionAlgorithm         string `json:"ISAKMPEncryptionAlgorithm,omitempty"`
	ISAKMPEncryptionKeyLifetime       int    `json:"ISAKMPEncryptionKeyLifetime,omitempty"`
	ISAKMPHashAlgorithm               string `json:"ISAKMPHashAlgorithm,omitempty"`
	Name                              string `json:"name,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	Sequence                          int    `json:"sequence,omitempty"`
	Description                       string `json:"description,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	IpsecSAReplayWindowSizeValue      int    `json:"ipsecSAReplayWindowSizeValue,omitempty"`
	AssociatedEnterpriseID            string `json:"associatedEnterpriseID,omitempty"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewIKEEncryptionprofile returns a new *IKEEncryptionprofile
func NewIKEEncryptionprofile() *IKEEncryptionprofile {

	return &IKEEncryptionprofile{
		DPDInterval:                       0,
		DPDMode:                           "REPLY_ONLY",
		DPDTimeout:                        0,
		IPsecAuthenticationAlgorithm:      "HMAC_SHA256",
		IPsecEncryptionAlgorithm:          "AES256",
		IPsecSALifetime:                   3600,
		IPsecSAReplayWindowSize:           "WINDOW_SIZE_32",
		ISAKMPAuthenticationMode:          "PRE_SHARED_KEY",
		ISAKMPDiffieHelmanGroupIdentifier: "GROUP_5_1536_BIT_DH",
		ISAKMPEncryptionAlgorithm:         "AES256",
		ISAKMPEncryptionKeyLifetime:       28800,
		ISAKMPHashAlgorithm:               "SHA256",
		IpsecSAReplayWindowSizeValue:      32,
	}
}

// Identity returns the Identity of the object.
func (o *IKEEncryptionprofile) Identity() bambou.Identity {

	return IKEEncryptionprofileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEEncryptionprofile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEEncryptionprofile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEEncryptionprofile from the server
func (o *IKEEncryptionprofile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEEncryptionprofile into the server
func (o *IKEEncryptionprofile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEEncryptionprofile from the server
func (o *IKEEncryptionprofile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEEncryptionprofile
func (o *IKEEncryptionprofile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEEncryptionprofile
func (o *IKEEncryptionprofile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEEncryptionprofile
func (o *IKEEncryptionprofile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEEncryptionprofile
func (o *IKEEncryptionprofile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
