package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610 struct {
	pusch_RepetitionTypeB_r16                        *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16                        `optional`
	ul_CancellationSelfCarrier_r16                   *FeatureSetUplink_v1610_ul_CancellationSelfCarrier_r16                   `optional`
	ul_CancellationCrossCarrier_r16                  *FeatureSetUplink_v1610_ul_CancellationCrossCarrier_r16                  `optional`
	ul_FullPwrMode2_MaxSRS_ResInSet_r16              *FeatureSetUplink_v1610_ul_FullPwrMode2_MaxSRS_ResInSet_r16              `optional`
	cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 *FeatureSetUplink_v1610_cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 `optional`
	cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 *FeatureSetUplink_v1610_cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 `optional`
	supportedSRS_PosResources_r16                    *SRS_AllPosResources_r16                                                 `optional`
	intraFreqDAPS_UL_r16                             *FeatureSetUplink_v1610_intraFreqDAPS_UL_r16                             `optional`
	intraBandFreqSeparationUL_v1620                  *FreqSeparationClassUL_v1620                                             `optional`
	multiPUCCH_r16                                   *FeatureSetUplink_v1610_multiPUCCH_r16                                   `optional`
	twoPUCCH_Type1_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type1_r16                               `optional`
	twoPUCCH_Type2_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type2_r16                               `optional`
	twoPUCCH_Type3_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type3_r16                               `optional`
	twoPUCCH_Type4_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type4_r16                               `optional`
	mux_SR_HARQ_ACK_r16                              *FeatureSetUplink_v1610_mux_SR_HARQ_ACK_r16                              `optional`
	dummy1                                           *FeatureSetUplink_v1610_dummy1                                           `optional`
	dummy2                                           *FeatureSetUplink_v1610_dummy2                                           `optional`
	twoPUCCH_Type5_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type5_r16                               `optional`
	twoPUCCH_Type6_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type6_r16                               `optional`
	twoPUCCH_Type7_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type7_r16                               `optional`
	twoPUCCH_Type8_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type8_r16                               `optional`
	twoPUCCH_Type9_r16                               *FeatureSetUplink_v1610_twoPUCCH_Type9_r16                               `optional`
	twoPUCCH_Type10_r16                              *FeatureSetUplink_v1610_twoPUCCH_Type10_r16                              `optional`
	twoPUCCH_Type11_r16                              *FeatureSetUplink_v1610_twoPUCCH_Type11_r16                              `optional`
	ul_IntraUE_Mux_r16                               *FeatureSetUplink_v1610_ul_IntraUE_Mux_r16                               `optional`
	ul_FullPwrMode_r16                               *FeatureSetUplink_v1610_ul_FullPwrMode_r16                               `optional`
	crossCarrierSchedulingProcessing_DiffSCS_r16     *CrossCarrierSchedulingProcessing_DiffSCS_r16                            `optional`
	ul_FullPwrMode1_r16                              *FeatureSetUplink_v1610_ul_FullPwrMode1_r16                              `optional`
	ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16    *FeatureSetUplink_v1610_ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16    `optional`
	ul_FullPwrMode2_TPMIGroup_r16                    *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16                    `lb:2,ub:2,optional`
}

func (ie *FeatureSetUplink_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pusch_RepetitionTypeB_r16 != nil, ie.ul_CancellationSelfCarrier_r16 != nil, ie.ul_CancellationCrossCarrier_r16 != nil, ie.ul_FullPwrMode2_MaxSRS_ResInSet_r16 != nil, ie.cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil, ie.cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil, ie.supportedSRS_PosResources_r16 != nil, ie.intraFreqDAPS_UL_r16 != nil, ie.intraBandFreqSeparationUL_v1620 != nil, ie.multiPUCCH_r16 != nil, ie.twoPUCCH_Type1_r16 != nil, ie.twoPUCCH_Type2_r16 != nil, ie.twoPUCCH_Type3_r16 != nil, ie.twoPUCCH_Type4_r16 != nil, ie.mux_SR_HARQ_ACK_r16 != nil, ie.dummy1 != nil, ie.dummy2 != nil, ie.twoPUCCH_Type5_r16 != nil, ie.twoPUCCH_Type6_r16 != nil, ie.twoPUCCH_Type7_r16 != nil, ie.twoPUCCH_Type8_r16 != nil, ie.twoPUCCH_Type9_r16 != nil, ie.twoPUCCH_Type10_r16 != nil, ie.twoPUCCH_Type11_r16 != nil, ie.ul_IntraUE_Mux_r16 != nil, ie.ul_FullPwrMode_r16 != nil, ie.crossCarrierSchedulingProcessing_DiffSCS_r16 != nil, ie.ul_FullPwrMode1_r16 != nil, ie.ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16 != nil, ie.ul_FullPwrMode2_TPMIGroup_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pusch_RepetitionTypeB_r16 != nil {
		if err = ie.pusch_RepetitionTypeB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_RepetitionTypeB_r16", err)
		}
	}
	if ie.ul_CancellationSelfCarrier_r16 != nil {
		if err = ie.ul_CancellationSelfCarrier_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_CancellationSelfCarrier_r16", err)
		}
	}
	if ie.ul_CancellationCrossCarrier_r16 != nil {
		if err = ie.ul_CancellationCrossCarrier_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_CancellationCrossCarrier_r16", err)
		}
	}
	if ie.ul_FullPwrMode2_MaxSRS_ResInSet_r16 != nil {
		if err = ie.ul_FullPwrMode2_MaxSRS_ResInSet_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FullPwrMode2_MaxSRS_ResInSet_r16", err)
		}
	}
	if ie.cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil {
		if err = ie.cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil {
		if err = ie.cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if ie.supportedSRS_PosResources_r16 != nil {
		if err = ie.supportedSRS_PosResources_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedSRS_PosResources_r16", err)
		}
	}
	if ie.intraFreqDAPS_UL_r16 != nil {
		if err = ie.intraFreqDAPS_UL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqDAPS_UL_r16", err)
		}
	}
	if ie.intraBandFreqSeparationUL_v1620 != nil {
		if err = ie.intraBandFreqSeparationUL_v1620.Encode(w); err != nil {
			return utils.WrapError("Encode intraBandFreqSeparationUL_v1620", err)
		}
	}
	if ie.multiPUCCH_r16 != nil {
		if err = ie.multiPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode multiPUCCH_r16", err)
		}
	}
	if ie.twoPUCCH_Type1_r16 != nil {
		if err = ie.twoPUCCH_Type1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type1_r16", err)
		}
	}
	if ie.twoPUCCH_Type2_r16 != nil {
		if err = ie.twoPUCCH_Type2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type2_r16", err)
		}
	}
	if ie.twoPUCCH_Type3_r16 != nil {
		if err = ie.twoPUCCH_Type3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type3_r16", err)
		}
	}
	if ie.twoPUCCH_Type4_r16 != nil {
		if err = ie.twoPUCCH_Type4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type4_r16", err)
		}
	}
	if ie.mux_SR_HARQ_ACK_r16 != nil {
		if err = ie.mux_SR_HARQ_ACK_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mux_SR_HARQ_ACK_r16", err)
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.dummy2 != nil {
		if err = ie.dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	if ie.twoPUCCH_Type5_r16 != nil {
		if err = ie.twoPUCCH_Type5_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type5_r16", err)
		}
	}
	if ie.twoPUCCH_Type6_r16 != nil {
		if err = ie.twoPUCCH_Type6_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type6_r16", err)
		}
	}
	if ie.twoPUCCH_Type7_r16 != nil {
		if err = ie.twoPUCCH_Type7_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type7_r16", err)
		}
	}
	if ie.twoPUCCH_Type8_r16 != nil {
		if err = ie.twoPUCCH_Type8_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type8_r16", err)
		}
	}
	if ie.twoPUCCH_Type9_r16 != nil {
		if err = ie.twoPUCCH_Type9_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type9_r16", err)
		}
	}
	if ie.twoPUCCH_Type10_r16 != nil {
		if err = ie.twoPUCCH_Type10_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type10_r16", err)
		}
	}
	if ie.twoPUCCH_Type11_r16 != nil {
		if err = ie.twoPUCCH_Type11_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Type11_r16", err)
		}
	}
	if ie.ul_IntraUE_Mux_r16 != nil {
		if err = ie.ul_IntraUE_Mux_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_IntraUE_Mux_r16", err)
		}
	}
	if ie.ul_FullPwrMode_r16 != nil {
		if err = ie.ul_FullPwrMode_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FullPwrMode_r16", err)
		}
	}
	if ie.crossCarrierSchedulingProcessing_DiffSCS_r16 != nil {
		if err = ie.crossCarrierSchedulingProcessing_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if ie.ul_FullPwrMode1_r16 != nil {
		if err = ie.ul_FullPwrMode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FullPwrMode1_r16", err)
		}
	}
	if ie.ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16 != nil {
		if err = ie.ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16", err)
		}
	}
	if ie.ul_FullPwrMode2_TPMIGroup_r16 != nil {
		if err = ie.ul_FullPwrMode2_TPMIGroup_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FullPwrMode2_TPMIGroup_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1610) Decode(r *uper.UperReader) error {
	var err error
	var pusch_RepetitionTypeB_r16Present bool
	if pusch_RepetitionTypeB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_CancellationSelfCarrier_r16Present bool
	if ul_CancellationSelfCarrier_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_CancellationCrossCarrier_r16Present bool
	if ul_CancellationCrossCarrier_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FullPwrMode2_MaxSRS_ResInSet_r16Present bool
	if ul_FullPwrMode2_MaxSRS_ResInSet_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16Present bool
	if cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16Present bool
	if cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedSRS_PosResources_r16Present bool
	if supportedSRS_PosResources_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqDAPS_UL_r16Present bool
	if intraFreqDAPS_UL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraBandFreqSeparationUL_v1620Present bool
	if intraBandFreqSeparationUL_v1620Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiPUCCH_r16Present bool
	if multiPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type1_r16Present bool
	if twoPUCCH_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type2_r16Present bool
	if twoPUCCH_Type2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type3_r16Present bool
	if twoPUCCH_Type3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type4_r16Present bool
	if twoPUCCH_Type4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_SR_HARQ_ACK_r16Present bool
	if mux_SR_HARQ_ACK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type5_r16Present bool
	if twoPUCCH_Type5_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type6_r16Present bool
	if twoPUCCH_Type6_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type7_r16Present bool
	if twoPUCCH_Type7_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type8_r16Present bool
	if twoPUCCH_Type8_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type9_r16Present bool
	if twoPUCCH_Type9_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type10_r16Present bool
	if twoPUCCH_Type10_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_Type11_r16Present bool
	if twoPUCCH_Type11_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_IntraUE_Mux_r16Present bool
	if ul_IntraUE_Mux_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FullPwrMode_r16Present bool
	if ul_FullPwrMode_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingProcessing_DiffSCS_r16Present bool
	if crossCarrierSchedulingProcessing_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FullPwrMode1_r16Present bool
	if ul_FullPwrMode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16Present bool
	if ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FullPwrMode2_TPMIGroup_r16Present bool
	if ul_FullPwrMode2_TPMIGroup_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pusch_RepetitionTypeB_r16Present {
		ie.pusch_RepetitionTypeB_r16 = new(FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16)
		if err = ie.pusch_RepetitionTypeB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_RepetitionTypeB_r16", err)
		}
	}
	if ul_CancellationSelfCarrier_r16Present {
		ie.ul_CancellationSelfCarrier_r16 = new(FeatureSetUplink_v1610_ul_CancellationSelfCarrier_r16)
		if err = ie.ul_CancellationSelfCarrier_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_CancellationSelfCarrier_r16", err)
		}
	}
	if ul_CancellationCrossCarrier_r16Present {
		ie.ul_CancellationCrossCarrier_r16 = new(FeatureSetUplink_v1610_ul_CancellationCrossCarrier_r16)
		if err = ie.ul_CancellationCrossCarrier_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_CancellationCrossCarrier_r16", err)
		}
	}
	if ul_FullPwrMode2_MaxSRS_ResInSet_r16Present {
		ie.ul_FullPwrMode2_MaxSRS_ResInSet_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_MaxSRS_ResInSet_r16)
		if err = ie.ul_FullPwrMode2_MaxSRS_ResInSet_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FullPwrMode2_MaxSRS_ResInSet_r16", err)
		}
	}
	if cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16Present {
		ie.cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16 = new(FeatureSetUplink_v1610_cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16)
		if err = ie.cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cbgPUSCH_ProcessingType1_DifferentTB_PerSlot_r16", err)
		}
	}
	if cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16Present {
		ie.cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16 = new(FeatureSetUplink_v1610_cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16)
		if err = ie.cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cbgPUSCH_ProcessingType2_DifferentTB_PerSlot_r16", err)
		}
	}
	if supportedSRS_PosResources_r16Present {
		ie.supportedSRS_PosResources_r16 = new(SRS_AllPosResources_r16)
		if err = ie.supportedSRS_PosResources_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportedSRS_PosResources_r16", err)
		}
	}
	if intraFreqDAPS_UL_r16Present {
		ie.intraFreqDAPS_UL_r16 = new(FeatureSetUplink_v1610_intraFreqDAPS_UL_r16)
		if err = ie.intraFreqDAPS_UL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqDAPS_UL_r16", err)
		}
	}
	if intraBandFreqSeparationUL_v1620Present {
		ie.intraBandFreqSeparationUL_v1620 = new(FreqSeparationClassUL_v1620)
		if err = ie.intraBandFreqSeparationUL_v1620.Decode(r); err != nil {
			return utils.WrapError("Decode intraBandFreqSeparationUL_v1620", err)
		}
	}
	if multiPUCCH_r16Present {
		ie.multiPUCCH_r16 = new(FeatureSetUplink_v1610_multiPUCCH_r16)
		if err = ie.multiPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode multiPUCCH_r16", err)
		}
	}
	if twoPUCCH_Type1_r16Present {
		ie.twoPUCCH_Type1_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type1_r16)
		if err = ie.twoPUCCH_Type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type1_r16", err)
		}
	}
	if twoPUCCH_Type2_r16Present {
		ie.twoPUCCH_Type2_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type2_r16)
		if err = ie.twoPUCCH_Type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type2_r16", err)
		}
	}
	if twoPUCCH_Type3_r16Present {
		ie.twoPUCCH_Type3_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type3_r16)
		if err = ie.twoPUCCH_Type3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type3_r16", err)
		}
	}
	if twoPUCCH_Type4_r16Present {
		ie.twoPUCCH_Type4_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type4_r16)
		if err = ie.twoPUCCH_Type4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type4_r16", err)
		}
	}
	if mux_SR_HARQ_ACK_r16Present {
		ie.mux_SR_HARQ_ACK_r16 = new(FeatureSetUplink_v1610_mux_SR_HARQ_ACK_r16)
		if err = ie.mux_SR_HARQ_ACK_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mux_SR_HARQ_ACK_r16", err)
		}
	}
	if dummy1Present {
		ie.dummy1 = new(FeatureSetUplink_v1610_dummy1)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if dummy2Present {
		ie.dummy2 = new(FeatureSetUplink_v1610_dummy2)
		if err = ie.dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
	}
	if twoPUCCH_Type5_r16Present {
		ie.twoPUCCH_Type5_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type5_r16)
		if err = ie.twoPUCCH_Type5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type5_r16", err)
		}
	}
	if twoPUCCH_Type6_r16Present {
		ie.twoPUCCH_Type6_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type6_r16)
		if err = ie.twoPUCCH_Type6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type6_r16", err)
		}
	}
	if twoPUCCH_Type7_r16Present {
		ie.twoPUCCH_Type7_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type7_r16)
		if err = ie.twoPUCCH_Type7_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type7_r16", err)
		}
	}
	if twoPUCCH_Type8_r16Present {
		ie.twoPUCCH_Type8_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type8_r16)
		if err = ie.twoPUCCH_Type8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type8_r16", err)
		}
	}
	if twoPUCCH_Type9_r16Present {
		ie.twoPUCCH_Type9_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type9_r16)
		if err = ie.twoPUCCH_Type9_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type9_r16", err)
		}
	}
	if twoPUCCH_Type10_r16Present {
		ie.twoPUCCH_Type10_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type10_r16)
		if err = ie.twoPUCCH_Type10_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type10_r16", err)
		}
	}
	if twoPUCCH_Type11_r16Present {
		ie.twoPUCCH_Type11_r16 = new(FeatureSetUplink_v1610_twoPUCCH_Type11_r16)
		if err = ie.twoPUCCH_Type11_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Type11_r16", err)
		}
	}
	if ul_IntraUE_Mux_r16Present {
		ie.ul_IntraUE_Mux_r16 = new(FeatureSetUplink_v1610_ul_IntraUE_Mux_r16)
		if err = ie.ul_IntraUE_Mux_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_IntraUE_Mux_r16", err)
		}
	}
	if ul_FullPwrMode_r16Present {
		ie.ul_FullPwrMode_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode_r16)
		if err = ie.ul_FullPwrMode_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FullPwrMode_r16", err)
		}
	}
	if crossCarrierSchedulingProcessing_DiffSCS_r16Present {
		ie.crossCarrierSchedulingProcessing_DiffSCS_r16 = new(CrossCarrierSchedulingProcessing_DiffSCS_r16)
		if err = ie.crossCarrierSchedulingProcessing_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingProcessing_DiffSCS_r16", err)
		}
	}
	if ul_FullPwrMode1_r16Present {
		ie.ul_FullPwrMode1_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode1_r16)
		if err = ie.ul_FullPwrMode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FullPwrMode1_r16", err)
		}
	}
	if ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16Present {
		ie.ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16)
		if err = ie.ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FullPwrMode2_SRSConfig_diffNumSRSPorts_r16", err)
		}
	}
	if ul_FullPwrMode2_TPMIGroup_r16Present {
		ie.ul_FullPwrMode2_TPMIGroup_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16)
		if err = ie.ul_FullPwrMode2_TPMIGroup_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FullPwrMode2_TPMIGroup_r16", err)
		}
	}
	return nil
}
