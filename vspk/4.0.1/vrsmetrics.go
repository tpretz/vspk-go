/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VrsMetricsIdentity represents the Identity of the object
var VrsMetricsIdentity = bambou.Identity{
	Name:     "vrsmetrics",
	Category: "vrsmetrics",
}

// VrsMetricsList represents a list of VrsMetrics
type VrsMetricsList []*VrsMetrics

// VrsMetricsAncestor is the interface that an ancestor of a VrsMetrics must implement.
// An Ancestor is defined as an entity that has VrsMetrics as a descendant.
// An Ancestor can get a list of its child VrsMetrics, but not necessarily create one.
type VrsMetricsAncestor interface {
	VrsMetrics(*bambou.FetchingInfo) (VrsMetricsList, *bambou.Error)
}

// VrsMetricsParent is the interface that a parent of a VrsMetrics must implement.
// A Parent is defined as an entity that has VrsMetrics as a child.
// A Parent is an Ancestor which can create a VrsMetrics.
type VrsMetricsParent interface {
	VrsMetricsAncestor
	CreateVrsMetrics(*VrsMetrics) *bambou.Error
}

// VrsMetrics represents the model of a vrsmetrics
type VrsMetrics struct {
	ID                       string  `json:"ID,omitempty"`
	ParentID                 string  `json:"parentID,omitempty"`
	ParentType               string  `json:"parentType,omitempty"`
	Owner                    string  `json:"owner,omitempty"`
	ALUbr0Status             bool    `json:"ALUbr0Status"`
	CPUUtilization           float64 `json:"CPUUtilization,omitempty"`
	VRSProcess               bool    `json:"VRSProcess"`
	VRSVSCStatus             bool    `json:"VRSVSCStatus"`
	ReceivingMetrics         bool    `json:"receivingMetrics"`
	MemoryUtilization        float64 `json:"memoryUtilization,omitempty"`
	JesxmonProcess           bool    `json:"jesxmonProcess"`
	AgentName                string  `json:"agentName,omitempty"`
	AssocVCenterHypervisorID string  `json:"assocVCenterHypervisorID,omitempty"`
}

// NewVrsMetrics returns a new *VrsMetrics
func NewVrsMetrics() *VrsMetrics {

	return &VrsMetrics{}
}

// Identity returns the Identity of the object.
func (o *VrsMetrics) Identity() bambou.Identity {

	return VrsMetricsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VrsMetrics) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VrsMetrics) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VrsMetrics from the server
func (o *VrsMetrics) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VrsMetrics into the server
func (o *VrsMetrics) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VrsMetrics from the server
func (o *VrsMetrics) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
