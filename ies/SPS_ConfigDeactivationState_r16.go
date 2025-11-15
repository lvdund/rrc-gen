package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SPS_ConfigDeactivationState_r16 struct {
	Value []SPS_ConfigIndex_r16 `lb:1,ub:maxNrofSPS_Config_r16,madatory`
}

func (ie *SPS_ConfigDeactivationState_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SPS_ConfigIndex_r16]([]*SPS_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSPS_Config_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SPS_ConfigDeactivationState_r16", err)
	}
	return nil
}

func (ie *SPS_ConfigDeactivationState_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SPS_ConfigIndex_r16]([]*SPS_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSPS_Config_r16}, false)
	fn := func() *SPS_ConfigIndex_r16 {
		return new(SPS_ConfigIndex_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SPS_ConfigDeactivationState_r16", err)
	}
	ie.Value = []SPS_ConfigIndex_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
