package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResume_v1700_IEs struct {
	sl_ConfigDedicatedNR_r17       *SL_ConfigDedicatedNR_r16          `optional,setuprelease`
	sl_L2RemoteUE_Config_r17       *SL_L2RemoteUE_Config_r17          `optional,setuprelease`
	needForGapNCSG_ConfigNR_r17    *NeedForGapNCSG_ConfigNR_r17       `optional,setuprelease`
	needForGapNCSG_ConfigEUTRA_r17 *NeedForGapNCSG_ConfigEUTRA_r17    `optional,setuprelease`
	scg_State_r17                  *RRCResume_v1700_IEs_scg_State_r17 `optional`
	appLayerMeasConfig_r17         *AppLayerMeasConfig_r17            `optional`
	nonCriticalExtension           interface{}                        `optional`
}

func (ie *RRCResume_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ConfigDedicatedNR_r17 != nil, ie.sl_L2RemoteUE_Config_r17 != nil, ie.needForGapNCSG_ConfigNR_r17 != nil, ie.needForGapNCSG_ConfigEUTRA_r17 != nil, ie.scg_State_r17 != nil, ie.appLayerMeasConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ConfigDedicatedNR_r17 != nil {
		tmp_sl_ConfigDedicatedNR_r17 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{
			Setup: ie.sl_ConfigDedicatedNR_r17,
		}
		if err = tmp_sl_ConfigDedicatedNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfigDedicatedNR_r17", err)
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
	if ie.needForGapNCSG_ConfigNR_r17 != nil {
		tmp_needForGapNCSG_ConfigNR_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigNR_r17]{
			Setup: ie.needForGapNCSG_ConfigNR_r17,
		}
		if err = tmp_needForGapNCSG_ConfigNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapNCSG_ConfigNR_r17", err)
		}
	}
	if ie.needForGapNCSG_ConfigEUTRA_r17 != nil {
		tmp_needForGapNCSG_ConfigEUTRA_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigEUTRA_r17]{
			Setup: ie.needForGapNCSG_ConfigEUTRA_r17,
		}
		if err = tmp_needForGapNCSG_ConfigEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapNCSG_ConfigEUTRA_r17", err)
		}
	}
	if ie.scg_State_r17 != nil {
		if err = ie.scg_State_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_State_r17", err)
		}
	}
	if ie.appLayerMeasConfig_r17 != nil {
		if err = ie.appLayerMeasConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode appLayerMeasConfig_r17", err)
		}
	}
	return nil
}

func (ie *RRCResume_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_ConfigDedicatedNR_r17Present bool
	if sl_ConfigDedicatedNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_L2RemoteUE_Config_r17Present bool
	if sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapNCSG_ConfigNR_r17Present bool
	if needForGapNCSG_ConfigNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapNCSG_ConfigEUTRA_r17Present bool
	if needForGapNCSG_ConfigEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_State_r17Present bool
	if scg_State_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var appLayerMeasConfig_r17Present bool
	if appLayerMeasConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ConfigDedicatedNR_r17Present {
		tmp_sl_ConfigDedicatedNR_r17 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{}
		if err = tmp_sl_ConfigDedicatedNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ConfigDedicatedNR_r17", err)
		}
		ie.sl_ConfigDedicatedNR_r17 = tmp_sl_ConfigDedicatedNR_r17.Setup
	}
	if sl_L2RemoteUE_Config_r17Present {
		tmp_sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{}
		if err = tmp_sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_L2RemoteUE_Config_r17", err)
		}
		ie.sl_L2RemoteUE_Config_r17 = tmp_sl_L2RemoteUE_Config_r17.Setup
	}
	if needForGapNCSG_ConfigNR_r17Present {
		tmp_needForGapNCSG_ConfigNR_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigNR_r17]{}
		if err = tmp_needForGapNCSG_ConfigNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapNCSG_ConfigNR_r17", err)
		}
		ie.needForGapNCSG_ConfigNR_r17 = tmp_needForGapNCSG_ConfigNR_r17.Setup
	}
	if needForGapNCSG_ConfigEUTRA_r17Present {
		tmp_needForGapNCSG_ConfigEUTRA_r17 := utils.SetupRelease[*NeedForGapNCSG_ConfigEUTRA_r17]{}
		if err = tmp_needForGapNCSG_ConfigEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapNCSG_ConfigEUTRA_r17", err)
		}
		ie.needForGapNCSG_ConfigEUTRA_r17 = tmp_needForGapNCSG_ConfigEUTRA_r17.Setup
	}
	if scg_State_r17Present {
		ie.scg_State_r17 = new(RRCResume_v1700_IEs_scg_State_r17)
		if err = ie.scg_State_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_State_r17", err)
		}
	}
	if appLayerMeasConfig_r17Present {
		ie.appLayerMeasConfig_r17 = new(AppLayerMeasConfig_r17)
		if err = ie.appLayerMeasConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode appLayerMeasConfig_r17", err)
		}
	}
	return nil
}
