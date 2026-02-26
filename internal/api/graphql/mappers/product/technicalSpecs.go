package product_mapper

import (
	"github.com/webomindapps-dev/coolaid-backend/internal/domain/techspec"
	"github.com/webomindapps-dev/coolaid-backend/internal/generated/graphql/model"
)

func MapTechnicalSpecsInput(
	in *model.ProductCategoryData,
) techspec.TechnicalSpecsInput {

	// default (empty)
	out := techspec.TechnicalSpecsInput{}

	if in == nil {
		return out
	}

	out.Type = in.Type

	switch in.Type {

	case "ACTUATOR":
		if in.Actuator == nil {
			return out
		}

		out.Actuator = &techspec.ActuatorInput{
			ConnectorType: in.Actuator.ConnectorType,
			Mounting:      in.Actuator.Mounting,
			Voltage:       in.Actuator.Voltage,
			RotationAngle: in.Actuator.RotationAngle,
			Notes:         in.Actuator.Notes,
		}

	case "BLOWER_MOTOR":
		if in.BlowerMotor == nil {
			return out
		}

		out.BlowerMotor = &techspec.BlowerMotorInput{
			Mounting:      in.BlowerMotor.Mounting,
			ConnectorType: in.BlowerMotor.ConnectorType,
			Impeller:      in.BlowerMotor.Impeller,
			Resistance:    in.BlowerMotor.Resistance,
			MotorMounting: in.BlowerMotor.MotorMounting,
			MotorType:     in.BlowerMotor.MotorType,
			Voltage:       in.BlowerMotor.Voltage,
			Notes:         in.BlowerMotor.Notes,
		}

	case "CABIN_FILTER":
		if in.CabinFilter == nil {
			return out
		}

		out.CabinFilter = &techspec.CabinFilterInput{
			Type:       in.CabinFilter.Type,
			Dimensions: in.CabinFilter.Dimensions,
			Material:   in.CabinFilter.Material,
			Notes:      in.CabinFilter.Notes,
		}

	case "CHILLER_UNIT":
		if in.ChillerUnit == nil {
			return out
		}

		out.ChillerUnit = &techspec.ChillerUnitInput{
			Type:    in.ChillerUnit.Type,
			Voltage: in.ChillerUnit.Voltage,
			Notes:   in.ChillerUnit.Notes,
		}

	case "CLUTCH_ASSY":
		if in.ClutchAssy == nil {
			return out
		}

		out.ClutchAssy = &techspec.ClutchAssyInput{
			PulleyRibs:        in.ClutchAssy.PulleyRibs,
			PulleySize:        in.ClutchAssy.PulleySize,
			CompressorDetails: in.ClutchAssy.CompressorDetails,
			ConnectorType:     in.ClutchAssy.ConnectorType,
			Voltage:           in.ClutchAssy.Voltage,
			ShaftType:         in.ClutchAssy.ShaftType,
			Notes:             in.ClutchAssy.Notes,
		}

	case "COMPRESSOR_VALVE":
		if in.CompressorValve == nil {
			return out
		}

		out.CompressorValve = &techspec.CompressorValveInput{
			Type:              in.CompressorValve.Type,
			Voltage:           in.CompressorValve.Voltage,
			ConnectorType:     in.CompressorValve.ConnectorType,
			CompressorDetails: in.CompressorValve.CompressorDetails,
			Notes:             in.CompressorValve.Notes,
		}

	case "COND_FAN_ASSY":
		if in.CondFanAssy == nil {
			return out
		}

		out.CondFanAssy = &techspec.CondFanAssyInput{
			Voltage:          in.CondFanAssy.Voltage,
			MotorType:        in.CondFanAssy.MotorType,
			Resistance:       in.CondFanAssy.Resistance,
			FanBladeDiameter: in.CondFanAssy.FanBladeDiameter,
			NumberOfBlades:   in.CondFanAssy.NumberOfBlades,
			Shroud:           in.CondFanAssy.Shroud,
			ConnectorType:    in.CondFanAssy.ConnectorType,
			Notes:            in.CondFanAssy.Notes,
		}

	case "CONDENSER":
		if in.Condenser == nil {
			return out
		}

		out.Condenser = &techspec.CondenserInput{
			Size:           in.Condenser.Size,
			PipeConnector:  in.Condenser.PipeConnector,
			Drier:          in.Condenser.Drier,
			PressureSwitch: in.Condenser.PressureSwitch,
			Notes:          in.Condenser.Notes,
		}

	case "EVAPORATOR":
		if in.Evaporator == nil {
			return out
		}

		out.Evaporator = &techspec.EvaporatorInput{
			Mounting:       in.Evaporator.Mounting,
			ExpValve:       in.Evaporator.ExpValve,
			AdditionalInfo: in.Evaporator.AdditionalInfo,
			Dimensions:     in.Evaporator.Dimensions,
			PipeConnector:  in.Evaporator.PipeConnector,
			Notes:          in.Evaporator.Notes,
		}

	case "EXPANSION_VALVE":
		if in.ExpansionValve == nil {
			return out
		}

		out.ExpansionValve = &techspec.ExpansionValveInput{
			Type:        in.ExpansionValve.Type,
			Material:    in.ExpansionValve.Material,
			Refrigerant: in.ExpansionValve.Refrigerant,
			Notes:       in.ExpansionValve.Notes,
		}

	case "FILTER_DRIER":
		if in.FilterDrier == nil {
			return out
		}

		out.FilterDrier = &techspec.FilterDrierInput{
			PipeConnector:  in.FilterDrier.PipeConnector,
			Size:           in.FilterDrier.Size,
			PressureSwitch: in.FilterDrier.PressureSwitch,
			Notes:          in.FilterDrier.Notes,
		}

	case "HEATER_CORE":
		if in.HeaterCore == nil {
			return out
		}

		out.HeaterCore = &techspec.HeaterCoreInput{
			Size:  in.HeaterCore.Size,
			Pipe:  in.HeaterCore.Pipe,
			Type:  in.HeaterCore.Type,
			Notes: in.HeaterCore.Notes,
		}

	case "INTERCOOLER":
		if in.Intercooler == nil {
			return out
		}

		out.Intercooler = &techspec.IntercoolerInput{
			Size:       in.Intercooler.Size,
			TempSensor: in.Intercooler.TempSensor,
			Notes:      in.Intercooler.Notes,
		}

	case "PRESSURE_SWITCH":
		if in.PressureSwitch == nil {
			return out
		}

		out.PressureSwitch = &techspec.PressureSwitchInput{
			ConnectorType: in.PressureSwitch.ConnectorType,
			ThreadType:    in.PressureSwitch.ThreadType,
			Notes:         in.PressureSwitch.Notes,
		}

	case "RADIATOR":
		if in.Radiator == nil {
			return out
		}

		out.Radiator = &techspec.RadiatorInput{
			Size:         in.Radiator.Size,
			Transmission: in.Radiator.Transmission,
			TempSensor:   in.Radiator.TempSensor,
			Tank:         in.Radiator.Tank,
			Notes:        in.Radiator.Notes,
		}

	case "RAD_FAN_ASSY":
		if in.RadFanAssy == nil {
			return out
		}

		out.RadFanAssy = &techspec.RadFanAssyInput{
			Voltage:          in.RadFanAssy.Voltage,
			MotorType:        in.RadFanAssy.MotorType,
			Resistance:       in.RadFanAssy.Resistance,
			NumberOfSockets:  in.RadFanAssy.NumberOfSockets,
			Shroud:           in.RadFanAssy.Shroud,
			ConnectorType:    in.RadFanAssy.ConnectorType,
			FanBladeDiameter: in.RadFanAssy.FanBladeDiameter,
			NumberOfBlades:   in.RadFanAssy.NumberOfBlades,
			Notes:            in.RadFanAssy.Notes,
		}

	case "RAD_FAN_MOTOR":
		if in.RadFanMotor == nil {
			return out
		}

		out.RadFanMotor = &techspec.RadFanMotorInput{
			FanBladeDiameter: in.RadFanMotor.FanBladeDiameter,
			NumberOfBlades:   in.RadFanMotor.NumberOfBlades,
			Voltage:          in.RadFanMotor.Voltage,
			NumberOfSockets:  in.RadFanMotor.NumberOfSockets,
			ConnectorType:    in.RadFanMotor.ConnectorType,
			Notes:            in.RadFanMotor.Notes,
		}

	case "RESISTOR":
		if in.Resistor == nil {
			return out
		}

		out.Resistor = &techspec.ResistorInput{
			Type:          in.Resistor.Type,
			ConnectorType: in.Resistor.ConnectorType,
			Voltage:       in.Resistor.Voltage,
			Notes:         in.Resistor.Notes,
		}

	case "ROTOR":
		if in.Rotor == nil {
			return out
		}

		out.Rotor = &techspec.RotorInput{
			PulleyRibs:        in.Rotor.PulleyRibs,
			PulleySize:        in.Rotor.PulleySize,
			CompressorDetails: in.Rotor.CompressorDetails,
			Notes:             in.Rotor.Notes,
		}

	case "STATOR":
		if in.Stator == nil {
			return out
		}

		out.Stator = &techspec.StatorInput{
			Voltage:           in.Stator.Voltage,
			CompressorDetails: in.Stator.CompressorDetails,
			Notes:             in.Stator.Notes,
		}

	case "COMPRESSOR":
		if in.Compressor == nil {
			return out
		}

		out.Compressor = &techspec.CompressorInput{
			CompressorID:  in.Compressor.CompressorID,
			Oil:           in.Compressor.Oil,
			Refrigerant:   in.Compressor.Refrigerant,
			Voltage:       in.Compressor.Voltage,
			PulleyRibs:    in.Compressor.PulleyRibs,
			PulleySize:    in.Compressor.PulleySize,
			PipeConnector: in.Compressor.PipeConnector,
			CompType:      in.Compressor.CompType,
			CompMounting:  in.Compressor.CompMounting,
			ConnectorType: in.Compressor.ConnectorType,
			Notes:         in.Compressor.Notes,
		}
	}

	return out
}
