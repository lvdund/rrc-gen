package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_DelayValueConfig_r16 struct {
	delay_DRBlist_r16 []DRB_Identity `lb:1,ub:maxDRB,madatory`
}

func (ie *UL_DelayValueConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_delay_DRBlist_r16 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.delay_DRBlist_r16 {
		tmp_delay_DRBlist_r16.Value = append(tmp_delay_DRBlist_r16.Value, &i)
	}
	if err = tmp_delay_DRBlist_r16.Encode(w); err != nil {
		return utils.WrapError("Encode delay_DRBlist_r16", err)
	}
	return nil
}

func (ie *UL_DelayValueConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_delay_DRBlist_r16 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn_delay_DRBlist_r16 := func() *DRB_Identity {
		return new(DRB_Identity)
	}
	if err = tmp_delay_DRBlist_r16.Decode(r, fn_delay_DRBlist_r16); err != nil {
		return utils.WrapError("Decode delay_DRBlist_r16", err)
	}
	ie.delay_DRBlist_r16 = []DRB_Identity{}
	for _, i := range tmp_delay_DRBlist_r16.Value {
		ie.delay_DRBlist_r16 = append(ie.delay_DRBlist_r16, *i)
	}
	return nil
}
