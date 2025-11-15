package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_Thres_RSRP_List_r16 struct {
	Value []SL_Thres_RSRP_r16 `lb:64,ub:64,madatory`
}

func (ie *SL_Thres_RSRP_List_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_Thres_RSRP_r16]([]*SL_Thres_RSRP_r16{}, uper.Constraint{Lb: 64, Ub: 64}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_Thres_RSRP_List_r16", err)
	}
	return nil
}

func (ie *SL_Thres_RSRP_List_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_Thres_RSRP_r16]([]*SL_Thres_RSRP_r16{}, uper.Constraint{Lb: 64, Ub: 64}, false)
	fn := func() *SL_Thres_RSRP_r16 {
		return new(SL_Thres_RSRP_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_Thres_RSRP_List_r16", err)
	}
	ie.Value = []SL_Thres_RSRP_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
