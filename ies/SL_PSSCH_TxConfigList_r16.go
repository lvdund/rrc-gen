package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_TxConfigList_r16 struct {
	Value []SL_PSSCH_TxConfig_r16 `lb:1,ub:maxPSSCH_TxConfig_r16,madatory`
}

func (ie *SL_PSSCH_TxConfigList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_PSSCH_TxConfig_r16]([]*SL_PSSCH_TxConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxPSSCH_TxConfig_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_PSSCH_TxConfigList_r16", err)
	}
	return nil
}

func (ie *SL_PSSCH_TxConfigList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_PSSCH_TxConfig_r16]([]*SL_PSSCH_TxConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxPSSCH_TxConfig_r16}, false)
	fn := func() *SL_PSSCH_TxConfig_r16 {
		return new(SL_PSSCH_TxConfig_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_PSSCH_TxConfigList_r16", err)
	}
	ie.Value = []SL_PSSCH_TxConfig_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
