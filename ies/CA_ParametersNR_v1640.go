package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1640 struct {
	uplinkTxDC_TwoCarrierReport_r16                            *CA_ParametersNR_v1640_uplinkTxDC_TwoCarrierReport_r16                            `optional`
	maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16         *PUCCH_Grp_CarrierTypes_r16                                                       `optional`
	maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16         *PUCCH_Grp_CarrierTypes_r16                                                       `optional`
	twoPUCCH_Grp_ConfigurationsList_r16                        []TwoPUCCH_Grp_Configurations_r16                                                 `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r16,optional`
	diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16           *CA_ParametersNR_v1640_diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16           `optional`
	diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 *CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 `optional`
	diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16  *CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16  `optional`
	pdcch_MonitoringCA_NonAlignedSpan_r16                      *int64                                                                            `lb:2,ub:16,optional`
	pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16            *CA_ParametersNR_v1640_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16            `lb:1,ub:15,optional`
}

func (ie *CA_ParametersNR_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxDC_TwoCarrierReport_r16 != nil, ie.maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil, ie.maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil, len(ie.twoPUCCH_Grp_ConfigurationsList_r16) > 0, ie.diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 != nil, ie.diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 != nil, ie.diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 != nil, ie.pdcch_MonitoringCA_NonAlignedSpan_r16 != nil, ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uplinkTxDC_TwoCarrierReport_r16 != nil {
		if err = ie.uplinkTxDC_TwoCarrierReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxDC_TwoCarrierReport_r16", err)
		}
	}
	if ie.maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil {
		if err = ie.maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if ie.maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16 != nil {
		if err = ie.maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if len(ie.twoPUCCH_Grp_ConfigurationsList_r16) > 0 {
		tmp_twoPUCCH_Grp_ConfigurationsList_r16 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r16]([]*TwoPUCCH_Grp_Configurations_r16{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r16}, false)
		for _, i := range ie.twoPUCCH_Grp_ConfigurationsList_r16 {
			tmp_twoPUCCH_Grp_ConfigurationsList_r16.Value = append(tmp_twoPUCCH_Grp_ConfigurationsList_r16.Value, &i)
		}
		if err = tmp_twoPUCCH_Grp_ConfigurationsList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Grp_ConfigurationsList_r16", err)
		}
	}
	if ie.diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 != nil {
		if err = ie.diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16", err)
		}
	}
	if ie.diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 != nil {
		if err = ie.diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16", err)
		}
	}
	if ie.diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 != nil {
		if err = ie.diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16", err)
		}
	}
	if ie.pdcch_MonitoringCA_NonAlignedSpan_r16 != nil {
		if err = w.WriteInteger(*ie.pdcch_MonitoringCA_NonAlignedSpan_r16, &uper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode pdcch_MonitoringCA_NonAlignedSpan_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 != nil {
		if err = ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1640) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxDC_TwoCarrierReport_r16Present bool
	if uplinkTxDC_TwoCarrierReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16Present bool
	if maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16Present bool
	if maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Grp_ConfigurationsList_r16Present bool
	if twoPUCCH_Grp_ConfigurationsList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16Present bool
	if diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16Present bool
	if diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16Present bool
	if diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_MonitoringCA_NonAlignedSpan_r16Present bool
	if pdcch_MonitoringCA_NonAlignedSpan_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16Present bool
	if pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if uplinkTxDC_TwoCarrierReport_r16Present {
		ie.uplinkTxDC_TwoCarrierReport_r16 = new(CA_ParametersNR_v1640_uplinkTxDC_TwoCarrierReport_r16)
		if err = ie.uplinkTxDC_TwoCarrierReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxDC_TwoCarrierReport_r16", err)
		}
	}
	if maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16Present {
		ie.maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16 = new(PUCCH_Grp_CarrierTypes_r16)
		if err = ie.maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxUpTo3Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16Present {
		ie.maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16 = new(PUCCH_Grp_CarrierTypes_r16)
		if err = ie.maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxUpTo4Diff_NumerologiesConfigSinglePUCCH_grp_r16", err)
		}
	}
	if twoPUCCH_Grp_ConfigurationsList_r16Present {
		tmp_twoPUCCH_Grp_ConfigurationsList_r16 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r16]([]*TwoPUCCH_Grp_Configurations_r16{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r16}, false)
		fn_twoPUCCH_Grp_ConfigurationsList_r16 := func() *TwoPUCCH_Grp_Configurations_r16 {
			return new(TwoPUCCH_Grp_Configurations_r16)
		}
		if err = tmp_twoPUCCH_Grp_ConfigurationsList_r16.Decode(r, fn_twoPUCCH_Grp_ConfigurationsList_r16); err != nil {
			return utils.WrapError("Decode twoPUCCH_Grp_ConfigurationsList_r16", err)
		}
		ie.twoPUCCH_Grp_ConfigurationsList_r16 = []TwoPUCCH_Grp_Configurations_r16{}
		for _, i := range tmp_twoPUCCH_Grp_ConfigurationsList_r16.Value {
			ie.twoPUCCH_Grp_ConfigurationsList_r16 = append(ie.twoPUCCH_Grp_ConfigurationsList_r16, *i)
		}
	}
	if diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16Present {
		ie.diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16 = new(CA_ParametersNR_v1640_diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16)
		if err = ie.diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode diffNumerologyAcrossPUCCH_Group_CarrierTypes_r16", err)
		}
	}
	if diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16Present {
		ie.diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16 = new(CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16)
		if err = ie.diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode diffNumerologyWithinPUCCH_GroupSmallerSCS_CarrierTypes_r16", err)
		}
	}
	if diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16Present {
		ie.diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16 = new(CA_ParametersNR_v1640_diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16)
		if err = ie.diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode diffNumerologyWithinPUCCH_GroupLargerSCS_CarrierTypes_r16", err)
		}
	}
	if pdcch_MonitoringCA_NonAlignedSpan_r16Present {
		var tmp_int_pdcch_MonitoringCA_NonAlignedSpan_r16 int64
		if tmp_int_pdcch_MonitoringCA_NonAlignedSpan_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode pdcch_MonitoringCA_NonAlignedSpan_r16", err)
		}
		ie.pdcch_MonitoringCA_NonAlignedSpan_r16 = &tmp_int_pdcch_MonitoringCA_NonAlignedSpan_r16
	}
	if pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16Present {
		ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16 = new(CA_ParametersNR_v1640_pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16)
		if err = ie.pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA_Mixed_NonAlignedSpan_r16", err)
		}
	}
	return nil
}
