package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ZoneConfigMCR_r16 struct {
	sl_ZoneConfigMCR_Index_r16 int64                                   `lb:0,ub:15,madatory`
	sl_TransRange_r16          *SL_ZoneConfigMCR_r16_sl_TransRange_r16 `optional`
	sl_ZoneConfig_r16          *SL_ZoneConfig_r16                      `optional`
}

func (ie *SL_ZoneConfigMCR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TransRange_r16 != nil, ie.sl_ZoneConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.sl_ZoneConfigMCR_Index_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger sl_ZoneConfigMCR_Index_r16", err)
	}
	if ie.sl_TransRange_r16 != nil {
		if err = ie.sl_TransRange_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TransRange_r16", err)
		}
	}
	if ie.sl_ZoneConfig_r16 != nil {
		if err = ie.sl_ZoneConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ZoneConfig_r16", err)
		}
	}
	return nil
}

func (ie *SL_ZoneConfigMCR_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_TransRange_r16Present bool
	if sl_TransRange_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ZoneConfig_r16Present bool
	if sl_ZoneConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_sl_ZoneConfigMCR_Index_r16 int64
	if tmp_int_sl_ZoneConfigMCR_Index_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger sl_ZoneConfigMCR_Index_r16", err)
	}
	ie.sl_ZoneConfigMCR_Index_r16 = tmp_int_sl_ZoneConfigMCR_Index_r16
	if sl_TransRange_r16Present {
		ie.sl_TransRange_r16 = new(SL_ZoneConfigMCR_r16_sl_TransRange_r16)
		if err = ie.sl_TransRange_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TransRange_r16", err)
		}
	}
	if sl_ZoneConfig_r16Present {
		ie.sl_ZoneConfig_r16 = new(SL_ZoneConfig_r16)
		if err = ie.sl_ZoneConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ZoneConfig_r16", err)
		}
	}
	return nil
}
