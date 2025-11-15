package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentTwoCarrierList_r16 struct {
	Value []UplinkTxDirectCurrentTwoCarrier_r16 `lb:1,ub:maxNrofTxDC_TwoCarrier_r16,madatory`
}

func (ie *UplinkTxDirectCurrentTwoCarrierList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UplinkTxDirectCurrentTwoCarrier_r16]([]*UplinkTxDirectCurrentTwoCarrier_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTxDC_TwoCarrier_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkTxDirectCurrentTwoCarrierList_r16", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrierList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UplinkTxDirectCurrentTwoCarrier_r16]([]*UplinkTxDirectCurrentTwoCarrier_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofTxDC_TwoCarrier_r16}, false)
	fn := func() *UplinkTxDirectCurrentTwoCarrier_r16 {
		return new(UplinkTxDirectCurrentTwoCarrier_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UplinkTxDirectCurrentTwoCarrierList_r16", err)
	}
	ie.Value = []UplinkTxDirectCurrentTwoCarrier_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
