package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1720 struct {
	parallelTxSRS_PUCCH_PUSCH_intraBand_r17           *CA_ParametersNR_v1720_parallelTxSRS_PUCCH_PUSCH_intraBand_r17           `optional`
	parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17     *CA_ParametersNR_v1720_parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17     `optional`
	semiStaticPUCCH_CellSwitchSingleGroup_r17         *CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17         `optional`
	semiStaticPUCCH_CellSwitchTwoGroups_r17           []TwoPUCCH_Grp_Configurations_r17                                        `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r17,optional`
	dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17  *CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17  `optional`
	dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17  *CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17  `optional`
	dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17    []TwoPUCCH_Grp_Configurations_r17                                        `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r17,optional`
	dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17    []TwoPUCCH_Grp_Configurations_r17                                        `lb:1,ub:maxTwoPUCCH_Grp_ConfigList_r17,optional`
	ack_NACK_FeedbackForMulticast_r17                 *CA_ParametersNR_v1720_ack_NACK_FeedbackForMulticast_r17                 `optional`
	ptp_Retx_Multicast_r17                            *CA_ParametersNR_v1720_ptp_Retx_Multicast_r17                            `optional`
	nack_OnlyFeedbackForMulticast_r17                 *CA_ParametersNR_v1720_nack_OnlyFeedbackForMulticast_r17                 `optional`
	nack_OnlyFeedbackSpecificResourceForMulticast_r17 *CA_ParametersNR_v1720_nack_OnlyFeedbackSpecificResourceForMulticast_r17 `optional`
	ack_NACK_FeedbackForSPS_Multicast_r17             *CA_ParametersNR_v1720_ack_NACK_FeedbackForSPS_Multicast_r17             `optional`
	ptp_Retx_SPS_Multicast_r17                        *CA_ParametersNR_v1720_ptp_Retx_SPS_Multicast_r17                        `optional`
	higherPowerLimit_r17                              *CA_ParametersNR_v1720_higherPowerLimit_r17                              `optional`
	parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17      *CA_ParametersNR_v1720_parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17      `optional`
	pdcch_MonitoringCA_r17                            *int64                                                                   `lb:4,ub:16,optional`
	pdcch_BlindDetectionMCG_SCG_List_r17              []PDCCH_BlindDetectionMCG_SCG_r17                                        `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
	pdcch_BlindDetectionMixedList1_r17                []PDCCH_BlindDetectionMixed_r17                                          `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
	pdcch_BlindDetectionMixedList2_r17                []PDCCH_BlindDetectionMixed_r17                                          `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
	pdcch_BlindDetectionMixedList3_r17                []PDCCH_BlindDetectionMixed1_r17                                         `lb:1,ub:maxNrofPdcch_BlindDetection_r17,optional`
}

func (ie *CA_ParametersNR_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.parallelTxSRS_PUCCH_PUSCH_intraBand_r17 != nil, ie.parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17 != nil, ie.semiStaticPUCCH_CellSwitchSingleGroup_r17 != nil, len(ie.semiStaticPUCCH_CellSwitchTwoGroups_r17) > 0, ie.dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 != nil, ie.dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 != nil, len(ie.dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17) > 0, len(ie.dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17) > 0, ie.ack_NACK_FeedbackForMulticast_r17 != nil, ie.ptp_Retx_Multicast_r17 != nil, ie.nack_OnlyFeedbackForMulticast_r17 != nil, ie.nack_OnlyFeedbackSpecificResourceForMulticast_r17 != nil, ie.ack_NACK_FeedbackForSPS_Multicast_r17 != nil, ie.ptp_Retx_SPS_Multicast_r17 != nil, ie.higherPowerLimit_r17 != nil, ie.parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17 != nil, ie.pdcch_MonitoringCA_r17 != nil, len(ie.pdcch_BlindDetectionMCG_SCG_List_r17) > 0, len(ie.pdcch_BlindDetectionMixedList1_r17) > 0, len(ie.pdcch_BlindDetectionMixedList2_r17) > 0, len(ie.pdcch_BlindDetectionMixedList3_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.parallelTxSRS_PUCCH_PUSCH_intraBand_r17 != nil {
		if err = ie.parallelTxSRS_PUCCH_PUSCH_intraBand_r17.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxSRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ie.parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17 != nil {
		if err = ie.parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ie.semiStaticPUCCH_CellSwitchSingleGroup_r17 != nil {
		if err = ie.semiStaticPUCCH_CellSwitchSingleGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode semiStaticPUCCH_CellSwitchSingleGroup_r17", err)
		}
	}
	if len(ie.semiStaticPUCCH_CellSwitchTwoGroups_r17) > 0 {
		tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		for _, i := range ie.semiStaticPUCCH_CellSwitchTwoGroups_r17 {
			tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17.Value = append(tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17.Value, &i)
		}
		if err = tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode semiStaticPUCCH_CellSwitchTwoGroups_r17", err)
		}
	}
	if ie.dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 != nil {
		if err = ie.dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17", err)
		}
	}
	if ie.dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 != nil {
		if err = ie.dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17", err)
		}
	}
	if len(ie.dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17) > 0 {
		tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		for _, i := range ie.dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 {
			tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Value = append(tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Value, &i)
		}
		if err = tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17", err)
		}
	}
	if len(ie.dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17) > 0 {
		tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		for _, i := range ie.dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 {
			tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Value = append(tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Value, &i)
		}
		if err = tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17", err)
		}
	}
	if ie.ack_NACK_FeedbackForMulticast_r17 != nil {
		if err = ie.ack_NACK_FeedbackForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ack_NACK_FeedbackForMulticast_r17", err)
		}
	}
	if ie.ptp_Retx_Multicast_r17 != nil {
		if err = ie.ptp_Retx_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ptp_Retx_Multicast_r17", err)
		}
	}
	if ie.nack_OnlyFeedbackForMulticast_r17 != nil {
		if err = ie.nack_OnlyFeedbackForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nack_OnlyFeedbackForMulticast_r17", err)
		}
	}
	if ie.nack_OnlyFeedbackSpecificResourceForMulticast_r17 != nil {
		if err = ie.nack_OnlyFeedbackSpecificResourceForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nack_OnlyFeedbackSpecificResourceForMulticast_r17", err)
		}
	}
	if ie.ack_NACK_FeedbackForSPS_Multicast_r17 != nil {
		if err = ie.ack_NACK_FeedbackForSPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ack_NACK_FeedbackForSPS_Multicast_r17", err)
		}
	}
	if ie.ptp_Retx_SPS_Multicast_r17 != nil {
		if err = ie.ptp_Retx_SPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ptp_Retx_SPS_Multicast_r17", err)
		}
	}
	if ie.higherPowerLimit_r17 != nil {
		if err = ie.higherPowerLimit_r17.Encode(w); err != nil {
			return utils.WrapError("Encode higherPowerLimit_r17", err)
		}
	}
	if ie.parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17 != nil {
		if err = ie.parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if ie.pdcch_MonitoringCA_r17 != nil {
		if err = w.WriteInteger(*ie.pdcch_MonitoringCA_r17, &uper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode pdcch_MonitoringCA_r17", err)
		}
	}
	if len(ie.pdcch_BlindDetectionMCG_SCG_List_r17) > 0 {
		tmp_pdcch_BlindDetectionMCG_SCG_List_r17 := utils.NewSequence[*PDCCH_BlindDetectionMCG_SCG_r17]([]*PDCCH_BlindDetectionMCG_SCG_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.pdcch_BlindDetectionMCG_SCG_List_r17 {
			tmp_pdcch_BlindDetectionMCG_SCG_List_r17.Value = append(tmp_pdcch_BlindDetectionMCG_SCG_List_r17.Value, &i)
		}
		if err = tmp_pdcch_BlindDetectionMCG_SCG_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionMCG_SCG_List_r17", err)
		}
	}
	if len(ie.pdcch_BlindDetectionMixedList1_r17) > 0 {
		tmp_pdcch_BlindDetectionMixedList1_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.pdcch_BlindDetectionMixedList1_r17 {
			tmp_pdcch_BlindDetectionMixedList1_r17.Value = append(tmp_pdcch_BlindDetectionMixedList1_r17.Value, &i)
		}
		if err = tmp_pdcch_BlindDetectionMixedList1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionMixedList1_r17", err)
		}
	}
	if len(ie.pdcch_BlindDetectionMixedList2_r17) > 0 {
		tmp_pdcch_BlindDetectionMixedList2_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.pdcch_BlindDetectionMixedList2_r17 {
			tmp_pdcch_BlindDetectionMixedList2_r17.Value = append(tmp_pdcch_BlindDetectionMixedList2_r17.Value, &i)
		}
		if err = tmp_pdcch_BlindDetectionMixedList2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionMixedList2_r17", err)
		}
	}
	if len(ie.pdcch_BlindDetectionMixedList3_r17) > 0 {
		tmp_pdcch_BlindDetectionMixedList3_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed1_r17]([]*PDCCH_BlindDetectionMixed1_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		for _, i := range ie.pdcch_BlindDetectionMixedList3_r17 {
			tmp_pdcch_BlindDetectionMixedList3_r17.Value = append(tmp_pdcch_BlindDetectionMixedList3_r17.Value, &i)
		}
		if err = tmp_pdcch_BlindDetectionMixedList3_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionMixedList3_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1720) Decode(r *uper.UperReader) error {
	var err error
	var parallelTxSRS_PUCCH_PUSCH_intraBand_r17Present bool
	if parallelTxSRS_PUCCH_PUSCH_intraBand_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17Present bool
	if parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var semiStaticPUCCH_CellSwitchSingleGroup_r17Present bool
	if semiStaticPUCCH_CellSwitchSingleGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var semiStaticPUCCH_CellSwitchTwoGroups_r17Present bool
	if semiStaticPUCCH_CellSwitchTwoGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17Present bool
	if dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17Present bool
	if dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17Present bool
	if dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17Present bool
	if dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ack_NACK_FeedbackForMulticast_r17Present bool
	if ack_NACK_FeedbackForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ptp_Retx_Multicast_r17Present bool
	if ptp_Retx_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nack_OnlyFeedbackForMulticast_r17Present bool
	if nack_OnlyFeedbackForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nack_OnlyFeedbackSpecificResourceForMulticast_r17Present bool
	if nack_OnlyFeedbackSpecificResourceForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ack_NACK_FeedbackForSPS_Multicast_r17Present bool
	if ack_NACK_FeedbackForSPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ptp_Retx_SPS_Multicast_r17Present bool
	if ptp_Retx_SPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var higherPowerLimit_r17Present bool
	if higherPowerLimit_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17Present bool
	if parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_MonitoringCA_r17Present bool
	if pdcch_MonitoringCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionMCG_SCG_List_r17Present bool
	if pdcch_BlindDetectionMCG_SCG_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionMixedList1_r17Present bool
	if pdcch_BlindDetectionMixedList1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionMixedList2_r17Present bool
	if pdcch_BlindDetectionMixedList2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionMixedList3_r17Present bool
	if pdcch_BlindDetectionMixedList3_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if parallelTxSRS_PUCCH_PUSCH_intraBand_r17Present {
		ie.parallelTxSRS_PUCCH_PUSCH_intraBand_r17 = new(CA_ParametersNR_v1720_parallelTxSRS_PUCCH_PUSCH_intraBand_r17)
		if err = ie.parallelTxSRS_PUCCH_PUSCH_intraBand_r17.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxSRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17Present {
		ie.parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17 = new(CA_ParametersNR_v1720_parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17)
		if err = ie.parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxPRACH_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if semiStaticPUCCH_CellSwitchSingleGroup_r17Present {
		ie.semiStaticPUCCH_CellSwitchSingleGroup_r17 = new(CA_ParametersNR_v1720_semiStaticPUCCH_CellSwitchSingleGroup_r17)
		if err = ie.semiStaticPUCCH_CellSwitchSingleGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode semiStaticPUCCH_CellSwitchSingleGroup_r17", err)
		}
	}
	if semiStaticPUCCH_CellSwitchTwoGroups_r17Present {
		tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		fn_semiStaticPUCCH_CellSwitchTwoGroups_r17 := func() *TwoPUCCH_Grp_Configurations_r17 {
			return new(TwoPUCCH_Grp_Configurations_r17)
		}
		if err = tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17.Decode(r, fn_semiStaticPUCCH_CellSwitchTwoGroups_r17); err != nil {
			return utils.WrapError("Decode semiStaticPUCCH_CellSwitchTwoGroups_r17", err)
		}
		ie.semiStaticPUCCH_CellSwitchTwoGroups_r17 = []TwoPUCCH_Grp_Configurations_r17{}
		for _, i := range tmp_semiStaticPUCCH_CellSwitchTwoGroups_r17.Value {
			ie.semiStaticPUCCH_CellSwitchTwoGroups_r17 = append(ie.semiStaticPUCCH_CellSwitchTwoGroups_r17, *i)
		}
	}
	if dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17Present {
		ie.dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 = new(CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17)
		if err = ie.dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17", err)
		}
	}
	if dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17Present {
		ie.dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 = new(CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17)
		if err = ie.dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17", err)
		}
	}
	if dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17Present {
		tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		fn_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 := func() *TwoPUCCH_Grp_Configurations_r17 {
			return new(TwoPUCCH_Grp_Configurations_r17)
		}
		if err = tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Decode(r, fn_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17); err != nil {
			return utils.WrapError("Decode dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17", err)
		}
		ie.dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 = []TwoPUCCH_Grp_Configurations_r17{}
		for _, i := range tmp_dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17.Value {
			ie.dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 = append(ie.dynamicPUCCH_CellSwitchSameLengthTwoGroups_r17, *i)
		}
	}
	if dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17Present {
		tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 := utils.NewSequence[*TwoPUCCH_Grp_Configurations_r17]([]*TwoPUCCH_Grp_Configurations_r17{}, uper.Constraint{Lb: 1, Ub: maxTwoPUCCH_Grp_ConfigList_r17}, false)
		fn_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 := func() *TwoPUCCH_Grp_Configurations_r17 {
			return new(TwoPUCCH_Grp_Configurations_r17)
		}
		if err = tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Decode(r, fn_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17); err != nil {
			return utils.WrapError("Decode dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17", err)
		}
		ie.dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 = []TwoPUCCH_Grp_Configurations_r17{}
		for _, i := range tmp_dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17.Value {
			ie.dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 = append(ie.dynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17, *i)
		}
	}
	if ack_NACK_FeedbackForMulticast_r17Present {
		ie.ack_NACK_FeedbackForMulticast_r17 = new(CA_ParametersNR_v1720_ack_NACK_FeedbackForMulticast_r17)
		if err = ie.ack_NACK_FeedbackForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ack_NACK_FeedbackForMulticast_r17", err)
		}
	}
	if ptp_Retx_Multicast_r17Present {
		ie.ptp_Retx_Multicast_r17 = new(CA_ParametersNR_v1720_ptp_Retx_Multicast_r17)
		if err = ie.ptp_Retx_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ptp_Retx_Multicast_r17", err)
		}
	}
	if nack_OnlyFeedbackForMulticast_r17Present {
		ie.nack_OnlyFeedbackForMulticast_r17 = new(CA_ParametersNR_v1720_nack_OnlyFeedbackForMulticast_r17)
		if err = ie.nack_OnlyFeedbackForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nack_OnlyFeedbackForMulticast_r17", err)
		}
	}
	if nack_OnlyFeedbackSpecificResourceForMulticast_r17Present {
		ie.nack_OnlyFeedbackSpecificResourceForMulticast_r17 = new(CA_ParametersNR_v1720_nack_OnlyFeedbackSpecificResourceForMulticast_r17)
		if err = ie.nack_OnlyFeedbackSpecificResourceForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nack_OnlyFeedbackSpecificResourceForMulticast_r17", err)
		}
	}
	if ack_NACK_FeedbackForSPS_Multicast_r17Present {
		ie.ack_NACK_FeedbackForSPS_Multicast_r17 = new(CA_ParametersNR_v1720_ack_NACK_FeedbackForSPS_Multicast_r17)
		if err = ie.ack_NACK_FeedbackForSPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ack_NACK_FeedbackForSPS_Multicast_r17", err)
		}
	}
	if ptp_Retx_SPS_Multicast_r17Present {
		ie.ptp_Retx_SPS_Multicast_r17 = new(CA_ParametersNR_v1720_ptp_Retx_SPS_Multicast_r17)
		if err = ie.ptp_Retx_SPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ptp_Retx_SPS_Multicast_r17", err)
		}
	}
	if higherPowerLimit_r17Present {
		ie.higherPowerLimit_r17 = new(CA_ParametersNR_v1720_higherPowerLimit_r17)
		if err = ie.higherPowerLimit_r17.Decode(r); err != nil {
			return utils.WrapError("Decode higherPowerLimit_r17", err)
		}
	}
	if parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17Present {
		ie.parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17 = new(CA_ParametersNR_v1720_parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17)
		if err = ie.parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxMsgA_SRS_PUCCH_PUSCH_intraBand_r17", err)
		}
	}
	if pdcch_MonitoringCA_r17Present {
		var tmp_int_pdcch_MonitoringCA_r17 int64
		if tmp_int_pdcch_MonitoringCA_r17, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode pdcch_MonitoringCA_r17", err)
		}
		ie.pdcch_MonitoringCA_r17 = &tmp_int_pdcch_MonitoringCA_r17
	}
	if pdcch_BlindDetectionMCG_SCG_List_r17Present {
		tmp_pdcch_BlindDetectionMCG_SCG_List_r17 := utils.NewSequence[*PDCCH_BlindDetectionMCG_SCG_r17]([]*PDCCH_BlindDetectionMCG_SCG_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_pdcch_BlindDetectionMCG_SCG_List_r17 := func() *PDCCH_BlindDetectionMCG_SCG_r17 {
			return new(PDCCH_BlindDetectionMCG_SCG_r17)
		}
		if err = tmp_pdcch_BlindDetectionMCG_SCG_List_r17.Decode(r, fn_pdcch_BlindDetectionMCG_SCG_List_r17); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionMCG_SCG_List_r17", err)
		}
		ie.pdcch_BlindDetectionMCG_SCG_List_r17 = []PDCCH_BlindDetectionMCG_SCG_r17{}
		for _, i := range tmp_pdcch_BlindDetectionMCG_SCG_List_r17.Value {
			ie.pdcch_BlindDetectionMCG_SCG_List_r17 = append(ie.pdcch_BlindDetectionMCG_SCG_List_r17, *i)
		}
	}
	if pdcch_BlindDetectionMixedList1_r17Present {
		tmp_pdcch_BlindDetectionMixedList1_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_pdcch_BlindDetectionMixedList1_r17 := func() *PDCCH_BlindDetectionMixed_r17 {
			return new(PDCCH_BlindDetectionMixed_r17)
		}
		if err = tmp_pdcch_BlindDetectionMixedList1_r17.Decode(r, fn_pdcch_BlindDetectionMixedList1_r17); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionMixedList1_r17", err)
		}
		ie.pdcch_BlindDetectionMixedList1_r17 = []PDCCH_BlindDetectionMixed_r17{}
		for _, i := range tmp_pdcch_BlindDetectionMixedList1_r17.Value {
			ie.pdcch_BlindDetectionMixedList1_r17 = append(ie.pdcch_BlindDetectionMixedList1_r17, *i)
		}
	}
	if pdcch_BlindDetectionMixedList2_r17Present {
		tmp_pdcch_BlindDetectionMixedList2_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed_r17]([]*PDCCH_BlindDetectionMixed_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_pdcch_BlindDetectionMixedList2_r17 := func() *PDCCH_BlindDetectionMixed_r17 {
			return new(PDCCH_BlindDetectionMixed_r17)
		}
		if err = tmp_pdcch_BlindDetectionMixedList2_r17.Decode(r, fn_pdcch_BlindDetectionMixedList2_r17); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionMixedList2_r17", err)
		}
		ie.pdcch_BlindDetectionMixedList2_r17 = []PDCCH_BlindDetectionMixed_r17{}
		for _, i := range tmp_pdcch_BlindDetectionMixedList2_r17.Value {
			ie.pdcch_BlindDetectionMixedList2_r17 = append(ie.pdcch_BlindDetectionMixedList2_r17, *i)
		}
	}
	if pdcch_BlindDetectionMixedList3_r17Present {
		tmp_pdcch_BlindDetectionMixedList3_r17 := utils.NewSequence[*PDCCH_BlindDetectionMixed1_r17]([]*PDCCH_BlindDetectionMixed1_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetection_r17}, false)
		fn_pdcch_BlindDetectionMixedList3_r17 := func() *PDCCH_BlindDetectionMixed1_r17 {
			return new(PDCCH_BlindDetectionMixed1_r17)
		}
		if err = tmp_pdcch_BlindDetectionMixedList3_r17.Decode(r, fn_pdcch_BlindDetectionMixedList3_r17); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionMixedList3_r17", err)
		}
		ie.pdcch_BlindDetectionMixedList3_r17 = []PDCCH_BlindDetectionMixed1_r17{}
		for _, i := range tmp_pdcch_BlindDetectionMixedList3_r17.Value {
			ie.pdcch_BlindDetectionMixedList3_r17 = append(ie.pdcch_BlindDetectionMixedList3_r17, *i)
		}
	}
	return nil
}
