package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetup_v1700_IEs struct {
	sl_ConfigDedicatedNR_r17 *SL_ConfigDedicatedNR_r16 `optional`
	sl_L2RemoteUE_Config_r17 *SL_L2RemoteUE_Config_r17 `optional`
	nonCriticalExtension     interface{}               `optional`
}

func (ie *RRCSetup_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ConfigDedicatedNR_r17 != nil, ie.sl_L2RemoteUE_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ConfigDedicatedNR_r17 != nil {
		if err = ie.sl_ConfigDedicatedNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfigDedicatedNR_r17", err)
		}
	}
	if ie.sl_L2RemoteUE_Config_r17 != nil {
		if err = ie.sl_L2RemoteUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_L2RemoteUE_Config_r17", err)
		}
	}
	return nil
}

func (ie *RRCSetup_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_ConfigDedicatedNR_r17Present bool
	if sl_ConfigDedicatedNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_L2RemoteUE_Config_r17Present bool
	if sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ConfigDedicatedNR_r17Present {
		ie.sl_ConfigDedicatedNR_r17 = new(SL_ConfigDedicatedNR_r16)
		if err = ie.sl_ConfigDedicatedNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ConfigDedicatedNR_r17", err)
		}
	}
	if sl_L2RemoteUE_Config_r17Present {
		ie.sl_L2RemoteUE_Config_r17 = new(SL_L2RemoteUE_Config_r17)
		if err = ie.sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_L2RemoteUE_Config_r17", err)
		}
	}
	return nil
}
