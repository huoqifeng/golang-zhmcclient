/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

//////////////////////////////////////////////////
// Adapter
//////////////////////////////////////////////////

type AdapterFamily string

const (
	ADAPTER_FAMILY_HIPERSOCKET AdapterFamily = "hipersockets"
	ADAPTER_FAMILY_OSA                       = "osa"
	ADAPTER_FAMILY_FICON                     = "ficon"
	ADAPTER_FAMILY_ROCE                      = "roce"
	ADAPTER_FAMILY_CRYPTO                    = "crypto"
	ADAPTER_FAMILY_ACCELERATOR               = "accelerator"
)

type AdapterType string

const (
	ADAPTER_TYPE_CRYPTO      AdapterType = "crypto"
	ADAPTER_TYPE_FCP                     = "fcp"
	ADAPTER_TYPE_HIPERSOCKET             = "hipersockets"
	ADAPTER_TYPE_OSD                     = "osd"
	ADAPTER_TYPE_OSM                     = "osm"
	ADAPTER_TYPE_ROCE                    = "roce"
	ADAPTER_TYPE_ZEDC                    = "zedc"
	ADAPTER_TYPE_FC                      = "fc"
	ADAPTER_TYPE_NOT_CFG                 = "not-configured"
)

type AdapterStatus string

const (
	ADAPTER_STATUS_ACTIVE       AdapterStatus = "active"
	ADAPTER_STATUS_NOT_ACTIVE                 = "not-active"
	ADAPTER_STATUS_NOT_DETECTED               = "not-detected"
	ADAPTER_STATUS_EXCEPTIONS                 = "exceptions"
)

type AdapterCardType string

const (
	ADPATER_CARD_HIPERSOCKETS              AdapterCardType = "hipersockets"
	ADPATER_CARD_OSA_EXPRESS_4S_1GB                        = "osa-express-4s-1gb"
	ADPATER_CARD_OSA_EXPRESS_4S_10GB                       = "osa-express-4s-10gb"
	ADPATER_CARD_OSA_EXPRESS_4S_1000BASE_T                 = "osa-express-4s-1000base-t"
	ADPATER_CARD_OSA_EXPRESS_5S_1GB                        = "osa-express-5s-1gb"
	ADPATER_CARD_OSA_EXPRESS_5S_10GB                       = "osa-express-5s-10gb"
	ADPATER_CARD_OSA_EXPRESS_5G_1000BASE_T                 = "osa-express-5s-1000base-t"
	ADPATER_CARD_OSA_EXPRESS_6S_1GB                        = "osa-express-6s-1gb"
	ADPATER_CARD_OSA_EXPRESS_6S_10GB                       = "osa-express-6s-10gb"
	ADPATER_CARD_OSA_EXPRESS_6G_1000BASE_T                 = "osa-express-6s-1000base-t"
	ADPATER_CARD_OSA_EXPRESS_7S_25GB                       = "osa-express-7s-25gb"
	ADPATER_CARD_10GBE_ROCE_EXPRESS                        = "10gbe-roce-express"
	ADPATER_CARD_ROCE_EXPRESS_2                            = "roce-express-2"
	ADAPTER_CARD_ROCE_EXPRESS_2_25GB                       = "roce-express-2-25gb"
	ADAPTER_CARD_CRYPTO_EXPRESS_5S                         = "crypto-express-5s"
	ADAPTER_CARD_CRYPTO_EXPRESS_6S                         = "crypto-express-6s"
	ADAPTER_CARD_FICON_EXPRESS_8                           = "ficon-express-8"
	ADAPTER_CARD_FICON_EXPRESS_8S                          = "ficon-express-8s"
	ADAPTER_CARD_FICON_EXPRESS_16S                         = "ficon-express-16s"
	ADAPTER_CARD_FICON_EXPRESS_16S_PLUS                    = "ficon-express-16s-plus"
	ADAPTER_CARD_FICON_EXPRESS_16SA                        = "ficon-express-16sa"
	ADAPTER_CARD_FCP_EXPRESS_32S                           = "fcp-express-32s"
	ADAPTER_CARD_ZEDC_EXPRESS                              = "zedc-express"
	ADAPTER_CARD_CLOUD_NETWORK_X5                          = "cloud-network-x5"
	ADAPTER_CARD_UNKNOWN                                   = "unknown"
)

type AdapterState string

const (
	ADAPTER_STATE_ONLINE   AdapterState = "online"
	ADAPTER_STATE_STAND_BY              = "stand-by"
	ADAPTER_STATE_RESERVED              = "reserved"
	ADAPTER_STATE_UNKNOWN               = "unknown"
)

type AdapterTransmissionUnit int

const (
	ADAPTER_TRANSMISSION_UNIT_8  AdapterTransmissionUnit = 8
	ADAPTER_TRANSMISSION_UNIT_16                         = 16
	ADAPTER_TRANSMISSION_UNIT_32                         = 32
	ADAPTER_TRANSMISSION_UNIT_56                         = 56
)

// Physical adapter channel connection status
type AdapterChannelStatus string

const (
	ADAPTER_CHANNEL_OPERATING                         AdapterChannelStatus = "operating"
	ADAPTER_CHANNEL_NO_POWER                                               = "no-power"
	ADAPTER_CHANNEL_SERVICE                                                = "service"
	ADAPTER_CHANNEL_STOPPED                                                = "stopped"
	ADAPTER_CHANNEL_NOT_DEFINED                                            = "not-defined"
	ADAPTER_CHANNEL_SUSPENDED                                              = "suspended"
	ADAPTER_CHANNEL_CHECK_STOPPED                                          = "check-stopped"
	ADAPTER_CHANNEL_WRAP_BLOCK                                             = "wrap-block"
	ADAPTER_CHANNEL_PERMANENT_ERROR                                        = "permanent-error"
	ADAPTER_CHANNEL_INITIALIZING                                           = "initializing"
	ADAPTER_CHANNEL_LOSS_OF_SIGNAL                                         = "loss-of-signal"
	ADAPTER_CHANNEL_LOSS_OF_SYNCHRONIZATION                                = "loss-of-synchronization"
	ADAPTER_CHANNEL_NOT_OPERATIONAL_LINK                                   = "not-operational-link"
	ADAPTER_CHANNEL_SEQUENCE_TIME_OUT                                      = "sequence-time-out"
	ADAPTER_CHANNEL_SEQUENCE_NOT_PERMITTED                                 = "sequence-not-permitted"
	ADAPTER_CHANNEL_TERMINAL_CONDITION                                     = "terminal-condition"
	ADAPTER_CHANNEL_OFFLINE_SIGNAL_RECEIVED                                = "offline-signal-received"
	ADAPTER_CHANNEL_FABRIC_LOGIN_SEQUENCE_FAILURE                          = "fabric-login-sequence-failure"
	ADAPTER_CHANNEL_POST_LOGIN_SEQUENCE_FAILURE                            = "post-login-sequence-failure"
	ADAPTER_CHANNEL_STATE_CHANGE_REGISTRATION_FAILURE                      = "state-change-registration-failure"
	ADAPTER_CHANNEL_INVALID_ATTACHMENT_FAILURE                             = "invalid-attachment-failure"
	ADAPTER_CHANNEL_TEST_MODE                                              = "test-mode"
	ADAPTER_CHANNEL_BIT_ERROR_THRESHOLD_EXCEEDED                           = "bit-error-threshold-exceeded"
	ADAPTER_CHANNEL_IFCC_THRESHOLD_EXCEEDED                                = "ifcc-threshold-exceeded"
	ADAPTER_CHANNEL_IO_SUPRESSED                                           = "io-supressed"
)

type Adapter struct {
	URI    string        `json:"object-uri,omitempty"`
	Name   string        `json:"name,omitempty"`
	ID     string        `json:"adapter-id,omitempty"`
	Family AdapterFamily `json:"adapter-family,omitempty"`
	Type   AdapterType   `json:"type,omitempty"`
	Status AdapterStatus `json:"status,omitempty"`
}

type AdaptersArray struct {
	ADAPTERS []Adapter `json:"adapters"`
}

type AdapterProperties struct {
	URI                        string                  `json:"object-uri,omitempty"`
	Name                       string                  `json:"name,omitempty"`
	ID                         string                  `json:"adapter-id,omitempty"`
	ObjectID                   string                  `json:"object-id,omitempty"`
	Description                string                  `json:"description,omitempty"`
	Family                     AdapterFamily           `json:"adapter-family,omitempty"`
	Type                       AdapterType             `json:"type,omitempty"`
	Status                     AdapterStatus           `json:"status,omitempty"`
	DetectedCardType           AdapterCardType         `json:"detected-card-type,omitempty"`
	CardLocation               string                  `json:"card-location,omitempty"`
	PortCount                  int                     `json:"port-count,omitempty"`
	NetworkAdapterPortURIs     []string                `json:"network-port-uris,omitempty"`
	StoragePortURIs            []string                `json:"storage-port-uris,omitempty"`
	State                      AdapterState            `json:"state,omitempty"`
	MAX_TRANSMISSION_UNIT_SIZE AdapterTransmissionUnit `json:"maximum-transmission-unit-size,omitempty"`
	PHYSICALCHANNELSTATUS      AdapterChannelStatus    `json:"physical-channel-status,omitempty"`
	USED_CAPACITY              int                     `json:"used-capacity,omitempty"`
	ALLOWED_CAPACITY           int                     `json:"allowed-capacity,omitempty"`
	MAXIMUM_TOTAL_CAPACITY     int                     `json:"maximum-total-capacity,omitempty"`
	CHANNEL_PATH_ID            string                  `json:"channel-path-id,omitempty"`
}

/**
* Sample
* {
* 	"name":
*   "description":
*   "port-description":
*   "maximum-transmission-unit-size":
* }
* @return *   "object-uri":"/api/adapters/542b9406-d033-11e5-9f39-020000000338"
 */
type HipersocketPayload struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	PortDescription string `json:"port-description,omitempty"`
	MaxUnitSize     int    `json:"maximum-transmission-unit-size,omitempty"`
}

type HipersocketCreateResponse struct {
	URI string `json:"object-uri"`
}

//////////////////////////////////////////////////
// Virtual Switches
//////////////////////////////////////////////////

type VirtualSwitchType string

const (
	VIRTUALSWITCH_TYPE_HIPERSOCKET VirtualSwitchType = "hipersockets"
	VIRTUALSWITCH_TYPE_OSD                           = "osd"
)

type VirtualSwitchesArray struct {
	VIRTUALSWITCHES []VirtualSwitch `json:"virtual-switches"`
}

type VirtualSwitch struct {
	URI  string            `json:"object-uri,omitempty"`
	Name string            `json:"name,omitempty"`
	Type VirtualSwitchType `json:"type,omitempty"`
}

type VirtualSwitchProperties struct {
	URI         string            `json:"object-uri,omitempty"`
	Name        string            `json:"name,omitempty"`
	Type        VirtualSwitchType `json:"type,omitempty"`
	ID          string            `json:"object-id,omitempty"`
	Parent      string            `json:"parent,omitempty"`
	Class       string            `json:"class,omitempty"` // "virtual-switch"
	Description string            `json:"description,omitempty"`
	AdapterURI  string            `json:"backing-adapter-uri,omitempty"`
	Port        int               `json:"port,omitempty"`
	VNicUris    []string          `json:"connected-vnic-uris,omitempty"`
}

//////////////////////////////////////////////////
// CPC
//////////////////////////////////////////////////

type CpcStatus string

const (
	CPC_STATUS_ACTIVE           CpcStatus = "active"
	CPC_STATUS_OPERATING                  = "operating"
	CPC_STATUS_NO_COMMUNICATING           = "not-communicating"
	CPC_STATUS_EXCEPTIONS                 = "exceptions"
	CPC_STATUS_STATUS_CHECK               = "status-check"
	CPC_STATUS_SERVICE                    = "service"
	CPC_STATUS_NOT_OPERATING              = "not-operating"
	CPC_STATUS_NO_POWER                   = "no-power"
	CPC_STATUS_SERVICE_REQUIRED           = "service-required"
	CPC_STATUS_DEGRADED                   = "degraded"
)

/**
* Sample:
* {
*    "dpm-enabled": true,
*    "has-unacceptable-status": true,
*    "name": "P0LXSMOZ",
*    "object-uri": "/api/cpcs/e8753ff5-8ea6-35d9-b047-83c2624ba8da",
*    "se-version": "2.13.1"
*    "status": "not-operating"
* }
 */
type CPC struct {
	URI                 string    `json:"object-uri,omitempty"`
	Name                string    `json:"name,omitempty"`
	Status              CpcStatus `json:"status,omitempty"`
	HasAcceptableStatus bool      `json:"has-unacceptable-status,omitempty"`
	DpmEnabled          bool      `json:"dpm-enabled,omitempty"`
	SeVersion           string    `json:"se-version,omitempty"`
}

type CpcsArray struct {
	CPCS []CPC `json:"cpcs"`
}

//////////////////////////////////////////////////
// JOB
//////////////////////////////////////////////////

type JobStatus string

const (
	JOB_STATUS_RUNNING        JobStatus = "running"
	JOB_STATUS_CANCEL_PENDING           = "cancel-pending"
	JOB_STATUS_CANCELED                 = "canceled"
	JOB_STATUS_COMPLETE                 = "complete"
)

type JobResults struct {
	Message string `json:"message,omitempty"`
}

type Job struct {
	URI           string     `json:"job-uri,omitempty"`
	Status        JobStatus  `json:"status,omitempty"`
	JobStatusCode int        `json:"job-status-code,omitempty"`
	JobReasonCode int        `json:"job-reason-code,omitempty"`
	JobResults    JobResults `json:"job-results,omitempty"`
}

//////////////////////////////////////////////////
// LPAR
//////////////////////////////////////////////////

type PartitionType string

const (
	PARTITION_TYPE_SSC   PartitionType = "ssc"
	PARTITION_TYPE_LINUX               = "linux"
	PARTITION_TYPE_ZVM                 = "zvm"
)

type PartitionStatus string

const (
	PARTITION_STATUS_NOT_ACTIVE   PartitionStatus = "communications-not-active"
	PARTITION_STATUS_STATUS_CHECK                 = "status-check"
	PARTITION_STATUS_STOPPED                      = "stopped"
	PARTITION_STATUS_TERMINATED                   = "terminated"
	PARTITION_STATUS_STARTING                     = "starting"
	PARTITION_STATUS_ACTIVE                       = "active"
	PARTITION_STATUS_STOPPING                     = "stopping"
	PARTITION_STATUS_DEGRADED                     = "degraded"
	PARTITION_STATUS_REV_ERR                      = "reservation-error"
	PARTITION_STATUS_PAUSED                       = "paused"
)

type PartitionProcessorMode string

const (
	PROCESSOR_MODE_DEDICATED PartitionProcessorMode = "dedicated"
	PROCESSOR_MODE_SHARED                           = "shared"
)

type PartitionBootDevice string

const (
	BOOT_DEVICE_STORAGE_ADAPTER PartitionBootDevice = "storage-adapter"
	BOOT_DEVICE_STORAGE_VOLUME                      = "storage-volume"
	BOOT_DEVICE_NETWORK_ADAPTER                     = "network-adapter"
	BOOT_DEVICE_FTP                                 = "ftp"
	BOOT_DEVICE_REMOVABLE_MEDIA                     = "removable-media"
	BOOT_DEVICE_ISO_IMAGE                           = "iso-image"
	BOOT_DEVICE_NONE                                = "none" // default
)

type PartionBootRemovableMediaType string

const (
	BOOT_REMOVABLE_MEDIA_CDROM PartionBootRemovableMediaType = "cdrom"
	BOOT_REMOVABLE_MEDIA_USB                                 = "usb"
)

type SscBootSelection string

const (
	SSC_BOOT_SELECTION_INSTALLER SscBootSelection = "installer"
	SSC_BOOT_SELECTION_APPLIANCE                  = "appliance"
)

type LPAR struct {
	URI    string          `json:"object-uri,omitempty"`
	Name   string          `json:"name,omitempty"`
	Status PartitionStatus `json:"status,omitempty"`
	Type   PartitionType   `json:"type,omitempty"`
}

type LPARsArray struct {
	LPARS []LPAR `json:"partitions"`
}

type PartitionFeatureInfo struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State       bool   `json:"state,omitempty"`
}

type LparProperties struct {
	URI                              string                        `json:"object-uri,omitempty"`
	CpcURI                           string                        `json:"parent,omitempty"`
	Class                            string                        `json:"class,omitempty"`
	Name                             string                        `json:"name,omitempty"`
	Description                      string                        `json:"description,omitempty"`
	Status                           PartitionStatus               `json:"status,omitempty"`
	Type                             PartitionType                 `json:"type,omitempty"`
	ShortName                        string                        `json:"short-name,omitempty"`
	ID                               string                        `json:"partition-id,omitempty"`
	AutoGenerateID                   bool                          `json:"autogenerate-partition-id,omitempty"`
	OsName                           string                        `json:"os-name,omitempty"`
	OsType                           string                        `json:"os-type,omitempty"`
	OsVersion                        string                        `json:"os-version,omitempty"`
	ReserveResources                 bool                          `json:"reserve-resources,omitempty"`
	DegradedAdapters                 []string                      `json:"degraded-adapters,omitempty"`
	ProcessorMode                    PartitionProcessorMode        `json:"processor-mode,omitempty"`
	CpProcessors                     int                           `json:"cp-processors,omitempty"`
	IflProcessors                    int                           `json:"ifl-processors,omitempty"`
	IflAbsoluteProcessorCapping      bool                          `json:"ifl-absolute-processor-capping,omitempty"`
	CpAbsoluteProcessorCapping       bool                          `json:"cp-absolute-processor-capping,omitempty"`
	IflAbsoluteProcessorCappingValue float64                       `json:"ifl-absolute-processor-capping-value,omitempty"`
	CpAbsoluteProcessorCappingValue  float64                       `json:"cp-absolute-processor-capping-value,omitempty"`
	IflProcessingWeightCapped        bool                          `json:"ifl-processing-weight-capped,omitempty"`
	CpProcessingWeightCapped         bool                          `json:"cp-processing-weight-capped,omitempty"`
	MinimumIflProcessingWeight       int                           `json:"minimum-ifl-processing-weight,omitempty"`
	MinimumCpProcessingWeight        int                           `json:"minimum-cp-processing-weight,omitempty"`
	InitialIflProcessingWeight       int                           `json:"initial-ifl-processing-weight,omitempty"`
	InitialCpProcessingWeight        int                           `json:"initial-cp-processing-weight,omitempty"`
	CurrentIflProcessingWeight       int                           `json:"current-ifl-processing-weight,omitempty"`
	CurrentCpProcessingWeight        int                           `json:"current-cp-processing-weight,omitempty"`
	MaximumIflProcessingWeight       int                           `json:"maximum-ifl-processing-weight,omitempty"`
	MaximumCpProcessingWeight        int                           `json:"maximum-cp-processing-weight,omitempty"`
	ProcessorManagementEnabled       bool                          `json:"processor-management-enabled,omitempty"`
	InitialMemory                    int                           `json:"initial-memory,omitempty"`
	ReservedMemory                   int                           `json:"reserved-memory,omitempty"`
	MaximumMemory                    int                           `json:"maximum-memory,omitempty"`
	AutoStart                        bool                          `json:"auto-start,omitempty"`
	BootDevice                       PartitionBootDevice           `json:"boot-device,omitempty"`
	BootNetworkDevice                string                        `json:"boot-network-device,omitempty"`
	BootFtpHost                      string                        `json:"boot-ftp-host,omitempty"`
	BootFtpUsername                  string                        `json:"boot-ftp-username,omitempty"`
	BootFtpPassword                  string                        `json:"boot-ftp-password,omitempty"`
	BootFtpInsfile                   string                        `json:"boot-ftp-insfile,omitempty"`
	BootRemovableMedia               string                        `json:"boot-removable-media,omitempty"`
	BootRemovableMediaType           PartionBootRemovableMediaType `json:"boot-removable-media-type,omitempty"`
	BootTimeout                      int                           `json:"boot-timeout,omitempty"`
	BootStorageDevice                string                        `json:"boot-storage-device,omitempty"`
	BootStorageVolume                string                        `json:"boot-storage-volume,omitempty"`
	BootLogicalUnitNumber            string                        `json:"boot-logical-unit-number,omitempty"`
	BootWorldWidePortName            string                        `json:"boot-world-wide-port-name,omitempty"`
	BootConfigurationSelector        int                           `json:"boot-configuration-selector,omitempty"`
	BootRecordLba                    string                        `json:"boot-record-lba,omitempty"`
	BootLoadParameters               string                        `json:"boot-load-parameters,omitempty"`
	BootOsSpecificParameters         string                        `json:"boot-os-specific-parameters,omitempty"`
	BootIsoImageName                 string                        `json:"boot-iso-image-name,omitempty"`
	BootIsoInsFile                   string                        `json:"boot-iso-ins-file,omitempty"`
	AccessGlobalPerformanceData      bool                          `json:"access-global-performance-data,omitempty"`
	PermitCrossPartitionCommands     bool                          `json:"permit-cross-partition-commands,omitempty"`
	AccessBasicCounterSet            bool                          `json:"access-basic-counter-set,omitempty"`
	AccessProblemStateCounterSet     bool                          `json:"access-problem-state-counter-set,omitempty"`
	AccessCryptoActivityCounterSet   bool                          `json:"access-crypto-activity-counter-set,omitempty"`
	AccessExtendedCounterSet         bool                          `json:"access-extended-counter-set,omitempty"`
	AccessCoprocessorGroupSet        bool                          `json:"access-coprocessor-group-set,omitempty"`
	AccessBasicSampling              bool                          `json:"access-basic-sampling,omitempty"`
	AccessDiagnosticSampling         bool                          `json:"access-diagnostic-sampling,omitempty"`
	PermitDesKeyImportFunctions      bool                          `json:"permit-des-key-import-functions,omitempty"`
	PermitAesKeyImportFunctions      bool                          `json:"permit-aes-key-import-functions,omitempty"`
	ThreadsPerProcessor              int                           `json:"threads-per-processor,omitempty"`
	VirtualFunctionUris              []string                      `json:"virtual-function-uris,omitempty"`
	NicUris                          []string                      `json:"nic-uris,omitempty"`
	HbaUris                          []string                      `json:"hba-uris,omitempty"`
	StorageGroupURIs                 []string                      `json:"storage-group-uris,omitempty"`
	CryptoConfiguration              CryptoConfig                  `json:"crypto-configuration,omitempty"`
	SscHostName                      string                        `json:"ssc-host-name,omitempty"`
	SscBootSelection                 SscBootSelection              `json:"ssc-boot-selection,omitempty"`
	SscIpv4Gateway                   string                        `json:"ssc-ipv4-gateway,omitempty"`
	SscIpv6Gateway                   string                        `json:"ssc-ipv6-gateway,omitempty"`
	SscDnsServers                    []string                      `json:"ssc-dns-servers,omitempty"`
	SscMasterUserid                  string                        `json:"ssc-master-userid,omitempty"`
	SscMasterPw                      string                        `json:"ssc-master-pw,omitempty"`
	AvailableFeaturesList            []PartitionFeatureInfo        `json:"available-features-list,omitempty"`
	Secureboot                       bool                          `json:"secure-boot,omitempty"`
}

type StartStopLparResponse struct {
	URI     string `json:"job-uri"`
	Message string `json:"message"`
}

//////////////////////////////////////////////////
// NIC
//////////////////////////////////////////////////

type NicType string

const (
	NIC_TYPE_ROCE NicType = "roce"
	NIC_TYPE_IQD          = "iqd"
	NIC_TYPE_OSD          = "osd"
)

type SscIpAddressType string

const (
	SSC_IP_TYPE_IPV4      SscIpAddressType = "ipv4"
	SSC_IP_TYPE_IPV6                       = "ipv6"
	SSC_IP_TYPE_LINKLOCAL                  = "linklocal"
	SSC_IP_TYPE_DHCP                       = "dhcp"
)

type VlanType string

const (
	VLAN_TYPE_ENFORCED VlanType = "enforced"
)

type NIC struct {
	ID     string `json:"element-id,omitempty"`
	URI    string `json:"element-uri,omitempty"`
	Parent string `json:"parent,omitempty"`
	Class  string `json:"class,omitempty"`
	/* below are payloads when create a new Nic */
	Name                  string           `json:"name,omitempty"`
	Description           string           `json:"description,omitempty"`
	DeviceNumber          string           `json:"device-number,omitempty"`
	NetworkAdapterPortURI string           `json:"network-adapter-port-uri,omitempty"`
	VirtualSwitchUriType  string           `json:"virtual-switch-uri-type,omitempty"`
	VirtualSwitchURI      string           `json:"virtual-switch-uri,omitempty"`
	Type                  NicType          `json:"type,omitempty"`
	SscManagmentNIC       bool             `json:"ssc-management-nic,omitempty"`
	SscIpAddressType      SscIpAddressType `json:"ssc-ip-address-type,omitempty"`
	SscIpAddress          string           `json:"ssc-ip-address,omitempty"`
	VlanID                int              `json:"vlan-id,omitempty"`
	MacAddress            string           `json:"mac-address,omitempty"`
	SscMaskPrefix         string           `json:"ssc-mask-prefix,omitempty"`
	VlanType              VlanType         `json:"vlan-type,omitempty"`
}

type NicCreateResponse struct {
	URI string `json:"element-uri"`
}

//////////////////////////////////////////////////
// Storage Groups
//////////////////////////////////////////////////

type StorageGroupOperation string

const (
	STORAGE_GROUP_MODIFY  StorageGroupOperation = "modify"
	STORAGE_VOLUME_CREATE StorageGroupOperation = "create"
)

type StorageGroupState string

const (
	STORAGE_GROUP_COMPLETE                StorageGroupState = "complete"
	STORAGE_GROUP_PENDING                                   = "pending"
	STORAGE_GROUP_INCOMPLETE                                = "incomplete"
	STORAGE_GROUP_PENDING_WITH_MISMATCHES                   = "pending-with-mismatches"
	STORAGE_GROUP_OVERPROVISIONED                           = "overprovisioned"
	STORAGE_GROUP_CHECKING_MIGRATION                        = "checking-migration"
	STORAGE_GROUP_CONFIGURATION_ERROR                       = "configuration-error"
)

type WWPNStatus string

const (
	WWPN_STATUS_VALIDATED     WWPNStatus = "validated"
	WWPN_STATUS_NOT_VALIDATED            = "not-validated"
	WWPN_STATUS_UNKNOWN                  = "unknown"
	WWPN_STATUS_INCOMPLETE               = "incomplete"
)

type StorageGroupArray struct {
	STORAGEGROUPS []StorageGroup `json:"storage-groups"`
}

type StorageVolumeArray struct {
	STORAGEVOLUMES []StorageVolume `json:"storage-volumes"`
}

type StorageGroupPayload struct {
	StorageGroupURI  string `json:"storage-group-uri"`
	StorageGroupName string `json:"storage-group-name,omitempty"`
	PartitionName    string `json:"partition-name,omitempty"`
}
type AttachStorageGroupPayload struct {
	StorageGroupURI string `json:"storage-group-uri"`
}

// Storage group class specific properties
type StorageGroupProperties struct {
	Class                      string            `json:"class,omitempty"`
	Connectivity               int               `json:"connectivity,omitempty"`
	ActiveConnectivity         int               `json:"active-connectivity,omitempty"`
	CpcURI                     string            `json:"cpc-uri,omitempty"`
	CandidatePortURIs          []string          `json:"candidate-adapter-port-uris,omitempty"`
	Description                string            `json:"description,omitempty"`
	DirectConnectionCount      int               `json:"direct-connection-count,omitempty"`
	FulfillmentState           StorageGroupState `json:"fulfillment-state,omitempty"`
	MaxPartitions              int               `json:"max-partitions,omitempty"`
	ActiveMaxPartitions        int               `json:"active-max-partitions,omitempty"`
	Name                       string            `json:"name,omitempty"`
	ObjectID                   string            `json:"object-id,omitempty"`
	ObjectURI                  string            `json:"object-uri,omitempty"`
	Parent                     string            `json:"parent,omitempty"`
	Shared                     bool              `json:"shared,omitempty"`
	StorageVolumes             []StorageVolume   `json:"storage-volumes,omitempty"`
	StorageVolumesURIs         []string          `json:"storage-volume-uris,omitempty"`
	UnAssignedWWPNs            []WWPN            `json:"unassigned-world-wide-port-names,omitempty"`
	VirtualStorageResourceURIs []string          `json:"virtual-storage-resource-uris,omitempty"`
	Type                       string            `json:"type,omitempty"`
}

type CreateStorageGroupProperties struct {
	CpcURI                string          `json:"cpc-uri,omitempty"`
	TemplateURI           string          `json:"template-uri,omitempty"`
	Name                  string          `json:"name,omitempty"`
	Description           string          `json:"description,omitempty"`
	Type                  string          `json:"type,omitempty"`
	Shared                bool            `json:"shared"`
	Connectivity          int             `json:"connectivity,omitempty"`
	MaxPartitions         int             `json:"max-partitions,omitempty"`
	DirectConnectionCount int             `json:"direct-connection-count,omitempty"`
	StorageVolumes        []StorageVolume `json:"storage-volumes,omitempty"`
	EmailToAddress        []string        `json:"email-to-address,omitempty"`
	EmailCcAddress        []string        `json:"email-cc-address,omitempty"`
	EmailInsert           string          `json:"email-insert,omitempty"`
}

type WWPN struct {
	Name   string     `json:"world-wide-port-name,omitempty"`
	Status WWPNStatus `json:"status,omitempty"`
}

// Storage group object
type StorageGroup struct {
	ObjectID         string            `json:"object-id,omitempty"`
	Parent           string            `json:"parent,omitempty"`
	CpcURI           string            `json:"cpc-uri,omitempty"`
	Description      string            `json:"description,omitempty"`
	FulfillmentState StorageGroupState `json:"fulfillment-state,omitempty"`
	Name             string            `json:"name,omitempty"`
	ObjectURI        string            `json:"object-uri,omitempty"`
	Type             string            `json:"type,omitempty"`
	StorageVolumes   []StorageVolume   `json:"storage-volumes,omitempty"`
}

type StorageGroupPartitions struct {
	GetStorageGroups []LparProperties `json:"partitions,omitempty"`
}

type StorageGroupModel string

const (
	STORAGE_GROUP_MODEL_1   StorageGroupModel = "1"
	STORAGE_GROUP_MODEL_2                     = "2"
	STORAGE_GROUP_MODEL_3                     = "3"
	STORAGE_GROUP_MODEL_9                     = "9"
	STORAGE_GROUP_MODEL_27                    = "27"
	STORAGE_GROUP_MODEL_54                    = "54"
	STORAGE_GROUP_MODEL_EAV                   = "EAV"
)

type EckdVolumetype string

const (
	STORAGE_VOLUME_ECKDTYPE_BASE  EckdVolumetype = "base"
	STORAGE_VOLUME_ECKDTYPE_ALIAS EckdVolumetype = "alias"
)

type StorageVolume struct {
	Operation        StorageGroupOperation `json:"operation,omitempty"`
	Class            string                `json:"class,omitempty"`
	Parent           string                `json:"parent,omitempty"`
	URI              string                `json:"element-uri,omitempty"`
	Name             string                `json:"name,omitempty"`
	Description      string                `json:"description,omitempty"`
	Size             float64               `json:"size,omitempty"`
	ActiveSize       float64               `json:"active-size,omitempty"`
	UUID             string                `json:"uuid,omitempty"`
	Model            string                `json:"model,omitempty"`
	ActiveModel      string                `json:"acitve-model,omitempty"`
	Usage            StorageVolumeUsage    `json:"usage,omitempty"`
	FulfillmentState StorageGroupState     `json:"fulfillment-state,omitempty"`
	Cylinders        int                   `json:"cylinders,omitempty"`
	DeviceNumber     string                `json:"device-number,omitempty"`
	ControlUnitURI   string                `json:"control-unit-uri,omitempty"`
	EckdType         string                `json:"eckd-type,omitempty"`
	UnitAddress      string                `json:"unit-address,omitempty"`
	Paths            []VolumePath          `json:"paths,omitempty"`
	AdapterURI       string                `json:"adapter-uri,omitempty"`
	SerialNumber     string                `json:"serial-number,omitempty"`
	FID              string                `json:"fid,omitempty"`
}

type StorageVolumeUsage string

const (
	BOOT_USAGE StorageVolumeUsage = "boot"
	DATA_USAGE StorageVolumeUsage = "data"
)

type VolumePath struct {
	PartitionURI      string `json:"partition-uri,omitempty"`
	DeviceNumber      string `json:"device-number,omitempty"`
	TargetWWPN        string `json:"target-world-wide-port-name,omitempty"`
	LogicalUnitNumber string `json:"logical-unit-number,omitempty"`
}
type StorageGroupCreateResponse struct {
	URI       []string                 `json:"element-uris,omitempty"`
	ObjectURI string                   `json:"object-uri,omitempty"`
	SvPaths   []StorageGroupVolumePath `json:"sv-paths,omitempty"`
}
type StorageGroupVolumePath struct {
	URI   string       `json:"element-uris,omitempty"`
	Paths []VolumePath `json:"paths,omitempty"`
}

// ASCIIConsolePayload
type AsciiConsoleURIPayload struct {
	ForceTakeover bool `json:"force-takeover,omitempty"`
}

type AsciiConsoleURIResponse struct {
	URI       string `json:"websocket-uri"`
	SessionID string `json:"session-id"`
}

// SSC crypto DomainInfo added as a substructure of CryptoConfig structure
type DomainInfo struct {
	DomainIdx  int    `json:"domain-index"`
	AccessMode string `json:"access-mode"`
}

// SSC crypto-configuration object properties added to Support Crypto Card working
type CryptoConfig struct {
	CryptoAdapterUris          []string     `json:"crypto-adapter-uris,omitempty"`
	CryptoDomainConfigurations []DomainInfo `json:"crypto-domain-configurations,omitempty"`
}
