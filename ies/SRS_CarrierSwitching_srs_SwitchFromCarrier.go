package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_CarrierSwitching_srs_SwitchFromCarrier_Enum_sUL uper.Enumerated = 0
	SRS_CarrierSwitching_srs_SwitchFromCarrier_Enum_nUL uper.Enumerated = 1
)

type SRS_CarrierSwitching_srs_SwitchFromCarrier struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SRS_CarrierSwitching_srs_SwitchFromCarrier) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SRS_CarrierSwitching_srs_SwitchFromCarrier", err)
	}
	return nil
}

func (ie *SRS_CarrierSwitching_srs_SwitchFromCarrier) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SRS_CarrierSwitching_srs_SwitchFromCarrier", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
