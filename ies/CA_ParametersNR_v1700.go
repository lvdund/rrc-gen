package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1700 struct {
	codebookParametersfetype2PerBC_r17          *CodebookParametersfetype2PerBC_r17                              `optional`
	demodulationEnhancementCA_r17               *CA_ParametersNR_v1700_demodulationEnhancementCA_r17             `optional`
	maxUplinkDutyCycle_interBandCA_PC2_r17      *CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17    `optional`
	maxUplinkDutyCycle_SULcombination_PC2_r17   *CA_ParametersNR_v1700_maxUplinkDutyCycle_SULcombination_PC2_r17 `optional`
	beamManagementType_CBM_r17                  *CA_ParametersNR_v1700_beamManagementType_CBM_r17                `optional`
	parallelTxPUCCH_PUSCH_r17                   *CA_ParametersNR_v1700_parallelTxPUCCH_PUSCH_r17                 `optional`
	codebookComboParameterMixedTypePerBC_r17    *CodebookComboParameterMixedTypePerBC_r17                        `optional`
	mTRP_CSI_EnhancementPerBC_r17               *CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17             `lb:1,ub:16,optional`
	codebookComboParameterMultiTRP_PerBC_r17    *CodebookComboParameterMultiTRP_PerBC_r17                        `optional`
	maxCC_32_DL_HARQ_ProcessFR2_2_r17           *CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17         `optional`
	maxCC_32_UL_HARQ_ProcessFR2_2_r17           *CA_ParametersNR_v1700_maxCC_32_UL_HARQ_ProcessFR2_2_r17         `optional`
	crossCarrierSchedulingSCell_SpCellTypeB_r17 *CrossCarrierSchedulingSCell_SpCell_r17                          `optional`
	crossCarrierSchedulingSCell_SpCellTypeA_r17 *CrossCarrierSchedulingSCell_SpCell_r17                          `optional`
	dci_FormatsPCellPSCellUSS_Sets_r17          *CA_ParametersNR_v1700_dci_FormatsPCellPSCellUSS_Sets_r17        `optional`
	disablingScalingFactorDeactSCell_r17        *CA_ParametersNR_v1700_disablingScalingFactorDeactSCell_r17      `optional`
	disablingScalingFactorDormantSCell_r17      *CA_ParametersNR_v1700_disablingScalingFactorDormantSCell_r17    `optional`
	non_AlignedFrameBoundaries_r17              *CA_ParametersNR_v1700_non_AlignedFrameBoundaries_r17            `lb:1,ub:496,optional`
}

func (ie *CA_ParametersNR_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.codebookParametersfetype2PerBC_r17 != nil, ie.demodulationEnhancementCA_r17 != nil, ie.maxUplinkDutyCycle_interBandCA_PC2_r17 != nil, ie.maxUplinkDutyCycle_SULcombination_PC2_r17 != nil, ie.beamManagementType_CBM_r17 != nil, ie.parallelTxPUCCH_PUSCH_r17 != nil, ie.codebookComboParameterMixedTypePerBC_r17 != nil, ie.mTRP_CSI_EnhancementPerBC_r17 != nil, ie.codebookComboParameterMultiTRP_PerBC_r17 != nil, ie.maxCC_32_DL_HARQ_ProcessFR2_2_r17 != nil, ie.maxCC_32_UL_HARQ_ProcessFR2_2_r17 != nil, ie.crossCarrierSchedulingSCell_SpCellTypeB_r17 != nil, ie.crossCarrierSchedulingSCell_SpCellTypeA_r17 != nil, ie.dci_FormatsPCellPSCellUSS_Sets_r17 != nil, ie.disablingScalingFactorDeactSCell_r17 != nil, ie.disablingScalingFactorDormantSCell_r17 != nil, ie.non_AlignedFrameBoundaries_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.codebookParametersfetype2PerBC_r17 != nil {
		if err = ie.codebookParametersfetype2PerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode codebookParametersfetype2PerBC_r17", err)
		}
	}
	if ie.demodulationEnhancementCA_r17 != nil {
		if err = ie.demodulationEnhancementCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode demodulationEnhancementCA_r17", err)
		}
	}
	if ie.maxUplinkDutyCycle_interBandCA_PC2_r17 != nil {
		if err = ie.maxUplinkDutyCycle_interBandCA_PC2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxUplinkDutyCycle_interBandCA_PC2_r17", err)
		}
	}
	if ie.maxUplinkDutyCycle_SULcombination_PC2_r17 != nil {
		if err = ie.maxUplinkDutyCycle_SULcombination_PC2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxUplinkDutyCycle_SULcombination_PC2_r17", err)
		}
	}
	if ie.beamManagementType_CBM_r17 != nil {
		if err = ie.beamManagementType_CBM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode beamManagementType_CBM_r17", err)
		}
	}
	if ie.parallelTxPUCCH_PUSCH_r17 != nil {
		if err = ie.parallelTxPUCCH_PUSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxPUCCH_PUSCH_r17", err)
		}
	}
	if ie.codebookComboParameterMixedTypePerBC_r17 != nil {
		if err = ie.codebookComboParameterMixedTypePerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode codebookComboParameterMixedTypePerBC_r17", err)
		}
	}
	if ie.mTRP_CSI_EnhancementPerBC_r17 != nil {
		if err = ie.mTRP_CSI_EnhancementPerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_CSI_EnhancementPerBC_r17", err)
		}
	}
	if ie.codebookComboParameterMultiTRP_PerBC_r17 != nil {
		if err = ie.codebookComboParameterMultiTRP_PerBC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode codebookComboParameterMultiTRP_PerBC_r17", err)
		}
	}
	if ie.maxCC_32_DL_HARQ_ProcessFR2_2_r17 != nil {
		if err = ie.maxCC_32_DL_HARQ_ProcessFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if ie.maxCC_32_UL_HARQ_ProcessFR2_2_r17 != nil {
		if err = ie.maxCC_32_UL_HARQ_ProcessFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxCC_32_UL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if ie.crossCarrierSchedulingSCell_SpCellTypeB_r17 != nil {
		if err = ie.crossCarrierSchedulingSCell_SpCellTypeB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingSCell_SpCellTypeB_r17", err)
		}
	}
	if ie.crossCarrierSchedulingSCell_SpCellTypeA_r17 != nil {
		if err = ie.crossCarrierSchedulingSCell_SpCellTypeA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingSCell_SpCellTypeA_r17", err)
		}
	}
	if ie.dci_FormatsPCellPSCellUSS_Sets_r17 != nil {
		if err = ie.dci_FormatsPCellPSCellUSS_Sets_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dci_FormatsPCellPSCellUSS_Sets_r17", err)
		}
	}
	if ie.disablingScalingFactorDeactSCell_r17 != nil {
		if err = ie.disablingScalingFactorDeactSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode disablingScalingFactorDeactSCell_r17", err)
		}
	}
	if ie.disablingScalingFactorDormantSCell_r17 != nil {
		if err = ie.disablingScalingFactorDormantSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode disablingScalingFactorDormantSCell_r17", err)
		}
	}
	if ie.non_AlignedFrameBoundaries_r17 != nil {
		if err = ie.non_AlignedFrameBoundaries_r17.Encode(w); err != nil {
			return utils.WrapError("Encode non_AlignedFrameBoundaries_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1700) Decode(r *uper.UperReader) error {
	var err error
	var codebookParametersfetype2PerBC_r17Present bool
	if codebookParametersfetype2PerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var demodulationEnhancementCA_r17Present bool
	if demodulationEnhancementCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxUplinkDutyCycle_interBandCA_PC2_r17Present bool
	if maxUplinkDutyCycle_interBandCA_PC2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxUplinkDutyCycle_SULcombination_PC2_r17Present bool
	if maxUplinkDutyCycle_SULcombination_PC2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamManagementType_CBM_r17Present bool
	if beamManagementType_CBM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var parallelTxPUCCH_PUSCH_r17Present bool
	if parallelTxPUCCH_PUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var codebookComboParameterMixedTypePerBC_r17Present bool
	if codebookComboParameterMixedTypePerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_CSI_EnhancementPerBC_r17Present bool
	if mTRP_CSI_EnhancementPerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var codebookComboParameterMultiTRP_PerBC_r17Present bool
	if codebookComboParameterMultiTRP_PerBC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCC_32_DL_HARQ_ProcessFR2_2_r17Present bool
	if maxCC_32_DL_HARQ_ProcessFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCC_32_UL_HARQ_ProcessFR2_2_r17Present bool
	if maxCC_32_UL_HARQ_ProcessFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingSCell_SpCellTypeB_r17Present bool
	if crossCarrierSchedulingSCell_SpCellTypeB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingSCell_SpCellTypeA_r17Present bool
	if crossCarrierSchedulingSCell_SpCellTypeA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_FormatsPCellPSCellUSS_Sets_r17Present bool
	if dci_FormatsPCellPSCellUSS_Sets_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var disablingScalingFactorDeactSCell_r17Present bool
	if disablingScalingFactorDeactSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var disablingScalingFactorDormantSCell_r17Present bool
	if disablingScalingFactorDormantSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var non_AlignedFrameBoundaries_r17Present bool
	if non_AlignedFrameBoundaries_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if codebookParametersfetype2PerBC_r17Present {
		ie.codebookParametersfetype2PerBC_r17 = new(CodebookParametersfetype2PerBC_r17)
		if err = ie.codebookParametersfetype2PerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode codebookParametersfetype2PerBC_r17", err)
		}
	}
	if demodulationEnhancementCA_r17Present {
		ie.demodulationEnhancementCA_r17 = new(CA_ParametersNR_v1700_demodulationEnhancementCA_r17)
		if err = ie.demodulationEnhancementCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode demodulationEnhancementCA_r17", err)
		}
	}
	if maxUplinkDutyCycle_interBandCA_PC2_r17Present {
		ie.maxUplinkDutyCycle_interBandCA_PC2_r17 = new(CA_ParametersNR_v1700_maxUplinkDutyCycle_interBandCA_PC2_r17)
		if err = ie.maxUplinkDutyCycle_interBandCA_PC2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxUplinkDutyCycle_interBandCA_PC2_r17", err)
		}
	}
	if maxUplinkDutyCycle_SULcombination_PC2_r17Present {
		ie.maxUplinkDutyCycle_SULcombination_PC2_r17 = new(CA_ParametersNR_v1700_maxUplinkDutyCycle_SULcombination_PC2_r17)
		if err = ie.maxUplinkDutyCycle_SULcombination_PC2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxUplinkDutyCycle_SULcombination_PC2_r17", err)
		}
	}
	if beamManagementType_CBM_r17Present {
		ie.beamManagementType_CBM_r17 = new(CA_ParametersNR_v1700_beamManagementType_CBM_r17)
		if err = ie.beamManagementType_CBM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode beamManagementType_CBM_r17", err)
		}
	}
	if parallelTxPUCCH_PUSCH_r17Present {
		ie.parallelTxPUCCH_PUSCH_r17 = new(CA_ParametersNR_v1700_parallelTxPUCCH_PUSCH_r17)
		if err = ie.parallelTxPUCCH_PUSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxPUCCH_PUSCH_r17", err)
		}
	}
	if codebookComboParameterMixedTypePerBC_r17Present {
		ie.codebookComboParameterMixedTypePerBC_r17 = new(CodebookComboParameterMixedTypePerBC_r17)
		if err = ie.codebookComboParameterMixedTypePerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode codebookComboParameterMixedTypePerBC_r17", err)
		}
	}
	if mTRP_CSI_EnhancementPerBC_r17Present {
		ie.mTRP_CSI_EnhancementPerBC_r17 = new(CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17)
		if err = ie.mTRP_CSI_EnhancementPerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_CSI_EnhancementPerBC_r17", err)
		}
	}
	if codebookComboParameterMultiTRP_PerBC_r17Present {
		ie.codebookComboParameterMultiTRP_PerBC_r17 = new(CodebookComboParameterMultiTRP_PerBC_r17)
		if err = ie.codebookComboParameterMultiTRP_PerBC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode codebookComboParameterMultiTRP_PerBC_r17", err)
		}
	}
	if maxCC_32_DL_HARQ_ProcessFR2_2_r17Present {
		ie.maxCC_32_DL_HARQ_ProcessFR2_2_r17 = new(CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17)
		if err = ie.maxCC_32_DL_HARQ_ProcessFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if maxCC_32_UL_HARQ_ProcessFR2_2_r17Present {
		ie.maxCC_32_UL_HARQ_ProcessFR2_2_r17 = new(CA_ParametersNR_v1700_maxCC_32_UL_HARQ_ProcessFR2_2_r17)
		if err = ie.maxCC_32_UL_HARQ_ProcessFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxCC_32_UL_HARQ_ProcessFR2_2_r17", err)
		}
	}
	if crossCarrierSchedulingSCell_SpCellTypeB_r17Present {
		ie.crossCarrierSchedulingSCell_SpCellTypeB_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17)
		if err = ie.crossCarrierSchedulingSCell_SpCellTypeB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingSCell_SpCellTypeB_r17", err)
		}
	}
	if crossCarrierSchedulingSCell_SpCellTypeA_r17Present {
		ie.crossCarrierSchedulingSCell_SpCellTypeA_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17)
		if err = ie.crossCarrierSchedulingSCell_SpCellTypeA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingSCell_SpCellTypeA_r17", err)
		}
	}
	if dci_FormatsPCellPSCellUSS_Sets_r17Present {
		ie.dci_FormatsPCellPSCellUSS_Sets_r17 = new(CA_ParametersNR_v1700_dci_FormatsPCellPSCellUSS_Sets_r17)
		if err = ie.dci_FormatsPCellPSCellUSS_Sets_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dci_FormatsPCellPSCellUSS_Sets_r17", err)
		}
	}
	if disablingScalingFactorDeactSCell_r17Present {
		ie.disablingScalingFactorDeactSCell_r17 = new(CA_ParametersNR_v1700_disablingScalingFactorDeactSCell_r17)
		if err = ie.disablingScalingFactorDeactSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode disablingScalingFactorDeactSCell_r17", err)
		}
	}
	if disablingScalingFactorDormantSCell_r17Present {
		ie.disablingScalingFactorDormantSCell_r17 = new(CA_ParametersNR_v1700_disablingScalingFactorDormantSCell_r17)
		if err = ie.disablingScalingFactorDormantSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode disablingScalingFactorDormantSCell_r17", err)
		}
	}
	if non_AlignedFrameBoundaries_r17Present {
		ie.non_AlignedFrameBoundaries_r17 = new(CA_ParametersNR_v1700_non_AlignedFrameBoundaries_r17)
		if err = ie.non_AlignedFrameBoundaries_r17.Decode(r); err != nil {
			return utils.WrapError("Decode non_AlignedFrameBoundaries_r17", err)
		}
	}
	return nil
}
