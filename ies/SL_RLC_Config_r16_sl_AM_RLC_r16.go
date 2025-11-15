package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_Config_r16_sl_AM_RLC_r16 struct {
	sl_SN_FieldLengthAM_r16 *SN_FieldLengthAM                                       `optional`
	sl_T_PollRetransmit_r16 T_PollRetransmit                                        `madatory`
	sl_PollPDU_r16          PollPDU                                                 `madatory`
	sl_PollByte_r16         PollByte                                                `madatory`
	sl_MaxRetxThreshold_r16 SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16 `madatory`
}

func (ie *SL_RLC_Config_r16_sl_AM_RLC_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_SN_FieldLengthAM_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_SN_FieldLengthAM_r16 != nil {
		if err = ie.sl_SN_FieldLengthAM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SN_FieldLengthAM_r16", err)
		}
	}
	if err = ie.sl_T_PollRetransmit_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_T_PollRetransmit_r16", err)
	}
	if err = ie.sl_PollPDU_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_PollPDU_r16", err)
	}
	if err = ie.sl_PollByte_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_PollByte_r16", err)
	}
	if err = ie.sl_MaxRetxThreshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_MaxRetxThreshold_r16", err)
	}
	return nil
}

func (ie *SL_RLC_Config_r16_sl_AM_RLC_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_SN_FieldLengthAM_r16Present bool
	if sl_SN_FieldLengthAM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_SN_FieldLengthAM_r16Present {
		ie.sl_SN_FieldLengthAM_r16 = new(SN_FieldLengthAM)
		if err = ie.sl_SN_FieldLengthAM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SN_FieldLengthAM_r16", err)
		}
	}
	if err = ie.sl_T_PollRetransmit_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_T_PollRetransmit_r16", err)
	}
	if err = ie.sl_PollPDU_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_PollPDU_r16", err)
	}
	if err = ie.sl_PollByte_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_PollByte_r16", err)
	}
	if err = ie.sl_MaxRetxThreshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_MaxRetxThreshold_r16", err)
	}
	return nil
}
