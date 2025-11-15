package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ModeIndication_r16 struct {
	sl_Mode_r16         SL_RLC_ModeIndication_r16_sl_Mode_r16 `madatory`
	sl_QoS_InfoList_r16 []SL_QoS_Info_r16                     `lb:1,ub:maxNrofSL_QFIsPerDest_r16,madatory`
}

func (ie *SL_RLC_ModeIndication_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_Mode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_Mode_r16", err)
	}
	tmp_sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
	for _, i := range ie.sl_QoS_InfoList_r16 {
		tmp_sl_QoS_InfoList_r16.Value = append(tmp_sl_QoS_InfoList_r16.Value, &i)
	}
	if err = tmp_sl_QoS_InfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_QoS_InfoList_r16", err)
	}
	return nil
}

func (ie *SL_RLC_ModeIndication_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_Mode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_Mode_r16", err)
	}
	tmp_sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
	fn_sl_QoS_InfoList_r16 := func() *SL_QoS_Info_r16 {
		return new(SL_QoS_Info_r16)
	}
	if err = tmp_sl_QoS_InfoList_r16.Decode(r, fn_sl_QoS_InfoList_r16); err != nil {
		return utils.WrapError("Decode sl_QoS_InfoList_r16", err)
	}
	ie.sl_QoS_InfoList_r16 = []SL_QoS_Info_r16{}
	for _, i := range tmp_sl_QoS_InfoList_r16.Value {
		ie.sl_QoS_InfoList_r16 = append(ie.sl_QoS_InfoList_r16, *i)
	}
	return nil
}
