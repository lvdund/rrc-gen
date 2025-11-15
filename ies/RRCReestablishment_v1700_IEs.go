package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishment_v1700_IEs struct {
	sl_L2RemoteUE_Config_r17 *SL_L2RemoteUE_Config_r17 `optional,setuprelease`
	nonCriticalExtension     interface{}               `optional`
}

func (ie *RRCReestablishment_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_L2RemoteUE_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_L2RemoteUE_Config_r17 != nil {
		tmp_sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{
			Setup: ie.sl_L2RemoteUE_Config_r17,
		}
		if err = tmp_sl_L2RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_L2RemoteUE_Config_r17", err)
		}
	}
	return nil
}

func (ie *RRCReestablishment_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_L2RemoteUE_Config_r17Present bool
	if sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_L2RemoteUE_Config_r17Present {
		tmp_sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{}
		if err = tmp_sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_L2RemoteUE_Config_r17", err)
		}
		ie.sl_L2RemoteUE_Config_r17 = tmp_sl_L2RemoteUE_Config_r17.Setup
	}
	return nil
}
