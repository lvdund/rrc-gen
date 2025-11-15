package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1700_IEs struct {
	ul_GapFR2_Preference_r17             *UL_GapFR2_Preference_r17                                         `optional`
	musim_Assistance_r17                 *MUSIM_Assistance_r17                                             `optional`
	overheatingAssistance_r17            *OverheatingAssistance_r17                                        `optional`
	maxBW_PreferenceFR2_2_r17            *MaxBW_PreferenceFR2_2_r17                                        `optional`
	maxMIMO_LayerPreferenceFR2_2_r17     *MaxMIMO_LayerPreferenceFR2_2_r17                                 `optional`
	minSchedulingOffsetPreferenceExt_r17 *MinSchedulingOffsetPreferenceExt_r17                             `optional`
	rlm_MeasRelaxationState_r17          *bool                                                             `optional`
	bfd_MeasRelaxationState_r17          *uper.BitString                                                   `lb:1,ub:maxNrofServingCells,optional`
	nonSDT_DataIndication_r17            *UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17      `optional`
	scg_DeactivationPreference_r17       *UEAssistanceInformation_v1700_IEs_scg_DeactivationPreference_r17 `optional`
	uplinkData_r17                       *UEAssistanceInformation_v1700_IEs_uplinkData_r17                 `optional`
	rrm_MeasRelaxationFulfilment_r17     *bool                                                             `optional`
	propagationDelayDifference_r17       *PropagationDelayDifference_r17                                   `optional`
	nonCriticalExtension                 interface{}                                                       `optional`
}

func (ie *UEAssistanceInformation_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_GapFR2_Preference_r17 != nil, ie.musim_Assistance_r17 != nil, ie.overheatingAssistance_r17 != nil, ie.maxBW_PreferenceFR2_2_r17 != nil, ie.maxMIMO_LayerPreferenceFR2_2_r17 != nil, ie.minSchedulingOffsetPreferenceExt_r17 != nil, ie.rlm_MeasRelaxationState_r17 != nil, ie.bfd_MeasRelaxationState_r17 != nil, ie.nonSDT_DataIndication_r17 != nil, ie.scg_DeactivationPreference_r17 != nil, ie.uplinkData_r17 != nil, ie.rrm_MeasRelaxationFulfilment_r17 != nil, ie.propagationDelayDifference_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_GapFR2_Preference_r17 != nil {
		if err = ie.ul_GapFR2_Preference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_GapFR2_Preference_r17", err)
		}
	}
	if ie.musim_Assistance_r17 != nil {
		if err = ie.musim_Assistance_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_Assistance_r17", err)
		}
	}
	if ie.overheatingAssistance_r17 != nil {
		if err = ie.overheatingAssistance_r17.Encode(w); err != nil {
			return utils.WrapError("Encode overheatingAssistance_r17", err)
		}
	}
	if ie.maxBW_PreferenceFR2_2_r17 != nil {
		if err = ie.maxBW_PreferenceFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxBW_PreferenceFR2_2_r17", err)
		}
	}
	if ie.maxMIMO_LayerPreferenceFR2_2_r17 != nil {
		if err = ie.maxMIMO_LayerPreferenceFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxMIMO_LayerPreferenceFR2_2_r17", err)
		}
	}
	if ie.minSchedulingOffsetPreferenceExt_r17 != nil {
		if err = ie.minSchedulingOffsetPreferenceExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode minSchedulingOffsetPreferenceExt_r17", err)
		}
	}
	if ie.rlm_MeasRelaxationState_r17 != nil {
		if err = w.WriteBoolean(*ie.rlm_MeasRelaxationState_r17); err != nil {
			return utils.WrapError("Encode rlm_MeasRelaxationState_r17", err)
		}
	}
	if ie.bfd_MeasRelaxationState_r17 != nil {
		if err = w.WriteBitString(ie.bfd_MeasRelaxationState_r17.Bytes, uint(ie.bfd_MeasRelaxationState_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Encode bfd_MeasRelaxationState_r17", err)
		}
	}
	if ie.nonSDT_DataIndication_r17 != nil {
		if err = ie.nonSDT_DataIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nonSDT_DataIndication_r17", err)
		}
	}
	if ie.scg_DeactivationPreference_r17 != nil {
		if err = ie.scg_DeactivationPreference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_DeactivationPreference_r17", err)
		}
	}
	if ie.uplinkData_r17 != nil {
		if err = ie.uplinkData_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkData_r17", err)
		}
	}
	if ie.rrm_MeasRelaxationFulfilment_r17 != nil {
		if err = w.WriteBoolean(*ie.rrm_MeasRelaxationFulfilment_r17); err != nil {
			return utils.WrapError("Encode rrm_MeasRelaxationFulfilment_r17", err)
		}
	}
	if ie.propagationDelayDifference_r17 != nil {
		if err = ie.propagationDelayDifference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode propagationDelayDifference_r17", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ul_GapFR2_Preference_r17Present bool
	if ul_GapFR2_Preference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_Assistance_r17Present bool
	if musim_Assistance_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var overheatingAssistance_r17Present bool
	if overheatingAssistance_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxBW_PreferenceFR2_2_r17Present bool
	if maxBW_PreferenceFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMIMO_LayerPreferenceFR2_2_r17Present bool
	if maxMIMO_LayerPreferenceFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var minSchedulingOffsetPreferenceExt_r17Present bool
	if minSchedulingOffsetPreferenceExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlm_MeasRelaxationState_r17Present bool
	if rlm_MeasRelaxationState_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bfd_MeasRelaxationState_r17Present bool
	if bfd_MeasRelaxationState_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonSDT_DataIndication_r17Present bool
	if nonSDT_DataIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_DeactivationPreference_r17Present bool
	if scg_DeactivationPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkData_r17Present bool
	if uplinkData_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rrm_MeasRelaxationFulfilment_r17Present bool
	if rrm_MeasRelaxationFulfilment_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var propagationDelayDifference_r17Present bool
	if propagationDelayDifference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_GapFR2_Preference_r17Present {
		ie.ul_GapFR2_Preference_r17 = new(UL_GapFR2_Preference_r17)
		if err = ie.ul_GapFR2_Preference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_GapFR2_Preference_r17", err)
		}
	}
	if musim_Assistance_r17Present {
		ie.musim_Assistance_r17 = new(MUSIM_Assistance_r17)
		if err = ie.musim_Assistance_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_Assistance_r17", err)
		}
	}
	if overheatingAssistance_r17Present {
		ie.overheatingAssistance_r17 = new(OverheatingAssistance_r17)
		if err = ie.overheatingAssistance_r17.Decode(r); err != nil {
			return utils.WrapError("Decode overheatingAssistance_r17", err)
		}
	}
	if maxBW_PreferenceFR2_2_r17Present {
		ie.maxBW_PreferenceFR2_2_r17 = new(MaxBW_PreferenceFR2_2_r17)
		if err = ie.maxBW_PreferenceFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxBW_PreferenceFR2_2_r17", err)
		}
	}
	if maxMIMO_LayerPreferenceFR2_2_r17Present {
		ie.maxMIMO_LayerPreferenceFR2_2_r17 = new(MaxMIMO_LayerPreferenceFR2_2_r17)
		if err = ie.maxMIMO_LayerPreferenceFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxMIMO_LayerPreferenceFR2_2_r17", err)
		}
	}
	if minSchedulingOffsetPreferenceExt_r17Present {
		ie.minSchedulingOffsetPreferenceExt_r17 = new(MinSchedulingOffsetPreferenceExt_r17)
		if err = ie.minSchedulingOffsetPreferenceExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode minSchedulingOffsetPreferenceExt_r17", err)
		}
	}
	if rlm_MeasRelaxationState_r17Present {
		var tmp_bool_rlm_MeasRelaxationState_r17 bool
		if tmp_bool_rlm_MeasRelaxationState_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode rlm_MeasRelaxationState_r17", err)
		}
		ie.rlm_MeasRelaxationState_r17 = &tmp_bool_rlm_MeasRelaxationState_r17
	}
	if bfd_MeasRelaxationState_r17Present {
		var tmp_bs_bfd_MeasRelaxationState_r17 []byte
		var n_bfd_MeasRelaxationState_r17 uint
		if tmp_bs_bfd_MeasRelaxationState_r17, n_bfd_MeasRelaxationState_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Decode bfd_MeasRelaxationState_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_bfd_MeasRelaxationState_r17,
			NumBits: uint64(n_bfd_MeasRelaxationState_r17),
		}
		ie.bfd_MeasRelaxationState_r17 = &tmp_bitstring
	}
	if nonSDT_DataIndication_r17Present {
		ie.nonSDT_DataIndication_r17 = new(UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17)
		if err = ie.nonSDT_DataIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nonSDT_DataIndication_r17", err)
		}
	}
	if scg_DeactivationPreference_r17Present {
		ie.scg_DeactivationPreference_r17 = new(UEAssistanceInformation_v1700_IEs_scg_DeactivationPreference_r17)
		if err = ie.scg_DeactivationPreference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_DeactivationPreference_r17", err)
		}
	}
	if uplinkData_r17Present {
		ie.uplinkData_r17 = new(UEAssistanceInformation_v1700_IEs_uplinkData_r17)
		if err = ie.uplinkData_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkData_r17", err)
		}
	}
	if rrm_MeasRelaxationFulfilment_r17Present {
		var tmp_bool_rrm_MeasRelaxationFulfilment_r17 bool
		if tmp_bool_rrm_MeasRelaxationFulfilment_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode rrm_MeasRelaxationFulfilment_r17", err)
		}
		ie.rrm_MeasRelaxationFulfilment_r17 = &tmp_bool_rrm_MeasRelaxationFulfilment_r17
	}
	if propagationDelayDifference_r17Present {
		ie.propagationDelayDifference_r17 = new(PropagationDelayDifference_r17)
		if err = ie.propagationDelayDifference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode propagationDelayDifference_r17", err)
		}
	}
	return nil
}
