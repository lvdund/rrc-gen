package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RemoteUE_Config_r17 struct {
	threshHighRemote_r17     *RSRP_Range               `optional`
	hystMaxRemote_r17        *Hysteresis               `optional`
	sl_ReselectionConfig_r17 *SL_ReselectionConfig_r17 `optional`
}

func (ie *SL_RemoteUE_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.threshHighRemote_r17 != nil, ie.hystMaxRemote_r17 != nil, ie.sl_ReselectionConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.threshHighRemote_r17 != nil {
		if err = ie.threshHighRemote_r17.Encode(w); err != nil {
			return utils.WrapError("Encode threshHighRemote_r17", err)
		}
	}
	if ie.hystMaxRemote_r17 != nil {
		if err = ie.hystMaxRemote_r17.Encode(w); err != nil {
			return utils.WrapError("Encode hystMaxRemote_r17", err)
		}
	}
	if ie.sl_ReselectionConfig_r17 != nil {
		if err = ie.sl_ReselectionConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ReselectionConfig_r17", err)
		}
	}
	return nil
}

func (ie *SL_RemoteUE_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var threshHighRemote_r17Present bool
	if threshHighRemote_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var hystMaxRemote_r17Present bool
	if hystMaxRemote_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ReselectionConfig_r17Present bool
	if sl_ReselectionConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if threshHighRemote_r17Present {
		ie.threshHighRemote_r17 = new(RSRP_Range)
		if err = ie.threshHighRemote_r17.Decode(r); err != nil {
			return utils.WrapError("Decode threshHighRemote_r17", err)
		}
	}
	if hystMaxRemote_r17Present {
		ie.hystMaxRemote_r17 = new(Hysteresis)
		if err = ie.hystMaxRemote_r17.Decode(r); err != nil {
			return utils.WrapError("Decode hystMaxRemote_r17", err)
		}
	}
	if sl_ReselectionConfig_r17Present {
		ie.sl_ReselectionConfig_r17 = new(SL_ReselectionConfig_r17)
		if err = ie.sl_ReselectionConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ReselectionConfig_r17", err)
		}
	}
	return nil
}
