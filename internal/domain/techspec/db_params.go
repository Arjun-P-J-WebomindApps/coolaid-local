package techspec

// =======================================================
// ACTUATOR
// =======================================================

type CreateActuatorParams struct {
	ID            string
	PartNo        string
	ConnectorType *string
	Mounting      *string
	Voltage       *string
	RotationAngle *string
	Notes         *string
}

type UpdateActuatorParams struct {
	PartNo        string
	ConnectorType *string
	Mounting      *string
	Voltage       *string
	RotationAngle *string
	Notes         *string
}

// =======================================================
// BLOWER MOTOR
// =======================================================

type CreateBlowerMotorParams struct {
	ID            string
	PartNo        string
	Mounting      *string
	ConnectorType *string
	Impeller      *string
	Resistance    *string
	MotorMounting *string
	MotorType     *string
	Voltage       *string
	Notes         *string
}

type UpdateBlowerMotorParams struct {
	PartNo        string
	Mounting      *string
	ConnectorType *string
	Impeller      *string
	Resistance    *string
	MotorMounting *string
	MotorType     *string
	Voltage       *string
	Notes         *string
}

// =======================================================
// CABIN FILTER
// =======================================================

type CreateCabinFilterParams struct {
	ID         string
	PartNo     string
	Type       *string
	Dimensions *string
	Material   *string
	Notes      *string
}

type UpdateCabinFilterParams struct {
	PartNo     string
	Type       *string
	Dimensions *string
	Material   *string
	Notes      *string
}

// =======================================================
// CHILLER UNIT
// =======================================================

type CreateChillerUnitParams struct {
	ID      string
	PartNo  string
	Type    *string
	Voltage *string
	Notes   *string
}

type UpdateChillerUnitParams struct {
	PartNo  string
	Type    *string
	Voltage *string
	Notes   *string
}

// =======================================================
// CLUTCH ASSY
// =======================================================

type CreateClutchAssyParams struct {
	ID                string
	PartNo            string
	PulleyRibs        *string
	PulleySize        *string
	CompressorDetails *string
	ConnectorType     *string
	Voltage           *string
	ShaftType         *string
	Notes             *string
}

type UpdateClutchAssyParams struct {
	PartNo            string
	PulleyRibs        *string
	PulleySize        *string
	CompressorDetails *string
	ConnectorType     *string
	Voltage           *string
	ShaftType         *string
	Notes             *string
}

// =======================================================
// COMPRESSOR VALVE
// =======================================================

type CreateCompressorValveParams struct {
	ID                string
	PartNo            string
	Type              *string
	Voltage           *string
	ConnectorType     *string
	CompressorDetails *string
	Notes             *string
}

type UpdateCompressorValveParams struct {
	PartNo            string
	Type              *string
	Voltage           *string
	ConnectorType     *string
	CompressorDetails *string
	Notes             *string
}

// =======================================================
// COND FAN ASSY
// =======================================================

type CreateCondFanAssyParams struct {
	ID               string
	PartNo           string
	Voltage          *string
	MotorType        *string
	Resistance       *string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Shroud           *string
	ConnectorType    *string
	Notes            *string
}

type UpdateCondFanAssyParams struct {
	PartNo           string
	Voltage          *string
	MotorType        *string
	Resistance       *string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Shroud           *string
	ConnectorType    *string
	Notes            *string
}

// =======================================================
// CONDENSER
// =======================================================

type CreateCondenserParams struct {
	ID             string
	PartNo         string
	Size           *string
	PipeConnector  *string
	Drier          *string
	PressureSwitch *string
	Notes          *string
}

type UpdateCondenserParams struct {
	PartNo         string
	Size           *string
	PipeConnector  *string
	Drier          *string
	PressureSwitch *string
	Notes          *string
}

// =======================================================
// EVAPORATOR
// =======================================================

type CreateEvaporatorParams struct {
	ID             string
	PartNo         string
	Mounting       *string
	ExpValve       *string
	AdditionalInfo *string
	Dimensions     *string
	PipeConnector  *string
	Notes          *string
}

type UpdateEvaporatorParams struct {
	PartNo         string
	Mounting       *string
	ExpValve       *string
	AdditionalInfo *string
	Dimensions     *string
	PipeConnector  *string
	Notes          *string
}

// =======================================================
// EXPANSION VALVE
// =======================================================

type CreateExpansionValveParams struct {
	ID          string
	PartNo      string
	Type        *string
	Material    *string
	Refrigerant *string
	Notes       *string
}

type UpdateExpansionValveParams struct {
	PartNo      string
	Type        *string
	Material    *string
	Refrigerant *string
	Notes       *string
}

// =======================================================
// FILTER DRIER
// =======================================================

type CreateFilterDrierParams struct {
	ID             string
	PartNo         string
	PipeConnector  *string
	Size           *string
	PressureSwitch *string
	Notes          *string
}

type UpdateFilterDrierParams struct {
	PartNo         string
	PipeConnector  *string
	Size           *string
	PressureSwitch *string
	Notes          *string
}

// =======================================================
// HEATER CORE
// =======================================================

type CreateHeaterCoreParams struct {
	ID     string
	PartNo string
	Size   *string
	Pipe   *string
	Type   *string
	Notes  *string
}

type UpdateHeaterCoreParams struct {
	PartNo string
	Size   *string
	Pipe   *string
	Type   *string
	Notes  *string
}

// =======================================================
// INTERCOOLER
// =======================================================

type CreateIntercoolerParams struct {
	ID         string
	PartNo     string
	Size       *string
	TempSensor *string
	Notes      *string
}

type UpdateIntercoolerParams struct {
	PartNo     string
	Size       *string
	TempSensor *string
	Notes      *string
}

// =======================================================
// PRESSURE SWITCH
// =======================================================

type CreatePressureSwitchParams struct {
	ID            string
	PartNo        string
	ConnectorType *string
	ThreadType    *string
	Notes         *string
}

type UpdatePressureSwitchParams struct {
	PartNo        string
	ConnectorType *string
	ThreadType    *string
	Notes         *string
}

// =======================================================
// RADIATOR
// =======================================================

type CreateRadiatorParams struct {
	ID           string
	PartNo       string
	Size         *string
	Transmission *string
	TempSensor   *string
	Tank         *string
	Notes        *string
}

type UpdateRadiatorParams struct {
	PartNo       string
	Size         *string
	Transmission *string
	TempSensor   *string
	Tank         *string
	Notes        *string
}

// =======================================================
// RAD FAN ASSY
// =======================================================

type CreateRadFanAssyParams struct {
	ID               string
	PartNo           string
	Voltage          *string
	MotorType        *string
	Resistance       *string
	NumberOfSockets  *int32
	Shroud           *string
	ConnectorType    *string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Notes            *string
}

type UpdateRadFanAssyParams struct {
	PartNo           string
	Voltage          *string
	MotorType        *string
	Resistance       *string
	NumberOfSockets  *int32
	Shroud           *string
	ConnectorType    *string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Notes            *string
}

// =======================================================
// RAD FAN MOTOR
// =======================================================

type CreateRadFanMotorParams struct {
	ID               string
	PartNo           string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Voltage          *string
	NumberOfSockets  *int32
	ConnectorType    *string
	Notes            *string
}

type UpdateRadFanMotorParams struct {
	PartNo           string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Voltage          *string
	NumberOfSockets  *int32
	ConnectorType    *string
	Notes            *string
}

// =======================================================
// RESISTOR
// =======================================================

type CreateResistorParams struct {
	ID            string
	PartNo        string
	Type          *string
	ConnectorType *string
	Voltage       *string
	Notes         *string
}

type UpdateResistorParams struct {
	PartNo        string
	Type          *string
	ConnectorType *string
	Voltage       *string
	Notes         *string
}

// =======================================================
// ROTOR
// =======================================================

type CreateRotorParams struct {
	ID                string
	PartNo            string
	PulleyRibs        *string
	PulleySize        *string
	CompressorDetails *string
	Notes             *string
}

type UpdateRotorParams struct {
	PartNo            string
	PulleyRibs        *string
	PulleySize        *string
	CompressorDetails *string
	Notes             *string
}

// =======================================================
// STATOR
// =======================================================

type CreateStatorParams struct {
	ID                string
	PartNo            string
	Voltage           *string
	CompressorDetails *string
	Notes             *string
}

type UpdateStatorParams struct {
	PartNo            string
	Voltage           *string
	CompressorDetails *string
	Notes             *string
}

// =======================================================
// COMPRESSOR
// =======================================================

type CreateCompressorParams struct {
	ID            string
	PartNo        string
	CompressorID  *string
	Oil           *string
	Refrigerant   *string
	Voltage       *string
	PulleyRibs    *string
	PulleySize    *string
	PipeConnector *string
	CompType      *string
	CompMounting  *string
	ConnectorType *string
	Notes         *string
}

type UpdateCompressorParams struct {
	PartNo        string
	CompressorID  *string
	Oil           *string
	Refrigerant   *string
	Voltage       *string
	PulleyRibs    *string
	PulleySize    *string
	PipeConnector *string
	CompType      *string
	CompMounting  *string
	ConnectorType *string
	Notes         *string
}
