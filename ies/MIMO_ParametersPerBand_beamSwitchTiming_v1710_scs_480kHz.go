package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz_Enum_sym56   uper.Enumerated = 0
	MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz_Enum_sym112  uper.Enumerated = 1
	MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz_Enum_sym192  uper.Enumerated = 2
	MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz_Enum_sym896  uper.Enumerated = 3
	MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz_Enum_sym1344 uper.Enumerated = 4
)

type MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
