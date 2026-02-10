package techspec

import (
	"context"
)

// =======================================================
// TECH SPEC SERVICE
// =======================================================

type Service struct {
	db DB
}

func NewService(db DB) *Service {
	return &Service{db: db}
}

// =======================================================
// CREATE
// =======================================================

func (s *Service) CreateTechSpec(
	ctx context.Context,
	partNo string,
	in *TechnicalSpecsInput,
) error {

	if in == nil {
		return ErrInvalidTechSpec
	}
	if partNo == "" {
		return ErrInvalidPartNo
	}

	tx, q, err := s.db.BeginTx(ctx)
	if err != nil {
		return ErrInternal
	}
	defer tx.Rollback()

	switch in.Type {

	case "ACTUATOR":
		_, err = CreateActuator(ctx, q, partNo, in.Actuator)

	case "BLOWER_MOTOR":
		_, err = CreateBlowerMotor(ctx, q, partNo, in.BlowerMotor)

	case "CABIN_FILTER":
		_, err = CreateCabinFilter(ctx, q, partNo, in.CabinFilter)

	case "CHILLER_UNIT":
		_, err = CreateChillerUnit(ctx, q, partNo, in.ChillerUnit)

	case "CLUTCH_ASSY":
		_, err = CreateClutchAssy(ctx, q, partNo, in.ClutchAssy)

	case "COMPRESSOR_VALVE":
		_, err = CreateCompressorValve(ctx, q, partNo, in.CompressorValve)

	case "COND_FAN_ASSY":
		_, err = CreateCondFanAssy(ctx, q, partNo, in.CondFanAssy)

	case "CONDENSER":
		_, err = CreateCondenser(ctx, q, partNo, in.Condenser)

	case "EVAPORATOR":
		_, err = CreateEvaporator(ctx, q, partNo, in.Evaporator)

	case "EXPANSION_VALVE":
		_, err = CreateExpansionValve(ctx, q, partNo, in.ExpansionValve)

	case "FILTER_DRIER":
		_, err = CreateFilterDrier(ctx, q, partNo, in.FilterDrier)

	case "HEATER_CORE":
		_, err = CreateHeaterCore(ctx, q, partNo, in.HeaterCore)

	case "INTERCOOLER":
		_, err = CreateIntercooler(ctx, q, partNo, in.Intercooler)

	case "PRESSURE_SWITCH":
		_, err = CreatePressureSwitch(ctx, q, partNo, in.PressureSwitch)

	case "RADIATOR":
		_, err = CreateRadiator(ctx, q, partNo, in.Radiator)

	case "RAD_FAN_ASSY":
		_, err = CreateRadFanAssy(ctx, q, partNo, in.RadFanAssy)

	case "RAD_FAN_MOTOR":
		_, err = CreateRadFanMotor(ctx, q, partNo, in.RadFanMotor)

	case "RESISTOR":
		_, err = CreateResistor(ctx, q, partNo, in.Resistor)

	case "ROTOR":
		_, err = CreateRotor(ctx, q, partNo, in.Rotor)

	case "STATOR":
		_, err = CreateStator(ctx, q, partNo, in.Stator)

	case "COMPRESSOR":
		_, err = CreateCompressor(ctx, q, partNo, in.Compressor)

	default:
		return ErrUnsupportedType
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}

// =======================================================
// UPDATE
// =======================================================

func (s *Service) UpdateTechSpec(
	ctx context.Context,
	partNo string,
	in *TechnicalSpecsInput,
) error {

	if in == nil {
		return ErrInvalidTechSpec
	}
	if partNo == "" {
		return ErrInvalidPartNo
	}

	tx, q, err := s.db.BeginTx(ctx)
	if err != nil {
		return ErrInternal
	}
	defer tx.Rollback()

	switch in.Type {

	case "ACTUATOR":
		_, err = UpdateActuator(ctx, q, partNo, in.Actuator)

	case "BLOWER_MOTOR":
		_, err = UpdateBlowerMotor(ctx, q, partNo, in.BlowerMotor)

	case "CABIN_FILTER":
		_, err = UpdateCabinFilter(ctx, q, partNo, in.CabinFilter)

	case "CHILLER_UNIT":
		_, err = UpdateChillerUnit(ctx, q, partNo, in.ChillerUnit)

	case "CLUTCH_ASSY":
		_, err = UpdateClutchAssy(ctx, q, partNo, in.ClutchAssy)

	case "COMPRESSOR_VALVE":
		_, err = UpdateCompressorValve(ctx, q, partNo, in.CompressorValve)

	case "COND_FAN_ASSY":
		_, err = UpdateCondFanAssy(ctx, q, partNo, in.CondFanAssy)

	case "CONDENSER":
		_, err = UpdateCondenser(ctx, q, partNo, in.Condenser)

	case "EVAPORATOR":
		_, err = UpdateEvaporator(ctx, q, partNo, in.Evaporator)

	case "EXPANSION_VALVE":
		_, err = UpdateExpansionValve(ctx, q, partNo, in.ExpansionValve)

	case "FILTER_DRIER":
		_, err = UpdateFilterDrier(ctx, q, partNo, in.FilterDrier)

	case "HEATER_CORE":
		_, err = UpdateHeaterCore(ctx, q, partNo, in.HeaterCore)

	case "INTERCOOLER":
		_, err = UpdateIntercooler(ctx, q, partNo, in.Intercooler)

	case "PRESSURE_SWITCH":
		_, err = UpdatePressureSwitch(ctx, q, partNo, in.PressureSwitch)

	case "RADIATOR":
		_, err = UpdateRadiator(ctx, q, partNo, in.Radiator)

	case "RAD_FAN_ASSY":
		_, err = UpdateRadFanAssy(ctx, q, partNo, in.RadFanAssy)

	case "RAD_FAN_MOTOR":
		_, err = UpdateRadFanMotor(ctx, q, partNo, in.RadFanMotor)

	case "RESISTOR":
		_, err = UpdateResistor(ctx, q, partNo, in.Resistor)

	case "ROTOR":
		_, err = UpdateRotor(ctx, q, partNo, in.Rotor)

	case "STATOR":
		_, err = UpdateStator(ctx, q, partNo, in.Stator)

	case "COMPRESSOR":
		_, err = UpdateCompressor(ctx, q, partNo, in.Compressor)

	default:
		return ErrUnsupportedType
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}

// =======================================================
// DELETE
// =======================================================

func (s *Service) DeleteTechSpec(
	ctx context.Context,
	partNo string,
	techType string,
) error {

	if partNo == "" {
		return ErrInvalidPartNo
	}

	tx, q, err := s.db.BeginTx(ctx)
	if err != nil {
		return ErrInternal
	}
	defer tx.Rollback()

	switch techType {

	case "ACTUATOR":
		err = DeleteActuator(ctx, q, partNo)

	case "BLOWER_MOTOR":
		err = DeleteBlowerMotor(ctx, q, partNo)

	case "CABIN_FILTER":
		err = DeleteCabinFilter(ctx, q, partNo)

	case "CHILLER_UNIT":
		err = DeleteChillerUnit(ctx, q, partNo)

	case "CLUTCH_ASSY":
		err = DeleteClutchAssy(ctx, q, partNo)

	case "COMPRESSOR_VALVE":
		err = DeleteCompressorValve(ctx, q, partNo)

	case "COND_FAN_ASSY":
		err = DeleteCondFanAssy(ctx, q, partNo)

	case "CONDENSER":
		err = DeleteCondenser(ctx, q, partNo)

	case "EVAPORATOR":
		err = DeleteEvaporator(ctx, q, partNo)

	case "EXPANSION_VALVE":
		err = DeleteExpansionValve(ctx, q, partNo)

	case "FILTER_DRIER":
		err = DeleteFilterDrier(ctx, q, partNo)

	case "HEATER_CORE":
		err = DeleteHeaterCore(ctx, q, partNo)

	case "INTERCOOLER":
		err = DeleteIntercooler(ctx, q, partNo)

	case "PRESSURE_SWITCH":
		err = DeletePressureSwitch(ctx, q, partNo)

	case "RADIATOR":
		err = DeleteRadiator(ctx, q, partNo)

	case "RAD_FAN_ASSY":
		err = DeleteRadFanAssy(ctx, q, partNo)

	case "RAD_FAN_MOTOR":
		err = DeleteRadFanMotor(ctx, q, partNo)

	case "RESISTOR":
		err = DeleteResistor(ctx, q, partNo)

	case "ROTOR":
		err = DeleteRotor(ctx, q, partNo)

	case "STATOR":
		err = DeleteStator(ctx, q, partNo)

	case "COMPRESSOR":
		err = DeleteCompressor(ctx, q, partNo)

	default:
		return ErrUnsupportedType
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}
