package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BT_NameList_r16 struct {
	Value []BT_Name_r16 `lb:1,ub:maxBT_Name_r16,madatory`
}

func (ie *BT_NameList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BT_Name_r16]([]*BT_Name_r16{}, uper.Constraint{Lb: 1, Ub: maxBT_Name_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BT_NameList_r16", err)
	}
	return nil
}

func (ie *BT_NameList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BT_Name_r16]([]*BT_Name_r16{}, uper.Constraint{Lb: 1, Ub: maxBT_Name_r16}, false)
	fn := func() *BT_Name_r16 {
		return new(BT_Name_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BT_NameList_r16", err)
	}
	ie.Value = []BT_Name_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
