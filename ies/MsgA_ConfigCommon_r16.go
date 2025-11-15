package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_ConfigCommon_r16 struct {
	rach_ConfigCommonTwoStepRA_r16 RACH_ConfigCommonTwoStepRA_r16 `madatory`
	msgA_PUSCH_Config_r16          *MsgA_PUSCH_Config_r16         `optional`
}

func (ie *MsgA_ConfigCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.msgA_PUSCH_Config_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rach_ConfigCommonTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rach_ConfigCommonTwoStepRA_r16", err)
	}
	if ie.msgA_PUSCH_Config_r16 != nil {
		if err = ie.msgA_PUSCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_PUSCH_Config_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_ConfigCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var msgA_PUSCH_Config_r16Present bool
	if msgA_PUSCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rach_ConfigCommonTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rach_ConfigCommonTwoStepRA_r16", err)
	}
	if msgA_PUSCH_Config_r16Present {
		ie.msgA_PUSCH_Config_r16 = new(MsgA_PUSCH_Config_r16)
		if err = ie.msgA_PUSCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_PUSCH_Config_r16", err)
		}
	}
	return nil
}
