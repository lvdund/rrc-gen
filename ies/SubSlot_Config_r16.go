package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SubSlot_Config_r16 struct {
	sub_SlotConfig_NCP_r16 *SubSlot_Config_r16_sub_SlotConfig_NCP_r16 `optional`
	sub_SlotConfig_ECP_r16 *SubSlot_Config_r16_sub_SlotConfig_ECP_r16 `optional`
}

func (ie *SubSlot_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sub_SlotConfig_NCP_r16 != nil, ie.sub_SlotConfig_ECP_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sub_SlotConfig_NCP_r16 != nil {
		if err = ie.sub_SlotConfig_NCP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sub_SlotConfig_NCP_r16", err)
		}
	}
	if ie.sub_SlotConfig_ECP_r16 != nil {
		if err = ie.sub_SlotConfig_ECP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sub_SlotConfig_ECP_r16", err)
		}
	}
	return nil
}

func (ie *SubSlot_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sub_SlotConfig_NCP_r16Present bool
	if sub_SlotConfig_NCP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sub_SlotConfig_ECP_r16Present bool
	if sub_SlotConfig_ECP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sub_SlotConfig_NCP_r16Present {
		ie.sub_SlotConfig_NCP_r16 = new(SubSlot_Config_r16_sub_SlotConfig_NCP_r16)
		if err = ie.sub_SlotConfig_NCP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sub_SlotConfig_NCP_r16", err)
		}
	}
	if sub_SlotConfig_ECP_r16Present {
		ie.sub_SlotConfig_ECP_r16 = new(SubSlot_Config_r16_sub_SlotConfig_ECP_r16)
		if err = ie.sub_SlotConfig_ECP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sub_SlotConfig_ECP_r16", err)
		}
	}
	return nil
}
