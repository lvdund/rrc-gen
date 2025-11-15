package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxProfileList_r17 struct {
	Value []SL_TxProfile_r17 `lb:1,ub:256,madatory`
}

func (ie *SL_TxProfileList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_TxProfile_r17]([]*SL_TxProfile_r17{}, uper.Constraint{Lb: 1, Ub: 256}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_TxProfileList_r17", err)
	}
	return nil
}

func (ie *SL_TxProfileList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_TxProfile_r17]([]*SL_TxProfile_r17{}, uper.Constraint{Lb: 1, Ub: 256}, false)
	fn := func() *SL_TxProfile_r17 {
		return new(SL_TxProfile_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_TxProfileList_r17", err)
	}
	ie.Value = []SL_TxProfile_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
