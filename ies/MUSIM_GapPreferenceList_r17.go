package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapPreferenceList_r17 struct {
	Value []MUSIM_GapInfo_r17 `lb:1,ub:4,madatory`
}

func (ie *MUSIM_GapPreferenceList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MUSIM_GapInfo_r17]([]*MUSIM_GapInfo_r17{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MUSIM_GapPreferenceList_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapPreferenceList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MUSIM_GapInfo_r17]([]*MUSIM_GapInfo_r17{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	fn := func() *MUSIM_GapInfo_r17 {
		return new(MUSIM_GapInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MUSIM_GapPreferenceList_r17", err)
	}
	ie.Value = []MUSIM_GapInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
