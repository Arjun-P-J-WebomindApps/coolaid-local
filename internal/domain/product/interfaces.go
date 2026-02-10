package techspec

import (
	"context"
	"database/sql"
)

// =======================================================
// QUERIES INTERFACE (READ + WRITE CONTRACT)
// =======================================================

type Queries interface {

	// ================= ACTUATOR =================
	GetActuatorByPartNo(ctx context.Context, partNo string) (*ActuatorRow, error)
	CreateActuator(ctx context.Context, p CreateActuatorParams) (*ActuatorRow, error)
	UpdateActuatorByPartNo(ctx context.Context, p UpdateActuatorParams) (*ActuatorRow, error)
	DeleteActuatorByPartNo(ctx context.Context, partNo string) error

	// ================= BLOWER MOTOR =================
	GetBlowerMotorByPartNo(ctx context.Context, partNo string) (*BlowerMotorRow, error)
	CreateBlowerMotor(ctx context.Context, p CreateBlowerMotorParams) (*BlowerMotorRow, error)
	UpdateBlowerMotorByPartNo(ctx context.Context, p UpdateBlowerMotorParams) (*BlowerMotorRow, error)
	DeleteBlowerMotorByPartNo(ctx context.Context, partNo string) error

	// ================= CABIN FILTER =================
	GetCabinFilterByPartNo(ctx context.Context, partNo string) (*CabinFilterRow, error)
	CreateCabinFilter(ctx context.Context, p CreateCabinFilterParams) (*CabinFilterRow, error)
	UpdateCabinFilterByPartNo(ctx context.Context, p UpdateCabinFilterParams) (*CabinFilterRow, error)
	DeleteCabinFilterByPartNo(ctx context.Context, partNo string) error

	// ================= CHILLER UNIT =================
	GetChillerUnitByPartNo(ctx context.Context, partNo string) (*ChillerUnitRow, error)
	CreateChillerUnit(ctx context.Context, p CreateChillerUnitParams) (*ChillerUnitRow, error)
	UpdateChillerUnitByPartNo(ctx context.Context, p UpdateChillerUnitParams) (*ChillerUnitRow, error)
	DeleteChillerUnitByPartNo(ctx context.Context, partNo string) error

	// ================= CLUTCH ASSY =================
	GetClutchAssyByPartNo(ctx context.Context, partNo string) (*ClutchAssyRow, error)
	CreateClutchAssy(ctx context.Context, p CreateClutchAssyParams) (*ClutchAssyRow, error)
	UpdateClutchAssyByPartNo(ctx context.Context, p UpdateClutchAssyParams) (*ClutchAssyRow, error)
	DeleteClutchAssyByPartNo(ctx context.Context, partNo string) error

	// ================= COMPRESSOR =================
	GetCompressorByPartNo(ctx context.Context, partNo string) (*CompressorRow, error)
	CreateCompressor(ctx context.Context, p CreateCompressorParams) (*CompressorRow, error)
	UpdateCompressorByPartNo(ctx context.Context, p UpdateCompressorParams) (*CompressorRow, error)
	DeleteCompressorByPartNo(ctx context.Context, partNo string) error

	// ================= COMPRESSOR VALVE =================
	GetCompressorValveByPartNo(ctx context.Context, partNo string) (*CompressorValveRow, error)
	CreateCompressorValve(ctx context.Context, p CreateCompressorValveParams) (*CompressorValveRow, error)
	UpdateCompressorValveByPartNo(ctx context.Context, p UpdateCompressorValveParams) (*CompressorValveRow, error)
	DeleteCompressorValveByPartNo(ctx context.Context, partNo string) error

	// ================= CONDENSER FAN ASSY =================
	GetCondFanAssyByPartNo(ctx context.Context, partNo string) (*CondFanAssyRow, error)
	CreateCondFanAssy(ctx context.Context, p CreateCondFanAssyParams) (*CondFanAssyRow, error)
	UpdateCondFanAssyByPartNo(ctx context.Context, p UpdateCondFanAssyParams) (*CondFanAssyRow, error)
	DeleteCondFanAssyByPartNo(ctx context.Context, partNo string) error

	// ================= CONDENSER =================
	GetCondenserByPartNo(ctx context.Context, partNo string) (*CondenserRow, error)
	CreateCondenser(ctx context.Context, p CreateCondenserParams) (*CondenserRow, error)
	UpdateCondenserByPartNo(ctx context.Context, p UpdateCondenserParams) (*CondenserRow, error)
	DeleteCondenserByPartNo(ctx context.Context, partNo string) error

	// ================= EVAPORATOR =================
	GetEvaporatorByPartNo(ctx context.Context, partNo string) (*EvaporatorRow, error)
	CreateEvaporator(ctx context.Context, p CreateEvaporatorParams) (*EvaporatorRow, error)
	UpdateEvaporatorByPartNo(ctx context.Context, p UpdateEvaporatorParams) (*EvaporatorRow, error)
	DeleteEvaporatorByPartNo(ctx context.Context, partNo string) error

	// ================= EXPANSION VALVE =================
	GetExpansionValveByPartNo(ctx context.Context, partNo string) (*ExpansionValveRow, error)
	CreateExpansionValve(ctx context.Context, p CreateExpansionValveParams) (*ExpansionValveRow, error)
	UpdateExpansionValveByPartNo(ctx context.Context, p UpdateExpansionValveParams) (*ExpansionValveRow, error)
	DeleteExpansionValveByPartNo(ctx context.Context, partNo string) error

	// ================= FILTER DRIER =================
	GetFilterDrierByPartNo(ctx context.Context, partNo string) (*FilterDrierRow, error)
	CreateFilterDrier(ctx context.Context, p CreateFilterDrierParams) (*FilterDrierRow, error)
	UpdateFilterDrierByPartNo(ctx context.Context, p UpdateFilterDrierParams) (*FilterDrierRow, error)
	DeleteFilterDrierByPartNo(ctx context.Context, partNo string) error

	// ================= HEATER CORE =================
	GetHeaterCoreByPartNo(ctx context.Context, partNo string) (*HeaterCoreRow, error)
	CreateHeaterCore(ctx context.Context, p CreateHeaterCoreParams) (*HeaterCoreRow, error)
	UpdateHeaterCoreByPartNo(ctx context.Context, p UpdateHeaterCoreParams) (*HeaterCoreRow, error)
	DeleteHeaterCoreByPartNo(ctx context.Context, partNo string) error

	// ================= INTERCOOLER =================
	GetIntercoolerByPartNo(ctx context.Context, partNo string) (*IntercoolerRow, error)
	CreateIntercooler(ctx context.Context, p CreateIntercoolerParams) (*IntercoolerRow, error)
	UpdateIntercoolerByPartNo(ctx context.Context, p UpdateIntercoolerParams) (*IntercoolerRow, error)
	DeleteIntercoolerByPartNo(ctx context.Context, partNo string) error

	// ================= PRESSURE SWITCH =================
	GetPressureSwitchByPartNo(ctx context.Context, partNo string) (*PressureSwitchRow, error)
	CreatePressureSwitch(ctx context.Context, p CreatePressureSwitchParams) (*PressureSwitchRow, error)
	UpdatePressureSwitchByPartNo(ctx context.Context, p UpdatePressureSwitchParams) (*PressureSwitchRow, error)
	DeletePressureSwitchByPartNo(ctx context.Context, partNo string) error

	// ================= RADIATOR =================
	GetRadiatorByPartNo(ctx context.Context, partNo string) (*RadiatorRow, error)
	CreateRadiator(ctx context.Context, p CreateRadiatorParams) (*RadiatorRow, error)
	UpdateRadiatorByPartNo(ctx context.Context, p UpdateRadiatorParams) (*RadiatorRow, error)
	DeleteRadiatorByPartNo(ctx context.Context, partNo string) error

	// ================= RAD FAN ASSY =================
	GetRadFanAssyByPartNo(ctx context.Context, partNo string) (*RadFanAssyRow, error)
	CreateRadFanAssy(ctx context.Context, p CreateRadFanAssyParams) (*RadFanAssyRow, error)
	UpdateRadFanAssyByPartNo(ctx context.Context, p UpdateRadFanAssyParams) (*RadFanAssyRow, error)
	DeleteRadFanAssyByPartNo(ctx context.Context, partNo string) error

	// ================= RAD FAN MOTOR =================
	GetRadFanMotorByPartNo(ctx context.Context, partNo string) (*RadFanMotorRow, error)
	CreateRadFanMotor(ctx context.Context, p CreateRadFanMotorParams) (*RadFanMotorRow, error)
	UpdateRadFanMotorByPartNo(ctx context.Context, p UpdateRadFanMotorParams) (*RadFanMotorRow, error)
	DeleteRadFanMotorByPartNo(ctx context.Context, partNo string) error

	// ================= RESISTOR =================
	GetResistorByPartNo(ctx context.Context, partNo string) (*ResistorRow, error)
	CreateResistor(ctx context.Context, p CreateResistorParams) (*ResistorRow, error)
	UpdateResistorByPartNo(ctx context.Context, p UpdateResistorParams) (*ResistorRow, error)
	DeleteResistorByPartNo(ctx context.Context, partNo string) error

	// ================= ROTOR =================
	GetRotorByPartNo(ctx context.Context, partNo string) (*RotorRow, error)
	CreateRotor(ctx context.Context, p CreateRotorParams) (*RotorRow, error)
	UpdateRotorByPartNo(ctx context.Context, p UpdateRotorParams) (*RotorRow, error)
	DeleteRotorByPartNo(ctx context.Context, partNo string) error

	// ================= STATOR =================
	GetStatorByPartNo(ctx context.Context, partNo string) (*StatorRow, error)
	CreateStator(ctx context.Context, p CreateStatorParams) (*StatorRow, error)
	UpdateStatorByPartNo(ctx context.Context, p UpdateStatorParams) (*StatorRow, error)
	DeleteStatorByPartNo(ctx context.Context, partNo string) error
}

// =======================================================
// DB INTERFACE (TRANSACTION BOUNDARY)
// =======================================================

type DB interface {
	BeginTx(ctx context.Context) (*sql.Tx, Queries, error)
	Queries() Queries
}
