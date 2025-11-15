package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1610_IEs struct {
	idc_Assistance_r16                *IDC_Assistance_r16                `optional`
	drx_Preference_r16                *DRX_Preference_r16                `optional`
	maxBW_Preference_r16              *MaxBW_Preference_r16              `optional`
	maxCC_Preference_r16              *MaxCC_Preference_r16              `optional`
	maxMIMO_LayerPreference_r16       *MaxMIMO_LayerPreference_r16       `optional`
	minSchedulingOffsetPreference_r16 *MinSchedulingOffsetPreference_r16 `optional`
	releasePreference_r16             *ReleasePreference_r16             `optional`
	sl_UE_AssistanceInformationNR_r16 *SL_UE_AssistanceInformationNR_r16 `optional`
	referenceTimeInfoPreference_r16   *bool                              `optional`
	nonCriticalExtension              *UEAssistanceInformation_v1700_IEs `optional`
}

func (ie *UEAssistanceInformation_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idc_Assistance_r16 != nil, ie.drx_Preference_r16 != nil, ie.maxBW_Preference_r16 != nil, ie.maxCC_Preference_r16 != nil, ie.maxMIMO_LayerPreference_r16 != nil, ie.minSchedulingOffsetPreference_r16 != nil, ie.releasePreference_r16 != nil, ie.sl_UE_AssistanceInformationNR_r16 != nil, ie.referenceTimeInfoPreference_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idc_Assistance_r16 != nil {
		if err = ie.idc_Assistance_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idc_Assistance_r16", err)
		}
	}
	if ie.drx_Preference_r16 != nil {
		if err = ie.drx_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode drx_Preference_r16", err)
		}
	}
	if ie.maxBW_Preference_r16 != nil {
		if err = ie.maxBW_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxBW_Preference_r16", err)
		}
	}
	if ie.maxCC_Preference_r16 != nil {
		if err = ie.maxCC_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxCC_Preference_r16", err)
		}
	}
	if ie.maxMIMO_LayerPreference_r16 != nil {
		if err = ie.maxMIMO_LayerPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxMIMO_LayerPreference_r16", err)
		}
	}
	if ie.minSchedulingOffsetPreference_r16 != nil {
		if err = ie.minSchedulingOffsetPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode minSchedulingOffsetPreference_r16", err)
		}
	}
	if ie.releasePreference_r16 != nil {
		if err = ie.releasePreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode releasePreference_r16", err)
		}
	}
	if ie.sl_UE_AssistanceInformationNR_r16 != nil {
		if err = ie.sl_UE_AssistanceInformationNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_UE_AssistanceInformationNR_r16", err)
		}
	}
	if ie.referenceTimeInfoPreference_r16 != nil {
		if err = w.WriteBoolean(*ie.referenceTimeInfoPreference_r16); err != nil {
			return utils.WrapError("Encode referenceTimeInfoPreference_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var idc_Assistance_r16Present bool
	if idc_Assistance_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_Preference_r16Present bool
	if drx_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxBW_Preference_r16Present bool
	if maxBW_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCC_Preference_r16Present bool
	if maxCC_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMIMO_LayerPreference_r16Present bool
	if maxMIMO_LayerPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var minSchedulingOffsetPreference_r16Present bool
	if minSchedulingOffsetPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var releasePreference_r16Present bool
	if releasePreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_UE_AssistanceInformationNR_r16Present bool
	if sl_UE_AssistanceInformationNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var referenceTimeInfoPreference_r16Present bool
	if referenceTimeInfoPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if idc_Assistance_r16Present {
		ie.idc_Assistance_r16 = new(IDC_Assistance_r16)
		if err = ie.idc_Assistance_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idc_Assistance_r16", err)
		}
	}
	if drx_Preference_r16Present {
		ie.drx_Preference_r16 = new(DRX_Preference_r16)
		if err = ie.drx_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode drx_Preference_r16", err)
		}
	}
	if maxBW_Preference_r16Present {
		ie.maxBW_Preference_r16 = new(MaxBW_Preference_r16)
		if err = ie.maxBW_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxBW_Preference_r16", err)
		}
	}
	if maxCC_Preference_r16Present {
		ie.maxCC_Preference_r16 = new(MaxCC_Preference_r16)
		if err = ie.maxCC_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxCC_Preference_r16", err)
		}
	}
	if maxMIMO_LayerPreference_r16Present {
		ie.maxMIMO_LayerPreference_r16 = new(MaxMIMO_LayerPreference_r16)
		if err = ie.maxMIMO_LayerPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxMIMO_LayerPreference_r16", err)
		}
	}
	if minSchedulingOffsetPreference_r16Present {
		ie.minSchedulingOffsetPreference_r16 = new(MinSchedulingOffsetPreference_r16)
		if err = ie.minSchedulingOffsetPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode minSchedulingOffsetPreference_r16", err)
		}
	}
	if releasePreference_r16Present {
		ie.releasePreference_r16 = new(ReleasePreference_r16)
		if err = ie.releasePreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode releasePreference_r16", err)
		}
	}
	if sl_UE_AssistanceInformationNR_r16Present {
		ie.sl_UE_AssistanceInformationNR_r16 = new(SL_UE_AssistanceInformationNR_r16)
		if err = ie.sl_UE_AssistanceInformationNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_UE_AssistanceInformationNR_r16", err)
		}
	}
	if referenceTimeInfoPreference_r16Present {
		var tmp_bool_referenceTimeInfoPreference_r16 bool
		if tmp_bool_referenceTimeInfoPreference_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode referenceTimeInfoPreference_r16", err)
		}
		ie.referenceTimeInfoPreference_r16 = &tmp_bool_referenceTimeInfoPreference_r16
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UEAssistanceInformation_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
