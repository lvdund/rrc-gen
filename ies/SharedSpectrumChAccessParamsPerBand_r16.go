package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_r16 struct {
	ul_DynamicChAccess_r16                   *SharedSpectrumChAccessParamsPerBand_r16_ul_DynamicChAccess_r16                   `optional`
	ul_Semi_StaticChAccess_r16               *SharedSpectrumChAccessParamsPerBand_r16_ul_Semi_StaticChAccess_r16               `optional`
	ssb_RRM_DynamicChAccess_r16              *SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_DynamicChAccess_r16              `optional`
	ssb_RRM_Semi_StaticChAccess_r16          *SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_Semi_StaticChAccess_r16          `optional`
	mib_Acquisition_r16                      *SharedSpectrumChAccessParamsPerBand_r16_mib_Acquisition_r16                      `optional`
	ssb_RLM_DynamicChAccess_r16              *SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_DynamicChAccess_r16              `optional`
	ssb_RLM_Semi_StaticChAccess_r16          *SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_Semi_StaticChAccess_r16          `optional`
	sib1_Acquisition_r16                     *SharedSpectrumChAccessParamsPerBand_r16_sib1_Acquisition_r16                     `optional`
	extRA_ResponseWindow_r16                 *SharedSpectrumChAccessParamsPerBand_r16_extRA_ResponseWindow_r16                 `optional`
	ssb_BFD_CBD_dynamicChannelAccess_r16     *SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_dynamicChannelAccess_r16     `optional`
	ssb_BFD_CBD_semi_staticChannelAccess_r16 *SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_semi_staticChannelAccess_r16 `optional`
	csi_RS_BFD_CBD_r16                       *SharedSpectrumChAccessParamsPerBand_r16_csi_RS_BFD_CBD_r16                       `optional`
	ul_ChannelBW_SCell_10mhz_r16             *SharedSpectrumChAccessParamsPerBand_r16_ul_ChannelBW_SCell_10mhz_r16             `optional`
	rssi_ChannelOccupancyReporting_r16       *SharedSpectrumChAccessParamsPerBand_r16_rssi_ChannelOccupancyReporting_r16       `optional`
	srs_StartAnyOFDM_Symbol_r16              *SharedSpectrumChAccessParamsPerBand_r16_srs_StartAnyOFDM_Symbol_r16              `optional`
	searchSpaceFreqMonitorLocation_r16       *int64                                                                            `lb:1,ub:5,optional`
	coreset_RB_Offset_r16                    *SharedSpectrumChAccessParamsPerBand_r16_coreset_RB_Offset_r16                    `optional`
	cgi_Acquisition_r16                      *SharedSpectrumChAccessParamsPerBand_r16_cgi_Acquisition_r16                      `optional`
	configuredUL_Tx_r16                      *SharedSpectrumChAccessParamsPerBand_r16_configuredUL_Tx_r16                      `optional`
	prach_Wideband_r16                       *SharedSpectrumChAccessParamsPerBand_r16_prach_Wideband_r16                       `optional`
	dci_AvailableRB_Set_r16                  *SharedSpectrumChAccessParamsPerBand_r16_dci_AvailableRB_Set_r16                  `optional`
	dci_ChOccupancyDuration_r16              *SharedSpectrumChAccessParamsPerBand_r16_dci_ChOccupancyDuration_r16              `optional`
	typeB_PDSCH_length_r16                   *SharedSpectrumChAccessParamsPerBand_r16_typeB_PDSCH_length_r16                   `optional`
	searchSpaceSwitchWithDCI_r16             *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16             `optional`
	searchSpaceSwitchWithoutDCI_r16          *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithoutDCI_r16          `optional`
	searchSpaceSwitchCapability2_r16         *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchCapability2_r16         `optional`
	non_numericalPDSCH_HARQ_timing_r16       *SharedSpectrumChAccessParamsPerBand_r16_non_numericalPDSCH_HARQ_timing_r16       `optional`
	enhancedDynamicHARQ_codebook_r16         *SharedSpectrumChAccessParamsPerBand_r16_enhancedDynamicHARQ_codebook_r16         `optional`
	oneShotHARQ_feedback_r16                 *SharedSpectrumChAccessParamsPerBand_r16_oneShotHARQ_feedback_r16                 `optional`
	multiPUSCH_UL_grant_r16                  *SharedSpectrumChAccessParamsPerBand_r16_multiPUSCH_UL_grant_r16                  `optional`
	csi_RS_RLM_r16                           *SharedSpectrumChAccessParamsPerBand_r16_csi_RS_RLM_r16                           `optional`
	dummy                                    *SharedSpectrumChAccessParamsPerBand_r16_dummy                                    `optional`
	periodicAndSemi_PersistentCSI_RS_r16     *SharedSpectrumChAccessParamsPerBand_r16_periodicAndSemi_PersistentCSI_RS_r16     `optional`
	pusch_PRB_interlace_r16                  *SharedSpectrumChAccessParamsPerBand_r16_pusch_PRB_interlace_r16                  `optional`
	pucch_F0_F1_PRB_Interlace_r16            *SharedSpectrumChAccessParamsPerBand_r16_pucch_F0_F1_PRB_Interlace_r16            `optional`
	occ_PRB_PF2_PF3_r16                      *SharedSpectrumChAccessParamsPerBand_r16_occ_PRB_PF2_PF3_r16                      `optional`
	extCP_rangeCG_PUSCH_r16                  *SharedSpectrumChAccessParamsPerBand_r16_extCP_rangeCG_PUSCH_r16                  `optional`
	configuredGrantWithReTx_r16              *SharedSpectrumChAccessParamsPerBand_r16_configuredGrantWithReTx_r16              `optional`
	ed_Threshold_r16                         *SharedSpectrumChAccessParamsPerBand_r16_ed_Threshold_r16                         `optional`
	ul_DL_COT_Sharing_r16                    *SharedSpectrumChAccessParamsPerBand_r16_ul_DL_COT_Sharing_r16                    `optional`
	mux_CG_UCI_HARQ_ACK_r16                  *SharedSpectrumChAccessParamsPerBand_r16_mux_CG_UCI_HARQ_ACK_r16                  `optional`
	cg_resourceConfig_r16                    *SharedSpectrumChAccessParamsPerBand_r16_cg_resourceConfig_r16                    `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_DynamicChAccess_r16 != nil, ie.ul_Semi_StaticChAccess_r16 != nil, ie.ssb_RRM_DynamicChAccess_r16 != nil, ie.ssb_RRM_Semi_StaticChAccess_r16 != nil, ie.mib_Acquisition_r16 != nil, ie.ssb_RLM_DynamicChAccess_r16 != nil, ie.ssb_RLM_Semi_StaticChAccess_r16 != nil, ie.sib1_Acquisition_r16 != nil, ie.extRA_ResponseWindow_r16 != nil, ie.ssb_BFD_CBD_dynamicChannelAccess_r16 != nil, ie.ssb_BFD_CBD_semi_staticChannelAccess_r16 != nil, ie.csi_RS_BFD_CBD_r16 != nil, ie.ul_ChannelBW_SCell_10mhz_r16 != nil, ie.rssi_ChannelOccupancyReporting_r16 != nil, ie.srs_StartAnyOFDM_Symbol_r16 != nil, ie.searchSpaceFreqMonitorLocation_r16 != nil, ie.coreset_RB_Offset_r16 != nil, ie.cgi_Acquisition_r16 != nil, ie.configuredUL_Tx_r16 != nil, ie.prach_Wideband_r16 != nil, ie.dci_AvailableRB_Set_r16 != nil, ie.dci_ChOccupancyDuration_r16 != nil, ie.typeB_PDSCH_length_r16 != nil, ie.searchSpaceSwitchWithDCI_r16 != nil, ie.searchSpaceSwitchWithoutDCI_r16 != nil, ie.searchSpaceSwitchCapability2_r16 != nil, ie.non_numericalPDSCH_HARQ_timing_r16 != nil, ie.enhancedDynamicHARQ_codebook_r16 != nil, ie.oneShotHARQ_feedback_r16 != nil, ie.multiPUSCH_UL_grant_r16 != nil, ie.csi_RS_RLM_r16 != nil, ie.dummy != nil, ie.periodicAndSemi_PersistentCSI_RS_r16 != nil, ie.pusch_PRB_interlace_r16 != nil, ie.pucch_F0_F1_PRB_Interlace_r16 != nil, ie.occ_PRB_PF2_PF3_r16 != nil, ie.extCP_rangeCG_PUSCH_r16 != nil, ie.configuredGrantWithReTx_r16 != nil, ie.ed_Threshold_r16 != nil, ie.ul_DL_COT_Sharing_r16 != nil, ie.mux_CG_UCI_HARQ_ACK_r16 != nil, ie.cg_resourceConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_DynamicChAccess_r16 != nil {
		if err = ie.ul_DynamicChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_DynamicChAccess_r16", err)
		}
	}
	if ie.ul_Semi_StaticChAccess_r16 != nil {
		if err = ie.ul_Semi_StaticChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_Semi_StaticChAccess_r16", err)
		}
	}
	if ie.ssb_RRM_DynamicChAccess_r16 != nil {
		if err = ie.ssb_RRM_DynamicChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RRM_DynamicChAccess_r16", err)
		}
	}
	if ie.ssb_RRM_Semi_StaticChAccess_r16 != nil {
		if err = ie.ssb_RRM_Semi_StaticChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RRM_Semi_StaticChAccess_r16", err)
		}
	}
	if ie.mib_Acquisition_r16 != nil {
		if err = ie.mib_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mib_Acquisition_r16", err)
		}
	}
	if ie.ssb_RLM_DynamicChAccess_r16 != nil {
		if err = ie.ssb_RLM_DynamicChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RLM_DynamicChAccess_r16", err)
		}
	}
	if ie.ssb_RLM_Semi_StaticChAccess_r16 != nil {
		if err = ie.ssb_RLM_Semi_StaticChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RLM_Semi_StaticChAccess_r16", err)
		}
	}
	if ie.sib1_Acquisition_r16 != nil {
		if err = ie.sib1_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sib1_Acquisition_r16", err)
		}
	}
	if ie.extRA_ResponseWindow_r16 != nil {
		if err = ie.extRA_ResponseWindow_r16.Encode(w); err != nil {
			return utils.WrapError("Encode extRA_ResponseWindow_r16", err)
		}
	}
	if ie.ssb_BFD_CBD_dynamicChannelAccess_r16 != nil {
		if err = ie.ssb_BFD_CBD_dynamicChannelAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_BFD_CBD_dynamicChannelAccess_r16", err)
		}
	}
	if ie.ssb_BFD_CBD_semi_staticChannelAccess_r16 != nil {
		if err = ie.ssb_BFD_CBD_semi_staticChannelAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_BFD_CBD_semi_staticChannelAccess_r16", err)
		}
	}
	if ie.csi_RS_BFD_CBD_r16 != nil {
		if err = ie.csi_RS_BFD_CBD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_BFD_CBD_r16", err)
		}
	}
	if ie.ul_ChannelBW_SCell_10mhz_r16 != nil {
		if err = ie.ul_ChannelBW_SCell_10mhz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_ChannelBW_SCell_10mhz_r16", err)
		}
	}
	if ie.rssi_ChannelOccupancyReporting_r16 != nil {
		if err = ie.rssi_ChannelOccupancyReporting_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rssi_ChannelOccupancyReporting_r16", err)
		}
	}
	if ie.srs_StartAnyOFDM_Symbol_r16 != nil {
		if err = ie.srs_StartAnyOFDM_Symbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode srs_StartAnyOFDM_Symbol_r16", err)
		}
	}
	if ie.searchSpaceFreqMonitorLocation_r16 != nil {
		if err = w.WriteInteger(*ie.searchSpaceFreqMonitorLocation_r16, &uper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
			return utils.WrapError("Encode searchSpaceFreqMonitorLocation_r16", err)
		}
	}
	if ie.coreset_RB_Offset_r16 != nil {
		if err = ie.coreset_RB_Offset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode coreset_RB_Offset_r16", err)
		}
	}
	if ie.cgi_Acquisition_r16 != nil {
		if err = ie.cgi_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cgi_Acquisition_r16", err)
		}
	}
	if ie.configuredUL_Tx_r16 != nil {
		if err = ie.configuredUL_Tx_r16.Encode(w); err != nil {
			return utils.WrapError("Encode configuredUL_Tx_r16", err)
		}
	}
	if ie.prach_Wideband_r16 != nil {
		if err = ie.prach_Wideband_r16.Encode(w); err != nil {
			return utils.WrapError("Encode prach_Wideband_r16", err)
		}
	}
	if ie.dci_AvailableRB_Set_r16 != nil {
		if err = ie.dci_AvailableRB_Set_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_AvailableRB_Set_r16", err)
		}
	}
	if ie.dci_ChOccupancyDuration_r16 != nil {
		if err = ie.dci_ChOccupancyDuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_ChOccupancyDuration_r16", err)
		}
	}
	if ie.typeB_PDSCH_length_r16 != nil {
		if err = ie.typeB_PDSCH_length_r16.Encode(w); err != nil {
			return utils.WrapError("Encode typeB_PDSCH_length_r16", err)
		}
	}
	if ie.searchSpaceSwitchWithDCI_r16 != nil {
		if err = ie.searchSpaceSwitchWithDCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSwitchWithDCI_r16", err)
		}
	}
	if ie.searchSpaceSwitchWithoutDCI_r16 != nil {
		if err = ie.searchSpaceSwitchWithoutDCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSwitchWithoutDCI_r16", err)
		}
	}
	if ie.searchSpaceSwitchCapability2_r16 != nil {
		if err = ie.searchSpaceSwitchCapability2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSwitchCapability2_r16", err)
		}
	}
	if ie.non_numericalPDSCH_HARQ_timing_r16 != nil {
		if err = ie.non_numericalPDSCH_HARQ_timing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode non_numericalPDSCH_HARQ_timing_r16", err)
		}
	}
	if ie.enhancedDynamicHARQ_codebook_r16 != nil {
		if err = ie.enhancedDynamicHARQ_codebook_r16.Encode(w); err != nil {
			return utils.WrapError("Encode enhancedDynamicHARQ_codebook_r16", err)
		}
	}
	if ie.oneShotHARQ_feedback_r16 != nil {
		if err = ie.oneShotHARQ_feedback_r16.Encode(w); err != nil {
			return utils.WrapError("Encode oneShotHARQ_feedback_r16", err)
		}
	}
	if ie.multiPUSCH_UL_grant_r16 != nil {
		if err = ie.multiPUSCH_UL_grant_r16.Encode(w); err != nil {
			return utils.WrapError("Encode multiPUSCH_UL_grant_r16", err)
		}
	}
	if ie.csi_RS_RLM_r16 != nil {
		if err = ie.csi_RS_RLM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_RLM_r16", err)
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.periodicAndSemi_PersistentCSI_RS_r16 != nil {
		if err = ie.periodicAndSemi_PersistentCSI_RS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode periodicAndSemi_PersistentCSI_RS_r16", err)
		}
	}
	if ie.pusch_PRB_interlace_r16 != nil {
		if err = ie.pusch_PRB_interlace_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_PRB_interlace_r16", err)
		}
	}
	if ie.pucch_F0_F1_PRB_Interlace_r16 != nil {
		if err = ie.pucch_F0_F1_PRB_Interlace_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F0_F1_PRB_Interlace_r16", err)
		}
	}
	if ie.occ_PRB_PF2_PF3_r16 != nil {
		if err = ie.occ_PRB_PF2_PF3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode occ_PRB_PF2_PF3_r16", err)
		}
	}
	if ie.extCP_rangeCG_PUSCH_r16 != nil {
		if err = ie.extCP_rangeCG_PUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode extCP_rangeCG_PUSCH_r16", err)
		}
	}
	if ie.configuredGrantWithReTx_r16 != nil {
		if err = ie.configuredGrantWithReTx_r16.Encode(w); err != nil {
			return utils.WrapError("Encode configuredGrantWithReTx_r16", err)
		}
	}
	if ie.ed_Threshold_r16 != nil {
		if err = ie.ed_Threshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ed_Threshold_r16", err)
		}
	}
	if ie.ul_DL_COT_Sharing_r16 != nil {
		if err = ie.ul_DL_COT_Sharing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_DL_COT_Sharing_r16", err)
		}
	}
	if ie.mux_CG_UCI_HARQ_ACK_r16 != nil {
		if err = ie.mux_CG_UCI_HARQ_ACK_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mux_CG_UCI_HARQ_ACK_r16", err)
		}
	}
	if ie.cg_resourceConfig_r16 != nil {
		if err = ie.cg_resourceConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cg_resourceConfig_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16) Decode(r *uper.UperReader) error {
	var err error
	var ul_DynamicChAccess_r16Present bool
	if ul_DynamicChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_Semi_StaticChAccess_r16Present bool
	if ul_Semi_StaticChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_RRM_DynamicChAccess_r16Present bool
	if ssb_RRM_DynamicChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_RRM_Semi_StaticChAccess_r16Present bool
	if ssb_RRM_Semi_StaticChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mib_Acquisition_r16Present bool
	if mib_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_RLM_DynamicChAccess_r16Present bool
	if ssb_RLM_DynamicChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_RLM_Semi_StaticChAccess_r16Present bool
	if ssb_RLM_Semi_StaticChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sib1_Acquisition_r16Present bool
	if sib1_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var extRA_ResponseWindow_r16Present bool
	if extRA_ResponseWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_BFD_CBD_dynamicChannelAccess_r16Present bool
	if ssb_BFD_CBD_dynamicChannelAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_BFD_CBD_semi_staticChannelAccess_r16Present bool
	if ssb_BFD_CBD_semi_staticChannelAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_BFD_CBD_r16Present bool
	if csi_RS_BFD_CBD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_ChannelBW_SCell_10mhz_r16Present bool
	if ul_ChannelBW_SCell_10mhz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rssi_ChannelOccupancyReporting_r16Present bool
	if rssi_ChannelOccupancyReporting_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_StartAnyOFDM_Symbol_r16Present bool
	if srs_StartAnyOFDM_Symbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceFreqMonitorLocation_r16Present bool
	if searchSpaceFreqMonitorLocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var coreset_RB_Offset_r16Present bool
	if coreset_RB_Offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cgi_Acquisition_r16Present bool
	if cgi_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredUL_Tx_r16Present bool
	if configuredUL_Tx_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var prach_Wideband_r16Present bool
	if prach_Wideband_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_AvailableRB_Set_r16Present bool
	if dci_AvailableRB_Set_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_ChOccupancyDuration_r16Present bool
	if dci_ChOccupancyDuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var typeB_PDSCH_length_r16Present bool
	if typeB_PDSCH_length_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSwitchWithDCI_r16Present bool
	if searchSpaceSwitchWithDCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSwitchWithoutDCI_r16Present bool
	if searchSpaceSwitchWithoutDCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSwitchCapability2_r16Present bool
	if searchSpaceSwitchCapability2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var non_numericalPDSCH_HARQ_timing_r16Present bool
	if non_numericalPDSCH_HARQ_timing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var enhancedDynamicHARQ_codebook_r16Present bool
	if enhancedDynamicHARQ_codebook_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var oneShotHARQ_feedback_r16Present bool
	if oneShotHARQ_feedback_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multiPUSCH_UL_grant_r16Present bool
	if multiPUSCH_UL_grant_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_RLM_r16Present bool
	if csi_RS_RLM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var periodicAndSemi_PersistentCSI_RS_r16Present bool
	if periodicAndSemi_PersistentCSI_RS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_PRB_interlace_r16Present bool
	if pusch_PRB_interlace_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F0_F1_PRB_Interlace_r16Present bool
	if pucch_F0_F1_PRB_Interlace_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var occ_PRB_PF2_PF3_r16Present bool
	if occ_PRB_PF2_PF3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var extCP_rangeCG_PUSCH_r16Present bool
	if extCP_rangeCG_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredGrantWithReTx_r16Present bool
	if configuredGrantWithReTx_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ed_Threshold_r16Present bool
	if ed_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_DL_COT_Sharing_r16Present bool
	if ul_DL_COT_Sharing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_CG_UCI_HARQ_ACK_r16Present bool
	if mux_CG_UCI_HARQ_ACK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_resourceConfig_r16Present bool
	if cg_resourceConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_DynamicChAccess_r16Present {
		ie.ul_DynamicChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_DynamicChAccess_r16)
		if err = ie.ul_DynamicChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_DynamicChAccess_r16", err)
		}
	}
	if ul_Semi_StaticChAccess_r16Present {
		ie.ul_Semi_StaticChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_Semi_StaticChAccess_r16)
		if err = ie.ul_Semi_StaticChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_Semi_StaticChAccess_r16", err)
		}
	}
	if ssb_RRM_DynamicChAccess_r16Present {
		ie.ssb_RRM_DynamicChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_DynamicChAccess_r16)
		if err = ie.ssb_RRM_DynamicChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RRM_DynamicChAccess_r16", err)
		}
	}
	if ssb_RRM_Semi_StaticChAccess_r16Present {
		ie.ssb_RRM_Semi_StaticChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RRM_Semi_StaticChAccess_r16)
		if err = ie.ssb_RRM_Semi_StaticChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RRM_Semi_StaticChAccess_r16", err)
		}
	}
	if mib_Acquisition_r16Present {
		ie.mib_Acquisition_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_mib_Acquisition_r16)
		if err = ie.mib_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mib_Acquisition_r16", err)
		}
	}
	if ssb_RLM_DynamicChAccess_r16Present {
		ie.ssb_RLM_DynamicChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_DynamicChAccess_r16)
		if err = ie.ssb_RLM_DynamicChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RLM_DynamicChAccess_r16", err)
		}
	}
	if ssb_RLM_Semi_StaticChAccess_r16Present {
		ie.ssb_RLM_Semi_StaticChAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_RLM_Semi_StaticChAccess_r16)
		if err = ie.ssb_RLM_Semi_StaticChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RLM_Semi_StaticChAccess_r16", err)
		}
	}
	if sib1_Acquisition_r16Present {
		ie.sib1_Acquisition_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_sib1_Acquisition_r16)
		if err = ie.sib1_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sib1_Acquisition_r16", err)
		}
	}
	if extRA_ResponseWindow_r16Present {
		ie.extRA_ResponseWindow_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_extRA_ResponseWindow_r16)
		if err = ie.extRA_ResponseWindow_r16.Decode(r); err != nil {
			return utils.WrapError("Decode extRA_ResponseWindow_r16", err)
		}
	}
	if ssb_BFD_CBD_dynamicChannelAccess_r16Present {
		ie.ssb_BFD_CBD_dynamicChannelAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_dynamicChannelAccess_r16)
		if err = ie.ssb_BFD_CBD_dynamicChannelAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_BFD_CBD_dynamicChannelAccess_r16", err)
		}
	}
	if ssb_BFD_CBD_semi_staticChannelAccess_r16Present {
		ie.ssb_BFD_CBD_semi_staticChannelAccess_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ssb_BFD_CBD_semi_staticChannelAccess_r16)
		if err = ie.ssb_BFD_CBD_semi_staticChannelAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_BFD_CBD_semi_staticChannelAccess_r16", err)
		}
	}
	if csi_RS_BFD_CBD_r16Present {
		ie.csi_RS_BFD_CBD_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_csi_RS_BFD_CBD_r16)
		if err = ie.csi_RS_BFD_CBD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_BFD_CBD_r16", err)
		}
	}
	if ul_ChannelBW_SCell_10mhz_r16Present {
		ie.ul_ChannelBW_SCell_10mhz_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_ChannelBW_SCell_10mhz_r16)
		if err = ie.ul_ChannelBW_SCell_10mhz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_ChannelBW_SCell_10mhz_r16", err)
		}
	}
	if rssi_ChannelOccupancyReporting_r16Present {
		ie.rssi_ChannelOccupancyReporting_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_rssi_ChannelOccupancyReporting_r16)
		if err = ie.rssi_ChannelOccupancyReporting_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rssi_ChannelOccupancyReporting_r16", err)
		}
	}
	if srs_StartAnyOFDM_Symbol_r16Present {
		ie.srs_StartAnyOFDM_Symbol_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_srs_StartAnyOFDM_Symbol_r16)
		if err = ie.srs_StartAnyOFDM_Symbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode srs_StartAnyOFDM_Symbol_r16", err)
		}
	}
	if searchSpaceFreqMonitorLocation_r16Present {
		var tmp_int_searchSpaceFreqMonitorLocation_r16 int64
		if tmp_int_searchSpaceFreqMonitorLocation_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode searchSpaceFreqMonitorLocation_r16", err)
		}
		ie.searchSpaceFreqMonitorLocation_r16 = &tmp_int_searchSpaceFreqMonitorLocation_r16
	}
	if coreset_RB_Offset_r16Present {
		ie.coreset_RB_Offset_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_coreset_RB_Offset_r16)
		if err = ie.coreset_RB_Offset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode coreset_RB_Offset_r16", err)
		}
	}
	if cgi_Acquisition_r16Present {
		ie.cgi_Acquisition_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_cgi_Acquisition_r16)
		if err = ie.cgi_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cgi_Acquisition_r16", err)
		}
	}
	if configuredUL_Tx_r16Present {
		ie.configuredUL_Tx_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_configuredUL_Tx_r16)
		if err = ie.configuredUL_Tx_r16.Decode(r); err != nil {
			return utils.WrapError("Decode configuredUL_Tx_r16", err)
		}
	}
	if prach_Wideband_r16Present {
		ie.prach_Wideband_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_prach_Wideband_r16)
		if err = ie.prach_Wideband_r16.Decode(r); err != nil {
			return utils.WrapError("Decode prach_Wideband_r16", err)
		}
	}
	if dci_AvailableRB_Set_r16Present {
		ie.dci_AvailableRB_Set_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_dci_AvailableRB_Set_r16)
		if err = ie.dci_AvailableRB_Set_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_AvailableRB_Set_r16", err)
		}
	}
	if dci_ChOccupancyDuration_r16Present {
		ie.dci_ChOccupancyDuration_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_dci_ChOccupancyDuration_r16)
		if err = ie.dci_ChOccupancyDuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_ChOccupancyDuration_r16", err)
		}
	}
	if typeB_PDSCH_length_r16Present {
		ie.typeB_PDSCH_length_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_typeB_PDSCH_length_r16)
		if err = ie.typeB_PDSCH_length_r16.Decode(r); err != nil {
			return utils.WrapError("Decode typeB_PDSCH_length_r16", err)
		}
	}
	if searchSpaceSwitchWithDCI_r16Present {
		ie.searchSpaceSwitchWithDCI_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16)
		if err = ie.searchSpaceSwitchWithDCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSwitchWithDCI_r16", err)
		}
	}
	if searchSpaceSwitchWithoutDCI_r16Present {
		ie.searchSpaceSwitchWithoutDCI_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithoutDCI_r16)
		if err = ie.searchSpaceSwitchWithoutDCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSwitchWithoutDCI_r16", err)
		}
	}
	if searchSpaceSwitchCapability2_r16Present {
		ie.searchSpaceSwitchCapability2_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchCapability2_r16)
		if err = ie.searchSpaceSwitchCapability2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSwitchCapability2_r16", err)
		}
	}
	if non_numericalPDSCH_HARQ_timing_r16Present {
		ie.non_numericalPDSCH_HARQ_timing_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_non_numericalPDSCH_HARQ_timing_r16)
		if err = ie.non_numericalPDSCH_HARQ_timing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode non_numericalPDSCH_HARQ_timing_r16", err)
		}
	}
	if enhancedDynamicHARQ_codebook_r16Present {
		ie.enhancedDynamicHARQ_codebook_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_enhancedDynamicHARQ_codebook_r16)
		if err = ie.enhancedDynamicHARQ_codebook_r16.Decode(r); err != nil {
			return utils.WrapError("Decode enhancedDynamicHARQ_codebook_r16", err)
		}
	}
	if oneShotHARQ_feedback_r16Present {
		ie.oneShotHARQ_feedback_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_oneShotHARQ_feedback_r16)
		if err = ie.oneShotHARQ_feedback_r16.Decode(r); err != nil {
			return utils.WrapError("Decode oneShotHARQ_feedback_r16", err)
		}
	}
	if multiPUSCH_UL_grant_r16Present {
		ie.multiPUSCH_UL_grant_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_multiPUSCH_UL_grant_r16)
		if err = ie.multiPUSCH_UL_grant_r16.Decode(r); err != nil {
			return utils.WrapError("Decode multiPUSCH_UL_grant_r16", err)
		}
	}
	if csi_RS_RLM_r16Present {
		ie.csi_RS_RLM_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_csi_RS_RLM_r16)
		if err = ie.csi_RS_RLM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_RLM_r16", err)
		}
	}
	if dummyPresent {
		ie.dummy = new(SharedSpectrumChAccessParamsPerBand_r16_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if periodicAndSemi_PersistentCSI_RS_r16Present {
		ie.periodicAndSemi_PersistentCSI_RS_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_periodicAndSemi_PersistentCSI_RS_r16)
		if err = ie.periodicAndSemi_PersistentCSI_RS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode periodicAndSemi_PersistentCSI_RS_r16", err)
		}
	}
	if pusch_PRB_interlace_r16Present {
		ie.pusch_PRB_interlace_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_pusch_PRB_interlace_r16)
		if err = ie.pusch_PRB_interlace_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_PRB_interlace_r16", err)
		}
	}
	if pucch_F0_F1_PRB_Interlace_r16Present {
		ie.pucch_F0_F1_PRB_Interlace_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_pucch_F0_F1_PRB_Interlace_r16)
		if err = ie.pucch_F0_F1_PRB_Interlace_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F0_F1_PRB_Interlace_r16", err)
		}
	}
	if occ_PRB_PF2_PF3_r16Present {
		ie.occ_PRB_PF2_PF3_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_occ_PRB_PF2_PF3_r16)
		if err = ie.occ_PRB_PF2_PF3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode occ_PRB_PF2_PF3_r16", err)
		}
	}
	if extCP_rangeCG_PUSCH_r16Present {
		ie.extCP_rangeCG_PUSCH_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_extCP_rangeCG_PUSCH_r16)
		if err = ie.extCP_rangeCG_PUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode extCP_rangeCG_PUSCH_r16", err)
		}
	}
	if configuredGrantWithReTx_r16Present {
		ie.configuredGrantWithReTx_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_configuredGrantWithReTx_r16)
		if err = ie.configuredGrantWithReTx_r16.Decode(r); err != nil {
			return utils.WrapError("Decode configuredGrantWithReTx_r16", err)
		}
	}
	if ed_Threshold_r16Present {
		ie.ed_Threshold_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ed_Threshold_r16)
		if err = ie.ed_Threshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ed_Threshold_r16", err)
		}
	}
	if ul_DL_COT_Sharing_r16Present {
		ie.ul_DL_COT_Sharing_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_ul_DL_COT_Sharing_r16)
		if err = ie.ul_DL_COT_Sharing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_DL_COT_Sharing_r16", err)
		}
	}
	if mux_CG_UCI_HARQ_ACK_r16Present {
		ie.mux_CG_UCI_HARQ_ACK_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_mux_CG_UCI_HARQ_ACK_r16)
		if err = ie.mux_CG_UCI_HARQ_ACK_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mux_CG_UCI_HARQ_ACK_r16", err)
		}
	}
	if cg_resourceConfig_r16Present {
		ie.cg_resourceConfig_r16 = new(SharedSpectrumChAccessParamsPerBand_r16_cg_resourceConfig_r16)
		if err = ie.cg_resourceConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cg_resourceConfig_r16", err)
		}
	}
	return nil
}
