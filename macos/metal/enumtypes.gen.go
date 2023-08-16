// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import "math"

// Options for specifying different kinds of instance types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureinstancedescriptortype?language=objc
type AccelerationStructureInstanceDescriptorType uint

const (
	AccelerationStructureInstanceDescriptorTypeDefault AccelerationStructureInstanceDescriptorType = 0
	AccelerationStructureInstanceDescriptorTypeMotion  AccelerationStructureInstanceDescriptorType = 2
	AccelerationStructureInstanceDescriptorTypeUserID  AccelerationStructureInstanceDescriptorType = 1
)

// Options for adjusting the behavior of an instanced acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureinstanceoptions?language=objc
type AccelerationStructureInstanceOptions uint32

const (
	AccelerationStructureInstanceOptionDisableTriangleCulling                     AccelerationStructureInstanceOptions = 1
	AccelerationStructureInstanceOptionNonOpaque                                  AccelerationStructureInstanceOptions = 8
	AccelerationStructureInstanceOptionNone                                       AccelerationStructureInstanceOptions = 0
	AccelerationStructureInstanceOptionOpaque                                     AccelerationStructureInstanceOptions = 4
	AccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise AccelerationStructureInstanceOptions = 2
)

// Options that describe which tasks you can perform on an acceleration structure and how the system performs those tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructureusage?language=objc
type AccelerationStructureUsage uint

const (
	AccelerationStructureUsageExtendedLimits  AccelerationStructureUsage = 4
	AccelerationStructureUsageNone            AccelerationStructureUsage = 0
	AccelerationStructureUsagePreferFastBuild AccelerationStructureUsage = 2
	AccelerationStructureUsageRefit           AccelerationStructureUsage = 1
)

// The values that determine the limits and capabilities of argument buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumentbufferstier?language=objc
type ArgumentBuffersTier uint

const (
	ArgumentBuffersTier1 ArgumentBuffersTier = 0
	ArgumentBuffersTier2 ArgumentBuffersTier = 1
)

// The resource type for an argument of a function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlargumenttype?language=objc
type ArgumentType uint

const (
	ArgumentTypeBuffer                         ArgumentType = 0
	ArgumentTypeImageblock                     ArgumentType = 17
	ArgumentTypeImageblockData                 ArgumentType = 16
	ArgumentTypeInstanceAccelerationStructure  ArgumentType = 26
	ArgumentTypeIntersectionFunctionTable      ArgumentType = 27
	ArgumentTypePrimitiveAccelerationStructure ArgumentType = 25
	ArgumentTypeSampler                        ArgumentType = 3
	ArgumentTypeTexture                        ArgumentType = 2
	ArgumentTypeThreadgroupMemory              ArgumentType = 1
	ArgumentTypeVisibleFunctionTable           ArgumentType = 24
)

// Values that specify the organization of function attribute data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlattributeformat?language=objc
type AttributeFormat uint

const (
	AttributeFormatChar                  AttributeFormat = 46
	AttributeFormatChar2                 AttributeFormat = 4
	AttributeFormatChar2Normalized       AttributeFormat = 10
	AttributeFormatChar3                 AttributeFormat = 5
	AttributeFormatChar3Normalized       AttributeFormat = 11
	AttributeFormatChar4                 AttributeFormat = 6
	AttributeFormatChar4Normalized       AttributeFormat = 12
	AttributeFormatCharNormalized        AttributeFormat = 48
	AttributeFormatFloat                 AttributeFormat = 28
	AttributeFormatFloat2                AttributeFormat = 29
	AttributeFormatFloat3                AttributeFormat = 30
	AttributeFormatFloat4                AttributeFormat = 31
	AttributeFormatHalf                  AttributeFormat = 53
	AttributeFormatHalf2                 AttributeFormat = 25
	AttributeFormatHalf3                 AttributeFormat = 26
	AttributeFormatHalf4                 AttributeFormat = 27
	AttributeFormatInt                   AttributeFormat = 32
	AttributeFormatInt1010102Normalized  AttributeFormat = 40
	AttributeFormatInt2                  AttributeFormat = 33
	AttributeFormatInt3                  AttributeFormat = 34
	AttributeFormatInt4                  AttributeFormat = 35
	AttributeFormatInvalid               AttributeFormat = 0
	AttributeFormatShort                 AttributeFormat = 50
	AttributeFormatShort2                AttributeFormat = 16
	AttributeFormatShort2Normalized      AttributeFormat = 22
	AttributeFormatShort3                AttributeFormat = 17
	AttributeFormatShort3Normalized      AttributeFormat = 23
	AttributeFormatShort4                AttributeFormat = 18
	AttributeFormatShort4Normalized      AttributeFormat = 24
	AttributeFormatShortNormalized       AttributeFormat = 52
	AttributeFormatUChar                 AttributeFormat = 45
	AttributeFormatUChar2                AttributeFormat = 1
	AttributeFormatUChar2Normalized      AttributeFormat = 7
	AttributeFormatUChar3                AttributeFormat = 2
	AttributeFormatUChar3Normalized      AttributeFormat = 8
	AttributeFormatUChar4                AttributeFormat = 3
	AttributeFormatUChar4Normalized      AttributeFormat = 9
	AttributeFormatUChar4Normalized_BGRA AttributeFormat = 42
	AttributeFormatUCharNormalized       AttributeFormat = 47
	AttributeFormatUInt                  AttributeFormat = 36
	AttributeFormatUInt1010102Normalized AttributeFormat = 41
	AttributeFormatUInt2                 AttributeFormat = 37
	AttributeFormatUInt3                 AttributeFormat = 38
	AttributeFormatUInt4                 AttributeFormat = 39
	AttributeFormatUShort                AttributeFormat = 49
	AttributeFormatUShort2               AttributeFormat = 13
	AttributeFormatUShort2Normalized     AttributeFormat = 19
	AttributeFormatUShort3               AttributeFormat = 14
	AttributeFormatUShort3Normalized     AttributeFormat = 20
	AttributeFormatUShort4               AttributeFormat = 15
	AttributeFormatUShort4Normalized     AttributeFormat = 21
	AttributeFormatUShortNormalized      AttributeFormat = 51
)

// Describes the types of resources that a barrier operates on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbarrierscope?language=objc
type BarrierScope uint

const (
	BarrierScopeBuffers       BarrierScope = 1
	BarrierScopeRenderTargets BarrierScope = 4
	BarrierScopeTextures      BarrierScope = 2
)

// Error codes when creating binary archives of compiled shader code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbinaryarchiveerror?language=objc
type BinaryArchiveError uint

const (
	BinaryArchiveErrorCompilationFailure BinaryArchiveError = 3
	BinaryArchiveErrorInvalidFile        BinaryArchiveError = 1
	BinaryArchiveErrorNone               BinaryArchiveError = 0
	BinaryArchiveErrorUnexpectedElement  BinaryArchiveError = 2
)

// The source and destination blend factors are often needed to complete specification of a blend operation. In most cases, the blend factor for both RGB values (F(rgb)) and alpha values (F(a)) are similar to one another, but in some cases, such as MTLBlendFactorSourceAlphaSaturated, the blend factor is slightly different. Four blend factors (MTLBlendFactorBlendColor, MTLBlendFactorOneMinusBlendColor, MTLBlendFactorBlendAlpha, and MTLBlendFactorOneMinusBlendAlpha) refer to a constant blend color value that is set by the setBlendColorRed:green:blue:alpha: method of MTLRenderCommandEncoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblendfactor?language=objc
type BlendFactor uint

const (
	BlendFactorBlendAlpha               BlendFactor = 13
	BlendFactorBlendColor               BlendFactor = 11
	BlendFactorDestinationAlpha         BlendFactor = 8
	BlendFactorDestinationColor         BlendFactor = 6
	BlendFactorOne                      BlendFactor = 1
	BlendFactorOneMinusBlendAlpha       BlendFactor = 14
	BlendFactorOneMinusBlendColor       BlendFactor = 12
	BlendFactorOneMinusDestinationAlpha BlendFactor = 9
	BlendFactorOneMinusDestinationColor BlendFactor = 7
	BlendFactorOneMinusSource1Alpha     BlendFactor = 18
	BlendFactorOneMinusSource1Color     BlendFactor = 16
	BlendFactorOneMinusSourceAlpha      BlendFactor = 5
	BlendFactorOneMinusSourceColor      BlendFactor = 3
	BlendFactorSource1Alpha             BlendFactor = 17
	BlendFactorSource1Color             BlendFactor = 15
	BlendFactorSourceAlpha              BlendFactor = 4
	BlendFactorSourceAlphaSaturated     BlendFactor = 10
	BlendFactorSourceColor              BlendFactor = 2
	BlendFactorZero                     BlendFactor = 0
)

// For every pixel, MTLBlendOperation determines how to combine and weight the source fragment values with the destination values. Some blend operations multiply the source values by a source blend factor (SBF), multiply the destination values by a destination blend factor (DBF), and then combine the results using addition or subtraction. Other blend operations use either a minimum or maximum function to determine the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblendoperation?language=objc
type BlendOperation uint

const (
	BlendOperationAdd             BlendOperation = 0
	BlendOperationMax             BlendOperation = 4
	BlendOperationMin             BlendOperation = 3
	BlendOperationReverseSubtract BlendOperation = 2
	BlendOperationSubtract        BlendOperation = 1
)

// The options that enable behavior for some blit operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitoption?language=objc
type BlitOption uint

const (
	BlitOptionDepthFromDepthStencil   BlitOption = 1
	BlitOptionNone                    BlitOption = 0
	BlitOptionRowLinearPVRTC          BlitOption = 4
	BlitOptionStencilFromDepthStencil BlitOption = 2
)

// Options for the CPU cache mode that define the CPU mapping of the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcpucachemode?language=objc
type CPUCacheMode uint

const (
	CPUCacheModeDefaultCache  CPUCacheMode = 0
	CPUCacheModeWriteCombined CPUCacheMode = 1
)

// The kinds of destinations for captured command data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcapturedestination?language=objc
type CaptureDestination int

const (
	CaptureDestinationDeveloperTools   CaptureDestination = 1
	CaptureDestinationGPUTraceDocument CaptureDestination = 2
)

// Errors returned by capture sessions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcaptureerror?language=objc
type CaptureError int

const (
	CaptureErrorAlreadyCapturing  CaptureError = 2
	CaptureErrorInvalidDescriptor CaptureError = 3
	CaptureErrorNotSupported      CaptureError = 1
)

// Values used to specify a mask to permit or restrict writing to color channels of a color value. The values [metal/mtlcolorwritemask/mtlcolorwritemaskred], [metal/mtlcolorwritemask/mtlcolorwritemaskgreen], [metal/mtlcolorwritemask/mtlcolorwritemaskblue], and [metal/mtlcolorwritemask/mtlcolorwritemaskalpha] select one color channel each, and they can be bitwise combined. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcolorwritemask?language=objc
type ColorWriteMask uint

const (
	ColorWriteMaskAll   ColorWriteMask = 15
	ColorWriteMaskAlpha ColorWriteMask = 1
	ColorWriteMaskBlue  ColorWriteMask = 2
	ColorWriteMaskGreen ColorWriteMask = 4
	ColorWriteMaskNone  ColorWriteMask = 0
	ColorWriteMaskRed   ColorWriteMask = 8
)

// Error codes that indicate why a command buffer is unable to finish its execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffererror?language=objc
type CommandBufferError uint

const (
	CommandBufferErrorAccessRevoked   CommandBufferError = 4
	CommandBufferErrorBlacklisted     CommandBufferError = 4
	CommandBufferErrorDeviceRemoved   CommandBufferError = 11
	CommandBufferErrorInternal        CommandBufferError = 1
	CommandBufferErrorInvalidResource CommandBufferError = 9
	CommandBufferErrorMemoryless      CommandBufferError = 10
	CommandBufferErrorNone            CommandBufferError = 0
	CommandBufferErrorNotPermitted    CommandBufferError = 7
	CommandBufferErrorOutOfMemory     CommandBufferError = 8
	CommandBufferErrorPageFault       CommandBufferError = 3
	CommandBufferErrorStackOverflow   CommandBufferError = 12
	CommandBufferErrorTimeout         CommandBufferError = 2
)

// Options for reporting errors from a command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffererroroption?language=objc
type CommandBufferErrorOption uint

const (
	CommandBufferErrorOptionEncoderExecutionStatus CommandBufferErrorOption = 1
	CommandBufferErrorOptionNone                   CommandBufferErrorOption = 0
)

// The discrete states for a command buffer that represent its life cycle stages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferstatus?language=objc
type CommandBufferStatus uint

const (
	CommandBufferStatusCommitted   CommandBufferStatus = 2
	CommandBufferStatusCompleted   CommandBufferStatus = 4
	CommandBufferStatusEnqueued    CommandBufferStatus = 1
	CommandBufferStatusError       CommandBufferStatus = 5
	CommandBufferStatusNotEnqueued CommandBufferStatus = 0
	CommandBufferStatusScheduled   CommandBufferStatus = 3
)

// Possible error conditions for the command encoder’s commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandencodererrorstate?language=objc
type CommandEncoderErrorState int

const (
	CommandEncoderErrorStateAffected  CommandEncoderErrorState = 2
	CommandEncoderErrorStateCompleted CommandEncoderErrorState = 1
	CommandEncoderErrorStateFaulted   CommandEncoderErrorState = 4
	CommandEncoderErrorStatePending   CommandEncoderErrorState = 3
	CommandEncoderErrorStateUnknown   CommandEncoderErrorState = 0
)

// The name of a specific counter that can appear in a GPU device’s counter sets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommoncounter?language=objc
type CommonCounter string

const (
	CommonCounterClipperInvocations                CommonCounter = "ClipperInvocations"
	CommonCounterClipperPrimitivesOut              CommonCounter = "ClipperPrimitivesOut"
	CommonCounterComputeKernelInvocations          CommonCounter = "KernelInvocations"
	CommonCounterFragmentCycles                    CommonCounter = "FragmentCycles"
	CommonCounterFragmentInvocations               CommonCounter = "FragmentInvocations"
	CommonCounterFragmentsPassed                   CommonCounter = "FragmentsPassed"
	CommonCounterPostTessellationVertexCycles      CommonCounter = "PostTessellationCycles"
	CommonCounterPostTessellationVertexInvocations CommonCounter = "PostTessellationVertexInvocations"
	CommonCounterRenderTargetWriteCycles           CommonCounter = "RenderTargetWriteCycles"
	CommonCounterTessellationCycles                CommonCounter = "TessellationCycles"
	CommonCounterTessellationInputPatches          CommonCounter = "TessellationInputPatches"
	CommonCounterTimestamp                         CommonCounter = "GPUTimestamp"
	CommonCounterTotalCycles                       CommonCounter = "TotalCycles"
	CommonCounterVertexCycles                      CommonCounter = "VertexCycles"
	CommonCounterVertexInvocations                 CommonCounter = "VertexInvocations"
)

// The name of a specific counter set that a GPU device can support. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommoncounterset?language=objc
type CommonCounterSet string

const (
	CommonCounterSetStageUtilization CommonCounterSet = "stageutilization"
	CommonCounterSetStatistic        CommonCounterSet = "statistic"
	CommonCounterSetTimestamp        CommonCounterSet = "timestamp"
)

// Options used to specify how a sample compare operation should be performed on a depth texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomparefunction?language=objc
type CompareFunction uint

const (
	CompareFunctionAlways       CompareFunction = 7
	CompareFunctionEqual        CompareFunction = 2
	CompareFunctionGreater      CompareFunction = 4
	CompareFunctionGreaterEqual CompareFunction = 6
	CompareFunctionLess         CompareFunction = 1
	CompareFunctionLessEqual    CompareFunction = 3
	CompareFunctionNever        CompareFunction = 0
	CompareFunctionNotEqual     CompareFunction = 5
)

// The underlying error code type that indicates why a GPU driver can’t create a counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplebuffererror?language=objc
type CounterSampleBufferError int

const (
	CounterSampleBufferErrorInternal    CounterSampleBufferError = 2
	CounterSampleBufferErrorInvalid     CounterSampleBufferError = 1
	CounterSampleBufferErrorOutOfMemory CounterSampleBufferError = 0
)

// Options for different times when you can sample GPU counters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcountersamplingpoint?language=objc
type CounterSamplingPoint uint

const (
	CounterSamplingPointAtBlitBoundary         CounterSamplingPoint = 4
	CounterSamplingPointAtDispatchBoundary     CounterSamplingPoint = 2
	CounterSamplingPointAtDrawBoundary         CounterSamplingPoint = 1
	CounterSamplingPointAtStageBoundary        CounterSamplingPoint = 0
	CounterSamplingPointAtTileDispatchBoundary CounterSamplingPoint = 3
)

// The mode that determines whether to perform culling and which type of primitive to cull. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcullmode?language=objc
type CullMode uint

const (
	CullModeBack  CullMode = 2
	CullModeFront CullMode = 1
	CullModeNone  CullMode = 0
)

// The types of GPU functions, including shaders and compute kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldatatype?language=objc
type DataType uint

const (
	DataTypeArray                          DataType = 2
	DataTypeBool                           DataType = 53
	DataTypeBool2                          DataType = 54
	DataTypeBool3                          DataType = 55
	DataTypeBool4                          DataType = 56
	DataTypeChar                           DataType = 45
	DataTypeChar2                          DataType = 46
	DataTypeChar3                          DataType = 47
	DataTypeChar4                          DataType = 48
	DataTypeComputePipeline                DataType = 79
	DataTypeFloat                          DataType = 3
	DataTypeFloat2                         DataType = 4
	DataTypeFloat2x2                       DataType = 7
	DataTypeFloat2x3                       DataType = 8
	DataTypeFloat2x4                       DataType = 9
	DataTypeFloat3                         DataType = 5
	DataTypeFloat3x2                       DataType = 10
	DataTypeFloat3x3                       DataType = 11
	DataTypeFloat3x4                       DataType = 12
	DataTypeFloat4                         DataType = 6
	DataTypeFloat4x2                       DataType = 13
	DataTypeFloat4x3                       DataType = 14
	DataTypeFloat4x4                       DataType = 15
	DataTypeHalf                           DataType = 16
	DataTypeHalf2                          DataType = 17
	DataTypeHalf2x2                        DataType = 20
	DataTypeHalf2x3                        DataType = 21
	DataTypeHalf2x4                        DataType = 22
	DataTypeHalf3                          DataType = 18
	DataTypeHalf3x2                        DataType = 23
	DataTypeHalf3x3                        DataType = 24
	DataTypeHalf3x4                        DataType = 25
	DataTypeHalf4                          DataType = 19
	DataTypeHalf4x2                        DataType = 26
	DataTypeHalf4x3                        DataType = 27
	DataTypeHalf4x4                        DataType = 28
	DataTypeIndirectCommandBuffer          DataType = 80
	DataTypeInstanceAccelerationStructure  DataType = 118
	DataTypeInt                            DataType = 29
	DataTypeInt2                           DataType = 30
	DataTypeInt3                           DataType = 31
	DataTypeInt4                           DataType = 32
	DataTypeIntersectionFunctionTable      DataType = 116
	DataTypeLong                           DataType = 81
	DataTypeLong2                          DataType = 82
	DataTypeLong3                          DataType = 83
	DataTypeLong4                          DataType = 84
	DataTypeNone                           DataType = 0
	DataTypePointer                        DataType = 60
	DataTypePrimitiveAccelerationStructure DataType = 117
	DataTypeR16Snorm                       DataType = 65
	DataTypeR16Unorm                       DataType = 64
	DataTypeR8Snorm                        DataType = 63
	DataTypeR8Unorm                        DataType = 62
	DataTypeRG11B10Float                   DataType = 76
	DataTypeRG16Snorm                      DataType = 69
	DataTypeRG16Unorm                      DataType = 68
	DataTypeRG8Snorm                       DataType = 67
	DataTypeRG8Unorm                       DataType = 66
	DataTypeRGB10A2Unorm                   DataType = 75
	DataTypeRGB9E5Float                    DataType = 77
	DataTypeRGBA16Snorm                    DataType = 74
	DataTypeRGBA16Unorm                    DataType = 73
	DataTypeRGBA8Snorm                     DataType = 72
	DataTypeRGBA8Unorm                     DataType = 70
	DataTypeRGBA8Unorm_sRGB                DataType = 71
	DataTypeRenderPipeline                 DataType = 78
	DataTypeSampler                        DataType = 59
	DataTypeShort                          DataType = 37
	DataTypeShort2                         DataType = 38
	DataTypeShort3                         DataType = 39
	DataTypeShort4                         DataType = 40
	DataTypeStruct                         DataType = 1
	DataTypeTexture                        DataType = 58
	DataTypeUChar                          DataType = 49
	DataTypeUChar2                         DataType = 50
	DataTypeUChar3                         DataType = 51
	DataTypeUChar4                         DataType = 52
	DataTypeUInt                           DataType = 33
	DataTypeUInt2                          DataType = 34
	DataTypeUInt3                          DataType = 35
	DataTypeUInt4                          DataType = 36
	DataTypeULong                          DataType = 85
	DataTypeULong2                         DataType = 86
	DataTypeULong3                         DataType = 87
	DataTypeULong4                         DataType = 88
	DataTypeUShort                         DataType = 41
	DataTypeUShort2                        DataType = 42
	DataTypeUShort3                        DataType = 43
	DataTypeUShort4                        DataType = 44
	DataTypeVisibleFunctionTable           DataType = 115
)

// The mode that determines how to deal with fragments outside of the near or far planes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthclipmode?language=objc
type DepthClipMode uint

const (
	DepthClipModeClamp DepthClipMode = 1
	DepthClipModeClip  DepthClipMode = 0
)

// Indicates the location of the GPU relative to the system it’s connect to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevicelocation?language=objc
type DeviceLocation uint

const (
	DeviceLocationBuiltIn     DeviceLocation = 0
	DeviceLocationExternal    DeviceLocation = 2
	DeviceLocationSlot        DeviceLocation = 1
	DeviceLocationUnspecified DeviceLocation = math.MaxUint
)

// A notification that represents a change to a GPU device in the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevicenotificationname?language=objc
type DeviceNotificationName string

const (
	DeviceRemovalRequestedNotification DeviceNotificationName = "MTLDeviceRemovalRequested"
	DeviceWasAddedNotification         DeviceNotificationName = "MTLDeviceWasAdded"
	DeviceWasRemovedNotification       DeviceNotificationName = "MTLDeviceWasRemoved"
)

// Instructs a compute command encoder to run commands serially or in parallel during a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldispatchtype?language=objc
type DispatchType uint

const (
	DispatchTypeConcurrent DispatchType = 1
	DispatchTypeSerial     DispatchType = 0
)

// Error codes that Metal can generate when creating dynamic libraries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibraryerror?language=objc
type DynamicLibraryError uint

const (
	DynamicLibraryErrorCompilationFailure    DynamicLibraryError = 2
	DynamicLibraryErrorDependencyLoadFailure DynamicLibraryError = 4
	DynamicLibraryErrorInvalidFile           DynamicLibraryError = 1
	DynamicLibraryErrorNone                  DynamicLibraryError = 0
	DynamicLibraryErrorUnresolvedInstallName DynamicLibraryError = 3
	DynamicLibraryErrorUnsupported           DynamicLibraryError = 5
)

// The device feature sets that define specific platform, hardware, and software configurations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfeatureset?language=objc
type FeatureSet uint

const (
	FeatureSet_macOS_GPUFamily1_v1         FeatureSet = 10000
	FeatureSet_macOS_GPUFamily1_v2         FeatureSet = 10001
	FeatureSet_macOS_GPUFamily1_v3         FeatureSet = 10003
	FeatureSet_macOS_GPUFamily1_v4         FeatureSet = 10004
	FeatureSet_macOS_GPUFamily2_v1         FeatureSet = 10005
	FeatureSet_macOS_ReadWriteTextureTier2 FeatureSet = 10002
)

// Options for different kinds of function logs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionlogtype?language=objc
type FunctionLogType uint

const (
	FunctionLogTypeValidation FunctionLogType = 0
)

// Options that define how Metal creates the function object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionoptions?language=objc
type FunctionOptions uint

const (
	FunctionOptionCompileToBinary FunctionOptions = 1
	FunctionOptionNone            FunctionOptions = 0
)

// The type of a top-level Metal Shading Language (MSL) function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiontype?language=objc
type FunctionType uint

const (
	FunctionTypeFragment     FunctionType = 2
	FunctionTypeIntersection FunctionType = 6
	FunctionTypeKernel       FunctionType = 3
	FunctionTypeVertex       FunctionType = 1
	FunctionTypeVisible      FunctionType = 5
)

// Represents the functionality for families of GPUs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlgpufamily?language=objc
type GPUFamily int

const (
	GPUFamilyApple1       GPUFamily = 1001
	GPUFamilyApple2       GPUFamily = 1002
	GPUFamilyApple3       GPUFamily = 1003
	GPUFamilyApple4       GPUFamily = 1004
	GPUFamilyApple5       GPUFamily = 1005
	GPUFamilyApple6       GPUFamily = 1006
	GPUFamilyApple7       GPUFamily = 1007
	GPUFamilyCommon1      GPUFamily = 3001
	GPUFamilyCommon2      GPUFamily = 3002
	GPUFamilyCommon3      GPUFamily = 3003
	GPUFamilyMac1         GPUFamily = 2001
	GPUFamilyMac2         GPUFamily = 2002
	GPUFamilyMacCatalyst1 GPUFamily = 4001
	GPUFamilyMacCatalyst2 GPUFamily = 4002
)

// The options you use to specify the hazard tracking mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlhazardtrackingmode?language=objc
type HazardTrackingMode uint

const (
	HazardTrackingModeDefault   HazardTrackingMode = 0
	HazardTrackingModeTracked   HazardTrackingMode = 2
	HazardTrackingModeUntracked HazardTrackingMode = 1
)

// The options you use to choose the heap type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlheaptype?language=objc
type HeapType int

const (
	HeapTypeAutomatic HeapType = 0
	HeapTypePlacement HeapType = 1
	HeapTypeSparse    HeapType = 2
)

// The index type for an index buffer that references vertices of geometric primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindextype?language=objc
type IndexType uint

const (
	IndexTypeUInt16 IndexType = 0
	IndexTypeUInt32 IndexType = 1
)

// The types of commands that you can encode into the indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcommandtype?language=objc
type IndirectCommandType uint

const (
	IndirectCommandTypeConcurrentDispatch        IndirectCommandType = 32
	IndirectCommandTypeConcurrentDispatchThreads IndirectCommandType = 64
	IndirectCommandTypeDraw                      IndirectCommandType = 1
	IndirectCommandTypeDrawIndexed               IndirectCommandType = 2
	IndirectCommandTypeDrawIndexedPatches        IndirectCommandType = 8
	IndirectCommandTypeDrawPatches               IndirectCommandType = 4
)

// Constants for specifying different types of custom intersection functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlintersectionfunctionsignature?language=objc
type IntersectionFunctionSignature uint

const (
	IntersectionFunctionSignatureExtendedLimits  IntersectionFunctionSignature = 32
	IntersectionFunctionSignatureInstanceMotion  IntersectionFunctionSignature = 8
	IntersectionFunctionSignatureInstancing      IntersectionFunctionSignature = 1
	IntersectionFunctionSignatureNone            IntersectionFunctionSignature = 0
	IntersectionFunctionSignaturePrimitiveMotion IntersectionFunctionSignature = 16
	IntersectionFunctionSignatureTriangleData    IntersectionFunctionSignature = 2
	IntersectionFunctionSignatureWorldSpaceData  IntersectionFunctionSignature = 4
)

// Metal shading language versions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllanguageversion?language=objc
type LanguageVersion uint

const (
	LanguageVersion1_1 LanguageVersion = 65537
	LanguageVersion1_2 LanguageVersion = 65538
	LanguageVersion2_0 LanguageVersion = 131072
	LanguageVersion2_1 LanguageVersion = 131073
	LanguageVersion2_2 LanguageVersion = 131074
	LanguageVersion2_3 LanguageVersion = 131075
	LanguageVersion2_4 LanguageVersion = 131076
)

// Error codes returned by Metal for library errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibraryerror?language=objc
type LibraryError uint

const (
	LibraryErrorCompileFailure   LibraryError = 3
	LibraryErrorCompileWarning   LibraryError = 4
	LibraryErrorFileNotFound     LibraryError = 6
	LibraryErrorFunctionNotFound LibraryError = 5
	LibraryErrorInternal         LibraryError = 2
	LibraryErrorUnsupported      LibraryError = 1
)

// A set of options for Metal library types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtllibrarytype?language=objc
type LibraryType int

const (
	LibraryTypeDynamic    LibraryType = 1
	LibraryTypeExecutable LibraryType = 0
)

// Types of actions performed for an attachment at the start of a rendering pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlloadaction?language=objc
type LoadAction uint

const (
	LoadActionClear    LoadAction = 2
	LoadActionDontCare LoadAction = 0
	LoadActionLoad     LoadAction = 1
)

// Options for specifying how the acceleration structure handles timestamps that are outside the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmotionbordermode?language=objc
type MotionBorderMode uint32

const (
	MotionBorderModeClamp  MotionBorderMode = 0
	MotionBorderModeVanish MotionBorderMode = 1
)

// Filtering options for controlling an MSAA depth resolve operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmultisampledepthresolvefilter?language=objc
type MultisampleDepthResolveFilter uint

const (
	MultisampleDepthResolveFilterMax     MultisampleDepthResolveFilter = 2
	MultisampleDepthResolveFilterMin     MultisampleDepthResolveFilter = 1
	MultisampleDepthResolveFilterSample0 MultisampleDepthResolveFilter = 0
)

// Constants used to control the multisample stencil resolve operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmultisamplestencilresolvefilter?language=objc
type MultisampleStencilResolveFilter uint

const (
	MultisampleStencilResolveFilterDepthResolvedSample MultisampleStencilResolveFilter = 1
	MultisampleStencilResolveFilterSample0             MultisampleStencilResolveFilter = 0
)

// The options that determine the mutability of a buffer's contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlmutability?language=objc
type Mutability uint

const (
	MutabilityDefault   Mutability = 0
	MutabilityImmutable Mutability = 2
	MutabilityMutable   Mutability = 1
)

// Types of tessellation patches that can be inputs of a post-tessellation vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpatchtype?language=objc
type PatchType uint

const (
	PatchTypeNone     PatchType = 0
	PatchTypeQuad     PatchType = 2
	PatchTypeTriangle PatchType = 1
)

// Options that determine how Metal prepares the pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpipelineoption?language=objc
type PipelineOption uint

const (
	PipelineOptionArgumentInfo            PipelineOption = 1
	PipelineOptionBufferTypeInfo          PipelineOption = 2
	PipelineOptionFailOnBinaryArchiveMiss PipelineOption = 4
	PipelineOptionNone                    PipelineOption = 0
)

// The data formats that describe the organization and characteristics of individual pixels in a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpixelformat?language=objc
type PixelFormat uint

const (
	PixelFormatA1BGR5Unorm           PixelFormat = 41
	PixelFormatA8Unorm               PixelFormat = 1
	PixelFormatABGR4Unorm            PixelFormat = 42
	PixelFormatASTC_10x10_HDR        PixelFormat = 234
	PixelFormatASTC_10x10_LDR        PixelFormat = 216
	PixelFormatASTC_10x10_sRGB       PixelFormat = 198
	PixelFormatASTC_10x5_HDR         PixelFormat = 231
	PixelFormatASTC_10x5_LDR         PixelFormat = 213
	PixelFormatASTC_10x5_sRGB        PixelFormat = 195
	PixelFormatASTC_10x6_HDR         PixelFormat = 232
	PixelFormatASTC_10x6_LDR         PixelFormat = 214
	PixelFormatASTC_10x6_sRGB        PixelFormat = 196
	PixelFormatASTC_10x8_HDR         PixelFormat = 233
	PixelFormatASTC_10x8_LDR         PixelFormat = 215
	PixelFormatASTC_10x8_sRGB        PixelFormat = 197
	PixelFormatASTC_12x10_HDR        PixelFormat = 235
	PixelFormatASTC_12x10_LDR        PixelFormat = 217
	PixelFormatASTC_12x10_sRGB       PixelFormat = 199
	PixelFormatASTC_12x12_HDR        PixelFormat = 236
	PixelFormatASTC_12x12_LDR        PixelFormat = 218
	PixelFormatASTC_12x12_sRGB       PixelFormat = 200
	PixelFormatASTC_4x4_HDR          PixelFormat = 222
	PixelFormatASTC_4x4_LDR          PixelFormat = 204
	PixelFormatASTC_4x4_sRGB         PixelFormat = 186
	PixelFormatASTC_5x4_HDR          PixelFormat = 223
	PixelFormatASTC_5x4_LDR          PixelFormat = 205
	PixelFormatASTC_5x4_sRGB         PixelFormat = 187
	PixelFormatASTC_5x5_HDR          PixelFormat = 224
	PixelFormatASTC_5x5_LDR          PixelFormat = 206
	PixelFormatASTC_5x5_sRGB         PixelFormat = 188
	PixelFormatASTC_6x5_HDR          PixelFormat = 225
	PixelFormatASTC_6x5_LDR          PixelFormat = 207
	PixelFormatASTC_6x5_sRGB         PixelFormat = 189
	PixelFormatASTC_6x6_HDR          PixelFormat = 226
	PixelFormatASTC_6x6_LDR          PixelFormat = 208
	PixelFormatASTC_6x6_sRGB         PixelFormat = 190
	PixelFormatASTC_8x5_HDR          PixelFormat = 228
	PixelFormatASTC_8x5_LDR          PixelFormat = 210
	PixelFormatASTC_8x5_sRGB         PixelFormat = 192
	PixelFormatASTC_8x6_HDR          PixelFormat = 229
	PixelFormatASTC_8x6_LDR          PixelFormat = 211
	PixelFormatASTC_8x6_sRGB         PixelFormat = 193
	PixelFormatASTC_8x8_HDR          PixelFormat = 230
	PixelFormatASTC_8x8_LDR          PixelFormat = 212
	PixelFormatASTC_8x8_sRGB         PixelFormat = 194
	PixelFormatB5G6R5Unorm           PixelFormat = 40
	PixelFormatBC1_RGBA              PixelFormat = 130
	PixelFormatBC1_RGBA_sRGB         PixelFormat = 131
	PixelFormatBC2_RGBA              PixelFormat = 132
	PixelFormatBC2_RGBA_sRGB         PixelFormat = 133
	PixelFormatBC3_RGBA              PixelFormat = 134
	PixelFormatBC3_RGBA_sRGB         PixelFormat = 135
	PixelFormatBC4_RSnorm            PixelFormat = 141
	PixelFormatBC4_RUnorm            PixelFormat = 140
	PixelFormatBC5_RGSnorm           PixelFormat = 143
	PixelFormatBC5_RGUnorm           PixelFormat = 142
	PixelFormatBC6H_RGBFloat         PixelFormat = 150
	PixelFormatBC6H_RGBUfloat        PixelFormat = 151
	PixelFormatBC7_RGBAUnorm         PixelFormat = 152
	PixelFormatBC7_RGBAUnorm_sRGB    PixelFormat = 153
	PixelFormatBGR10A2Unorm          PixelFormat = 94
	PixelFormatBGR10_XR              PixelFormat = 554
	PixelFormatBGR10_XR_sRGB         PixelFormat = 555
	PixelFormatBGR5A1Unorm           PixelFormat = 43
	PixelFormatBGRA10_XR             PixelFormat = 552
	PixelFormatBGRA10_XR_sRGB        PixelFormat = 553
	PixelFormatBGRA8Unorm            PixelFormat = 80
	PixelFormatBGRA8Unorm_sRGB       PixelFormat = 81
	PixelFormatBGRG422               PixelFormat = 241
	PixelFormatDepth16Unorm          PixelFormat = 250
	PixelFormatDepth24Unorm_Stencil8 PixelFormat = 255
	PixelFormatDepth32Float          PixelFormat = 252
	PixelFormatDepth32Float_Stencil8 PixelFormat = 260
	PixelFormatEAC_R11Snorm          PixelFormat = 172
	PixelFormatEAC_R11Unorm          PixelFormat = 170
	PixelFormatEAC_RG11Snorm         PixelFormat = 176
	PixelFormatEAC_RG11Unorm         PixelFormat = 174
	PixelFormatEAC_RGBA8             PixelFormat = 178
	PixelFormatEAC_RGBA8_sRGB        PixelFormat = 179
	PixelFormatETC2_RGB8             PixelFormat = 180
	PixelFormatETC2_RGB8A1           PixelFormat = 182
	PixelFormatETC2_RGB8A1_sRGB      PixelFormat = 183
	PixelFormatETC2_RGB8_sRGB        PixelFormat = 181
	PixelFormatGBGR422               PixelFormat = 240
	PixelFormatInvalid               PixelFormat = 0
	PixelFormatPVRTC_RGBA_2BPP       PixelFormat = 164
	PixelFormatPVRTC_RGBA_2BPP_sRGB  PixelFormat = 165
	PixelFormatPVRTC_RGBA_4BPP       PixelFormat = 166
	PixelFormatPVRTC_RGBA_4BPP_sRGB  PixelFormat = 167
	PixelFormatPVRTC_RGB_2BPP        PixelFormat = 160
	PixelFormatPVRTC_RGB_2BPP_sRGB   PixelFormat = 161
	PixelFormatPVRTC_RGB_4BPP        PixelFormat = 162
	PixelFormatPVRTC_RGB_4BPP_sRGB   PixelFormat = 163
	PixelFormatR16Float              PixelFormat = 25
	PixelFormatR16Sint               PixelFormat = 24
	PixelFormatR16Snorm              PixelFormat = 22
	PixelFormatR16Uint               PixelFormat = 23
	PixelFormatR16Unorm              PixelFormat = 20
	PixelFormatR32Float              PixelFormat = 55
	PixelFormatR32Sint               PixelFormat = 54
	PixelFormatR32Uint               PixelFormat = 53
	PixelFormatR8Sint                PixelFormat = 14
	PixelFormatR8Snorm               PixelFormat = 12
	PixelFormatR8Uint                PixelFormat = 13
	PixelFormatR8Unorm               PixelFormat = 10
	PixelFormatR8Unorm_sRGB          PixelFormat = 11
	PixelFormatRG11B10Float          PixelFormat = 92
	PixelFormatRG16Float             PixelFormat = 65
	PixelFormatRG16Sint              PixelFormat = 64
	PixelFormatRG16Snorm             PixelFormat = 62
	PixelFormatRG16Uint              PixelFormat = 63
	PixelFormatRG16Unorm             PixelFormat = 60
	PixelFormatRG32Float             PixelFormat = 105
	PixelFormatRG32Sint              PixelFormat = 104
	PixelFormatRG32Uint              PixelFormat = 103
	PixelFormatRG8Sint               PixelFormat = 34
	PixelFormatRG8Snorm              PixelFormat = 32
	PixelFormatRG8Uint               PixelFormat = 33
	PixelFormatRG8Unorm              PixelFormat = 30
	PixelFormatRG8Unorm_sRGB         PixelFormat = 31
	PixelFormatRGB10A2Uint           PixelFormat = 91
	PixelFormatRGB10A2Unorm          PixelFormat = 90
	PixelFormatRGB9E5Float           PixelFormat = 93
	PixelFormatRGBA16Float           PixelFormat = 115
	PixelFormatRGBA16Sint            PixelFormat = 114
	PixelFormatRGBA16Snorm           PixelFormat = 112
	PixelFormatRGBA16Uint            PixelFormat = 113
	PixelFormatRGBA16Unorm           PixelFormat = 110
	PixelFormatRGBA32Float           PixelFormat = 125
	PixelFormatRGBA32Sint            PixelFormat = 124
	PixelFormatRGBA32Uint            PixelFormat = 123
	PixelFormatRGBA8Sint             PixelFormat = 74
	PixelFormatRGBA8Snorm            PixelFormat = 72
	PixelFormatRGBA8Uint             PixelFormat = 73
	PixelFormatRGBA8Unorm            PixelFormat = 70
	PixelFormatRGBA8Unorm_sRGB       PixelFormat = 71
	PixelFormatStencil8              PixelFormat = 253
	PixelFormatX24_Stencil8          PixelFormat = 262
	PixelFormatX32_Stencil8          PixelFormat = 261
)

// The primitive topologies available for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitivetopologyclass?language=objc
type PrimitiveTopologyClass uint

const (
	PrimitiveTopologyClassLine        PrimitiveTopologyClass = 2
	PrimitiveTopologyClassPoint       PrimitiveTopologyClass = 1
	PrimitiveTopologyClassTriangle    PrimitiveTopologyClass = 3
	PrimitiveTopologyClassUnspecified PrimitiveTopologyClass = 0
)

// The geometric primitive type for drawing commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlprimitivetype?language=objc
type PrimitiveType uint

const (
	PrimitiveTypeLine          PrimitiveType = 1
	PrimitiveTypeLineStrip     PrimitiveType = 2
	PrimitiveTypePoint         PrimitiveType = 0
	PrimitiveTypeTriangle      PrimitiveType = 3
	PrimitiveTypeTriangleStrip PrimitiveType = 4
)

// The purgeable state of the resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlpurgeablestate?language=objc
type PurgeableState uint

const (
	PurgeableStateEmpty       PurgeableState = 4
	PurgeableStateKeepCurrent PurgeableState = 1
	PurgeableStateNonVolatile PurgeableState = 2
	PurgeableStateVolatile    PurgeableState = 3
)

// The support level for read-write texture formats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlreadwritetexturetier?language=objc
type ReadWriteTextureTier uint

const (
	ReadWriteTextureTier1    ReadWriteTextureTier = 1
	ReadWriteTextureTier2    ReadWriteTextureTier = 2
	ReadWriteTextureTierNone ReadWriteTextureTier = 0
)

// The stages in a render pass that triggers a synchronization command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrenderstages?language=objc
type RenderStages uint

const (
	RenderStageFragment RenderStages = 2
	RenderStageTile     RenderStages = 4
	RenderStageVertex   RenderStages = 1
)

// Optional arguments used to set the behavior of a resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourceoptions?language=objc
type ResourceOptions uint

const (
	ResourceCPUCacheModeDefaultCache        ResourceOptions = 0
	ResourceCPUCacheModeWriteCombined       ResourceOptions = 1
	ResourceHazardTrackingModeDefault       ResourceOptions = 0
	ResourceHazardTrackingModeTracked       ResourceOptions = 512
	ResourceHazardTrackingModeUntracked     ResourceOptions = 256
	ResourceOptionCPUCacheModeDefault       ResourceOptions = 0
	ResourceOptionCPUCacheModeWriteCombined ResourceOptions = 1
	ResourceStorageModeManaged              ResourceOptions = 16
	ResourceStorageModeMemoryless           ResourceOptions = 48
	ResourceStorageModePrivate              ResourceOptions = 32
	ResourceStorageModeShared               ResourceOptions = 0
)

// Options that describe how a graphics or compute function uses an argument buffer’s resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresourceusage?language=objc
type ResourceUsage uint

const (
	ResourceUsageRead   ResourceUsage = 1
	ResourceUsageSample ResourceUsage = 4
	ResourceUsageWrite  ResourceUsage = 2
)

// Modes that determine the texture coordinate at each pixel when a fetch falls outside the bounds of a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsampleraddressmode?language=objc
type SamplerAddressMode uint

const (
	SamplerAddressModeClampToBorderColor SamplerAddressMode = 5
	SamplerAddressModeClampToEdge        SamplerAddressMode = 0
	SamplerAddressModeClampToZero        SamplerAddressMode = 4
	SamplerAddressModeMirrorClampToEdge  SamplerAddressMode = 1
	SamplerAddressModeMirrorRepeat       SamplerAddressMode = 3
	SamplerAddressModeRepeat             SamplerAddressMode = 2
)

// Values that determine the border color for clamped texture values when the sampler address mode is [metal/mtlsampleraddressmode/mtlsampleraddressmodeclamptobordercolor]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerbordercolor?language=objc
type SamplerBorderColor uint

const (
	SamplerBorderColorOpaqueBlack      SamplerBorderColor = 1
	SamplerBorderColorOpaqueWhite      SamplerBorderColor = 2
	SamplerBorderColorTransparentBlack SamplerBorderColor = 0
)

// Filtering options for determining which pixel value is returned within a mipmap level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplerminmagfilter?language=objc
type SamplerMinMagFilter uint

const (
	SamplerMinMagFilterLinear  SamplerMinMagFilter = 1
	SamplerMinMagFilterNearest SamplerMinMagFilter = 0
)

// Filtering options for determining what pixel value is returned with multiple mipmap levels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsamplermipfilter?language=objc
type SamplerMipFilter uint

const (
	SamplerMipFilterLinear       SamplerMipFilter = 2
	SamplerMipFilterNearest      SamplerMipFilter = 1
	SamplerMipFilterNotMipmapped SamplerMipFilter = 0
)

// Options for sparse texture mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsparsetexturemappingmode?language=objc
type SparseTextureMappingMode uint

const (
	SparseTextureMappingModeMap   SparseTextureMappingMode = 0
	SparseTextureMappingModeUnmap SparseTextureMappingMode = 1
)

// Options used when converting between a pixel-based region within a texture to a tile-based region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsparsetextureregionalignmentmode?language=objc
type SparseTextureRegionAlignmentMode uint

const (
	SparseTextureRegionAlignmentModeInward  SparseTextureRegionAlignmentMode = 1
	SparseTextureRegionAlignmentModeOutward SparseTextureRegionAlignmentMode = 0
)

// The operation performed on a currently stored stencil value when a comparison test passes or fails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstenciloperation?language=objc
type StencilOperation uint

const (
	StencilOperationDecrementClamp StencilOperation = 4
	StencilOperationDecrementWrap  StencilOperation = 7
	StencilOperationIncrementClamp StencilOperation = 3
	StencilOperationIncrementWrap  StencilOperation = 6
	StencilOperationInvert         StencilOperation = 5
	StencilOperationKeep           StencilOperation = 0
	StencilOperationReplace        StencilOperation = 2
	StencilOperationZero           StencilOperation = 1
)

// The frequency with which a function fetches attribute data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstepfunction?language=objc
type StepFunction uint

const (
	StepFunctionConstant                     StepFunction = 0
	StepFunctionPerInstance                  StepFunction = 2
	StepFunctionPerPatch                     StepFunction = 3
	StepFunctionPerPatchControlPoint         StepFunction = 4
	StepFunctionPerVertex                    StepFunction = 1
	StepFunctionThreadPositionInGridX        StepFunction = 5
	StepFunctionThreadPositionInGridXIndexed StepFunction = 7
	StepFunctionThreadPositionInGridY        StepFunction = 6
	StepFunctionThreadPositionInGridYIndexed StepFunction = 8
)

// Options for the memory location and access permissions for a resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstoragemode?language=objc
type StorageMode uint

const (
	StorageModeManaged    StorageMode = 1
	StorageModeMemoryless StorageMode = 3
	StorageModePrivate    StorageMode = 2
	StorageModeShared     StorageMode = 0
)

// Types of actions performed for an attachment at the end of a rendering pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstoreaction?language=objc
type StoreAction uint

const (
	StoreActionCustomSampleDepthStore     StoreAction = 5
	StoreActionDontCare                   StoreAction = 0
	StoreActionMultisampleResolve         StoreAction = 2
	StoreActionStore                      StoreAction = 1
	StoreActionStoreAndMultisampleResolve StoreAction = 3
	StoreActionUnknown                    StoreAction = 4
)

// Options that modify a store action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstoreactionoptions?language=objc
type StoreActionOptions uint

const (
	StoreActionOptionCustomSamplePositions StoreActionOptions = 1
	StoreActionOptionNone                  StoreActionOptions = 0
)

// Options for specifying the size of the control point indices in a control point index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltessellationcontrolpointindextype?language=objc
type TessellationControlPointIndexType uint

const (
	TessellationControlPointIndexTypeNone   TessellationControlPointIndexType = 0
	TessellationControlPointIndexTypeUInt16 TessellationControlPointIndexType = 1
	TessellationControlPointIndexTypeUInt32 TessellationControlPointIndexType = 2
)

// Options for specifying the format of the tessellation factors in a tessellation factor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltessellationfactorformat?language=objc
type TessellationFactorFormat uint

const (
	TessellationFactorFormatHalf TessellationFactorFormat = 0
)

// Options for specifying the step function that determines the tessellation factors for a patch from the tessellation factor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltessellationfactorstepfunction?language=objc
type TessellationFactorStepFunction uint

const (
	TessellationFactorStepFunctionConstant               TessellationFactorStepFunction = 0
	TessellationFactorStepFunctionPerInstance            TessellationFactorStepFunction = 2
	TessellationFactorStepFunctionPerPatch               TessellationFactorStepFunction = 1
	TessellationFactorStepFunctionPerPatchAndPerInstance TessellationFactorStepFunction = 3
)

// Options for specifying the partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltessellationpartitionmode?language=objc
type TessellationPartitionMode uint

const (
	TessellationPartitionModeFractionalEven TessellationPartitionMode = 3
	TessellationPartitionModeFractionalOdd  TessellationPartitionMode = 2
	TessellationPartitionModeInteger        TessellationPartitionMode = 1
	TessellationPartitionModePow2           TessellationPartitionMode = 0
)

// A set of options to choose from when creating a texture swizzle pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltextureswizzle?language=objc
type TextureSwizzle uint8

const (
	TextureSwizzleAlpha TextureSwizzle = 5
	TextureSwizzleBlue  TextureSwizzle = 4
	TextureSwizzleGreen TextureSwizzle = 3
	TextureSwizzleOne   TextureSwizzle = 1
	TextureSwizzleRed   TextureSwizzle = 2
	TextureSwizzleZero  TextureSwizzle = 0
)

// The dimension of each image, including whether multiple images are arranged into an array or a cube. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltexturetype?language=objc
type TextureType uint

const (
	TextureType1D                 TextureType = 0
	TextureType1DArray            TextureType = 1
	TextureType2D                 TextureType = 2
	TextureType2DArray            TextureType = 3
	TextureType2DMultisample      TextureType = 4
	TextureType2DMultisampleArray TextureType = 8
	TextureType3D                 TextureType = 7
	TextureTypeCube               TextureType = 5
	TextureTypeCubeArray          TextureType = 6
	TextureTypeTextureBuffer      TextureType = 9
)

// An enumeration for the various options that determine how you can use a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltextureusage?language=objc
type TextureUsage uint

const (
	TextureUsagePixelFormatView TextureUsage = 16
	TextureUsageRenderTarget    TextureUsage = 4
	TextureUsageShaderRead      TextureUsage = 1
	TextureUsageShaderWrite     TextureUsage = 2
	TextureUsageUnknown         TextureUsage = 0
)

// The number of nanoseconds for a point in absolute time or Mach absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltimestamp?language=objc
type Timestamp uint64

// Specifies how to rasterize triangle and triangle strip primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtltrianglefillmode?language=objc
type TriangleFillMode uint

const (
	TriangleFillModeFill  TriangleFillMode = 0
	TriangleFillModeLines TriangleFillMode = 1
)

// Values that specify the organization of function vertex data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexformat?language=objc
type VertexFormat uint

const (
	VertexFormatChar                  VertexFormat = 46
	VertexFormatChar2                 VertexFormat = 4
	VertexFormatChar2Normalized       VertexFormat = 10
	VertexFormatChar3                 VertexFormat = 5
	VertexFormatChar3Normalized       VertexFormat = 11
	VertexFormatChar4                 VertexFormat = 6
	VertexFormatChar4Normalized       VertexFormat = 12
	VertexFormatCharNormalized        VertexFormat = 48
	VertexFormatFloat                 VertexFormat = 28
	VertexFormatFloat2                VertexFormat = 29
	VertexFormatFloat3                VertexFormat = 30
	VertexFormatFloat4                VertexFormat = 31
	VertexFormatHalf                  VertexFormat = 53
	VertexFormatHalf2                 VertexFormat = 25
	VertexFormatHalf3                 VertexFormat = 26
	VertexFormatHalf4                 VertexFormat = 27
	VertexFormatInt                   VertexFormat = 32
	VertexFormatInt1010102Normalized  VertexFormat = 40
	VertexFormatInt2                  VertexFormat = 33
	VertexFormatInt3                  VertexFormat = 34
	VertexFormatInt4                  VertexFormat = 35
	VertexFormatInvalid               VertexFormat = 0
	VertexFormatShort                 VertexFormat = 50
	VertexFormatShort2                VertexFormat = 16
	VertexFormatShort2Normalized      VertexFormat = 22
	VertexFormatShort3                VertexFormat = 17
	VertexFormatShort3Normalized      VertexFormat = 23
	VertexFormatShort4                VertexFormat = 18
	VertexFormatShort4Normalized      VertexFormat = 24
	VertexFormatShortNormalized       VertexFormat = 52
	VertexFormatUChar                 VertexFormat = 45
	VertexFormatUChar2                VertexFormat = 1
	VertexFormatUChar2Normalized      VertexFormat = 7
	VertexFormatUChar3                VertexFormat = 2
	VertexFormatUChar3Normalized      VertexFormat = 8
	VertexFormatUChar4                VertexFormat = 3
	VertexFormatUChar4Normalized      VertexFormat = 9
	VertexFormatUChar4Normalized_BGRA VertexFormat = 42
	VertexFormatUCharNormalized       VertexFormat = 47
	VertexFormatUInt                  VertexFormat = 36
	VertexFormatUInt1010102Normalized VertexFormat = 41
	VertexFormatUInt2                 VertexFormat = 37
	VertexFormatUInt3                 VertexFormat = 38
	VertexFormatUInt4                 VertexFormat = 39
	VertexFormatUShort                VertexFormat = 49
	VertexFormatUShort2               VertexFormat = 13
	VertexFormatUShort2Normalized     VertexFormat = 19
	VertexFormatUShort3               VertexFormat = 14
	VertexFormatUShort3Normalized     VertexFormat = 20
	VertexFormatUShort4               VertexFormat = 15
	VertexFormatUShort4Normalized     VertexFormat = 21
	VertexFormatUShortNormalized      VertexFormat = 51
)

// The frequency with which the vertex function or post-tessellation vertex function fetches attribute data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexstepfunction?language=objc
type VertexStepFunction uint

const (
	VertexStepFunctionConstant             VertexStepFunction = 0
	VertexStepFunctionPerInstance          VertexStepFunction = 2
	VertexStepFunctionPerPatch             VertexStepFunction = 3
	VertexStepFunctionPerPatchControlPoint VertexStepFunction = 4
	VertexStepFunctionPerVertex            VertexStepFunction = 1
)

// The mode that determines what, if anything, the GPU writes to the results buffer, after the GPU executes the render pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvisibilityresultmode?language=objc
type VisibilityResultMode uint

const (
	VisibilityResultModeBoolean  VisibilityResultMode = 1
	VisibilityResultModeCounting VisibilityResultMode = 2
	VisibilityResultModeDisabled VisibilityResultMode = 0
)

// The vertex winding rule that determines a front-facing primitive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlwinding?language=objc
type Winding uint

const (
	WindingClockwise        Winding = 0
	WindingCounterClockwise Winding = 1
)