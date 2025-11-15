package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ExcessDelay_DRB_IdentityInfo_r17 struct {
	drb_IdentityList []DRB_Identity                                  `lb:1,ub:maxDRB,madatory`
	delayThreshold   ExcessDelay_DRB_IdentityInfo_r17_delayThreshold `madatory`
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_drb_IdentityList := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.drb_IdentityList {
		tmp_drb_IdentityList.Value = append(tmp_drb_IdentityList.Value, &i)
	}
	if err = tmp_drb_IdentityList.Encode(w); err != nil {
		return utils.WrapError("Encode drb_IdentityList", err)
	}
	if err = ie.delayThreshold.Encode(w); err != nil {
		return utils.WrapError("Encode delayThreshold", err)
	}
	return nil
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_drb_IdentityList := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn_drb_IdentityList := func() *DRB_Identity {
		return new(DRB_Identity)
	}
	if err = tmp_drb_IdentityList.Decode(r, fn_drb_IdentityList); err != nil {
		return utils.WrapError("Decode drb_IdentityList", err)
	}
	ie.drb_IdentityList = []DRB_Identity{}
	for _, i := range tmp_drb_IdentityList.Value {
		ie.drb_IdentityList = append(ie.drb_IdentityList, *i)
	}
	if err = ie.delayThreshold.Decode(r); err != nil {
		return utils.WrapError("Decode delayThreshold", err)
	}
	return nil
}
