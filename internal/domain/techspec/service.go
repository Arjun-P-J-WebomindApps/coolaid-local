package techspec

import (
	"context"
)

// =======================================================
// TECH SPEC SERVICE
// =======================================================

type Service struct {
	DB DB
}

func NewService(db DB) *Service {
	return &Service{DB: db}
}

// =======================================================
// CREATE
// =======================================================

func (s *Service) CreateTechSpec(
	ctx context.Context,
	Q Queries,
	partNo string,
	in *TechnicalSpecsInput,
) error {

	if in == nil {
		return ErrInvalidTechSpec
	}
	if partNo == "" {
		return ErrInvalidPartNo
	}

	var err error

	switch in.Type {

	case "ACTUATOR":
		_, err = CreateActuator(ctx, Q, partNo, in.Actuator)

	case "BLOWER_MOTOR":
		_, err = CreateBlowerMotor(ctx, Q, partNo, in.BlowerMotor)

	case "CABIN_FILTER":
		_, err = CreateCabinFilter(ctx, Q, partNo, in.CabinFilter)

	case "CHILLER_UNIT":
		_, err = CreateChillerUnit(ctx, Q, partNo, in.ChillerUnit)

	case "CLUTCH_ASSY":
		_, err = CreateClutchAssy(ctx, Q, partNo, in.ClutchAssy)

	case "COMPRESSOR_VALVE":
		_, err = CreateCompressorValve(ctx, Q, partNo, in.CompressorValve)

	case "COND_FAN_ASSY":
		_, err = CreateCondFanAssy(ctx, Q, partNo, in.CondFanAssy)

	case "CONDENSER":
		_, err = CreateCondenser(ctx, Q, partNo, in.Condenser)

	case "EVAPORATOR":
		_, err = CreateEvaporator(ctx, Q, partNo, in.Evaporator)

	case "EXPANSION_VALVE":
		_, err = CreateExpansionValve(ctx, Q, partNo, in.ExpansionValve)

	case "FILTER_DRIER":
		_, err = CreateFilterDrier(ctx, Q, partNo, in.FilterDrier)

	case "HEATER_CORE":
		_, err = CreateHeaterCore(ctx, Q, partNo, in.HeaterCore)

	case "INTERCOOLER":
		_, err = CreateIntercooler(ctx, Q, partNo, in.Intercooler)

	case "PRESSURE_SWITCH":
		_, err = CreatePressureSwitch(ctx, Q, partNo, in.PressureSwitch)

	case "RADIATOR":
		_, err = CreateRadiator(ctx, Q, partNo, in.Radiator)

	case "RAD_FAN_ASSY":
		_, err = CreateRadFanAssy(ctx, Q, partNo, in.RadFanAssy)

	case "RAD_FAN_MOTOR":
		_, err = CreateRadFanMotor(ctx, Q, partNo, in.RadFanMotor)

	case "RESISTOR":
		_, err = CreateResistor(ctx, Q, partNo, in.Resistor)

	case "ROTOR":
		_, err = CreateRotor(ctx, Q, partNo, in.Rotor)

	case "STATOR":
		_, err = CreateStator(ctx, Q, partNo, in.Stator)

	case "COMPRESSOR":
		_, err = CreateCompressor(ctx, Q, partNo, in.Compressor)

	default:
		return ErrUnsupportedType
	}

	if err != nil {
		return err
	}

	return err
}

// =======================================================
// UPDATE
// =======================================================

func (s *Service) UpdateTechSpec(
	ctx context.Context,
	Q Queries,
	partNo string,
	in *TechnicalSpecsInput,
) error {

	if in == nil {
		return ErrInvalidTechSpec
	}
	if partNo == "" {
		return ErrInvalidPartNo
	}

	var err error

	switch in.Type {

	case "ACTUATOR":
		_, err = UpdateActuator(ctx, Q, partNo, in.Actuator)

	case "BLOWER_MOTOR":
		_, err = UpdateBlowerMotor(ctx, Q, partNo, in.BlowerMotor)

	case "CABIN_FILTER":
		_, err = UpdateCabinFilter(ctx, Q, partNo, in.CabinFilter)

	case "CHILLER_UNIT":
		_, err = UpdateChillerUnit(ctx, Q, partNo, in.ChillerUnit)

	case "CLUTCH_ASSY":
		_, err = UpdateClutchAssy(ctx, Q, partNo, in.ClutchAssy)

	case "COMPRESSOR_VALVE":
		_, err = UpdateCompressorValve(ctx, Q, partNo, in.CompressorValve)

	case "COND_FAN_ASSY":
		_, err = UpdateCondFanAssy(ctx, Q, partNo, in.CondFanAssy)

	case "CONDENSER":
		_, err = UpdateCondenser(ctx, Q, partNo, in.Condenser)

	case "EVAPORATOR":
		_, err = UpdateEvaporator(ctx, Q, partNo, in.Evaporator)

	case "EXPANSION_VALVE":
		_, err = UpdateExpansionValve(ctx, Q, partNo, in.ExpansionValve)

	case "FILTER_DRIER":
		_, err = UpdateFilterDrier(ctx, Q, partNo, in.FilterDrier)

	case "HEATER_CORE":
		_, err = UpdateHeaterCore(ctx, Q, partNo, in.HeaterCore)

	case "INTERCOOLER":
		_, err = UpdateIntercooler(ctx, Q, partNo, in.Intercooler)

	case "PRESSURE_SWITCH":
		_, err = UpdatePressureSwitch(ctx, Q, partNo, in.PressureSwitch)

	case "RADIATOR":
		_, err = UpdateRadiator(ctx, Q, partNo, in.Radiator)

	case "RAD_FAN_ASSY":
		_, err = UpdateRadFanAssy(ctx, Q, partNo, in.RadFanAssy)

	case "RAD_FAN_MOTOR":
		_, err = UpdateRadFanMotor(ctx, Q, partNo, in.RadFanMotor)

	case "RESISTOR":
		_, err = UpdateResistor(ctx, Q, partNo, in.Resistor)

	case "ROTOR":
		_, err = UpdateRotor(ctx, Q, partNo, in.Rotor)

	case "STATOR":
		_, err = UpdateStator(ctx, Q, partNo, in.Stator)

	case "COMPRESSOR":
		_, err = UpdateCompressor(ctx, Q, partNo, in.Compressor)

	default:
		return ErrUnsupportedType
	}

	return err
}

// =======================================================
// DELETE
// =======================================================

func (s *Service) DeleteTechSpec(
	ctx context.Context,
	Q Queries,
	partNo string,
	techType string,
) error {

	if partNo == "" {
		return ErrInvalidPartNo
	}

	var err error

	switch techType {

	case "ACTUATOR":
		err = DeleteActuator(ctx, Q, partNo)

	case "BLOWER_MOTOR":
		err = DeleteBlowerMotor(ctx, Q, partNo)

	case "CABIN_FILTER":
		err = DeleteCabinFilter(ctx, Q, partNo)

	case "CHILLER_UNIT":
		err = DeleteChillerUnit(ctx, Q, partNo)

	case "CLUTCH_ASSY":
		err = DeleteClutchAssy(ctx, Q, partNo)

	case "COMPRESSOR_VALVE":
		err = DeleteCompressorValve(ctx, Q, partNo)

	case "COND_FAN_ASSY":
		err = DeleteCondFanAssy(ctx, Q, partNo)

	case "CONDENSER":
		err = DeleteCondenser(ctx, Q, partNo)

	case "EVAPORATOR":
		err = DeleteEvaporator(ctx, Q, partNo)

	case "EXPANSION_VALVE":
		err = DeleteExpansionValve(ctx, Q, partNo)

	case "FILTER_DRIER":
		err = DeleteFilterDrier(ctx, Q, partNo)

	case "HEATER_CORE":
		err = DeleteHeaterCore(ctx, Q, partNo)

	case "INTERCOOLER":
		err = DeleteIntercooler(ctx, Q, partNo)

	case "PRESSURE_SWITCH":
		err = DeletePressureSwitch(ctx, Q, partNo)

	case "RADIATOR":
		err = DeleteRadiator(ctx, Q, partNo)

	case "RAD_FAN_ASSY":
		err = DeleteRadFanAssy(ctx, Q, partNo)

	case "RAD_FAN_MOTOR":
		err = DeleteRadFanMotor(ctx, Q, partNo)

	case "RESISTOR":
		err = DeleteResistor(ctx, Q, partNo)

	case "ROTOR":
		err = DeleteRotor(ctx, Q, partNo)

	case "STATOR":
		err = DeleteStator(ctx, Q, partNo)

	case "COMPRESSOR":
		err = DeleteCompressor(ctx, Q, partNo)

	default:
		return ErrUnsupportedType
	}

	return err
}
