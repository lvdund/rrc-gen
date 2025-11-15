package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfig struct {
	tdd_UL_DL_ConfigurationDedicated              *TDD_UL_DL_ConfigDedicated                                       `optional`
	initialDownlinkBWP                            *BWP_DownlinkDedicated                                           `optional`
	downlinkBWP_ToReleaseList                     []BWP_Id                                                         `lb:1,ub:maxNrofBWPs,optional`
	downlinkBWP_ToAddModList                      []BWP_Downlink                                                   `lb:1,ub:maxNrofBWPs,optional`
	firstActiveDownlinkBWP_Id                     *BWP_Id                                                          `optional`
	bwp_InactivityTimer                           *ServingCellConfig_bwp_InactivityTimer                           `optional`
	defaultDownlinkBWP_Id                         *BWP_Id                                                          `optional`
	uplinkConfig                                  *UplinkConfig                                                    `optional`
	supplementaryUplink                           *UplinkConfig                                                    `optional`
	pdcch_ServingCellConfig                       *PDCCH_ServingCellConfig                                         `optional,setuprelease`
	pdsch_ServingCellConfig                       *PDSCH_ServingCellConfig                                         `optional,setuprelease`
	csi_MeasConfig                                *CSI_MeasConfig                                                  `optional,setuprelease`
	sCellDeactivationTimer                        *ServingCellConfig_sCellDeactivationTimer                        `optional`
	crossCarrierSchedulingConfig                  *CrossCarrierSchedulingConfig                                    `optional`
	tag_Id                                        TAG_Id                                                           `madatory`
	dummy1                                        *ServingCellConfig_dummy1                                        `optional`
	pathlossReferenceLinking                      *ServingCellConfig_pathlossReferenceLinking                      `optional`
	servingCellMO                                 *MeasObjectId                                                    `optional`
	lte_CRS_ToMatchAround                         *RateMatchPatternLTE_CRS                                         `optional,ext-1,setuprelease`
	rateMatchPatternToAddModList                  []RateMatchPattern                                               `lb:1,ub:maxNrofRateMatchPatterns,optional,ext-1`
	rateMatchPatternToReleaseList                 []RateMatchPatternId                                             `lb:1,ub:maxNrofRateMatchPatterns,optional,ext-1`
	downlinkChannelBW_PerSCS_List                 []SCS_SpecificCarrier                                            `lb:1,ub:maxSCSs,optional,ext-1`
	supplementaryUplinkRelease_r16                *ServingCellConfig_supplementaryUplinkRelease_r16                `optional,ext-2`
	tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16   *TDD_UL_DL_ConfigDedicated_IAB_MT_r16                            `optional,ext-2`
	dormantBWP_Config_r16                         *DormantBWP_Config_r16                                           `optional,ext-2,setuprelease`
	ca_SlotOffset_r16                             *ServingCellConfig_ca_SlotOffset_r16                             `lb:-2,ub:2,optional,ext-2`
	dummy2                                        *DummyJ                                                          `optional,ext-2,setuprelease`
	intraCellGuardBandsDL_List_r16                []IntraCellGuardBandsPerSCS_r16                                  `lb:1,ub:maxSCSs,optional,ext-2`
	intraCellGuardBandsUL_List_r16                []IntraCellGuardBandsPerSCS_r16                                  `lb:1,ub:maxSCSs,optional,ext-2`
	csi_RS_ValidationWithDCI_r16                  *ServingCellConfig_csi_RS_ValidationWithDCI_r16                  `optional,ext-2`
	lte_CRS_PatternList1_r16                      *LTE_CRS_PatternList_r16                                         `optional,ext-2,setuprelease`
	lte_CRS_PatternList2_r16                      *LTE_CRS_PatternList_r16                                         `optional,ext-2,setuprelease`
	crs_RateMatch_PerCORESETPoolIndex_r16         *ServingCellConfig_crs_RateMatch_PerCORESETPoolIndex_r16         `optional,ext-2`
	enableTwoDefaultTCI_States_r16                *ServingCellConfig_enableTwoDefaultTCI_States_r16                `optional,ext-2`
	enableDefaultTCI_StatePerCoresetPoolIndex_r16 *ServingCellConfig_enableDefaultTCI_StatePerCoresetPoolIndex_r16 `optional,ext-2`
	enableBeamSwitchTiming_r16                    *ServingCellConfig_enableBeamSwitchTiming_r16                    `optional,ext-2`
	cbg_TxDiffTBsProcessingType1_r16              *ServingCellConfig_cbg_TxDiffTBsProcessingType1_r16              `optional,ext-2`
	cbg_TxDiffTBsProcessingType2_r16              *ServingCellConfig_cbg_TxDiffTBsProcessingType2_r16              `optional,ext-2`
	directionalCollisionHandling_r16              *ServingCellConfig_directionalCollisionHandling_r16              `optional,ext-3`
	channelAccessConfig_r16                       *ChannelAccessConfig_r16                                         `optional,ext-3,setuprelease`
	nr_dl_PRS_PDC_Info_r17                        *NR_DL_PRS_PDC_Info_r17                                          `optional,ext-4,setuprelease`
	semiStaticChannelAccessConfigUE_r17           *SemiStaticChannelAccessConfigUE_r17                             `optional,ext-4,setuprelease`
	mimoParam_r17                                 *MIMOParam_r17                                                   `optional,ext-4,setuprelease`
	channelAccessMode2_r17                        *ServingCellConfig_channelAccessMode2_r17                        `optional,ext-4`
	timeDomainHARQ_BundlingType1_r17              *ServingCellConfig_timeDomainHARQ_BundlingType1_r17              `optional,ext-4`
	nrofHARQ_BundlingGroups_r17                   *ServingCellConfig_nrofHARQ_BundlingGroups_r17                   `optional,ext-4`
	fdmed_ReceptionMulticast_r17                  *ServingCellConfig_fdmed_ReceptionMulticast_r17                  `optional,ext-4`
	moreThanOneNackOnlyMode_r17                   *ServingCellConfig_moreThanOneNackOnlyMode_r17                   `optional,ext-4`
	tci_ActivatedConfig_r17                       *TCI_ActivatedConfig_r17                                         `optional,ext-4`
	directionalCollisionHandling_DC_r17           *ServingCellConfig_directionalCollisionHandling_DC_r17           `optional,ext-4`
	lte_NeighCellsCRS_AssistInfoList_r17          *LTE_NeighCellsCRS_AssistInfoList_r17                            `optional,ext-4,setuprelease`
	lte_NeighCellsCRS_Assumptions_r17             *ServingCellConfig_lte_NeighCellsCRS_Assumptions_r17             `optional,ext-5`
}

func (ie *ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.lte_CRS_ToMatchAround != nil || len(ie.rateMatchPatternToAddModList) > 0 || len(ie.rateMatchPatternToReleaseList) > 0 || len(ie.downlinkChannelBW_PerSCS_List) > 0 || ie.supplementaryUplinkRelease_r16 != nil || ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil || ie.dormantBWP_Config_r16 != nil || ie.ca_SlotOffset_r16 != nil || ie.dummy2 != nil || len(ie.intraCellGuardBandsDL_List_r16) > 0 || len(ie.intraCellGuardBandsUL_List_r16) > 0 || ie.csi_RS_ValidationWithDCI_r16 != nil || ie.lte_CRS_PatternList1_r16 != nil || ie.lte_CRS_PatternList2_r16 != nil || ie.crs_RateMatch_PerCORESETPoolIndex_r16 != nil || ie.enableTwoDefaultTCI_States_r16 != nil || ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil || ie.enableBeamSwitchTiming_r16 != nil || ie.cbg_TxDiffTBsProcessingType1_r16 != nil || ie.cbg_TxDiffTBsProcessingType2_r16 != nil || ie.directionalCollisionHandling_r16 != nil || ie.channelAccessConfig_r16 != nil || ie.nr_dl_PRS_PDC_Info_r17 != nil || ie.semiStaticChannelAccessConfigUE_r17 != nil || ie.mimoParam_r17 != nil || ie.channelAccessMode2_r17 != nil || ie.timeDomainHARQ_BundlingType1_r17 != nil || ie.nrofHARQ_BundlingGroups_r17 != nil || ie.fdmed_ReceptionMulticast_r17 != nil || ie.moreThanOneNackOnlyMode_r17 != nil || ie.tci_ActivatedConfig_r17 != nil || ie.directionalCollisionHandling_DC_r17 != nil || ie.lte_NeighCellsCRS_AssistInfoList_r17 != nil || ie.lte_NeighCellsCRS_Assumptions_r17 != nil
	preambleBits := []bool{hasExtensions, ie.tdd_UL_DL_ConfigurationDedicated != nil, ie.initialDownlinkBWP != nil, len(ie.downlinkBWP_ToReleaseList) > 0, len(ie.downlinkBWP_ToAddModList) > 0, ie.firstActiveDownlinkBWP_Id != nil, ie.bwp_InactivityTimer != nil, ie.defaultDownlinkBWP_Id != nil, ie.uplinkConfig != nil, ie.supplementaryUplink != nil, ie.pdcch_ServingCellConfig != nil, ie.pdsch_ServingCellConfig != nil, ie.csi_MeasConfig != nil, ie.sCellDeactivationTimer != nil, ie.crossCarrierSchedulingConfig != nil, ie.dummy1 != nil, ie.pathlossReferenceLinking != nil, ie.servingCellMO != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tdd_UL_DL_ConfigurationDedicated != nil {
		if err = ie.tdd_UL_DL_ConfigurationDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_UL_DL_ConfigurationDedicated", err)
		}
	}
	if ie.initialDownlinkBWP != nil {
		if err = ie.initialDownlinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode initialDownlinkBWP", err)
		}
	}
	if len(ie.downlinkBWP_ToReleaseList) > 0 {
		tmp_downlinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.downlinkBWP_ToReleaseList {
			tmp_downlinkBWP_ToReleaseList.Value = append(tmp_downlinkBWP_ToReleaseList.Value, &i)
		}
		if err = tmp_downlinkBWP_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode downlinkBWP_ToReleaseList", err)
		}
	}
	if len(ie.downlinkBWP_ToAddModList) > 0 {
		tmp_downlinkBWP_ToAddModList := utils.NewSequence[*BWP_Downlink]([]*BWP_Downlink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.downlinkBWP_ToAddModList {
			tmp_downlinkBWP_ToAddModList.Value = append(tmp_downlinkBWP_ToAddModList.Value, &i)
		}
		if err = tmp_downlinkBWP_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode downlinkBWP_ToAddModList", err)
		}
	}
	if ie.firstActiveDownlinkBWP_Id != nil {
		if err = ie.firstActiveDownlinkBWP_Id.Encode(w); err != nil {
			return utils.WrapError("Encode firstActiveDownlinkBWP_Id", err)
		}
	}
	if ie.bwp_InactivityTimer != nil {
		if err = ie.bwp_InactivityTimer.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_InactivityTimer", err)
		}
	}
	if ie.defaultDownlinkBWP_Id != nil {
		if err = ie.defaultDownlinkBWP_Id.Encode(w); err != nil {
			return utils.WrapError("Encode defaultDownlinkBWP_Id", err)
		}
	}
	if ie.uplinkConfig != nil {
		if err = ie.uplinkConfig.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkConfig", err)
		}
	}
	if ie.supplementaryUplink != nil {
		if err = ie.supplementaryUplink.Encode(w); err != nil {
			return utils.WrapError("Encode supplementaryUplink", err)
		}
	}
	if ie.pdcch_ServingCellConfig != nil {
		tmp_pdcch_ServingCellConfig := utils.SetupRelease[*PDCCH_ServingCellConfig]{
			Setup: ie.pdcch_ServingCellConfig,
		}
		if err = tmp_pdcch_ServingCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_ServingCellConfig", err)
		}
	}
	if ie.pdsch_ServingCellConfig != nil {
		tmp_pdsch_ServingCellConfig := utils.SetupRelease[*PDSCH_ServingCellConfig]{
			Setup: ie.pdsch_ServingCellConfig,
		}
		if err = tmp_pdsch_ServingCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ServingCellConfig", err)
		}
	}
	if ie.csi_MeasConfig != nil {
		tmp_csi_MeasConfig := utils.SetupRelease[*CSI_MeasConfig]{
			Setup: ie.csi_MeasConfig,
		}
		if err = tmp_csi_MeasConfig.Encode(w); err != nil {
			return utils.WrapError("Encode csi_MeasConfig", err)
		}
	}
	if ie.sCellDeactivationTimer != nil {
		if err = ie.sCellDeactivationTimer.Encode(w); err != nil {
			return utils.WrapError("Encode sCellDeactivationTimer", err)
		}
	}
	if ie.crossCarrierSchedulingConfig != nil {
		if err = ie.crossCarrierSchedulingConfig.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingConfig", err)
		}
	}
	if err = ie.tag_Id.Encode(w); err != nil {
		return utils.WrapError("Encode tag_Id", err)
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.pathlossReferenceLinking != nil {
		if err = ie.pathlossReferenceLinking.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceLinking", err)
		}
	}
	if ie.servingCellMO != nil {
		if err = ie.servingCellMO.Encode(w); err != nil {
			return utils.WrapError("Encode servingCellMO", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.lte_CRS_ToMatchAround != nil || len(ie.rateMatchPatternToAddModList) > 0 || len(ie.rateMatchPatternToReleaseList) > 0 || len(ie.downlinkChannelBW_PerSCS_List) > 0, ie.supplementaryUplinkRelease_r16 != nil || ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil || ie.dormantBWP_Config_r16 != nil || ie.ca_SlotOffset_r16 != nil || ie.dummy2 != nil || len(ie.intraCellGuardBandsDL_List_r16) > 0 || len(ie.intraCellGuardBandsUL_List_r16) > 0 || ie.csi_RS_ValidationWithDCI_r16 != nil || ie.lte_CRS_PatternList1_r16 != nil || ie.lte_CRS_PatternList2_r16 != nil || ie.crs_RateMatch_PerCORESETPoolIndex_r16 != nil || ie.enableTwoDefaultTCI_States_r16 != nil || ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil || ie.enableBeamSwitchTiming_r16 != nil || ie.cbg_TxDiffTBsProcessingType1_r16 != nil || ie.cbg_TxDiffTBsProcessingType2_r16 != nil, ie.directionalCollisionHandling_r16 != nil || ie.channelAccessConfig_r16 != nil, ie.nr_dl_PRS_PDC_Info_r17 != nil || ie.semiStaticChannelAccessConfigUE_r17 != nil || ie.mimoParam_r17 != nil || ie.channelAccessMode2_r17 != nil || ie.timeDomainHARQ_BundlingType1_r17 != nil || ie.nrofHARQ_BundlingGroups_r17 != nil || ie.fdmed_ReceptionMulticast_r17 != nil || ie.moreThanOneNackOnlyMode_r17 != nil || ie.tci_ActivatedConfig_r17 != nil || ie.directionalCollisionHandling_DC_r17 != nil || ie.lte_NeighCellsCRS_AssistInfoList_r17 != nil, ie.lte_NeighCellsCRS_Assumptions_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.lte_CRS_ToMatchAround != nil, len(ie.rateMatchPatternToAddModList) > 0, len(ie.rateMatchPatternToReleaseList) > 0, len(ie.downlinkChannelBW_PerSCS_List) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode lte_CRS_ToMatchAround optional
			if ie.lte_CRS_ToMatchAround != nil {
				tmp_lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{
					Setup: ie.lte_CRS_ToMatchAround,
				}
				if err = tmp_lte_CRS_ToMatchAround.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lte_CRS_ToMatchAround", err)
				}
			}
			// encode rateMatchPatternToAddModList optional
			if len(ie.rateMatchPatternToAddModList) > 0 {
				tmp_rateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				for _, i := range ie.rateMatchPatternToAddModList {
					tmp_rateMatchPatternToAddModList.Value = append(tmp_rateMatchPatternToAddModList.Value, &i)
				}
				if err = tmp_rateMatchPatternToAddModList.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rateMatchPatternToAddModList", err)
				}
			}
			// encode rateMatchPatternToReleaseList optional
			if len(ie.rateMatchPatternToReleaseList) > 0 {
				tmp_rateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				for _, i := range ie.rateMatchPatternToReleaseList {
					tmp_rateMatchPatternToReleaseList.Value = append(tmp_rateMatchPatternToReleaseList.Value, &i)
				}
				if err = tmp_rateMatchPatternToReleaseList.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rateMatchPatternToReleaseList", err)
				}
			}
			// encode downlinkChannelBW_PerSCS_List optional
			if len(ie.downlinkChannelBW_PerSCS_List) > 0 {
				tmp_downlinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.downlinkChannelBW_PerSCS_List {
					tmp_downlinkChannelBW_PerSCS_List.Value = append(tmp_downlinkChannelBW_PerSCS_List.Value, &i)
				}
				if err = tmp_downlinkChannelBW_PerSCS_List.Encode(extWriter); err != nil {
					return utils.WrapError("Encode downlinkChannelBW_PerSCS_List", err)
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
			optionals_ext_2 := []bool{ie.supplementaryUplinkRelease_r16 != nil, ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil, ie.dormantBWP_Config_r16 != nil, ie.ca_SlotOffset_r16 != nil, ie.dummy2 != nil, len(ie.intraCellGuardBandsDL_List_r16) > 0, len(ie.intraCellGuardBandsUL_List_r16) > 0, ie.csi_RS_ValidationWithDCI_r16 != nil, ie.lte_CRS_PatternList1_r16 != nil, ie.lte_CRS_PatternList2_r16 != nil, ie.crs_RateMatch_PerCORESETPoolIndex_r16 != nil, ie.enableTwoDefaultTCI_States_r16 != nil, ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil, ie.enableBeamSwitchTiming_r16 != nil, ie.cbg_TxDiffTBsProcessingType1_r16 != nil, ie.cbg_TxDiffTBsProcessingType2_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supplementaryUplinkRelease_r16 optional
			if ie.supplementaryUplinkRelease_r16 != nil {
				if err = ie.supplementaryUplinkRelease_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supplementaryUplinkRelease_r16", err)
				}
			}
			// encode tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 optional
			if ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil {
				if err = ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16", err)
				}
			}
			// encode dormantBWP_Config_r16 optional
			if ie.dormantBWP_Config_r16 != nil {
				tmp_dormantBWP_Config_r16 := utils.SetupRelease[*DormantBWP_Config_r16]{
					Setup: ie.dormantBWP_Config_r16,
				}
				if err = tmp_dormantBWP_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dormantBWP_Config_r16", err)
				}
			}
			// encode ca_SlotOffset_r16 optional
			if ie.ca_SlotOffset_r16 != nil {
				if err = ie.ca_SlotOffset_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ca_SlotOffset_r16", err)
				}
			}
			// encode dummy2 optional
			if ie.dummy2 != nil {
				tmp_dummy2 := utils.SetupRelease[*DummyJ]{
					Setup: ie.dummy2,
				}
				if err = tmp_dummy2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy2", err)
				}
			}
			// encode intraCellGuardBandsDL_List_r16 optional
			if len(ie.intraCellGuardBandsDL_List_r16) > 0 {
				tmp_intraCellGuardBandsDL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.intraCellGuardBandsDL_List_r16 {
					tmp_intraCellGuardBandsDL_List_r16.Value = append(tmp_intraCellGuardBandsDL_List_r16.Value, &i)
				}
				if err = tmp_intraCellGuardBandsDL_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraCellGuardBandsDL_List_r16", err)
				}
			}
			// encode intraCellGuardBandsUL_List_r16 optional
			if len(ie.intraCellGuardBandsUL_List_r16) > 0 {
				tmp_intraCellGuardBandsUL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.intraCellGuardBandsUL_List_r16 {
					tmp_intraCellGuardBandsUL_List_r16.Value = append(tmp_intraCellGuardBandsUL_List_r16.Value, &i)
				}
				if err = tmp_intraCellGuardBandsUL_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraCellGuardBandsUL_List_r16", err)
				}
			}
			// encode csi_RS_ValidationWithDCI_r16 optional
			if ie.csi_RS_ValidationWithDCI_r16 != nil {
				if err = ie.csi_RS_ValidationWithDCI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_RS_ValidationWithDCI_r16", err)
				}
			}
			// encode lte_CRS_PatternList1_r16 optional
			if ie.lte_CRS_PatternList1_r16 != nil {
				tmp_lte_CRS_PatternList1_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{
					Setup: ie.lte_CRS_PatternList1_r16,
				}
				if err = tmp_lte_CRS_PatternList1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lte_CRS_PatternList1_r16", err)
				}
			}
			// encode lte_CRS_PatternList2_r16 optional
			if ie.lte_CRS_PatternList2_r16 != nil {
				tmp_lte_CRS_PatternList2_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{
					Setup: ie.lte_CRS_PatternList2_r16,
				}
				if err = tmp_lte_CRS_PatternList2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lte_CRS_PatternList2_r16", err)
				}
			}
			// encode crs_RateMatch_PerCORESETPoolIndex_r16 optional
			if ie.crs_RateMatch_PerCORESETPoolIndex_r16 != nil {
				if err = ie.crs_RateMatch_PerCORESETPoolIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode crs_RateMatch_PerCORESETPoolIndex_r16", err)
				}
			}
			// encode enableTwoDefaultTCI_States_r16 optional
			if ie.enableTwoDefaultTCI_States_r16 != nil {
				if err = ie.enableTwoDefaultTCI_States_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableTwoDefaultTCI_States_r16", err)
				}
			}
			// encode enableDefaultTCI_StatePerCoresetPoolIndex_r16 optional
			if ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil {
				if err = ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableDefaultTCI_StatePerCoresetPoolIndex_r16", err)
				}
			}
			// encode enableBeamSwitchTiming_r16 optional
			if ie.enableBeamSwitchTiming_r16 != nil {
				if err = ie.enableBeamSwitchTiming_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableBeamSwitchTiming_r16", err)
				}
			}
			// encode cbg_TxDiffTBsProcessingType1_r16 optional
			if ie.cbg_TxDiffTBsProcessingType1_r16 != nil {
				if err = ie.cbg_TxDiffTBsProcessingType1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cbg_TxDiffTBsProcessingType1_r16", err)
				}
			}
			// encode cbg_TxDiffTBsProcessingType2_r16 optional
			if ie.cbg_TxDiffTBsProcessingType2_r16 != nil {
				if err = ie.cbg_TxDiffTBsProcessingType2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cbg_TxDiffTBsProcessingType2_r16", err)
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
			optionals_ext_3 := []bool{ie.directionalCollisionHandling_r16 != nil, ie.channelAccessConfig_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode directionalCollisionHandling_r16 optional
			if ie.directionalCollisionHandling_r16 != nil {
				if err = ie.directionalCollisionHandling_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode directionalCollisionHandling_r16", err)
				}
			}
			// encode channelAccessConfig_r16 optional
			if ie.channelAccessConfig_r16 != nil {
				tmp_channelAccessConfig_r16 := utils.SetupRelease[*ChannelAccessConfig_r16]{
					Setup: ie.channelAccessConfig_r16,
				}
				if err = tmp_channelAccessConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessConfig_r16", err)
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
			optionals_ext_4 := []bool{ie.nr_dl_PRS_PDC_Info_r17 != nil, ie.semiStaticChannelAccessConfigUE_r17 != nil, ie.mimoParam_r17 != nil, ie.channelAccessMode2_r17 != nil, ie.timeDomainHARQ_BundlingType1_r17 != nil, ie.nrofHARQ_BundlingGroups_r17 != nil, ie.fdmed_ReceptionMulticast_r17 != nil, ie.moreThanOneNackOnlyMode_r17 != nil, ie.tci_ActivatedConfig_r17 != nil, ie.directionalCollisionHandling_DC_r17 != nil, ie.lte_NeighCellsCRS_AssistInfoList_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nr_dl_PRS_PDC_Info_r17 optional
			if ie.nr_dl_PRS_PDC_Info_r17 != nil {
				tmp_nr_dl_PRS_PDC_Info_r17 := utils.SetupRelease[*NR_DL_PRS_PDC_Info_r17]{
					Setup: ie.nr_dl_PRS_PDC_Info_r17,
				}
				if err = tmp_nr_dl_PRS_PDC_Info_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_dl_PRS_PDC_Info_r17", err)
				}
			}
			// encode semiStaticChannelAccessConfigUE_r17 optional
			if ie.semiStaticChannelAccessConfigUE_r17 != nil {
				tmp_semiStaticChannelAccessConfigUE_r17 := utils.SetupRelease[*SemiStaticChannelAccessConfigUE_r17]{
					Setup: ie.semiStaticChannelAccessConfigUE_r17,
				}
				if err = tmp_semiStaticChannelAccessConfigUE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode semiStaticChannelAccessConfigUE_r17", err)
				}
			}
			// encode mimoParam_r17 optional
			if ie.mimoParam_r17 != nil {
				tmp_mimoParam_r17 := utils.SetupRelease[*MIMOParam_r17]{
					Setup: ie.mimoParam_r17,
				}
				if err = tmp_mimoParam_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mimoParam_r17", err)
				}
			}
			// encode channelAccessMode2_r17 optional
			if ie.channelAccessMode2_r17 != nil {
				if err = ie.channelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessMode2_r17", err)
				}
			}
			// encode timeDomainHARQ_BundlingType1_r17 optional
			if ie.timeDomainHARQ_BundlingType1_r17 != nil {
				if err = ie.timeDomainHARQ_BundlingType1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode timeDomainHARQ_BundlingType1_r17", err)
				}
			}
			// encode nrofHARQ_BundlingGroups_r17 optional
			if ie.nrofHARQ_BundlingGroups_r17 != nil {
				if err = ie.nrofHARQ_BundlingGroups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrofHARQ_BundlingGroups_r17", err)
				}
			}
			// encode fdmed_ReceptionMulticast_r17 optional
			if ie.fdmed_ReceptionMulticast_r17 != nil {
				if err = ie.fdmed_ReceptionMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fdmed_ReceptionMulticast_r17", err)
				}
			}
			// encode moreThanOneNackOnlyMode_r17 optional
			if ie.moreThanOneNackOnlyMode_r17 != nil {
				if err = ie.moreThanOneNackOnlyMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode moreThanOneNackOnlyMode_r17", err)
				}
			}
			// encode tci_ActivatedConfig_r17 optional
			if ie.tci_ActivatedConfig_r17 != nil {
				if err = ie.tci_ActivatedConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tci_ActivatedConfig_r17", err)
				}
			}
			// encode directionalCollisionHandling_DC_r17 optional
			if ie.directionalCollisionHandling_DC_r17 != nil {
				if err = ie.directionalCollisionHandling_DC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode directionalCollisionHandling_DC_r17", err)
				}
			}
			// encode lte_NeighCellsCRS_AssistInfoList_r17 optional
			if ie.lte_NeighCellsCRS_AssistInfoList_r17 != nil {
				tmp_lte_NeighCellsCRS_AssistInfoList_r17 := utils.SetupRelease[*LTE_NeighCellsCRS_AssistInfoList_r17]{
					Setup: ie.lte_NeighCellsCRS_AssistInfoList_r17,
				}
				if err = tmp_lte_NeighCellsCRS_AssistInfoList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lte_NeighCellsCRS_AssistInfoList_r17", err)
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
			optionals_ext_5 := []bool{ie.lte_NeighCellsCRS_Assumptions_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode lte_NeighCellsCRS_Assumptions_r17 optional
			if ie.lte_NeighCellsCRS_Assumptions_r17 != nil {
				if err = ie.lte_NeighCellsCRS_Assumptions_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lte_NeighCellsCRS_Assumptions_r17", err)
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

func (ie *ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_UL_DL_ConfigurationDedicatedPresent bool
	if tdd_UL_DL_ConfigurationDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var initialDownlinkBWPPresent bool
	if initialDownlinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var downlinkBWP_ToReleaseListPresent bool
	if downlinkBWP_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var downlinkBWP_ToAddModListPresent bool
	if downlinkBWP_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var firstActiveDownlinkBWP_IdPresent bool
	if firstActiveDownlinkBWP_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_InactivityTimerPresent bool
	if bwp_InactivityTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var defaultDownlinkBWP_IdPresent bool
	if defaultDownlinkBWP_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkConfigPresent bool
	if uplinkConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supplementaryUplinkPresent bool
	if supplementaryUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_ServingCellConfigPresent bool
	if pdcch_ServingCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ServingCellConfigPresent bool
	if pdsch_ServingCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_MeasConfigPresent bool
	if csi_MeasConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sCellDeactivationTimerPresent bool
	if sCellDeactivationTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingConfigPresent bool
	if crossCarrierSchedulingConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceLinkingPresent bool
	if pathlossReferenceLinkingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var servingCellMOPresent bool
	if servingCellMOPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if tdd_UL_DL_ConfigurationDedicatedPresent {
		ie.tdd_UL_DL_ConfigurationDedicated = new(TDD_UL_DL_ConfigDedicated)
		if err = ie.tdd_UL_DL_ConfigurationDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_UL_DL_ConfigurationDedicated", err)
		}
	}
	if initialDownlinkBWPPresent {
		ie.initialDownlinkBWP = new(BWP_DownlinkDedicated)
		if err = ie.initialDownlinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode initialDownlinkBWP", err)
		}
	}
	if downlinkBWP_ToReleaseListPresent {
		tmp_downlinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_downlinkBWP_ToReleaseList := func() *BWP_Id {
			return new(BWP_Id)
		}
		if err = tmp_downlinkBWP_ToReleaseList.Decode(r, fn_downlinkBWP_ToReleaseList); err != nil {
			return utils.WrapError("Decode downlinkBWP_ToReleaseList", err)
		}
		ie.downlinkBWP_ToReleaseList = []BWP_Id{}
		for _, i := range tmp_downlinkBWP_ToReleaseList.Value {
			ie.downlinkBWP_ToReleaseList = append(ie.downlinkBWP_ToReleaseList, *i)
		}
	}
	if downlinkBWP_ToAddModListPresent {
		tmp_downlinkBWP_ToAddModList := utils.NewSequence[*BWP_Downlink]([]*BWP_Downlink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_downlinkBWP_ToAddModList := func() *BWP_Downlink {
			return new(BWP_Downlink)
		}
		if err = tmp_downlinkBWP_ToAddModList.Decode(r, fn_downlinkBWP_ToAddModList); err != nil {
			return utils.WrapError("Decode downlinkBWP_ToAddModList", err)
		}
		ie.downlinkBWP_ToAddModList = []BWP_Downlink{}
		for _, i := range tmp_downlinkBWP_ToAddModList.Value {
			ie.downlinkBWP_ToAddModList = append(ie.downlinkBWP_ToAddModList, *i)
		}
	}
	if firstActiveDownlinkBWP_IdPresent {
		ie.firstActiveDownlinkBWP_Id = new(BWP_Id)
		if err = ie.firstActiveDownlinkBWP_Id.Decode(r); err != nil {
			return utils.WrapError("Decode firstActiveDownlinkBWP_Id", err)
		}
	}
	if bwp_InactivityTimerPresent {
		ie.bwp_InactivityTimer = new(ServingCellConfig_bwp_InactivityTimer)
		if err = ie.bwp_InactivityTimer.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_InactivityTimer", err)
		}
	}
	if defaultDownlinkBWP_IdPresent {
		ie.defaultDownlinkBWP_Id = new(BWP_Id)
		if err = ie.defaultDownlinkBWP_Id.Decode(r); err != nil {
			return utils.WrapError("Decode defaultDownlinkBWP_Id", err)
		}
	}
	if uplinkConfigPresent {
		ie.uplinkConfig = new(UplinkConfig)
		if err = ie.uplinkConfig.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkConfig", err)
		}
	}
	if supplementaryUplinkPresent {
		ie.supplementaryUplink = new(UplinkConfig)
		if err = ie.supplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode supplementaryUplink", err)
		}
	}
	if pdcch_ServingCellConfigPresent {
		tmp_pdcch_ServingCellConfig := utils.SetupRelease[*PDCCH_ServingCellConfig]{}
		if err = tmp_pdcch_ServingCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_ServingCellConfig", err)
		}
		ie.pdcch_ServingCellConfig = tmp_pdcch_ServingCellConfig.Setup
	}
	if pdsch_ServingCellConfigPresent {
		tmp_pdsch_ServingCellConfig := utils.SetupRelease[*PDSCH_ServingCellConfig]{}
		if err = tmp_pdsch_ServingCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ServingCellConfig", err)
		}
		ie.pdsch_ServingCellConfig = tmp_pdsch_ServingCellConfig.Setup
	}
	if csi_MeasConfigPresent {
		tmp_csi_MeasConfig := utils.SetupRelease[*CSI_MeasConfig]{}
		if err = tmp_csi_MeasConfig.Decode(r); err != nil {
			return utils.WrapError("Decode csi_MeasConfig", err)
		}
		ie.csi_MeasConfig = tmp_csi_MeasConfig.Setup
	}
	if sCellDeactivationTimerPresent {
		ie.sCellDeactivationTimer = new(ServingCellConfig_sCellDeactivationTimer)
		if err = ie.sCellDeactivationTimer.Decode(r); err != nil {
			return utils.WrapError("Decode sCellDeactivationTimer", err)
		}
	}
	if crossCarrierSchedulingConfigPresent {
		ie.crossCarrierSchedulingConfig = new(CrossCarrierSchedulingConfig)
		if err = ie.crossCarrierSchedulingConfig.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingConfig", err)
		}
	}
	if err = ie.tag_Id.Decode(r); err != nil {
		return utils.WrapError("Decode tag_Id", err)
	}
	if dummy1Present {
		ie.dummy1 = new(ServingCellConfig_dummy1)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if pathlossReferenceLinkingPresent {
		ie.pathlossReferenceLinking = new(ServingCellConfig_pathlossReferenceLinking)
		if err = ie.pathlossReferenceLinking.Decode(r); err != nil {
			return utils.WrapError("Decode pathlossReferenceLinking", err)
		}
	}
	if servingCellMOPresent {
		ie.servingCellMO = new(MeasObjectId)
		if err = ie.servingCellMO.Decode(r); err != nil {
			return utils.WrapError("Decode servingCellMO", err)
		}
	}

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

			lte_CRS_ToMatchAroundPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rateMatchPatternToAddModListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rateMatchPatternToReleaseListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			downlinkChannelBW_PerSCS_ListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode lte_CRS_ToMatchAround optional
			if lte_CRS_ToMatchAroundPresent {
				tmp_lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{}
				if err = tmp_lte_CRS_ToMatchAround.Decode(extReader); err != nil {
					return utils.WrapError("Decode lte_CRS_ToMatchAround", err)
				}
				ie.lte_CRS_ToMatchAround = tmp_lte_CRS_ToMatchAround.Setup
			}
			// decode rateMatchPatternToAddModList optional
			if rateMatchPatternToAddModListPresent {
				tmp_rateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				fn_rateMatchPatternToAddModList := func() *RateMatchPattern {
					return new(RateMatchPattern)
				}
				if err = tmp_rateMatchPatternToAddModList.Decode(extReader, fn_rateMatchPatternToAddModList); err != nil {
					return utils.WrapError("Decode rateMatchPatternToAddModList", err)
				}
				ie.rateMatchPatternToAddModList = []RateMatchPattern{}
				for _, i := range tmp_rateMatchPatternToAddModList.Value {
					ie.rateMatchPatternToAddModList = append(ie.rateMatchPatternToAddModList, *i)
				}
			}
			// decode rateMatchPatternToReleaseList optional
			if rateMatchPatternToReleaseListPresent {
				tmp_rateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
				fn_rateMatchPatternToReleaseList := func() *RateMatchPatternId {
					return new(RateMatchPatternId)
				}
				if err = tmp_rateMatchPatternToReleaseList.Decode(extReader, fn_rateMatchPatternToReleaseList); err != nil {
					return utils.WrapError("Decode rateMatchPatternToReleaseList", err)
				}
				ie.rateMatchPatternToReleaseList = []RateMatchPatternId{}
				for _, i := range tmp_rateMatchPatternToReleaseList.Value {
					ie.rateMatchPatternToReleaseList = append(ie.rateMatchPatternToReleaseList, *i)
				}
			}
			// decode downlinkChannelBW_PerSCS_List optional
			if downlinkChannelBW_PerSCS_ListPresent {
				tmp_downlinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_downlinkChannelBW_PerSCS_List := func() *SCS_SpecificCarrier {
					return new(SCS_SpecificCarrier)
				}
				if err = tmp_downlinkChannelBW_PerSCS_List.Decode(extReader, fn_downlinkChannelBW_PerSCS_List); err != nil {
					return utils.WrapError("Decode downlinkChannelBW_PerSCS_List", err)
				}
				ie.downlinkChannelBW_PerSCS_List = []SCS_SpecificCarrier{}
				for _, i := range tmp_downlinkChannelBW_PerSCS_List.Value {
					ie.downlinkChannelBW_PerSCS_List = append(ie.downlinkChannelBW_PerSCS_List, *i)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supplementaryUplinkRelease_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dormantBWP_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ca_SlotOffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dummy2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraCellGuardBandsDL_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraCellGuardBandsUL_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_RS_ValidationWithDCI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lte_CRS_PatternList1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lte_CRS_PatternList2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			crs_RateMatch_PerCORESETPoolIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableTwoDefaultTCI_States_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableDefaultTCI_StatePerCoresetPoolIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableBeamSwitchTiming_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cbg_TxDiffTBsProcessingType1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cbg_TxDiffTBsProcessingType2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supplementaryUplinkRelease_r16 optional
			if supplementaryUplinkRelease_r16Present {
				ie.supplementaryUplinkRelease_r16 = new(ServingCellConfig_supplementaryUplinkRelease_r16)
				if err = ie.supplementaryUplinkRelease_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supplementaryUplinkRelease_r16", err)
				}
			}
			// decode tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 optional
			if tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16Present {
				ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 = new(TDD_UL_DL_ConfigDedicated_IAB_MT_r16)
				if err = ie.tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16", err)
				}
			}
			// decode dormantBWP_Config_r16 optional
			if dormantBWP_Config_r16Present {
				tmp_dormantBWP_Config_r16 := utils.SetupRelease[*DormantBWP_Config_r16]{}
				if err = tmp_dormantBWP_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dormantBWP_Config_r16", err)
				}
				ie.dormantBWP_Config_r16 = tmp_dormantBWP_Config_r16.Setup
			}
			// decode ca_SlotOffset_r16 optional
			if ca_SlotOffset_r16Present {
				ie.ca_SlotOffset_r16 = new(ServingCellConfig_ca_SlotOffset_r16)
				if err = ie.ca_SlotOffset_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ca_SlotOffset_r16", err)
				}
			}
			// decode dummy2 optional
			if dummy2Present {
				tmp_dummy2 := utils.SetupRelease[*DummyJ]{}
				if err = tmp_dummy2.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy2", err)
				}
				ie.dummy2 = tmp_dummy2.Setup
			}
			// decode intraCellGuardBandsDL_List_r16 optional
			if intraCellGuardBandsDL_List_r16Present {
				tmp_intraCellGuardBandsDL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_intraCellGuardBandsDL_List_r16 := func() *IntraCellGuardBandsPerSCS_r16 {
					return new(IntraCellGuardBandsPerSCS_r16)
				}
				if err = tmp_intraCellGuardBandsDL_List_r16.Decode(extReader, fn_intraCellGuardBandsDL_List_r16); err != nil {
					return utils.WrapError("Decode intraCellGuardBandsDL_List_r16", err)
				}
				ie.intraCellGuardBandsDL_List_r16 = []IntraCellGuardBandsPerSCS_r16{}
				for _, i := range tmp_intraCellGuardBandsDL_List_r16.Value {
					ie.intraCellGuardBandsDL_List_r16 = append(ie.intraCellGuardBandsDL_List_r16, *i)
				}
			}
			// decode intraCellGuardBandsUL_List_r16 optional
			if intraCellGuardBandsUL_List_r16Present {
				tmp_intraCellGuardBandsUL_List_r16 := utils.NewSequence[*IntraCellGuardBandsPerSCS_r16]([]*IntraCellGuardBandsPerSCS_r16{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_intraCellGuardBandsUL_List_r16 := func() *IntraCellGuardBandsPerSCS_r16 {
					return new(IntraCellGuardBandsPerSCS_r16)
				}
				if err = tmp_intraCellGuardBandsUL_List_r16.Decode(extReader, fn_intraCellGuardBandsUL_List_r16); err != nil {
					return utils.WrapError("Decode intraCellGuardBandsUL_List_r16", err)
				}
				ie.intraCellGuardBandsUL_List_r16 = []IntraCellGuardBandsPerSCS_r16{}
				for _, i := range tmp_intraCellGuardBandsUL_List_r16.Value {
					ie.intraCellGuardBandsUL_List_r16 = append(ie.intraCellGuardBandsUL_List_r16, *i)
				}
			}
			// decode csi_RS_ValidationWithDCI_r16 optional
			if csi_RS_ValidationWithDCI_r16Present {
				ie.csi_RS_ValidationWithDCI_r16 = new(ServingCellConfig_csi_RS_ValidationWithDCI_r16)
				if err = ie.csi_RS_ValidationWithDCI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_RS_ValidationWithDCI_r16", err)
				}
			}
			// decode lte_CRS_PatternList1_r16 optional
			if lte_CRS_PatternList1_r16Present {
				tmp_lte_CRS_PatternList1_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{}
				if err = tmp_lte_CRS_PatternList1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lte_CRS_PatternList1_r16", err)
				}
				ie.lte_CRS_PatternList1_r16 = tmp_lte_CRS_PatternList1_r16.Setup
			}
			// decode lte_CRS_PatternList2_r16 optional
			if lte_CRS_PatternList2_r16Present {
				tmp_lte_CRS_PatternList2_r16 := utils.SetupRelease[*LTE_CRS_PatternList_r16]{}
				if err = tmp_lte_CRS_PatternList2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lte_CRS_PatternList2_r16", err)
				}
				ie.lte_CRS_PatternList2_r16 = tmp_lte_CRS_PatternList2_r16.Setup
			}
			// decode crs_RateMatch_PerCORESETPoolIndex_r16 optional
			if crs_RateMatch_PerCORESETPoolIndex_r16Present {
				ie.crs_RateMatch_PerCORESETPoolIndex_r16 = new(ServingCellConfig_crs_RateMatch_PerCORESETPoolIndex_r16)
				if err = ie.crs_RateMatch_PerCORESETPoolIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode crs_RateMatch_PerCORESETPoolIndex_r16", err)
				}
			}
			// decode enableTwoDefaultTCI_States_r16 optional
			if enableTwoDefaultTCI_States_r16Present {
				ie.enableTwoDefaultTCI_States_r16 = new(ServingCellConfig_enableTwoDefaultTCI_States_r16)
				if err = ie.enableTwoDefaultTCI_States_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableTwoDefaultTCI_States_r16", err)
				}
			}
			// decode enableDefaultTCI_StatePerCoresetPoolIndex_r16 optional
			if enableDefaultTCI_StatePerCoresetPoolIndex_r16Present {
				ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16 = new(ServingCellConfig_enableDefaultTCI_StatePerCoresetPoolIndex_r16)
				if err = ie.enableDefaultTCI_StatePerCoresetPoolIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableDefaultTCI_StatePerCoresetPoolIndex_r16", err)
				}
			}
			// decode enableBeamSwitchTiming_r16 optional
			if enableBeamSwitchTiming_r16Present {
				ie.enableBeamSwitchTiming_r16 = new(ServingCellConfig_enableBeamSwitchTiming_r16)
				if err = ie.enableBeamSwitchTiming_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableBeamSwitchTiming_r16", err)
				}
			}
			// decode cbg_TxDiffTBsProcessingType1_r16 optional
			if cbg_TxDiffTBsProcessingType1_r16Present {
				ie.cbg_TxDiffTBsProcessingType1_r16 = new(ServingCellConfig_cbg_TxDiffTBsProcessingType1_r16)
				if err = ie.cbg_TxDiffTBsProcessingType1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cbg_TxDiffTBsProcessingType1_r16", err)
				}
			}
			// decode cbg_TxDiffTBsProcessingType2_r16 optional
			if cbg_TxDiffTBsProcessingType2_r16Present {
				ie.cbg_TxDiffTBsProcessingType2_r16 = new(ServingCellConfig_cbg_TxDiffTBsProcessingType2_r16)
				if err = ie.cbg_TxDiffTBsProcessingType2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cbg_TxDiffTBsProcessingType2_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			directionalCollisionHandling_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelAccessConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode directionalCollisionHandling_r16 optional
			if directionalCollisionHandling_r16Present {
				ie.directionalCollisionHandling_r16 = new(ServingCellConfig_directionalCollisionHandling_r16)
				if err = ie.directionalCollisionHandling_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode directionalCollisionHandling_r16", err)
				}
			}
			// decode channelAccessConfig_r16 optional
			if channelAccessConfig_r16Present {
				tmp_channelAccessConfig_r16 := utils.SetupRelease[*ChannelAccessConfig_r16]{}
				if err = tmp_channelAccessConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessConfig_r16", err)
				}
				ie.channelAccessConfig_r16 = tmp_channelAccessConfig_r16.Setup
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			nr_dl_PRS_PDC_Info_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			semiStaticChannelAccessConfigUE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mimoParam_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			timeDomainHARQ_BundlingType1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrofHARQ_BundlingGroups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			fdmed_ReceptionMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			moreThanOneNackOnlyMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tci_ActivatedConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			directionalCollisionHandling_DC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lte_NeighCellsCRS_AssistInfoList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nr_dl_PRS_PDC_Info_r17 optional
			if nr_dl_PRS_PDC_Info_r17Present {
				tmp_nr_dl_PRS_PDC_Info_r17 := utils.SetupRelease[*NR_DL_PRS_PDC_Info_r17]{}
				if err = tmp_nr_dl_PRS_PDC_Info_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_dl_PRS_PDC_Info_r17", err)
				}
				ie.nr_dl_PRS_PDC_Info_r17 = tmp_nr_dl_PRS_PDC_Info_r17.Setup
			}
			// decode semiStaticChannelAccessConfigUE_r17 optional
			if semiStaticChannelAccessConfigUE_r17Present {
				tmp_semiStaticChannelAccessConfigUE_r17 := utils.SetupRelease[*SemiStaticChannelAccessConfigUE_r17]{}
				if err = tmp_semiStaticChannelAccessConfigUE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode semiStaticChannelAccessConfigUE_r17", err)
				}
				ie.semiStaticChannelAccessConfigUE_r17 = tmp_semiStaticChannelAccessConfigUE_r17.Setup
			}
			// decode mimoParam_r17 optional
			if mimoParam_r17Present {
				tmp_mimoParam_r17 := utils.SetupRelease[*MIMOParam_r17]{}
				if err = tmp_mimoParam_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mimoParam_r17", err)
				}
				ie.mimoParam_r17 = tmp_mimoParam_r17.Setup
			}
			// decode channelAccessMode2_r17 optional
			if channelAccessMode2_r17Present {
				ie.channelAccessMode2_r17 = new(ServingCellConfig_channelAccessMode2_r17)
				if err = ie.channelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessMode2_r17", err)
				}
			}
			// decode timeDomainHARQ_BundlingType1_r17 optional
			if timeDomainHARQ_BundlingType1_r17Present {
				ie.timeDomainHARQ_BundlingType1_r17 = new(ServingCellConfig_timeDomainHARQ_BundlingType1_r17)
				if err = ie.timeDomainHARQ_BundlingType1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode timeDomainHARQ_BundlingType1_r17", err)
				}
			}
			// decode nrofHARQ_BundlingGroups_r17 optional
			if nrofHARQ_BundlingGroups_r17Present {
				ie.nrofHARQ_BundlingGroups_r17 = new(ServingCellConfig_nrofHARQ_BundlingGroups_r17)
				if err = ie.nrofHARQ_BundlingGroups_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrofHARQ_BundlingGroups_r17", err)
				}
			}
			// decode fdmed_ReceptionMulticast_r17 optional
			if fdmed_ReceptionMulticast_r17Present {
				ie.fdmed_ReceptionMulticast_r17 = new(ServingCellConfig_fdmed_ReceptionMulticast_r17)
				if err = ie.fdmed_ReceptionMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode fdmed_ReceptionMulticast_r17", err)
				}
			}
			// decode moreThanOneNackOnlyMode_r17 optional
			if moreThanOneNackOnlyMode_r17Present {
				ie.moreThanOneNackOnlyMode_r17 = new(ServingCellConfig_moreThanOneNackOnlyMode_r17)
				if err = ie.moreThanOneNackOnlyMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode moreThanOneNackOnlyMode_r17", err)
				}
			}
			// decode tci_ActivatedConfig_r17 optional
			if tci_ActivatedConfig_r17Present {
				ie.tci_ActivatedConfig_r17 = new(TCI_ActivatedConfig_r17)
				if err = ie.tci_ActivatedConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode tci_ActivatedConfig_r17", err)
				}
			}
			// decode directionalCollisionHandling_DC_r17 optional
			if directionalCollisionHandling_DC_r17Present {
				ie.directionalCollisionHandling_DC_r17 = new(ServingCellConfig_directionalCollisionHandling_DC_r17)
				if err = ie.directionalCollisionHandling_DC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode directionalCollisionHandling_DC_r17", err)
				}
			}
			// decode lte_NeighCellsCRS_AssistInfoList_r17 optional
			if lte_NeighCellsCRS_AssistInfoList_r17Present {
				tmp_lte_NeighCellsCRS_AssistInfoList_r17 := utils.SetupRelease[*LTE_NeighCellsCRS_AssistInfoList_r17]{}
				if err = tmp_lte_NeighCellsCRS_AssistInfoList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode lte_NeighCellsCRS_AssistInfoList_r17", err)
				}
				ie.lte_NeighCellsCRS_AssistInfoList_r17 = tmp_lte_NeighCellsCRS_AssistInfoList_r17.Setup
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			lte_NeighCellsCRS_Assumptions_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode lte_NeighCellsCRS_Assumptions_r17 optional
			if lte_NeighCellsCRS_Assumptions_r17Present {
				ie.lte_NeighCellsCRS_Assumptions_r17 = new(ServingCellConfig_lte_NeighCellsCRS_Assumptions_r17)
				if err = ie.lte_NeighCellsCRS_Assumptions_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode lte_NeighCellsCRS_Assumptions_r17", err)
				}
			}
		}
	}
	return nil
}
