package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1730 struct {
	dmrs_BundlingPUSCH_RepTypeAPerBC_r17                  *CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeAPerBC_r17                  `optional`
	dmrs_BundlingPUSCH_RepTypeBPerBC_r17                  *CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeBPerBC_r17                  `optional`
	dmrs_BundlingPUSCH_multiSlotPerBC_r17                 *CA_ParametersNR_v1730_dmrs_BundlingPUSCH_multiSlotPerBC_r17                 `optional`
	dmrs_BundlingPUCCH_RepPerBC_r17                       *CA_ParametersNR_v1730_dmrs_BundlingPUCCH_RepPerBC_r17                       `optional`
	dmrs_BundlingRestartPerBC_r17                         *CA_ParametersNR_v1730_dmrs_BundlingRestartPerBC_r17                         `optional`
	dmrs_BundlingNonBackToBackTX_PerBC_r17                *CA_ParametersNR_v1730_dmrs_BundlingNonBackToBackTX_PerBC_r17                `optional`
	stayOnTargetCC_SRS_CarrierSwitch_r17                  *CA_ParametersNR_v1730_stayOnTargetCC_SRS_CarrierSwitch_r17                  `optional`
	fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17       *CA_ParametersNR_v1730_fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17       `optional`
	mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 *CA_ParametersNR_v1730_mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 `optional`
	mode1_ForType1_CodebookGeneration_r17                 *CA_ParametersNR_v1730_mode1_ForType1_CodebookGeneration_r17                 `optional`
	nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 *CA_ParametersNR_v1730_nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 `optional`
	multiPUCCH_ConfigForMulticast_r17                     *CA_ParametersNR_v1730_multiPUCCH_ConfigForMulticast_r17                     `optional`
	pucch_ConfigForSPS_Multicast_r17                      *CA_ParametersNR_v1730_pucch_ConfigForSPS_Multicast_r17                      `optional`
	maxNumberG_RNTI_HARQ_ACK_Codebook_r17                 *int64                                                                       `lb:1,ub:4,optional`
	mux_HARQ_ACK_UnicastMulticast_r17                     *CA_ParametersNR_v1730_mux_HARQ_ACK_UnicastMulticast_r17                     `optional`
}

func (ie *CA_ParametersNR_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dmrs_BundlingPUSCH_RepTypeAPerBC_r17 != nil, ie.dmrs_BundlingPUSCH_RepTypeBPerBC_r17 != nil, ie.dmrs_BundlingPUSCH_multiSlotPerBC_r17 != nil, ie.dmrs_BundlingPUCCH_RepPerBC_r17 != nil, ie.dmrs_BundlingRestartPerBC_r17 != nil, ie.dmrs_BundlingNonBackToBackTX_PerBC_r17 != nil, ie.stayOnTargetCC_SRS_CarrierSwitch_r17 != nil, ie.fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil, ie.mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil, ie.mode1_ForType1_CodebookGeneration_r17 != nil, ie.nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 != nil, ie.multiPUCCH_ConfigForMulticast_r17 != nil, ie.pucch_ConfigForSPS_Multicast_r17 != nil, ie.maxNumberG_RNTI_HARQ_ACK_Codebook_r17 != nil, ie.mux_HARQ_ACK_UnicastMulticast_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dmrs_BundlingPUSCH_RepTypeAPerBC_r17 != nil {
		if err = ie.dmrs_BundlingPUSCH_RepTypeAPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_BundlingPUSCH_RepTypeAPerBC_r17", err)
		}
	}
	if ie.dmrs_BundlingPUSCH_RepTypeBPerBC_r17 != nil {
		if err = ie.dmrs_BundlingPUSCH_RepTypeBPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_BundlingPUSCH_RepTypeBPerBC_r17", err)
		}
	}
	if ie.dmrs_BundlingPUSCH_multiSlotPerBC_r17 != nil {
		if err = ie.dmrs_BundlingPUSCH_multiSlotPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_BundlingPUSCH_multiSlotPerBC_r17", err)
		}
	}
	if ie.dmrs_BundlingPUCCH_RepPerBC_r17 != nil {
		if err = ie.dmrs_BundlingPUCCH_RepPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_BundlingPUCCH_RepPerBC_r17", err)
		}
	}
	if ie.dmrs_BundlingRestartPerBC_r17 != nil {
		if err = ie.dmrs_BundlingRestartPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_BundlingRestartPerBC_r17", err)
		}
	}
	if ie.dmrs_BundlingNonBackToBackTX_PerBC_r17 != nil {
		if err = ie.dmrs_BundlingNonBackToBackTX_PerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_BundlingNonBackToBackTX_PerBC_r17", err)
		}
	}
	if ie.stayOnTargetCC_SRS_CarrierSwitch_r17 != nil {
		if err = ie.stayOnTargetCC_SRS_CarrierSwitch_r17.Encode(w); err != nil {
			return utils.WrapError("Encode stayOnTargetCC_SRS_CarrierSwitch_r17", err)
		}
	}
	if ie.fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil {
		if err = ie.fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if ie.mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 != nil {
		if err = ie.mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if ie.mode1_ForType1_CodebookGeneration_r17 != nil {
		if err = ie.mode1_ForType1_CodebookGeneration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mode1_ForType1_CodebookGeneration_r17", err)
		}
	}
	if ie.nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 != nil {
		if err = ie.nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17", err)
		}
	}
	if ie.multiPUCCH_ConfigForMulticast_r17 != nil {
		if err = ie.multiPUCCH_ConfigForMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode multiPUCCH_ConfigForMulticast_r17", err)
		}
	}
	if ie.pucch_ConfigForSPS_Multicast_r17 != nil {
		if err = ie.pucch_ConfigForSPS_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_ConfigForSPS_Multicast_r17", err)
		}
	}
	if ie.maxNumberG_RNTI_HARQ_ACK_Codebook_r17 != nil {
		if err = w.WriteInteger(*ie.maxNumberG_RNTI_HARQ_ACK_Codebook_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode maxNumberG_RNTI_HARQ_ACK_Codebook_r17", err)
		}
	}
	if ie.mux_HARQ_ACK_UnicastMulticast_r17 != nil {
		if err = ie.mux_HARQ_ACK_UnicastMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mux_HARQ_ACK_UnicastMulticast_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1730) Decode(r *uper.UperReader) error {
	var err error
	var dmrs_BundlingPUSCH_RepTypeAPerBC_r17Present bool
	if dmrs_BundlingPUSCH_RepTypeAPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_BundlingPUSCH_RepTypeBPerBC_r17Present bool
	if dmrs_BundlingPUSCH_RepTypeBPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_BundlingPUSCH_multiSlotPerBC_r17Present bool
	if dmrs_BundlingPUSCH_multiSlotPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_BundlingPUCCH_RepPerBC_r17Present bool
	if dmrs_BundlingPUCCH_RepPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_BundlingRestartPerBC_r17Present bool
	if dmrs_BundlingRestartPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_BundlingNonBackToBackTX_PerBC_r17Present bool
	if dmrs_BundlingNonBackToBackTX_PerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var stayOnTargetCC_SRS_CarrierSwitch_r17Present bool
	if stayOnTargetCC_SRS_CarrierSwitch_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present bool
	if fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present bool
	if mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mode1_ForType1_CodebookGeneration_r17Present bool
	if mode1_ForType1_CodebookGeneration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17Present bool
	if nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiPUCCH_ConfigForMulticast_r17Present bool
	if multiPUCCH_ConfigForMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_ConfigForSPS_Multicast_r17Present bool
	if pucch_ConfigForSPS_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberG_RNTI_HARQ_ACK_Codebook_r17Present bool
	if maxNumberG_RNTI_HARQ_ACK_Codebook_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_HARQ_ACK_UnicastMulticast_r17Present bool
	if mux_HARQ_ACK_UnicastMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dmrs_BundlingPUSCH_RepTypeAPerBC_r17Present {
		ie.dmrs_BundlingPUSCH_RepTypeAPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeAPerBC_r17)
		if err = ie.dmrs_BundlingPUSCH_RepTypeAPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_BundlingPUSCH_RepTypeAPerBC_r17", err)
		}
	}
	if dmrs_BundlingPUSCH_RepTypeBPerBC_r17Present {
		ie.dmrs_BundlingPUSCH_RepTypeBPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUSCH_RepTypeBPerBC_r17)
		if err = ie.dmrs_BundlingPUSCH_RepTypeBPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_BundlingPUSCH_RepTypeBPerBC_r17", err)
		}
	}
	if dmrs_BundlingPUSCH_multiSlotPerBC_r17Present {
		ie.dmrs_BundlingPUSCH_multiSlotPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUSCH_multiSlotPerBC_r17)
		if err = ie.dmrs_BundlingPUSCH_multiSlotPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_BundlingPUSCH_multiSlotPerBC_r17", err)
		}
	}
	if dmrs_BundlingPUCCH_RepPerBC_r17Present {
		ie.dmrs_BundlingPUCCH_RepPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingPUCCH_RepPerBC_r17)
		if err = ie.dmrs_BundlingPUCCH_RepPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_BundlingPUCCH_RepPerBC_r17", err)
		}
	}
	if dmrs_BundlingRestartPerBC_r17Present {
		ie.dmrs_BundlingRestartPerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingRestartPerBC_r17)
		if err = ie.dmrs_BundlingRestartPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_BundlingRestartPerBC_r17", err)
		}
	}
	if dmrs_BundlingNonBackToBackTX_PerBC_r17Present {
		ie.dmrs_BundlingNonBackToBackTX_PerBC_r17 = new(CA_ParametersNR_v1730_dmrs_BundlingNonBackToBackTX_PerBC_r17)
		if err = ie.dmrs_BundlingNonBackToBackTX_PerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_BundlingNonBackToBackTX_PerBC_r17", err)
		}
	}
	if stayOnTargetCC_SRS_CarrierSwitch_r17Present {
		ie.stayOnTargetCC_SRS_CarrierSwitch_r17 = new(CA_ParametersNR_v1730_stayOnTargetCC_SRS_CarrierSwitch_r17)
		if err = ie.stayOnTargetCC_SRS_CarrierSwitch_r17.Decode(r); err != nil {
			return utils.WrapError("Decode stayOnTargetCC_SRS_CarrierSwitch_r17", err)
		}
	}
	if fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present {
		ie.fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17 = new(CA_ParametersNR_v1730_fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17)
		if err = ie.fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fdm_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17Present {
		ie.mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17 = new(CA_ParametersNR_v1730_mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17)
		if err = ie.mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mode2_TDM_CodebookForMux_UnicastMulticastHARQ_ACK_r17", err)
		}
	}
	if mode1_ForType1_CodebookGeneration_r17Present {
		ie.mode1_ForType1_CodebookGeneration_r17 = new(CA_ParametersNR_v1730_mode1_ForType1_CodebookGeneration_r17)
		if err = ie.mode1_ForType1_CodebookGeneration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mode1_ForType1_CodebookGeneration_r17", err)
		}
	}
	if nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17Present {
		ie.nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17 = new(CA_ParametersNR_v1730_nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17)
		if err = ie.nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nack_OnlyFeedbackSpecificResourceForSPS_Multicast_r17", err)
		}
	}
	if multiPUCCH_ConfigForMulticast_r17Present {
		ie.multiPUCCH_ConfigForMulticast_r17 = new(CA_ParametersNR_v1730_multiPUCCH_ConfigForMulticast_r17)
		if err = ie.multiPUCCH_ConfigForMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode multiPUCCH_ConfigForMulticast_r17", err)
		}
	}
	if pucch_ConfigForSPS_Multicast_r17Present {
		ie.pucch_ConfigForSPS_Multicast_r17 = new(CA_ParametersNR_v1730_pucch_ConfigForSPS_Multicast_r17)
		if err = ie.pucch_ConfigForSPS_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_ConfigForSPS_Multicast_r17", err)
		}
	}
	if maxNumberG_RNTI_HARQ_ACK_Codebook_r17Present {
		var tmp_int_maxNumberG_RNTI_HARQ_ACK_Codebook_r17 int64
		if tmp_int_maxNumberG_RNTI_HARQ_ACK_Codebook_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode maxNumberG_RNTI_HARQ_ACK_Codebook_r17", err)
		}
		ie.maxNumberG_RNTI_HARQ_ACK_Codebook_r17 = &tmp_int_maxNumberG_RNTI_HARQ_ACK_Codebook_r17
	}
	if mux_HARQ_ACK_UnicastMulticast_r17Present {
		ie.mux_HARQ_ACK_UnicastMulticast_r17 = new(CA_ParametersNR_v1730_mux_HARQ_ACK_UnicastMulticast_r17)
		if err = ie.mux_HARQ_ACK_UnicastMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mux_HARQ_ACK_UnicastMulticast_r17", err)
		}
	}
	return nil
}
