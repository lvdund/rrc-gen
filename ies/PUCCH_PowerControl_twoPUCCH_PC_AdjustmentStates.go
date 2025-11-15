package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates_Enum_twoStates uper.Enumerated = 0
)

type PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates", err)
	}
	return nil
}

func (ie *PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
