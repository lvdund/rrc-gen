package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RF_ParametersMRDC_srs_SwitchingTimeRequested_Enum_true uper.Enumerated = 0
)

type RF_ParametersMRDC_srs_SwitchingTimeRequested struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RF_ParametersMRDC_srs_SwitchingTimeRequested) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RF_ParametersMRDC_srs_SwitchingTimeRequested", err)
	}
	return nil
}

func (ie *RF_ParametersMRDC_srs_SwitchingTimeRequested) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RF_ParametersMRDC_srs_SwitchingTimeRequested", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
