package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1630_IEs_uac_BarringInfo_v1630 struct {
	uac_AC1_SelectAssistInfo_r16 []UAC_AC1_SelectAssistInfo_r16 `lb:2,ub:maxPLMN,madatory`
}

func (ie *SIB1_v1630_IEs_uac_BarringInfo_v1630) Encode(w *uper.UperWriter) error {
	var err error
	tmp_uac_AC1_SelectAssistInfo_r16 := utils.NewSequence[*UAC_AC1_SelectAssistInfo_r16]([]*UAC_AC1_SelectAssistInfo_r16{}, uper.Constraint{Lb: 2, Ub: maxPLMN}, false)
	for _, i := range ie.uac_AC1_SelectAssistInfo_r16 {
		tmp_uac_AC1_SelectAssistInfo_r16.Value = append(tmp_uac_AC1_SelectAssistInfo_r16.Value, &i)
	}
	if err = tmp_uac_AC1_SelectAssistInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode uac_AC1_SelectAssistInfo_r16", err)
	}
	return nil
}

func (ie *SIB1_v1630_IEs_uac_BarringInfo_v1630) Decode(r *uper.UperReader) error {
	var err error
	tmp_uac_AC1_SelectAssistInfo_r16 := utils.NewSequence[*UAC_AC1_SelectAssistInfo_r16]([]*UAC_AC1_SelectAssistInfo_r16{}, uper.Constraint{Lb: 2, Ub: maxPLMN}, false)
	fn_uac_AC1_SelectAssistInfo_r16 := func() *UAC_AC1_SelectAssistInfo_r16 {
		return new(UAC_AC1_SelectAssistInfo_r16)
	}
	if err = tmp_uac_AC1_SelectAssistInfo_r16.Decode(r, fn_uac_AC1_SelectAssistInfo_r16); err != nil {
		return utils.WrapError("Decode uac_AC1_SelectAssistInfo_r16", err)
	}
	ie.uac_AC1_SelectAssistInfo_r16 = []UAC_AC1_SelectAssistInfo_r16{}
	for _, i := range tmp_uac_AC1_SelectAssistInfo_r16.Value {
		ie.uac_AC1_SelectAssistInfo_r16 = append(ie.uac_AC1_SelectAssistInfo_r16, *i)
	}
	return nil
}
