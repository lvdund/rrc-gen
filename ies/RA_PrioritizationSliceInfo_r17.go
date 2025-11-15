package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RA_PrioritizationSliceInfo_r17 struct {
	nsag_ID_List_r17      []NSAG_ID_r17     `lb:1,ub:maxSliceInfo_r17,madatory`
	ra_Prioritization_r17 RA_Prioritization `madatory`
}

func (ie *RA_PrioritizationSliceInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_nsag_ID_List_r17 := utils.NewSequence[*NSAG_ID_r17]([]*NSAG_ID_r17{}, uper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	for _, i := range ie.nsag_ID_List_r17 {
		tmp_nsag_ID_List_r17.Value = append(tmp_nsag_ID_List_r17.Value, &i)
	}
	if err = tmp_nsag_ID_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nsag_ID_List_r17", err)
	}
	if err = ie.ra_Prioritization_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ra_Prioritization_r17", err)
	}
	return nil
}

func (ie *RA_PrioritizationSliceInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_nsag_ID_List_r17 := utils.NewSequence[*NSAG_ID_r17]([]*NSAG_ID_r17{}, uper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	fn_nsag_ID_List_r17 := func() *NSAG_ID_r17 {
		return new(NSAG_ID_r17)
	}
	if err = tmp_nsag_ID_List_r17.Decode(r, fn_nsag_ID_List_r17); err != nil {
		return utils.WrapError("Decode nsag_ID_List_r17", err)
	}
	ie.nsag_ID_List_r17 = []NSAG_ID_r17{}
	for _, i := range tmp_nsag_ID_List_r17.Value {
		ie.nsag_ID_List_r17 = append(ie.nsag_ID_List_r17, *i)
	}
	if err = ie.ra_Prioritization_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ra_Prioritization_r17", err)
	}
	return nil
}
