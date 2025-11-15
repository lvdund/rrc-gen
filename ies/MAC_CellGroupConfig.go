package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_CellGroupConfig struct {
	drx_Config                              *DRX_Config                                             `optional,setuprelease`
	schedulingRequestConfig                 *SchedulingRequestConfig                                `optional`
	bsr_Config                              *BSR_Config                                             `optional`
	tag_Config                              *TAG_Config                                             `optional`
	phr_Config                              *PHR_Config                                             `optional,setuprelease`
	skipUplinkTxDynamic                     bool                                                    `madatory`
	csi_Mask                                *bool                                                   `optional,ext-1`
	dataInactivityTimer                     *DataInactivityTimer                                    `optional,ext-1,setuprelease`
	usePreBSR_r16                           *MAC_CellGroupConfig_usePreBSR_r16                      `optional,ext-2`
	schedulingRequestID_LBT_SCell_r16       *SchedulingRequestId                                    `optional,ext-2`
	lch_BasedPrioritization_r16             *MAC_CellGroupConfig_lch_BasedPrioritization_r16        `optional,ext-2`
	schedulingRequestID_BFR_SCell_r16       *SchedulingRequestId                                    `optional,ext-2`
	drx_ConfigSecondaryGroup_r16            *DRX_ConfigSecondaryGroup_r16                           `optional,ext-2,setuprelease`
	enhancedSkipUplinkTxDynamic_r16         *MAC_CellGroupConfig_enhancedSkipUplinkTxDynamic_r16    `optional,ext-3`
	enhancedSkipUplinkTxConfigured_r16      *MAC_CellGroupConfig_enhancedSkipUplinkTxConfigured_r16 `optional,ext-3`
	intraCG_Prioritization_r17              *MAC_CellGroupConfig_intraCG_Prioritization_r17         `optional,ext-4`
	drx_ConfigSL_r17                        *DRX_ConfigSL_r17                                       `optional,ext-4,setuprelease`
	drx_ConfigExt_v1700                     *DRX_ConfigExt_v1700                                    `optional,ext-4,setuprelease`
	schedulingRequestID_BFR_r17             *SchedulingRequestId                                    `optional,ext-4`
	schedulingRequestID_BFR2_r17            *SchedulingRequestId                                    `optional,ext-4`
	schedulingRequestConfig_v1700           *SchedulingRequestConfig_v1700                          `optional,ext-4`
	tar_Config_r17                          *TAR_Config_r17                                         `optional,ext-4,setuprelease`
	g_RNTI_ConfigToAddModList_r17           []MBS_RNTI_SpecificConfig_r17                           `lb:1,ub:maxG_RNTI_r17,optional,ext-4`
	g_RNTI_ConfigToReleaseList_r17          []MBS_RNTI_SpecificConfigId_r17                         `lb:1,ub:maxG_RNTI_r17,optional,ext-4`
	g_CS_RNTI_ConfigToAddModList_r17        []MBS_RNTI_SpecificConfig_r17                           `lb:1,ub:maxG_CS_RNTI_r17,optional,ext-4`
	g_CS_RNTI_ConfigToReleaseList_r17       []MBS_RNTI_SpecificConfigId_r17                         `lb:1,ub:maxG_CS_RNTI_r17,optional,ext-4`
	allowCSI_SRS_Tx_MulticastDRX_Active_r17 *bool                                                   `optional,ext-4`
	schedulingRequestID_PosMG_Request_r17   *SchedulingRequestId                                    `optional,ext-5`
	drx_LastTransmissionUL_r17              *MAC_CellGroupConfig_drx_LastTransmissionUL_r17         `optional,ext-5`
}

func (ie *MAC_CellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.csi_Mask != nil || ie.dataInactivityTimer != nil || ie.usePreBSR_r16 != nil || ie.schedulingRequestID_LBT_SCell_r16 != nil || ie.lch_BasedPrioritization_r16 != nil || ie.schedulingRequestID_BFR_SCell_r16 != nil || ie.drx_ConfigSecondaryGroup_r16 != nil || ie.enhancedSkipUplinkTxDynamic_r16 != nil || ie.enhancedSkipUplinkTxConfigured_r16 != nil || ie.intraCG_Prioritization_r17 != nil || ie.drx_ConfigSL_r17 != nil || ie.drx_ConfigExt_v1700 != nil || ie.schedulingRequestID_BFR_r17 != nil || ie.schedulingRequestID_BFR2_r17 != nil || ie.schedulingRequestConfig_v1700 != nil || ie.tar_Config_r17 != nil || len(ie.g_RNTI_ConfigToAddModList_r17) > 0 || len(ie.g_RNTI_ConfigToReleaseList_r17) > 0 || len(ie.g_CS_RNTI_ConfigToAddModList_r17) > 0 || len(ie.g_CS_RNTI_ConfigToReleaseList_r17) > 0 || ie.allowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil || ie.schedulingRequestID_PosMG_Request_r17 != nil || ie.drx_LastTransmissionUL_r17 != nil
	preambleBits := []bool{hasExtensions, ie.drx_Config != nil, ie.schedulingRequestConfig != nil, ie.bsr_Config != nil, ie.tag_Config != nil, ie.phr_Config != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drx_Config != nil {
		tmp_drx_Config := utils.SetupRelease[*DRX_Config]{
			Setup: ie.drx_Config,
		}
		if err = tmp_drx_Config.Encode(w); err != nil {
			return utils.WrapError("Encode drx_Config", err)
		}
	}
	if ie.schedulingRequestConfig != nil {
		if err = ie.schedulingRequestConfig.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestConfig", err)
		}
	}
	if ie.bsr_Config != nil {
		if err = ie.bsr_Config.Encode(w); err != nil {
			return utils.WrapError("Encode bsr_Config", err)
		}
	}
	if ie.tag_Config != nil {
		if err = ie.tag_Config.Encode(w); err != nil {
			return utils.WrapError("Encode tag_Config", err)
		}
	}
	if ie.phr_Config != nil {
		tmp_phr_Config := utils.SetupRelease[*PHR_Config]{
			Setup: ie.phr_Config,
		}
		if err = tmp_phr_Config.Encode(w); err != nil {
			return utils.WrapError("Encode phr_Config", err)
		}
	}
	if err = w.WriteBoolean(ie.skipUplinkTxDynamic); err != nil {
		return utils.WrapError("WriteBoolean skipUplinkTxDynamic", err)
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.csi_Mask != nil || ie.dataInactivityTimer != nil, ie.usePreBSR_r16 != nil || ie.schedulingRequestID_LBT_SCell_r16 != nil || ie.lch_BasedPrioritization_r16 != nil || ie.schedulingRequestID_BFR_SCell_r16 != nil || ie.drx_ConfigSecondaryGroup_r16 != nil, ie.enhancedSkipUplinkTxDynamic_r16 != nil || ie.enhancedSkipUplinkTxConfigured_r16 != nil, ie.intraCG_Prioritization_r17 != nil || ie.drx_ConfigSL_r17 != nil || ie.drx_ConfigExt_v1700 != nil || ie.schedulingRequestID_BFR_r17 != nil || ie.schedulingRequestID_BFR2_r17 != nil || ie.schedulingRequestConfig_v1700 != nil || ie.tar_Config_r17 != nil || len(ie.g_RNTI_ConfigToAddModList_r17) > 0 || len(ie.g_RNTI_ConfigToReleaseList_r17) > 0 || len(ie.g_CS_RNTI_ConfigToAddModList_r17) > 0 || len(ie.g_CS_RNTI_ConfigToReleaseList_r17) > 0 || ie.allowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil, ie.schedulingRequestID_PosMG_Request_r17 != nil || ie.drx_LastTransmissionUL_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_CellGroupConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.csi_Mask != nil, ie.dataInactivityTimer != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode csi_Mask optional
			if ie.csi_Mask != nil {
				if err = extWriter.WriteBoolean(*ie.csi_Mask); err != nil {
					return utils.WrapError("Encode csi_Mask", err)
				}
			}
			// encode dataInactivityTimer optional
			if ie.dataInactivityTimer != nil {
				tmp_dataInactivityTimer := utils.SetupRelease[*DataInactivityTimer]{
					Setup: ie.dataInactivityTimer,
				}
				if err = tmp_dataInactivityTimer.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dataInactivityTimer", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.usePreBSR_r16 != nil, ie.schedulingRequestID_LBT_SCell_r16 != nil, ie.lch_BasedPrioritization_r16 != nil, ie.schedulingRequestID_BFR_SCell_r16 != nil, ie.drx_ConfigSecondaryGroup_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode usePreBSR_r16 optional
			if ie.usePreBSR_r16 != nil {
				if err = ie.usePreBSR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode usePreBSR_r16", err)
				}
			}
			// encode schedulingRequestID_LBT_SCell_r16 optional
			if ie.schedulingRequestID_LBT_SCell_r16 != nil {
				if err = ie.schedulingRequestID_LBT_SCell_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestID_LBT_SCell_r16", err)
				}
			}
			// encode lch_BasedPrioritization_r16 optional
			if ie.lch_BasedPrioritization_r16 != nil {
				if err = ie.lch_BasedPrioritization_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lch_BasedPrioritization_r16", err)
				}
			}
			// encode schedulingRequestID_BFR_SCell_r16 optional
			if ie.schedulingRequestID_BFR_SCell_r16 != nil {
				if err = ie.schedulingRequestID_BFR_SCell_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestID_BFR_SCell_r16", err)
				}
			}
			// encode drx_ConfigSecondaryGroup_r16 optional
			if ie.drx_ConfigSecondaryGroup_r16 != nil {
				tmp_drx_ConfigSecondaryGroup_r16 := utils.SetupRelease[*DRX_ConfigSecondaryGroup_r16]{
					Setup: ie.drx_ConfigSecondaryGroup_r16,
				}
				if err = tmp_drx_ConfigSecondaryGroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drx_ConfigSecondaryGroup_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.enhancedSkipUplinkTxDynamic_r16 != nil, ie.enhancedSkipUplinkTxConfigured_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enhancedSkipUplinkTxDynamic_r16 optional
			if ie.enhancedSkipUplinkTxDynamic_r16 != nil {
				if err = ie.enhancedSkipUplinkTxDynamic_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// encode enhancedSkipUplinkTxConfigured_r16 optional
			if ie.enhancedSkipUplinkTxConfigured_r16 != nil {
				if err = ie.enhancedSkipUplinkTxConfigured_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedSkipUplinkTxConfigured_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.intraCG_Prioritization_r17 != nil, ie.drx_ConfigSL_r17 != nil, ie.drx_ConfigExt_v1700 != nil, ie.schedulingRequestID_BFR_r17 != nil, ie.schedulingRequestID_BFR2_r17 != nil, ie.schedulingRequestConfig_v1700 != nil, ie.tar_Config_r17 != nil, len(ie.g_RNTI_ConfigToAddModList_r17) > 0, len(ie.g_RNTI_ConfigToReleaseList_r17) > 0, len(ie.g_CS_RNTI_ConfigToAddModList_r17) > 0, len(ie.g_CS_RNTI_ConfigToReleaseList_r17) > 0, ie.allowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode intraCG_Prioritization_r17 optional
			if ie.intraCG_Prioritization_r17 != nil {
				if err = ie.intraCG_Prioritization_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraCG_Prioritization_r17", err)
				}
			}
			// encode drx_ConfigSL_r17 optional
			if ie.drx_ConfigSL_r17 != nil {
				tmp_drx_ConfigSL_r17 := utils.SetupRelease[*DRX_ConfigSL_r17]{
					Setup: ie.drx_ConfigSL_r17,
				}
				if err = tmp_drx_ConfigSL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drx_ConfigSL_r17", err)
				}
			}
			// encode drx_ConfigExt_v1700 optional
			if ie.drx_ConfigExt_v1700 != nil {
				tmp_drx_ConfigExt_v1700 := utils.SetupRelease[*DRX_ConfigExt_v1700]{
					Setup: ie.drx_ConfigExt_v1700,
				}
				if err = tmp_drx_ConfigExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drx_ConfigExt_v1700", err)
				}
			}
			// encode schedulingRequestID_BFR_r17 optional
			if ie.schedulingRequestID_BFR_r17 != nil {
				if err = ie.schedulingRequestID_BFR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestID_BFR_r17", err)
				}
			}
			// encode schedulingRequestID_BFR2_r17 optional
			if ie.schedulingRequestID_BFR2_r17 != nil {
				if err = ie.schedulingRequestID_BFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestID_BFR2_r17", err)
				}
			}
			// encode schedulingRequestConfig_v1700 optional
			if ie.schedulingRequestConfig_v1700 != nil {
				if err = ie.schedulingRequestConfig_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestConfig_v1700", err)
				}
			}
			// encode tar_Config_r17 optional
			if ie.tar_Config_r17 != nil {
				tmp_tar_Config_r17 := utils.SetupRelease[*TAR_Config_r17]{
					Setup: ie.tar_Config_r17,
				}
				if err = tmp_tar_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tar_Config_r17", err)
				}
			}
			// encode g_RNTI_ConfigToAddModList_r17 optional
			if len(ie.g_RNTI_ConfigToAddModList_r17) > 0 {
				tmp_g_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				for _, i := range ie.g_RNTI_ConfigToAddModList_r17 {
					tmp_g_RNTI_ConfigToAddModList_r17.Value = append(tmp_g_RNTI_ConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_g_RNTI_ConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode g_RNTI_ConfigToAddModList_r17", err)
				}
			}
			// encode g_RNTI_ConfigToReleaseList_r17 optional
			if len(ie.g_RNTI_ConfigToReleaseList_r17) > 0 {
				tmp_g_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				for _, i := range ie.g_RNTI_ConfigToReleaseList_r17 {
					tmp_g_RNTI_ConfigToReleaseList_r17.Value = append(tmp_g_RNTI_ConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_g_RNTI_ConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode g_RNTI_ConfigToReleaseList_r17", err)
				}
			}
			// encode g_CS_RNTI_ConfigToAddModList_r17 optional
			if len(ie.g_CS_RNTI_ConfigToAddModList_r17) > 0 {
				tmp_g_CS_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				for _, i := range ie.g_CS_RNTI_ConfigToAddModList_r17 {
					tmp_g_CS_RNTI_ConfigToAddModList_r17.Value = append(tmp_g_CS_RNTI_ConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_g_CS_RNTI_ConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode g_CS_RNTI_ConfigToAddModList_r17", err)
				}
			}
			// encode g_CS_RNTI_ConfigToReleaseList_r17 optional
			if len(ie.g_CS_RNTI_ConfigToReleaseList_r17) > 0 {
				tmp_g_CS_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				for _, i := range ie.g_CS_RNTI_ConfigToReleaseList_r17 {
					tmp_g_CS_RNTI_ConfigToReleaseList_r17.Value = append(tmp_g_CS_RNTI_ConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_g_CS_RNTI_ConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode g_CS_RNTI_ConfigToReleaseList_r17", err)
				}
			}
			// encode allowCSI_SRS_Tx_MulticastDRX_Active_r17 optional
			if ie.allowCSI_SRS_Tx_MulticastDRX_Active_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.allowCSI_SRS_Tx_MulticastDRX_Active_r17); err != nil {
					return utils.WrapError("Encode allowCSI_SRS_Tx_MulticastDRX_Active_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.schedulingRequestID_PosMG_Request_r17 != nil, ie.drx_LastTransmissionUL_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode schedulingRequestID_PosMG_Request_r17 optional
			if ie.schedulingRequestID_PosMG_Request_r17 != nil {
				if err = ie.schedulingRequestID_PosMG_Request_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestID_PosMG_Request_r17", err)
				}
			}
			// encode drx_LastTransmissionUL_r17 optional
			if ie.drx_LastTransmissionUL_r17 != nil {
				if err = ie.drx_LastTransmissionUL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drx_LastTransmissionUL_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MAC_CellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_ConfigPresent bool
	if drx_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var schedulingRequestConfigPresent bool
	if schedulingRequestConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bsr_ConfigPresent bool
	if bsr_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tag_ConfigPresent bool
	if tag_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var phr_ConfigPresent bool
	if phr_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if drx_ConfigPresent {
		tmp_drx_Config := utils.SetupRelease[*DRX_Config]{}
		if err = tmp_drx_Config.Decode(r); err != nil {
			return utils.WrapError("Decode drx_Config", err)
		}
		ie.drx_Config = tmp_drx_Config.Setup
	}
	if schedulingRequestConfigPresent {
		ie.schedulingRequestConfig = new(SchedulingRequestConfig)
		if err = ie.schedulingRequestConfig.Decode(r); err != nil {
			return utils.WrapError("Decode schedulingRequestConfig", err)
		}
	}
	if bsr_ConfigPresent {
		ie.bsr_Config = new(BSR_Config)
		if err = ie.bsr_Config.Decode(r); err != nil {
			return utils.WrapError("Decode bsr_Config", err)
		}
	}
	if tag_ConfigPresent {
		ie.tag_Config = new(TAG_Config)
		if err = ie.tag_Config.Decode(r); err != nil {
			return utils.WrapError("Decode tag_Config", err)
		}
	}
	if phr_ConfigPresent {
		tmp_phr_Config := utils.SetupRelease[*PHR_Config]{}
		if err = tmp_phr_Config.Decode(r); err != nil {
			return utils.WrapError("Decode phr_Config", err)
		}
		ie.phr_Config = tmp_phr_Config.Setup
	}
	var tmp_bool_skipUplinkTxDynamic bool
	if tmp_bool_skipUplinkTxDynamic, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean skipUplinkTxDynamic", err)
	}
	ie.skipUplinkTxDynamic = tmp_bool_skipUplinkTxDynamic

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			csi_MaskPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dataInactivityTimerPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode csi_Mask optional
			if csi_MaskPresent {
				var tmp_bool_csi_Mask bool
				if tmp_bool_csi_Mask, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode csi_Mask", err)
				}
				ie.csi_Mask = &tmp_bool_csi_Mask
			}
			// decode dataInactivityTimer optional
			if dataInactivityTimerPresent {
				tmp_dataInactivityTimer := utils.SetupRelease[*DataInactivityTimer]{}
				if err = tmp_dataInactivityTimer.Decode(extReader); err != nil {
					return utils.WrapError("Decode dataInactivityTimer", err)
				}
				ie.dataInactivityTimer = tmp_dataInactivityTimer.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			usePreBSR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestID_LBT_SCell_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lch_BasedPrioritization_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestID_BFR_SCell_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			drx_ConfigSecondaryGroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode usePreBSR_r16 optional
			if usePreBSR_r16Present {
				ie.usePreBSR_r16 = new(MAC_CellGroupConfig_usePreBSR_r16)
				if err = ie.usePreBSR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode usePreBSR_r16", err)
				}
			}
			// decode schedulingRequestID_LBT_SCell_r16 optional
			if schedulingRequestID_LBT_SCell_r16Present {
				ie.schedulingRequestID_LBT_SCell_r16 = new(SchedulingRequestId)
				if err = ie.schedulingRequestID_LBT_SCell_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode schedulingRequestID_LBT_SCell_r16", err)
				}
			}
			// decode lch_BasedPrioritization_r16 optional
			if lch_BasedPrioritization_r16Present {
				ie.lch_BasedPrioritization_r16 = new(MAC_CellGroupConfig_lch_BasedPrioritization_r16)
				if err = ie.lch_BasedPrioritization_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lch_BasedPrioritization_r16", err)
				}
			}
			// decode schedulingRequestID_BFR_SCell_r16 optional
			if schedulingRequestID_BFR_SCell_r16Present {
				ie.schedulingRequestID_BFR_SCell_r16 = new(SchedulingRequestId)
				if err = ie.schedulingRequestID_BFR_SCell_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode schedulingRequestID_BFR_SCell_r16", err)
				}
			}
			// decode drx_ConfigSecondaryGroup_r16 optional
			if drx_ConfigSecondaryGroup_r16Present {
				tmp_drx_ConfigSecondaryGroup_r16 := utils.SetupRelease[*DRX_ConfigSecondaryGroup_r16]{}
				if err = tmp_drx_ConfigSecondaryGroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode drx_ConfigSecondaryGroup_r16", err)
				}
				ie.drx_ConfigSecondaryGroup_r16 = tmp_drx_ConfigSecondaryGroup_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			enhancedSkipUplinkTxDynamic_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enhancedSkipUplinkTxConfigured_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enhancedSkipUplinkTxDynamic_r16 optional
			if enhancedSkipUplinkTxDynamic_r16Present {
				ie.enhancedSkipUplinkTxDynamic_r16 = new(MAC_CellGroupConfig_enhancedSkipUplinkTxDynamic_r16)
				if err = ie.enhancedSkipUplinkTxDynamic_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// decode enhancedSkipUplinkTxConfigured_r16 optional
			if enhancedSkipUplinkTxConfigured_r16Present {
				ie.enhancedSkipUplinkTxConfigured_r16 = new(MAC_CellGroupConfig_enhancedSkipUplinkTxConfigured_r16)
				if err = ie.enhancedSkipUplinkTxConfigured_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedSkipUplinkTxConfigured_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			intraCG_Prioritization_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			drx_ConfigSL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			drx_ConfigExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestID_BFR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestID_BFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestConfig_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tar_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			g_RNTI_ConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			g_RNTI_ConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			g_CS_RNTI_ConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			g_CS_RNTI_ConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			allowCSI_SRS_Tx_MulticastDRX_Active_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode intraCG_Prioritization_r17 optional
			if intraCG_Prioritization_r17Present {
				ie.intraCG_Prioritization_r17 = new(MAC_CellGroupConfig_intraCG_Prioritization_r17)
				if err = ie.intraCG_Prioritization_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraCG_Prioritization_r17", err)
				}
			}
			// decode drx_ConfigSL_r17 optional
			if drx_ConfigSL_r17Present {
				tmp_drx_ConfigSL_r17 := utils.SetupRelease[*DRX_ConfigSL_r17]{}
				if err = tmp_drx_ConfigSL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode drx_ConfigSL_r17", err)
				}
				ie.drx_ConfigSL_r17 = tmp_drx_ConfigSL_r17.Setup
			}
			// decode drx_ConfigExt_v1700 optional
			if drx_ConfigExt_v1700Present {
				tmp_drx_ConfigExt_v1700 := utils.SetupRelease[*DRX_ConfigExt_v1700]{}
				if err = tmp_drx_ConfigExt_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode drx_ConfigExt_v1700", err)
				}
				ie.drx_ConfigExt_v1700 = tmp_drx_ConfigExt_v1700.Setup
			}
			// decode schedulingRequestID_BFR_r17 optional
			if schedulingRequestID_BFR_r17Present {
				ie.schedulingRequestID_BFR_r17 = new(SchedulingRequestId)
				if err = ie.schedulingRequestID_BFR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode schedulingRequestID_BFR_r17", err)
				}
			}
			// decode schedulingRequestID_BFR2_r17 optional
			if schedulingRequestID_BFR2_r17Present {
				ie.schedulingRequestID_BFR2_r17 = new(SchedulingRequestId)
				if err = ie.schedulingRequestID_BFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode schedulingRequestID_BFR2_r17", err)
				}
			}
			// decode schedulingRequestConfig_v1700 optional
			if schedulingRequestConfig_v1700Present {
				ie.schedulingRequestConfig_v1700 = new(SchedulingRequestConfig_v1700)
				if err = ie.schedulingRequestConfig_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode schedulingRequestConfig_v1700", err)
				}
			}
			// decode tar_Config_r17 optional
			if tar_Config_r17Present {
				tmp_tar_Config_r17 := utils.SetupRelease[*TAR_Config_r17]{}
				if err = tmp_tar_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode tar_Config_r17", err)
				}
				ie.tar_Config_r17 = tmp_tar_Config_r17.Setup
			}
			// decode g_RNTI_ConfigToAddModList_r17 optional
			if g_RNTI_ConfigToAddModList_r17Present {
				tmp_g_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				fn_g_RNTI_ConfigToAddModList_r17 := func() *MBS_RNTI_SpecificConfig_r17 {
					return new(MBS_RNTI_SpecificConfig_r17)
				}
				if err = tmp_g_RNTI_ConfigToAddModList_r17.Decode(extReader, fn_g_RNTI_ConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode g_RNTI_ConfigToAddModList_r17", err)
				}
				ie.g_RNTI_ConfigToAddModList_r17 = []MBS_RNTI_SpecificConfig_r17{}
				for _, i := range tmp_g_RNTI_ConfigToAddModList_r17.Value {
					ie.g_RNTI_ConfigToAddModList_r17 = append(ie.g_RNTI_ConfigToAddModList_r17, *i)
				}
			}
			// decode g_RNTI_ConfigToReleaseList_r17 optional
			if g_RNTI_ConfigToReleaseList_r17Present {
				tmp_g_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_RNTI_r17}, false)
				fn_g_RNTI_ConfigToReleaseList_r17 := func() *MBS_RNTI_SpecificConfigId_r17 {
					return new(MBS_RNTI_SpecificConfigId_r17)
				}
				if err = tmp_g_RNTI_ConfigToReleaseList_r17.Decode(extReader, fn_g_RNTI_ConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode g_RNTI_ConfigToReleaseList_r17", err)
				}
				ie.g_RNTI_ConfigToReleaseList_r17 = []MBS_RNTI_SpecificConfigId_r17{}
				for _, i := range tmp_g_RNTI_ConfigToReleaseList_r17.Value {
					ie.g_RNTI_ConfigToReleaseList_r17 = append(ie.g_RNTI_ConfigToReleaseList_r17, *i)
				}
			}
			// decode g_CS_RNTI_ConfigToAddModList_r17 optional
			if g_CS_RNTI_ConfigToAddModList_r17Present {
				tmp_g_CS_RNTI_ConfigToAddModList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfig_r17]([]*MBS_RNTI_SpecificConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				fn_g_CS_RNTI_ConfigToAddModList_r17 := func() *MBS_RNTI_SpecificConfig_r17 {
					return new(MBS_RNTI_SpecificConfig_r17)
				}
				if err = tmp_g_CS_RNTI_ConfigToAddModList_r17.Decode(extReader, fn_g_CS_RNTI_ConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode g_CS_RNTI_ConfigToAddModList_r17", err)
				}
				ie.g_CS_RNTI_ConfigToAddModList_r17 = []MBS_RNTI_SpecificConfig_r17{}
				for _, i := range tmp_g_CS_RNTI_ConfigToAddModList_r17.Value {
					ie.g_CS_RNTI_ConfigToAddModList_r17 = append(ie.g_CS_RNTI_ConfigToAddModList_r17, *i)
				}
			}
			// decode g_CS_RNTI_ConfigToReleaseList_r17 optional
			if g_CS_RNTI_ConfigToReleaseList_r17Present {
				tmp_g_CS_RNTI_ConfigToReleaseList_r17 := utils.NewSequence[*MBS_RNTI_SpecificConfigId_r17]([]*MBS_RNTI_SpecificConfigId_r17{}, uper.Constraint{Lb: 1, Ub: maxG_CS_RNTI_r17}, false)
				fn_g_CS_RNTI_ConfigToReleaseList_r17 := func() *MBS_RNTI_SpecificConfigId_r17 {
					return new(MBS_RNTI_SpecificConfigId_r17)
				}
				if err = tmp_g_CS_RNTI_ConfigToReleaseList_r17.Decode(extReader, fn_g_CS_RNTI_ConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode g_CS_RNTI_ConfigToReleaseList_r17", err)
				}
				ie.g_CS_RNTI_ConfigToReleaseList_r17 = []MBS_RNTI_SpecificConfigId_r17{}
				for _, i := range tmp_g_CS_RNTI_ConfigToReleaseList_r17.Value {
					ie.g_CS_RNTI_ConfigToReleaseList_r17 = append(ie.g_CS_RNTI_ConfigToReleaseList_r17, *i)
				}
			}
			// decode allowCSI_SRS_Tx_MulticastDRX_Active_r17 optional
			if allowCSI_SRS_Tx_MulticastDRX_Active_r17Present {
				var tmp_bool_allowCSI_SRS_Tx_MulticastDRX_Active_r17 bool
				if tmp_bool_allowCSI_SRS_Tx_MulticastDRX_Active_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode allowCSI_SRS_Tx_MulticastDRX_Active_r17", err)
				}
				ie.allowCSI_SRS_Tx_MulticastDRX_Active_r17 = &tmp_bool_allowCSI_SRS_Tx_MulticastDRX_Active_r17
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			schedulingRequestID_PosMG_Request_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			drx_LastTransmissionUL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode schedulingRequestID_PosMG_Request_r17 optional
			if schedulingRequestID_PosMG_Request_r17Present {
				ie.schedulingRequestID_PosMG_Request_r17 = new(SchedulingRequestId)
				if err = ie.schedulingRequestID_PosMG_Request_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode schedulingRequestID_PosMG_Request_r17", err)
				}
			}
			// decode drx_LastTransmissionUL_r17 optional
			if drx_LastTransmissionUL_r17Present {
				ie.drx_LastTransmissionUL_r17 = new(MAC_CellGroupConfig_drx_LastTransmissionUL_r17)
				if err = ie.drx_LastTransmissionUL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode drx_LastTransmissionUL_r17", err)
				}
			}
		}
	}
	return nil
}
