package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_CBR_PriorityTxConfigList_r16 struct {
	Value []SL_PriorityTxConfigIndex_r16 `lb:1,ub:8,madatory`
}

func (ie *SL_CBR_PriorityTxConfigList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_PriorityTxConfigIndex_r16]([]*SL_PriorityTxConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_CBR_PriorityTxConfigList_r16", err)
	}
	return nil
}

func (ie *SL_CBR_PriorityTxConfigList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_PriorityTxConfigIndex_r16]([]*SL_PriorityTxConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
	fn := func() *SL_PriorityTxConfigIndex_r16 {
		return new(SL_PriorityTxConfigIndex_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_CBR_PriorityTxConfigList_r16", err)
	}
	ie.Value = []SL_PriorityTxConfigIndex_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
