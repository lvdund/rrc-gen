package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NSAG_List_r17 struct {
	Value []NSAG_ID_r17 `lb:1,ub:maxSliceInfo_r17,madatory`
}

func (ie *NSAG_List_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*NSAG_ID_r17]([]*NSAG_ID_r17{}, uper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NSAG_List_r17", err)
	}
	return nil
}

func (ie *NSAG_List_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*NSAG_ID_r17]([]*NSAG_ID_r17{}, uper.Constraint{Lb: 1, Ub: maxSliceInfo_r17}, false)
	fn := func() *NSAG_ID_r17 {
		return new(NSAG_ID_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NSAG_List_r17", err)
	}
	ie.Value = []NSAG_ID_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
