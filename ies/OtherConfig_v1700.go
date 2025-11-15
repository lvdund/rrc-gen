package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_v1700 struct {
	ul_GapFR2_PreferenceConfig_r17             *OtherConfig_v1700_ul_GapFR2_PreferenceConfig_r17             `optional`
	musim_GapAssistanceConfig_r17              *MUSIM_GapAssistanceConfig_r17                                `optional,setuprelease`
	musim_LeaveAssistanceConfig_r17            *MUSIM_LeaveAssistanceConfig_r17                              `optional,setuprelease`
	successHO_Config_r17                       *SuccessHO_Config_r17                                         `optional,setuprelease`
	maxBW_PreferenceConfigFR2_2_r17            *OtherConfig_v1700_maxBW_PreferenceConfigFR2_2_r17            `optional`
	maxMIMO_LayerPreferenceConfigFR2_2_r17     *OtherConfig_v1700_maxMIMO_LayerPreferenceConfigFR2_2_r17     `optional`
	minSchedulingOffsetPreferenceConfigExt_r17 *OtherConfig_v1700_minSchedulingOffsetPreferenceConfigExt_r17 `optional`
	rlm_RelaxationReportingConfig_r17          *RLM_RelaxationReportingConfig_r17                            `optional,setuprelease`
	bfd_RelaxationReportingConfig_r17          *BFD_RelaxationReportingConfig_r17                            `optional,setuprelease`
	scg_DeactivationPreferenceConfig_r17       *SCG_DeactivationPreferenceConfig_r17                         `optional,setuprelease`
	rrm_MeasRelaxationReportingConfig_r17      *RRM_MeasRelaxationReportingConfig_r17                        `optional,setuprelease`
	propDelayDiffReportConfig_r17              *PropDelayDiffReportConfig_r17                                `optional,setuprelease`
}

func (ie *OtherConfig_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_GapFR2_PreferenceConfig_r17 != nil, ie.musim_GapAssistanceConfig_r17 != nil, ie.musim_LeaveAssistanceConfig_r17 != nil, ie.successHO_Config_r17 != nil, ie.maxBW_PreferenceConfigFR2_2_r17 != nil, ie.maxMIMO_LayerPreferenceConfigFR2_2_r17 != nil, ie.minSchedulingOffsetPreferenceConfigExt_r17 != nil, ie.rlm_RelaxationReportingConfig_r17 != nil, ie.bfd_RelaxationReportingConfig_r17 != nil, ie.scg_DeactivationPreferenceConfig_r17 != nil, ie.rrm_MeasRelaxationReportingConfig_r17 != nil, ie.propDelayDiffReportConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_GapFR2_PreferenceConfig_r17 != nil {
		if err = ie.ul_GapFR2_PreferenceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_GapFR2_PreferenceConfig_r17", err)
		}
	}
	if ie.musim_GapAssistanceConfig_r17 != nil {
		tmp_musim_GapAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_GapAssistanceConfig_r17]{
			Setup: ie.musim_GapAssistanceConfig_r17,
		}
		if err = tmp_musim_GapAssistanceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapAssistanceConfig_r17", err)
		}
	}
	if ie.musim_LeaveAssistanceConfig_r17 != nil {
		tmp_musim_LeaveAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_LeaveAssistanceConfig_r17]{
			Setup: ie.musim_LeaveAssistanceConfig_r17,
		}
		if err = tmp_musim_LeaveAssistanceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_LeaveAssistanceConfig_r17", err)
		}
	}
	if ie.successHO_Config_r17 != nil {
		tmp_successHO_Config_r17 := utils.SetupRelease[*SuccessHO_Config_r17]{
			Setup: ie.successHO_Config_r17,
		}
		if err = tmp_successHO_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode successHO_Config_r17", err)
		}
	}
	if ie.maxBW_PreferenceConfigFR2_2_r17 != nil {
		if err = ie.maxBW_PreferenceConfigFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxBW_PreferenceConfigFR2_2_r17", err)
		}
	}
	if ie.maxMIMO_LayerPreferenceConfigFR2_2_r17 != nil {
		if err = ie.maxMIMO_LayerPreferenceConfigFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxMIMO_LayerPreferenceConfigFR2_2_r17", err)
		}
	}
	if ie.minSchedulingOffsetPreferenceConfigExt_r17 != nil {
		if err = ie.minSchedulingOffsetPreferenceConfigExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode minSchedulingOffsetPreferenceConfigExt_r17", err)
		}
	}
	if ie.rlm_RelaxationReportingConfig_r17 != nil {
		tmp_rlm_RelaxationReportingConfig_r17 := utils.SetupRelease[*RLM_RelaxationReportingConfig_r17]{
			Setup: ie.rlm_RelaxationReportingConfig_r17,
		}
		if err = tmp_rlm_RelaxationReportingConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rlm_RelaxationReportingConfig_r17", err)
		}
	}
	if ie.bfd_RelaxationReportingConfig_r17 != nil {
		tmp_bfd_RelaxationReportingConfig_r17 := utils.SetupRelease[*BFD_RelaxationReportingConfig_r17]{
			Setup: ie.bfd_RelaxationReportingConfig_r17,
		}
		if err = tmp_bfd_RelaxationReportingConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bfd_RelaxationReportingConfig_r17", err)
		}
	}
	if ie.scg_DeactivationPreferenceConfig_r17 != nil {
		tmp_scg_DeactivationPreferenceConfig_r17 := utils.SetupRelease[*SCG_DeactivationPreferenceConfig_r17]{
			Setup: ie.scg_DeactivationPreferenceConfig_r17,
		}
		if err = tmp_scg_DeactivationPreferenceConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_DeactivationPreferenceConfig_r17", err)
		}
	}
	if ie.rrm_MeasRelaxationReportingConfig_r17 != nil {
		tmp_rrm_MeasRelaxationReportingConfig_r17 := utils.SetupRelease[*RRM_MeasRelaxationReportingConfig_r17]{
			Setup: ie.rrm_MeasRelaxationReportingConfig_r17,
		}
		if err = tmp_rrm_MeasRelaxationReportingConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rrm_MeasRelaxationReportingConfig_r17", err)
		}
	}
	if ie.propDelayDiffReportConfig_r17 != nil {
		tmp_propDelayDiffReportConfig_r17 := utils.SetupRelease[*PropDelayDiffReportConfig_r17]{
			Setup: ie.propDelayDiffReportConfig_r17,
		}
		if err = tmp_propDelayDiffReportConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode propDelayDiffReportConfig_r17", err)
		}
	}
	return nil
}

func (ie *OtherConfig_v1700) Decode(r *uper.UperReader) error {
	var err error
	var ul_GapFR2_PreferenceConfig_r17Present bool
	if ul_GapFR2_PreferenceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_GapAssistanceConfig_r17Present bool
	if musim_GapAssistanceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_LeaveAssistanceConfig_r17Present bool
	if musim_LeaveAssistanceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var successHO_Config_r17Present bool
	if successHO_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxBW_PreferenceConfigFR2_2_r17Present bool
	if maxBW_PreferenceConfigFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMIMO_LayerPreferenceConfigFR2_2_r17Present bool
	if maxMIMO_LayerPreferenceConfigFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var minSchedulingOffsetPreferenceConfigExt_r17Present bool
	if minSchedulingOffsetPreferenceConfigExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlm_RelaxationReportingConfig_r17Present bool
	if rlm_RelaxationReportingConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bfd_RelaxationReportingConfig_r17Present bool
	if bfd_RelaxationReportingConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_DeactivationPreferenceConfig_r17Present bool
	if scg_DeactivationPreferenceConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rrm_MeasRelaxationReportingConfig_r17Present bool
	if rrm_MeasRelaxationReportingConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var propDelayDiffReportConfig_r17Present bool
	if propDelayDiffReportConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_GapFR2_PreferenceConfig_r17Present {
		ie.ul_GapFR2_PreferenceConfig_r17 = new(OtherConfig_v1700_ul_GapFR2_PreferenceConfig_r17)
		if err = ie.ul_GapFR2_PreferenceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_GapFR2_PreferenceConfig_r17", err)
		}
	}
	if musim_GapAssistanceConfig_r17Present {
		tmp_musim_GapAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_GapAssistanceConfig_r17]{}
		if err = tmp_musim_GapAssistanceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_GapAssistanceConfig_r17", err)
		}
		ie.musim_GapAssistanceConfig_r17 = tmp_musim_GapAssistanceConfig_r17.Setup
	}
	if musim_LeaveAssistanceConfig_r17Present {
		tmp_musim_LeaveAssistanceConfig_r17 := utils.SetupRelease[*MUSIM_LeaveAssistanceConfig_r17]{}
		if err = tmp_musim_LeaveAssistanceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_LeaveAssistanceConfig_r17", err)
		}
		ie.musim_LeaveAssistanceConfig_r17 = tmp_musim_LeaveAssistanceConfig_r17.Setup
	}
	if successHO_Config_r17Present {
		tmp_successHO_Config_r17 := utils.SetupRelease[*SuccessHO_Config_r17]{}
		if err = tmp_successHO_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode successHO_Config_r17", err)
		}
		ie.successHO_Config_r17 = tmp_successHO_Config_r17.Setup
	}
	if maxBW_PreferenceConfigFR2_2_r17Present {
		ie.maxBW_PreferenceConfigFR2_2_r17 = new(OtherConfig_v1700_maxBW_PreferenceConfigFR2_2_r17)
		if err = ie.maxBW_PreferenceConfigFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxBW_PreferenceConfigFR2_2_r17", err)
		}
	}
	if maxMIMO_LayerPreferenceConfigFR2_2_r17Present {
		ie.maxMIMO_LayerPreferenceConfigFR2_2_r17 = new(OtherConfig_v1700_maxMIMO_LayerPreferenceConfigFR2_2_r17)
		if err = ie.maxMIMO_LayerPreferenceConfigFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxMIMO_LayerPreferenceConfigFR2_2_r17", err)
		}
	}
	if minSchedulingOffsetPreferenceConfigExt_r17Present {
		ie.minSchedulingOffsetPreferenceConfigExt_r17 = new(OtherConfig_v1700_minSchedulingOffsetPreferenceConfigExt_r17)
		if err = ie.minSchedulingOffsetPreferenceConfigExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode minSchedulingOffsetPreferenceConfigExt_r17", err)
		}
	}
	if rlm_RelaxationReportingConfig_r17Present {
		tmp_rlm_RelaxationReportingConfig_r17 := utils.SetupRelease[*RLM_RelaxationReportingConfig_r17]{}
		if err = tmp_rlm_RelaxationReportingConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rlm_RelaxationReportingConfig_r17", err)
		}
		ie.rlm_RelaxationReportingConfig_r17 = tmp_rlm_RelaxationReportingConfig_r17.Setup
	}
	if bfd_RelaxationReportingConfig_r17Present {
		tmp_bfd_RelaxationReportingConfig_r17 := utils.SetupRelease[*BFD_RelaxationReportingConfig_r17]{}
		if err = tmp_bfd_RelaxationReportingConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bfd_RelaxationReportingConfig_r17", err)
		}
		ie.bfd_RelaxationReportingConfig_r17 = tmp_bfd_RelaxationReportingConfig_r17.Setup
	}
	if scg_DeactivationPreferenceConfig_r17Present {
		tmp_scg_DeactivationPreferenceConfig_r17 := utils.SetupRelease[*SCG_DeactivationPreferenceConfig_r17]{}
		if err = tmp_scg_DeactivationPreferenceConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_DeactivationPreferenceConfig_r17", err)
		}
		ie.scg_DeactivationPreferenceConfig_r17 = tmp_scg_DeactivationPreferenceConfig_r17.Setup
	}
	if rrm_MeasRelaxationReportingConfig_r17Present {
		tmp_rrm_MeasRelaxationReportingConfig_r17 := utils.SetupRelease[*RRM_MeasRelaxationReportingConfig_r17]{}
		if err = tmp_rrm_MeasRelaxationReportingConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rrm_MeasRelaxationReportingConfig_r17", err)
		}
		ie.rrm_MeasRelaxationReportingConfig_r17 = tmp_rrm_MeasRelaxationReportingConfig_r17.Setup
	}
	if propDelayDiffReportConfig_r17Present {
		tmp_propDelayDiffReportConfig_r17 := utils.SetupRelease[*PropDelayDiffReportConfig_r17]{}
		if err = tmp_propDelayDiffReportConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode propDelayDiffReportConfig_r17", err)
		}
		ie.propDelayDiffReportConfig_r17 = tmp_propDelayDiffReportConfig_r17.Setup
	}
	return nil
}
