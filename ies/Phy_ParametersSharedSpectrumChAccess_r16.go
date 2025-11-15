package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersSharedSpectrumChAccess_r16 struct {
	ss_SINR_Meas_r16                           *Phy_ParametersSharedSpectrumChAccess_r16_ss_SINR_Meas_r16                           `optional`
	sp_CSI_ReportPUCCH_r16                     *Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUCCH_r16                     `optional`
	sp_CSI_ReportPUSCH_r16                     *Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUSCH_r16                     `optional`
	dynamicSFI_r16                             *Phy_ParametersSharedSpectrumChAccess_r16_dynamicSFI_r16                             `optional`
	mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16  *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16  `optional`
	mux_SR_HARQ_ACK_PUCCH_r16                  *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_PUCCH_r16                  `optional`
	mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 `optional`
	mux_HARQ_ACK_PUSCH_DiffSymbol_r16          *Phy_ParametersSharedSpectrumChAccess_r16_mux_HARQ_ACK_PUSCH_DiffSymbol_r16          `optional`
	pucch_Repetition_F1_3_4_r16                *Phy_ParametersSharedSpectrumChAccess_r16_pucch_Repetition_F1_3_4_r16                `optional`
	type1_PUSCH_RepetitionMultiSlots_r16       *Phy_ParametersSharedSpectrumChAccess_r16_type1_PUSCH_RepetitionMultiSlots_r16       `optional`
	type2_PUSCH_RepetitionMultiSlots_r16       *Phy_ParametersSharedSpectrumChAccess_r16_type2_PUSCH_RepetitionMultiSlots_r16       `optional`
	pusch_RepetitionMultiSlots_r16             *Phy_ParametersSharedSpectrumChAccess_r16_pusch_RepetitionMultiSlots_r16             `optional`
	pdsch_RepetitionMultiSlots_r16             *Phy_ParametersSharedSpectrumChAccess_r16_pdsch_RepetitionMultiSlots_r16             `optional`
	downlinkSPS_r16                            *Phy_ParametersSharedSpectrumChAccess_r16_downlinkSPS_r16                            `optional`
	configuredUL_GrantType1_r16                *Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType1_r16                `optional`
	configuredUL_GrantType2_r16                *Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType2_r16                `optional`
	pre_EmptIndication_DL_r16                  *Phy_ParametersSharedSpectrumChAccess_r16_pre_EmptIndication_DL_r16                  `optional`
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ss_SINR_Meas_r16 != nil, ie.sp_CSI_ReportPUCCH_r16 != nil, ie.sp_CSI_ReportPUSCH_r16 != nil, ie.dynamicSFI_r16 != nil, ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 != nil, ie.mux_SR_HARQ_ACK_PUCCH_r16 != nil, ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 != nil, ie.mux_HARQ_ACK_PUSCH_DiffSymbol_r16 != nil, ie.pucch_Repetition_F1_3_4_r16 != nil, ie.type1_PUSCH_RepetitionMultiSlots_r16 != nil, ie.type2_PUSCH_RepetitionMultiSlots_r16 != nil, ie.pusch_RepetitionMultiSlots_r16 != nil, ie.pdsch_RepetitionMultiSlots_r16 != nil, ie.downlinkSPS_r16 != nil, ie.configuredUL_GrantType1_r16 != nil, ie.configuredUL_GrantType2_r16 != nil, ie.pre_EmptIndication_DL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ss_SINR_Meas_r16 != nil {
		if err = ie.ss_SINR_Meas_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ss_SINR_Meas_r16", err)
		}
	}
	if ie.sp_CSI_ReportPUCCH_r16 != nil {
		if err = ie.sp_CSI_ReportPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_ReportPUCCH_r16", err)
		}
	}
	if ie.sp_CSI_ReportPUSCH_r16 != nil {
		if err = ie.sp_CSI_ReportPUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_ReportPUSCH_r16", err)
		}
	}
	if ie.dynamicSFI_r16 != nil {
		if err = ie.dynamicSFI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicSFI_r16", err)
		}
	}
	if ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 != nil {
		if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16", err)
		}
	}
	if ie.mux_SR_HARQ_ACK_PUCCH_r16 != nil {
		if err = ie.mux_SR_HARQ_ACK_PUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mux_SR_HARQ_ACK_PUCCH_r16", err)
		}
	}
	if ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 != nil {
		if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16", err)
		}
	}
	if ie.mux_HARQ_ACK_PUSCH_DiffSymbol_r16 != nil {
		if err = ie.mux_HARQ_ACK_PUSCH_DiffSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mux_HARQ_ACK_PUSCH_DiffSymbol_r16", err)
		}
	}
	if ie.pucch_Repetition_F1_3_4_r16 != nil {
		if err = ie.pucch_Repetition_F1_3_4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_Repetition_F1_3_4_r16", err)
		}
	}
	if ie.type1_PUSCH_RepetitionMultiSlots_r16 != nil {
		if err = ie.type1_PUSCH_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.type2_PUSCH_RepetitionMultiSlots_r16 != nil {
		if err = ie.type2_PUSCH_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type2_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.pusch_RepetitionMultiSlots_r16 != nil {
		if err = ie.pusch_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.pdsch_RepetitionMultiSlots_r16 != nil {
		if err = ie.pdsch_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.downlinkSPS_r16 != nil {
		if err = ie.downlinkSPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode downlinkSPS_r16", err)
		}
	}
	if ie.configuredUL_GrantType1_r16 != nil {
		if err = ie.configuredUL_GrantType1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode configuredUL_GrantType1_r16", err)
		}
	}
	if ie.configuredUL_GrantType2_r16 != nil {
		if err = ie.configuredUL_GrantType2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode configuredUL_GrantType2_r16", err)
		}
	}
	if ie.pre_EmptIndication_DL_r16 != nil {
		if err = ie.pre_EmptIndication_DL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pre_EmptIndication_DL_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16) Decode(r *uper.UperReader) error {
	var err error
	var ss_SINR_Meas_r16Present bool
	if ss_SINR_Meas_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_ReportPUCCH_r16Present bool
	if sp_CSI_ReportPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_ReportPUSCH_r16Present bool
	if sp_CSI_ReportPUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicSFI_r16Present bool
	if dynamicSFI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16Present bool
	if mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_SR_HARQ_ACK_PUCCH_r16Present bool
	if mux_SR_HARQ_ACK_PUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16Present bool
	if mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_HARQ_ACK_PUSCH_DiffSymbol_r16Present bool
	if mux_HARQ_ACK_PUSCH_DiffSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_Repetition_F1_3_4_r16Present bool
	if pucch_Repetition_F1_3_4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1_PUSCH_RepetitionMultiSlots_r16Present bool
	if type1_PUSCH_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_PUSCH_RepetitionMultiSlots_r16Present bool
	if type2_PUSCH_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_RepetitionMultiSlots_r16Present bool
	if pusch_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_RepetitionMultiSlots_r16Present bool
	if pdsch_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var downlinkSPS_r16Present bool
	if downlinkSPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredUL_GrantType1_r16Present bool
	if configuredUL_GrantType1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredUL_GrantType2_r16Present bool
	if configuredUL_GrantType2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pre_EmptIndication_DL_r16Present bool
	if pre_EmptIndication_DL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ss_SINR_Meas_r16Present {
		ie.ss_SINR_Meas_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_ss_SINR_Meas_r16)
		if err = ie.ss_SINR_Meas_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ss_SINR_Meas_r16", err)
		}
	}
	if sp_CSI_ReportPUCCH_r16Present {
		ie.sp_CSI_ReportPUCCH_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUCCH_r16)
		if err = ie.sp_CSI_ReportPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_ReportPUCCH_r16", err)
		}
	}
	if sp_CSI_ReportPUSCH_r16Present {
		ie.sp_CSI_ReportPUSCH_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUSCH_r16)
		if err = ie.sp_CSI_ReportPUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_ReportPUSCH_r16", err)
		}
	}
	if dynamicSFI_r16Present {
		ie.dynamicSFI_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_dynamicSFI_r16)
		if err = ie.dynamicSFI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicSFI_r16", err)
		}
	}
	if mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16Present {
		ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16)
		if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16", err)
		}
	}
	if mux_SR_HARQ_ACK_PUCCH_r16Present {
		ie.mux_SR_HARQ_ACK_PUCCH_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_PUCCH_r16)
		if err = ie.mux_SR_HARQ_ACK_PUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mux_SR_HARQ_ACK_PUCCH_r16", err)
		}
	}
	if mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16Present {
		ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16)
		if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16", err)
		}
	}
	if mux_HARQ_ACK_PUSCH_DiffSymbol_r16Present {
		ie.mux_HARQ_ACK_PUSCH_DiffSymbol_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_HARQ_ACK_PUSCH_DiffSymbol_r16)
		if err = ie.mux_HARQ_ACK_PUSCH_DiffSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mux_HARQ_ACK_PUSCH_DiffSymbol_r16", err)
		}
	}
	if pucch_Repetition_F1_3_4_r16Present {
		ie.pucch_Repetition_F1_3_4_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pucch_Repetition_F1_3_4_r16)
		if err = ie.pucch_Repetition_F1_3_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_Repetition_F1_3_4_r16", err)
		}
	}
	if type1_PUSCH_RepetitionMultiSlots_r16Present {
		ie.type1_PUSCH_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_type1_PUSCH_RepetitionMultiSlots_r16)
		if err = ie.type1_PUSCH_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type1_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if type2_PUSCH_RepetitionMultiSlots_r16Present {
		ie.type2_PUSCH_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_type2_PUSCH_RepetitionMultiSlots_r16)
		if err = ie.type2_PUSCH_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type2_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if pusch_RepetitionMultiSlots_r16Present {
		ie.pusch_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pusch_RepetitionMultiSlots_r16)
		if err = ie.pusch_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_RepetitionMultiSlots_r16", err)
		}
	}
	if pdsch_RepetitionMultiSlots_r16Present {
		ie.pdsch_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pdsch_RepetitionMultiSlots_r16)
		if err = ie.pdsch_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_RepetitionMultiSlots_r16", err)
		}
	}
	if downlinkSPS_r16Present {
		ie.downlinkSPS_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_downlinkSPS_r16)
		if err = ie.downlinkSPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode downlinkSPS_r16", err)
		}
	}
	if configuredUL_GrantType1_r16Present {
		ie.configuredUL_GrantType1_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType1_r16)
		if err = ie.configuredUL_GrantType1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode configuredUL_GrantType1_r16", err)
		}
	}
	if configuredUL_GrantType2_r16Present {
		ie.configuredUL_GrantType2_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType2_r16)
		if err = ie.configuredUL_GrantType2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode configuredUL_GrantType2_r16", err)
		}
	}
	if pre_EmptIndication_DL_r16Present {
		ie.pre_EmptIndication_DL_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pre_EmptIndication_DL_r16)
		if err = ie.pre_EmptIndication_DL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pre_EmptIndication_DL_r16", err)
		}
	}
	return nil
}
