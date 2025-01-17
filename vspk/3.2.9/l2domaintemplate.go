/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// L2DomainTemplateIdentity represents the Identity of the object
var L2DomainTemplateIdentity = bambou.Identity{
	Name:     "l2domaintemplate",
	Category: "l2domaintemplates",
}

// L2DomainTemplatesList represents a list of L2DomainTemplates
type L2DomainTemplatesList []*L2DomainTemplate

// L2DomainTemplatesAncestor is the interface that an ancestor of a L2DomainTemplate must implement.
// An Ancestor is defined as an entity that has L2DomainTemplate as a descendant.
// An Ancestor can get a list of its child L2DomainTemplates, but not necessarily create one.
type L2DomainTemplatesAncestor interface {
	L2DomainTemplates(*bambou.FetchingInfo) (L2DomainTemplatesList, *bambou.Error)
}

// L2DomainTemplatesParent is the interface that a parent of a L2DomainTemplate must implement.
// A Parent is defined as an entity that has L2DomainTemplate as a child.
// A Parent is an Ancestor which can create a L2DomainTemplate.
type L2DomainTemplatesParent interface {
	L2DomainTemplatesAncestor
	CreateL2DomainTemplate(*L2DomainTemplate) *bambou.Error
}

// L2DomainTemplate represents the model of a l2domaintemplate
type L2DomainTemplate struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	DHCPManaged                     bool   `json:"DHCPManaged"`
	IPType                          string `json:"IPType,omitempty"`
	Name                            string `json:"name,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	Gateway                         string `json:"gateway,omitempty"`
	Address                         string `json:"address,omitempty"`
	Description                     string `json:"description,omitempty"`
	Netmask                         string `json:"netmask,omitempty"`
	Encryption                      string `json:"encryption,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	PolicyChangeStatus              string `json:"policyChangeStatus,omitempty"`
	AssociatedMulticastChannelMapID string `json:"associatedMulticastChannelMapID,omitempty"`
	Multicast                       string `json:"multicast,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
}

// NewL2DomainTemplate returns a new *L2DomainTemplate
func NewL2DomainTemplate() *L2DomainTemplate {

	return &L2DomainTemplate{}
}

// Identity returns the Identity of the object.
func (o *L2DomainTemplate) Identity() bambou.Identity {

	return L2DomainTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *L2DomainTemplate) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *L2DomainTemplate) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the L2DomainTemplate from the server
func (o *L2DomainTemplate) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the L2DomainTemplate into the server
func (o *L2DomainTemplate) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the L2DomainTemplate from the server
func (o *L2DomainTemplate) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// L2Domains retrieves the list of child L2Domains of the L2DomainTemplate
func (o *L2DomainTemplate) L2Domains(info *bambou.FetchingInfo) (L2DomainsList, *bambou.Error) {

	var list L2DomainsList
	err := bambou.CurrentSession().FetchChildren(o, L2DomainIdentity, &list, info)
	return list, err
}

// AddressRanges retrieves the list of child AddressRanges of the L2DomainTemplate
func (o *L2DomainTemplate) AddressRanges(info *bambou.FetchingInfo) (AddressRangesList, *bambou.Error) {

	var list AddressRangesList
	err := bambou.CurrentSession().FetchChildren(o, AddressRangeIdentity, &list, info)
	return list, err
}

// CreateAddressRange creates a new child AddressRange under the L2DomainTemplate
func (o *L2DomainTemplate) CreateAddressRange(child *AddressRange) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// RedirectionTargetTemplates retrieves the list of child RedirectionTargetTemplates of the L2DomainTemplate
func (o *L2DomainTemplate) RedirectionTargetTemplates(info *bambou.FetchingInfo) (RedirectionTargetTemplatesList, *bambou.Error) {

	var list RedirectionTargetTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, RedirectionTargetTemplateIdentity, &list, info)
	return list, err
}

// CreateRedirectionTargetTemplate creates a new child RedirectionTargetTemplate under the L2DomainTemplate
func (o *L2DomainTemplate) CreateRedirectionTargetTemplate(child *RedirectionTargetTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Permissions retrieves the list of child Permissions of the L2DomainTemplate
func (o *L2DomainTemplate) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the L2DomainTemplate
func (o *L2DomainTemplate) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the L2DomainTemplate
func (o *L2DomainTemplate) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the L2DomainTemplate
func (o *L2DomainTemplate) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EgressACLTemplates retrieves the list of child EgressACLTemplates of the L2DomainTemplate
func (o *L2DomainTemplate) EgressACLTemplates(info *bambou.FetchingInfo) (EgressACLTemplatesList, *bambou.Error) {

	var list EgressACLTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, EgressACLTemplateIdentity, &list, info)
	return list, err
}

// CreateEgressACLTemplate creates a new child EgressACLTemplate under the L2DomainTemplate
func (o *L2DomainTemplate) CreateEgressACLTemplate(child *EgressACLTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the L2DomainTemplate
func (o *L2DomainTemplate) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the L2DomainTemplate
func (o *L2DomainTemplate) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IngressACLTemplates retrieves the list of child IngressACLTemplates of the L2DomainTemplate
func (o *L2DomainTemplate) IngressACLTemplates(info *bambou.FetchingInfo) (IngressACLTemplatesList, *bambou.Error) {

	var list IngressACLTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, IngressACLTemplateIdentity, &list, info)
	return list, err
}

// CreateIngressACLTemplate creates a new child IngressACLTemplate under the L2DomainTemplate
func (o *L2DomainTemplate) CreateIngressACLTemplate(child *IngressACLTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IngressAdvFwdTemplates retrieves the list of child IngressAdvFwdTemplates of the L2DomainTemplate
func (o *L2DomainTemplate) IngressAdvFwdTemplates(info *bambou.FetchingInfo) (IngressAdvFwdTemplatesList, *bambou.Error) {

	var list IngressAdvFwdTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, IngressAdvFwdTemplateIdentity, &list, info)
	return list, err
}

// CreateIngressAdvFwdTemplate creates a new child IngressAdvFwdTemplate under the L2DomainTemplate
func (o *L2DomainTemplate) CreateIngressAdvFwdTemplate(child *IngressAdvFwdTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// IngressExternalServiceTemplates retrieves the list of child IngressExternalServiceTemplates of the L2DomainTemplate
func (o *L2DomainTemplate) IngressExternalServiceTemplates(info *bambou.FetchingInfo) (IngressExternalServiceTemplatesList, *bambou.Error) {

	var list IngressExternalServiceTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, IngressExternalServiceTemplateIdentity, &list, info)
	return list, err
}

// CreateIngressExternalServiceTemplate creates a new child IngressExternalServiceTemplate under the L2DomainTemplate
func (o *L2DomainTemplate) CreateIngressExternalServiceTemplate(child *IngressExternalServiceTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// CreateJob creates a new child Job under the L2DomainTemplate
func (o *L2DomainTemplate) CreateJob(child *Job) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// PolicyGroupTemplates retrieves the list of child PolicyGroupTemplates of the L2DomainTemplate
func (o *L2DomainTemplate) PolicyGroupTemplates(info *bambou.FetchingInfo) (PolicyGroupTemplatesList, *bambou.Error) {

	var list PolicyGroupTemplatesList
	err := bambou.CurrentSession().FetchChildren(o, PolicyGroupTemplateIdentity, &list, info)
	return list, err
}

// CreatePolicyGroupTemplate creates a new child PolicyGroupTemplate under the L2DomainTemplate
func (o *L2DomainTemplate) CreatePolicyGroupTemplate(child *PolicyGroupTemplate) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// QOSs retrieves the list of child QOSs of the L2DomainTemplate
func (o *L2DomainTemplate) QOSs(info *bambou.FetchingInfo) (QOSsList, *bambou.Error) {

	var list QOSsList
	err := bambou.CurrentSession().FetchChildren(o, QOSIdentity, &list, info)
	return list, err
}

// CreateQOS creates a new child QOS under the L2DomainTemplate
func (o *L2DomainTemplate) CreateQOS(child *QOS) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Groups retrieves the list of child Groups of the L2DomainTemplate
func (o *L2DomainTemplate) Groups(info *bambou.FetchingInfo) (GroupsList, *bambou.Error) {

	var list GroupsList
	err := bambou.CurrentSession().FetchChildren(o, GroupIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the L2DomainTemplate
func (o *L2DomainTemplate) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
