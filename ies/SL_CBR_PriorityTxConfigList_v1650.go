package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_CBR_PriorityTxConfigList_v1650 struct {
	Value []SL_PriorityTxConfigIndex_v1650 `lb:1,ub:8,madatory`
}

func (ie *SL_CBR_PriorityTxConfigList_v1650) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_PriorityTxConfigIndex_v1650]([]*SL_PriorityTxConfigIndex_v1650{}, uper.Constraint{Lb: 1, Ub: 8}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_CBR_PriorityTxConfigList_v1650", err)
	}
	return nil
}

func (ie *SL_CBR_PriorityTxConfigList_v1650) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_PriorityTxConfigIndex_v1650]([]*SL_PriorityTxConfigIndex_v1650{}, uper.Constraint{Lb: 1, Ub: 8}, false)
	fn := func() *SL_PriorityTxConfigIndex_v1650 {
		return new(SL_PriorityTxConfigIndex_v1650)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_CBR_PriorityTxConfigList_v1650", err)
	}
	ie.Value = []SL_PriorityTxConfigIndex_v1650{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
