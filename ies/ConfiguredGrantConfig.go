package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfig struct {
	frequencyHopping                  *ConfiguredGrantConfig_frequencyHopping           `optional`
	cg_DMRS_Configuration             DMRS_UplinkConfig                                 `madatory`
	mcs_Table                         *ConfiguredGrantConfig_mcs_Table                  `optional`
	mcs_TableTransformPrecoder        *ConfiguredGrantConfig_mcs_TableTransformPrecoder `optional`
	uci_OnPUSCH                       *CG_UCI_OnPUSCH                                   `optional,setuprelease`
	resourceAllocation                ConfiguredGrantConfig_resourceAllocation          `madatory`
	rbg_Size                          *ConfiguredGrantConfig_rbg_Size                   `optional`
	powerControlLoopToUse             ConfiguredGrantConfig_powerControlLoopToUse       `madatory`
	p0_PUSCH_Alpha                    P0_PUSCH_AlphaSetId                               `madatory`
	transformPrecoder                 *ConfiguredGrantConfig_transformPrecoder          `optional`
	nrofHARQ_Processes                int64                                             `lb:1,ub:16,madatory`
	repK                              ConfiguredGrantConfig_repK                        `madatory`
	repK_RV                           *ConfiguredGrantConfig_repK_RV                    `optional`
	periodicity                       ConfiguredGrantConfig_periodicity                 `madatory`
	configuredGrantTimer              *int64                                            `lb:1,ub:64,optional`
	rrc_ConfiguredUplinkGrant         *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant  `lb:18,ub:18,optional`
	cg_RetransmissionTimer_r16        *int64                                            `lb:1,ub:64,optional,ext-3`
	cg_minDFI_Delay_r16               *ConfiguredGrantConfig_cg_minDFI_Delay_r16        `optional,ext-3`
	cg_nrofPUSCH_InSlot_r16           *int64                                            `lb:1,ub:7,optional,ext-3`
	cg_nrofSlots_r16                  *int64                                            `lb:1,ub:40,optional,ext-3`
	cg_StartingOffsets_r16            *CG_StartingOffsets_r16                           `optional,ext-3`
	cg_UCI_Multiplexing_r16           *ConfiguredGrantConfig_cg_UCI_Multiplexing_r16    `optional,ext-3`
	cg_COT_SharingOffset_r16          *int64                                            `lb:1,ub:39,optional,ext-3`
	betaOffsetCG_UCI_r16              *int64                                            `lb:0,ub:31,optional,ext-3`
	cg_COT_SharingList_r16            []CG_COT_Sharing_r16                              `lb:1,ub:1709,optional,ext-3`
	harq_ProcID_Offset_r16            *int64                                            `lb:0,ub:15,optional,ext-3`
	harq_ProcID_Offset2_r16           *int64                                            `lb:0,ub:15,optional,ext-3`
	configuredGrantConfigIndex_r16    *ConfiguredGrantConfigIndex_r16                   `optional,ext-3`
	configuredGrantConfigIndexMAC_r16 *ConfiguredGrantConfigIndexMAC_r16                `optional,ext-3`
	periodicityExt_r16                *int64                                            `lb:1,ub:5120,optional,ext-3`
	startingFromRV0_r16               *ConfiguredGrantConfig_startingFromRV0_r16        `optional,ext-3`
	phy_PriorityIndex_r16             *ConfiguredGrantConfig_phy_PriorityIndex_r16      `optional,ext-3`
	autonomousTx_r16                  *ConfiguredGrantConfig_autonomousTx_r16           `optional,ext-3`
	cg_betaOffsetsCrossPri0_r17       *BetaOffsetsCrossPriSelCG_r17                     `optional,ext-4,setuprelease`
	cg_betaOffsetsCrossPri1_r17       *BetaOffsetsCrossPriSelCG_r17                     `optional,ext-4,setuprelease`
	mappingPattern_r17                *ConfiguredGrantConfig_mappingPattern_r17         `optional,ext-4`
	sequenceOffsetForRV_r17           *int64                                            `lb:0,ub:3,optional,ext-4`
	p0_PUSCH_Alpha2_r17               *P0_PUSCH_AlphaSetId                              `optional,ext-4`
	powerControlLoopToUse2_r17        *ConfiguredGrantConfig_powerControlLoopToUse2_r17 `optional,ext-4`
	cg_COT_SharingList_r17            []CG_COT_Sharing_r17                              `lb:1,ub:50722,optional,ext-4`
	periodicityExt_r17                *int64                                            `lb:1,ub:40960,optional,ext-4`
	repK_v1710                        *ConfiguredGrantConfig_repK_v1710                 `optional,ext-4`
	nrofHARQ_Processes_v1700          *int64                                            `lb:17,ub:32,optional,ext-4`
	harq_ProcID_Offset2_v1700         *int64                                            `lb:16,ub:31,optional,ext-4`
	configuredGrantTimer_v1700        *int64                                            `lb:33,ub:288,optional,ext-4`
	cg_minDFI_Delay_v1710             *int64                                            `lb:238,ub:3584,optional,ext-4`
	harq_ProcID_Offset_v1730          *int64                                            `lb:16,ub:31,optional,ext-5`
	cg_nrofSlots_r17                  *int64                                            `lb:1,ub:320,optional,ext-5`
}

func (ie *ConfiguredGrantConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.cg_RetransmissionTimer_r16 != nil || ie.cg_minDFI_Delay_r16 != nil || ie.cg_nrofPUSCH_InSlot_r16 != nil || ie.cg_nrofSlots_r16 != nil || ie.cg_StartingOffsets_r16 != nil || ie.cg_UCI_Multiplexing_r16 != nil || ie.cg_COT_SharingOffset_r16 != nil || ie.betaOffsetCG_UCI_r16 != nil || len(ie.cg_COT_SharingList_r16) > 0 || ie.harq_ProcID_Offset_r16 != nil || ie.harq_ProcID_Offset2_r16 != nil || ie.configuredGrantConfigIndex_r16 != nil || ie.configuredGrantConfigIndexMAC_r16 != nil || ie.periodicityExt_r16 != nil || ie.startingFromRV0_r16 != nil || ie.phy_PriorityIndex_r16 != nil || ie.autonomousTx_r16 != nil || ie.cg_betaOffsetsCrossPri0_r17 != nil || ie.cg_betaOffsetsCrossPri1_r17 != nil || ie.mappingPattern_r17 != nil || ie.sequenceOffsetForRV_r17 != nil || ie.p0_PUSCH_Alpha2_r17 != nil || ie.powerControlLoopToUse2_r17 != nil || len(ie.cg_COT_SharingList_r17) > 0 || ie.periodicityExt_r17 != nil || ie.repK_v1710 != nil || ie.nrofHARQ_Processes_v1700 != nil || ie.harq_ProcID_Offset2_v1700 != nil || ie.configuredGrantTimer_v1700 != nil || ie.cg_minDFI_Delay_v1710 != nil || ie.harq_ProcID_Offset_v1730 != nil || ie.cg_nrofSlots_r17 != nil
	preambleBits := []bool{hasExtensions, ie.frequencyHopping != nil, ie.mcs_Table != nil, ie.mcs_TableTransformPrecoder != nil, ie.uci_OnPUSCH != nil, ie.rbg_Size != nil, ie.transformPrecoder != nil, ie.repK_RV != nil, ie.configuredGrantTimer != nil, ie.rrc_ConfiguredUplinkGrant != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.frequencyHopping != nil {
		if err = ie.frequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyHopping", err)
		}
	}
	if err = ie.cg_DMRS_Configuration.Encode(w); err != nil {
		return utils.WrapError("Encode cg_DMRS_Configuration", err)
	}
	if ie.mcs_Table != nil {
		if err = ie.mcs_Table.Encode(w); err != nil {
			return utils.WrapError("Encode mcs_Table", err)
		}
	}
	if ie.mcs_TableTransformPrecoder != nil {
		if err = ie.mcs_TableTransformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode mcs_TableTransformPrecoder", err)
		}
	}
	if ie.uci_OnPUSCH != nil {
		tmp_uci_OnPUSCH := utils.SetupRelease[*CG_UCI_OnPUSCH]{
			Setup: ie.uci_OnPUSCH,
		}
		if err = tmp_uci_OnPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode uci_OnPUSCH", err)
		}
	}
	if err = ie.resourceAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode resourceAllocation", err)
	}
	if ie.rbg_Size != nil {
		if err = ie.rbg_Size.Encode(w); err != nil {
			return utils.WrapError("Encode rbg_Size", err)
		}
	}
	if err = ie.powerControlLoopToUse.Encode(w); err != nil {
		return utils.WrapError("Encode powerControlLoopToUse", err)
	}
	if err = ie.p0_PUSCH_Alpha.Encode(w); err != nil {
		return utils.WrapError("Encode p0_PUSCH_Alpha", err)
	}
	if ie.transformPrecoder != nil {
		if err = ie.transformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode transformPrecoder", err)
		}
	}
	if err = w.WriteInteger(ie.nrofHARQ_Processes, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger nrofHARQ_Processes", err)
	}
	if err = ie.repK.Encode(w); err != nil {
		return utils.WrapError("Encode repK", err)
	}
	if ie.repK_RV != nil {
		if err = ie.repK_RV.Encode(w); err != nil {
			return utils.WrapError("Encode repK_RV", err)
		}
	}
	if err = ie.periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode periodicity", err)
	}
	if ie.configuredGrantTimer != nil {
		if err = w.WriteInteger(*ie.configuredGrantTimer, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode configuredGrantTimer", err)
		}
	}
	if ie.rrc_ConfiguredUplinkGrant != nil {
		if err = ie.rrc_ConfiguredUplinkGrant.Encode(w); err != nil {
			return utils.WrapError("Encode rrc_ConfiguredUplinkGrant", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.cg_RetransmissionTimer_r16 != nil || ie.cg_minDFI_Delay_r16 != nil || ie.cg_nrofPUSCH_InSlot_r16 != nil || ie.cg_nrofSlots_r16 != nil || ie.cg_StartingOffsets_r16 != nil || ie.cg_UCI_Multiplexing_r16 != nil || ie.cg_COT_SharingOffset_r16 != nil || ie.betaOffsetCG_UCI_r16 != nil || len(ie.cg_COT_SharingList_r16) > 0 || ie.harq_ProcID_Offset_r16 != nil || ie.harq_ProcID_Offset2_r16 != nil || ie.configuredGrantConfigIndex_r16 != nil || ie.configuredGrantConfigIndexMAC_r16 != nil || ie.periodicityExt_r16 != nil || ie.startingFromRV0_r16 != nil || ie.phy_PriorityIndex_r16 != nil || ie.autonomousTx_r16 != nil, ie.cg_betaOffsetsCrossPri0_r17 != nil || ie.cg_betaOffsetsCrossPri1_r17 != nil || ie.mappingPattern_r17 != nil || ie.sequenceOffsetForRV_r17 != nil || ie.p0_PUSCH_Alpha2_r17 != nil || ie.powerControlLoopToUse2_r17 != nil || len(ie.cg_COT_SharingList_r17) > 0 || ie.periodicityExt_r17 != nil || ie.repK_v1710 != nil || ie.nrofHARQ_Processes_v1700 != nil || ie.harq_ProcID_Offset2_v1700 != nil || ie.configuredGrantTimer_v1700 != nil || ie.cg_minDFI_Delay_v1710 != nil, ie.harq_ProcID_Offset_v1730 != nil || ie.cg_nrofSlots_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ConfiguredGrantConfig", err)
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.cg_RetransmissionTimer_r16 != nil, ie.cg_minDFI_Delay_r16 != nil, ie.cg_nrofPUSCH_InSlot_r16 != nil, ie.cg_nrofSlots_r16 != nil, ie.cg_StartingOffsets_r16 != nil, ie.cg_UCI_Multiplexing_r16 != nil, ie.cg_COT_SharingOffset_r16 != nil, ie.betaOffsetCG_UCI_r16 != nil, len(ie.cg_COT_SharingList_r16) > 0, ie.harq_ProcID_Offset_r16 != nil, ie.harq_ProcID_Offset2_r16 != nil, ie.configuredGrantConfigIndex_r16 != nil, ie.configuredGrantConfigIndexMAC_r16 != nil, ie.periodicityExt_r16 != nil, ie.startingFromRV0_r16 != nil, ie.phy_PriorityIndex_r16 != nil, ie.autonomousTx_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cg_RetransmissionTimer_r16 optional
			if ie.cg_RetransmissionTimer_r16 != nil {
				if err = extWriter.WriteInteger(*ie.cg_RetransmissionTimer_r16, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
					return utils.WrapError("Encode cg_RetransmissionTimer_r16", err)
				}
			}
			// encode cg_minDFI_Delay_r16 optional
			if ie.cg_minDFI_Delay_r16 != nil {
				if err = ie.cg_minDFI_Delay_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_minDFI_Delay_r16", err)
				}
			}
			// encode cg_nrofPUSCH_InSlot_r16 optional
			if ie.cg_nrofPUSCH_InSlot_r16 != nil {
				if err = extWriter.WriteInteger(*ie.cg_nrofPUSCH_InSlot_r16, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
					return utils.WrapError("Encode cg_nrofPUSCH_InSlot_r16", err)
				}
			}
			// encode cg_nrofSlots_r16 optional
			if ie.cg_nrofSlots_r16 != nil {
				if err = extWriter.WriteInteger(*ie.cg_nrofSlots_r16, &uper.Constraint{Lb: 1, Ub: 40}, false); err != nil {
					return utils.WrapError("Encode cg_nrofSlots_r16", err)
				}
			}
			// encode cg_StartingOffsets_r16 optional
			if ie.cg_StartingOffsets_r16 != nil {
				if err = ie.cg_StartingOffsets_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_StartingOffsets_r16", err)
				}
			}
			// encode cg_UCI_Multiplexing_r16 optional
			if ie.cg_UCI_Multiplexing_r16 != nil {
				if err = ie.cg_UCI_Multiplexing_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_UCI_Multiplexing_r16", err)
				}
			}
			// encode cg_COT_SharingOffset_r16 optional
			if ie.cg_COT_SharingOffset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.cg_COT_SharingOffset_r16, &uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
					return utils.WrapError("Encode cg_COT_SharingOffset_r16", err)
				}
			}
			// encode betaOffsetCG_UCI_r16 optional
			if ie.betaOffsetCG_UCI_r16 != nil {
				if err = extWriter.WriteInteger(*ie.betaOffsetCG_UCI_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode betaOffsetCG_UCI_r16", err)
				}
			}
			// encode cg_COT_SharingList_r16 optional
			if len(ie.cg_COT_SharingList_r16) > 0 {
				tmp_cg_COT_SharingList_r16 := utils.NewSequence[*CG_COT_Sharing_r16]([]*CG_COT_Sharing_r16{}, uper.Constraint{Lb: 1, Ub: 1709}, false)
				for _, i := range ie.cg_COT_SharingList_r16 {
					tmp_cg_COT_SharingList_r16.Value = append(tmp_cg_COT_SharingList_r16.Value, &i)
				}
				if err = tmp_cg_COT_SharingList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_COT_SharingList_r16", err)
				}
			}
			// encode harq_ProcID_Offset_r16 optional
			if ie.harq_ProcID_Offset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcID_Offset_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode harq_ProcID_Offset_r16", err)
				}
			}
			// encode harq_ProcID_Offset2_r16 optional
			if ie.harq_ProcID_Offset2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcID_Offset2_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode harq_ProcID_Offset2_r16", err)
				}
			}
			// encode configuredGrantConfigIndex_r16 optional
			if ie.configuredGrantConfigIndex_r16 != nil {
				if err = ie.configuredGrantConfigIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredGrantConfigIndex_r16", err)
				}
			}
			// encode configuredGrantConfigIndexMAC_r16 optional
			if ie.configuredGrantConfigIndexMAC_r16 != nil {
				if err = ie.configuredGrantConfigIndexMAC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredGrantConfigIndexMAC_r16", err)
				}
			}
			// encode periodicityExt_r16 optional
			if ie.periodicityExt_r16 != nil {
				if err = extWriter.WriteInteger(*ie.periodicityExt_r16, &uper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Encode periodicityExt_r16", err)
				}
			}
			// encode startingFromRV0_r16 optional
			if ie.startingFromRV0_r16 != nil {
				if err = ie.startingFromRV0_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode startingFromRV0_r16", err)
				}
			}
			// encode phy_PriorityIndex_r16 optional
			if ie.phy_PriorityIndex_r16 != nil {
				if err = ie.phy_PriorityIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode phy_PriorityIndex_r16", err)
				}
			}
			// encode autonomousTx_r16 optional
			if ie.autonomousTx_r16 != nil {
				if err = ie.autonomousTx_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode autonomousTx_r16", err)
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
			optionals_ext_4 := []bool{ie.cg_betaOffsetsCrossPri0_r17 != nil, ie.cg_betaOffsetsCrossPri1_r17 != nil, ie.mappingPattern_r17 != nil, ie.sequenceOffsetForRV_r17 != nil, ie.p0_PUSCH_Alpha2_r17 != nil, ie.powerControlLoopToUse2_r17 != nil, len(ie.cg_COT_SharingList_r17) > 0, ie.periodicityExt_r17 != nil, ie.repK_v1710 != nil, ie.nrofHARQ_Processes_v1700 != nil, ie.harq_ProcID_Offset2_v1700 != nil, ie.configuredGrantTimer_v1700 != nil, ie.cg_minDFI_Delay_v1710 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cg_betaOffsetsCrossPri0_r17 optional
			if ie.cg_betaOffsetsCrossPri0_r17 != nil {
				tmp_cg_betaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{
					Setup: ie.cg_betaOffsetsCrossPri0_r17,
				}
				if err = tmp_cg_betaOffsetsCrossPri0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_betaOffsetsCrossPri0_r17", err)
				}
			}
			// encode cg_betaOffsetsCrossPri1_r17 optional
			if ie.cg_betaOffsetsCrossPri1_r17 != nil {
				tmp_cg_betaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{
					Setup: ie.cg_betaOffsetsCrossPri1_r17,
				}
				if err = tmp_cg_betaOffsetsCrossPri1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_betaOffsetsCrossPri1_r17", err)
				}
			}
			// encode mappingPattern_r17 optional
			if ie.mappingPattern_r17 != nil {
				if err = ie.mappingPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mappingPattern_r17", err)
				}
			}
			// encode sequenceOffsetForRV_r17 optional
			if ie.sequenceOffsetForRV_r17 != nil {
				if err = extWriter.WriteInteger(*ie.sequenceOffsetForRV_r17, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode sequenceOffsetForRV_r17", err)
				}
			}
			// encode p0_PUSCH_Alpha2_r17 optional
			if ie.p0_PUSCH_Alpha2_r17 != nil {
				if err = ie.p0_PUSCH_Alpha2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p0_PUSCH_Alpha2_r17", err)
				}
			}
			// encode powerControlLoopToUse2_r17 optional
			if ie.powerControlLoopToUse2_r17 != nil {
				if err = ie.powerControlLoopToUse2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode powerControlLoopToUse2_r17", err)
				}
			}
			// encode cg_COT_SharingList_r17 optional
			if len(ie.cg_COT_SharingList_r17) > 0 {
				tmp_cg_COT_SharingList_r17 := utils.NewSequence[*CG_COT_Sharing_r17]([]*CG_COT_Sharing_r17{}, uper.Constraint{Lb: 1, Ub: 50722}, false)
				for _, i := range ie.cg_COT_SharingList_r17 {
					tmp_cg_COT_SharingList_r17.Value = append(tmp_cg_COT_SharingList_r17.Value, &i)
				}
				if err = tmp_cg_COT_SharingList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_COT_SharingList_r17", err)
				}
			}
			// encode periodicityExt_r17 optional
			if ie.periodicityExt_r17 != nil {
				if err = extWriter.WriteInteger(*ie.periodicityExt_r17, &uper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Encode periodicityExt_r17", err)
				}
			}
			// encode repK_v1710 optional
			if ie.repK_v1710 != nil {
				if err = ie.repK_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode repK_v1710", err)
				}
			}
			// encode nrofHARQ_Processes_v1700 optional
			if ie.nrofHARQ_Processes_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.nrofHARQ_Processes_v1700, &uper.Constraint{Lb: 17, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode nrofHARQ_Processes_v1700", err)
				}
			}
			// encode harq_ProcID_Offset2_v1700 optional
			if ie.harq_ProcID_Offset2_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcID_Offset2_v1700, &uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode harq_ProcID_Offset2_v1700", err)
				}
			}
			// encode configuredGrantTimer_v1700 optional
			if ie.configuredGrantTimer_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.configuredGrantTimer_v1700, &uper.Constraint{Lb: 33, Ub: 288}, false); err != nil {
					return utils.WrapError("Encode configuredGrantTimer_v1700", err)
				}
			}
			// encode cg_minDFI_Delay_v1710 optional
			if ie.cg_minDFI_Delay_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.cg_minDFI_Delay_v1710, &uper.Constraint{Lb: 238, Ub: 3584}, false); err != nil {
					return utils.WrapError("Encode cg_minDFI_Delay_v1710", err)
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
			optionals_ext_5 := []bool{ie.harq_ProcID_Offset_v1730 != nil, ie.cg_nrofSlots_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode harq_ProcID_Offset_v1730 optional
			if ie.harq_ProcID_Offset_v1730 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcID_Offset_v1730, &uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode harq_ProcID_Offset_v1730", err)
				}
			}
			// encode cg_nrofSlots_r17 optional
			if ie.cg_nrofSlots_r17 != nil {
				if err = extWriter.WriteInteger(*ie.cg_nrofSlots_r17, &uper.Constraint{Lb: 1, Ub: 320}, false); err != nil {
					return utils.WrapError("Encode cg_nrofSlots_r17", err)
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

func (ie *ConfiguredGrantConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyHoppingPresent bool
	if frequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mcs_TablePresent bool
	if mcs_TablePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mcs_TableTransformPrecoderPresent bool
	if mcs_TableTransformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uci_OnPUSCHPresent bool
	if uci_OnPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rbg_SizePresent bool
	if rbg_SizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var transformPrecoderPresent bool
	if transformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var repK_RVPresent bool
	if repK_RVPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredGrantTimerPresent bool
	if configuredGrantTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rrc_ConfiguredUplinkGrantPresent bool
	if rrc_ConfiguredUplinkGrantPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyHoppingPresent {
		ie.frequencyHopping = new(ConfiguredGrantConfig_frequencyHopping)
		if err = ie.frequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyHopping", err)
		}
	}
	if err = ie.cg_DMRS_Configuration.Decode(r); err != nil {
		return utils.WrapError("Decode cg_DMRS_Configuration", err)
	}
	if mcs_TablePresent {
		ie.mcs_Table = new(ConfiguredGrantConfig_mcs_Table)
		if err = ie.mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_Table", err)
		}
	}
	if mcs_TableTransformPrecoderPresent {
		ie.mcs_TableTransformPrecoder = new(ConfiguredGrantConfig_mcs_TableTransformPrecoder)
		if err = ie.mcs_TableTransformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_TableTransformPrecoder", err)
		}
	}
	if uci_OnPUSCHPresent {
		tmp_uci_OnPUSCH := utils.SetupRelease[*CG_UCI_OnPUSCH]{}
		if err = tmp_uci_OnPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode uci_OnPUSCH", err)
		}
		ie.uci_OnPUSCH = tmp_uci_OnPUSCH.Setup
	}
	if err = ie.resourceAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode resourceAllocation", err)
	}
	if rbg_SizePresent {
		ie.rbg_Size = new(ConfiguredGrantConfig_rbg_Size)
		if err = ie.rbg_Size.Decode(r); err != nil {
			return utils.WrapError("Decode rbg_Size", err)
		}
	}
	if err = ie.powerControlLoopToUse.Decode(r); err != nil {
		return utils.WrapError("Decode powerControlLoopToUse", err)
	}
	if err = ie.p0_PUSCH_Alpha.Decode(r); err != nil {
		return utils.WrapError("Decode p0_PUSCH_Alpha", err)
	}
	if transformPrecoderPresent {
		ie.transformPrecoder = new(ConfiguredGrantConfig_transformPrecoder)
		if err = ie.transformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode transformPrecoder", err)
		}
	}
	var tmp_int_nrofHARQ_Processes int64
	if tmp_int_nrofHARQ_Processes, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger nrofHARQ_Processes", err)
	}
	ie.nrofHARQ_Processes = tmp_int_nrofHARQ_Processes
	if err = ie.repK.Decode(r); err != nil {
		return utils.WrapError("Decode repK", err)
	}
	if repK_RVPresent {
		ie.repK_RV = new(ConfiguredGrantConfig_repK_RV)
		if err = ie.repK_RV.Decode(r); err != nil {
			return utils.WrapError("Decode repK_RV", err)
		}
	}
	if err = ie.periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode periodicity", err)
	}
	if configuredGrantTimerPresent {
		var tmp_int_configuredGrantTimer int64
		if tmp_int_configuredGrantTimer, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode configuredGrantTimer", err)
		}
		ie.configuredGrantTimer = &tmp_int_configuredGrantTimer
	}
	if rrc_ConfiguredUplinkGrantPresent {
		ie.rrc_ConfiguredUplinkGrant = new(ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant)
		if err = ie.rrc_ConfiguredUplinkGrant.Decode(r); err != nil {
			return utils.WrapError("Decode rrc_ConfiguredUplinkGrant", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			cg_RetransmissionTimer_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_minDFI_Delay_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_nrofPUSCH_InSlot_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_nrofSlots_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_StartingOffsets_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_UCI_Multiplexing_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_COT_SharingOffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			betaOffsetCG_UCI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_COT_SharingList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcID_Offset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcID_Offset2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredGrantConfigIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredGrantConfigIndexMAC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			periodicityExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			startingFromRV0_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			phy_PriorityIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			autonomousTx_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cg_RetransmissionTimer_r16 optional
			if cg_RetransmissionTimer_r16Present {
				var tmp_int_cg_RetransmissionTimer_r16 int64
				if tmp_int_cg_RetransmissionTimer_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
					return utils.WrapError("Decode cg_RetransmissionTimer_r16", err)
				}
				ie.cg_RetransmissionTimer_r16 = &tmp_int_cg_RetransmissionTimer_r16
			}
			// decode cg_minDFI_Delay_r16 optional
			if cg_minDFI_Delay_r16Present {
				ie.cg_minDFI_Delay_r16 = new(ConfiguredGrantConfig_cg_minDFI_Delay_r16)
				if err = ie.cg_minDFI_Delay_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_minDFI_Delay_r16", err)
				}
			}
			// decode cg_nrofPUSCH_InSlot_r16 optional
			if cg_nrofPUSCH_InSlot_r16Present {
				var tmp_int_cg_nrofPUSCH_InSlot_r16 int64
				if tmp_int_cg_nrofPUSCH_InSlot_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
					return utils.WrapError("Decode cg_nrofPUSCH_InSlot_r16", err)
				}
				ie.cg_nrofPUSCH_InSlot_r16 = &tmp_int_cg_nrofPUSCH_InSlot_r16
			}
			// decode cg_nrofSlots_r16 optional
			if cg_nrofSlots_r16Present {
				var tmp_int_cg_nrofSlots_r16 int64
				if tmp_int_cg_nrofSlots_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 40}, false); err != nil {
					return utils.WrapError("Decode cg_nrofSlots_r16", err)
				}
				ie.cg_nrofSlots_r16 = &tmp_int_cg_nrofSlots_r16
			}
			// decode cg_StartingOffsets_r16 optional
			if cg_StartingOffsets_r16Present {
				ie.cg_StartingOffsets_r16 = new(CG_StartingOffsets_r16)
				if err = ie.cg_StartingOffsets_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_StartingOffsets_r16", err)
				}
			}
			// decode cg_UCI_Multiplexing_r16 optional
			if cg_UCI_Multiplexing_r16Present {
				ie.cg_UCI_Multiplexing_r16 = new(ConfiguredGrantConfig_cg_UCI_Multiplexing_r16)
				if err = ie.cg_UCI_Multiplexing_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_UCI_Multiplexing_r16", err)
				}
			}
			// decode cg_COT_SharingOffset_r16 optional
			if cg_COT_SharingOffset_r16Present {
				var tmp_int_cg_COT_SharingOffset_r16 int64
				if tmp_int_cg_COT_SharingOffset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
					return utils.WrapError("Decode cg_COT_SharingOffset_r16", err)
				}
				ie.cg_COT_SharingOffset_r16 = &tmp_int_cg_COT_SharingOffset_r16
			}
			// decode betaOffsetCG_UCI_r16 optional
			if betaOffsetCG_UCI_r16Present {
				var tmp_int_betaOffsetCG_UCI_r16 int64
				if tmp_int_betaOffsetCG_UCI_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode betaOffsetCG_UCI_r16", err)
				}
				ie.betaOffsetCG_UCI_r16 = &tmp_int_betaOffsetCG_UCI_r16
			}
			// decode cg_COT_SharingList_r16 optional
			if cg_COT_SharingList_r16Present {
				tmp_cg_COT_SharingList_r16 := utils.NewSequence[*CG_COT_Sharing_r16]([]*CG_COT_Sharing_r16{}, uper.Constraint{Lb: 1, Ub: 1709}, false)
				fn_cg_COT_SharingList_r16 := func() *CG_COT_Sharing_r16 {
					return new(CG_COT_Sharing_r16)
				}
				if err = tmp_cg_COT_SharingList_r16.Decode(extReader, fn_cg_COT_SharingList_r16); err != nil {
					return utils.WrapError("Decode cg_COT_SharingList_r16", err)
				}
				ie.cg_COT_SharingList_r16 = []CG_COT_Sharing_r16{}
				for _, i := range tmp_cg_COT_SharingList_r16.Value {
					ie.cg_COT_SharingList_r16 = append(ie.cg_COT_SharingList_r16, *i)
				}
			}
			// decode harq_ProcID_Offset_r16 optional
			if harq_ProcID_Offset_r16Present {
				var tmp_int_harq_ProcID_Offset_r16 int64
				if tmp_int_harq_ProcID_Offset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode harq_ProcID_Offset_r16", err)
				}
				ie.harq_ProcID_Offset_r16 = &tmp_int_harq_ProcID_Offset_r16
			}
			// decode harq_ProcID_Offset2_r16 optional
			if harq_ProcID_Offset2_r16Present {
				var tmp_int_harq_ProcID_Offset2_r16 int64
				if tmp_int_harq_ProcID_Offset2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode harq_ProcID_Offset2_r16", err)
				}
				ie.harq_ProcID_Offset2_r16 = &tmp_int_harq_ProcID_Offset2_r16
			}
			// decode configuredGrantConfigIndex_r16 optional
			if configuredGrantConfigIndex_r16Present {
				ie.configuredGrantConfigIndex_r16 = new(ConfiguredGrantConfigIndex_r16)
				if err = ie.configuredGrantConfigIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredGrantConfigIndex_r16", err)
				}
			}
			// decode configuredGrantConfigIndexMAC_r16 optional
			if configuredGrantConfigIndexMAC_r16Present {
				ie.configuredGrantConfigIndexMAC_r16 = new(ConfiguredGrantConfigIndexMAC_r16)
				if err = ie.configuredGrantConfigIndexMAC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredGrantConfigIndexMAC_r16", err)
				}
			}
			// decode periodicityExt_r16 optional
			if periodicityExt_r16Present {
				var tmp_int_periodicityExt_r16 int64
				if tmp_int_periodicityExt_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Decode periodicityExt_r16", err)
				}
				ie.periodicityExt_r16 = &tmp_int_periodicityExt_r16
			}
			// decode startingFromRV0_r16 optional
			if startingFromRV0_r16Present {
				ie.startingFromRV0_r16 = new(ConfiguredGrantConfig_startingFromRV0_r16)
				if err = ie.startingFromRV0_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode startingFromRV0_r16", err)
				}
			}
			// decode phy_PriorityIndex_r16 optional
			if phy_PriorityIndex_r16Present {
				ie.phy_PriorityIndex_r16 = new(ConfiguredGrantConfig_phy_PriorityIndex_r16)
				if err = ie.phy_PriorityIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode phy_PriorityIndex_r16", err)
				}
			}
			// decode autonomousTx_r16 optional
			if autonomousTx_r16Present {
				ie.autonomousTx_r16 = new(ConfiguredGrantConfig_autonomousTx_r16)
				if err = ie.autonomousTx_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode autonomousTx_r16", err)
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

			cg_betaOffsetsCrossPri0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_betaOffsetsCrossPri1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mappingPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sequenceOffsetForRV_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			p0_PUSCH_Alpha2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			powerControlLoopToUse2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_COT_SharingList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			periodicityExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			repK_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrofHARQ_Processes_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcID_Offset2_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredGrantTimer_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_minDFI_Delay_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cg_betaOffsetsCrossPri0_r17 optional
			if cg_betaOffsetsCrossPri0_r17Present {
				tmp_cg_betaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{}
				if err = tmp_cg_betaOffsetsCrossPri0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_betaOffsetsCrossPri0_r17", err)
				}
				ie.cg_betaOffsetsCrossPri0_r17 = tmp_cg_betaOffsetsCrossPri0_r17.Setup
			}
			// decode cg_betaOffsetsCrossPri1_r17 optional
			if cg_betaOffsetsCrossPri1_r17Present {
				tmp_cg_betaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelCG_r17]{}
				if err = tmp_cg_betaOffsetsCrossPri1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_betaOffsetsCrossPri1_r17", err)
				}
				ie.cg_betaOffsetsCrossPri1_r17 = tmp_cg_betaOffsetsCrossPri1_r17.Setup
			}
			// decode mappingPattern_r17 optional
			if mappingPattern_r17Present {
				ie.mappingPattern_r17 = new(ConfiguredGrantConfig_mappingPattern_r17)
				if err = ie.mappingPattern_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mappingPattern_r17", err)
				}
			}
			// decode sequenceOffsetForRV_r17 optional
			if sequenceOffsetForRV_r17Present {
				var tmp_int_sequenceOffsetForRV_r17 int64
				if tmp_int_sequenceOffsetForRV_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Decode sequenceOffsetForRV_r17", err)
				}
				ie.sequenceOffsetForRV_r17 = &tmp_int_sequenceOffsetForRV_r17
			}
			// decode p0_PUSCH_Alpha2_r17 optional
			if p0_PUSCH_Alpha2_r17Present {
				ie.p0_PUSCH_Alpha2_r17 = new(P0_PUSCH_AlphaSetId)
				if err = ie.p0_PUSCH_Alpha2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode p0_PUSCH_Alpha2_r17", err)
				}
			}
			// decode powerControlLoopToUse2_r17 optional
			if powerControlLoopToUse2_r17Present {
				ie.powerControlLoopToUse2_r17 = new(ConfiguredGrantConfig_powerControlLoopToUse2_r17)
				if err = ie.powerControlLoopToUse2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode powerControlLoopToUse2_r17", err)
				}
			}
			// decode cg_COT_SharingList_r17 optional
			if cg_COT_SharingList_r17Present {
				tmp_cg_COT_SharingList_r17 := utils.NewSequence[*CG_COT_Sharing_r17]([]*CG_COT_Sharing_r17{}, uper.Constraint{Lb: 1, Ub: 50722}, false)
				fn_cg_COT_SharingList_r17 := func() *CG_COT_Sharing_r17 {
					return new(CG_COT_Sharing_r17)
				}
				if err = tmp_cg_COT_SharingList_r17.Decode(extReader, fn_cg_COT_SharingList_r17); err != nil {
					return utils.WrapError("Decode cg_COT_SharingList_r17", err)
				}
				ie.cg_COT_SharingList_r17 = []CG_COT_Sharing_r17{}
				for _, i := range tmp_cg_COT_SharingList_r17.Value {
					ie.cg_COT_SharingList_r17 = append(ie.cg_COT_SharingList_r17, *i)
				}
			}
			// decode periodicityExt_r17 optional
			if periodicityExt_r17Present {
				var tmp_int_periodicityExt_r17 int64
				if tmp_int_periodicityExt_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Decode periodicityExt_r17", err)
				}
				ie.periodicityExt_r17 = &tmp_int_periodicityExt_r17
			}
			// decode repK_v1710 optional
			if repK_v1710Present {
				ie.repK_v1710 = new(ConfiguredGrantConfig_repK_v1710)
				if err = ie.repK_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode repK_v1710", err)
				}
			}
			// decode nrofHARQ_Processes_v1700 optional
			if nrofHARQ_Processes_v1700Present {
				var tmp_int_nrofHARQ_Processes_v1700 int64
				if tmp_int_nrofHARQ_Processes_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 17, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode nrofHARQ_Processes_v1700", err)
				}
				ie.nrofHARQ_Processes_v1700 = &tmp_int_nrofHARQ_Processes_v1700
			}
			// decode harq_ProcID_Offset2_v1700 optional
			if harq_ProcID_Offset2_v1700Present {
				var tmp_int_harq_ProcID_Offset2_v1700 int64
				if tmp_int_harq_ProcID_Offset2_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode harq_ProcID_Offset2_v1700", err)
				}
				ie.harq_ProcID_Offset2_v1700 = &tmp_int_harq_ProcID_Offset2_v1700
			}
			// decode configuredGrantTimer_v1700 optional
			if configuredGrantTimer_v1700Present {
				var tmp_int_configuredGrantTimer_v1700 int64
				if tmp_int_configuredGrantTimer_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 33, Ub: 288}, false); err != nil {
					return utils.WrapError("Decode configuredGrantTimer_v1700", err)
				}
				ie.configuredGrantTimer_v1700 = &tmp_int_configuredGrantTimer_v1700
			}
			// decode cg_minDFI_Delay_v1710 optional
			if cg_minDFI_Delay_v1710Present {
				var tmp_int_cg_minDFI_Delay_v1710 int64
				if tmp_int_cg_minDFI_Delay_v1710, err = extReader.ReadInteger(&uper.Constraint{Lb: 238, Ub: 3584}, false); err != nil {
					return utils.WrapError("Decode cg_minDFI_Delay_v1710", err)
				}
				ie.cg_minDFI_Delay_v1710 = &tmp_int_cg_minDFI_Delay_v1710
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			harq_ProcID_Offset_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_nrofSlots_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode harq_ProcID_Offset_v1730 optional
			if harq_ProcID_Offset_v1730Present {
				var tmp_int_harq_ProcID_Offset_v1730 int64
				if tmp_int_harq_ProcID_Offset_v1730, err = extReader.ReadInteger(&uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode harq_ProcID_Offset_v1730", err)
				}
				ie.harq_ProcID_Offset_v1730 = &tmp_int_harq_ProcID_Offset_v1730
			}
			// decode cg_nrofSlots_r17 optional
			if cg_nrofSlots_r17Present {
				var tmp_int_cg_nrofSlots_r17 int64
				if tmp_int_cg_nrofSlots_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 320}, false); err != nil {
					return utils.WrapError("Decode cg_nrofSlots_r17", err)
				}
				ie.cg_nrofSlots_r17 = &tmp_int_cg_nrofSlots_r17
			}
		}
	}
	return nil
}
