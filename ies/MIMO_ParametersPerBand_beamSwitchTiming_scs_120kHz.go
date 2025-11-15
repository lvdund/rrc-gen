package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym14  uper.Enumerated = 0
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym28  uper.Enumerated = 1
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym48  uper.Enumerated = 2
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym224 uper.Enumerated = 3
	MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz_Enum_sym336 uper.Enumerated = 4
)

type MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_beamSwitchTiming_scs_120kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
