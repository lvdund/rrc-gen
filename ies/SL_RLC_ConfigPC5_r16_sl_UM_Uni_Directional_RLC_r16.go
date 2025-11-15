package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ConfigPC5_r16_sl_UM_Uni_Directional_RLC_r16 struct {
	sl_SN_FieldLengthUM_r16 *SN_FieldLengthUM `optional`
}

func (ie *SL_RLC_ConfigPC5_r16_sl_UM_Uni_Directional_RLC_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_SN_FieldLengthUM_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_SN_FieldLengthUM_r16 != nil {
		if err = ie.sl_SN_FieldLengthUM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SN_FieldLengthUM_r16", err)
		}
	}
	return nil
}

func (ie *SL_RLC_ConfigPC5_r16_sl_UM_Uni_Directional_RLC_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_SN_FieldLengthUM_r16Present bool
	if sl_SN_FieldLengthUM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_SN_FieldLengthUM_r16Present {
		ie.sl_SN_FieldLengthUM_r16 = new(SN_FieldLengthUM)
		if err = ie.sl_SN_FieldLengthUM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SN_FieldLengthUM_r16", err)
		}
	}
	return nil
}
