package techspec

/*
Domain-level representations of DB rows.
These are NOT sqlc / ORM models.
The repository layer adapts DB rows into these.
*/

type ActuatorRow struct {
	ID            string
	PartNo        string
	ConnectorType *string
	Mounting      *string
	Voltage       *string
	RotationAngle *string
	Notes         *string
}

type BlowerMotorRow struct {
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

type CabinFilterRow struct {
	ID         string
	PartNo     string
	Type       *string
	Dimensions *string
	Material   *string
	Notes      *string
}

type ChillerUnitRow struct {
	ID      string
	PartNo  string
	Type    *string
	Voltage *string
	Notes   *string
}

type ClutchAssyRow struct {
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

type CompressorValveRow struct {
	ID                string
	PartNo            string
	Type              *string
	Voltage           *string
	ConnectorType     *string
	CompressorDetails *string
	Notes             *string
}

type CondFanAssyRow struct {
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

type CondenserRow struct {
	ID             string
	PartNo         string
	Size           *string
	PipeConnector  *string
	Drier          *string
	PressureSwitch *string
	Notes          *string
}

type EvaporatorRow struct {
	ID             string
	PartNo         string
	Mounting       *string
	ExpValve       *string
	AdditionalInfo *string
	Dimensions     *string
	PipeConnector  *string
	Notes          *string
}

type ExpansionValveRow struct {
	ID          string
	PartNo      string
	Type        *string
	Material    *string
	Refrigerant *string
	Notes       *string
}

type FilterDrierRow struct {
	ID             string
	PartNo         string
	PipeConnector  *string
	Size           *string
	PressureSwitch *string
	Notes          *string
}

type HeaterCoreRow struct {
	ID     string
	PartNo string
	Size   *string
	Pipe   *string
	Type   *string
	Notes  *string
}

type IntercoolerRow struct {
	ID         string
	PartNo     string
	Size       *string
	TempSensor *string
	Notes      *string
}

type PressureSwitchRow struct {
	ID            string
	PartNo        string
	ConnectorType *string
	ThreadType    *string
	Notes         *string
}

type RadiatorRow struct {
	ID           string
	PartNo       string
	Size         *string
	Transmission *string
	TempSensor   *string
	Tank         *string
	Notes        *string
}

type RadFanAssyRow struct {
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

type RadFanMotorRow struct {
	ID               string
	PartNo           string
	FanBladeDiameter *string
	NumberOfBlades   *int32
	Voltage          *string
	NumberOfSockets  *int32
	ConnectorType    *string
	Notes            *string
}

type ResistorRow struct {
	ID            string
	PartNo        string
	Type          *string
	ConnectorType *string
	Voltage       *string
	Notes         *string
}

type RotorRow struct {
	ID                string
	PartNo            string
	PulleyRibs        *string
	PulleySize        *string
	CompressorDetails *string
	Notes             *string
}

type StatorRow struct {
	ID                string
	PartNo            string
	Voltage           *string
	CompressorDetails *string
	Notes             *string
}

type CompressorRow struct {
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
