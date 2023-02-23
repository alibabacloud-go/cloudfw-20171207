// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAddressBookRequest struct {
	// The addresses that you want to add to the address book. Separate multiple addresses with commas (,).
	//
	// > If you set GroupType to `ip`, `port` or `domain`, you must specify the AddressList parameter.
	// >
	// > * If you set GroupType to `ip`, you must add IP addresses to the address book. Example: 192.0.XX.XX/32, 192.0.XX.XX/24.
	// > * If you set GroupType to `port`, you must add port numbers or port ranges to the address book. Example: 80, 100/200.
	// > * If you set GroupType to `domain`, you must add domain names to the address book. Example: example.com, aliyundoc.com.
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of ECS instances to the address book if the instances match the specified tags. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no (default)
	AutoAddTagEcs *string `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book. Valid values:
	//
	// * **ip**: IP address book
	// * **domain**: domain address book
	// * **port**: port address book
	// * **tag**: ECS tag-based address book
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ECS tags that you want to match.
	TagList []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relation among the ECS tags that you want to match. Valid values:
	//
	// *   **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book. This is the default value.
	// *   **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the address book.
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s AddAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookRequest) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequest) SetAddressList(v string) *AddAddressBookRequest {
	s.AddressList = &v
	return s
}

func (s *AddAddressBookRequest) SetAutoAddTagEcs(v string) *AddAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *AddAddressBookRequest) SetDescription(v string) *AddAddressBookRequest {
	s.Description = &v
	return s
}

func (s *AddAddressBookRequest) SetGroupName(v string) *AddAddressBookRequest {
	s.GroupName = &v
	return s
}

func (s *AddAddressBookRequest) SetGroupType(v string) *AddAddressBookRequest {
	s.GroupType = &v
	return s
}

func (s *AddAddressBookRequest) SetLang(v string) *AddAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *AddAddressBookRequest) SetSourceIp(v string) *AddAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAddressBookRequest) SetTagList(v []*AddAddressBookRequestTagList) *AddAddressBookRequest {
	s.TagList = v
	return s
}

func (s *AddAddressBookRequest) SetTagRelation(v string) *AddAddressBookRequest {
	s.TagRelation = &v
	return s
}

type AddAddressBookRequestTagList struct {
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s AddAddressBookRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestTagList) SetTagKey(v string) *AddAddressBookRequestTagList {
	s.TagKey = &v
	return s
}

func (s *AddAddressBookRequestTagList) SetTagValue(v string) *AddAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

type AddAddressBookResponseBody struct {
	// The UUID of the returned address book.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *AddAddressBookResponseBody) SetGroupUuid(v string) *AddAddressBookResponseBody {
	s.GroupUuid = &v
	return s
}

func (s *AddAddressBookResponseBody) SetRequestId(v string) *AddAddressBookResponseBody {
	s.RequestId = &v
	return s
}

type AddAddressBookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookResponse) GoString() string {
	return s.String()
}

func (s *AddAddressBookResponse) SetHeaders(v map[string]*string) *AddAddressBookResponse {
	s.Headers = v
	return s
}

func (s *AddAddressBookResponse) SetStatusCode(v int32) *AddAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAddressBookResponse) SetBody(v *AddAddressBookResponseBody) *AddAddressBookResponse {
	s.Body = v
	return s
}

type AddControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// * **accept**: allows the traffic.
	// * **drop**: denies the traffic.
	// * **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The type of the application that the access control policy supports. Valid values:
	//
	// * **FTP**
	// * **HTTP**
	// * **HTTPS**
	// * **Memcache**
	// * **MongoDB**
	// * **MQTT**
	// * **MySQL**
	// * **RDP**
	// * **Redis**
	// * **SMTP**
	// * **SMTPS**
	// * **SSH**
	// * **SSL_No_Cert**
	// * **SSL**
	// * **VNC**
	// * **ANY**: all types of applications
	//
	// > The value of this parameter depends on the value of Proto. If Proto is set to TCP, you can set ApplicationName to any valid value. If Proto is set to UDP, ICMP, or ANY, you can set ApplicationName only to ANY.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The types of the application that the access control policy supports.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy. Valid values:
	//
	// * If Proto is set to ICMP, the value of DestPort is empty.
	//
	// > If Proto is set to ICMP, access control does not take effect on the destination port.
	//
	// * If Proto is set to TCP, UDP, or ANY and DestPortType is set to group, the value of DestPort is empty.
	//
	// > If DestPortType is set to group, you do not need to specify the destination port number. All ports that the access control policy controls are included in the destination port address book.
	//
	// * If Proto is set to TCP, UDP, or ANY and DestPortType is set to port, the value of DestPort is the destination port number.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// >  If DestPortType is set to group, you must specify the name of the destination port address book.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy.
	//
	// Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// Valid values:
	//
	// * If DestinationType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.2.XX.XX/24
	//
	// * If DestinationType is set to group, the value of this parameter is an address book.
	//
	//     Example: db_group
	//
	// * If DestinationType is set to domain, the value of this parameter is a domain name.
	//
	//     Example: \*.aliyuncs.com
	//
	// * If DestinationType is set to location, the value of this parameter is a location.
	//
	//     Example: \["BJ11", "ZB"]
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// * **net**: destination CIDR block
	// * **group**: destination address book
	// * **domain**: destination domain name
	// * **location**: destination location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// * **in**: inbound traffic
	// * **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version of the address in the access control policy.
	//
	// Valid values:
	//
	// * **4**: IPv4
	// * **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// * **zh**: Chinese (default)
	// * **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// * **ANY**: any protocol type
	// * **TCP**
	// * **UDP**
	// * **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**: The access control policy is enabled.
	// *   **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// * If SourceType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.1.XX.XX/24
	//
	// * If SourceType is set to group, the value of this parameter is an address book.
	//
	//     Example: db_group
	//
	// * If SourceType is set to location, the value of this parameter is a location.
	//
	//     Example: \["BJ11", "ZB"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the source address book in the access control policy. Valid values:
	//
	// * **net**: source CIDR block
	// * **group**: source address book
	// * **location**: source location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s AddControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddControlPolicyRequest) SetAclAction(v string) *AddControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationName(v string) *AddControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationNameList(v []*string) *AddControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *AddControlPolicyRequest) SetDescription(v string) *AddControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPort(v string) *AddControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortGroup(v string) *AddControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortType(v string) *AddControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestination(v string) *AddControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestinationType(v string) *AddControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDirection(v string) *AddControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *AddControlPolicyRequest) SetIpVersion(v string) *AddControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *AddControlPolicyRequest) SetLang(v string) *AddControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *AddControlPolicyRequest) SetNewOrder(v string) *AddControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *AddControlPolicyRequest) SetProto(v string) *AddControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *AddControlPolicyRequest) SetRelease(v string) *AddControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *AddControlPolicyRequest) SetSource(v string) *AddControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceIp(v string) *AddControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceType(v string) *AddControlPolicyRequest {
	s.SourceType = &v
	return s
}

type AddControlPolicyResponseBody struct {
	// The ID of the access control policy that is created on the Internet firewall.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponseBody) SetAclUuid(v string) *AddControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *AddControlPolicyResponseBody) SetRequestId(v string) *AddControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AddControlPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponse) SetHeaders(v map[string]*string) *AddControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddControlPolicyResponse) SetStatusCode(v int32) *AddControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddControlPolicyResponse) SetBody(v *AddControlPolicyResponseBody) *AddControlPolicyResponse {
	s.Body = v
	return s
}

type AddInstanceMembersRequest struct {
	// The members that you want to add to Cloud Firewall.
	Members []*AddInstanceMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s AddInstanceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersRequest) SetMembers(v []*AddInstanceMembersRequestMembers) *AddInstanceMembersRequest {
	s.Members = v
	return s
}

type AddInstanceMembersRequestMembers struct {
	// The remarks of member that you want to add to Cloud Firewall. The remarks must be 1 to 256 characters in length.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The UID of member that you want to add to Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s AddInstanceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersRequestMembers) SetMemberDesc(v string) *AddInstanceMembersRequestMembers {
	s.MemberDesc = &v
	return s
}

func (s *AddInstanceMembersRequestMembers) SetMemberUid(v int64) *AddInstanceMembersRequestMembers {
	s.MemberUid = &v
	return s
}

type AddInstanceMembersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddInstanceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersResponseBody) SetRequestId(v string) *AddInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

type AddInstanceMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddInstanceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersResponse) SetHeaders(v map[string]*string) *AddInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddInstanceMembersResponse) SetStatusCode(v int32) *AddInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceMembersResponse) SetBody(v *AddInstanceMembersResponseBody) *AddInstanceMembersResponse {
	s.Body = v
	return s
}

type BatchCopyVpcFirewallControlPolicyRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy group of the source VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
	// *   If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the IDs of policy groups.
	SourceVpcFirewallId *string `json:"SourceVpcFirewallId,omitempty" xml:"SourceVpcFirewallId,omitempty"`
	// The ID of the policy group of the destination VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a VBR, or a CCN instance.
	// *   If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the IDs of policy groups.
	TargetVpcFirewallId *string `json:"TargetVpcFirewallId,omitempty" xml:"TargetVpcFirewallId,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetLang(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetSourceIp(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetSourceVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.SourceVpcFirewallId = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetTargetVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.TargetVpcFirewallId = &v
	return s
}

type BatchCopyVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *BatchCopyVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type BatchCopyVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchCopyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCopyVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *BatchCopyVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *BatchCopyVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetBody(v *BatchCopyVpcFirewallControlPolicyResponseBody) *BatchCopyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallCenConfigureRequest struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// *   **open**: After you create the VPC firewall, the VPC firewall is automatically enabled. This is the default value.
	// *   **close**: After you create the VPC firewall, the VPC firewall is disabled. You can call the [ModifyVpcFirewallCenSwitchStatus](~~345780~~) operation to manually enable the VPC firewall.
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the VPC for which you want to create the VPC firewall.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The ID of the region to which the VPC belongs.
	//
	// >  For more information about the regions, see [Supported regions](~~195657~~).
	VpcRegion *string `json:"VpcRegion,omitempty" xml:"VpcRegion,omitempty"`
}

func (s CreateVpcFirewallCenConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureRequest) SetCenId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.CenId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallSwitch(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetLang(v string) *CreateVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetNetworkInstanceId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVSwitchId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcRegion(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcRegion = &v
	return s
}

type CreateVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallCenConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallCenConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcFirewallCenConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) SetBody(v *CreateVpcFirewallCenConfigureResponseBody) *CreateVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallConfigureRequest struct {
	// The status of the VPC firewall after you create the firewall. Valid values:
	//
	// *   **open**: After you create the VPC firewall, the firewall is automatically enabled. This is the default value.
	// *   **close**: After you create the VPC firewall, the firewall is not automatically enabled. To enable the firewall, you can call the [ModifyVpcFirewallSwitchStatus](~~342935~~) operation.
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The CIDR blocks of the local VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the local VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the local VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the local VPC.
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	// The ID of the local VPC.
	LocalVpcId *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	// The region ID of the local VPC.
	//
	// >  For more information about regions in which Cloud Firewall is supported, see [Supported regions](~~195657~~).
	LocalVpcRegion *string `json:"LocalVpcRegion,omitempty" xml:"LocalVpcRegion,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The CIDR blocks of the peer VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the peer VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the peer VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the peer VPC.
	PeerVpcCidrTableList *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	// The ID of the peer VPC.
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The region ID of the peer VPC.
	//
	// >  For more information about regions in which Cloud Firewall is supported, see [Supported regions](~~195657~~).
	PeerVpcRegion *string `json:"PeerVpcRegion,omitempty" xml:"PeerVpcRegion,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s CreateVpcFirewallConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureRequest) SetFirewallSwitch(v string) *CreateVpcFirewallConfigureRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLang(v string) *CreateVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcCidrTableList = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcId(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcId = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcRegion(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcRegion = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcCidrTableList = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcId(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcId = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcRegion(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcRegion = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

type CreateVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcFirewallConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) SetBody(v *CreateVpcFirewallConfigureResponseBody) *CreateVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// - **accept**: allows the traffic.
	// - **drop**: blocks the traffic.
	// - **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The type of the applications that the access control policy supports. Valid values:
	//
	// - **FTP**
	// - **HTTP**
	// - **HTTPS**
	// - **MySQL**
	// - **SMTP**
	// - **SMTPS**
	// - **RDP**
	// - **VNC**
	// - **SSH**
	// - **Redis**
	// - **MQTT**
	// - **MongoDB**
	// - **Memcache**
	// - **SSL**
	// - **ANY**: all types of applications
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	//
	// >  If **DestPortType** is set to `port`, you must specify this parameter.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// >  If **DestPortType** is set to `group`, you must specify this parameter.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// - **port**: port
	// - **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Valid values:
	//
	// - If **DestinationType** is set to `net`, the value of **Destination** must be a CIDR block.
	// - If **DestinationType** is set to `group`, the value of **Destination** must be an address book.
	// - If **DestinationType** is set to `domain`, the value of **Destination** must be a domain name.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	// - **group**: address book
	// - **domain**: domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// - **ANY** (If you are not sure about the protocol type, you can set this parameter to ANY.)
	// - **TCP**
	// - **UDP**
	// - **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// - **true**: enables the access control policy.
	// - **false**: disables the access control policy.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy.
	//
	// - If SourceType is set to `net`, the value of Source must be a CIDR block.
	// - If SourceType is set to `group`, the value of Source must be an address book.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	// - **group**: address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the policy group in which you want to create the access control policy.
	//
	// - If a VPC firewall protects the traffic between two VPCs that are connected by using a CEN instance, the value of this parameter must be the ID of the CEN instance.
	// - If a VPC firewall protects the traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](https://www.alibabacloud.com/help/en/cloud-firewall/latest/describevpcfirewallaclgrouplist) operation to query the IDs.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyRequest) SetAclAction(v string) *CreateVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetApplicationName(v string) *CreateVpcFirewallControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDescription(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPort(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestination(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestinationType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetLang(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetMemberUid(v string) *CreateVpcFirewallControlPolicyRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetNewOrder(v string) *CreateVpcFirewallControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetProto(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRelease(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSource(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSourceType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *CreateVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallControlPolicyResponseBody struct {
	// The ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetAclUuid(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *CreateVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) SetBody(v *CreateVpcFirewallControlPolicyResponseBody) *CreateVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteAddressBookRequest struct {
	// The ID of the address book.
	//
	// To delete the address book, you must provide the ID of the address book. You can call the DescribeAddressBook operation to query the ID.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookRequest) SetGroupUuid(v string) *DeleteAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *DeleteAddressBookRequest) SetLang(v string) *DeleteAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAddressBookRequest) SetSourceIp(v string) *DeleteAddressBookRequest {
	s.SourceIp = &v
	return s
}

type DeleteAddressBookResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookResponseBody) SetRequestId(v string) *DeleteAddressBookResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAddressBookResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookResponse) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookResponse) SetHeaders(v map[string]*string) *DeleteAddressBookResponse {
	s.Headers = v
	return s
}

func (s *DeleteAddressBookResponse) SetStatusCode(v int32) *DeleteAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAddressBookResponse) SetBody(v *DeleteAddressBookResponseBody) *DeleteAddressBookResponse {
	s.Body = v
	return s
}

type DeleteControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// To delete an access control policy, you must provide the ID of the policy. You can call the [DescribeControlPolicy](~~138866~~) operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the traffic.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) SetAclUuid(v string) *DeleteControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetDirection(v string) *DeleteControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetLang(v string) *DeleteControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetSourceIp(v string) *DeleteControlPolicyRequest {
	s.SourceIp = &v
	return s
}

type DeleteControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponseBody) SetRequestId(v string) *DeleteControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyResponse) SetStatusCode(v int32) *DeleteControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyResponse) SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteInstanceMembersRequest struct {
	// The unique identifiers (UID) of members that you want to remove from Cloud Firewall.
	MemberUids []*int64 `json:"MemberUids,omitempty" xml:"MemberUids,omitempty" type:"Repeated"`
}

func (s DeleteInstanceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersRequest) SetMemberUids(v []*int64) *DeleteInstanceMembersRequest {
	s.MemberUids = v
	return s
}

type DeleteInstanceMembersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersResponseBody) SetRequestId(v string) *DeleteInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceMembersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersResponse) SetHeaders(v map[string]*string) *DeleteInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceMembersResponse) SetStatusCode(v int32) *DeleteInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceMembersResponse) SetBody(v *DeleteInstanceMembersResponseBody) *DeleteInstanceMembersResponse {
	s.Body = v
	return s
}

type DeleteVpcFirewallCenConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The list of VPC firewall IDs.
	VpcFirewallIdList []*string `json:"VpcFirewallIdList,omitempty" xml:"VpcFirewallIdList,omitempty" type:"Repeated"`
}

func (s DeleteVpcFirewallCenConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetLang(v string) *DeleteVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetMemberUid(v string) *DeleteVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallCenConfigureRequest {
	s.VpcFirewallIdList = v
	return s
}

type DeleteVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallCenConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *DeleteVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcFirewallCenConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *DeleteVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetBody(v *DeleteVpcFirewallCenConfigureResponseBody) *DeleteVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

type DeleteVpcFirewallConfigureRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The list of the VPC firewall IDs.
	VpcFirewallIdList []*string `json:"VpcFirewallIdList,omitempty" xml:"VpcFirewallIdList,omitempty" type:"Repeated"`
}

func (s DeleteVpcFirewallConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureRequest) SetLang(v string) *DeleteVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) SetMemberUid(v string) *DeleteVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallConfigureRequest {
	s.VpcFirewallIdList = v
	return s
}

type DeleteVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureResponseBody) SetRequestId(v string) *DeleteVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcFirewallConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) SetStatusCode(v int32) *DeleteVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) SetBody(v *DeleteVpcFirewallConfigureResponseBody) *DeleteVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

type DeleteVpcFirewallControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// To delete an access control policy, you must provide the ID of the policy. You can call the **DescribeVpcFirewallControlPolicy** operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the group to which the access control policy belongs. You can call the **DescribeVpcFirewallAclGroupList** operation to query the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a CEN instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall.
	//
	// Example: vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetLang(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type DeleteVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *DeleteVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *DeleteVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetBody(v *DeleteVpcFirewallControlPolicyResponseBody) *DeleteVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeAddressBookRequest struct {
	// The port that is included in the address book. This parameter takes effect only when the **GroupType** parameter is set to **port**.
	ContainPort *string `json:"ContainPort,omitempty" xml:"ContainPort,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the address book. Valid values:
	//
	// * **ip**: IP address book
	// * **domain**: domain address book
	// * **port**: port address book
	// * **tag**: Elastic Compute Service (ECS) tag-based address book
	// * **allCloud**: cloud service address book
	// * **threat**: threat intelligence address book
	//
	// > If you do not specify a type, the domain address books and ECS tag-based address books are queried.
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language of the content within the request. Valid values:
	//
	// * **zh**: Chinese (default)
	// * **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition that is used to search for the address book.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s DescribeAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookRequest) SetContainPort(v string) *DescribeAddressBookRequest {
	s.ContainPort = &v
	return s
}

func (s *DescribeAddressBookRequest) SetCurrentPage(v string) *DescribeAddressBookRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAddressBookRequest) SetGroupType(v string) *DescribeAddressBookRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookRequest) SetLang(v string) *DescribeAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAddressBookRequest) SetPageSize(v string) *DescribeAddressBookRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookRequest) SetQuery(v string) *DescribeAddressBookRequest {
	s.Query = &v
	return s
}

type DescribeAddressBookResponseBody struct {
	// The information about the address book.
	Acls []*DescribeAddressBookResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The page number of the current page.
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned address books.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBody) SetAcls(v []*DescribeAddressBookResponseBodyAcls) *DescribeAddressBookResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeAddressBookResponseBody) SetPageNo(v string) *DescribeAddressBookResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetPageSize(v string) *DescribeAddressBookResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetRequestId(v string) *DescribeAddressBookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetTotalCount(v string) *DescribeAddressBookResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAddressBookResponseBodyAcls struct {
	// The addresses in the address book.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The number of addresses in the address book.
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// Indicates whether the public IP addresses of ECS instances are automatically added to the address book if the instances match the specified tags. The setting takes effect on both newly purchased ECS instances whose tag settings are complete and ECS instances whose tag settings are modified. Valid values:
	//
	// * **1**: yes
	// * **0**: no
	AutoAddTagEcs *int32 `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book. Valid values:
	//
	// * **ip**: IP address book
	// * **domain**: domain address book
	// * **port**: port address book
	// * **tag**: ECS tag-based address book
	// * **allCloud**: cloud service address book
	// * **threat**: threat intelligence address book
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The unique ID of the address book.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The number of times that the address book is referenced.
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The details about the ECS tags that can be automatically added to the address book.
	TagList []*DescribeAddressBookResponseBodyAclsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among ECS tags. Valid values:
	//
	// * **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	// * **or**: The public IP addresses of ECS instances that match any of the specified tags can be added to the address book.
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s DescribeAddressBookResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddressList(v []*string) *DescribeAddressBookResponseBodyAcls {
	s.AddressList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddressListCount(v int32) *DescribeAddressBookResponseBodyAcls {
	s.AddressListCount = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAutoAddTagEcs(v int32) *DescribeAddressBookResponseBodyAcls {
	s.AutoAddTagEcs = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetDescription(v string) *DescribeAddressBookResponseBodyAcls {
	s.Description = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupName(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupName = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupType(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupUuid(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupUuid = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetReferenceCount(v int32) *DescribeAddressBookResponseBodyAcls {
	s.ReferenceCount = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagList(v []*DescribeAddressBookResponseBodyAclsTagList) *DescribeAddressBookResponseBodyAcls {
	s.TagList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagRelation(v string) *DescribeAddressBookResponseBodyAcls {
	s.TagRelation = &v
	return s
}

type DescribeAddressBookResponseBodyAclsTagList struct {
	// The key of the ECS tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the ECS tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsTagList) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsTagList) SetTagKey(v string) *DescribeAddressBookResponseBodyAclsTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsTagList) SetTagValue(v string) *DescribeAddressBookResponseBodyAclsTagList {
	s.TagValue = &v
	return s
}

type DescribeAddressBookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponse) SetHeaders(v map[string]*string) *DescribeAddressBookResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddressBookResponse) SetStatusCode(v int32) *DescribeAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddressBookResponse) SetBody(v *DescribeAddressBookResponseBody) *DescribeAddressBookResponse {
	s.Body = v
	return s
}

type DescribeAssetListRequest struct {
	// The number of the page to return.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall. Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is added in Cloud Firewall.
	MemberUid      *int64  `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// The number of entries to return on each page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which Cloud Firewall is supported.
	//
	// >  For more information about the regions in which Cloud Firewall is supported, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The type of the asset. Valid values:
	//
	// *   **BastionHostEgressIP**: the egress IP address of a bastion host
	// *   **BastionHostIngressIP**: the ingress IP address of a bastion host
	// *   **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	// *   **EcsPublicIP**: the public IP address of an ECS instance
	// *   **EIP**: the EIP
	// *   **EniEIP**: the EIP of an elastic network interface (ENI)
	// *   **NatEIP**: the EIP of a Network Address Translation (NAT) gateway
	// *   **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance
	// *   **SlbPublicIP**: the public IP address of an SLB instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **HAVIP**: the high-availability virtual IP address (HAVIP)
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The instance ID or the IP address of the asset.
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// *   **pass**: delivered
	//
	// *   **block**: undelivered
	//
	// *   **unsupport**: unsupported
	//
	// > If you do not specify this parameter, the assets on which security group policies in all states take effect are queried.
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **open**: The firewall is enabled.
	// *   **opening**: The firewall is being enabled.
	// *   **closed**: The firewall is disabled.
	// *   **closing**: The firewall is being disabled.
	//
	// >  If you do not specify this parameter, the assets that are configured for firewalls in all states are queried.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is deprecated.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The edition of Cloud Firewall. Valid values:
	//
	// *   **buy**: a paid edition (default)
	// *   **free**: a free edition
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeAssetListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetListRequest) SetCurrentPage(v string) *DescribeAssetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAssetListRequest) SetIpVersion(v string) *DescribeAssetListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetListRequest) SetLang(v string) *DescribeAssetListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetListRequest) SetMemberUid(v int64) *DescribeAssetListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetListRequest) SetNewResourceTag(v string) *DescribeAssetListRequest {
	s.NewResourceTag = &v
	return s
}

func (s *DescribeAssetListRequest) SetPageSize(v string) *DescribeAssetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAssetListRequest) SetRegionNo(v string) *DescribeAssetListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAssetListRequest) SetResourceType(v string) *DescribeAssetListRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAssetListRequest) SetSearchItem(v string) *DescribeAssetListRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribeAssetListRequest) SetSgStatus(v string) *DescribeAssetListRequest {
	s.SgStatus = &v
	return s
}

func (s *DescribeAssetListRequest) SetStatus(v string) *DescribeAssetListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAssetListRequest) SetType(v string) *DescribeAssetListRequest {
	s.Type = &v
	return s
}

func (s *DescribeAssetListRequest) SetUserType(v string) *DescribeAssetListRequest {
	s.UserType = &v
	return s
}

type DescribeAssetListResponseBody struct {
	// The details about the assets that are protected by Cloud Firewall.
	Assets []*DescribeAssetListResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets that are protected by Cloud Firewall.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAssetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponseBody) SetAssets(v []*DescribeAssetListResponseBodyAssets) *DescribeAssetListResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeAssetListResponseBody) SetRequestId(v string) *DescribeAssetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetListResponseBody) SetTotalCount(v int32) *DescribeAssetListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAssetListResponseBodyAssets struct {
	// The UID of the Alibaba Cloud account.
	//
	// >  The value of this parameter indicates the management account to which the member is added.
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The instance ID of the asset that is bound to Cloud Firewall.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The instance name of the asset that is bound to Cloud Firewall.
	BindInstanceName *string `json:"BindInstanceName,omitempty" xml:"BindInstanceName,omitempty"`
	CreateTimeStamp  *string `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The public IP address of the server.
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// The internal IP address of the server.
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4**: IPv4
	// *   **6**: IPv6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The UID of the member that is added in Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance name of the asset that is protected by Cloud Firewall.
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// The remarks of the asset. Valid values:
	//
	// *   **REGION\_NOT\_SUPPORT**: The region is not supported.
	// *   **NETWORK\_NOT\_SUPPORT**: The network is not supported.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **open**: The firewall is enabled.
	// *   **opening**: The firewall is being enabled.
	// *   **closed**: The firewall is disabled.
	// *   **closing**: The firewall is being disabled.
	ProtectStatus *string `json:"ProtectStatus,omitempty" xml:"ProtectStatus,omitempty"`
	// The ID of the region in which the asset resides.
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// Indicates whether the firewall is supported in the region in which the asset resides. Valid values:
	//
	// *   **enable**: supported
	// *   **disable**: unsupported
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The instance ID of the asset.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The type of the asset. Valid values:
	//
	// *   **BastionHostEgressIP**: the egress IP address of a bastion host
	// *   **BastionHostIngressIP**: the ingress IP address of a bastion host
	// *   **EcsEIP**: the EIP of an ECS instance
	// *   **EcsPublicIP**: the public IP address of an ECS instance
	// *   **EIP**: the EIP
	// *   **EniEIP**: the EIP of an ENI
	// *   **NatEIP**: the EIP of a NAT gateway
	// *   **SlbEIP**: the EIP of an SLB instance
	// *   **SlbPublicIP**: the public IP address of an SLB instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **HAVIP**: the HAVIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The risk level of the asset. Valid values:
	//
	// *   **low**: low
	// *   **middle**: medium
	// *   **hight**: high
	//
	// >  The value of this parameter is returned only when the UserType parameter is set to free.
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// *   **pass**: delivered
	// *   **block**: undelivered
	// *   **unsupport**: unsupported
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The time when the status of the security group was last checked. The value is a UNIX timestamp. Unit: seconds.
	SgStatusTime *int64 `json:"SgStatusTime,omitempty" xml:"SgStatusTime,omitempty"`
	// The status of traffic redirection for the asset. Valid values:
	//
	// *   **enable**: Traffic redirection is enabled.
	// *   **disable**: Traffic redirection is disabled.
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	// This parameter is deprecated.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetListResponseBodyAssets) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponseBodyAssets) SetAliUid(v int64) *DescribeAssetListResponseBodyAssets {
	s.AliUid = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetBindInstanceId(v string) *DescribeAssetListResponseBodyAssets {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetBindInstanceName(v string) *DescribeAssetListResponseBodyAssets {
	s.BindInstanceName = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetCreateTimeStamp(v string) *DescribeAssetListResponseBodyAssets {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetInternetAddress(v string) *DescribeAssetListResponseBodyAssets {
	s.InternetAddress = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetIntranetAddress(v string) *DescribeAssetListResponseBodyAssets {
	s.IntranetAddress = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetIpVersion(v int32) *DescribeAssetListResponseBodyAssets {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetMemberUid(v int64) *DescribeAssetListResponseBodyAssets {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetName(v string) *DescribeAssetListResponseBodyAssets {
	s.Name = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetNewResourceTag(v string) *DescribeAssetListResponseBodyAssets {
	s.NewResourceTag = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetNote(v string) *DescribeAssetListResponseBodyAssets {
	s.Note = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetProtectStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.ProtectStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRegionID(v string) *DescribeAssetListResponseBodyAssets {
	s.RegionID = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRegionStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.RegionStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetResourceInstanceId(v string) *DescribeAssetListResponseBodyAssets {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetResourceType(v string) *DescribeAssetListResponseBodyAssets {
	s.ResourceType = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRiskLevel(v string) *DescribeAssetListResponseBodyAssets {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSgStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SgStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSgStatusTime(v int64) *DescribeAssetListResponseBodyAssets {
	s.SgStatusTime = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSyncStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SyncStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetType(v string) *DescribeAssetListResponseBodyAssets {
	s.Type = &v
	return s
}

type DescribeAssetListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssetListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponse) SetHeaders(v map[string]*string) *DescribeAssetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetListResponse) SetStatusCode(v int32) *DescribeAssetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetListResponse) SetBody(v *DescribeAssetListResponseBody) *DescribeAssetListResponse {
	s.Body = v
	return s
}

type DescribeControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	//
	// >  If you do not specify this parameter, access control policies of all action types are queried.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// >  If you do not specify this parameter, access control policies that have descriptions are queried.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value of this parameter depends on the value of the DestinationType parameter.
	//
	// *   If DestinationType is set to `net`, the value of Destination must be a CIDR block. Example: 10.0.3.0/24.
	// *   If DestinationType is set to `domain`, the value of Destination must be a domain name. Example: aliyun.
	// *   If DestinationType is set to `group`, the value of Destination must be the name of an address book. Example: db_group.
	// *   If DestinationType is set to `location`, the value of Destination must be a location. Example: beijing.
	//
	// >  If you do not specify this parameter, access control policies of all destination address types are queried.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version of the address in the access control policy. Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// * **TCP**
	// * **UDP**
	// * **ICMP**
	// * **ANY**: all types of protocols
	//
	// >  If you do not specify this parameter, access control policies of all protocol types are queried.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**: The access control policy is enabled.
	// *   **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported. The value of this parameter depends on the value of the SourceType parameter.
	//
	// *   If SourceType is set to `net`, the value of Source must be a CIDR block. Example: 192.0.XX.XX/24.
	// *   If SourceType is set to `group`, the value of Source must be the name of an address book. Example: db_group. If the db_group address book does not contain addresses, all source addresses are queried.
	// *   If SourceType is set to `location`, the value of Source must be a location. Example: beijing.
	//
	// >  If you do not specify this parameter, access control policies of all source address types are queried.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyRequest) SetAclAction(v string) *DescribeControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetAclUuid(v string) *DescribeControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetCurrentPage(v string) *DescribeControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetDescription(v string) *DescribeControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetDestination(v string) *DescribeControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetDirection(v string) *DescribeControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetIpVersion(v string) *DescribeControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetLang(v string) *DescribeControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetPageSize(v string) *DescribeControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetProto(v string) *DescribeControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetRelease(v string) *DescribeControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetSource(v string) *DescribeControlPolicyRequest {
	s.Source = &v
	return s
}

type DescribeControlPolicyResponseBody struct {
	// The page number of the returned page.
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details about the access control policy.
	Policys []*DescribeControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned access control policies.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponseBody) SetPageNo(v string) *DescribeControlPolicyResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetPageSize(v string) *DescribeControlPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetPolicys(v []*DescribeControlPolicyResponseBodyPolicys) *DescribeControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetRequestId(v string) *DescribeControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetTotalCount(v string) *DescribeControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application ID in the access control policy.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The type of the application that the access control policy supports. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL**
	// *   **VNC**
	// *   **ANY**: all types of applications
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The names of applications.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time at which the access control policy was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. The value of this parameter depends on the value of the DestinationType parameter. Valid values:
	//
	// *   If **DestinationType** is set to **net**, the value of Destination is a CIDR block. Example: 192.0.XX.XX/24.
	// *   If **DestinationType** is set to **domain**, the value of Destination is a domain name. Example: aliyuncs.com.
	// *   If **DestinationType** is set to **group**, the value of Destination is the name of an address book. Example: db_group.
	// *   If **DestinationType** is set to **location**, the value of Destination is a location. For more information about location codes, see [AddControlPolicy](~~138867~~). Example: \["BJ11", "ZB"].
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more IP addresses
	// *   **tag**: an ECS tag-based address book that includes the IP addresses of the ECS instances with one or more specific tags
	// *   **domain**: an address book that includes one or more domain names
	// *   **threat**: an address book that includes one or more malicious IP addresses or domain names
	// *   **backsrc**: an address book that includes one or more back-to-origin addresses of Anti-DDoS Pro or Anti-DDoS Premium instances or WAF instances
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: destination CIDR block
	// *   **group**: destination address book
	// *   **domain**: destination domain name
	// *   **location**: destination location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The DNS resolution result.
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The timestamp of the DNS resolution result. The value is a UNIX timestamp. Unit: seconds.
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The timestamp when the access control policy was last hit. The value is a UNIX timestamp. Unit: seconds.
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy.
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The IP version of the address in the access control policy. Valid values:
	//
	// *   **4**: IPv4
	// *   **6**: IPv6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The time at which the access control policy was modified.
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A small priority value indicates a high priority.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Indicates whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**: The access control policy is enabled.
	// *   **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of Source is a CIDR block. Example: 192.0.XX.XX/24.
	// *   If **SourceType** is set to `group`, the value of Source is the name of an address book. Example: db_group.
	// *   If **SourceType** is set to `location`, the value of Source is a location. For more information about location codes, see [AddControlPolicy](~~138867~~). Example: \["BJ11", "ZB"].
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more IP addresses
	// *   **tag**: an Elastic Compute Service (ECS) tag-based address book that includes the IP addresses of the ECS instances with one or more specific tags
	// *   **domain**: an address book that includes one or more domain names
	// *   **threat**: an address book that includes one or more malicious IP addresses or domain names
	// *   **backsrc**: an address book that includes one or more back-to-origin addresses of Anti-DDoS Pro or Anti-DDoS Premium instances or Web Application Firewall (WAF) instances
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: source CIDR block
	// *   **group**: source address book
	// *   **location**: source location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The total quota consumed by the returned access control policies, which is the sum of the quota consumed by each policy.
	//
	// Quota that is consumed by an access control policy = Number of source CIDR blocks × Number of destination CIDR blocks, destination locations, or IP addresses that are resolved from destination domain names × Number of applications × Number of ports.
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
}

func (s DescribeControlPolicyResponseBodyPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDirection(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Direction = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDnsResult(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DnsResult = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDnsResultTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.DnsResultTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetIpVersion(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSpreadCnt(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

type DescribeControlPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeControlPolicyResponse) SetStatusCode(v int32) *DescribeControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeControlPolicyResponse) SetBody(v *DescribeControlPolicyResponseBody) *DescribeControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeDomainResolveRequest struct {
	// The domain name whose DNS record you want to query.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The natural language of the response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainResolveRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveRequest) SetDomain(v string) *DescribeDomainResolveRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetIpVersion(v string) *DescribeDomainResolveRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetLang(v string) *DescribeDomainResolveRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetSourceIp(v string) *DescribeDomainResolveRequest {
	s.SourceIp = &v
	return s
}

type DescribeDomainResolveResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the DNS record of the domain name.
	ResolveResult *DescribeDomainResolveResponseBodyResolveResult `json:"ResolveResult,omitempty" xml:"ResolveResult,omitempty" type:"Struct"`
}

func (s DescribeDomainResolveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponseBody) SetRequestId(v string) *DescribeDomainResolveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResolveResponseBody) SetResolveResult(v *DescribeDomainResolveResponseBodyResolveResult) *DescribeDomainResolveResponseBody {
	s.ResolveResult = v
	return s
}

type DescribeDomainResolveResponseBodyResolveResult struct {
	// The IP address to which the domain name is resolved. Multiple IP addresses are separated by commas (,).
	IpAddrs *string `json:"IpAddrs,omitempty" xml:"IpAddrs,omitempty"`
	// The time when the domain name was resolved. The value of this parameter is a timestamp. Unit: seconds.
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDomainResolveResponseBodyResolveResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveResponseBodyResolveResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponseBodyResolveResult) SetIpAddrs(v string) *DescribeDomainResolveResponseBodyResolveResult {
	s.IpAddrs = &v
	return s
}

func (s *DescribeDomainResolveResponseBodyResolveResult) SetUpdateTime(v int64) *DescribeDomainResolveResponseBodyResolveResult {
	s.UpdateTime = &v
	return s
}

type DescribeDomainResolveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainResolveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainResolveResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponse) SetHeaders(v map[string]*string) *DescribeDomainResolveResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResolveResponse) SetStatusCode(v int32) *DescribeDomainResolveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResolveResponse) SetBody(v *DescribeDomainResolveResponseBody) *DescribeDomainResolveResponse {
	s.Body = v
	return s
}

type DescribeInstanceMembersRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The remarks of the member in Cloud Firewall. The length is 1 ~ 256 characters.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The name of the member in Cloud Firewall.
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	// The unique identifier (UID) of the member in Cloud Firewall.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **20**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersRequest) SetCurrentPage(v string) *DescribeInstanceMembersRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberDesc(v string) *DescribeInstanceMembersRequest {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberDisplayName(v string) *DescribeInstanceMembersRequest {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberUid(v string) *DescribeInstanceMembersRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetPageSize(v string) *DescribeInstanceMembersRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceMembersResponseBody struct {
	// The information about the member in Cloud Firewall.
	Members []*DescribeInstanceMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeInstanceMembersResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBody) SetMembers(v []*DescribeInstanceMembersResponseBodyMembers) *DescribeInstanceMembersResponseBody {
	s.Members = v
	return s
}

func (s *DescribeInstanceMembersResponseBody) SetPageInfo(v *DescribeInstanceMembersResponseBodyPageInfo) *DescribeInstanceMembersResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInstanceMembersResponseBody) SetRequestId(v string) *DescribeInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceMembersResponseBodyMembers struct {
	// The time when the member was added to Cloud Firewall.
	//
	// >  The value is a UNIX timestamp. Unit: seconds.
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The remarks of the member in Cloud Firewall.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The name of the member in Cloud Firewall.
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	// The status of the member in Cloud Firewall. Valid values:
	//
	// *   **normal**
	// *   **deleting**
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// The UID of the member in Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the member in Cloud Firewall was last modified.
	//
	// >  The value is a UNIX timestamp. Unit: seconds.
	ModifyTime *int32 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeInstanceMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetCreateTime(v int32) *DescribeInstanceMembersResponseBodyMembers {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberDesc(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberDisplayName(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberStatus(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberStatus = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberUid(v int64) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetModifyTime(v int32) *DescribeInstanceMembersResponseBodyMembers {
	s.ModifyTime = &v
	return s
}

type DescribeInstanceMembersResponseBodyPageInfo struct {
	// The page number of the current page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of the members in Cloud Firewall.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceMembersResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetPageSize(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeInstanceMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponse) SetHeaders(v map[string]*string) *DescribeInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMembersResponse) SetStatusCode(v int32) *DescribeInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMembersResponse) SetBody(v *DescribeInstanceMembersResponseBody) *DescribeInstanceMembersResponse {
	s.Body = v
	return s
}

type DescribeInvadeEventListRequest struct {
	// The IP address of the affected asset.
	AssetsIP *string `json:"AssetsIP,omitempty" xml:"AssetsIP,omitempty"`
	// The ID of the instance.
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the instance.
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the breach awareness event.
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the breach awareness event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The UUID of the breach awareness event.
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// Specifies whether to ignore the breach awareness event. Valid values:
	//
	// *   **true**: ignores the breach awareness event.
	// *   **false**: does not ignore the breach awareness event.
	IsIgnore *string `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the member.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 10.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of process statuses.
	ProcessStatusList []*int32 `json:"ProcessStatusList,omitempty" xml:"ProcessStatusList,omitempty" type:"Repeated"`
	// The list of risk levels.
	RiskLevel []*int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListRequest) SetAssetsIP(v string) *DescribeInvadeEventListRequest {
	s.AssetsIP = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetAssetsInstanceId(v string) *DescribeInvadeEventListRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetAssetsInstanceName(v string) *DescribeInvadeEventListRequest {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetCurrentPage(v string) *DescribeInvadeEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEndTime(v string) *DescribeInvadeEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventKey(v string) *DescribeInvadeEventListRequest {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventName(v string) *DescribeInvadeEventListRequest {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventUuid(v string) *DescribeInvadeEventListRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetIsIgnore(v string) *DescribeInvadeEventListRequest {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetLang(v string) *DescribeInvadeEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetMemberUid(v int64) *DescribeInvadeEventListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetPageSize(v string) *DescribeInvadeEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetProcessStatusList(v []*int32) *DescribeInvadeEventListRequest {
	s.ProcessStatusList = v
	return s
}

func (s *DescribeInvadeEventListRequest) SetRiskLevel(v []*int32) *DescribeInvadeEventListRequest {
	s.RiskLevel = v
	return s
}

func (s *DescribeInvadeEventListRequest) SetSourceIp(v string) *DescribeInvadeEventListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetStartTime(v string) *DescribeInvadeEventListRequest {
	s.StartTime = &v
	return s
}

type DescribeInvadeEventListResponseBody struct {
	// An array that consists of breach awareness events.
	EventList []*DescribeInvadeEventListResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// The ratio of high-risk events.
	HighLevelPercent *int32 `json:"HighLevelPercent,omitempty" xml:"HighLevelPercent,omitempty"`
	// The ratio of low-risk events.
	LowLevelPercent *int32 `json:"LowLevelPercent,omitempty" xml:"LowLevelPercent,omitempty"`
	// The ratio of medium-risk events.
	MiddleLevelPercent *int32 `json:"MiddleLevelPercent,omitempty" xml:"MiddleLevelPercent,omitempty"`
	// The pagination information.
	PageInfo *DescribeInvadeEventListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvadeEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBody) SetEventList(v []*DescribeInvadeEventListResponseBodyEventList) *DescribeInvadeEventListResponseBody {
	s.EventList = v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetHighLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.HighLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetLowLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.LowLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetMiddleLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.MiddleLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetPageInfo(v *DescribeInvadeEventListResponseBodyPageInfo) *DescribeInvadeEventListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetRequestId(v string) *DescribeInvadeEventListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvadeEventListResponseBodyEventList struct {
	// The ID of the affected asset.
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the affected asset.
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The type of the affected asset. Valid values:
	//
	// * **BastionHostIP**: the egress IP address of a bastion host
	// * **BastionHostIngressIP**: the ingress IP address of a bastion host
	// * **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	// * **EcsPublicIP**: the public IP address of an ECS instance
	// * **EIP**: the EIP
	// * **EniEIP**: the EIP of an elastic network interface (ENI)
	// * **NatEIP**: the EIP of a NAT gateway
	// * **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance
	// * **SlbPublicIP**: the public IP address of an SLB instance
	// * **NatPublicIP**: the public IP address of a NAT gateway
	// * **HAVIP**: the high-availability virtual IP address (HAVIP)
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The ID of the breach awareness event.
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the breach awareness event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the breach awareness event. Valid values:
	//
	// *   **IPS**: intrusion prevention event
	// *   **offline**: disconnection event
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// The UUID of the breach awareness event.
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The time when the breach awareness event first occurred. The value is a UNIX timestamp. Unit: seconds.
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether the breach awareness event is ignored. Valid values:
	//
	// *   **true**: The breach awareness event is ignored.
	// *   **false**: The breach awareness event is not ignored.
	IsIgnore *bool `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// The time when the breach awareness event last occurred. The value is a UNIX timestamp. Unit: seconds.
	LastTime *int32 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The ID of the member.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The private IP address of the affected asset.
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The handling status of the breach awareness event. Valid values:
	//
	// *   **0**: unhandled
	// *   **20**: handled
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The public IP address of the affected asset.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The type of the affected asset. Valid values:
	//
	// * **BastionHostIP**: the egress IP address of a bastion host
	// * **BastionHostIngressIP**: the ingress IP address of a bastion host
	// * **EcsEIP**: the EIP of an ECS instance
	// * **EcsPublicIP**: the public IP address of an ECS instance
	// * **EIP**: the EIP
	// * **EniEIP**: the EIP of an ENI
	// * **NatEIP**: the EIP of a NAT gateway
	// * **SlbEIP**: the EIP of an SLB instance
	// * **SlbPublicIP**: the public IP address of an SLB instance
	// * **NatPublicIP**: the public IP address of a NAT gateway
	// * **HAVIP**: the HAVIP
	PublicIpType *string `json:"PublicIpType,omitempty" xml:"PublicIpType,omitempty"`
	// The risk level. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeInvadeEventListResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsInstanceId(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsInstanceName(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsType(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsType = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventKey(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventName(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventSrc(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventSrc = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventUuid(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetFirstTime(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.FirstTime = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetIsIgnore(v bool) *DescribeInvadeEventListResponseBodyEventList {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetLastTime(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.LastTime = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetMemberUid(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.MemberUid = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPrivateIP(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetProcessStatus(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPublicIP(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PublicIP = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPublicIpType(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PublicIpType = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetRiskLevel(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.RiskLevel = &v
	return s
}

type DescribeInvadeEventListResponseBodyPageInfo struct {
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of breach awareness events.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvadeEventListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetPageSize(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeInvadeEventListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInvadeEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInvadeEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventListResponse) SetStatusCode(v int32) *DescribeInvadeEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventListResponse) SetBody(v *DescribeInvadeEventListResponseBody) *DescribeInvadeEventListResponse {
	s.Body = v
	return s
}

type DescribeOutgoingDestinationIPRequest struct {
	// The application type in the access control policy. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL_No_Cert**
	// *   **SSL**
	// *   **VNC**
	//
	// >  The value of this parameter depends on the value of Proto. If you set Proto to TCP, you can set ApplicationNameList to any valid value. If you specify both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the service to which the destination IP address belongs. This parameter is left empty by default. Valid values:
	//
	// *   **All**: all services
	// *   **RiskDomain**: risky domain names
	// *   **RiskIP**: risky IP addresses
	// *   **AliYun**: Alibaba Cloud services
	// *   **NotAliYun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The destination IP address in the outbound connection that is initiated to access a domain name.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order in which you want to sort the queried information. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 10.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port number.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the ECS instance that initiates the outbound connection.
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The public IP address of the Elastic Compute Service (ECS) instance that initiates the outbound connection.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field based on which you want to sort the queried information. Valid values:
	//
	// *   **SessionCount**: the number of requests. This is the default value.
	// *   **TotalBytes**: the total volume of traffic.
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tag. Valid values:
	//
	// *   **AliYun**: Alibaba Cloud service
	// *   **RiskDomain**: risky domain name
	// *   **RiskIP**: risky IP address
	// *   **TrustedDomain**: trusted website
	// *   **AliPay**: Alipay
	// *   **DingDing**: DingTalk
	// *   **WeChat**: WeChat
	// *   **QQ**: Tencent QQ
	// *   **SecurityService**: security service
	// *   **Microsoft**: Microsoft
	// *   **Amazon**: Amazon Web Services (AWS)
	// *   **Pan**: cloud disk
	// *   **Map**: map
	// *   **Code**: code hosting
	// *   **SystemService**: system service
	// *   **Taobao**: Taobao
	// *   **Google**: Google
	// *   **ThirdPartyService**: third-party service
	// *   **FirstFlow**: the first time
	// *   **Downloader**: malicious download
	// *   **Alexa Top1M**: popular website
	// *   **Miner**: mining pool
	// *   **Intelligence**: threat intelligence
	// *   **DDoS**: DDoS trojan
	// *   **Ransomware**: ransomware
	// *   **Spyware**: spyware
	// *   **Rogue**: rogue software
	// *   **Botnet**: botnet
	// *   **Suspicious**: suspicious website
	// *   **C\&C**: command and control (C\&C)
	// *   **Gang**: gang
	// *   **CVE**: Common Vulnerabilities and Exposures (CVE)
	// *   **Backdoor**: webshell
	// *   **Phishing**: phishing website
	// *   **APT**: advanced persistent threat (APT) attack
	// *   **Supply Chain Attack**: supply chain attack
	// *   **Malicious software**: malware
	TagIdNew *string `json:"TagIdNew,omitempty" xml:"TagIdNew,omitempty"`
}

func (s DescribeOutgoingDestinationIPRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPRequest) SetApplicationName(v string) *DescribeOutgoingDestinationIPRequest {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetCategoryId(v string) *DescribeOutgoingDestinationIPRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetCurrentPage(v string) *DescribeOutgoingDestinationIPRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetDstIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetEndTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetLang(v string) *DescribeOutgoingDestinationIPRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetOrder(v string) *DescribeOutgoingDestinationIPRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPageSize(v string) *DescribeOutgoingDestinationIPRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Port = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPrivateIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPublicIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetSort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetStartTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetTagIdNew(v string) *DescribeOutgoingDestinationIPRequest {
	s.TagIdNew = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBody struct {
	// The destination IP addresses in outbound connections.
	DstIPList []*DescribeOutgoingDestinationIPResponseBodyDstIPList `json:"DstIPList,omitempty" xml:"DstIPList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of destination IP addresses in outbound connections.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetDstIPList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPList) *DescribeOutgoingDestinationIPResponseBody {
	s.DstIPList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetRequestId(v string) *DescribeOutgoingDestinationIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetTotalCount(v int32) *DescribeOutgoingDestinationIPResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPList struct {
	// Indicates whether an access control policy is configured. Valid values:
	//
	// *   **Uncovered**: No access control policies are configured.
	// *   **FullCoverage**: An access control policy is configured.
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// The suggestion in an access control policy.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The state of the access control policy. Valid values:
	//
	// *   **Normal**: healthy
	// *   **Abnormal**: unhealthy
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The information about the address book.
	AddressGroupList []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList `json:"AddressGroupList,omitempty" xml:"AddressGroupList,omitempty" type:"Repeated"`
	// An array that consists of application ports.
	ApplicationPortList []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList `json:"ApplicationPortList,omitempty" xml:"ApplicationPortList,omitempty" type:"Repeated"`
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The ID of the service to which the destination IP address belongs. Valid values:
	//
	// *   **Aliyun**: Alibaba Cloud services
	// *   **NotAliyun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of the service to which the destination IP address belongs. Valid values:
	//
	// *   **Alibaba Cloud services**
	// *   **third-party services**
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The destination IP address in the outbound connection that is initiated to access a domain name.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The name of the group to which the access control policy belongs.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an access control policy is configured. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an access control policy is recommended. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The inbound traffic. Unit: bytes.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the destination IP address is added to a whitelist. Valid values:
	//
	// *   **true**: added
	// *   **false**: not added
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The outbound traffic. Unit: bytes.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The UUID of the access control policy.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the access control policy.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The reason why the domain name is secure.
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The suggestion to handle the traffic of the domain name in outbound connections. Valid values:
	//
	// *   **pass**: allow
	// *   **alert**: deny
	// *   **drop**: monitor
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The tags.
	TagList []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total volume of traffic. Unit: bytes.
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclCoverage(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclRecommendDetail(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclStatus(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAddressGroupList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AddressGroupList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetApplicationPortList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.ApplicationPortList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryClassId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetDstIP(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetHasAcl(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.HasAcl = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetHasAclRecommend(v bool) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetInBytes(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetIsMarkNormal(v bool) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetOutBytes(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetRuleId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetRuleName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSecurityReason(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SecurityReason = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSecuritySuggest(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSessionCount(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetTagList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetTotalBytes(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.TotalBytes = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList struct {
	// The name of the address book.
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) SetAddressGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) SetAddressGroupUUID(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	s.AddressGroupUUID = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList struct {
	// The application type in the access control policy. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL_No_Cert**
	// *   **SSL**
	// *   **VNC**
	//
	// >  The value of this parameter depends on the value of Proto. If you set Proto to TCP, you can set ApplicationNameList to any valid value. If you specify both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The port of the application.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetApplicationName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetPort(v int32) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.Port = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListTagList struct {
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk level. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the tag.
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The ID of the tag.
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetClassId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetRiskLevel(v int32) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagDescribe(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagName = &v
	return s
}

type DescribeOutgoingDestinationIPResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOutgoingDestinationIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOutgoingDestinationIPResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDestinationIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) SetStatusCode(v int32) *DescribeOutgoingDestinationIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) SetBody(v *DescribeOutgoingDestinationIPResponseBody) *DescribeOutgoingDestinationIPResponse {
	s.Body = v
	return s
}

type DescribeOutgoingDomainRequest struct {
	// The type of the service. This parameter is empty by default. Valid values:
	//
	// *   **All**: all services
	// *   **RiskDomain**: risky domain names
	// *   **RiskIP**: risky IP addresses
	// *   **AliYun**: Alibaba Cloud services
	// *   **NotAliYun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain name in outbound connections.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 100.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The public IP address of the Elastic Compute Service (ECS) instance that initiates outbound connections.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field based on which you want to sort the query results. Valid values:
	//
	// *   **SessionCount**: the number of requests. This is the default value.
	// *   **TotalBytes**: the total volume of traffic.
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tag. Valid values:
	//
	// *   **AliYun**: Alibaba Cloud service
	// *   **RiskDomain**: risky domain name
	// *   **RiskIP**: risky IP address
	// *   **TrustedDomain**: trusted website
	// *   **AliPay**: Alipay
	// *   **DingDing**: DingTalk
	// *   **WeChat**: WeChat
	// *   **QQ**: Tencent QQ
	// *   **SecurityService**: security service
	// *   **Microsoft**: Microsoft
	// *   **Amazon**: Amazon Web Services (AWS)
	// *   **Pan**: cloud disk
	// *   **Map**: map
	// *   **Code**: code hosting
	// *   **SystemService**: system service
	// *   **Taobao**: Taobao
	// *   **Google**: Google
	// *   **ThirdPartyService**: third-party service
	// *   **FirstFlow**: the first time when an outbound connection is initiated
	// *   **Downloader**: malicious download
	// *   **Alexa Top1M**: popular website
	// *   **Miner**: mining pool
	// *   **Intelligence**: threat intelligence
	// *   **DDoS**: DDoS trojan
	// *   **Ransomware**: ransomware
	// *   **Spyware**: spyware
	// *   **Rogue**: rogue software
	// *   **Botnet**: botnet
	// *   **Suspicious**: suspicious website
	// *   **C\&C**: command and control (C\&C)
	// *   **Gang**: gang
	// *   **CVE**: Common Vulnerabilities and Exposures (CVE)
	// *   **Backdoor**: webshell
	// *   **Phishing**: phishing website
	// *   **APT**: advanced persistent threat (APT) attack
	// *   **Supply Chain Attack**: supply chain attack
	// *   **Malicious software**: malware
	TagIdNew *string `json:"TagIdNew,omitempty" xml:"TagIdNew,omitempty"`
}

func (s DescribeOutgoingDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainRequest) SetCategoryId(v string) *DescribeOutgoingDomainRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetCurrentPage(v string) *DescribeOutgoingDomainRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetDomain(v string) *DescribeOutgoingDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetEndTime(v string) *DescribeOutgoingDomainRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetLang(v string) *DescribeOutgoingDomainRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetOrder(v string) *DescribeOutgoingDomainRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetPageSize(v string) *DescribeOutgoingDomainRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetPublicIP(v string) *DescribeOutgoingDomainRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetSort(v string) *DescribeOutgoingDomainRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetStartTime(v string) *DescribeOutgoingDomainRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetTagIdNew(v string) *DescribeOutgoingDomainRequest {
	s.TagIdNew = &v
	return s
}

type DescribeOutgoingDomainResponseBody struct {
	// An array that consists of the domain names in outbound connections.
	DomainList []*DescribeOutgoingDomainResponseBodyDomainList `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the domain names in outbound connections.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBody) SetDomainList(v []*DescribeOutgoingDomainResponseBodyDomainList) *DescribeOutgoingDomainResponseBody {
	s.DomainList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) SetRequestId(v string) *DescribeOutgoingDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) SetTotalCount(v int32) *DescribeOutgoingDomainResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOutgoingDomainResponseBodyDomainList struct {
	// Indicates whether an access control policy is configured. Valid values:
	//
	// *   **Uncovered**: no
	// *   **FullCoverage**: yes
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// The suggestion in an access control policy.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The state of the access control policy. Valid values:
	//
	// *   **normal**: healthy
	// *   **abnormal**: unhealthy
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The name of the address book.
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
	// The website service.
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The type ID of the service to which the domain name belongs. Valid values:
	//
	// *   **Aliyun**: Alibaba Cloud services
	// *   **NotAliyun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of the service to which the domain name belongs. Valid values:
	//
	// *   **Alibaba Cloud services**
	// *   **Third-party services**
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The domain name in outbound connections.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the group to which the access control policy belongs.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an `access control policy` is configured for the domain name. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an access control policy is recommended. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The volume of inbound traffic.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the domain name is marked as normal. Valid values:
	//
	// *   **true**: normal
	// *   **false**: abnormal
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The name of the organization.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The volume of outbound traffic.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The ID of the access control policy.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the access control policy.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The reason why the domain name is secure.
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The suggestion to handle the traffic of the domain name in outbound connections. Valid values:
	//
	// *   **pass**: allow
	// *   **alert**: monitor
	// *   **drop**: deny
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// An array that consists of tags.
	TagList []*DescribeOutgoingDomainResponseBodyDomainListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total volume of traffic. Unit: bytes.
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclCoverage(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclRecommendDetail(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclStatus(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAddressGroupName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAddressGroupUUID(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AddressGroupUUID = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetBusiness(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Business = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryClassId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetDomain(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetGroupName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetHasAcl(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.HasAcl = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetHasAclRecommend(v bool) *DescribeOutgoingDomainResponseBodyDomainList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetInBytes(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetIsMarkNormal(v bool) *DescribeOutgoingDomainResponseBodyDomainList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetOrganization(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Organization = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetOutBytes(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetRuleId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetRuleName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSecurityReason(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SecurityReason = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSecuritySuggest(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSessionCount(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetTagList(v []*DescribeOutgoingDomainResponseBodyDomainListTagList) *DescribeOutgoingDomainResponseBodyDomainList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetTotalBytes(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.TotalBytes = &v
	return s
}

type DescribeOutgoingDomainResponseBodyDomainListTagList struct {
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk level. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the tag.
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The ID of the tag.
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetClassId(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetRiskLevel(v int32) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagDescribe(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagId(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagName(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagName = &v
	return s
}

type DescribeOutgoingDomainResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOutgoingDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOutgoingDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDomainResponse) SetStatusCode(v int32) *DescribeOutgoingDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDomainResponse) SetBody(v *DescribeOutgoingDomainResponseBody) *DescribeOutgoingDomainResponse {
	s.Body = v
	return s
}

type DescribePolicyAdvancedConfigRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePolicyAdvancedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigRequest) SetLang(v string) *DescribePolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyAdvancedConfigRequest) SetSourceIp(v string) *DescribePolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribePolicyAdvancedConfigResponseBody struct {
	// Indicates whether the strict mode is enabled for the access control policy. Valid values:
	//
	// *   **on**: The strict mode is enabled.
	// *   **off**: The strict mode is disabled.
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyAdvancedConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetInternetSwitch(v string) *DescribePolicyAdvancedConfigResponseBody {
	s.InternetSwitch = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetRequestId(v string) *DescribePolicyAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyAdvancedConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolicyAdvancedConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigResponse) SetHeaders(v map[string]*string) *DescribePolicyAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) SetStatusCode(v int32) *DescribePolicyAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) SetBody(v *DescribePolicyAdvancedConfigResponseBody) *DescribePolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

type DescribePolicyPriorUsedRequest struct {
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedRequest) SetDirection(v string) *DescribePolicyPriorUsedRequest {
	s.Direction = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetIpVersion(v string) *DescribePolicyPriorUsedRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetLang(v string) *DescribePolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetSourceIp(v string) *DescribePolicyPriorUsedRequest {
	s.SourceIp = &v
	return s
}

type DescribePolicyPriorUsedResponseBody struct {
	// The lowest priority of existing access control policies.
	//
	// >  The value -1 indicates the lowest priority.
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority of existing access control policies.
	//
	// >  The value 0 indicates the highest priority.
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribePolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponseBody) SetEnd(v int32) *DescribePolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetRequestId(v string) *DescribePolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetStart(v int32) *DescribePolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

type DescribePolicyPriorUsedResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolicyPriorUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribePolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyPriorUsedResponse) SetStatusCode(v int32) *DescribePolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyPriorUsedResponse) SetBody(v *DescribePolicyPriorUsedResponseBody) *DescribePolicyPriorUsedResponse {
	s.Body = v
	return s
}

type DescribeRiskEventGroupRequest struct {
	// The names of attacked applications. Set the value in the `["AttackApp1","AttackApp2"]` format.
	AttackApp []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	// The attack type of the intrusion events. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leak
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: trojan backdoor
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	//
	// > If you do not specify this parameter, the intrusion events of all attack types are queried.
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The edition of Cloud Firewall that you purchase. Valid values:
	//
	// *   **2**: Premium Edition
	// *   **3**: Enterprise Edition
	// *   **4**: Ultimate Edition
	// *   **10**: Cloud Firewall that uses the pay-as-you-go billing method
	BuyVersion *int64 `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// The number of the page to return. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the risk events.\
	// Set the value to **session**, which indicates intrusion events.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The direction of the traffic for the intrusion events. Valid values:
	//
	// *   **in**: inbound
	// *   **out**: outbound
	//
	// > If you do not specify this parameter, the intrusion events in both inbound and outbound directions are queried.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address to query. If you specify this parameter, all intrusion events with the specified destination IP address are queried.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The ID of the destination VPC.
	//
	// > If the FirewallType parameter is set to VpcFirewall, you must specify this parameter.
	DstNetworkInstanceId *string `json:"DstNetworkInstanceId,omitempty" xml:"DstNetworkInstanceId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the firewall. Valid values:
	//
	// *   **VpcFirewall**: virtual private cloud (VPC) firewall
	// *   **InternetFirewall**: Internet firewall (default)
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to query the information about the geographical locations of IP addresses.
	//
	// *   **true**: does not query the information about the geographical locations of IP addresses.
	// *   **false**: queries the information about the geographical locations of IP addresses. This is the default value.
	NoLocation *string `json:"NoLocation,omitempty" xml:"NoLocation,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **6**. Maximum value: **10**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **1**: alerting
	// *   **2**: blocking
	//
	// > If you do not specify this parameter, the intrusion events that are detected by firewalls in both states are queried.
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The module of the rule that is used to detect the intrusion events. Valid values:
	//
	// *   **1**: basic protection
	// *   **2**: virtual patching
	// *   **4**: threat intelligence
	//
	// > If you do not specify this parameter, the intrusion events that are detected by using all rules are queried.
	RuleSource *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The field based on which the results are sorted. Valid values:
	//
	// *   **VulLevel**: The results are sorted based on the risk level field. This is the default value.
	// *   **LastTime**: The results are sorted based on the most recent occurrence time.
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The source IP address to query. If you specify this parameter, all intrusion events with the specified source IP address are queried.
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The ID of the source VPC.
	//
	// > If the FirewallType parameter is set to VpcFirewall, you must specify this parameter.
	SrcNetworkInstanceId *string `json:"SrcNetworkInstanceId,omitempty" xml:"SrcNetworkInstanceId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The risk level of the intrusion events. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	//
	// > If you do not specify this parameter, the intrusion events that are at all risk levels are queried.
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeRiskEventGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupRequest) SetAttackApp(v []*string) *DescribeRiskEventGroupRequest {
	s.AttackApp = v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetAttackType(v string) *DescribeRiskEventGroupRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetBuyVersion(v int64) *DescribeRiskEventGroupRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetCurrentPage(v string) *DescribeRiskEventGroupRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDataType(v string) *DescribeRiskEventGroupRequest {
	s.DataType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDirection(v string) *DescribeRiskEventGroupRequest {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDstIP(v string) *DescribeRiskEventGroupRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDstNetworkInstanceId(v string) *DescribeRiskEventGroupRequest {
	s.DstNetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetEndTime(v string) *DescribeRiskEventGroupRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetFirewallType(v string) *DescribeRiskEventGroupRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetLang(v string) *DescribeRiskEventGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetNoLocation(v string) *DescribeRiskEventGroupRequest {
	s.NoLocation = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetOrder(v string) *DescribeRiskEventGroupRequest {
	s.Order = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetPageSize(v string) *DescribeRiskEventGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetRuleResult(v string) *DescribeRiskEventGroupRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetRuleSource(v string) *DescribeRiskEventGroupRequest {
	s.RuleSource = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSort(v string) *DescribeRiskEventGroupRequest {
	s.Sort = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSrcIP(v string) *DescribeRiskEventGroupRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSrcNetworkInstanceId(v string) *DescribeRiskEventGroupRequest {
	s.SrcNetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetStartTime(v string) *DescribeRiskEventGroupRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetVulLevel(v string) *DescribeRiskEventGroupRequest {
	s.VulLevel = &v
	return s
}

type DescribeRiskEventGroupResponseBody struct {
	// An array that consists of the details of the intrusion events.
	DataList []*DescribeRiskEventGroupResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of intrusion events.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRiskEventGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBody) SetDataList(v []*DescribeRiskEventGroupResponseBodyDataList) *DescribeRiskEventGroupResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) SetRequestId(v string) *DescribeRiskEventGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) SetTotalCount(v int32) *DescribeRiskEventGroupResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataList struct {
	// The name of the attacked application.
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The attack type of the intrusion event. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leak
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: trojan backdoor
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The description of the intrusion event.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The direction of the traffic for the intrusion events. Valid values:
	//
	// *   **in**: inbound
	// *   **out**: outbound
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address that is included in the intrusion event.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The number of intrusion events.
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The ID of the intrusion event.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the intrusion event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The time when the intrusion event was first detected. The value is a UNIX timestamp. Unit: seconds.
	FirstEventTime *int32 `json:"FirstEventTime,omitempty" xml:"FirstEventTime,omitempty"`
	// The information about the geographical location of the IP address. The value is a struct that contains the following parameters: **CityId**, **CityName**, **CountryId**, and **CountryName**.\
	// ****************
	IPLocationInfo *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo `json:"IPLocationInfo,omitempty" xml:"IPLocationInfo,omitempty" type:"Struct"`
	// The time when the intrusion event was last detected. The value is a UNIX timestamp. Unit: seconds.
	LastEventTime *int32 `json:"LastEventTime,omitempty" xml:"LastEventTime,omitempty"`
	// The information about the private IP address of the intrusion event. The value is an array that contains the following parameters: **RegionNo**, **ResourceInstanceId**, **ResourceInstanceName**, and **ResourcePrivateIP**.\
	// ****************
	ResourcePrivateIPList []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList `json:"ResourcePrivateIPList,omitempty" xml:"ResourcePrivateIPList,omitempty" type:"Repeated"`
	// The type of the public IP address in the intrusion event. Valid values:
	//
	// *   **EIP**: the elastic IP address (EIP)
	// *   **EcsPublicIP**: the public IP address of an Elastic Compute Service (ECS) instance
	// *   **EcsEIP**: the EIP of an ECS instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **NatEIP**: the EIP of a NAT gateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the rule that is used to detect the intrusion event.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **1**: alerting
	// *   **2**: blocking
	RuleResult *int32 `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The module of the rule that is used to detect the intrusion event. Valid values:
	//
	// *   **1**: basic protection
	// *   **2**: virtual patching
	// *   **4**: threat intelligence
	RuleSource *int32 `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The source IP address that is included in the intrusion event.
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The tag added to the source IP address. The tag helps identify whether the source IP address is a back-to-origin IP address for a cloud service.
	SrcIPTag *string `json:"SrcIPTag,omitempty" xml:"SrcIPTag,omitempty"`
	// The source private IP addresses of the intrusion event.
	SrcPrivateIPList []*string `json:"SrcPrivateIPList,omitempty" xml:"SrcPrivateIPList,omitempty" type:"Repeated"`
	// The tag added to the threat intelligence that is provided for major events.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The information about the destination VPC of the intrusion event. The value is a struct that contains the following parameters: **EcsInstanceId**, **EcsInstanceName**, **NetworkInstanceId**, **NetworkInstanceName**, and **RegionNo**.\
	// ********************
	VpcDstInfo *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo `json:"VpcDstInfo,omitempty" xml:"VpcDstInfo,omitempty" type:"Struct"`
	// The information about the source VPC of the intrusion event. The value is a struct that contains the following parameters: **EcsInstanceId**, **EcsInstanceName**, **NetworkInstanceId**, **NetworkInstanceName**, and **RegionNo**.\
	// ********************
	VpcSrcInfo *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo `json:"VpcSrcInfo,omitempty" xml:"VpcSrcInfo,omitempty" type:"Struct"`
	// The risk level of the intrusion event. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	VulLevel *int32 `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetAttackApp(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.AttackApp = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetAttackType(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDescription(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDirection(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDstIP(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventCount(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventCount = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventId(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventName(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetFirstEventTime(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.FirstEventTime = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetIPLocationInfo(v *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.IPLocationInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetLastEventTime(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.LastEventTime = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetResourcePrivateIPList(v []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) *DescribeRiskEventGroupResponseBodyDataList {
	s.ResourcePrivateIPList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetResourceType(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleId(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleResult(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleResult = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleSource(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleSource = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIP(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIPTag(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIPTag = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcPrivateIPList(v []*string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcPrivateIPList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetTag(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Tag = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVpcDstInfo(v *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.VpcDstInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVpcSrcInfo(v *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.VpcSrcInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVulLevel(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.VulLevel = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListIPLocationInfo struct {
	// The ID of the city to which the IP address belongs.
	CityId *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	// The name of the city to which the IP address belongs.
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The ID of the country to which the IP address belongs.
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The name of the country to which the IP address belongs.
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCityId(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CityId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCityName(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CityName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCountryId(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CountryId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCountryName(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CountryName = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList struct {
	// The ID of the region to which the private IP address belongs.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the instance that uses the private IP address.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The name of the instance that uses the private IP address.
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// The private IP address.
	ResourcePrivateIP *string `json:"ResourcePrivateIP,omitempty" xml:"ResourcePrivateIP,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourceInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourceInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourcePrivateIP(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourcePrivateIP = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListVpcDstInfo struct {
	// The ID of the ECS instance.
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the ECS instance.
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The ID of the VPC.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the VPC.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The ID of the region in which the destination VPC resides.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetEcsInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetEcsInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetNetworkInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetNetworkInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.RegionNo = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo struct {
	// The ID of the ECS instance.
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the ECS instance.
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The ID of the VPC.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the VPC.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The ID of the region in which the source VPC resides.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetEcsInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetEcsInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetNetworkInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetNetworkInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.RegionNo = &v
	return s
}

type DescribeRiskEventGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRiskEventGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskEventGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponse) SetHeaders(v map[string]*string) *DescribeRiskEventGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventGroupResponse) SetStatusCode(v int32) *DescribeRiskEventGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventGroupResponse) SetBody(v *DescribeRiskEventGroupResponseBody) *DescribeRiskEventGroupResponse {
	s.Body = v
	return s
}

type DescribeUserAssetIPTrafficInfoRequest struct {
	// The IP address of the asset.
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The point in time to query. The value is a UNIX timestamp. Unit: seconds.
	TrafficTime *string `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetAssetIP(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.AssetIP = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetLang(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetTrafficTime(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.TrafficTime = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponseBody struct {
	// The end of the time range that is queried. The value is a UNIX timestamp. Unit: seconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The network throughput, which indicates the inbound traffic rate. Unit: bit/s.
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The network throughput, which indicates the inbound packet rate. Unit: packets per second (pps).
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The number of new connections.
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// The network throughput, which indicates the outbound traffic rate. Unit: bit/s.
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The network throughput, which indicates the outbound packet rate. Unit: pps.
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The beginning of the time range that is queried. The value is a UNIX timestamp. Unit: seconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetEndTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetNewConn(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.NewConn = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetRequestId(v string) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetSessionCount(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.SessionCount = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetStartTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.StartTime = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserAssetIPTrafficInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserAssetIPTrafficInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetHeaders(v map[string]*string) *DescribeUserAssetIPTrafficInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetStatusCode(v int32) *DescribeUserAssetIPTrafficInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetBody(v *DescribeUserAssetIPTrafficInfoResponseBody) *DescribeUserAssetIPTrafficInfoResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallAclGroupListRequest struct {
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether VPC firewalls are configured. Valid values:
	//
	// - **notconfigured**: VPC firewalls are not configured.
	// - **configured**: VPC firewalls are configured.
	// - If this parameter is left empty, all policy groups of access control policies are queried.
	FirewallConfigureStatus *string `json:"FirewallConfigureStatus,omitempty" xml:"FirewallConfigureStatus,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetCurrentPage(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetFirewallConfigureStatus(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.FirewallConfigureStatus = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetLang(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetPageSize(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.PageSize = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponseBody struct {
	// An array that consists of the information about the policy group.
	AclGroupList []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList `json:"AclGroupList,omitempty" xml:"AclGroupList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the policy groups that are returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetAclGroupList(v []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) *DescribeVpcFirewallAclGroupListResponseBody {
	s.AclGroupList = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetRequestId(v string) *DescribeVpcFirewallAclGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallAclGroupListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponseBodyAclGroupList struct {
	// The ID of the policy group.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall instance.
	//
	// Example: vfw-a42bbb7b887148c9****
	AclGroupId *string `json:"AclGroupId,omitempty" xml:"AclGroupId,omitempty"`
	// The name of the policy group. Valid values:
	//
	// - If the VPC firewall is used to protect a CEN instance, the value of this parameter is the name of the CEN instance.
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the name of the VPC firewall instance.
	AclGroupName *string `json:"AclGroupName,omitempty" xml:"AclGroupName,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclGroupId(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclGroupId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclGroupName(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclGroupName = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetMemberUid(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.MemberUid = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallAclGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallAclGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallAclGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetStatusCode(v int32) *DescribeVpcFirewallAclGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetBody(v *DescribeVpcFirewallAclGroupListResponseBody) *DescribeVpcFirewallAclGroupListResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallCenDetailRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the VPC for which the VPC firewall is created.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailRequest) SetLang(v string) *DescribeVpcFirewallCenDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBody struct {
	// The connection type of the VPC firewall. The value is fixed as **cen**, which indicates CEN instances.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: enabled
	// *   **closed**: disabled
	// *   **notconfigured**: not configured
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The details about the VPC.
	LocalVpc *DescribeVpcFirewallCenDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetConnectType(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetLocalVpc(v *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) *DescribeVpcFirewallCenDetailResponseBody {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetVpcFirewallName(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpc struct {
	// The ID of the connection between two network instances.
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The name of the connection between two network instances.
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The CIDR blocks that are protected by the VPC firewall.
	DefendCidrList []*string `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	// The Elastic Network Interfaces (ENIs).
	EniList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList `json:"EniList,omitempty" xml:"EniList,omitempty" type:"Repeated"`
	// The ID of the vSwitch. The value of this parameter is returned only when the RouteMode parameter is set to manual.
	ManualVSwitchId *string `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	// The ID of the VPC for which the VPC firewall is created.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the network instance.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The type of the network instance. The value is fixed as **VPC**.
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// The UID of the Alibaba Cloud account to which the VPC belongs.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the VPC resides.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode. Valid values:
	//
	// *   auto
	// *   manual
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Indicates whether the routing mode can be set to manual. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	SupportManualMode *string `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	// The instance ID of the CEN transit router.
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The edition of the CEN transit router. Valid values:
	//
	// *   **Basic**: Basic Edition
	// *   **Enterprise**: Enterprise Edition
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// The CIDR blocks of the VPC.
	VpcCidrTableList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetAttachmentId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.AttachmentId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetAttachmentName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.AttachmentName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetDefendCidrList(v []*string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.DefendCidrList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetEniList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.EniList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetManualVSwitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.ManualVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceType(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetOwnerId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetRouteMode(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetSupportManualMode(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.SupportManualMode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetTransitRouterId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetTransitRouterType(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList struct {
	// The ID of the ENI that belongs to the VPC.
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the ENI that belongs to the VPC.
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniPrivateIpAddress = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList struct {
	// The route entries of the VPC.
	RouteEntryList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallCenDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallCenDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) SetBody(v *DescribeVpcFirewallCenDetailResponseBody) *DescribeVpcFirewallCenDetailResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallCenListRequest struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	// *   **configured**: The VPC firewall is created but is not enabled.
	//
	// >  If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is manged by your Alibaba Cloud account. The member is also an Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the network instance.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPC.
	//
	// >  For more information about the regions, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **auto**: automatic mode
	// *   **manual**: manual mode
	//
	// >  If you do not specify this parameter, VPC firewalls in all routing modes are queried.
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The type of the transit router. Valid values:
	//
	// *   **Basic**: Basic Edition transit router
	// *   **Enterprise**: Enterprise Edition transit router
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListRequest) SetCenId(v string) *DescribeVpcFirewallCenListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetCurrentPage(v string) *DescribeVpcFirewallCenListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetLang(v string) *DescribeVpcFirewallCenListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetMemberUid(v string) *DescribeVpcFirewallCenListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetOwnerId(v string) *DescribeVpcFirewallCenListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetPageSize(v string) *DescribeVpcFirewallCenListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetRegionNo(v string) *DescribeVpcFirewallCenListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetRouteMode(v string) *DescribeVpcFirewallCenListRequest {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetTransitRouterType(v string) *DescribeVpcFirewallCenListRequest {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPC firewalls.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details about the VPC firewall.
	VpcFirewalls []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls `json:"VpcFirewalls,omitempty" xml:"VpcFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallCenListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallCenListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) SetVpcFirewalls(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls) *DescribeVpcFirewallCenListResponseBody {
	s.VpcFirewalls = v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewalls struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The name of the CEN instance.
	CenName *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as cen, which indicates a CEN instance.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The information about the intrusion prevention system (IPS) configuration.
	IpsConfig *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The details about the VPC.
	LocalVpc *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is manged by your Alibaba Cloud account. The member is also an Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// Indicates whether the VPC firewall can be automatically enabled to protect VPC traffic based on route learning. Valid values:
	//
	// - **passed**: The VPC firewall can be automatically enabled.
	// - **failed**: The VPC firewall cannot be automatically enabled.
	// - **unknown**: The VPC firewall is in an unknown state.
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// Indicates whether you can create a VPC firewall in a region. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation that creates the VPC firewall. Valid values:
	//
	// *   **Unauthorized**: Cloud Firewall is not authorized to access the VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **RegionDisable**: .VPC Firewall is not supported in the region of the VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **OpsDisable**: You are not allowed to create the VPC firewall.
	// *   **VbrNotSupport**: The VPC firewall cannot be created for a VBR that is attached to the CEN instance.
	// *   Empty string: You can create a VPC firewall for the network instance.
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewalls) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetCenId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetCenName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.CenName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetConnectType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetIpsConfig(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetLocalVpc(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetMemberUid(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetPrecheckStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetRegionStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetResultCode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// - **1**: yes
	// - **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// - **1**: yes
	// - **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// - **1**: block mode
	// - **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetBasicRules(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetRunMode(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc struct {
	// Indicates whether the VPC is granted the required permissions. The value is fixed as **authorized**, which indicates that the VPC is granted the required permissions.
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The CIDR blocks that are protected by the VPC firewall.
	DefendCidrList []*string `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	// The ID of the specified vSwitch when the routing mode is manual.
	ManualVSwitchId *string `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	// The ID of the network instance.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the network instance.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The type of the network instance. Valid values:
	//
	// *   **VPC**
	// *   **VBR**
	// *   **CCN**
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **auto**: automatic mode
	// *   **manual**: manual mode
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Indicates whether the manual routing mode is supported. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	SupportManualMode *string `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	// The type of the CEN transit router. Valid values:
	//
	// *   **Basic**: Basic Edition transit router
	// *   **Enterprise**: Enterprise Edition transit router
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// The CIDR block of the VPC.
	VpcCidrTableList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetDefendCidrList(v []*string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.DefendCidrList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetManualVSwitchId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.ManualVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetOwnerId(v int64) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetRouteMode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetSupportManualMode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.SupportManualMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetTransitRouterType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList struct {
	// The route entries for the VPC.
	RouteEntryList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The route table ID of the VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID for the next hop of the VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallCenListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallCenListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallCenListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) SetBody(v *DescribeVpcFirewallCenListResponseBody) *DescribeVpcFirewallCenListResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: blocks the traffic.
	// *   **log**: monitors the traffic.
	//
	// >  If you do not specify this parameter, access control policies are queried based on all actions.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The number of the page to return.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy match is supported.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported.
	//
	// >  The value of this parameter can be a CIDR block, a domain name, or an address book name.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The language of the content within the request and response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	// *   **ANY**: all protocol types
	//
	// >  If you do not specify this parameter, access control policies are queried based on all protocol types.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Indicates whether the access control policy is enabled. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// *   **true**: The access control policy is enabled.
	// *   **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported.
	//
	// >  The value of this parameter can be a CIDR block or an address book name.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects the traffic between two VPCs that are connected by using a CEN instance, the value of this parameter must be the ID of the CEN instance.
	// *   If the VPC firewall protects the traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the IDs.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclAction(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetCurrentPage(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDescription(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDestination(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetLang(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetPageSize(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetProto(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetRelease(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetSource(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBody struct {
	// The details of the access control policies.
	Policys []*DescribeVpcFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the access control policies that are returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetPolicys(v []*DescribeVpcFirewallControlPolicyResponseBodyPolicys) *DescribeVpcFirewallControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetTotalCount(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: blocks the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application ID in the access control policy.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type in the access control policy. Valid values:
	//
	// *   **HTTP**
	// *   **HTTPS**
	// *   **MySQL**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **RDP**
	// *   **VNC**
	// *   **SSH**
	// *   **Redis**
	// *   **MQTT**
	// *   **MongoDB**
	// *   **Memcache**
	// *   **SSL**
	// *   **ANY**: all application types
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book of the access control policy.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Valid values:
	//
	// *   If **DestinationType** is set to `net`, the value of this parameter is a CIDR block.
	// *   If **DestinationType** is set to `domain`, the value of this parameter is a domain name.
	// *   If **DestinationType** is set to `group`, the value of this parameter is the name of an address book name.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book of the access control policy.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more CIDR blocks
	// *   **domain**: an address book that includes one or more domain names
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The number of hits for the access control policy.
	HitTimes *int32 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	// *   **ANY**: all protocol types
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Indicates whether the access control policy is enabled. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// *   **true**: The access control policy is enabled.
	// *   **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of this parameter is a CIDR block.
	// *   If **SourceType** is set to `group`, the value of this parameter is an address book name.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book of the access control policy.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address in the access control policy. The value is fixed as **ip**. The value indicates an address book that includes one or more CIDR blocks.
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetHitTimes(v int32) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *DescribeVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetBody(v *DescribeVpcFirewallControlPolicyResponseBody) *DescribeVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallDefaultIPSConfigRequest struct {
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the IDs of CEN instances.
	// *   If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall. You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) SetMemberUid(v string) *DescribeVpcFirewallDefaultIPSConfigRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallDefaultIPSConfigRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallDefaultIPSConfigResponseBody struct {
	// Indicates whether basic policies are enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The mode of the intrusion prevention system (IPS). Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetBasicRules(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetEnableAllPatch(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRequestId(v string) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRunMode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RunMode = &v
	return s
}

type DescribeVpcFirewallDefaultIPSConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallDefaultIPSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetStatusCode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetBody(v *DescribeVpcFirewallDefaultIPSConfigResponseBody) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallDetailRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the local VPC.
	LocalVpcId *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	// The ID of the peer VPC.
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailRequest) SetLang(v string) *DescribeVpcFirewallDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetLocalVpcId(v string) *DescribeVpcFirewallDetailRequest {
	s.LocalVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetPeerVpcId(v string) *DescribeVpcFirewallDetailRequest {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallDetailRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBody struct {
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as **expressconnect**, which indicates Express Connect circuits.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not configured.
	// *   **configured**: The VPC firewall is configured.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The details about the local VPC.
	LocalVpc *DescribeVpcFirewallDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The details about the peer VPC.
	PeerVpc *DescribeVpcFirewallDetailResponseBodyPeerVpc `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBody) SetBandwidth(v int32) *DescribeVpcFirewallDetailResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetConnectType(v string) *DescribeVpcFirewallDetailResponseBody {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallDetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetLocalVpc(v *DescribeVpcFirewallDetailResponseBodyLocalVpc) *DescribeVpcFirewallDetailResponseBody {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetMemberUid(v string) *DescribeVpcFirewallDetailResponseBody {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetPeerVpc(v *DescribeVpcFirewallDetailResponseBodyPeerVpc) *DescribeVpcFirewallDetailResponseBody {
	s.PeerVpc = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetVpcFirewallId(v string) *DescribeVpcFirewallDetailResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetVpcFirewallName(v string) *DescribeVpcFirewallDetailResponseBody {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyLocalVpc struct {
	// The ID of the ENI for the local VPC.
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the elastic network interface (ENI) for the local VPC.
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The region ID of the local VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The router interface ID of the local VPC.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The CIDR blocks of the local VPC.
	VpcCidrTableList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the local VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the local VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetEniId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetRouterInterfaceId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcName(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList struct {
	// The route entries of the local VPC.
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the local VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the local VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the local VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyPeerVpc struct {
	// The ID of the ENI for the peer VPC.
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the ENI for the peer VPC.
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The region ID of the peer VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The router interface ID of the peer VPC.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The CIDR blocks of the peer VPC.
	VpcCidrTableList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the peer VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the peer VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetEniId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetRegionNo(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetRouterInterfaceId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcName(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList struct {
	// The route entries of the peer VPC.
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the peer VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the peer VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the peer VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) SetBody(v *DescribeVpcFirewallDetailResponseBody) *DescribeVpcFirewallDetailResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallListRequest struct {
	// The sub-type of the connection. Valid values:
	//
	// *   **vpc2vpc**: Express Connect connection
	// *   **vpcpeer**: peer connection
	ConnectSubType *string `json:"ConnectSubType,omitempty" xml:"ConnectSubType,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page **1**. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	// *   **configured**: The VPC firewall is created.
	//
	// >  If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **10**.**** Maximum value: **50**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UID of the Alibaba Cloud account to which the peer VPC belongs.
	PeerUid *string `json:"PeerUid,omitempty" xml:"PeerUid,omitempty"`
	// The region ID of the VPC.
	//
	// >  For more information about the regions, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListRequest) SetConnectSubType(v string) *DescribeVpcFirewallListRequest {
	s.ConnectSubType = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetCurrentPage(v string) *DescribeVpcFirewallListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetLang(v string) *DescribeVpcFirewallListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetMemberUid(v string) *DescribeVpcFirewallListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetPageSize(v string) *DescribeVpcFirewallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetPeerUid(v string) *DescribeVpcFirewallListRequest {
	s.PeerUid = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetRegionNo(v string) *DescribeVpcFirewallListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallListRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcId(v string) *DescribeVpcFirewallListRequest {
	s.VpcId = &v
	return s
}

type DescribeVpcFirewallListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPC firewalls.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details about the VPC firewalls.
	VpcFirewalls []*DescribeVpcFirewallListResponseBodyVpcFirewalls `json:"VpcFirewalls,omitempty" xml:"VpcFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBody) SetRequestId(v string) *DescribeVpcFirewallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) SetVpcFirewalls(v []*DescribeVpcFirewallListResponseBodyVpcFirewalls) *DescribeVpcFirewallListResponseBody {
	s.VpcFirewalls = v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewalls struct {
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The sub-type of the connection. Valid values:
	//
	// *   **vpc2vpc**: Express Connect connection
	// *   **vpcpeer**: peer connection
	ConnectSubType *string `json:"ConnectSubType,omitempty" xml:"ConnectSubType,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as **expressconnect**, which indicates an Express Connect connection.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The information about the intrusion prevention system (IPS) configuration.
	IpsConfig *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The details about the local VPC.
	LocalVpc *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The details about the peer VPC.
	PeerVpc *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	// Indicates whether you can create a VPC firewall in a specified region. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation that creates the VPC firewall. Valid values:
	//
	// *   **Unauthorized**: Cloud Firewall is not authorized to access a VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **RegionDisable**: VPC Firewall is not supported in the region of a VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **Empty string**: You can create a VPC firewall for the network instance.
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewalls) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetBandwidth(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetConnectSubType(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ConnectSubType = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetConnectType(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetIpsConfig(v *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetLocalVpc(v *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetMemberUid(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetPeerVpc(v *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.PeerVpc = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetRegionStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetResultCode(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetBasicRules(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetRunMode(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc struct {
	// Indicates whether Cloud Firewall is authorized to access the local VPC. The value is fixed as authorized, which indicates that Cloud Firewall is authorized to access the local VPC.
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The UID of the Alibaba Cloud account to which the local VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the local VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The CIDR blocks of the local VPC.
	VpcCidrTableList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the local VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the local VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetOwnerId(v int64) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList struct {
	// The route entries of the local VPC.
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the local VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the local VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the local VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc struct {
	// Indicates whether Cloud Firewall is authorized to access the peer VPC. The value is fixed as **authorized**, which indicates that Cloud Firewall is authorized to access the peer VPC.
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The UID of the Alibaba Cloud account to which the peer VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the peer VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The CIDR blocks of the peer VPC.
	VpcCidrTableList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the peer VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the peer VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetOwnerId(v int64) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetRegionNo(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList struct {
	// The route entries of the peer VPC.
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the peer VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the peer VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the peer VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallListResponse) SetStatusCode(v int32) *DescribeVpcFirewallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallListResponse) SetBody(v *DescribeVpcFirewallListResponseBody) *DescribeVpcFirewallListResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedRequest struct {
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the policy group to which the access control policy belongs. You can call the DescribeVpcFirewallAclGroupList operation to query the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall instance.
	//
	// Example: vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetLang(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedResponseBody struct {
	// The lowest priority for the access control policy.
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority for the access control policy.
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetEnd(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetRequestId(v string) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetStart(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetStatusCode(v int32) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetBody(v *DescribeVpcFirewallPolicyPriorUsedResponseBody) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Body = v
	return s
}

type DescribeVulnerabilityProtectedListRequest struct {
	// The attack type of the intrusion event. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leakage
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: webshell
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	//
	// >  If you do not specify this parameter, the intrusion events of all attack types are queried.
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The edition of Cloud Firewall.
	BuyVersion *int64 `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// The number of the page to return. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The order in which you want to sort the queried information. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The dimension based on which you want to sort the queried information. Set the value to **attackCnt**, which indicates the number of attacks.
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the user. Set the value to **buy**, which indicates users of a paid edition of Cloud Firewall.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID of the vulnerability.
	VulnCveName *string `json:"VulnCveName,omitempty" xml:"VulnCveName,omitempty"`
	// The risk level of the vulnerability. Valid values:
	//
	// *   **high**
	// *   **medium**
	// *   **low**
	VulnLevel *string `json:"VulnLevel,omitempty" xml:"VulnLevel,omitempty"`
	// The number of assets that are affected by the vulnerability.
	VulnResource *string `json:"VulnResource,omitempty" xml:"VulnResource,omitempty"`
	// The status of vulnerability protection. Valid values:
	//
	// *   **partProtected**: partially protected
	// *   **protected**: protected
	// *   **unProtected**: unprotected
	VulnStatus *string `json:"VulnStatus,omitempty" xml:"VulnStatus,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// *   **App**: application vulnerability
	// *   **emg**: urgent vulnerability
	// *   **cms**: Web-CMS vulnerability
	VulnType *string `json:"VulnType,omitempty" xml:"VulnType,omitempty"`
}

func (s DescribeVulnerabilityProtectedListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListRequest) SetAttackType(v string) *DescribeVulnerabilityProtectedListRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetBuyVersion(v int64) *DescribeVulnerabilityProtectedListRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetCurrentPage(v string) *DescribeVulnerabilityProtectedListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetEndTime(v string) *DescribeVulnerabilityProtectedListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetLang(v string) *DescribeVulnerabilityProtectedListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetMemberUid(v string) *DescribeVulnerabilityProtectedListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetOrder(v string) *DescribeVulnerabilityProtectedListRequest {
	s.Order = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetPageSize(v string) *DescribeVulnerabilityProtectedListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetSortKey(v string) *DescribeVulnerabilityProtectedListRequest {
	s.SortKey = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetSourceIp(v string) *DescribeVulnerabilityProtectedListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetStartTime(v string) *DescribeVulnerabilityProtectedListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetUserType(v string) *DescribeVulnerabilityProtectedListRequest {
	s.UserType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnCveName(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnCveName = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnLevel(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnLevel = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnResource(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnResource = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnStatus(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnStatus = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnType(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnType = &v
	return s
}

type DescribeVulnerabilityProtectedListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities that are detected by Cloud Firewall.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// An array that consists of the information about the vulnerabilities.
	VulnList []*DescribeVulnerabilityProtectedListResponseBodyVulnList `json:"VulnList,omitempty" xml:"VulnList,omitempty" type:"Repeated"`
	// The number of assets on which no vulnerabilities are detected.
	ZeroResourceCount *int32 `json:"ZeroResourceCount,omitempty" xml:"ZeroResourceCount,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetRequestId(v string) *DescribeVulnerabilityProtectedListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetTotalCount(v int32) *DescribeVulnerabilityProtectedListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetVulnList(v []*DescribeVulnerabilityProtectedListResponseBodyVulnList) *DescribeVulnerabilityProtectedListResponseBody {
	s.VulnList = v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetZeroResourceCount(v int32) *DescribeVulnerabilityProtectedListResponseBody {
	s.ZeroResourceCount = &v
	return s
}

type DescribeVulnerabilityProtectedListResponseBodyVulnList struct {
	// The number of vulnerabilities.
	AttackCnt *int32 `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// The attack type of the intrusion event. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leakage
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: webshell
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	//
	// >  If no attack type is specified, the intrusion events of all attack types are queried.
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The IDs of associated basic protection policies.
	BasicRuleIds *string `json:"BasicRuleIds,omitempty" xml:"BasicRuleIds,omitempty"`
	// The CVE IDs.
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The time when the first attack was launched.
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether you need to pay special attention to the vulnerability. Valid values:
	//
	// *   **0**: no
	// *   **1**: yes
	HighlightTag *int32 `json:"HighlightTag,omitempty" xml:"HighlightTag,omitempty"`
	// The time when the last attack was launched.
	LastTime  *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The status of basic protection. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	//
	// >  If the value of this parameter is true, you must configure the intrusion prevention feature when you enable protection.
	NeedOpenBasicRule *bool `json:"NeedOpenBasicRule,omitempty" xml:"NeedOpenBasicRule,omitempty"`
	// The UUIDs of the basic protection policies for which you want to set the Current Action parameter to Block.
	NeedOpenBasicRuleUuids *string `json:"NeedOpenBasicRuleUuids,omitempty" xml:"NeedOpenBasicRuleUuids,omitempty"`
	// Indicates whether the intrusion prevention feature needs to be configured when you enable protection. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	NeedOpenRunMode *bool `json:"NeedOpenRunMode,omitempty" xml:"NeedOpenRunMode,omitempty"`
	// The status of virtual patching. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	//
	// >  If the value of this parameter is true, you must configure the intrusion prevention feature when you enable protection.
	NeedOpenVirtualPatche *bool `json:"NeedOpenVirtualPatche,omitempty" xml:"NeedOpenVirtualPatche,omitempty"`
	// The UUIDs of the virtual patching policies for which you want to set the Current Action parameter to Block.
	NeedOpenVirtualPatcheUuids *string `json:"NeedOpenVirtualPatcheUuids,omitempty" xml:"NeedOpenVirtualPatcheUuids,omitempty"`
	// The Rule Group that you want to specify. Valid values:
	//
	// *   **1**: Loose (default)
	// *   **2**: Medium
	// *   **3**: Strict
	NeedRuleClass *int32 `json:"NeedRuleClass,omitempty" xml:"NeedRuleClass,omitempty"`
	// The number of assets on which vulnerabilities are detected.
	ResourceCnt *int32 `json:"ResourceCnt,omitempty" xml:"ResourceCnt,omitempty"`
	// An array consisting of the assets on which vulnerabilities are detected.
	ResourceList []*DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The IDs of associated virtual patching policies.
	VirtualPatcheIds *string `json:"VirtualPatcheIds,omitempty" xml:"VirtualPatcheIds,omitempty"`
	// The code of the vulnerability.
	VulnKey *string `json:"VulnKey,omitempty" xml:"VulnKey,omitempty"`
	// The risk level of the vulnerability. Valid values:
	//
	// *   **high**
	// *   **medium**
	// *   **low**
	VulnLevel *string `json:"VulnLevel,omitempty" xml:"VulnLevel,omitempty"`
	// The name of the vulnerability.
	VulnName *string `json:"VulnName,omitempty" xml:"VulnName,omitempty"`
	// The status of vulnerability protection. Valid values:
	//
	// *   **partProtected**: partially protected
	// *   **protected**: protected
	// *   **unProtected**: unprotected
	VulnStatus *string `json:"VulnStatus,omitempty" xml:"VulnStatus,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// *   **cve**: Windows vulnerability
	// *   **cms**: Web-CMS vulnerability
	// *   **App**: application vulnerability
	VulnType *string `json:"VulnType,omitempty" xml:"VulnType,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnList) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetAttackCnt(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.AttackCnt = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetAttackType(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.AttackType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetBasicRuleIds(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.BasicRuleIds = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetCveId(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.CveId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetFirstTime(v int64) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.FirstTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetHighlightTag(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.HighlightTag = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetLastTime(v int64) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.LastTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetMemberUid(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.MemberUid = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenBasicRule(v bool) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenBasicRule = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenBasicRuleUuids(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenBasicRuleUuids = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenRunMode(v bool) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenRunMode = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenVirtualPatche(v bool) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenVirtualPatche = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenVirtualPatcheUuids(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenVirtualPatcheUuids = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedRuleClass(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedRuleClass = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetResourceCnt(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.ResourceCnt = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetResourceList(v []*DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.ResourceList = v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVirtualPatcheIds(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VirtualPatcheIds = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnKey(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnKey = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnLevel(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnLevel = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnName(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnName = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnStatus(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnStatus = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnType(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnType = &v
	return s
}

type DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList struct {
	// The elastic IP address (EIP) that is associated with the instance.
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The public IP address of the instance.
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the instance.
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the region in which Cloud Firewall is supported.
	//
	// >  For more information about the regions, see [Supported regions](~~195657~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the asset. Valid values:
	//
	// *   **SLB**
	// *   **EIP**
	// *   **ECS**
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of vulnerability protection. Valid values:
	//
	// *   **partProtected**: partially protected
	// *   **protected**: protected
	// *   **unProtected**: unprotected
	VulnStatus *string `json:"VulnStatus,omitempty" xml:"VulnStatus,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetEip(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.Eip = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetInternetIp(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.InternetIp = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetIntranetIp(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetRegionId(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.RegionId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetResourceId(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.ResourceId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetResourceName(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.ResourceName = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetResourceType(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.ResourceType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetVulnStatus(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.VulnStatus = &v
	return s
}

type DescribeVulnerabilityProtectedListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVulnerabilityProtectedListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulnerabilityProtectedListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponse) SetHeaders(v map[string]*string) *DescribeVulnerabilityProtectedListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponse) SetStatusCode(v int32) *DescribeVulnerabilityProtectedListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponse) SetBody(v *DescribeVulnerabilityProtectedListResponseBody) *DescribeVulnerabilityProtectedListResponse {
	s.Body = v
	return s
}

type ModifyAddressBookRequest struct {
	// The addresses in the address book. Separate multiple addresses with commas (,). If you set GroupType to **ip**, **port**, or **domain**, you must specify this parameter.
	//
	// *   If you set GroupType to **ip**, you must specify IP addresses for the address book. Example: 1.2.XX.XX/32,1.2.XX.XX/24.
	// *   If you set GroupType to **port**, you must specify port numbers or port ranges for the address book. Example: 80/80,100/200.
	// *   If you set GroupType to **domain**, you must specify domain names for the address book. Example: demo1.aliyun.com,demo2.aliyun.com.
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of Elastic Compute Service (ECS) instances to the address book if the instances match the specified tags. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	AutoAddTagEcs *string `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the address book.
	//
	// >  To modify the address book, you must provide the ID of the address book. You can call the [DescribeAddressBook](~~138869~~) operation to query the ID.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ECS tags that you want to match.
	TagList []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among ECS tags. Valid values:
	//
	// *   **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	// *   **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the address book.
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s ModifyAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookRequest) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequest) SetAddressList(v string) *ModifyAddressBookRequest {
	s.AddressList = &v
	return s
}

func (s *ModifyAddressBookRequest) SetAutoAddTagEcs(v string) *ModifyAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *ModifyAddressBookRequest) SetDescription(v string) *ModifyAddressBookRequest {
	s.Description = &v
	return s
}

func (s *ModifyAddressBookRequest) SetGroupName(v string) *ModifyAddressBookRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyAddressBookRequest) SetGroupUuid(v string) *ModifyAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *ModifyAddressBookRequest) SetLang(v string) *ModifyAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAddressBookRequest) SetSourceIp(v string) *ModifyAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAddressBookRequest) SetTagList(v []*ModifyAddressBookRequestTagList) *ModifyAddressBookRequest {
	s.TagList = v
	return s
}

func (s *ModifyAddressBookRequest) SetTagRelation(v string) *ModifyAddressBookRequest {
	s.TagRelation = &v
	return s
}

type ModifyAddressBookRequestTagList struct {
	// The key of ECS tag N that you want to match.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of ECS tag N that you want to match.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ModifyAddressBookRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestTagList) SetTagKey(v string) *ModifyAddressBookRequestTagList {
	s.TagKey = &v
	return s
}

func (s *ModifyAddressBookRequestTagList) SetTagValue(v string) *ModifyAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

type ModifyAddressBookResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookResponseBody) SetRequestId(v string) *ModifyAddressBookResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAddressBookResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookResponse) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookResponse) SetHeaders(v map[string]*string) *ModifyAddressBookResponse {
	s.Headers = v
	return s
}

func (s *ModifyAddressBookResponse) SetStatusCode(v int32) *ModifyAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAddressBookResponse) SetBody(v *ModifyAddressBookResponseBody) *ModifyAddressBookResponse {
	s.Body = v
	return s
}

type ModifyControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The ID of the access control policy.
	//
	// >  If you want to modify the configurations of an access control policy, you must provide the ID of the policy. You can call the [DescribeControlPolicy](~~138866~~) operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The type of the application that the access control policy supports. Valid values:
	//
	// *   **ANY**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **MySQL**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **RDP**
	// *   **VNC**
	// *   **SSH**
	// *   **Redis**
	// *   **MQTT**
	// *   **MongoDB**
	// *   **Memcache**
	// *   **SSL**
	//
	// >  The value **ANY** indicates all types of applications.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application names. You can specify multiple application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// *   If **DestinationType** is set to net, the value of **Destination** is a CIDR block. Example: 1.2.XX.XX/24
	// *   If **DestinationType** is set to group, the value of **Destination** is an address book. Example: db_group
	// *   If **DestinationType** is set to domain, the value of **Destination** is a domain name. Example: \*.aliyuncs.com
	// *   If **DestinationType** is set to location, the value of **Destination** is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: \["BJ11", "ZB"]
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: destination CIDR block
	// *   **group**: destination address book
	// *   **domain**: destination domain name
	// *   **location**: destination location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	//
	// >  The value **ANY** indicates all types of protocols.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// *   true: enabled
	// *   false: disabled
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy.
	//
	// *   If **SourceType** is set to net, the value of **Source** is a CIDR block. Example: 1.2.XX.XX/24
	// *   If **SourceType** is set to group, the value of **Source** is an address book. Example: db_group
	// *   If **SourceType** is set to location, the value of **Source** is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: \["BJ11", "ZB"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: source CIDR block
	// *   **group**: source address book
	// *   **location**: source location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ModifyControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyRequest) SetAclAction(v string) *ModifyControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetAclUuid(v string) *ModifyControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationName(v string) *ModifyControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyControlPolicyRequest) SetDescription(v string) *ModifyControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPort(v string) *ModifyControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortGroup(v string) *ModifyControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortType(v string) *ModifyControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestination(v string) *ModifyControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestinationType(v string) *ModifyControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDirection(v string) *ModifyControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetLang(v string) *ModifyControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetProto(v string) *ModifyControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRelease(v string) *ModifyControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSource(v string) *ModifyControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSourceType(v string) *ModifyControlPolicyRequest {
	s.SourceType = &v
	return s
}

type ModifyControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyResponseBody) SetRequestId(v string) *ModifyControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyResponse) SetStatusCode(v int32) *ModifyControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyResponse) SetBody(v *ModifyControlPolicyResponseBody) *ModifyControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyControlPolicyPositionRequest struct {
	// The direction of the traffic to which the IPv4 access control policy applies. Valid values:
	//
	// *   in: inbound traffic
	// *   out: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   zh: Chinese (default)
	// *   en: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the IPv4 access control policy.
	//
	// You must specify a numeric value for this parameter. The value 1 indicates the highest priority. A larger value indicates a lower priority.
	//
	// >  The value of this parameter must be within the priority range of existing IPv4 access control policies. Otherwise, an error occurs when you call this operation.
	//
	// We recommend that you first call the [DescribePolicyPriorUsed](~~138862~~) operation to query the priority range of existing IPv4 access control policies that apply to the traffic of the specified direction.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the IPv4 access control policy.
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionRequest) SetDirection(v string) *ModifyControlPolicyPositionRequest {
	s.Direction = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetLang(v string) *ModifyControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetNewOrder(v string) *ModifyControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetOldOrder(v string) *ModifyControlPolicyPositionRequest {
	s.OldOrder = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetSourceIp(v string) *ModifyControlPolicyPositionRequest {
	s.SourceIp = &v
	return s
}

type ModifyControlPolicyPositionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyControlPolicyPositionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyControlPolicyPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyPositionResponse) SetBody(v *ModifyControlPolicyPositionResponseBody) *ModifyControlPolicyPositionResponse {
	s.Body = v
	return s
}

type ModifyInstanceMemberAttributesRequest struct {
	// The members that to be modified.
	Members []*ModifyInstanceMemberAttributesRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s ModifyInstanceMemberAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesRequest) SetMembers(v []*ModifyInstanceMemberAttributesRequestMembers) *ModifyInstanceMemberAttributesRequest {
	s.Members = v
	return s
}

type ModifyInstanceMemberAttributesRequestMembers struct {
	// The remarks of the member in Cloud Firewall.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The UID of the member in Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s ModifyInstanceMemberAttributesRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesRequestMembers) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesRequestMembers) SetMemberDesc(v string) *ModifyInstanceMemberAttributesRequestMembers {
	s.MemberDesc = &v
	return s
}

func (s *ModifyInstanceMemberAttributesRequestMembers) SetMemberUid(v int64) *ModifyInstanceMemberAttributesRequestMembers {
	s.MemberUid = &v
	return s
}

type ModifyInstanceMemberAttributesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMemberAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesResponseBody) SetRequestId(v string) *ModifyInstanceMemberAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMemberAttributesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceMemberAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceMemberAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesResponse) SetHeaders(v map[string]*string) *ModifyInstanceMemberAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) SetStatusCode(v int32) *ModifyInstanceMemberAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) SetBody(v *ModifyInstanceMemberAttributesResponseBody) *ModifyInstanceMemberAttributesResponse {
	s.Body = v
	return s
}

type ModifyPolicyAdvancedConfigRequest struct {
	// Specifies whether to enable the strict mode for the access control policy. Valid values:
	//
	// *   **on**: enables the strict mode.
	// *   **off**: disables the strict mode.
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyPolicyAdvancedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigRequest) SetInternetSwitch(v string) *ModifyPolicyAdvancedConfigRequest {
	s.InternetSwitch = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetLang(v string) *ModifyPolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetSourceIp(v string) *ModifyPolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

type ModifyPolicyAdvancedConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyAdvancedConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigResponseBody) SetRequestId(v string) *ModifyPolicyAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyAdvancedConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPolicyAdvancedConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigResponse) SetHeaders(v map[string]*string) *ModifyPolicyAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) SetStatusCode(v int32) *ModifyPolicyAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) SetBody(v *ModifyPolicyAdvancedConfigResponseBody) *ModifyPolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallCenConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetLang(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetMemberUid(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

type ModifyVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *ModifyVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallCenConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *ModifyVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetBody(v *ModifyVpcFirewallCenConfigureResponseBody) *ModifyVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallCenSwitchStatusRequest struct {
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// *   **open**: yes
	// *   **close**: no
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetFirewallSwitch(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetLang(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetMemberUid(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallCenSwitchStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusResponseBody) SetRequestId(v string) *ModifyVpcFirewallCenSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallCenSwitchStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallCenSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallCenSwitchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetStatusCode(v int32) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetBody(v *ModifyVpcFirewallCenSwitchStatusResponseBody) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallConfigureRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The CIDR blocks of the local VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the local VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the local VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the local VPC.
	//
	// >  You can call the [DescribeVpcFirewallDetail](~~342892~~) operation to query the CIDR blocks of local VPCs for VPC firewalls.
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The CIDR blocks of the peer VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the peer VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the peer VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the peer VPC.
	//
	// >  You can call the [DescribeVpcFirewallDetail](~~342892~~) operation to query the CIDR blocks of peer VPCs for VPC firewalls.
	PeerVpcCidrTableList *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s ModifyVpcFirewallConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureRequest) SetLang(v string) *ModifyVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetLocalVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest {
	s.LocalVpcCidrTableList = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetMemberUid(v string) *ModifyVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetPeerVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest {
	s.PeerVpcCidrTableList = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallConfigureRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetVpcFirewallName(v string) *ModifyVpcFirewallConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

type ModifyVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureResponseBody) SetRequestId(v string) *ModifyVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) SetStatusCode(v int32) *ModifyVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) SetBody(v *ModifyVpcFirewallConfigureResponseBody) *ModifyVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// - **accept**: allows the traffic.
	// - **drop**: denies the traffic.
	// - **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The ID of the access control policy.
	//
	// If you want to modify the configurations of an access control policy, you must provide the ID of the policy. You can call the [DescribeVpcFirewallControlPolicy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/describevpcfirewallcontrolpolicy#doc-api-Cloudfw-DescribeVpcFirewallControlPolicy) operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The type of the application that the access control policy supports.
	//
	// Valid values:
	//
	// - FTP
	// - HTTP
	// - HTTPS
	// - MySQL
	// - SMTP
	// - SMTPS
	// - RDP
	// - VNC
	// - SSH
	// - Redis
	// - MQTT
	// - MongoDB
	// - Memcache
	// - SSL
	// - ANY: all types of applications
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy.
	//
	// - **port**: port
	// - **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// - If **DestinationType** is set to `net`, the value of Destination is a CIDR block.
	//
	// Example: 10.2.3.0/24
	// - If **DestinationType** is set to `group`, the value of Destination is an address book.
	//
	// Example: db_group
	// - If **DestinationType** is set to `domain`, the value of Destination is a domain name.
	//
	// Example: *.aliyuncs.com
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy.
	//
	// Valid values:
	//
	// - **net**: destination CIDR block
	// - **group**: destination address book
	// - **domain**: destination domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the protocol in the access control policy.
	//
	// Valid values:
	//
	// - ANY: all types of protocols
	// - TCP
	// - UDP
	// - ICMP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Indicates whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// - **true**: The access control policy is enabled.
	// - **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The source address in the access control policy.
	//
	// Valid values:
	//
	// - If **SourceType** is set to `net`, the value of Source is a CIDR block.
	//
	// Example: 10.2.4.0/24
	// - If **SourceType** is set to `group`, the value of Source is an address book.
	//
	// Example: db_group
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy.
	//
	// Valid values:
	//
	// - **net**: source CIDR block
	// - **group**: source address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the policy group to which the access control policy belongs. You can call the DescribeVpcFirewallAclGroupList operation to query the ID.
	//
	// - If the VPC firewall is used to protect a CEN instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	// Example: vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclAction(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclUuid(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetApplicationName(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDescription(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPort(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestination(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestinationType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetProto(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRelease(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSource(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSourceType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *ModifyVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetBody(v *ModifyVpcFirewallControlPolicyResponseBody) *ModifyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallControlPolicyPositionRequest struct {
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the access control policy.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the access control policy.
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// The ID of the policy group to which the access control policy belongs. You can call the DescribeVpcFirewallAclGroupList operation to query the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	// Example: vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetNewOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetOldOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.OldOrder = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallControlPolicyPositionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyVpcFirewallControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallControlPolicyPositionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallControlPolicyPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetBody(v *ModifyVpcFirewallControlPolicyPositionResponseBody) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallDefaultIPSConfigRequest struct {
	// Specifies whether to enable basic protection. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *string `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Specifies whether to enable virtual patching. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *string `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The mode of the intrusion prevention system (IPS). Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the IDs of CEN instances.
	// *   If the VPC firewall protects mutual access traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the ID of the VPC firewall. You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetBasicRules(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.BasicRules = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetEnableAllPatch(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.EnableAllPatch = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetLang(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetMemberUid(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetRunMode(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.RunMode = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetSourceIp(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallDefaultIPSConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponseBody) SetRequestId(v string) *ModifyVpcFirewallDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallDefaultIPSConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallDefaultIPSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetStatusCode(v int32) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetBody(v *ModifyVpcFirewallDefaultIPSConfigResponseBody) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallSwitchStatusRequest struct {
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// *   **open**: enables the VPC firewall.
	// *   **close**: disables the VPC firewall.
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetFirewallSwitch(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetLang(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetMemberUid(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallSwitchStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusResponseBody) SetRequestId(v string) *ModifyVpcFirewallSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallSwitchStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcFirewallSwitchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetStatusCode(v int32) *ModifyVpcFirewallSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetBody(v *ModifyVpcFirewallSwitchStatusResponseBody) *ModifyVpcFirewallSwitchStatusResponse {
	s.Body = v
	return s
}

type PutDisableAllFwSwitchRequest struct {
	// The instance ID of your Cloud Firewall.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutDisableAllFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchRequest) SetInstanceId(v string) *PutDisableAllFwSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetLang(v string) *PutDisableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetSourceIp(v string) *PutDisableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutDisableAllFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableAllFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchResponseBody) SetRequestId(v string) *PutDisableAllFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutDisableAllFwSwitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDisableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutDisableAllFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchResponse) SetHeaders(v map[string]*string) *PutDisableAllFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutDisableAllFwSwitchResponse) SetStatusCode(v int32) *PutDisableAllFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDisableAllFwSwitchResponse) SetBody(v *PutDisableAllFwSwitchResponseBody) *PutDisableAllFwSwitchResponse {
	s.Body = v
	return s
}

type PutDisableFwSwitchRequest struct {
	// The IP addresses.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The types of the assets.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutDisableFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchRequest) SetIpaddrList(v []*string) *PutDisableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetLang(v string) *PutDisableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetRegionList(v []*string) *PutDisableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetResourceTypeList(v []*string) *PutDisableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetSourceIp(v string) *PutDisableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutDisableFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponseBody) SetRequestId(v string) *PutDisableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutDisableFwSwitchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDisableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutDisableFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponse) SetHeaders(v map[string]*string) *PutDisableFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutDisableFwSwitchResponse) SetStatusCode(v int32) *PutDisableFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDisableFwSwitchResponse) SetBody(v *PutDisableFwSwitchResponseBody) *PutDisableFwSwitchResponse {
	s.Body = v
	return s
}

type PutEnableAllFwSwitchRequest struct {
	// The instance ID of your Cloud Firewall.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutEnableAllFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchRequest) SetInstanceId(v string) *PutEnableAllFwSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetLang(v string) *PutEnableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetSourceIp(v string) *PutEnableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutEnableAllFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableAllFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchResponseBody) SetRequestId(v string) *PutEnableAllFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutEnableAllFwSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutEnableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEnableAllFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchResponse) SetHeaders(v map[string]*string) *PutEnableAllFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutEnableAllFwSwitchResponse) SetStatusCode(v int32) *PutEnableAllFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnableAllFwSwitchResponse) SetBody(v *PutEnableAllFwSwitchResponseBody) *PutEnableAllFwSwitchResponse {
	s.Body = v
	return s
}

type PutEnableFwSwitchRequest struct {
	// The list of IP addresses.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, ResourceTypeList parameters.
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the content within the request and response.
	//
	// *   **zh**: Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of regions.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, ResourceTypeList parameters.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The list of asset types.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, ResourceTypeList parameters.
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutEnableFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchRequest) SetIpaddrList(v []*string) *PutEnableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetLang(v string) *PutEnableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableFwSwitchRequest) SetRegionList(v []*string) *PutEnableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetResourceTypeList(v []*string) *PutEnableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetSourceIp(v string) *PutEnableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutEnableFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponseBody) SetRequestId(v string) *PutEnableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutEnableFwSwitchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutEnableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEnableFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponse) SetHeaders(v map[string]*string) *PutEnableFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutEnableFwSwitchResponse) SetStatusCode(v int32) *PutEnableFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnableFwSwitchResponse) SetBody(v *PutEnableFwSwitchResponseBody) *PutEnableFwSwitchResponse {
	s.Body = v
	return s
}

type ResetVpcFirewallRuleHitCountRequest struct {
	// The ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetAclUuid(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.AclUuid = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetLang(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.Lang = &v
	return s
}

type ResetVpcFirewallRuleHitCountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountResponseBody) SetRequestId(v string) *ResetVpcFirewallRuleHitCountResponseBody {
	s.RequestId = &v
	return s
}

type ResetVpcFirewallRuleHitCountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetVpcFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetVpcFirewallRuleHitCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountResponse) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetHeaders(v map[string]*string) *ResetVpcFirewallRuleHitCountResponse {
	s.Headers = v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetStatusCode(v int32) *ResetVpcFirewallRuleHitCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetBody(v *ResetVpcFirewallRuleHitCountResponseBody) *ResetVpcFirewallRuleHitCountResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"ap-southeast-1": tea.String("cloudfw.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou":    tea.String("cloudfw.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudfw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the AddAddressBook operation to create an address book for access control. The address book can be an IP address book, an ECS tag-based address book, a port address book, or a domain address book.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddAddressBookResponse
 */
func (client *Client) AddAddressBookWithOptions(request *AddAddressBookRequest, runtime *util.RuntimeOptions) (_result *AddAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressList)) {
		query["AddressList"] = request.AddressList
	}

	if !tea.BoolValue(util.IsUnset(request.AutoAddTagEcs)) {
		query["AutoAddTagEcs"] = request.AutoAddTagEcs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TagList)) {
		query["TagList"] = request.TagList
	}

	if !tea.BoolValue(util.IsUnset(request.TagRelation)) {
		query["TagRelation"] = request.TagRelation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the AddAddressBook operation to create an address book for access control. The address book can be an IP address book, an ECS tag-based address book, a port address book, or a domain address book.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddAddressBookRequest
 * @return AddAddressBookResponse
 */
func (client *Client) AddAddressBook(request *AddAddressBookRequest) (_result *AddAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAddressBookResponse{}
	_body, _err := client.AddAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the AddControlPolicy operation to create an access control policy to allow, deny, or monitor traffic that passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddControlPolicyResponse
 */
func (client *Client) AddControlPolicyWithOptions(request *AddControlPolicyRequest, runtime *util.RuntimeOptions) (_result *AddControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the AddControlPolicy operation to create an access control policy to allow, deny, or monitor traffic that passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddControlPolicyRequest
 * @return AddControlPolicyResponse
 */
func (client *Client) AddControlPolicy(request *AddControlPolicyRequest) (_result *AddControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.AddControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the AddInstanceMembers operation to add members to Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddInstanceMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddInstanceMembersResponse
 */
func (client *Client) AddInstanceMembersWithOptions(request *AddInstanceMembersRequest, runtime *util.RuntimeOptions) (_result *AddInstanceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		query["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddInstanceMembers"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the AddInstanceMembers operation to add members to Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddInstanceMembersRequest
 * @return AddInstanceMembersResponse
 */
func (client *Client) AddInstanceMembers(request *AddInstanceMembersRequest) (_result *AddInstanceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddInstanceMembersResponse{}
	_body, _err := client.AddInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the BatchCopyVpcFirewallControlPolicy operation to copy all access control policies from a policy group of a source VPC firewall to a policy group of a destination VPC firewall.
 * Before you call this operation, we recommend that you back up access control policies. For more information about how to back up an access control policy, see [Back up an access control policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/back-up-and-roll-back-an-access-control-policy).
 * After you call this operation, all the access control policies in the policy group of the destination VPC firewall are replaced.
 * The policy groups of the source VPC firewall and the destination VPC firewall must belong to the same Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request BatchCopyVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BatchCopyVpcFirewallControlPolicyResponse
 */
func (client *Client) BatchCopyVpcFirewallControlPolicyWithOptions(request *BatchCopyVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceVpcFirewallId)) {
		query["SourceVpcFirewallId"] = request.SourceVpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVpcFirewallId)) {
		query["TargetVpcFirewallId"] = request.TargetVpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCopyVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCopyVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the BatchCopyVpcFirewallControlPolicy operation to copy all access control policies from a policy group of a source VPC firewall to a policy group of a destination VPC firewall.
 * Before you call this operation, we recommend that you back up access control policies. For more information about how to back up an access control policy, see [Back up an access control policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/back-up-and-roll-back-an-access-control-policy).
 * After you call this operation, all the access control policies in the policy group of the destination VPC firewall are replaced.
 * The policy groups of the source VPC firewall and the destination VPC firewall must belong to the same Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request BatchCopyVpcFirewallControlPolicyRequest
 * @return BatchCopyVpcFirewallControlPolicyResponse
 */
func (client *Client) BatchCopyVpcFirewallControlPolicy(request *BatchCopyVpcFirewallControlPolicyRequest) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCopyVpcFirewallControlPolicyResponse{}
	_body, _err := client.BatchCopyVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallCenConfigure operation to create a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. The VPC firewall cannot protect mutual access traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC firewall limits](https://www.alibabacloud.com/help/en/cloud-firewall/latest/vpc-firewall-limits).
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request CreateVpcFirewallCenConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcFirewallCenConfigureResponse
 */
func (client *Client) CreateVpcFirewallCenConfigureWithOptions(request *CreateVpcFirewallCenConfigureRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegion)) {
		query["VpcRegion"] = request.VpcRegion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcFirewallCenConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallCenConfigure operation to create a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. The VPC firewall cannot protect mutual access traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC firewall limits](https://www.alibabacloud.com/help/en/cloud-firewall/latest/vpc-firewall-limits).
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request CreateVpcFirewallCenConfigureRequest
 * @return CreateVpcFirewallCenConfigureResponse
 */
func (client *Client) CreateVpcFirewallCenConfigure(request *CreateVpcFirewallCenConfigureRequest) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcFirewallCenConfigureResponse{}
	_body, _err := client.CreateVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallConfigure operation to create a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. The VPC firewall does not control the mutual access traffic between VPCs that reside in different regions or belong to different Alibaba Cloud accounts. The firewall also does not control the mutual access traffic between VPCs and virtual border routers (VBRs). For more information, see [Limits on VPC firewalls](https://www.alibabacloud.com/help/en/cloud-firewall/latest/vpc-firewall-limits).
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcFirewallConfigureResponse
 */
func (client *Client) CreateVpcFirewallConfigureWithOptions(request *CreateVpcFirewallConfigureRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcCidrTableList)) {
		query["LocalVpcCidrTableList"] = request.LocalVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcId)) {
		query["LocalVpcId"] = request.LocalVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcRegion)) {
		query["LocalVpcRegion"] = request.LocalVpcRegion
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcCidrTableList)) {
		query["PeerVpcCidrTableList"] = request.PeerVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcId)) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcRegion)) {
		query["PeerVpcRegion"] = request.PeerVpcRegion
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcFirewallConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallConfigure operation to create a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. The VPC firewall does not control the mutual access traffic between VPCs that reside in different regions or belong to different Alibaba Cloud accounts. The firewall also does not control the mutual access traffic between VPCs and virtual border routers (VBRs). For more information, see [Limits on VPC firewalls](https://www.alibabacloud.com/help/en/cloud-firewall/latest/vpc-firewall-limits).
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallConfigureRequest
 * @return CreateVpcFirewallConfigureResponse
 */
func (client *Client) CreateVpcFirewallConfigure(request *CreateVpcFirewallConfigureRequest) (_result *CreateVpcFirewallConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcFirewallConfigureResponse{}
	_body, _err := client.CreateVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallControlPolicy operation to create an access control policy in a specified policy group for a VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcFirewallControlPolicyResponse
 */
func (client *Client) CreateVpcFirewallControlPolicyWithOptions(request *CreateVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallControlPolicy operation to create an access control policy in a specified policy group for a VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallControlPolicyRequest
 * @return CreateVpcFirewallControlPolicyResponse
 */
func (client *Client) CreateVpcFirewallControlPolicy(request *CreateVpcFirewallControlPolicyRequest) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.CreateVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteAddressBook operation to delete an address book for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAddressBookResponse
 */
func (client *Client) DeleteAddressBookWithOptions(request *DeleteAddressBookRequest, runtime *util.RuntimeOptions) (_result *DeleteAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupUuid)) {
		query["GroupUuid"] = request.GroupUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteAddressBook operation to delete an address book for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteAddressBookRequest
 * @return DeleteAddressBookResponse
 */
func (client *Client) DeleteAddressBook(request *DeleteAddressBookRequest) (_result *DeleteAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.DeleteAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteControlPolicy operation to delete an access control policy that applies to inbound or outbound traffic.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteControlPolicyResponse
 */
func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteControlPolicy operation to delete an access control policy that applies to inbound or outbound traffic.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteControlPolicyRequest
 * @return DeleteControlPolicyResponse
 */
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (_result *DeleteControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DeleteControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteInstanceMembers operation to remove members from Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteInstanceMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteInstanceMembersResponse
 */
func (client *Client) DeleteInstanceMembersWithOptions(request *DeleteInstanceMembersRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberUids)) {
		query["MemberUids"] = request.MemberUids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceMembers"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteInstanceMembers operation to remove members from Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteInstanceMembersRequest
 * @return DeleteInstanceMembersResponse
 */
func (client *Client) DeleteInstanceMembers(request *DeleteInstanceMembersRequest) (_result *DeleteInstanceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceMembersResponse{}
	_body, _err := client.DeleteInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallCenConfigure operation to delete a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request DeleteVpcFirewallCenConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcFirewallCenConfigureResponse
 */
func (client *Client) DeleteVpcFirewallCenConfigureWithOptions(request *DeleteVpcFirewallCenConfigureRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallIdList)) {
		query["VpcFirewallIdList"] = request.VpcFirewallIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcFirewallCenConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallCenConfigure operation to delete a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request DeleteVpcFirewallCenConfigureRequest
 * @return DeleteVpcFirewallCenConfigureResponse
 */
func (client *Client) DeleteVpcFirewallCenConfigure(request *DeleteVpcFirewallCenConfigureRequest) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcFirewallCenConfigureResponse{}
	_body, _err := client.DeleteVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallConfigure operation to delete a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcFirewallConfigureResponse
 */
func (client *Client) DeleteVpcFirewallConfigureWithOptions(request *DeleteVpcFirewallConfigureRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallIdList)) {
		query["VpcFirewallIdList"] = request.VpcFirewallIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcFirewallConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallConfigure operation to delete a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallConfigureRequest
 * @return DeleteVpcFirewallConfigureResponse
 */
func (client *Client) DeleteVpcFirewallConfigure(request *DeleteVpcFirewallConfigureRequest) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcFirewallConfigureResponse{}
	_body, _err := client.DeleteVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallControlPolicy operation to delete an access control policy from a specific policy group for a VPC firewall. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcFirewallControlPolicyResponse
 */
func (client *Client) DeleteVpcFirewallControlPolicyWithOptions(request *DeleteVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallControlPolicy operation to delete an access control policy from a specific policy group for a VPC firewall. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallControlPolicyRequest
 * @return DeleteVpcFirewallControlPolicyResponse
 */
func (client *Client) DeleteVpcFirewallControlPolicy(request *DeleteVpcFirewallControlPolicyRequest) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.DeleteVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeAddressBook operation to query the details about an address book for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAddressBookResponse
 */
func (client *Client) DescribeAddressBookWithOptions(request *DescribeAddressBookRequest, runtime *util.RuntimeOptions) (_result *DescribeAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainPort)) {
		query["ContainPort"] = request.ContainPort
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeAddressBook operation to query the details about an address book for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAddressBookRequest
 * @return DescribeAddressBookResponse
 */
func (client *Client) DescribeAddressBook(request *DescribeAddressBookRequest) (_result *DescribeAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAddressBookResponse{}
	_body, _err := client.DescribeAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeAssetList operation to query the assets that are protected by Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAssetListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAssetListResponse
 */
func (client *Client) DescribeAssetListWithOptions(request *DescribeAssetListRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceTag)) {
		query["NewResourceTag"] = request.NewResourceTag
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchItem)) {
		query["SearchItem"] = request.SearchItem
	}

	if !tea.BoolValue(util.IsUnset(request.SgStatus)) {
		query["SgStatus"] = request.SgStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeAssetList operation to query the assets that are protected by Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAssetListRequest
 * @return DescribeAssetListResponse
 */
func (client *Client) DescribeAssetList(request *DescribeAssetListRequest) (_result *DescribeAssetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetListResponse{}
	_body, _err := client.DescribeAssetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeControlPolicy operation to query the details about access control policies by page.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeControlPolicyResponse
 */
func (client *Client) DescribeControlPolicyWithOptions(request *DescribeControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeControlPolicy operation to query the details about access control policies by page.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeControlPolicyRequest
 * @return DescribeControlPolicyResponse
 */
func (client *Client) DescribeControlPolicy(request *DescribeControlPolicyRequest) (_result *DescribeControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeControlPolicyResponse{}
	_body, _err := client.DescribeControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeDomainResolve operation to query the DNS record of a domain name. This operation can retrieve DNS records only from Alibaba Cloud DNS. Before you can call this operation, make sure that your domain name is hosted on Alibaba Cloud DNS.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDomainResolveRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDomainResolveResponse
 */
func (client *Client) DescribeDomainResolveWithOptions(request *DescribeDomainResolveRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainResolveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainResolve"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainResolveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeDomainResolve operation to query the DNS record of a domain name. This operation can retrieve DNS records only from Alibaba Cloud DNS. Before you can call this operation, make sure that your domain name is hosted on Alibaba Cloud DNS.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDomainResolveRequest
 * @return DescribeDomainResolveResponse
 */
func (client *Client) DescribeDomainResolve(request *DescribeDomainResolveRequest) (_result *DescribeDomainResolveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainResolveResponse{}
	_body, _err := client.DescribeDomainResolveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeInstanceMembers operation to query the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeInstanceMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeInstanceMembersResponse
 */
func (client *Client) DescribeInstanceMembersWithOptions(request *DescribeInstanceMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.MemberDesc)) {
		query["MemberDesc"] = request.MemberDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MemberDisplayName)) {
		query["MemberDisplayName"] = request.MemberDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceMembers"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeInstanceMembers operation to query the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeInstanceMembersRequest
 * @return DescribeInstanceMembersResponse
 */
func (client *Client) DescribeInstanceMembers(request *DescribeInstanceMembersRequest) (_result *DescribeInstanceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceMembersResponse{}
	_body, _err := client.DescribeInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvadeEventListWithOptions(request *DescribeInvadeEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeInvadeEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetsIP)) {
		query["AssetsIP"] = request.AssetsIP
	}

	if !tea.BoolValue(util.IsUnset(request.AssetsInstanceId)) {
		query["AssetsInstanceId"] = request.AssetsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AssetsInstanceName)) {
		query["AssetsInstanceName"] = request.AssetsInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventKey)) {
		query["EventKey"] = request.EventKey
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventUuid)) {
		query["EventUuid"] = request.EventUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IsIgnore)) {
		query["IsIgnore"] = request.IsIgnore
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessStatusList)) {
		query["ProcessStatusList"] = request.ProcessStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvadeEventList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvadeEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvadeEventList(request *DescribeInvadeEventListRequest) (_result *DescribeInvadeEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvadeEventListResponse{}
	_body, _err := client.DescribeInvadeEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOutgoingDestinationIPWithOptions(request *DescribeOutgoingDestinationIPRequest, runtime *util.RuntimeOptions) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DstIP)) {
		query["DstIP"] = request.DstIP
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIP)) {
		query["PrivateIP"] = request.PrivateIP
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIP)) {
		query["PublicIP"] = request.PublicIP
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdNew)) {
		query["TagIdNew"] = request.TagIdNew
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOutgoingDestinationIP"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOutgoingDestinationIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOutgoingDestinationIP(request *DescribeOutgoingDestinationIPRequest) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOutgoingDestinationIPResponse{}
	_body, _err := client.DescribeOutgoingDestinationIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOutgoingDomainWithOptions(request *DescribeOutgoingDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeOutgoingDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIP)) {
		query["PublicIP"] = request.PublicIP
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdNew)) {
		query["TagIdNew"] = request.TagIdNew
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOutgoingDomain"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOutgoingDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOutgoingDomain(request *DescribeOutgoingDomainRequest) (_result *DescribeOutgoingDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOutgoingDomainResponse{}
	_body, _err := client.DescribeOutgoingDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribePolicyAdvancedConfig operation to query whether the strict mode is enabled for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyAdvancedConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribePolicyAdvancedConfigResponse
 */
func (client *Client) DescribePolicyAdvancedConfigWithOptions(request *DescribePolicyAdvancedConfigRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyAdvancedConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyAdvancedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribePolicyAdvancedConfig operation to query whether the strict mode is enabled for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyAdvancedConfigRequest
 * @return DescribePolicyAdvancedConfigResponse
 */
func (client *Client) DescribePolicyAdvancedConfig(request *DescribePolicyAdvancedConfigRequest) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyAdvancedConfigResponse{}
	_body, _err := client.DescribePolicyAdvancedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribePolicyPriorUsed operation to query the priority range of the access control policies that match specific query conditions.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyPriorUsedRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribePolicyPriorUsedResponse
 */
func (client *Client) DescribePolicyPriorUsedWithOptions(request *DescribePolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyPriorUsed"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribePolicyPriorUsed operation to query the priority range of the access control policies that match specific query conditions.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyPriorUsedRequest
 * @return DescribePolicyPriorUsedResponse
 */
func (client *Client) DescribePolicyPriorUsed(request *DescribePolicyPriorUsedRequest) (_result *DescribePolicyPriorUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.DescribePolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeRiskEventGroup operation to query and download the details of intrusion events. We recommend that you query the details of 5 to 10 intrusion events at a time. If you do not need to query the information about the geographical locations of IP addresses, set the NoLocation parameter to true. This prevents query timeout.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRiskEventGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeRiskEventGroupResponse
 */
func (client *Client) DescribeRiskEventGroupWithOptions(request *DescribeRiskEventGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskEventGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttackApp)) {
		query["AttackApp"] = request.AttackApp
	}

	if !tea.BoolValue(util.IsUnset(request.AttackType)) {
		query["AttackType"] = request.AttackType
	}

	if !tea.BoolValue(util.IsUnset(request.BuyVersion)) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.DstIP)) {
		query["DstIP"] = request.DstIP
	}

	if !tea.BoolValue(util.IsUnset(request.DstNetworkInstanceId)) {
		query["DstNetworkInstanceId"] = request.DstNetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallType)) {
		query["FirewallType"] = request.FirewallType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NoLocation)) {
		query["NoLocation"] = request.NoLocation
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleResult)) {
		query["RuleResult"] = request.RuleResult
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSource)) {
		query["RuleSource"] = request.RuleSource
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.SrcIP)) {
		query["SrcIP"] = request.SrcIP
	}

	if !tea.BoolValue(util.IsUnset(request.SrcNetworkInstanceId)) {
		query["SrcNetworkInstanceId"] = request.SrcNetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.VulLevel)) {
		query["VulLevel"] = request.VulLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRiskEventGroup"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRiskEventGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeRiskEventGroup operation to query and download the details of intrusion events. We recommend that you query the details of 5 to 10 intrusion events at a time. If you do not need to query the information about the geographical locations of IP addresses, set the NoLocation parameter to true. This prevents query timeout.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRiskEventGroupRequest
 * @return DescribeRiskEventGroupResponse
 */
func (client *Client) DescribeRiskEventGroup(request *DescribeRiskEventGroupRequest) (_result *DescribeRiskEventGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskEventGroupResponse{}
	_body, _err := client.DescribeRiskEventGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserAssetIPTrafficInfoWithOptions(request *DescribeUserAssetIPTrafficInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserAssetIPTrafficInfo"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserAssetIPTrafficInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserAssetIPTrafficInfo(request *DescribeUserAssetIPTrafficInfoRequest) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserAssetIPTrafficInfoResponse{}
	_body, _err := client.DescribeUserAssetIPTrafficInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallAclGroupList operation to query the information about all policy groups of access control policies that are created for VPC firewalls.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallAclGroupListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallAclGroupListResponse
 */
func (client *Client) DescribeVpcFirewallAclGroupListWithOptions(request *DescribeVpcFirewallAclGroupListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallConfigureStatus)) {
		query["FirewallConfigureStatus"] = request.FirewallConfigureStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallAclGroupList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallAclGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallAclGroupList operation to query the information about all policy groups of access control policies that are created for VPC firewalls.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallAclGroupListRequest
 * @return DescribeVpcFirewallAclGroupListResponse
 */
func (client *Client) DescribeVpcFirewallAclGroupList(request *DescribeVpcFirewallAclGroupListRequest) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallAclGroupListResponse{}
	_body, _err := client.DescribeVpcFirewallAclGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenDetail operation to query the details about a VPC firewall. The VPC firewall controls mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallCenDetailResponse
 */
func (client *Client) DescribeVpcFirewallCenDetailWithOptions(request *DescribeVpcFirewallCenDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallCenDetail"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallCenDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenDetail operation to query the details about a VPC firewall. The VPC firewall controls mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenDetailRequest
 * @return DescribeVpcFirewallCenDetailResponse
 */
func (client *Client) DescribeVpcFirewallCenDetail(request *DescribeVpcFirewallCenDetailRequest) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenDetailResponse{}
	_body, _err := client.DescribeVpcFirewallCenDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenList operation to query VPC firewalls. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallCenListResponse
 */
func (client *Client) DescribeVpcFirewallCenListWithOptions(request *DescribeVpcFirewallCenListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallCenListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitchStatus)) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMode)) {
		query["RouteMode"] = request.RouteMode
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterType)) {
		query["TransitRouterType"] = request.TransitRouterType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallCenList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallCenListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenList operation to query VPC firewalls. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenListRequest
 * @return DescribeVpcFirewallCenListResponse
 */
func (client *Client) DescribeVpcFirewallCenList(request *DescribeVpcFirewallCenListRequest) (_result *DescribeVpcFirewallCenListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenListResponse{}
	_body, _err := client.DescribeVpcFirewallCenListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallControlPolicy operation to query the details of all access control policies that are created for a specified VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 *
 * @param request DescribeVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallControlPolicyResponse
 */
func (client *Client) DescribeVpcFirewallControlPolicyWithOptions(request *DescribeVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallControlPolicy operation to query the details of all access control policies that are created for a specified VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 *
 * @param request DescribeVpcFirewallControlPolicyRequest
 * @return DescribeVpcFirewallControlPolicyResponse
 */
func (client *Client) DescribeVpcFirewallControlPolicy(request *DescribeVpcFirewallControlPolicyRequest) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.DescribeVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDefaultIPSConfig operation to query the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDefaultIPSConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) DescribeVpcFirewallDefaultIPSConfigWithOptions(request *DescribeVpcFirewallDefaultIPSConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallDefaultIPSConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDefaultIPSConfig operation to query the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDefaultIPSConfigRequest
 * @return DescribeVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) DescribeVpcFirewallDefaultIPSConfig(request *DescribeVpcFirewallDefaultIPSConfigRequest) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.DescribeVpcFirewallDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDetail operation to query the details about a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallDetailResponse
 */
func (client *Client) DescribeVpcFirewallDetailWithOptions(request *DescribeVpcFirewallDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcId)) {
		query["LocalVpcId"] = request.LocalVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcId)) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallDetail"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDetail operation to query the details about a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDetailRequest
 * @return DescribeVpcFirewallDetailResponse
 */
func (client *Client) DescribeVpcFirewallDetail(request *DescribeVpcFirewallDetailRequest) (_result *DescribeVpcFirewallDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallDetailResponse{}
	_body, _err := client.DescribeVpcFirewallDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallList operation to query the details about VPC firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallListResponse
 */
func (client *Client) DescribeVpcFirewallListWithOptions(request *DescribeVpcFirewallListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectSubType)) {
		query["ConnectSubType"] = request.ConnectSubType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitchStatus)) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeerUid)) {
		query["PeerUid"] = request.PeerUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallList operation to query the details about VPC firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallListRequest
 * @return DescribeVpcFirewallListResponse
 */
func (client *Client) DescribeVpcFirewallList(request *DescribeVpcFirewallListRequest) (_result *DescribeVpcFirewallListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallListResponse{}
	_body, _err := client.DescribeVpcFirewallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallPolicyPriorUsed operation to query the priority range of access control policies that are created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallPolicyPriorUsedRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallPolicyPriorUsedResponse
 */
func (client *Client) DescribeVpcFirewallPolicyPriorUsedWithOptions(request *DescribeVpcFirewallPolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallPolicyPriorUsed"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallPolicyPriorUsed operation to query the priority range of access control policies that are created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallPolicyPriorUsedRequest
 * @return DescribeVpcFirewallPolicyPriorUsedResponse
 */
func (client *Client) DescribeVpcFirewallPolicyPriorUsed(request *DescribeVpcFirewallPolicyPriorUsedRequest) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.DescribeVpcFirewallPolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulnerabilityProtectedListWithOptions(request *DescribeVulnerabilityProtectedListRequest, runtime *util.RuntimeOptions) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttackType)) {
		query["AttackType"] = request.AttackType
	}

	if !tea.BoolValue(util.IsUnset(request.BuyVersion)) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["SortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	if !tea.BoolValue(util.IsUnset(request.VulnCveName)) {
		query["VulnCveName"] = request.VulnCveName
	}

	if !tea.BoolValue(util.IsUnset(request.VulnLevel)) {
		query["VulnLevel"] = request.VulnLevel
	}

	if !tea.BoolValue(util.IsUnset(request.VulnResource)) {
		query["VulnResource"] = request.VulnResource
	}

	if !tea.BoolValue(util.IsUnset(request.VulnStatus)) {
		query["VulnStatus"] = request.VulnStatus
	}

	if !tea.BoolValue(util.IsUnset(request.VulnType)) {
		query["VulnType"] = request.VulnType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVulnerabilityProtectedList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVulnerabilityProtectedListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulnerabilityProtectedList(request *DescribeVulnerabilityProtectedListRequest) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulnerabilityProtectedListResponse{}
	_body, _err := client.DescribeVulnerabilityProtectedListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyAddressBook operation to modify the address book that is configured for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAddressBookResponse
 */
func (client *Client) ModifyAddressBookWithOptions(request *ModifyAddressBookRequest, runtime *util.RuntimeOptions) (_result *ModifyAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressList)) {
		query["AddressList"] = request.AddressList
	}

	if !tea.BoolValue(util.IsUnset(request.AutoAddTagEcs)) {
		query["AutoAddTagEcs"] = request.AutoAddTagEcs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUuid)) {
		query["GroupUuid"] = request.GroupUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TagList)) {
		query["TagList"] = request.TagList
	}

	if !tea.BoolValue(util.IsUnset(request.TagRelation)) {
		query["TagRelation"] = request.TagRelation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyAddressBook operation to modify the address book that is configured for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyAddressBookRequest
 * @return ModifyAddressBookResponse
 */
func (client *Client) ModifyAddressBook(request *ModifyAddressBookRequest) (_result *ModifyAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.ModifyAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyControlPolicy operation to modify the configurations of an access control policy. The policy allows Cloud Firewall to allow, deny, or monitor the traffic that passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyControlPolicyResponse
 */
func (client *Client) ModifyControlPolicyWithOptions(request *ModifyControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyControlPolicy operation to modify the configurations of an access control policy. The policy allows Cloud Firewall to allow, deny, or monitor the traffic that passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyRequest
 * @return ModifyControlPolicyResponse
 */
func (client *Client) ModifyControlPolicy(request *ModifyControlPolicyRequest) (_result *ModifyControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.ModifyControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyControlPolicyPosition operation to modify the priority of an IPv4 access control policy for the Internet firewall. No API operations are provided for you to modify the priority of an IPv6 access control policy for the Internet firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyPositionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyControlPolicyPositionResponse
 */
func (client *Client) ModifyControlPolicyPositionWithOptions(request *ModifyControlPolicyPositionRequest, runtime *util.RuntimeOptions) (_result *ModifyControlPolicyPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OldOrder)) {
		query["OldOrder"] = request.OldOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyControlPolicyPosition"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyControlPolicyPosition operation to modify the priority of an IPv4 access control policy for the Internet firewall. No API operations are provided for you to modify the priority of an IPv6 access control policy for the Internet firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyPositionRequest
 * @return ModifyControlPolicyPositionResponse
 */
func (client *Client) ModifyControlPolicyPosition(request *ModifyControlPolicyPositionRequest) (_result *ModifyControlPolicyPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.ModifyControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyInstanceMemberAttributes operation to update the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second for each account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyInstanceMemberAttributesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyInstanceMemberAttributesResponse
 */
func (client *Client) ModifyInstanceMemberAttributesWithOptions(request *ModifyInstanceMemberAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		query["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceMemberAttributes"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceMemberAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyInstanceMemberAttributes operation to update the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second for each account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyInstanceMemberAttributesRequest
 * @return ModifyInstanceMemberAttributesResponse
 */
func (client *Client) ModifyInstanceMemberAttributes(request *ModifyInstanceMemberAttributesRequest) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMemberAttributesResponse{}
	_body, _err := client.ModifyInstanceMemberAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyPolicyAdvancedConfig operation to enable or disable the strict mode for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyPolicyAdvancedConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyPolicyAdvancedConfigResponse
 */
func (client *Client) ModifyPolicyAdvancedConfigWithOptions(request *ModifyPolicyAdvancedConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternetSwitch)) {
		query["InternetSwitch"] = request.InternetSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicyAdvancedConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyPolicyAdvancedConfig operation to enable or disable the strict mode for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyPolicyAdvancedConfigRequest
 * @return ModifyPolicyAdvancedConfigResponse
 */
func (client *Client) ModifyPolicyAdvancedConfig(request *ModifyPolicyAdvancedConfigRequest) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.ModifyPolicyAdvancedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenConfigure operation to modify the configurations of a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallCenConfigureResponse
 */
func (client *Client) ModifyVpcFirewallCenConfigureWithOptions(request *ModifyVpcFirewallCenConfigureRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallCenConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenConfigure operation to modify the configurations of a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenConfigureRequest
 * @return ModifyVpcFirewallCenConfigureResponse
 */
func (client *Client) ModifyVpcFirewallCenConfigure(request *ModifyVpcFirewallCenConfigureRequest) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallCenConfigureResponse{}
	_body, _err := client.ModifyVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenSwitchStatus operation to enable or disable a VPC firewall. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. After you enable the VPC firewall, the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. After you disable the VPC firewall, the VPC firewall no longer protect mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenSwitchStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallCenSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallCenSwitchStatusWithOptions(request *ModifyVpcFirewallCenSwitchStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallCenSwitchStatus"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallCenSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenSwitchStatus operation to enable or disable a VPC firewall. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. After you enable the VPC firewall, the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. After you disable the VPC firewall, the VPC firewall no longer protect mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenSwitchStatusRequest
 * @return ModifyVpcFirewallCenSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallCenSwitchStatus(request *ModifyVpcFirewallCenSwitchStatusRequest) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallCenSwitchStatusResponse{}
	_body, _err := client.ModifyVpcFirewallCenSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallConfigure operation to modify the configurations of a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallConfigureResponse
 */
func (client *Client) ModifyVpcFirewallConfigureWithOptions(request *ModifyVpcFirewallConfigureRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcCidrTableList)) {
		query["LocalVpcCidrTableList"] = request.LocalVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcCidrTableList)) {
		query["PeerVpcCidrTableList"] = request.PeerVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallConfigure operation to modify the configurations of a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallConfigureRequest
 * @return ModifyVpcFirewallConfigureResponse
 */
func (client *Client) ModifyVpcFirewallConfigure(request *ModifyVpcFirewallConfigureRequest) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallConfigureResponse{}
	_body, _err := client.ModifyVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallControlPolicy operation to modify the configurations of an access control policy that is created for a VPC firewall in a specific policy group. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallControlPolicyResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicyWithOptions(request *ModifyVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallControlPolicy operation to modify the configurations of an access control policy that is created for a VPC firewall in a specific policy group. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyRequest
 * @return ModifyVpcFirewallControlPolicyResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicy(request *ModifyVpcFirewallControlPolicyRequest) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.ModifyVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallControlPolicyPosition operation to modify the priority of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyPositionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallControlPolicyPositionResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicyPositionWithOptions(request *ModifyVpcFirewallControlPolicyPositionRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OldOrder)) {
		query["OldOrder"] = request.OldOrder
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallControlPolicyPosition"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallControlPolicyPosition operation to modify the priority of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyPositionRequest
 * @return ModifyVpcFirewallControlPolicyPositionResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicyPosition(request *ModifyVpcFirewallControlPolicyPositionRequest) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.ModifyVpcFirewallControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallDefaultIPSConfig operation to modify the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallDefaultIPSConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) ModifyVpcFirewallDefaultIPSConfigWithOptions(request *ModifyVpcFirewallDefaultIPSConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BasicRules)) {
		query["BasicRules"] = request.BasicRules
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAllPatch)) {
		query["EnableAllPatch"] = request.EnableAllPatch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.RunMode)) {
		query["RunMode"] = request.RunMode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallDefaultIPSConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallDefaultIPSConfig operation to modify the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallDefaultIPSConfigRequest
 * @return ModifyVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) ModifyVpcFirewallDefaultIPSConfig(request *ModifyVpcFirewallDefaultIPSConfigRequest) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.ModifyVpcFirewallDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallSwitchStatus operation to enable or disable a VPC firewall. The VPC firewall can control traffic between two VPCs that are connected by using an Express Connect circuit. After you enable the VPC firewall, the VPC firewall controls mutual access traffic between two VPCs that are connected by using an Express Connect circuit. After you disable the VPC firewall, the VPC firewall no longer controls mutual access traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallSwitchStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallSwitchStatusWithOptions(request *ModifyVpcFirewallSwitchStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallSwitchStatus"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallSwitchStatus operation to enable or disable a VPC firewall. The VPC firewall can control traffic between two VPCs that are connected by using an Express Connect circuit. After you enable the VPC firewall, the VPC firewall controls mutual access traffic between two VPCs that are connected by using an Express Connect circuit. After you disable the VPC firewall, the VPC firewall no longer controls mutual access traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallSwitchStatusRequest
 * @return ModifyVpcFirewallSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallSwitchStatus(request *ModifyVpcFirewallSwitchStatusRequest) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallSwitchStatusResponse{}
	_body, _err := client.ModifyVpcFirewallSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutDisableAllFwSwitch operation to turn off all firewall switches.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableAllFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutDisableAllFwSwitchResponse
 */
func (client *Client) PutDisableAllFwSwitchWithOptions(request *PutDisableAllFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutDisableAllFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDisableAllFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutDisableAllFwSwitch operation to turn off all firewall switches.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableAllFwSwitchRequest
 * @return PutDisableAllFwSwitchResponse
 */
func (client *Client) PutDisableAllFwSwitch(request *PutDisableAllFwSwitchRequest) (_result *PutDisableAllFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.PutDisableAllFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutDisableFwSwitch operation to disable a firewall for specific assets. After you disable the firewall, traffic does not pass through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutDisableFwSwitchResponse
 */
func (client *Client) PutDisableFwSwitchWithOptions(request *PutDisableFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutDisableFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpaddrList)) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionList)) {
		query["RegionList"] = request.RegionList
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeList)) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDisableFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutDisableFwSwitch operation to disable a firewall for specific assets. After you disable the firewall, traffic does not pass through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableFwSwitchRequest
 * @return PutDisableFwSwitchResponse
 */
func (client *Client) PutDisableFwSwitch(request *PutDisableFwSwitchRequest) (_result *PutDisableFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.PutDisableFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutEnableAllFwSwitch operation to enable a firewall for all public IP addresses within your Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutEnableAllFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutEnableAllFwSwitchResponse
 */
func (client *Client) PutEnableAllFwSwitchWithOptions(request *PutEnableAllFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutEnableAllFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnableAllFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutEnableAllFwSwitch operation to enable a firewall for all public IP addresses within your Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutEnableAllFwSwitchRequest
 * @return PutEnableAllFwSwitchResponse
 */
func (client *Client) PutEnableAllFwSwitch(request *PutEnableAllFwSwitchRequest) (_result *PutEnableAllFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.PutEnableAllFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutEnableFwSwitch operation to enable a firewall. After you enable a firewall, traffic passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutEnableFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutEnableFwSwitchResponse
 */
func (client *Client) PutEnableFwSwitchWithOptions(request *PutEnableFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutEnableFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpaddrList)) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionList)) {
		query["RegionList"] = request.RegionList
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeList)) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnableFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutEnableFwSwitch operation to enable a firewall. After you enable a firewall, traffic passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutEnableFwSwitchRequest
 * @return PutEnableFwSwitchResponse
 */
func (client *Client) PutEnableFwSwitch(request *PutEnableFwSwitchRequest) (_result *PutEnableFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.PutEnableFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ResetVpcFirewallRuleHitCount operation to clear the count on hits of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ResetVpcFirewallRuleHitCountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetVpcFirewallRuleHitCountResponse
 */
func (client *Client) ResetVpcFirewallRuleHitCountWithOptions(request *ResetVpcFirewallRuleHitCountRequest, runtime *util.RuntimeOptions) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetVpcFirewallRuleHitCount"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ResetVpcFirewallRuleHitCount operation to clear the count on hits of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ResetVpcFirewallRuleHitCountRequest
 * @return ResetVpcFirewallRuleHitCountResponse
 */
func (client *Client) ResetVpcFirewallRuleHitCount(request *ResetVpcFirewallRuleHitCountRequest) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.ResetVpcFirewallRuleHitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
