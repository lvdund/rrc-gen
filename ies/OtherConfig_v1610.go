package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_v1610 struct {
	idc_AssistanceConfig_r16                *IDC_AssistanceConfig_r16                               `optional,setuprelease`
	drx_PreferenceConfig_r16                *DRX_PreferenceConfig_r16                               `optional,setuprelease`
	maxBW_PreferenceConfig_r16              *MaxBW_PreferenceConfig_r16                             `optional,setuprelease`
	maxCC_PreferenceConfig_r16              *MaxCC_PreferenceConfig_r16                             `optional,setuprelease`
	maxMIMO_LayerPreferenceConfig_r16       *MaxMIMO_LayerPreferenceConfig_r16                      `optional,setuprelease`
	minSchedulingOffsetPreferenceConfig_r16 *MinSchedulingOffsetPreferenceConfig_r16                `optional,setuprelease`
	releasePreferenceConfig_r16             *ReleasePreferenceConfig_r16                            `optional,setuprelease`
	referenceTimePreferenceReporting_r16    *OtherConfig_v1610_referenceTimePreferenceReporting_r16 `optional`
	btNameList_r16                          *BT_NameList_r16                                        `optional,setuprelease`
	wlanNameList_r16                        *WLAN_NameList_r16                                      `optional,setuprelease`
	sensorNameList_r16                      *Sensor_NameList_r16                                    `optional,setuprelease`
	obtainCommonLocation_r16                *OtherConfig_v1610_obtainCommonLocation_r16             `optional`
	sl_AssistanceConfigNR_r16               *OtherConfig_v1610_sl_AssistanceConfigNR_r16            `optional`
}

func (ie *OtherConfig_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idc_AssistanceConfig_r16 != nil, ie.drx_PreferenceConfig_r16 != nil, ie.maxBW_PreferenceConfig_r16 != nil, ie.maxCC_PreferenceConfig_r16 != nil, ie.maxMIMO_LayerPreferenceConfig_r16 != nil, ie.minSchedulingOffsetPreferenceConfig_r16 != nil, ie.releasePreferenceConfig_r16 != nil, ie.referenceTimePreferenceReporting_r16 != nil, ie.btNameList_r16 != nil, ie.wlanNameList_r16 != nil, ie.sensorNameList_r16 != nil, ie.obtainCommonLocation_r16 != nil, ie.sl_AssistanceConfigNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idc_AssistanceConfig_r16 != nil {
		tmp_idc_AssistanceConfig_r16 := utils.SetupRelease[*IDC_AssistanceConfig_r16]{
			Setup: ie.idc_AssistanceConfig_r16,
		}
		if err = tmp_idc_AssistanceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idc_AssistanceConfig_r16", err)
		}
	}
	if ie.drx_PreferenceConfig_r16 != nil {
		tmp_drx_PreferenceConfig_r16 := utils.SetupRelease[*DRX_PreferenceConfig_r16]{
			Setup: ie.drx_PreferenceConfig_r16,
		}
		if err = tmp_drx_PreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode drx_PreferenceConfig_r16", err)
		}
	}
	if ie.maxBW_PreferenceConfig_r16 != nil {
		tmp_maxBW_PreferenceConfig_r16 := utils.SetupRelease[*MaxBW_PreferenceConfig_r16]{
			Setup: ie.maxBW_PreferenceConfig_r16,
		}
		if err = tmp_maxBW_PreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxBW_PreferenceConfig_r16", err)
		}
	}
	if ie.maxCC_PreferenceConfig_r16 != nil {
		tmp_maxCC_PreferenceConfig_r16 := utils.SetupRelease[*MaxCC_PreferenceConfig_r16]{
			Setup: ie.maxCC_PreferenceConfig_r16,
		}
		if err = tmp_maxCC_PreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxCC_PreferenceConfig_r16", err)
		}
	}
	if ie.maxMIMO_LayerPreferenceConfig_r16 != nil {
		tmp_maxMIMO_LayerPreferenceConfig_r16 := utils.SetupRelease[*MaxMIMO_LayerPreferenceConfig_r16]{
			Setup: ie.maxMIMO_LayerPreferenceConfig_r16,
		}
		if err = tmp_maxMIMO_LayerPreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxMIMO_LayerPreferenceConfig_r16", err)
		}
	}
	if ie.minSchedulingOffsetPreferenceConfig_r16 != nil {
		tmp_minSchedulingOffsetPreferenceConfig_r16 := utils.SetupRelease[*MinSchedulingOffsetPreferenceConfig_r16]{
			Setup: ie.minSchedulingOffsetPreferenceConfig_r16,
		}
		if err = tmp_minSchedulingOffsetPreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode minSchedulingOffsetPreferenceConfig_r16", err)
		}
	}
	if ie.releasePreferenceConfig_r16 != nil {
		tmp_releasePreferenceConfig_r16 := utils.SetupRelease[*ReleasePreferenceConfig_r16]{
			Setup: ie.releasePreferenceConfig_r16,
		}
		if err = tmp_releasePreferenceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode releasePreferenceConfig_r16", err)
		}
	}
	if ie.referenceTimePreferenceReporting_r16 != nil {
		if err = ie.referenceTimePreferenceReporting_r16.Encode(w); err != nil {
			return utils.WrapError("Encode referenceTimePreferenceReporting_r16", err)
		}
	}
	if ie.btNameList_r16 != nil {
		tmp_btNameList_r16 := utils.SetupRelease[*BT_NameList_r16]{
			Setup: ie.btNameList_r16,
		}
		if err = tmp_btNameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode btNameList_r16", err)
		}
	}
	if ie.wlanNameList_r16 != nil {
		tmp_wlanNameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{
			Setup: ie.wlanNameList_r16,
		}
		if err = tmp_wlanNameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode wlanNameList_r16", err)
		}
	}
	if ie.sensorNameList_r16 != nil {
		tmp_sensorNameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{
			Setup: ie.sensorNameList_r16,
		}
		if err = tmp_sensorNameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sensorNameList_r16", err)
		}
	}
	if ie.obtainCommonLocation_r16 != nil {
		if err = ie.obtainCommonLocation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode obtainCommonLocation_r16", err)
		}
	}
	if ie.sl_AssistanceConfigNR_r16 != nil {
		if err = ie.sl_AssistanceConfigNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_AssistanceConfigNR_r16", err)
		}
	}
	return nil
}

func (ie *OtherConfig_v1610) Decode(r *uper.UperReader) error {
	var err error
	var idc_AssistanceConfig_r16Present bool
	if idc_AssistanceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_PreferenceConfig_r16Present bool
	if drx_PreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxBW_PreferenceConfig_r16Present bool
	if maxBW_PreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCC_PreferenceConfig_r16Present bool
	if maxCC_PreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMIMO_LayerPreferenceConfig_r16Present bool
	if maxMIMO_LayerPreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var minSchedulingOffsetPreferenceConfig_r16Present bool
	if minSchedulingOffsetPreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var releasePreferenceConfig_r16Present bool
	if releasePreferenceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var referenceTimePreferenceReporting_r16Present bool
	if referenceTimePreferenceReporting_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var btNameList_r16Present bool
	if btNameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var wlanNameList_r16Present bool
	if wlanNameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sensorNameList_r16Present bool
	if sensorNameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var obtainCommonLocation_r16Present bool
	if obtainCommonLocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_AssistanceConfigNR_r16Present bool
	if sl_AssistanceConfigNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if idc_AssistanceConfig_r16Present {
		tmp_idc_AssistanceConfig_r16 := utils.SetupRelease[*IDC_AssistanceConfig_r16]{}
		if err = tmp_idc_AssistanceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idc_AssistanceConfig_r16", err)
		}
		ie.idc_AssistanceConfig_r16 = tmp_idc_AssistanceConfig_r16.Setup
	}
	if drx_PreferenceConfig_r16Present {
		tmp_drx_PreferenceConfig_r16 := utils.SetupRelease[*DRX_PreferenceConfig_r16]{}
		if err = tmp_drx_PreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode drx_PreferenceConfig_r16", err)
		}
		ie.drx_PreferenceConfig_r16 = tmp_drx_PreferenceConfig_r16.Setup
	}
	if maxBW_PreferenceConfig_r16Present {
		tmp_maxBW_PreferenceConfig_r16 := utils.SetupRelease[*MaxBW_PreferenceConfig_r16]{}
		if err = tmp_maxBW_PreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxBW_PreferenceConfig_r16", err)
		}
		ie.maxBW_PreferenceConfig_r16 = tmp_maxBW_PreferenceConfig_r16.Setup
	}
	if maxCC_PreferenceConfig_r16Present {
		tmp_maxCC_PreferenceConfig_r16 := utils.SetupRelease[*MaxCC_PreferenceConfig_r16]{}
		if err = tmp_maxCC_PreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxCC_PreferenceConfig_r16", err)
		}
		ie.maxCC_PreferenceConfig_r16 = tmp_maxCC_PreferenceConfig_r16.Setup
	}
	if maxMIMO_LayerPreferenceConfig_r16Present {
		tmp_maxMIMO_LayerPreferenceConfig_r16 := utils.SetupRelease[*MaxMIMO_LayerPreferenceConfig_r16]{}
		if err = tmp_maxMIMO_LayerPreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxMIMO_LayerPreferenceConfig_r16", err)
		}
		ie.maxMIMO_LayerPreferenceConfig_r16 = tmp_maxMIMO_LayerPreferenceConfig_r16.Setup
	}
	if minSchedulingOffsetPreferenceConfig_r16Present {
		tmp_minSchedulingOffsetPreferenceConfig_r16 := utils.SetupRelease[*MinSchedulingOffsetPreferenceConfig_r16]{}
		if err = tmp_minSchedulingOffsetPreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode minSchedulingOffsetPreferenceConfig_r16", err)
		}
		ie.minSchedulingOffsetPreferenceConfig_r16 = tmp_minSchedulingOffsetPreferenceConfig_r16.Setup
	}
	if releasePreferenceConfig_r16Present {
		tmp_releasePreferenceConfig_r16 := utils.SetupRelease[*ReleasePreferenceConfig_r16]{}
		if err = tmp_releasePreferenceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode releasePreferenceConfig_r16", err)
		}
		ie.releasePreferenceConfig_r16 = tmp_releasePreferenceConfig_r16.Setup
	}
	if referenceTimePreferenceReporting_r16Present {
		ie.referenceTimePreferenceReporting_r16 = new(OtherConfig_v1610_referenceTimePreferenceReporting_r16)
		if err = ie.referenceTimePreferenceReporting_r16.Decode(r); err != nil {
			return utils.WrapError("Decode referenceTimePreferenceReporting_r16", err)
		}
	}
	if btNameList_r16Present {
		tmp_btNameList_r16 := utils.SetupRelease[*BT_NameList_r16]{}
		if err = tmp_btNameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode btNameList_r16", err)
		}
		ie.btNameList_r16 = tmp_btNameList_r16.Setup
	}
	if wlanNameList_r16Present {
		tmp_wlanNameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{}
		if err = tmp_wlanNameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode wlanNameList_r16", err)
		}
		ie.wlanNameList_r16 = tmp_wlanNameList_r16.Setup
	}
	if sensorNameList_r16Present {
		tmp_sensorNameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{}
		if err = tmp_sensorNameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sensorNameList_r16", err)
		}
		ie.sensorNameList_r16 = tmp_sensorNameList_r16.Setup
	}
	if obtainCommonLocation_r16Present {
		ie.obtainCommonLocation_r16 = new(OtherConfig_v1610_obtainCommonLocation_r16)
		if err = ie.obtainCommonLocation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode obtainCommonLocation_r16", err)
		}
	}
	if sl_AssistanceConfigNR_r16Present {
		ie.sl_AssistanceConfigNR_r16 = new(OtherConfig_v1610_sl_AssistanceConfigNR_r16)
		if err = ie.sl_AssistanceConfigNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_AssistanceConfigNR_r16", err)
		}
	}
	return nil
}
