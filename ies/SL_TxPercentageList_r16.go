package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxPercentageList_r16 struct {
	Value []SL_TxPercentageConfig_r16 `lb:8,ub:8,madatory`
}

func (ie *SL_TxPercentageList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_TxPercentageConfig_r16]([]*SL_TxPercentageConfig_r16{}, uper.Constraint{Lb: 8, Ub: 8}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_TxPercentageList_r16", err)
	}
	return nil
}

func (ie *SL_TxPercentageList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_TxPercentageConfig_r16]([]*SL_TxPercentageConfig_r16{}, uper.Constraint{Lb: 8, Ub: 8}, false)
	fn := func() *SL_TxPercentageConfig_r16 {
		return new(SL_TxPercentageConfig_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_TxPercentageList_r16", err)
	}
	ie.Value = []SL_TxPercentageConfig_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
