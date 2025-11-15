package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationSidelink_v1700_IEs struct {
	sl_DRX_ConfigUC_PC5_r17            *SL_DRX_ConfigUC_r17           `optional,setuprelease`
	sl_LatencyBoundIUC_Report_r17      *SL_LatencyBoundIUC_Report_r17 `optional,setuprelease`
	sl_RLC_ChannelToReleaseListPC5_r17 []SL_RLC_ChannelID_r17         `lb:1,ub:maxSL_LCID_r16,optional`
	sl_RLC_ChannelToAddModListPC5_r17  []SL_RLC_ChannelConfigPC5_r17  `lb:1,ub:maxSL_LCID_r16,optional`
	nonCriticalExtension               interface{}                    `optional`
}

func (ie *RRCReconfigurationSidelink_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_DRX_ConfigUC_PC5_r17 != nil, ie.sl_LatencyBoundIUC_Report_r17 != nil, len(ie.sl_RLC_ChannelToReleaseListPC5_r17) > 0, len(ie.sl_RLC_ChannelToAddModListPC5_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_DRX_ConfigUC_PC5_r17 != nil {
		tmp_sl_DRX_ConfigUC_PC5_r17 := utils.SetupRelease[*SL_DRX_ConfigUC_r17]{
			Setup: ie.sl_DRX_ConfigUC_PC5_r17,
		}
		if err = tmp_sl_DRX_ConfigUC_PC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DRX_ConfigUC_PC5_r17", err)
		}
	}
	if ie.sl_LatencyBoundIUC_Report_r17 != nil {
		tmp_sl_LatencyBoundIUC_Report_r17 := utils.SetupRelease[*SL_LatencyBoundIUC_Report_r17]{
			Setup: ie.sl_LatencyBoundIUC_Report_r17,
		}
		if err = tmp_sl_LatencyBoundIUC_Report_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_LatencyBoundIUC_Report_r17", err)
		}
	}
	if len(ie.sl_RLC_ChannelToReleaseListPC5_r17) > 0 {
		tmp_sl_RLC_ChannelToReleaseListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelID_r17]([]*SL_RLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.sl_RLC_ChannelToReleaseListPC5_r17 {
			tmp_sl_RLC_ChannelToReleaseListPC5_r17.Value = append(tmp_sl_RLC_ChannelToReleaseListPC5_r17.Value, &i)
		}
		if err = tmp_sl_RLC_ChannelToReleaseListPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_ChannelToReleaseListPC5_r17", err)
		}
	}
	if len(ie.sl_RLC_ChannelToAddModListPC5_r17) > 0 {
		tmp_sl_RLC_ChannelToAddModListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelConfigPC5_r17]([]*SL_RLC_ChannelConfigPC5_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.sl_RLC_ChannelToAddModListPC5_r17 {
			tmp_sl_RLC_ChannelToAddModListPC5_r17.Value = append(tmp_sl_RLC_ChannelToAddModListPC5_r17.Value, &i)
		}
		if err = tmp_sl_RLC_ChannelToAddModListPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_ChannelToAddModListPC5_r17", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationSidelink_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_DRX_ConfigUC_PC5_r17Present bool
	if sl_DRX_ConfigUC_PC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_LatencyBoundIUC_Report_r17Present bool
	if sl_LatencyBoundIUC_Report_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_ChannelToReleaseListPC5_r17Present bool
	if sl_RLC_ChannelToReleaseListPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_ChannelToAddModListPC5_r17Present bool
	if sl_RLC_ChannelToAddModListPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_DRX_ConfigUC_PC5_r17Present {
		tmp_sl_DRX_ConfigUC_PC5_r17 := utils.SetupRelease[*SL_DRX_ConfigUC_r17]{}
		if err = tmp_sl_DRX_ConfigUC_PC5_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DRX_ConfigUC_PC5_r17", err)
		}
		ie.sl_DRX_ConfigUC_PC5_r17 = tmp_sl_DRX_ConfigUC_PC5_r17.Setup
	}
	if sl_LatencyBoundIUC_Report_r17Present {
		tmp_sl_LatencyBoundIUC_Report_r17 := utils.SetupRelease[*SL_LatencyBoundIUC_Report_r17]{}
		if err = tmp_sl_LatencyBoundIUC_Report_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_LatencyBoundIUC_Report_r17", err)
		}
		ie.sl_LatencyBoundIUC_Report_r17 = tmp_sl_LatencyBoundIUC_Report_r17.Setup
	}
	if sl_RLC_ChannelToReleaseListPC5_r17Present {
		tmp_sl_RLC_ChannelToReleaseListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelID_r17]([]*SL_RLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_sl_RLC_ChannelToReleaseListPC5_r17 := func() *SL_RLC_ChannelID_r17 {
			return new(SL_RLC_ChannelID_r17)
		}
		if err = tmp_sl_RLC_ChannelToReleaseListPC5_r17.Decode(r, fn_sl_RLC_ChannelToReleaseListPC5_r17); err != nil {
			return utils.WrapError("Decode sl_RLC_ChannelToReleaseListPC5_r17", err)
		}
		ie.sl_RLC_ChannelToReleaseListPC5_r17 = []SL_RLC_ChannelID_r17{}
		for _, i := range tmp_sl_RLC_ChannelToReleaseListPC5_r17.Value {
			ie.sl_RLC_ChannelToReleaseListPC5_r17 = append(ie.sl_RLC_ChannelToReleaseListPC5_r17, *i)
		}
	}
	if sl_RLC_ChannelToAddModListPC5_r17Present {
		tmp_sl_RLC_ChannelToAddModListPC5_r17 := utils.NewSequence[*SL_RLC_ChannelConfigPC5_r17]([]*SL_RLC_ChannelConfigPC5_r17{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_sl_RLC_ChannelToAddModListPC5_r17 := func() *SL_RLC_ChannelConfigPC5_r17 {
			return new(SL_RLC_ChannelConfigPC5_r17)
		}
		if err = tmp_sl_RLC_ChannelToAddModListPC5_r17.Decode(r, fn_sl_RLC_ChannelToAddModListPC5_r17); err != nil {
			return utils.WrapError("Decode sl_RLC_ChannelToAddModListPC5_r17", err)
		}
		ie.sl_RLC_ChannelToAddModListPC5_r17 = []SL_RLC_ChannelConfigPC5_r17{}
		for _, i := range tmp_sl_RLC_ChannelToAddModListPC5_r17.Value {
			ie.sl_RLC_ChannelToAddModListPC5_r17 = append(ie.sl_RLC_ChannelToAddModListPC5_r17, *i)
		}
	}
	return nil
}
