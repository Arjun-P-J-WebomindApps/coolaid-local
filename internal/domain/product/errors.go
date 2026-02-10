package techspec

import "errors"

// =======================================================
// GENERIC TECHSPEC ERRORS
// =======================================================

var (
	// Common
	ErrInvalidPartNo     = errors.New("invalid part number")
	ErrTechSpecNotFound  = errors.New("technical specification not found")
	ErrDuplicateTechSpec = errors.New("technical specification already exists")
	ErrInvalidTechSpec   = errors.New("invalid technical specification")
	ErrUnsupportedType   = errors.New("unsupported technical specification type")
	ErrInternal          = errors.New("internal error")
)

// =======================================================
// CREATE ERRORS
// =======================================================

var (
	ErrCreateActuatorFailed        = errors.New("failed to create actuator")
	ErrCreateBlowerMotorFailed     = errors.New("failed to create blower motor")
	ErrCreateCabinFilterFailed     = errors.New("failed to create cabin filter")
	ErrCreateChillerUnitFailed     = errors.New("failed to create chiller unit")
	ErrCreateClutchAssyFailed      = errors.New("failed to create clutch assembly")
	ErrCreateCompressorValveFailed = errors.New("failed to create compressor valve")
	ErrCreateCondFanAssyFailed     = errors.New("failed to create condenser fan assembly")
	ErrCreateCondenserFailed       = errors.New("failed to create condenser")
	ErrCreateEvaporatorFailed      = errors.New("failed to create evaporator")
	ErrCreateExpansionValveFailed  = errors.New("failed to create expansion valve")
	ErrCreateFilterDrierFailed     = errors.New("failed to create filter drier")
	ErrCreateHeaterCoreFailed      = errors.New("failed to create heater core")
	ErrCreateIntercoolerFailed     = errors.New("failed to create intercooler")
	ErrCreatePressureSwitchFailed  = errors.New("failed to create pressure switch")
	ErrCreateRadiatorFailed        = errors.New("failed to create radiator")
	ErrCreateRadFanAssyFailed      = errors.New("failed to create radiator fan assembly")
	ErrCreateRadFanMotorFailed     = errors.New("failed to create radiator fan motor")
	ErrCreateResistorFailed        = errors.New("failed to create resistor")
	ErrCreateRotorFailed           = errors.New("failed to create rotor")
	ErrCreateStatorFailed          = errors.New("failed to create stator")
	ErrCreateCompressorFailed      = errors.New("failed to create compressor")
)

// =======================================================
// UPDATE ERRORS
// =======================================================

var (
	ErrUpdateActuatorFailed        = errors.New("failed to update actuator")
	ErrUpdateBlowerMotorFailed     = errors.New("failed to update blower motor")
	ErrUpdateCabinFilterFailed     = errors.New("failed to update cabin filter")
	ErrUpdateChillerUnitFailed     = errors.New("failed to update chiller unit")
	ErrUpdateClutchAssyFailed      = errors.New("failed to update clutch assembly")
	ErrUpdateCompressorValveFailed = errors.New("failed to update compressor valve")
	ErrUpdateCondFanAssyFailed     = errors.New("failed to update condenser fan assembly")
	ErrUpdateCondenserFailed       = errors.New("failed to update condenser")
	ErrUpdateEvaporatorFailed      = errors.New("failed to update evaporator")
	ErrUpdateExpansionValveFailed  = errors.New("failed to update expansion valve")
	ErrUpdateFilterDrierFailed     = errors.New("failed to update filter drier")
	ErrUpdateHeaterCoreFailed      = errors.New("failed to update heater core")
	ErrUpdateIntercoolerFailed     = errors.New("failed to update intercooler")
	ErrUpdatePressureSwitchFailed  = errors.New("failed to update pressure switch")
	ErrUpdateRadiatorFailed        = errors.New("failed to update radiator")
	ErrUpdateRadFanAssyFailed      = errors.New("failed to update radiator fan assembly")
	ErrUpdateRadFanMotorFailed     = errors.New("failed to update radiator fan motor")
	ErrUpdateResistorFailed        = errors.New("failed to update resistor")
	ErrUpdateRotorFailed           = errors.New("failed to update rotor")
	ErrUpdateStatorFailed          = errors.New("failed to update stator")
	ErrUpdateCompressorFailed      = errors.New("failed to update compressor")
)

// =======================================================
// DELETE ERRORS
// =======================================================

var (
	ErrDeleteActuatorFailed        = errors.New("failed to delete actuator")
	ErrDeleteBlowerMotorFailed     = errors.New("failed to delete blower motor")
	ErrDeleteCabinFilterFailed     = errors.New("failed to delete cabin filter")
	ErrDeleteChillerUnitFailed     = errors.New("failed to delete chiller unit")
	ErrDeleteClutchAssyFailed      = errors.New("failed to delete clutch assembly")
	ErrDeleteCompressorValveFailed = errors.New("failed to delete compressor valve")
	ErrDeleteCondFanAssyFailed     = errors.New("failed to delete condenser fan assembly")
	ErrDeleteCondenserFailed       = errors.New("failed to delete condenser")
	ErrDeleteEvaporatorFailed      = errors.New("failed to delete evaporator")
	ErrDeleteExpansionValveFailed  = errors.New("failed to delete expansion valve")
	ErrDeleteFilterDrierFailed     = errors.New("failed to delete filter drier")
	ErrDeleteHeaterCoreFailed      = errors.New("failed to delete heater core")
	ErrDeleteIntercoolerFailed     = errors.New("failed to delete intercooler")
	ErrDeletePressureSwitchFailed  = errors.New("failed to delete pressure switch")
	ErrDeleteRadiatorFailed        = errors.New("failed to delete radiator")
	ErrDeleteRadFanAssyFailed      = errors.New("failed to delete radiator fan assembly")
	ErrDeleteRadFanMotorFailed     = errors.New("failed to delete radiator fan motor")
	ErrDeleteResistorFailed        = errors.New("failed to delete resistor")
	ErrDeleteRotorFailed           = errors.New("failed to delete rotor")
	ErrDeleteStatorFailed          = errors.New("failed to delete stator")
	ErrDeleteCompressorFailed      = errors.New("failed to delete compressor")
)
