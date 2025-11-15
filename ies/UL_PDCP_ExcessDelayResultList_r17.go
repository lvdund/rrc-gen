package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_PDCP_ExcessDelayResultList_r17 struct {
	Value []UL_PDCP_ExcessDelayResult_r17 `lb:1,ub:maxDRB,madatory`
}

func (ie *UL_PDCP_ExcessDelayResultList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UL_PDCP_ExcessDelayResult_r17]([]*UL_PDCP_ExcessDelayResult_r17{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UL_PDCP_ExcessDelayResultList_r17", err)
	}
	return nil
}

func (ie *UL_PDCP_ExcessDelayResultList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UL_PDCP_ExcessDelayResult_r17]([]*UL_PDCP_ExcessDelayResult_r17{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn := func() *UL_PDCP_ExcessDelayResult_r17 {
		return new(UL_PDCP_ExcessDelayResult_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UL_PDCP_ExcessDelayResultList_r17", err)
	}
	ie.Value = []UL_PDCP_ExcessDelayResult_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
