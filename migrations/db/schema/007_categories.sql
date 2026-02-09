-- +goose Up
-- +goose StatementBegin

-- 1. Compressors
CREATE TABLE IF NOT EXISTS compressors (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    compressor_id TEXT,
    oil TEXT,
    refrigerant TEXT,
    voltage TEXT,
    pulley_ribs TEXT,
    pulley_size TEXT,
    pipe_connector TEXT,
    comp_type TEXT,
    comp_mounting TEXT,
    connector_type TEXT,
    notes TEXT
);

-- 2. Condensers
CREATE TABLE IF NOT EXISTS condensers (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    size TEXT,
    pipe_connector TEXT,
    drier TEXT,
    pressure_switch TEXT,
    notes TEXT
);

-- 3. Actuators
CREATE TABLE IF NOT EXISTS actuators (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    connector_type TEXT,
    mounting TEXT,
    voltage TEXT,
    rotation_angle TEXT,
    notes TEXT
);

-- 4. Blower Motors
CREATE TABLE IF NOT EXISTS blower_motors (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    mounting TEXT,
    connector_type TEXT,
    impeller TEXT,
    resistance TEXT,
    motor_mounting TEXT,
    motor_type TEXT,
    voltage TEXT,
    notes TEXT
);

-- 5. Cabin Filters
CREATE TABLE IF NOT EXISTS cabin_filters (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    type TEXT,
    dimensions TEXT,
    material TEXT,
    notes TEXT
);

-- 6. Chiller Units
CREATE TABLE IF NOT EXISTS chiller_units (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    type TEXT,
    voltage TEXT,
    notes TEXT
);

-- 7. Clutch Assys
CREATE TABLE IF NOT EXISTS clutch_assys (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    pulley_ribs TEXT,
    pulley_size TEXT,
    compressor_details TEXT,
    connector_type TEXT,
    voltage TEXT,
    shaft_type TEXT,
    notes TEXT
);

-- 8. Compressor Valves
CREATE TABLE IF NOT EXISTS compressor_valves (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    type TEXT,
    voltage TEXT,
    connector_type TEXT,
    compressor_details TEXT,
    notes TEXT
);

-- 9. Cond Fan Assys
CREATE TABLE IF NOT EXISTS cond_fan_assys (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    voltage TEXT,
    motor_type TEXT,
    resistance TEXT,
    fan_blade_diameter TEXT,
    number_of_blades INT,
    shroud TEXT,
    connector_type TEXT,
    notes TEXT
);

-- 10. Evaporators
CREATE TABLE IF NOT EXISTS evaporators (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    mounting TEXT,
    exp_valve TEXT,
    additional_info TEXT,
    dimensions TEXT,
    pipe_connector TEXT,
    notes TEXT
);

-- 11. Expansion Valves
CREATE TABLE IF NOT EXISTS expansion_valves (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    type TEXT,
    material TEXT,
    refrigerant TEXT,
    notes TEXT
);

-- 12. Filter Driers
CREATE TABLE IF NOT EXISTS filter_driers (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    pipe_connector TEXT,
    size TEXT,
    pressure_switch TEXT,
    notes TEXT
);

-- 13. Heater Cores
CREATE TABLE IF NOT EXISTS heater_cores (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    size TEXT,
    pipe TEXT,
    type TEXT,
    notes TEXT
);

-- 14. Intercoolers
CREATE TABLE IF NOT EXISTS intercoolers (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    size TEXT,
    temp_sensor TEXT,
    notes TEXT
);

-- 15. Pressure Switches
CREATE TABLE IF NOT EXISTS pressure_switches (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    connector_type TEXT,
    thread_type TEXT,
    notes TEXT
);

-- 16. Radiators
CREATE TABLE IF NOT EXISTS radiators (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    size TEXT,
    transmission TEXT,
    temp_sensor TEXT,
    tank TEXT,
    notes TEXT
);

-- 17. Rad Fan Assys
CREATE TABLE IF NOT EXISTS rad_fan_assys (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    voltage TEXT,
    motor_type TEXT,
    resistance TEXT,
    number_of_sockets INT,
    shroud TEXT,
    connector_type TEXT,
    fan_blade_diameter TEXT,
    number_of_blades INT,
    notes TEXT
);

-- 18. Rad Fan Motors
CREATE TABLE IF NOT EXISTS rad_fan_motors (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    fan_blade_diameter TEXT,
    number_of_blades INT,
    voltage TEXT,
    number_of_sockets INT,
    connector_type TEXT,
    notes TEXT
);

-- 19. Resistors
CREATE TABLE IF NOT EXISTS resistors (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    type TEXT,
    connector_type TEXT,
    voltage TEXT,
    notes TEXT
);

-- 20. Rotors
CREATE TABLE IF NOT EXISTS rotors (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    pulley_ribs TEXT,
    pulley_size TEXT,
    compressor_details TEXT,
    notes TEXT
);

-- 21. Stators
CREATE TABLE IF NOT EXISTS stators (
    id UUID PRIMARY KEY,
    part_no TEXT NOT NULL UNIQUE,
    voltage TEXT,
    compressor_details TEXT,
    notes TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS stators;
DROP TABLE IF EXISTS rotors;
DROP TABLE IF EXISTS resistors;
DROP TABLE IF EXISTS rad_fan_motors;
DROP TABLE IF EXISTS rad_fan_assys;
DROP TABLE IF EXISTS radiators;
DROP TABLE IF EXISTS pressure_switches;
DROP TABLE IF EXISTS intercoolers;
DROP TABLE IF EXISTS heater_cores;
DROP TABLE IF EXISTS filter_driers;
DROP TABLE IF EXISTS expansion_valves;
DROP TABLE IF EXISTS evaporators;
DROP TABLE IF EXISTS cond_fan_assys;
DROP TABLE IF EXISTS compressor_valves;
DROP TABLE IF EXISTS clutch_assys;
DROP TABLE IF EXISTS chiller_units;
DROP TABLE IF EXISTS cabin_filters;
DROP TABLE IF EXISTS blower_motors;
DROP TABLE IF EXISTS actuators;
DROP TABLE IF EXISTS condensers;
DROP TABLE IF EXISTS compressors;

-- +goose StatementEnd
