package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DiscConfig_r17 struct {
	sl_RelayUE_Config_r17  *SL_RelayUE_Config_r17  `optional,setuprelease`
	sl_RemoteUE_Config_r17 *SL_RemoteUE_Config_r17 `optional,setuprelease`
}

func (ie *SL_DiscConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RelayUE_Config_r17 != nil, ie.sl_RemoteUE_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_RelayUE_Config_r17 != nil {
		tmp_sl_RelayUE_Config_r17 := utils.SetupRelease[*SL_RelayUE_Config_r17]{
			Setup: ie.sl_RelayUE_Config_r17,
		}
		if err = tmp_sl_RelayUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RelayUE_Config_r17", err)
		}
	}
	if ie.sl_RemoteUE_Config_r17 != nil {
		tmp_sl_RemoteUE_Config_r17 := utils.SetupRelease[*SL_RemoteUE_Config_r17]{
			Setup: ie.sl_RemoteUE_Config_r17,
		}
		if err = tmp_sl_RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RemoteUE_Config_r17", err)
		}
	}
	return nil
}

func (ie *SL_DiscConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_RelayUE_Config_r17Present bool
	if sl_RelayUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RemoteUE_Config_r17Present bool
	if sl_RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RelayUE_Config_r17Present {
		tmp_sl_RelayUE_Config_r17 := utils.SetupRelease[*SL_RelayUE_Config_r17]{}
		if err = tmp_sl_RelayUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RelayUE_Config_r17", err)
		}
		ie.sl_RelayUE_Config_r17 = tmp_sl_RelayUE_Config_r17.Setup
	}
	if sl_RemoteUE_Config_r17Present {
		tmp_sl_RemoteUE_Config_r17 := utils.SetupRelease[*SL_RemoteUE_Config_r17]{}
		if err = tmp_sl_RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RemoteUE_Config_r17", err)
		}
		ie.sl_RemoteUE_Config_r17 = tmp_sl_RemoteUE_Config_r17.Setup
	}
	return nil
}
