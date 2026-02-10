package techspec

// =======================================================
// ROOT TECH SPECS INPUT
// =======================================================

type TechnicalSpecsInput struct {
	Type string `json:"type"`

	Actuator        *ActuatorInput
	BlowerMotor     *BlowerMotorInput
	CabinFilter     *CabinFilterInput
	ChillerUnit     *ChillerUnitInput
	ClutchAssy      *ClutchAssyInput
	CompressorValve *CompressorValveInput
	CondFanAssy     *CondFanAssyInput
	Condenser       *CondenserInput
	Evaporator      *EvaporatorInput
	ExpansionValve  *ExpansionValveInput
	FilterDrier     *FilterDrierInput
	HeaterCore      *HeaterCoreInput
	Intercooler     *IntercoolerInput
	PressureSwitch  *PressureSwitchInput
	Radiator        *RadiatorInput
	RadFanAssy      *RadFanAssyInput
	RadFanMotor     *RadFanMotorInput
	Resistor        *ResistorInput
	Rotor           *RotorInput
	Stator          *StatorInput
	Compressor      *CompressorInput
}

// =======================================================
// INPUT STRUCTS (flat, no overthinking)
// =======================================================

type ActuatorInput struct {
	ConnectorType string
	Mounting      string
	Voltage       string
	RotationAngle string
	Notes         string
}

type BlowerMotorInput struct {
	Mounting      string
	ConnectorType string
	Impeller      string
	Resistance    string
	MotorMounting string
	MotorType     string
	Voltage       string
	Notes         string
}

type CabinFilterInput struct {
	Type       string
	Dimensions string
	Material   string
	Notes      string
}

type ChillerUnitInput struct {
	Type    string
	Voltage string
	Notes   string
}

type ClutchAssyInput struct {
	PulleyRibs        string
	PulleySize        string
	CompressorDetails string
	ConnectorType     string
	Voltage           string
	ShaftType         string
	Notes             string
}

type CompressorValveInput struct {
	Type              string
	Voltage           string
	ConnectorType     string
	CompressorDetails string
	Notes             string
}

type CondFanAssyInput struct {
	Voltage          string
	MotorType        string
	Resistance       string
	FanBladeDiameter string
	NumberOfBlades   *int32
	Shroud           string
	ConnectorType    string
	Notes            string
}

type CondenserInput struct {
	Size           string
	PipeConnector  string
	Drier          string
	PressureSwitch string
	Notes          string
}

type EvaporatorInput struct {
	Mounting       string
	ExpValve       string
	AdditionalInfo string
	Dimensions     string
	PipeConnector  string
	Notes          string
}

type ExpansionValveInput struct {
	Type        string
	Material    string
	Refrigerant string
	Notes       string
}

type FilterDrierInput struct {
	PipeConnector  string
	Size           string
	PressureSwitch string
	Notes          string
}

type HeaterCoreInput struct {
	Size  string
	Pipe  string
	Type  string
	Notes string
}

type IntercoolerInput struct {
	Size       string
	TempSensor string
	Notes      string
}

type PressureSwitchInput struct {
	ConnectorType string
	ThreadType    string
	Notes         string
}

type RadiatorInput struct {
	Size         string
	Transmission string
	TempSensor   string
	Tank         string
	Notes        string
}

type RadFanAssyInput struct {
	Voltage          string
	MotorType        string
	Resistance       string
	NumberOfSockets  *int32
	Shroud           string
	ConnectorType    string
	FanBladeDiameter string
	NumberOfBlades   *int32
	Notes            string
}

type RadFanMotorInput struct {
	FanBladeDiameter string
	NumberOfBlades   *int32
	Voltage          string
	NumberOfSockets  *int32
	ConnectorType    string
	Notes            string
}

type ResistorInput struct {
	Type          string
	ConnectorType string
	Voltage       string
	Notes         string
}

type RotorInput struct {
	PulleyRibs        string
	PulleySize        string
	CompressorDetails string
	Notes             string
}

type StatorInput struct {
	Voltage           string
	CompressorDetails string
	Notes             string
}

type CompressorInput struct {
	CompressorID  string
	Oil           string
	Refrigerant   string
	Voltage       string
	PulleyRibs    string
	PulleySize    string
	PipeConnector string
	CompType      string
	CompMounting  string
	ConnectorType string
	Notes         string
}
