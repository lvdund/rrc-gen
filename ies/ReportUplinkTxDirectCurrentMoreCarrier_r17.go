package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportUplinkTxDirectCurrentMoreCarrier_r17 struct {
	Value []IntraBandCC_CombinationReqList_r17 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *ReportUplinkTxDirectCurrentMoreCarrier_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*IntraBandCC_CombinationReqList_r17]([]*IntraBandCC_CombinationReqList_r17{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ReportUplinkTxDirectCurrentMoreCarrier_r17", err)
	}
	return nil
}

func (ie *ReportUplinkTxDirectCurrentMoreCarrier_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*IntraBandCC_CombinationReqList_r17]([]*IntraBandCC_CombinationReqList_r17{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *IntraBandCC_CombinationReqList_r17 {
		return new(IntraBandCC_CombinationReqList_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ReportUplinkTxDirectCurrentMoreCarrier_r17", err)
	}
	ie.Value = []IntraBandCC_CombinationReqList_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
