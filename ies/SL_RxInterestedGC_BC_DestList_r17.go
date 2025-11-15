package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RxInterestedGC_BC_DestList_r17 struct {
	Value []SL_RxInterestedGC_BC_Dest_r17 `lb:1,ub:maxNrofSL_Dest_r16,madatory`
}

func (ie *SL_RxInterestedGC_BC_DestList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_RxInterestedGC_BC_Dest_r17]([]*SL_RxInterestedGC_BC_Dest_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_RxInterestedGC_BC_DestList_r17", err)
	}
	return nil
}

func (ie *SL_RxInterestedGC_BC_DestList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_RxInterestedGC_BC_Dest_r17]([]*SL_RxInterestedGC_BC_Dest_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
	fn := func() *SL_RxInterestedGC_BC_Dest_r17 {
		return new(SL_RxInterestedGC_BC_Dest_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_RxInterestedGC_BC_DestList_r17", err)
	}
	ie.Value = []SL_RxInterestedGC_BC_Dest_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
