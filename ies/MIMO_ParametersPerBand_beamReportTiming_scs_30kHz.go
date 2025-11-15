package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym4  uper.Enumerated = 0
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym8  uper.Enumerated = 1
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym14 uper.Enumerated = 2
	MIMO_ParametersPerBand_beamReportTiming_scs_30kHz_Enum_sym28 uper.Enumerated = 3
)

type MIMO_ParametersPerBand_beamReportTiming_scs_30kHz struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MIMO_ParametersPerBand_beamReportTiming_scs_30kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_beamReportTiming_scs_30kHz", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamReportTiming_scs_30kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_beamReportTiming_scs_30kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
