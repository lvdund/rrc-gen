package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationSidelink_r16_IEs struct {
	slrb_ConfigToAddModList_r16   []SLRB_Config_r16                                      `lb:1,ub:maxNrofSLRB_r16,optional`
	slrb_ConfigToReleaseList_r16  []SLRB_PC5_ConfigIndex_r16                             `lb:1,ub:maxNrofSLRB_r16,optional`
	sl_MeasConfig_r16             *SL_MeasConfig_r16                                     `optional,setuprelease`
	sl_CSI_RS_Config_r16          *SL_CSI_RS_Config_r16                                  `optional,setuprelease`
	sl_ResetConfig_r16            *RRCReconfigurationSidelink_r16_IEs_sl_ResetConfig_r16 `optional`
	sl_LatencyBoundCSI_Report_r16 *int64                                                 `lb:3,ub:160,optional`
	lateNonCriticalExtension      *[]byte                                                `optional`
	nonCriticalExtension          *RRCReconfigurationSidelink_v1700_IEs                  `optional`
}

func (ie *RRCReconfigurationSidelink_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.slrb_ConfigToAddModList_r16) > 0, len(ie.slrb_ConfigToReleaseList_r16) > 0, ie.sl_MeasConfig_r16 != nil, ie.sl_CSI_RS_Config_r16 != nil, ie.sl_ResetConfig_r16 != nil, ie.sl_LatencyBoundCSI_Report_r16 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.slrb_ConfigToAddModList_r16) > 0 {
		tmp_slrb_ConfigToAddModList_r16 := utils.NewSequence[*SLRB_Config_r16]([]*SLRB_Config_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.slrb_ConfigToAddModList_r16 {
			tmp_slrb_ConfigToAddModList_r16.Value = append(tmp_slrb_ConfigToAddModList_r16.Value, &i)
		}
		if err = tmp_slrb_ConfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode slrb_ConfigToAddModList_r16", err)
		}
	}
	if len(ie.slrb_ConfigToReleaseList_r16) > 0 {
		tmp_slrb_ConfigToReleaseList_r16 := utils.NewSequence[*SLRB_PC5_ConfigIndex_r16]([]*SLRB_PC5_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.slrb_ConfigToReleaseList_r16 {
			tmp_slrb_ConfigToReleaseList_r16.Value = append(tmp_slrb_ConfigToReleaseList_r16.Value, &i)
		}
		if err = tmp_slrb_ConfigToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode slrb_ConfigToReleaseList_r16", err)
		}
	}
	if ie.sl_MeasConfig_r16 != nil {
		tmp_sl_MeasConfig_r16 := utils.SetupRelease[*SL_MeasConfig_r16]{
			Setup: ie.sl_MeasConfig_r16,
		}
		if err = tmp_sl_MeasConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasConfig_r16", err)
		}
	}
	if ie.sl_CSI_RS_Config_r16 != nil {
		tmp_sl_CSI_RS_Config_r16 := utils.SetupRelease[*SL_CSI_RS_Config_r16]{
			Setup: ie.sl_CSI_RS_Config_r16,
		}
		if err = tmp_sl_CSI_RS_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CSI_RS_Config_r16", err)
		}
	}
	if ie.sl_ResetConfig_r16 != nil {
		if err = ie.sl_ResetConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ResetConfig_r16", err)
		}
	}
	if ie.sl_LatencyBoundCSI_Report_r16 != nil {
		if err = w.WriteInteger(*ie.sl_LatencyBoundCSI_Report_r16, &uper.Constraint{Lb: 3, Ub: 160}, false); err != nil {
			return utils.WrapError("Encode sl_LatencyBoundCSI_Report_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationSidelink_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var slrb_ConfigToAddModList_r16Present bool
	if slrb_ConfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var slrb_ConfigToReleaseList_r16Present bool
	if slrb_ConfigToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasConfig_r16Present bool
	if sl_MeasConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CSI_RS_Config_r16Present bool
	if sl_CSI_RS_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ResetConfig_r16Present bool
	if sl_ResetConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_LatencyBoundCSI_Report_r16Present bool
	if sl_LatencyBoundCSI_Report_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if slrb_ConfigToAddModList_r16Present {
		tmp_slrb_ConfigToAddModList_r16 := utils.NewSequence[*SLRB_Config_r16]([]*SLRB_Config_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_slrb_ConfigToAddModList_r16 := func() *SLRB_Config_r16 {
			return new(SLRB_Config_r16)
		}
		if err = tmp_slrb_ConfigToAddModList_r16.Decode(r, fn_slrb_ConfigToAddModList_r16); err != nil {
			return utils.WrapError("Decode slrb_ConfigToAddModList_r16", err)
		}
		ie.slrb_ConfigToAddModList_r16 = []SLRB_Config_r16{}
		for _, i := range tmp_slrb_ConfigToAddModList_r16.Value {
			ie.slrb_ConfigToAddModList_r16 = append(ie.slrb_ConfigToAddModList_r16, *i)
		}
	}
	if slrb_ConfigToReleaseList_r16Present {
		tmp_slrb_ConfigToReleaseList_r16 := utils.NewSequence[*SLRB_PC5_ConfigIndex_r16]([]*SLRB_PC5_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_slrb_ConfigToReleaseList_r16 := func() *SLRB_PC5_ConfigIndex_r16 {
			return new(SLRB_PC5_ConfigIndex_r16)
		}
		if err = tmp_slrb_ConfigToReleaseList_r16.Decode(r, fn_slrb_ConfigToReleaseList_r16); err != nil {
			return utils.WrapError("Decode slrb_ConfigToReleaseList_r16", err)
		}
		ie.slrb_ConfigToReleaseList_r16 = []SLRB_PC5_ConfigIndex_r16{}
		for _, i := range tmp_slrb_ConfigToReleaseList_r16.Value {
			ie.slrb_ConfigToReleaseList_r16 = append(ie.slrb_ConfigToReleaseList_r16, *i)
		}
	}
	if sl_MeasConfig_r16Present {
		tmp_sl_MeasConfig_r16 := utils.SetupRelease[*SL_MeasConfig_r16]{}
		if err = tmp_sl_MeasConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasConfig_r16", err)
		}
		ie.sl_MeasConfig_r16 = tmp_sl_MeasConfig_r16.Setup
	}
	if sl_CSI_RS_Config_r16Present {
		tmp_sl_CSI_RS_Config_r16 := utils.SetupRelease[*SL_CSI_RS_Config_r16]{}
		if err = tmp_sl_CSI_RS_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CSI_RS_Config_r16", err)
		}
		ie.sl_CSI_RS_Config_r16 = tmp_sl_CSI_RS_Config_r16.Setup
	}
	if sl_ResetConfig_r16Present {
		ie.sl_ResetConfig_r16 = new(RRCReconfigurationSidelink_r16_IEs_sl_ResetConfig_r16)
		if err = ie.sl_ResetConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ResetConfig_r16", err)
		}
	}
	if sl_LatencyBoundCSI_Report_r16Present {
		var tmp_int_sl_LatencyBoundCSI_Report_r16 int64
		if tmp_int_sl_LatencyBoundCSI_Report_r16, err = r.ReadInteger(&uper.Constraint{Lb: 3, Ub: 160}, false); err != nil {
			return utils.WrapError("Decode sl_LatencyBoundCSI_Report_r16", err)
		}
		ie.sl_LatencyBoundCSI_Report_r16 = &tmp_int_sl_LatencyBoundCSI_Report_r16
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfigurationSidelink_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
