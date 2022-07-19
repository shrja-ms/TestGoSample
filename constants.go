//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.7.6, generator: @autorest/go@4.0.0-preview.42)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

const host = "https://pluginBaseUrl.com"

type AbsoluteMarker string

const (
	AbsoluteMarkerAllBackup AbsoluteMarker = "AllBackup"
	AbsoluteMarkerFirstOfDay AbsoluteMarker = "FirstOfDay"
	AbsoluteMarkerFirstOfMonth AbsoluteMarker = "FirstOfMonth"
	AbsoluteMarkerFirstOfWeek AbsoluteMarker = "FirstOfWeek"
	AbsoluteMarkerFirstOfYear AbsoluteMarker = "FirstOfYear"
	AbsoluteMarkerInvalid AbsoluteMarker = "Invalid"
)

// PossibleAbsoluteMarkerValues returns the possible values for the AbsoluteMarker const type.
func PossibleAbsoluteMarkerValues() []AbsoluteMarker {
	return []AbsoluteMarker{	
		AbsoluteMarkerAllBackup,
		AbsoluteMarkerFirstOfDay,
		AbsoluteMarkerFirstOfMonth,
		AbsoluteMarkerFirstOfWeek,
		AbsoluteMarkerFirstOfYear,
		AbsoluteMarkerInvalid,
	}
}

// AutoHealStatus - Policy controlled toggle
type AutoHealStatus string

const (
	AutoHealStatusOff AutoHealStatus = "Off"
	AutoHealStatusOn AutoHealStatus = "On"
)

// PossibleAutoHealStatusValues returns the possible values for the AutoHealStatus const type.
func PossibleAutoHealStatusValues() []AutoHealStatus {
	return []AutoHealStatus{	
		AutoHealStatusOff,
		AutoHealStatusOn,
	}
}

// DataStoreTypes - type of datastore; SnapShot/Hot/Archive
type DataStoreTypes string

const (
	DataStoreTypesArchiveStore DataStoreTypes = "ArchiveStore"
	DataStoreTypesOperationalStore DataStoreTypes = "OperationalStore"
	DataStoreTypesVaultStore DataStoreTypes = "VaultStore"
)

// PossibleDataStoreTypesValues returns the possible values for the DataStoreTypes const type.
func PossibleDataStoreTypesValues() []DataStoreTypes {
	return []DataStoreTypes{	
		DataStoreTypesArchiveStore,
		DataStoreTypesOperationalStore,
		DataStoreTypesVaultStore,
	}
}

type DayOfWeek string

const (
	DayOfWeekFriday DayOfWeek = "Friday"
	DayOfWeekMonday DayOfWeek = "Monday"
	DayOfWeekSaturday DayOfWeek = "Saturday"
	DayOfWeekSunday DayOfWeek = "Sunday"
	DayOfWeekThursday DayOfWeek = "Thursday"
	DayOfWeekTuesday DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{	
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// ExecutionStatus - Service-set extensible enum indicating operation?s kind
type ExecutionStatus string

const (
	ExecutionStatusCancelled ExecutionStatus = "Cancelled"
	ExecutionStatusFailed ExecutionStatus = "Failed"
	ExecutionStatusNotStarted ExecutionStatus = "NotStarted"
	ExecutionStatusRunning ExecutionStatus = "Running"
	ExecutionStatusSucceeded ExecutionStatus = "Succeeded"
)

// PossibleExecutionStatusValues returns the possible values for the ExecutionStatus const type.
func PossibleExecutionStatusValues() []ExecutionStatus {
	return []ExecutionStatus{	
		ExecutionStatusCancelled,
		ExecutionStatusFailed,
		ExecutionStatusNotStarted,
		ExecutionStatusRunning,
		ExecutionStatusSucceeded,
	}
}

type Month string

const (
	MonthApril Month = "April"
	MonthAugust Month = "August"
	MonthDecember Month = "December"
	MonthFebruary Month = "February"
	MonthJanuary Month = "January"
	MonthJuly Month = "July"
	MonthJune Month = "June"
	MonthMarch Month = "March"
	MonthMay Month = "May"
	MonthNovember Month = "November"
	MonthOctober Month = "October"
	MonthSeptember Month = "September"
)

// PossibleMonthValues returns the possible values for the Month const type.
func PossibleMonthValues() []Month {
	return []Month{	
		MonthApril,
		MonthAugust,
		MonthDecember,
		MonthFebruary,
		MonthJanuary,
		MonthJuly,
		MonthJune,
		MonthMarch,
		MonthMay,
		MonthNovember,
		MonthOctober,
		MonthSeptember,
	}
}

// RestoreTargetLocationType - Denotes the target location where the data will be restored, string value for the enum {Microsoft.Internal.AzureBackup.DataProtection.Common.Interface.RestoreTargetLocationType}
type RestoreTargetLocationType string

const (
	RestoreTargetLocationTypeAzureBlobs RestoreTargetLocationType = "AzureBlobs"
	RestoreTargetLocationTypeAzureDisks RestoreTargetLocationType = "AzureDisks"
	RestoreTargetLocationTypeAzureFiles RestoreTargetLocationType = "AzureFiles"
	RestoreTargetLocationTypeAzureKubernetesWorkload RestoreTargetLocationType = "AzureKubernetesWorkload"
	RestoreTargetLocationTypeInvalid RestoreTargetLocationType = "Invalid"
)

// PossibleRestoreTargetLocationTypeValues returns the possible values for the RestoreTargetLocationType const type.
func PossibleRestoreTargetLocationTypeValues() []RestoreTargetLocationType {
	return []RestoreTargetLocationType{	
		RestoreTargetLocationTypeAzureBlobs,
		RestoreTargetLocationTypeAzureDisks,
		RestoreTargetLocationTypeAzureFiles,
		RestoreTargetLocationTypeAzureKubernetesWorkload,
		RestoreTargetLocationTypeInvalid,
	}
}

// RestoreType - Restore type: OLR, ALR, ILR, Copy
type RestoreType string

const (
	RestoreTypeAlternateLocationRecovery RestoreType = "AlternateLocationRecovery"
	RestoreTypeInvalid RestoreType = "Invalid"
	RestoreTypeItemLevelLocationRecovery RestoreType = "ItemLevelLocationRecovery"
	RestoreTypeOriginalLocationRecovery RestoreType = "OriginalLocationRecovery"
)

// PossibleRestoreTypeValues returns the possible values for the RestoreType const type.
func PossibleRestoreTypeValues() []RestoreType {
	return []RestoreType{	
		RestoreTypeAlternateLocationRecovery,
		RestoreTypeInvalid,
		RestoreTypeItemLevelLocationRecovery,
		RestoreTypeOriginalLocationRecovery,
	}
}

// TriggerType - Trigger Type ? Adhoc/Scheduled/Custom event etc.
type TriggerType string

const (
	TriggerTypeAdhoc TriggerType = "Adhoc"
	TriggerTypeScheduled TriggerType = "Scheduled"
)

// PossibleTriggerTypeValues returns the possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{	
		TriggerTypeAdhoc,
		TriggerTypeScheduled,
	}
}

type WeekNumber string

const (
	WeekNumberFirst WeekNumber = "First"
	WeekNumberFourth WeekNumber = "Fourth"
	WeekNumberInvalid WeekNumber = "Invalid"
	WeekNumberLast WeekNumber = "Last"
	WeekNumberSecond WeekNumber = "Second"
	WeekNumberThird WeekNumber = "Third"
)

// PossibleWeekNumberValues returns the possible values for the WeekNumber const type.
func PossibleWeekNumberValues() []WeekNumber {
	return []WeekNumber{	
		WeekNumberFirst,
		WeekNumberFourth,
		WeekNumberInvalid,
		WeekNumberLast,
		WeekNumberSecond,
		WeekNumberThird,
	}
}

