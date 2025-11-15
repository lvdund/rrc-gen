package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDT_MAC_PHY_CG_Config_r17 struct {
	cg_SDT_ConfigLCH_RestrictionToAddModList_r17  []CG_SDT_ConfigLCH_Restriction_r17 `lb:1,ub:maxLC_ID,optional`
	cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 []LogicalChannelIdentity           `lb:1,ub:maxLC_ID,optional`
	cg_SDT_ConfigInitialBWP_NUL_r17               *BWP_UplinkDedicatedSDT_r17        `optional,setuprelease`
	cg_SDT_ConfigInitialBWP_SUL_r17               *BWP_UplinkDedicatedSDT_r17        `optional,setuprelease`
	cg_SDT_ConfigInitialBWP_DL_r17                *BWP_DownlinkDedicatedSDT_r17      `optional`
	cg_SDT_TimeAlignmentTimer_r17                 *TimeAlignmentTimer                `optional`
	cg_SDT_RSRP_ThresholdSSB_r17                  *RSRP_Range                        `optional`
	cg_SDT_TA_ValidationConfig_r17                *CG_SDT_TA_ValidationConfig_r17    `optional,setuprelease`
	cg_SDT_CS_RNTI_r17                            *RNTI_Value                        `optional`
}

func (ie *SDT_MAC_PHY_CG_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.cg_SDT_ConfigLCH_RestrictionToAddModList_r17) > 0, len(ie.cg_SDT_ConfigLCH_RestrictionToReleaseList_r17) > 0, ie.cg_SDT_ConfigInitialBWP_NUL_r17 != nil, ie.cg_SDT_ConfigInitialBWP_SUL_r17 != nil, ie.cg_SDT_ConfigInitialBWP_DL_r17 != nil, ie.cg_SDT_TimeAlignmentTimer_r17 != nil, ie.cg_SDT_RSRP_ThresholdSSB_r17 != nil, ie.cg_SDT_TA_ValidationConfig_r17 != nil, ie.cg_SDT_CS_RNTI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.cg_SDT_ConfigLCH_RestrictionToAddModList_r17) > 0 {
		tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17 := utils.NewSequence[*CG_SDT_ConfigLCH_Restriction_r17]([]*CG_SDT_ConfigLCH_Restriction_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.cg_SDT_ConfigLCH_RestrictionToAddModList_r17 {
			tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Value = append(tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Value, &i)
		}
		if err = tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_ConfigLCH_RestrictionToAddModList_r17", err)
		}
	}
	if len(ie.cg_SDT_ConfigLCH_RestrictionToReleaseList_r17) > 0 {
		tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 {
			tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Value = append(tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Value, &i)
		}
		if err = tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_ConfigLCH_RestrictionToReleaseList_r17", err)
		}
	}
	if ie.cg_SDT_ConfigInitialBWP_NUL_r17 != nil {
		tmp_cg_SDT_ConfigInitialBWP_NUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{
			Setup: ie.cg_SDT_ConfigInitialBWP_NUL_r17,
		}
		if err = tmp_cg_SDT_ConfigInitialBWP_NUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_ConfigInitialBWP_NUL_r17", err)
		}
	}
	if ie.cg_SDT_ConfigInitialBWP_SUL_r17 != nil {
		tmp_cg_SDT_ConfigInitialBWP_SUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{
			Setup: ie.cg_SDT_ConfigInitialBWP_SUL_r17,
		}
		if err = tmp_cg_SDT_ConfigInitialBWP_SUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_ConfigInitialBWP_SUL_r17", err)
		}
	}
	if ie.cg_SDT_ConfigInitialBWP_DL_r17 != nil {
		if err = ie.cg_SDT_ConfigInitialBWP_DL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_ConfigInitialBWP_DL_r17", err)
		}
	}
	if ie.cg_SDT_TimeAlignmentTimer_r17 != nil {
		if err = ie.cg_SDT_TimeAlignmentTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_TimeAlignmentTimer_r17", err)
		}
	}
	if ie.cg_SDT_RSRP_ThresholdSSB_r17 != nil {
		if err = ie.cg_SDT_RSRP_ThresholdSSB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_RSRP_ThresholdSSB_r17", err)
		}
	}
	if ie.cg_SDT_TA_ValidationConfig_r17 != nil {
		tmp_cg_SDT_TA_ValidationConfig_r17 := utils.SetupRelease[*CG_SDT_TA_ValidationConfig_r17]{
			Setup: ie.cg_SDT_TA_ValidationConfig_r17,
		}
		if err = tmp_cg_SDT_TA_ValidationConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_TA_ValidationConfig_r17", err)
		}
	}
	if ie.cg_SDT_CS_RNTI_r17 != nil {
		if err = ie.cg_SDT_CS_RNTI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_SDT_CS_RNTI_r17", err)
		}
	}
	return nil
}

func (ie *SDT_MAC_PHY_CG_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var cg_SDT_ConfigLCH_RestrictionToAddModList_r17Present bool
	if cg_SDT_ConfigLCH_RestrictionToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_ConfigLCH_RestrictionToReleaseList_r17Present bool
	if cg_SDT_ConfigLCH_RestrictionToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_ConfigInitialBWP_NUL_r17Present bool
	if cg_SDT_ConfigInitialBWP_NUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_ConfigInitialBWP_SUL_r17Present bool
	if cg_SDT_ConfigInitialBWP_SUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_ConfigInitialBWP_DL_r17Present bool
	if cg_SDT_ConfigInitialBWP_DL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_TimeAlignmentTimer_r17Present bool
	if cg_SDT_TimeAlignmentTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_RSRP_ThresholdSSB_r17Present bool
	if cg_SDT_RSRP_ThresholdSSB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_TA_ValidationConfig_r17Present bool
	if cg_SDT_TA_ValidationConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_SDT_CS_RNTI_r17Present bool
	if cg_SDT_CS_RNTI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cg_SDT_ConfigLCH_RestrictionToAddModList_r17Present {
		tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17 := utils.NewSequence[*CG_SDT_ConfigLCH_Restriction_r17]([]*CG_SDT_ConfigLCH_Restriction_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_cg_SDT_ConfigLCH_RestrictionToAddModList_r17 := func() *CG_SDT_ConfigLCH_Restriction_r17 {
			return new(CG_SDT_ConfigLCH_Restriction_r17)
		}
		if err = tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Decode(r, fn_cg_SDT_ConfigLCH_RestrictionToAddModList_r17); err != nil {
			return utils.WrapError("Decode cg_SDT_ConfigLCH_RestrictionToAddModList_r17", err)
		}
		ie.cg_SDT_ConfigLCH_RestrictionToAddModList_r17 = []CG_SDT_ConfigLCH_Restriction_r17{}
		for _, i := range tmp_cg_SDT_ConfigLCH_RestrictionToAddModList_r17.Value {
			ie.cg_SDT_ConfigLCH_RestrictionToAddModList_r17 = append(ie.cg_SDT_ConfigLCH_RestrictionToAddModList_r17, *i)
		}
	}
	if cg_SDT_ConfigLCH_RestrictionToReleaseList_r17Present {
		tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 := func() *LogicalChannelIdentity {
			return new(LogicalChannelIdentity)
		}
		if err = tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Decode(r, fn_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17); err != nil {
			return utils.WrapError("Decode cg_SDT_ConfigLCH_RestrictionToReleaseList_r17", err)
		}
		ie.cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 = []LogicalChannelIdentity{}
		for _, i := range tmp_cg_SDT_ConfigLCH_RestrictionToReleaseList_r17.Value {
			ie.cg_SDT_ConfigLCH_RestrictionToReleaseList_r17 = append(ie.cg_SDT_ConfigLCH_RestrictionToReleaseList_r17, *i)
		}
	}
	if cg_SDT_ConfigInitialBWP_NUL_r17Present {
		tmp_cg_SDT_ConfigInitialBWP_NUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{}
		if err = tmp_cg_SDT_ConfigInitialBWP_NUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_ConfigInitialBWP_NUL_r17", err)
		}
		ie.cg_SDT_ConfigInitialBWP_NUL_r17 = tmp_cg_SDT_ConfigInitialBWP_NUL_r17.Setup
	}
	if cg_SDT_ConfigInitialBWP_SUL_r17Present {
		tmp_cg_SDT_ConfigInitialBWP_SUL_r17 := utils.SetupRelease[*BWP_UplinkDedicatedSDT_r17]{}
		if err = tmp_cg_SDT_ConfigInitialBWP_SUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_ConfigInitialBWP_SUL_r17", err)
		}
		ie.cg_SDT_ConfigInitialBWP_SUL_r17 = tmp_cg_SDT_ConfigInitialBWP_SUL_r17.Setup
	}
	if cg_SDT_ConfigInitialBWP_DL_r17Present {
		ie.cg_SDT_ConfigInitialBWP_DL_r17 = new(BWP_DownlinkDedicatedSDT_r17)
		if err = ie.cg_SDT_ConfigInitialBWP_DL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_ConfigInitialBWP_DL_r17", err)
		}
	}
	if cg_SDT_TimeAlignmentTimer_r17Present {
		ie.cg_SDT_TimeAlignmentTimer_r17 = new(TimeAlignmentTimer)
		if err = ie.cg_SDT_TimeAlignmentTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_TimeAlignmentTimer_r17", err)
		}
	}
	if cg_SDT_RSRP_ThresholdSSB_r17Present {
		ie.cg_SDT_RSRP_ThresholdSSB_r17 = new(RSRP_Range)
		if err = ie.cg_SDT_RSRP_ThresholdSSB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_RSRP_ThresholdSSB_r17", err)
		}
	}
	if cg_SDT_TA_ValidationConfig_r17Present {
		tmp_cg_SDT_TA_ValidationConfig_r17 := utils.SetupRelease[*CG_SDT_TA_ValidationConfig_r17]{}
		if err = tmp_cg_SDT_TA_ValidationConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_TA_ValidationConfig_r17", err)
		}
		ie.cg_SDT_TA_ValidationConfig_r17 = tmp_cg_SDT_TA_ValidationConfig_r17.Setup
	}
	if cg_SDT_CS_RNTI_r17Present {
		ie.cg_SDT_CS_RNTI_r17 = new(RNTI_Value)
		if err = ie.cg_SDT_CS_RNTI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cg_SDT_CS_RNTI_r17", err)
		}
	}
	return nil
}
