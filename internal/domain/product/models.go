package techspec

// =======================================================
// DOMAIN MODELS
// =======================================================
//
// These are domain-level models used by services / handlers.
// They are NOT DB rows and NOT sqlc models.
// DB rows are mapped into these if needed.
//

type Actuator struct {
	PartNo        string
	ConnectorType string
	Mounting      string
	Voltage       string
	RotationAngle string
	Notes         string
}

type BlowerMotor struct {
	PartNo        string
	Mounting      string
	ConnectorType string
	Impeller      string
	Resistance    string
	MotorMounting string
	MotorType     string
	Voltage       string
	Notes         string
}

type CabinFilter struct {
	PartNo     string
	Type       string
	Dimensions string
	Material   string
	Notes      string
}

type ChillerUnit struct {
	PartNo  string
	Type    string
	Voltage string
	Notes   string
}

type ClutchAssy struct {
	PartNo            string
	PulleyRibs        string
	PulleySize        string
	CompressorDetails string
	ConnectorType     string
	Voltage           string
	ShaftType         string
	Notes             string
}

type CompressorValve struct {
	PartNo            string
	Type              string
	Voltage           string
	ConnectorType     string
	CompressorDetails string
	Notes             string
}

type CondFanAssy struct {
	PartNo           string
	Voltage          string
	MotorType        string
	Resistance       string
	FanBladeDiameter string
	NumberOfBlades   *int32
	Shroud           string
	ConnectorType    string
	Notes            string
}

type Condenser struct {
	PartNo         string
	Size           string
	PipeConnector  string
	Drier          string
	PressureSwitch string
	Notes          string
}

type Evaporator struct {
	PartNo         string
	Mounting       string
	ExpValve       string
	AdditionalInfo string
	Dimensions     string
	PipeConnector  string
	Notes          string
}

type ExpansionValve struct {
	PartNo      string
	Type        string
	Material    string
	Refrigerant string
	Notes       string
}

type FilterDrier struct {
	PartNo         string
	PipeConnector  string
	Size           string
	PressureSwitch string
	Notes          string
}

type HeaterCore struct {
	PartNo string
	Size   string
	Pipe   string
	Type   string
	Notes  string
}

type Intercooler struct {
	PartNo     string
	Size       string
	TempSensor string
	Notes      string
}

type PressureSwitch struct {
	PartNo        string
	ConnectorType string
	ThreadType    string
	Notes         string
}

type Radiator struct {
	PartNo       string
	Size         string
	Transmission string
	TempSensor   string
	Tank         string
	Notes        string
}

type RadFanAssy struct {
	PartNo           string
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

type RadFanMotor struct {
	PartNo           string
	FanBladeDiameter string
	NumberOfBlades   *int32
	Voltage          string
	NumberOfSockets  *int32
	ConnectorType    string
	Notes            string
}

type Resistor struct {
	PartNo        string
	Type          string
	ConnectorType string
	Voltage       string
	Notes         string
}

type Rotor struct {
	PartNo            string
	PulleyRibs        string
	PulleySize        string
	CompressorDetails string
	Notes             string
}

type Stator struct {
	PartNo            string
	Voltage           string
	CompressorDetails string
	Notes             string
}

type Compressor struct {
	PartNo        string
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
