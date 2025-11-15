package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_Config struct {
	dataScramblingIdentityPUSCH                     *int64                                                      `lb:0,ub:1023,optional`
	txConfig                                        *PUSCH_Config_txConfig                                      `optional`
	dmrs_UplinkForPUSCH_MappingTypeA                *DMRS_UplinkConfig                                          `optional,setuprelease`
	dmrs_UplinkForPUSCH_MappingTypeB                *DMRS_UplinkConfig                                          `optional,setuprelease`
	pusch_PowerControl                              *PUSCH_PowerControl                                         `optional`
	frequencyHopping                                *PUSCH_Config_frequencyHopping                              `optional`
	frequencyHoppingOffsetLists                     []int64                                                     `lb:1,ub:4,e_lb:1,e_ub:maxNrofPhysicalResourceBlocks_1,optional`
	resourceAllocation                              PUSCH_Config_resourceAllocation                             `madatory`
	pusch_TimeDomainAllocationList                  *PUSCH_TimeDomainResourceAllocationList                     `optional,setuprelease`
	pusch_AggregationFactor                         *PUSCH_Config_pusch_AggregationFactor                       `optional`
	mcs_Table                                       *PUSCH_Config_mcs_Table                                     `optional`
	mcs_TableTransformPrecoder                      *PUSCH_Config_mcs_TableTransformPrecoder                    `optional`
	transformPrecoder                               *PUSCH_Config_transformPrecoder                             `optional`
	codebookSubset                                  *PUSCH_Config_codebookSubset                                `optional`
	maxRank                                         *int64                                                      `lb:1,ub:4,optional`
	rbg_Size                                        *PUSCH_Config_rbg_Size                                      `optional`
	uci_OnPUSCH                                     *UCI_OnPUSCH                                                `optional,setuprelease`
	tp_pi2BPSK                                      *PUSCH_Config_tp_pi2BPSK                                    `optional`
	minimumSchedulingOffsetK2_r16                   *MinSchedulingOffsetK2_Values_r16                           `optional,ext-1,setuprelease`
	ul_AccessConfigListDCI_0_1_r16                  *UL_AccessConfigListDCI_0_1_r16                             `optional,ext-1,setuprelease`
	harq_ProcessNumberSizeDCI_0_2_r16               *int64                                                      `lb:0,ub:4,optional,ext-1`
	dmrs_SequenceInitializationDCI_0_2_r16          *PUSCH_Config_dmrs_SequenceInitializationDCI_0_2_r16        `optional,ext-1`
	numberOfBitsForRV_DCI_0_2_r16                   *int64                                                      `lb:0,ub:2,optional,ext-1`
	antennaPortsFieldPresenceDCI_0_2_r16            *PUSCH_Config_antennaPortsFieldPresenceDCI_0_2_r16          `optional,ext-1`
	dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16    *DMRS_UplinkConfig                                          `optional,ext-1,setuprelease`
	dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16    *DMRS_UplinkConfig                                          `optional,ext-1,setuprelease`
	frequencyHoppingDCI_0_2_r16                     *PUSCH_Config_frequencyHoppingDCI_0_2_r16                   `optional,ext-1`
	frequencyHoppingOffsetListsDCI_0_2_r16          *FrequencyHoppingOffsetListsDCI_0_2_r16                     `optional,ext-1,setuprelease`
	codebookSubsetDCI_0_2_r16                       *PUSCH_Config_codebookSubsetDCI_0_2_r16                     `optional,ext-1`
	invalidSymbolPatternIndicatorDCI_0_2_r16        *PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_2_r16      `optional,ext-1`
	maxRankDCI_0_2_r16                              *int64                                                      `lb:1,ub:4,optional,ext-1`
	mcs_TableDCI_0_2_r16                            *PUSCH_Config_mcs_TableDCI_0_2_r16                          `optional,ext-1`
	mcs_TableTransformPrecoderDCI_0_2_r16           *PUSCH_Config_mcs_TableTransformPrecoderDCI_0_2_r16         `optional,ext-1`
	priorityIndicatorDCI_0_2_r16                    *PUSCH_Config_priorityIndicatorDCI_0_2_r16                  `optional,ext-1`
	pusch_RepTypeIndicatorDCI_0_2_r16               *PUSCH_Config_pusch_RepTypeIndicatorDCI_0_2_r16             `optional,ext-1`
	resourceAllocationDCI_0_2_r16                   *PUSCH_Config_resourceAllocationDCI_0_2_r16                 `optional,ext-1`
	resourceAllocationType1GranularityDCI_0_2_r16   *PUSCH_Config_resourceAllocationType1GranularityDCI_0_2_r16 `optional,ext-1`
	uci_OnPUSCH_ListDCI_0_2_r16                     *UCI_OnPUSCH_ListDCI_0_2_r16                                `optional,ext-1,setuprelease`
	pusch_TimeDomainAllocationListDCI_0_2_r16       *PUSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	pusch_TimeDomainAllocationListDCI_0_1_r16       *PUSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	invalidSymbolPatternIndicatorDCI_0_1_r16        *PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_1_r16      `optional,ext-1`
	priorityIndicatorDCI_0_1_r16                    *PUSCH_Config_priorityIndicatorDCI_0_1_r16                  `optional,ext-1`
	pusch_RepTypeIndicatorDCI_0_1_r16               *PUSCH_Config_pusch_RepTypeIndicatorDCI_0_1_r16             `optional,ext-1`
	frequencyHoppingDCI_0_1_r16                     *PUSCH_Config_frequencyHoppingDCI_0_1_r16                   `optional,ext-1`
	uci_OnPUSCH_ListDCI_0_1_r16                     *UCI_OnPUSCH_ListDCI_0_1_r16                                `optional,ext-1,setuprelease`
	invalidSymbolPattern_r16                        *InvalidSymbolPattern_r16                                   `optional,ext-1`
	pusch_PowerControl_v1610                        *PUSCH_PowerControl_v1610                                   `optional,ext-1,setuprelease`
	ul_FullPowerTransmission_r16                    *PUSCH_Config_ul_FullPowerTransmission_r16                  `optional,ext-1`
	pusch_TimeDomainAllocationListForMultiPUSCH_r16 *PUSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	numberOfInvalidSymbolsForDL_UL_Switching_r16    *int64                                                      `lb:1,ub:4,optional,ext-1`
	ul_AccessConfigListDCI_0_2_r17                  *UL_AccessConfigListDCI_0_2_r17                             `optional,ext-2,setuprelease`
	betaOffsetsCrossPri0_r17                        *BetaOffsetsCrossPriSel_r17                                 `optional,ext-2,setuprelease`
	betaOffsetsCrossPri1_r17                        *BetaOffsetsCrossPriSel_r17                                 `optional,ext-2,setuprelease`
	betaOffsetsCrossPri0DCI_0_2_r17                 *BetaOffsetsCrossPriSelDCI_0_2_r17                          `optional,ext-2,setuprelease`
	betaOffsetsCrossPri1DCI_0_2_r17                 *BetaOffsetsCrossPriSelDCI_0_2_r17                          `optional,ext-2,setuprelease`
	mappingPattern_r17                              *PUSCH_Config_mappingPattern_r17                            `optional,ext-2`
	secondTPCFieldDCI_0_1_r17                       *PUSCH_Config_secondTPCFieldDCI_0_1_r17                     `optional,ext-2`
	secondTPCFieldDCI_0_2_r17                       *PUSCH_Config_secondTPCFieldDCI_0_2_r17                     `optional,ext-2`
	sequenceOffsetForRV_r17                         *int64                                                      `lb:0,ub:3,optional,ext-2`
	ul_AccessConfigListDCI_0_1_r17                  *UL_AccessConfigListDCI_0_1_r17                             `optional,ext-2,setuprelease`
	minimumSchedulingOffsetK2_r17                   *MinSchedulingOffsetK2_Values_r17                           `optional,ext-2,setuprelease`
	availableSlotCounting_r17                       *PUSCH_Config_availableSlotCounting_r17                     `optional,ext-2`
	dmrs_BundlingPUSCH_Config_r17                   *DMRS_BundlingPUSCH_Config_r17                              `optional,ext-2,setuprelease`
	harq_ProcessNumberSizeDCI_0_2_v1700             *int64                                                      `lb:5,ub:5,optional,ext-2`
	harq_ProcessNumberSizeDCI_0_1_r17               *int64                                                      `lb:5,ub:5,optional,ext-2`
	mpe_ResourcePoolToAddModList_r17                []MPE_Resource_r17                                          `lb:1,ub:maxMPE_Resources_r17,optional,ext-2`
	mpe_ResourcePoolToReleaseList_r17               []MPE_ResourceId_r17                                        `lb:1,ub:maxMPE_Resources_r17,optional,ext-2`
}

func (ie *PUSCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.minimumSchedulingOffsetK2_r16 != nil || ie.ul_AccessConfigListDCI_0_1_r16 != nil || ie.harq_ProcessNumberSizeDCI_0_2_r16 != nil || ie.dmrs_SequenceInitializationDCI_0_2_r16 != nil || ie.numberOfBitsForRV_DCI_0_2_r16 != nil || ie.antennaPortsFieldPresenceDCI_0_2_r16 != nil || ie.dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil || ie.dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil || ie.frequencyHoppingDCI_0_2_r16 != nil || ie.frequencyHoppingOffsetListsDCI_0_2_r16 != nil || ie.codebookSubsetDCI_0_2_r16 != nil || ie.invalidSymbolPatternIndicatorDCI_0_2_r16 != nil || ie.maxRankDCI_0_2_r16 != nil || ie.mcs_TableDCI_0_2_r16 != nil || ie.mcs_TableTransformPrecoderDCI_0_2_r16 != nil || ie.priorityIndicatorDCI_0_2_r16 != nil || ie.pusch_RepTypeIndicatorDCI_0_2_r16 != nil || ie.resourceAllocationDCI_0_2_r16 != nil || ie.resourceAllocationType1GranularityDCI_0_2_r16 != nil || ie.uci_OnPUSCH_ListDCI_0_2_r16 != nil || ie.pusch_TimeDomainAllocationListDCI_0_2_r16 != nil || ie.pusch_TimeDomainAllocationListDCI_0_1_r16 != nil || ie.invalidSymbolPatternIndicatorDCI_0_1_r16 != nil || ie.priorityIndicatorDCI_0_1_r16 != nil || ie.pusch_RepTypeIndicatorDCI_0_1_r16 != nil || ie.frequencyHoppingDCI_0_1_r16 != nil || ie.uci_OnPUSCH_ListDCI_0_1_r16 != nil || ie.invalidSymbolPattern_r16 != nil || ie.pusch_PowerControl_v1610 != nil || ie.ul_FullPowerTransmission_r16 != nil || ie.pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil || ie.numberOfInvalidSymbolsForDL_UL_Switching_r16 != nil || ie.ul_AccessConfigListDCI_0_2_r17 != nil || ie.betaOffsetsCrossPri0_r17 != nil || ie.betaOffsetsCrossPri1_r17 != nil || ie.betaOffsetsCrossPri0DCI_0_2_r17 != nil || ie.betaOffsetsCrossPri1DCI_0_2_r17 != nil || ie.mappingPattern_r17 != nil || ie.secondTPCFieldDCI_0_1_r17 != nil || ie.secondTPCFieldDCI_0_2_r17 != nil || ie.sequenceOffsetForRV_r17 != nil || ie.ul_AccessConfigListDCI_0_1_r17 != nil || ie.minimumSchedulingOffsetK2_r17 != nil || ie.availableSlotCounting_r17 != nil || ie.dmrs_BundlingPUSCH_Config_r17 != nil || ie.harq_ProcessNumberSizeDCI_0_2_v1700 != nil || ie.harq_ProcessNumberSizeDCI_0_1_r17 != nil || len(ie.mpe_ResourcePoolToAddModList_r17) > 0 || len(ie.mpe_ResourcePoolToReleaseList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.dataScramblingIdentityPUSCH != nil, ie.txConfig != nil, ie.dmrs_UplinkForPUSCH_MappingTypeA != nil, ie.dmrs_UplinkForPUSCH_MappingTypeB != nil, ie.pusch_PowerControl != nil, ie.frequencyHopping != nil, len(ie.frequencyHoppingOffsetLists) > 0, ie.pusch_TimeDomainAllocationList != nil, ie.pusch_AggregationFactor != nil, ie.mcs_Table != nil, ie.mcs_TableTransformPrecoder != nil, ie.transformPrecoder != nil, ie.codebookSubset != nil, ie.maxRank != nil, ie.rbg_Size != nil, ie.uci_OnPUSCH != nil, ie.tp_pi2BPSK != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dataScramblingIdentityPUSCH != nil {
		if err = w.WriteInteger(*ie.dataScramblingIdentityPUSCH, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode dataScramblingIdentityPUSCH", err)
		}
	}
	if ie.txConfig != nil {
		if err = ie.txConfig.Encode(w); err != nil {
			return utils.WrapError("Encode txConfig", err)
		}
	}
	if ie.dmrs_UplinkForPUSCH_MappingTypeA != nil {
		tmp_dmrs_UplinkForPUSCH_MappingTypeA := utils.SetupRelease[*DMRS_UplinkConfig]{
			Setup: ie.dmrs_UplinkForPUSCH_MappingTypeA,
		}
		if err = tmp_dmrs_UplinkForPUSCH_MappingTypeA.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_UplinkForPUSCH_MappingTypeA", err)
		}
	}
	if ie.dmrs_UplinkForPUSCH_MappingTypeB != nil {
		tmp_dmrs_UplinkForPUSCH_MappingTypeB := utils.SetupRelease[*DMRS_UplinkConfig]{
			Setup: ie.dmrs_UplinkForPUSCH_MappingTypeB,
		}
		if err = tmp_dmrs_UplinkForPUSCH_MappingTypeB.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_UplinkForPUSCH_MappingTypeB", err)
		}
	}
	if ie.pusch_PowerControl != nil {
		if err = ie.pusch_PowerControl.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_PowerControl", err)
		}
	}
	if ie.frequencyHopping != nil {
		if err = ie.frequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyHopping", err)
		}
	}
	if len(ie.frequencyHoppingOffsetLists) > 0 {
		tmp_frequencyHoppingOffsetLists := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.frequencyHoppingOffsetLists {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false)
			tmp_frequencyHoppingOffsetLists.Value = append(tmp_frequencyHoppingOffsetLists.Value, &tmp_ie)
		}
		if err = tmp_frequencyHoppingOffsetLists.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyHoppingOffsetLists", err)
		}
	}
	if err = ie.resourceAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode resourceAllocation", err)
	}
	if ie.pusch_TimeDomainAllocationList != nil {
		tmp_pusch_TimeDomainAllocationList := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList]{
			Setup: ie.pusch_TimeDomainAllocationList,
		}
		if err = tmp_pusch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_TimeDomainAllocationList", err)
		}
	}
	if ie.pusch_AggregationFactor != nil {
		if err = ie.pusch_AggregationFactor.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_AggregationFactor", err)
		}
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
	if ie.transformPrecoder != nil {
		if err = ie.transformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode transformPrecoder", err)
		}
	}
	if ie.codebookSubset != nil {
		if err = ie.codebookSubset.Encode(w); err != nil {
			return utils.WrapError("Encode codebookSubset", err)
		}
	}
	if ie.maxRank != nil {
		if err = w.WriteInteger(*ie.maxRank, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode maxRank", err)
		}
	}
	if ie.rbg_Size != nil {
		if err = ie.rbg_Size.Encode(w); err != nil {
			return utils.WrapError("Encode rbg_Size", err)
		}
	}
	if ie.uci_OnPUSCH != nil {
		tmp_uci_OnPUSCH := utils.SetupRelease[*UCI_OnPUSCH]{
			Setup: ie.uci_OnPUSCH,
		}
		if err = tmp_uci_OnPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode uci_OnPUSCH", err)
		}
	}
	if ie.tp_pi2BPSK != nil {
		if err = ie.tp_pi2BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode tp_pi2BPSK", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.minimumSchedulingOffsetK2_r16 != nil || ie.ul_AccessConfigListDCI_0_1_r16 != nil || ie.harq_ProcessNumberSizeDCI_0_2_r16 != nil || ie.dmrs_SequenceInitializationDCI_0_2_r16 != nil || ie.numberOfBitsForRV_DCI_0_2_r16 != nil || ie.antennaPortsFieldPresenceDCI_0_2_r16 != nil || ie.dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil || ie.dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil || ie.frequencyHoppingDCI_0_2_r16 != nil || ie.frequencyHoppingOffsetListsDCI_0_2_r16 != nil || ie.codebookSubsetDCI_0_2_r16 != nil || ie.invalidSymbolPatternIndicatorDCI_0_2_r16 != nil || ie.maxRankDCI_0_2_r16 != nil || ie.mcs_TableDCI_0_2_r16 != nil || ie.mcs_TableTransformPrecoderDCI_0_2_r16 != nil || ie.priorityIndicatorDCI_0_2_r16 != nil || ie.pusch_RepTypeIndicatorDCI_0_2_r16 != nil || ie.resourceAllocationDCI_0_2_r16 != nil || ie.resourceAllocationType1GranularityDCI_0_2_r16 != nil || ie.uci_OnPUSCH_ListDCI_0_2_r16 != nil || ie.pusch_TimeDomainAllocationListDCI_0_2_r16 != nil || ie.pusch_TimeDomainAllocationListDCI_0_1_r16 != nil || ie.invalidSymbolPatternIndicatorDCI_0_1_r16 != nil || ie.priorityIndicatorDCI_0_1_r16 != nil || ie.pusch_RepTypeIndicatorDCI_0_1_r16 != nil || ie.frequencyHoppingDCI_0_1_r16 != nil || ie.uci_OnPUSCH_ListDCI_0_1_r16 != nil || ie.invalidSymbolPattern_r16 != nil || ie.pusch_PowerControl_v1610 != nil || ie.ul_FullPowerTransmission_r16 != nil || ie.pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil || ie.numberOfInvalidSymbolsForDL_UL_Switching_r16 != nil, ie.ul_AccessConfigListDCI_0_2_r17 != nil || ie.betaOffsetsCrossPri0_r17 != nil || ie.betaOffsetsCrossPri1_r17 != nil || ie.betaOffsetsCrossPri0DCI_0_2_r17 != nil || ie.betaOffsetsCrossPri1DCI_0_2_r17 != nil || ie.mappingPattern_r17 != nil || ie.secondTPCFieldDCI_0_1_r17 != nil || ie.secondTPCFieldDCI_0_2_r17 != nil || ie.sequenceOffsetForRV_r17 != nil || ie.ul_AccessConfigListDCI_0_1_r17 != nil || ie.minimumSchedulingOffsetK2_r17 != nil || ie.availableSlotCounting_r17 != nil || ie.dmrs_BundlingPUSCH_Config_r17 != nil || ie.harq_ProcessNumberSizeDCI_0_2_v1700 != nil || ie.harq_ProcessNumberSizeDCI_0_1_r17 != nil || len(ie.mpe_ResourcePoolToAddModList_r17) > 0 || len(ie.mpe_ResourcePoolToReleaseList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.minimumSchedulingOffsetK2_r16 != nil, ie.ul_AccessConfigListDCI_0_1_r16 != nil, ie.harq_ProcessNumberSizeDCI_0_2_r16 != nil, ie.dmrs_SequenceInitializationDCI_0_2_r16 != nil, ie.numberOfBitsForRV_DCI_0_2_r16 != nil, ie.antennaPortsFieldPresenceDCI_0_2_r16 != nil, ie.dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil, ie.dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil, ie.frequencyHoppingDCI_0_2_r16 != nil, ie.frequencyHoppingOffsetListsDCI_0_2_r16 != nil, ie.codebookSubsetDCI_0_2_r16 != nil, ie.invalidSymbolPatternIndicatorDCI_0_2_r16 != nil, ie.maxRankDCI_0_2_r16 != nil, ie.mcs_TableDCI_0_2_r16 != nil, ie.mcs_TableTransformPrecoderDCI_0_2_r16 != nil, ie.priorityIndicatorDCI_0_2_r16 != nil, ie.pusch_RepTypeIndicatorDCI_0_2_r16 != nil, ie.resourceAllocationDCI_0_2_r16 != nil, ie.resourceAllocationType1GranularityDCI_0_2_r16 != nil, ie.uci_OnPUSCH_ListDCI_0_2_r16 != nil, ie.pusch_TimeDomainAllocationListDCI_0_2_r16 != nil, ie.pusch_TimeDomainAllocationListDCI_0_1_r16 != nil, ie.invalidSymbolPatternIndicatorDCI_0_1_r16 != nil, ie.priorityIndicatorDCI_0_1_r16 != nil, ie.pusch_RepTypeIndicatorDCI_0_1_r16 != nil, ie.frequencyHoppingDCI_0_1_r16 != nil, ie.uci_OnPUSCH_ListDCI_0_1_r16 != nil, ie.invalidSymbolPattern_r16 != nil, ie.pusch_PowerControl_v1610 != nil, ie.ul_FullPowerTransmission_r16 != nil, ie.pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil, ie.numberOfInvalidSymbolsForDL_UL_Switching_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode minimumSchedulingOffsetK2_r16 optional
			if ie.minimumSchedulingOffsetK2_r16 != nil {
				tmp_minimumSchedulingOffsetK2_r16 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r16]{
					Setup: ie.minimumSchedulingOffsetK2_r16,
				}
				if err = tmp_minimumSchedulingOffsetK2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode minimumSchedulingOffsetK2_r16", err)
				}
			}
			// encode ul_AccessConfigListDCI_0_1_r16 optional
			if ie.ul_AccessConfigListDCI_0_1_r16 != nil {
				tmp_ul_AccessConfigListDCI_0_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r16]{
					Setup: ie.ul_AccessConfigListDCI_0_1_r16,
				}
				if err = tmp_ul_AccessConfigListDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_AccessConfigListDCI_0_1_r16", err)
				}
			}
			// encode harq_ProcessNumberSizeDCI_0_2_r16 optional
			if ie.harq_ProcessNumberSizeDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcessNumberSizeDCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode harq_ProcessNumberSizeDCI_0_2_r16", err)
				}
			}
			// encode dmrs_SequenceInitializationDCI_0_2_r16 optional
			if ie.dmrs_SequenceInitializationDCI_0_2_r16 != nil {
				if err = ie.dmrs_SequenceInitializationDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_SequenceInitializationDCI_0_2_r16", err)
				}
			}
			// encode numberOfBitsForRV_DCI_0_2_r16 optional
			if ie.numberOfBitsForRV_DCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.numberOfBitsForRV_DCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode numberOfBitsForRV_DCI_0_2_r16", err)
				}
			}
			// encode antennaPortsFieldPresenceDCI_0_2_r16 optional
			if ie.antennaPortsFieldPresenceDCI_0_2_r16 != nil {
				if err = ie.antennaPortsFieldPresenceDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode antennaPortsFieldPresenceDCI_0_2_r16", err)
				}
			}
			// encode dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 optional
			if ie.dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil {
				tmp_dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{
					Setup: ie.dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16,
				}
				if err = tmp_dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16", err)
				}
			}
			// encode dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 optional
			if ie.dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil {
				tmp_dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{
					Setup: ie.dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16,
				}
				if err = tmp_dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16", err)
				}
			}
			// encode frequencyHoppingDCI_0_2_r16 optional
			if ie.frequencyHoppingDCI_0_2_r16 != nil {
				if err = ie.frequencyHoppingDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode frequencyHoppingDCI_0_2_r16", err)
				}
			}
			// encode frequencyHoppingOffsetListsDCI_0_2_r16 optional
			if ie.frequencyHoppingOffsetListsDCI_0_2_r16 != nil {
				tmp_frequencyHoppingOffsetListsDCI_0_2_r16 := utils.SetupRelease[*FrequencyHoppingOffsetListsDCI_0_2_r16]{
					Setup: ie.frequencyHoppingOffsetListsDCI_0_2_r16,
				}
				if err = tmp_frequencyHoppingOffsetListsDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode frequencyHoppingOffsetListsDCI_0_2_r16", err)
				}
			}
			// encode codebookSubsetDCI_0_2_r16 optional
			if ie.codebookSubsetDCI_0_2_r16 != nil {
				if err = ie.codebookSubsetDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookSubsetDCI_0_2_r16", err)
				}
			}
			// encode invalidSymbolPatternIndicatorDCI_0_2_r16 optional
			if ie.invalidSymbolPatternIndicatorDCI_0_2_r16 != nil {
				if err = ie.invalidSymbolPatternIndicatorDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode invalidSymbolPatternIndicatorDCI_0_2_r16", err)
				}
			}
			// encode maxRankDCI_0_2_r16 optional
			if ie.maxRankDCI_0_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.maxRankDCI_0_2_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode maxRankDCI_0_2_r16", err)
				}
			}
			// encode mcs_TableDCI_0_2_r16 optional
			if ie.mcs_TableDCI_0_2_r16 != nil {
				if err = ie.mcs_TableDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_TableDCI_0_2_r16", err)
				}
			}
			// encode mcs_TableTransformPrecoderDCI_0_2_r16 optional
			if ie.mcs_TableTransformPrecoderDCI_0_2_r16 != nil {
				if err = ie.mcs_TableTransformPrecoderDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_TableTransformPrecoderDCI_0_2_r16", err)
				}
			}
			// encode priorityIndicatorDCI_0_2_r16 optional
			if ie.priorityIndicatorDCI_0_2_r16 != nil {
				if err = ie.priorityIndicatorDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorDCI_0_2_r16", err)
				}
			}
			// encode pusch_RepTypeIndicatorDCI_0_2_r16 optional
			if ie.pusch_RepTypeIndicatorDCI_0_2_r16 != nil {
				if err = ie.pusch_RepTypeIndicatorDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_RepTypeIndicatorDCI_0_2_r16", err)
				}
			}
			// encode resourceAllocationDCI_0_2_r16 optional
			if ie.resourceAllocationDCI_0_2_r16 != nil {
				if err = ie.resourceAllocationDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceAllocationDCI_0_2_r16", err)
				}
			}
			// encode resourceAllocationType1GranularityDCI_0_2_r16 optional
			if ie.resourceAllocationType1GranularityDCI_0_2_r16 != nil {
				if err = ie.resourceAllocationType1GranularityDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceAllocationType1GranularityDCI_0_2_r16", err)
				}
			}
			// encode uci_OnPUSCH_ListDCI_0_2_r16 optional
			if ie.uci_OnPUSCH_ListDCI_0_2_r16 != nil {
				tmp_uci_OnPUSCH_ListDCI_0_2_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_2_r16]{
					Setup: ie.uci_OnPUSCH_ListDCI_0_2_r16,
				}
				if err = tmp_uci_OnPUSCH_ListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uci_OnPUSCH_ListDCI_0_2_r16", err)
				}
			}
			// encode pusch_TimeDomainAllocationListDCI_0_2_r16 optional
			if ie.pusch_TimeDomainAllocationListDCI_0_2_r16 != nil {
				tmp_pusch_TimeDomainAllocationListDCI_0_2_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.pusch_TimeDomainAllocationListDCI_0_2_r16,
				}
				if err = tmp_pusch_TimeDomainAllocationListDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_TimeDomainAllocationListDCI_0_2_r16", err)
				}
			}
			// encode pusch_TimeDomainAllocationListDCI_0_1_r16 optional
			if ie.pusch_TimeDomainAllocationListDCI_0_1_r16 != nil {
				tmp_pusch_TimeDomainAllocationListDCI_0_1_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.pusch_TimeDomainAllocationListDCI_0_1_r16,
				}
				if err = tmp_pusch_TimeDomainAllocationListDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_TimeDomainAllocationListDCI_0_1_r16", err)
				}
			}
			// encode invalidSymbolPatternIndicatorDCI_0_1_r16 optional
			if ie.invalidSymbolPatternIndicatorDCI_0_1_r16 != nil {
				if err = ie.invalidSymbolPatternIndicatorDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode invalidSymbolPatternIndicatorDCI_0_1_r16", err)
				}
			}
			// encode priorityIndicatorDCI_0_1_r16 optional
			if ie.priorityIndicatorDCI_0_1_r16 != nil {
				if err = ie.priorityIndicatorDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorDCI_0_1_r16", err)
				}
			}
			// encode pusch_RepTypeIndicatorDCI_0_1_r16 optional
			if ie.pusch_RepTypeIndicatorDCI_0_1_r16 != nil {
				if err = ie.pusch_RepTypeIndicatorDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_RepTypeIndicatorDCI_0_1_r16", err)
				}
			}
			// encode frequencyHoppingDCI_0_1_r16 optional
			if ie.frequencyHoppingDCI_0_1_r16 != nil {
				if err = ie.frequencyHoppingDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode frequencyHoppingDCI_0_1_r16", err)
				}
			}
			// encode uci_OnPUSCH_ListDCI_0_1_r16 optional
			if ie.uci_OnPUSCH_ListDCI_0_1_r16 != nil {
				tmp_uci_OnPUSCH_ListDCI_0_1_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_1_r16]{
					Setup: ie.uci_OnPUSCH_ListDCI_0_1_r16,
				}
				if err = tmp_uci_OnPUSCH_ListDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uci_OnPUSCH_ListDCI_0_1_r16", err)
				}
			}
			// encode invalidSymbolPattern_r16 optional
			if ie.invalidSymbolPattern_r16 != nil {
				if err = ie.invalidSymbolPattern_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode invalidSymbolPattern_r16", err)
				}
			}
			// encode pusch_PowerControl_v1610 optional
			if ie.pusch_PowerControl_v1610 != nil {
				tmp_pusch_PowerControl_v1610 := utils.SetupRelease[*PUSCH_PowerControl_v1610]{
					Setup: ie.pusch_PowerControl_v1610,
				}
				if err = tmp_pusch_PowerControl_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_PowerControl_v1610", err)
				}
			}
			// encode ul_FullPowerTransmission_r16 optional
			if ie.ul_FullPowerTransmission_r16 != nil {
				if err = ie.ul_FullPowerTransmission_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_FullPowerTransmission_r16", err)
				}
			}
			// encode pusch_TimeDomainAllocationListForMultiPUSCH_r16 optional
			if ie.pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil {
				tmp_pusch_TimeDomainAllocationListForMultiPUSCH_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.pusch_TimeDomainAllocationListForMultiPUSCH_r16,
				}
				if err = tmp_pusch_TimeDomainAllocationListForMultiPUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_TimeDomainAllocationListForMultiPUSCH_r16", err)
				}
			}
			// encode numberOfInvalidSymbolsForDL_UL_Switching_r16 optional
			if ie.numberOfInvalidSymbolsForDL_UL_Switching_r16 != nil {
				if err = extWriter.WriteInteger(*ie.numberOfInvalidSymbolsForDL_UL_Switching_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode numberOfInvalidSymbolsForDL_UL_Switching_r16", err)
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
			optionals_ext_2 := []bool{ie.ul_AccessConfigListDCI_0_2_r17 != nil, ie.betaOffsetsCrossPri0_r17 != nil, ie.betaOffsetsCrossPri1_r17 != nil, ie.betaOffsetsCrossPri0DCI_0_2_r17 != nil, ie.betaOffsetsCrossPri1DCI_0_2_r17 != nil, ie.mappingPattern_r17 != nil, ie.secondTPCFieldDCI_0_1_r17 != nil, ie.secondTPCFieldDCI_0_2_r17 != nil, ie.sequenceOffsetForRV_r17 != nil, ie.ul_AccessConfigListDCI_0_1_r17 != nil, ie.minimumSchedulingOffsetK2_r17 != nil, ie.availableSlotCounting_r17 != nil, ie.dmrs_BundlingPUSCH_Config_r17 != nil, ie.harq_ProcessNumberSizeDCI_0_2_v1700 != nil, ie.harq_ProcessNumberSizeDCI_0_1_r17 != nil, len(ie.mpe_ResourcePoolToAddModList_r17) > 0, len(ie.mpe_ResourcePoolToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ul_AccessConfigListDCI_0_2_r17 optional
			if ie.ul_AccessConfigListDCI_0_2_r17 != nil {
				tmp_ul_AccessConfigListDCI_0_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_2_r17]{
					Setup: ie.ul_AccessConfigListDCI_0_2_r17,
				}
				if err = tmp_ul_AccessConfigListDCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_AccessConfigListDCI_0_2_r17", err)
				}
			}
			// encode betaOffsetsCrossPri0_r17 optional
			if ie.betaOffsetsCrossPri0_r17 != nil {
				tmp_betaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{
					Setup: ie.betaOffsetsCrossPri0_r17,
				}
				if err = tmp_betaOffsetsCrossPri0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode betaOffsetsCrossPri0_r17", err)
				}
			}
			// encode betaOffsetsCrossPri1_r17 optional
			if ie.betaOffsetsCrossPri1_r17 != nil {
				tmp_betaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{
					Setup: ie.betaOffsetsCrossPri1_r17,
				}
				if err = tmp_betaOffsetsCrossPri1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode betaOffsetsCrossPri1_r17", err)
				}
			}
			// encode betaOffsetsCrossPri0DCI_0_2_r17 optional
			if ie.betaOffsetsCrossPri0DCI_0_2_r17 != nil {
				tmp_betaOffsetsCrossPri0DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{
					Setup: ie.betaOffsetsCrossPri0DCI_0_2_r17,
				}
				if err = tmp_betaOffsetsCrossPri0DCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode betaOffsetsCrossPri0DCI_0_2_r17", err)
				}
			}
			// encode betaOffsetsCrossPri1DCI_0_2_r17 optional
			if ie.betaOffsetsCrossPri1DCI_0_2_r17 != nil {
				tmp_betaOffsetsCrossPri1DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{
					Setup: ie.betaOffsetsCrossPri1DCI_0_2_r17,
				}
				if err = tmp_betaOffsetsCrossPri1DCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode betaOffsetsCrossPri1DCI_0_2_r17", err)
				}
			}
			// encode mappingPattern_r17 optional
			if ie.mappingPattern_r17 != nil {
				if err = ie.mappingPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mappingPattern_r17", err)
				}
			}
			// encode secondTPCFieldDCI_0_1_r17 optional
			if ie.secondTPCFieldDCI_0_1_r17 != nil {
				if err = ie.secondTPCFieldDCI_0_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode secondTPCFieldDCI_0_1_r17", err)
				}
			}
			// encode secondTPCFieldDCI_0_2_r17 optional
			if ie.secondTPCFieldDCI_0_2_r17 != nil {
				if err = ie.secondTPCFieldDCI_0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode secondTPCFieldDCI_0_2_r17", err)
				}
			}
			// encode sequenceOffsetForRV_r17 optional
			if ie.sequenceOffsetForRV_r17 != nil {
				if err = extWriter.WriteInteger(*ie.sequenceOffsetForRV_r17, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode sequenceOffsetForRV_r17", err)
				}
			}
			// encode ul_AccessConfigListDCI_0_1_r17 optional
			if ie.ul_AccessConfigListDCI_0_1_r17 != nil {
				tmp_ul_AccessConfigListDCI_0_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r17]{
					Setup: ie.ul_AccessConfigListDCI_0_1_r17,
				}
				if err = tmp_ul_AccessConfigListDCI_0_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_AccessConfigListDCI_0_1_r17", err)
				}
			}
			// encode minimumSchedulingOffsetK2_r17 optional
			if ie.minimumSchedulingOffsetK2_r17 != nil {
				tmp_minimumSchedulingOffsetK2_r17 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r17]{
					Setup: ie.minimumSchedulingOffsetK2_r17,
				}
				if err = tmp_minimumSchedulingOffsetK2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode minimumSchedulingOffsetK2_r17", err)
				}
			}
			// encode availableSlotCounting_r17 optional
			if ie.availableSlotCounting_r17 != nil {
				if err = ie.availableSlotCounting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availableSlotCounting_r17", err)
				}
			}
			// encode dmrs_BundlingPUSCH_Config_r17 optional
			if ie.dmrs_BundlingPUSCH_Config_r17 != nil {
				tmp_dmrs_BundlingPUSCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUSCH_Config_r17]{
					Setup: ie.dmrs_BundlingPUSCH_Config_r17,
				}
				if err = tmp_dmrs_BundlingPUSCH_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingPUSCH_Config_r17", err)
				}
			}
			// encode harq_ProcessNumberSizeDCI_0_2_v1700 optional
			if ie.harq_ProcessNumberSizeDCI_0_2_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcessNumberSizeDCI_0_2_v1700, &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode harq_ProcessNumberSizeDCI_0_2_v1700", err)
				}
			}
			// encode harq_ProcessNumberSizeDCI_0_1_r17 optional
			if ie.harq_ProcessNumberSizeDCI_0_1_r17 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcessNumberSizeDCI_0_1_r17, &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode harq_ProcessNumberSizeDCI_0_1_r17", err)
				}
			}
			// encode mpe_ResourcePoolToAddModList_r17 optional
			if len(ie.mpe_ResourcePoolToAddModList_r17) > 0 {
				tmp_mpe_ResourcePoolToAddModList_r17 := utils.NewSequence[*MPE_Resource_r17]([]*MPE_Resource_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				for _, i := range ie.mpe_ResourcePoolToAddModList_r17 {
					tmp_mpe_ResourcePoolToAddModList_r17.Value = append(tmp_mpe_ResourcePoolToAddModList_r17.Value, &i)
				}
				if err = tmp_mpe_ResourcePoolToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpe_ResourcePoolToAddModList_r17", err)
				}
			}
			// encode mpe_ResourcePoolToReleaseList_r17 optional
			if len(ie.mpe_ResourcePoolToReleaseList_r17) > 0 {
				tmp_mpe_ResourcePoolToReleaseList_r17 := utils.NewSequence[*MPE_ResourceId_r17]([]*MPE_ResourceId_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				for _, i := range ie.mpe_ResourcePoolToReleaseList_r17 {
					tmp_mpe_ResourcePoolToReleaseList_r17.Value = append(tmp_mpe_ResourcePoolToReleaseList_r17.Value, &i)
				}
				if err = tmp_mpe_ResourcePoolToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpe_ResourcePoolToReleaseList_r17", err)
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

func (ie *PUSCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dataScramblingIdentityPUSCHPresent bool
	if dataScramblingIdentityPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var txConfigPresent bool
	if txConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_UplinkForPUSCH_MappingTypeAPresent bool
	if dmrs_UplinkForPUSCH_MappingTypeAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_UplinkForPUSCH_MappingTypeBPresent bool
	if dmrs_UplinkForPUSCH_MappingTypeBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_PowerControlPresent bool
	if pusch_PowerControlPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyHoppingPresent bool
	if frequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyHoppingOffsetListsPresent bool
	if frequencyHoppingOffsetListsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_TimeDomainAllocationListPresent bool
	if pusch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_AggregationFactorPresent bool
	if pusch_AggregationFactorPresent, err = r.ReadBool(); err != nil {
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
	var transformPrecoderPresent bool
	if transformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var codebookSubsetPresent bool
	if codebookSubsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxRankPresent bool
	if maxRankPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rbg_SizePresent bool
	if rbg_SizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uci_OnPUSCHPresent bool
	if uci_OnPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tp_pi2BPSKPresent bool
	if tp_pi2BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dataScramblingIdentityPUSCHPresent {
		var tmp_int_dataScramblingIdentityPUSCH int64
		if tmp_int_dataScramblingIdentityPUSCH, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode dataScramblingIdentityPUSCH", err)
		}
		ie.dataScramblingIdentityPUSCH = &tmp_int_dataScramblingIdentityPUSCH
	}
	if txConfigPresent {
		ie.txConfig = new(PUSCH_Config_txConfig)
		if err = ie.txConfig.Decode(r); err != nil {
			return utils.WrapError("Decode txConfig", err)
		}
	}
	if dmrs_UplinkForPUSCH_MappingTypeAPresent {
		tmp_dmrs_UplinkForPUSCH_MappingTypeA := utils.SetupRelease[*DMRS_UplinkConfig]{}
		if err = tmp_dmrs_UplinkForPUSCH_MappingTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_UplinkForPUSCH_MappingTypeA", err)
		}
		ie.dmrs_UplinkForPUSCH_MappingTypeA = tmp_dmrs_UplinkForPUSCH_MappingTypeA.Setup
	}
	if dmrs_UplinkForPUSCH_MappingTypeBPresent {
		tmp_dmrs_UplinkForPUSCH_MappingTypeB := utils.SetupRelease[*DMRS_UplinkConfig]{}
		if err = tmp_dmrs_UplinkForPUSCH_MappingTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_UplinkForPUSCH_MappingTypeB", err)
		}
		ie.dmrs_UplinkForPUSCH_MappingTypeB = tmp_dmrs_UplinkForPUSCH_MappingTypeB.Setup
	}
	if pusch_PowerControlPresent {
		ie.pusch_PowerControl = new(PUSCH_PowerControl)
		if err = ie.pusch_PowerControl.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_PowerControl", err)
		}
	}
	if frequencyHoppingPresent {
		ie.frequencyHopping = new(PUSCH_Config_frequencyHopping)
		if err = ie.frequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyHopping", err)
		}
	}
	if frequencyHoppingOffsetListsPresent {
		tmp_frequencyHoppingOffsetLists := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn_frequencyHoppingOffsetLists := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks_1}, false)
			return &ie
		}
		if err = tmp_frequencyHoppingOffsetLists.Decode(r, fn_frequencyHoppingOffsetLists); err != nil {
			return utils.WrapError("Decode frequencyHoppingOffsetLists", err)
		}
		ie.frequencyHoppingOffsetLists = []int64{}
		for _, i := range tmp_frequencyHoppingOffsetLists.Value {
			ie.frequencyHoppingOffsetLists = append(ie.frequencyHoppingOffsetLists, int64(i.Value))
		}
	}
	if err = ie.resourceAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode resourceAllocation", err)
	}
	if pusch_TimeDomainAllocationListPresent {
		tmp_pusch_TimeDomainAllocationList := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList]{}
		if err = tmp_pusch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_TimeDomainAllocationList", err)
		}
		ie.pusch_TimeDomainAllocationList = tmp_pusch_TimeDomainAllocationList.Setup
	}
	if pusch_AggregationFactorPresent {
		ie.pusch_AggregationFactor = new(PUSCH_Config_pusch_AggregationFactor)
		if err = ie.pusch_AggregationFactor.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_AggregationFactor", err)
		}
	}
	if mcs_TablePresent {
		ie.mcs_Table = new(PUSCH_Config_mcs_Table)
		if err = ie.mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_Table", err)
		}
	}
	if mcs_TableTransformPrecoderPresent {
		ie.mcs_TableTransformPrecoder = new(PUSCH_Config_mcs_TableTransformPrecoder)
		if err = ie.mcs_TableTransformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_TableTransformPrecoder", err)
		}
	}
	if transformPrecoderPresent {
		ie.transformPrecoder = new(PUSCH_Config_transformPrecoder)
		if err = ie.transformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode transformPrecoder", err)
		}
	}
	if codebookSubsetPresent {
		ie.codebookSubset = new(PUSCH_Config_codebookSubset)
		if err = ie.codebookSubset.Decode(r); err != nil {
			return utils.WrapError("Decode codebookSubset", err)
		}
	}
	if maxRankPresent {
		var tmp_int_maxRank int64
		if tmp_int_maxRank, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode maxRank", err)
		}
		ie.maxRank = &tmp_int_maxRank
	}
	if rbg_SizePresent {
		ie.rbg_Size = new(PUSCH_Config_rbg_Size)
		if err = ie.rbg_Size.Decode(r); err != nil {
			return utils.WrapError("Decode rbg_Size", err)
		}
	}
	if uci_OnPUSCHPresent {
		tmp_uci_OnPUSCH := utils.SetupRelease[*UCI_OnPUSCH]{}
		if err = tmp_uci_OnPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode uci_OnPUSCH", err)
		}
		ie.uci_OnPUSCH = tmp_uci_OnPUSCH.Setup
	}
	if tp_pi2BPSKPresent {
		ie.tp_pi2BPSK = new(PUSCH_Config_tp_pi2BPSK)
		if err = ie.tp_pi2BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode tp_pi2BPSK", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			minimumSchedulingOffsetK2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_AccessConfigListDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcessNumberSizeDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_SequenceInitializationDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfBitsForRV_DCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			antennaPortsFieldPresenceDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			frequencyHoppingDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			frequencyHoppingOffsetListsDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookSubsetDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			invalidSymbolPatternIndicatorDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxRankDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_TableDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_TableTransformPrecoderDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			priorityIndicatorDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_RepTypeIndicatorDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceAllocationDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceAllocationType1GranularityDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uci_OnPUSCH_ListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_TimeDomainAllocationListDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_TimeDomainAllocationListDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			invalidSymbolPatternIndicatorDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			priorityIndicatorDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_RepTypeIndicatorDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			frequencyHoppingDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uci_OnPUSCH_ListDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			invalidSymbolPattern_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_PowerControl_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_FullPowerTransmission_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_TimeDomainAllocationListForMultiPUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfInvalidSymbolsForDL_UL_Switching_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode minimumSchedulingOffsetK2_r16 optional
			if minimumSchedulingOffsetK2_r16Present {
				tmp_minimumSchedulingOffsetK2_r16 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r16]{}
				if err = tmp_minimumSchedulingOffsetK2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode minimumSchedulingOffsetK2_r16", err)
				}
				ie.minimumSchedulingOffsetK2_r16 = tmp_minimumSchedulingOffsetK2_r16.Setup
			}
			// decode ul_AccessConfigListDCI_0_1_r16 optional
			if ul_AccessConfigListDCI_0_1_r16Present {
				tmp_ul_AccessConfigListDCI_0_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r16]{}
				if err = tmp_ul_AccessConfigListDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_AccessConfigListDCI_0_1_r16", err)
				}
				ie.ul_AccessConfigListDCI_0_1_r16 = tmp_ul_AccessConfigListDCI_0_1_r16.Setup
			}
			// decode harq_ProcessNumberSizeDCI_0_2_r16 optional
			if harq_ProcessNumberSizeDCI_0_2_r16Present {
				var tmp_int_harq_ProcessNumberSizeDCI_0_2_r16 int64
				if tmp_int_harq_ProcessNumberSizeDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode harq_ProcessNumberSizeDCI_0_2_r16", err)
				}
				ie.harq_ProcessNumberSizeDCI_0_2_r16 = &tmp_int_harq_ProcessNumberSizeDCI_0_2_r16
			}
			// decode dmrs_SequenceInitializationDCI_0_2_r16 optional
			if dmrs_SequenceInitializationDCI_0_2_r16Present {
				ie.dmrs_SequenceInitializationDCI_0_2_r16 = new(PUSCH_Config_dmrs_SequenceInitializationDCI_0_2_r16)
				if err = ie.dmrs_SequenceInitializationDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_SequenceInitializationDCI_0_2_r16", err)
				}
			}
			// decode numberOfBitsForRV_DCI_0_2_r16 optional
			if numberOfBitsForRV_DCI_0_2_r16Present {
				var tmp_int_numberOfBitsForRV_DCI_0_2_r16 int64
				if tmp_int_numberOfBitsForRV_DCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode numberOfBitsForRV_DCI_0_2_r16", err)
				}
				ie.numberOfBitsForRV_DCI_0_2_r16 = &tmp_int_numberOfBitsForRV_DCI_0_2_r16
			}
			// decode antennaPortsFieldPresenceDCI_0_2_r16 optional
			if antennaPortsFieldPresenceDCI_0_2_r16Present {
				ie.antennaPortsFieldPresenceDCI_0_2_r16 = new(PUSCH_Config_antennaPortsFieldPresenceDCI_0_2_r16)
				if err = ie.antennaPortsFieldPresenceDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode antennaPortsFieldPresenceDCI_0_2_r16", err)
				}
			}
			// decode dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 optional
			if dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16Present {
				tmp_dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{}
				if err = tmp_dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16", err)
				}
				ie.dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 = tmp_dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16.Setup
			}
			// decode dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 optional
			if dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16Present {
				tmp_dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 := utils.SetupRelease[*DMRS_UplinkConfig]{}
				if err = tmp_dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16", err)
				}
				ie.dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 = tmp_dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16.Setup
			}
			// decode frequencyHoppingDCI_0_2_r16 optional
			if frequencyHoppingDCI_0_2_r16Present {
				ie.frequencyHoppingDCI_0_2_r16 = new(PUSCH_Config_frequencyHoppingDCI_0_2_r16)
				if err = ie.frequencyHoppingDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode frequencyHoppingDCI_0_2_r16", err)
				}
			}
			// decode frequencyHoppingOffsetListsDCI_0_2_r16 optional
			if frequencyHoppingOffsetListsDCI_0_2_r16Present {
				tmp_frequencyHoppingOffsetListsDCI_0_2_r16 := utils.SetupRelease[*FrequencyHoppingOffsetListsDCI_0_2_r16]{}
				if err = tmp_frequencyHoppingOffsetListsDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode frequencyHoppingOffsetListsDCI_0_2_r16", err)
				}
				ie.frequencyHoppingOffsetListsDCI_0_2_r16 = tmp_frequencyHoppingOffsetListsDCI_0_2_r16.Setup
			}
			// decode codebookSubsetDCI_0_2_r16 optional
			if codebookSubsetDCI_0_2_r16Present {
				ie.codebookSubsetDCI_0_2_r16 = new(PUSCH_Config_codebookSubsetDCI_0_2_r16)
				if err = ie.codebookSubsetDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookSubsetDCI_0_2_r16", err)
				}
			}
			// decode invalidSymbolPatternIndicatorDCI_0_2_r16 optional
			if invalidSymbolPatternIndicatorDCI_0_2_r16Present {
				ie.invalidSymbolPatternIndicatorDCI_0_2_r16 = new(PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_2_r16)
				if err = ie.invalidSymbolPatternIndicatorDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode invalidSymbolPatternIndicatorDCI_0_2_r16", err)
				}
			}
			// decode maxRankDCI_0_2_r16 optional
			if maxRankDCI_0_2_r16Present {
				var tmp_int_maxRankDCI_0_2_r16 int64
				if tmp_int_maxRankDCI_0_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode maxRankDCI_0_2_r16", err)
				}
				ie.maxRankDCI_0_2_r16 = &tmp_int_maxRankDCI_0_2_r16
			}
			// decode mcs_TableDCI_0_2_r16 optional
			if mcs_TableDCI_0_2_r16Present {
				ie.mcs_TableDCI_0_2_r16 = new(PUSCH_Config_mcs_TableDCI_0_2_r16)
				if err = ie.mcs_TableDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mcs_TableDCI_0_2_r16", err)
				}
			}
			// decode mcs_TableTransformPrecoderDCI_0_2_r16 optional
			if mcs_TableTransformPrecoderDCI_0_2_r16Present {
				ie.mcs_TableTransformPrecoderDCI_0_2_r16 = new(PUSCH_Config_mcs_TableTransformPrecoderDCI_0_2_r16)
				if err = ie.mcs_TableTransformPrecoderDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mcs_TableTransformPrecoderDCI_0_2_r16", err)
				}
			}
			// decode priorityIndicatorDCI_0_2_r16 optional
			if priorityIndicatorDCI_0_2_r16Present {
				ie.priorityIndicatorDCI_0_2_r16 = new(PUSCH_Config_priorityIndicatorDCI_0_2_r16)
				if err = ie.priorityIndicatorDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorDCI_0_2_r16", err)
				}
			}
			// decode pusch_RepTypeIndicatorDCI_0_2_r16 optional
			if pusch_RepTypeIndicatorDCI_0_2_r16Present {
				ie.pusch_RepTypeIndicatorDCI_0_2_r16 = new(PUSCH_Config_pusch_RepTypeIndicatorDCI_0_2_r16)
				if err = ie.pusch_RepTypeIndicatorDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_RepTypeIndicatorDCI_0_2_r16", err)
				}
			}
			// decode resourceAllocationDCI_0_2_r16 optional
			if resourceAllocationDCI_0_2_r16Present {
				ie.resourceAllocationDCI_0_2_r16 = new(PUSCH_Config_resourceAllocationDCI_0_2_r16)
				if err = ie.resourceAllocationDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourceAllocationDCI_0_2_r16", err)
				}
			}
			// decode resourceAllocationType1GranularityDCI_0_2_r16 optional
			if resourceAllocationType1GranularityDCI_0_2_r16Present {
				ie.resourceAllocationType1GranularityDCI_0_2_r16 = new(PUSCH_Config_resourceAllocationType1GranularityDCI_0_2_r16)
				if err = ie.resourceAllocationType1GranularityDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourceAllocationType1GranularityDCI_0_2_r16", err)
				}
			}
			// decode uci_OnPUSCH_ListDCI_0_2_r16 optional
			if uci_OnPUSCH_ListDCI_0_2_r16Present {
				tmp_uci_OnPUSCH_ListDCI_0_2_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_2_r16]{}
				if err = tmp_uci_OnPUSCH_ListDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uci_OnPUSCH_ListDCI_0_2_r16", err)
				}
				ie.uci_OnPUSCH_ListDCI_0_2_r16 = tmp_uci_OnPUSCH_ListDCI_0_2_r16.Setup
			}
			// decode pusch_TimeDomainAllocationListDCI_0_2_r16 optional
			if pusch_TimeDomainAllocationListDCI_0_2_r16Present {
				tmp_pusch_TimeDomainAllocationListDCI_0_2_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_pusch_TimeDomainAllocationListDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_TimeDomainAllocationListDCI_0_2_r16", err)
				}
				ie.pusch_TimeDomainAllocationListDCI_0_2_r16 = tmp_pusch_TimeDomainAllocationListDCI_0_2_r16.Setup
			}
			// decode pusch_TimeDomainAllocationListDCI_0_1_r16 optional
			if pusch_TimeDomainAllocationListDCI_0_1_r16Present {
				tmp_pusch_TimeDomainAllocationListDCI_0_1_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_pusch_TimeDomainAllocationListDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_TimeDomainAllocationListDCI_0_1_r16", err)
				}
				ie.pusch_TimeDomainAllocationListDCI_0_1_r16 = tmp_pusch_TimeDomainAllocationListDCI_0_1_r16.Setup
			}
			// decode invalidSymbolPatternIndicatorDCI_0_1_r16 optional
			if invalidSymbolPatternIndicatorDCI_0_1_r16Present {
				ie.invalidSymbolPatternIndicatorDCI_0_1_r16 = new(PUSCH_Config_invalidSymbolPatternIndicatorDCI_0_1_r16)
				if err = ie.invalidSymbolPatternIndicatorDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode invalidSymbolPatternIndicatorDCI_0_1_r16", err)
				}
			}
			// decode priorityIndicatorDCI_0_1_r16 optional
			if priorityIndicatorDCI_0_1_r16Present {
				ie.priorityIndicatorDCI_0_1_r16 = new(PUSCH_Config_priorityIndicatorDCI_0_1_r16)
				if err = ie.priorityIndicatorDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorDCI_0_1_r16", err)
				}
			}
			// decode pusch_RepTypeIndicatorDCI_0_1_r16 optional
			if pusch_RepTypeIndicatorDCI_0_1_r16Present {
				ie.pusch_RepTypeIndicatorDCI_0_1_r16 = new(PUSCH_Config_pusch_RepTypeIndicatorDCI_0_1_r16)
				if err = ie.pusch_RepTypeIndicatorDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_RepTypeIndicatorDCI_0_1_r16", err)
				}
			}
			// decode frequencyHoppingDCI_0_1_r16 optional
			if frequencyHoppingDCI_0_1_r16Present {
				ie.frequencyHoppingDCI_0_1_r16 = new(PUSCH_Config_frequencyHoppingDCI_0_1_r16)
				if err = ie.frequencyHoppingDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode frequencyHoppingDCI_0_1_r16", err)
				}
			}
			// decode uci_OnPUSCH_ListDCI_0_1_r16 optional
			if uci_OnPUSCH_ListDCI_0_1_r16Present {
				tmp_uci_OnPUSCH_ListDCI_0_1_r16 := utils.SetupRelease[*UCI_OnPUSCH_ListDCI_0_1_r16]{}
				if err = tmp_uci_OnPUSCH_ListDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uci_OnPUSCH_ListDCI_0_1_r16", err)
				}
				ie.uci_OnPUSCH_ListDCI_0_1_r16 = tmp_uci_OnPUSCH_ListDCI_0_1_r16.Setup
			}
			// decode invalidSymbolPattern_r16 optional
			if invalidSymbolPattern_r16Present {
				ie.invalidSymbolPattern_r16 = new(InvalidSymbolPattern_r16)
				if err = ie.invalidSymbolPattern_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode invalidSymbolPattern_r16", err)
				}
			}
			// decode pusch_PowerControl_v1610 optional
			if pusch_PowerControl_v1610Present {
				tmp_pusch_PowerControl_v1610 := utils.SetupRelease[*PUSCH_PowerControl_v1610]{}
				if err = tmp_pusch_PowerControl_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_PowerControl_v1610", err)
				}
				ie.pusch_PowerControl_v1610 = tmp_pusch_PowerControl_v1610.Setup
			}
			// decode ul_FullPowerTransmission_r16 optional
			if ul_FullPowerTransmission_r16Present {
				ie.ul_FullPowerTransmission_r16 = new(PUSCH_Config_ul_FullPowerTransmission_r16)
				if err = ie.ul_FullPowerTransmission_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_FullPowerTransmission_r16", err)
				}
			}
			// decode pusch_TimeDomainAllocationListForMultiPUSCH_r16 optional
			if pusch_TimeDomainAllocationListForMultiPUSCH_r16Present {
				tmp_pusch_TimeDomainAllocationListForMultiPUSCH_r16 := utils.SetupRelease[*PUSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_pusch_TimeDomainAllocationListForMultiPUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_TimeDomainAllocationListForMultiPUSCH_r16", err)
				}
				ie.pusch_TimeDomainAllocationListForMultiPUSCH_r16 = tmp_pusch_TimeDomainAllocationListForMultiPUSCH_r16.Setup
			}
			// decode numberOfInvalidSymbolsForDL_UL_Switching_r16 optional
			if numberOfInvalidSymbolsForDL_UL_Switching_r16Present {
				var tmp_int_numberOfInvalidSymbolsForDL_UL_Switching_r16 int64
				if tmp_int_numberOfInvalidSymbolsForDL_UL_Switching_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode numberOfInvalidSymbolsForDL_UL_Switching_r16", err)
				}
				ie.numberOfInvalidSymbolsForDL_UL_Switching_r16 = &tmp_int_numberOfInvalidSymbolsForDL_UL_Switching_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ul_AccessConfigListDCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			betaOffsetsCrossPri0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			betaOffsetsCrossPri1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			betaOffsetsCrossPri0DCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			betaOffsetsCrossPri1DCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mappingPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			secondTPCFieldDCI_0_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			secondTPCFieldDCI_0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sequenceOffsetForRV_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_AccessConfigListDCI_0_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			minimumSchedulingOffsetK2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			availableSlotCounting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingPUSCH_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcessNumberSizeDCI_0_2_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcessNumberSizeDCI_0_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mpe_ResourcePoolToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mpe_ResourcePoolToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ul_AccessConfigListDCI_0_2_r17 optional
			if ul_AccessConfigListDCI_0_2_r17Present {
				tmp_ul_AccessConfigListDCI_0_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_2_r17]{}
				if err = tmp_ul_AccessConfigListDCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_AccessConfigListDCI_0_2_r17", err)
				}
				ie.ul_AccessConfigListDCI_0_2_r17 = tmp_ul_AccessConfigListDCI_0_2_r17.Setup
			}
			// decode betaOffsetsCrossPri0_r17 optional
			if betaOffsetsCrossPri0_r17Present {
				tmp_betaOffsetsCrossPri0_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{}
				if err = tmp_betaOffsetsCrossPri0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode betaOffsetsCrossPri0_r17", err)
				}
				ie.betaOffsetsCrossPri0_r17 = tmp_betaOffsetsCrossPri0_r17.Setup
			}
			// decode betaOffsetsCrossPri1_r17 optional
			if betaOffsetsCrossPri1_r17Present {
				tmp_betaOffsetsCrossPri1_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSel_r17]{}
				if err = tmp_betaOffsetsCrossPri1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode betaOffsetsCrossPri1_r17", err)
				}
				ie.betaOffsetsCrossPri1_r17 = tmp_betaOffsetsCrossPri1_r17.Setup
			}
			// decode betaOffsetsCrossPri0DCI_0_2_r17 optional
			if betaOffsetsCrossPri0DCI_0_2_r17Present {
				tmp_betaOffsetsCrossPri0DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{}
				if err = tmp_betaOffsetsCrossPri0DCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode betaOffsetsCrossPri0DCI_0_2_r17", err)
				}
				ie.betaOffsetsCrossPri0DCI_0_2_r17 = tmp_betaOffsetsCrossPri0DCI_0_2_r17.Setup
			}
			// decode betaOffsetsCrossPri1DCI_0_2_r17 optional
			if betaOffsetsCrossPri1DCI_0_2_r17Present {
				tmp_betaOffsetsCrossPri1DCI_0_2_r17 := utils.SetupRelease[*BetaOffsetsCrossPriSelDCI_0_2_r17]{}
				if err = tmp_betaOffsetsCrossPri1DCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode betaOffsetsCrossPri1DCI_0_2_r17", err)
				}
				ie.betaOffsetsCrossPri1DCI_0_2_r17 = tmp_betaOffsetsCrossPri1DCI_0_2_r17.Setup
			}
			// decode mappingPattern_r17 optional
			if mappingPattern_r17Present {
				ie.mappingPattern_r17 = new(PUSCH_Config_mappingPattern_r17)
				if err = ie.mappingPattern_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mappingPattern_r17", err)
				}
			}
			// decode secondTPCFieldDCI_0_1_r17 optional
			if secondTPCFieldDCI_0_1_r17Present {
				ie.secondTPCFieldDCI_0_1_r17 = new(PUSCH_Config_secondTPCFieldDCI_0_1_r17)
				if err = ie.secondTPCFieldDCI_0_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode secondTPCFieldDCI_0_1_r17", err)
				}
			}
			// decode secondTPCFieldDCI_0_2_r17 optional
			if secondTPCFieldDCI_0_2_r17Present {
				ie.secondTPCFieldDCI_0_2_r17 = new(PUSCH_Config_secondTPCFieldDCI_0_2_r17)
				if err = ie.secondTPCFieldDCI_0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode secondTPCFieldDCI_0_2_r17", err)
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
			// decode ul_AccessConfigListDCI_0_1_r17 optional
			if ul_AccessConfigListDCI_0_1_r17Present {
				tmp_ul_AccessConfigListDCI_0_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_0_1_r17]{}
				if err = tmp_ul_AccessConfigListDCI_0_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_AccessConfigListDCI_0_1_r17", err)
				}
				ie.ul_AccessConfigListDCI_0_1_r17 = tmp_ul_AccessConfigListDCI_0_1_r17.Setup
			}
			// decode minimumSchedulingOffsetK2_r17 optional
			if minimumSchedulingOffsetK2_r17Present {
				tmp_minimumSchedulingOffsetK2_r17 := utils.SetupRelease[*MinSchedulingOffsetK2_Values_r17]{}
				if err = tmp_minimumSchedulingOffsetK2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode minimumSchedulingOffsetK2_r17", err)
				}
				ie.minimumSchedulingOffsetK2_r17 = tmp_minimumSchedulingOffsetK2_r17.Setup
			}
			// decode availableSlotCounting_r17 optional
			if availableSlotCounting_r17Present {
				ie.availableSlotCounting_r17 = new(PUSCH_Config_availableSlotCounting_r17)
				if err = ie.availableSlotCounting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode availableSlotCounting_r17", err)
				}
			}
			// decode dmrs_BundlingPUSCH_Config_r17 optional
			if dmrs_BundlingPUSCH_Config_r17Present {
				tmp_dmrs_BundlingPUSCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUSCH_Config_r17]{}
				if err = tmp_dmrs_BundlingPUSCH_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingPUSCH_Config_r17", err)
				}
				ie.dmrs_BundlingPUSCH_Config_r17 = tmp_dmrs_BundlingPUSCH_Config_r17.Setup
			}
			// decode harq_ProcessNumberSizeDCI_0_2_v1700 optional
			if harq_ProcessNumberSizeDCI_0_2_v1700Present {
				var tmp_int_harq_ProcessNumberSizeDCI_0_2_v1700 int64
				if tmp_int_harq_ProcessNumberSizeDCI_0_2_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode harq_ProcessNumberSizeDCI_0_2_v1700", err)
				}
				ie.harq_ProcessNumberSizeDCI_0_2_v1700 = &tmp_int_harq_ProcessNumberSizeDCI_0_2_v1700
			}
			// decode harq_ProcessNumberSizeDCI_0_1_r17 optional
			if harq_ProcessNumberSizeDCI_0_1_r17Present {
				var tmp_int_harq_ProcessNumberSizeDCI_0_1_r17 int64
				if tmp_int_harq_ProcessNumberSizeDCI_0_1_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode harq_ProcessNumberSizeDCI_0_1_r17", err)
				}
				ie.harq_ProcessNumberSizeDCI_0_1_r17 = &tmp_int_harq_ProcessNumberSizeDCI_0_1_r17
			}
			// decode mpe_ResourcePoolToAddModList_r17 optional
			if mpe_ResourcePoolToAddModList_r17Present {
				tmp_mpe_ResourcePoolToAddModList_r17 := utils.NewSequence[*MPE_Resource_r17]([]*MPE_Resource_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				fn_mpe_ResourcePoolToAddModList_r17 := func() *MPE_Resource_r17 {
					return new(MPE_Resource_r17)
				}
				if err = tmp_mpe_ResourcePoolToAddModList_r17.Decode(extReader, fn_mpe_ResourcePoolToAddModList_r17); err != nil {
					return utils.WrapError("Decode mpe_ResourcePoolToAddModList_r17", err)
				}
				ie.mpe_ResourcePoolToAddModList_r17 = []MPE_Resource_r17{}
				for _, i := range tmp_mpe_ResourcePoolToAddModList_r17.Value {
					ie.mpe_ResourcePoolToAddModList_r17 = append(ie.mpe_ResourcePoolToAddModList_r17, *i)
				}
			}
			// decode mpe_ResourcePoolToReleaseList_r17 optional
			if mpe_ResourcePoolToReleaseList_r17Present {
				tmp_mpe_ResourcePoolToReleaseList_r17 := utils.NewSequence[*MPE_ResourceId_r17]([]*MPE_ResourceId_r17{}, uper.Constraint{Lb: 1, Ub: maxMPE_Resources_r17}, false)
				fn_mpe_ResourcePoolToReleaseList_r17 := func() *MPE_ResourceId_r17 {
					return new(MPE_ResourceId_r17)
				}
				if err = tmp_mpe_ResourcePoolToReleaseList_r17.Decode(extReader, fn_mpe_ResourcePoolToReleaseList_r17); err != nil {
					return utils.WrapError("Decode mpe_ResourcePoolToReleaseList_r17", err)
				}
				ie.mpe_ResourcePoolToReleaseList_r17 = []MPE_ResourceId_r17{}
				for _, i := range tmp_mpe_ResourcePoolToReleaseList_r17.Value {
					ie.mpe_ResourcePoolToReleaseList_r17 = append(ie.mpe_ResourcePoolToReleaseList_r17, *i)
				}
			}
		}
	}
	return nil
}
