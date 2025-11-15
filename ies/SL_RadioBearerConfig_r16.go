package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RadioBearerConfig_r16 struct {
	slrb_Uu_ConfigIndex_r16 SLRB_Uu_ConfigIndex_r16                     `madatory`
	sl_SDAP_Config_r16      *SL_SDAP_Config_r16                         `optional`
	sl_PDCP_Config_r16      *SL_PDCP_Config_r16                         `optional`
	sl_TransRange_r16       *SL_RadioBearerConfig_r16_sl_TransRange_r16 `optional`
}

func (ie *SL_RadioBearerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_SDAP_Config_r16 != nil, ie.sl_PDCP_Config_r16 != nil, ie.sl_TransRange_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.slrb_Uu_ConfigIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode slrb_Uu_ConfigIndex_r16", err)
	}
	if ie.sl_SDAP_Config_r16 != nil {
		if err = ie.sl_SDAP_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SDAP_Config_r16", err)
		}
	}
	if ie.sl_PDCP_Config_r16 != nil {
		if err = ie.sl_PDCP_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PDCP_Config_r16", err)
		}
	}
	if ie.sl_TransRange_r16 != nil {
		if err = ie.sl_TransRange_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TransRange_r16", err)
		}
	}
	return nil
}

func (ie *SL_RadioBearerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_SDAP_Config_r16Present bool
	if sl_SDAP_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PDCP_Config_r16Present bool
	if sl_PDCP_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TransRange_r16Present bool
	if sl_TransRange_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.slrb_Uu_ConfigIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode slrb_Uu_ConfigIndex_r16", err)
	}
	if sl_SDAP_Config_r16Present {
		ie.sl_SDAP_Config_r16 = new(SL_SDAP_Config_r16)
		if err = ie.sl_SDAP_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SDAP_Config_r16", err)
		}
	}
	if sl_PDCP_Config_r16Present {
		ie.sl_PDCP_Config_r16 = new(SL_PDCP_Config_r16)
		if err = ie.sl_PDCP_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PDCP_Config_r16", err)
		}
	}
	if sl_TransRange_r16Present {
		ie.sl_TransRange_r16 = new(SL_RadioBearerConfig_r16_sl_TransRange_r16)
		if err = ie.sl_TransRange_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TransRange_r16", err)
		}
	}
	return nil
}
