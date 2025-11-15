package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1700_IEs struct {
	otherConfig_v1700                  *OtherConfig_v1700                          `optional`
	sl_L2RelayUE_Config_r17            *SL_L2RelayUE_Config_r17                    `optional,setuprelease`
	sl_L2RemoteUE_Config_r17           *SL_L2RemoteUE_Config_r17                   `optional,setuprelease`
	dedicatedPagingDelivery_r17        *[]byte                                     `optional`
	needForGapNCSG_ConfigNR_r17        *NeedForGapNCSG_ConfigNR_r17                `optional,setuprelease`
	needForGapNCSG_ConfigEUTRA_r17     *NeedForGapNCSG_ConfigEUTRA_r17             `optional,setuprelease`
	musim_GapConfig_r17                *MUSIM_GapConfig_r17                        `optional,setuprelease`
	ul_GapFR2_Config_r17               *UL_GapFR2_Config_r17                       `optional,setuprelease`
	scg_State_r17                      *RRCReconfiguration_v1700_IEs_scg_State_r17 `optional`
	appLayerMeasConfig_r17             *AppLayerMeasConfig_r17                     `optional`
	ue_TxTEG_RequestUL_TDOA_Config_r17 *UE_TxTEG_RequestUL_TDOA_Config_r17         `optional,setuprelease`
	nonCriticalExtension               interface{}                                 `optional`
}

func (ie *RRCReconfiguration_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.otherConfig_v1700 != nil, ie.sl_L2RelayUE_Config_r17 != nil, ie.sl_L2RemoteUE_Config_r17 != nil, ie.dedicatedPagingDelivery_r17 != nil, ie.needForGapNCSG_ConfigNR_r17 != nil, ie.needForGapNCSG_ConfigEUTRA_r17 != nil, ie.musim_GapConfig_r17 != nil, ie.ul_GapFR2_Config_r17 != nil, ie.scg_State_r17 != nil, ie.appLayerMeasConfig_r17 != nil, ie.ue_TxTEG_RequestUL_TDOA_Config_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.otherConfig_v1700 != nil {
		if err = ie.otherConfig_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode otherConfig_v1700", err)
		}
	}
	if ie.sl_L2RelayUE_Config_r17 != nil {
		tmp_sl_L2RelayUE_Config_r17 := utils.SetupRelease[*SL_L2RelayUE_Config_r17]{
			Setup: ie.sl_L2RelayUE_Config_r17,
		}
		if err = tmp_sl_L2RelayUE_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_L2RelayUE_Config_r17", err)
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
	if ie.dedicatedPagingDelivery_r17 != nil {
		if err = w.WriteOctetString(*ie.dedicatedPagingDelivery_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dedicatedPagingDelivery_r17", err)
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
	if ie.musim_GapConfig_r17 != nil {
		tmp_musim_GapConfig_r17 := utils.SetupRelease[*MUSIM_GapConfig_r17]{
			Setup: ie.musim_GapConfig_r17,
		}
		if err = tmp_musim_GapConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapConfig_r17", err)
		}
	}
	if ie.ul_GapFR2_Config_r17 != nil {
		tmp_ul_GapFR2_Config_r17 := utils.SetupRelease[*UL_GapFR2_Config_r17]{
			Setup: ie.ul_GapFR2_Config_r17,
		}
		if err = tmp_ul_GapFR2_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_GapFR2_Config_r17", err)
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
	if ie.ue_TxTEG_RequestUL_TDOA_Config_r17 != nil {
		tmp_ue_TxTEG_RequestUL_TDOA_Config_r17 := utils.SetupRelease[*UE_TxTEG_RequestUL_TDOA_Config_r17]{
			Setup: ie.ue_TxTEG_RequestUL_TDOA_Config_r17,
		}
		if err = tmp_ue_TxTEG_RequestUL_TDOA_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_TxTEG_RequestUL_TDOA_Config_r17", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var otherConfig_v1700Present bool
	if otherConfig_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_L2RelayUE_Config_r17Present bool
	if sl_L2RelayUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_L2RemoteUE_Config_r17Present bool
	if sl_L2RemoteUE_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dedicatedPagingDelivery_r17Present bool
	if dedicatedPagingDelivery_r17Present, err = r.ReadBool(); err != nil {
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
	var musim_GapConfig_r17Present bool
	if musim_GapConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_GapFR2_Config_r17Present bool
	if ul_GapFR2_Config_r17Present, err = r.ReadBool(); err != nil {
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
	var ue_TxTEG_RequestUL_TDOA_Config_r17Present bool
	if ue_TxTEG_RequestUL_TDOA_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if otherConfig_v1700Present {
		ie.otherConfig_v1700 = new(OtherConfig_v1700)
		if err = ie.otherConfig_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode otherConfig_v1700", err)
		}
	}
	if sl_L2RelayUE_Config_r17Present {
		tmp_sl_L2RelayUE_Config_r17 := utils.SetupRelease[*SL_L2RelayUE_Config_r17]{}
		if err = tmp_sl_L2RelayUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_L2RelayUE_Config_r17", err)
		}
		ie.sl_L2RelayUE_Config_r17 = tmp_sl_L2RelayUE_Config_r17.Setup
	}
	if sl_L2RemoteUE_Config_r17Present {
		tmp_sl_L2RemoteUE_Config_r17 := utils.SetupRelease[*SL_L2RemoteUE_Config_r17]{}
		if err = tmp_sl_L2RemoteUE_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_L2RemoteUE_Config_r17", err)
		}
		ie.sl_L2RemoteUE_Config_r17 = tmp_sl_L2RemoteUE_Config_r17.Setup
	}
	if dedicatedPagingDelivery_r17Present {
		var tmp_os_dedicatedPagingDelivery_r17 []byte
		if tmp_os_dedicatedPagingDelivery_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dedicatedPagingDelivery_r17", err)
		}
		ie.dedicatedPagingDelivery_r17 = &tmp_os_dedicatedPagingDelivery_r17
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
	if musim_GapConfig_r17Present {
		tmp_musim_GapConfig_r17 := utils.SetupRelease[*MUSIM_GapConfig_r17]{}
		if err = tmp_musim_GapConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_GapConfig_r17", err)
		}
		ie.musim_GapConfig_r17 = tmp_musim_GapConfig_r17.Setup
	}
	if ul_GapFR2_Config_r17Present {
		tmp_ul_GapFR2_Config_r17 := utils.SetupRelease[*UL_GapFR2_Config_r17]{}
		if err = tmp_ul_GapFR2_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_GapFR2_Config_r17", err)
		}
		ie.ul_GapFR2_Config_r17 = tmp_ul_GapFR2_Config_r17.Setup
	}
	if scg_State_r17Present {
		ie.scg_State_r17 = new(RRCReconfiguration_v1700_IEs_scg_State_r17)
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
	if ue_TxTEG_RequestUL_TDOA_Config_r17Present {
		tmp_ue_TxTEG_RequestUL_TDOA_Config_r17 := utils.SetupRelease[*UE_TxTEG_RequestUL_TDOA_Config_r17]{}
		if err = tmp_ue_TxTEG_RequestUL_TDOA_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_TxTEG_RequestUL_TDOA_Config_r17", err)
		}
		ie.ue_TxTEG_RequestUL_TDOA_Config_r17 = tmp_ue_TxTEG_RequestUL_TDOA_Config_r17.Setup
	}
	return nil
}
