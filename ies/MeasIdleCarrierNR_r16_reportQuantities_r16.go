package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasIdleCarrierNR_r16_reportQuantities_r16_Enum_rsrp uper.Enumerated = 0
	MeasIdleCarrierNR_r16_reportQuantities_r16_Enum_rsrq uper.Enumerated = 1
	MeasIdleCarrierNR_r16_reportQuantities_r16_Enum_both uper.Enumerated = 2
)

type MeasIdleCarrierNR_r16_reportQuantities_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MeasIdleCarrierNR_r16_reportQuantities_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MeasIdleCarrierNR_r16_reportQuantities_r16", err)
	}
	return nil
}

func (ie *MeasIdleCarrierNR_r16_reportQuantities_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MeasIdleCarrierNR_r16_reportQuantities_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
