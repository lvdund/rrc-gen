package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_ExcessDelayConfig_r17 struct {
	excessDelay_DRBlist_r17 []ExcessDelay_DRB_IdentityInfo_r17 `lb:1,ub:maxDRB,madatory`
}

func (ie *UL_ExcessDelayConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_excessDelay_DRBlist_r17 := utils.NewSequence[*ExcessDelay_DRB_IdentityInfo_r17]([]*ExcessDelay_DRB_IdentityInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.excessDelay_DRBlist_r17 {
		tmp_excessDelay_DRBlist_r17.Value = append(tmp_excessDelay_DRBlist_r17.Value, &i)
	}
	if err = tmp_excessDelay_DRBlist_r17.Encode(w); err != nil {
		return utils.WrapError("Encode excessDelay_DRBlist_r17", err)
	}
	return nil
}

func (ie *UL_ExcessDelayConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_excessDelay_DRBlist_r17 := utils.NewSequence[*ExcessDelay_DRB_IdentityInfo_r17]([]*ExcessDelay_DRB_IdentityInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn_excessDelay_DRBlist_r17 := func() *ExcessDelay_DRB_IdentityInfo_r17 {
		return new(ExcessDelay_DRB_IdentityInfo_r17)
	}
	if err = tmp_excessDelay_DRBlist_r17.Decode(r, fn_excessDelay_DRBlist_r17); err != nil {
		return utils.WrapError("Decode excessDelay_DRBlist_r17", err)
	}
	ie.excessDelay_DRBlist_r17 = []ExcessDelay_DRB_IdentityInfo_r17{}
	for _, i := range tmp_excessDelay_DRBlist_r17.Value {
		ie.excessDelay_DRBlist_r17 = append(ie.excessDelay_DRBlist_r17, *i)
	}
	return nil
}
