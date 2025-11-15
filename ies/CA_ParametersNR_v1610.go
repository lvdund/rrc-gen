package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610 struct {
	parallelTxMsgA_SRS_PUCCH_PUSCH_r16       *CA_ParametersNR_v1610_parallelTxMsgA_SRS_PUCCH_PUSCH_r16    `optional`
	msgA_SUL_r16                             *CA_ParametersNR_v1610_msgA_SUL_r16                          `optional`
	jointSearchSpaceSwitchAcrossCells_r16    *CA_ParametersNR_v1610_jointSearchSpaceSwitchAcrossCells_r16 `optional`
	half_DuplexTDD_CA_SameSCS_r16            *CA_ParametersNR_v1610_half_DuplexTDD_CA_SameSCS_r16         `optional`
	scellDormancyWithinActiveTime_r16        *CA_ParametersNR_v1610_scellDormancyWithinActiveTime_r16     `optional`
	scellDormancyOutsideActiveTime_r16       *CA_ParametersNR_v1610_scellDormancyOutsideActiveTime_r16    `optional`
	crossCarrierA_CSI_trigDiffSCS_r16        *CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16     `optional`
	defaultQCL_CrossCarrierA_CSI_Trig_r16    *CA_ParametersNR_v1610_defaultQCL_CrossCarrierA_CSI_Trig_r16 `optional`
	interCA_NonAlignedFrame_r16              *CA_ParametersNR_v1610_interCA_NonAlignedFrame_r16           `optional`
	simul_SRS_Trans_BC_r16                   *CA_ParametersNR_v1610_simul_SRS_Trans_BC_r16                `optional`
	interFreqDAPS_r16                        *CA_ParametersNR_v1610_interFreqDAPS_r16                     `optional`
	codebookParametersPerBC_r16              *CodebookParameters_v1610                                    `optional`
	blindDetectFactor_r16                    *int64                                                       `lb:1,ub:2,optional`
	pdcch_MonitoringCA_r16                   *CA_ParametersNR_v1610_pdcch_MonitoringCA_r16                `lb:2,ub:16,optional`
	pdcch_BlindDetectionCA_Mixed_r16         *CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16      `lb:1,ub:15,optional`
	pdcch_BlindDetectionMCG_UE_r16           *int64                                                       `lb:1,ub:14,optional`
	pdcch_BlindDetectionSCG_UE_r16           *int64                                                       `lb:1,ub:14,optional`
	pdcch_BlindDetectionMCG_UE_Mixed_r16     *CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16  `lb:0,ub:15,optional`
	pdcch_BlindDetectionSCG_UE_Mixed_r16     *CA_ParametersNR_v1610_pdcch_BlindDetectionSCG_UE_Mixed_r16  `lb:0,ub:15,optional`
	crossCarrierSchedulingDL_DiffSCS_r16     *CA_ParametersNR_v1610_crossCarrierSchedulingDL_DiffSCS_r16  `optional`
	crossCarrierSchedulingDefaultQCL_r16     *CA_ParametersNR_v1610_crossCarrierSchedulingDefaultQCL_r16  `optional`
	crossCarrierSchedulingUL_DiffSCS_r16     *CA_ParametersNR_v1610_crossCarrierSchedulingUL_DiffSCS_r16  `optional`
	simul_SRS_MIMO_Trans_BC_r16              *CA_ParametersNR_v1610_simul_SRS_MIMO_Trans_BC_r16           `optional`
	codebookParametersAdditionPerBC_r16      *CodebookParametersAdditionPerBC_r16                         `optional`
	codebookComboParametersAdditionPerBC_r16 *CodebookComboParametersAdditionPerBC_r16                    `optional`
}

func (ie *CA_ParametersNR_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.parallelTxMsgA_SRS_PUCCH_PUSCH_r16 != nil, ie.msgA_SUL_r16 != nil, ie.jointSearchSpaceSwitchAcrossCells_r16 != nil, ie.half_DuplexTDD_CA_SameSCS_r16 != nil, ie.scellDormancyWithinActiveTime_r16 != nil, ie.scellDormancyOutsideActiveTime_r16 != nil, ie.crossCarrierA_CSI_trigDiffSCS_r16 != nil, ie.defaultQCL_CrossCarrierA_CSI_Trig_r16 != nil, ie.interCA_NonAlignedFrame_r16 != nil, ie.simul_SRS_Trans_BC_r16 != nil, ie.interFreqDAPS_r16 != nil, ie.codebookParametersPerBC_r16 != nil, ie.blindDetectFactor_r16 != nil, ie.pdcch_MonitoringCA_r16 != nil, ie.pdcch_BlindDetectionCA_Mixed_r16 != nil, ie.pdcch_BlindDetectionMCG_UE_r16 != nil, ie.pdcch_BlindDetectionSCG_UE_r16 != nil, ie.pdcch_BlindDetectionMCG_UE_Mixed_r16 != nil, ie.pdcch_BlindDetectionSCG_UE_Mixed_r16 != nil, ie.crossCarrierSchedulingDL_DiffSCS_r16 != nil, ie.crossCarrierSchedulingDefaultQCL_r16 != nil, ie.crossCarrierSchedulingUL_DiffSCS_r16 != nil, ie.simul_SRS_MIMO_Trans_BC_r16 != nil, ie.codebookParametersAdditionPerBC_r16 != nil, ie.codebookComboParametersAdditionPerBC_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.parallelTxMsgA_SRS_PUCCH_PUSCH_r16 != nil {
		if err = ie.parallelTxMsgA_SRS_PUCCH_PUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxMsgA_SRS_PUCCH_PUSCH_r16", err)
		}
	}
	if ie.msgA_SUL_r16 != nil {
		if err = ie.msgA_SUL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_SUL_r16", err)
		}
	}
	if ie.jointSearchSpaceSwitchAcrossCells_r16 != nil {
		if err = ie.jointSearchSpaceSwitchAcrossCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode jointSearchSpaceSwitchAcrossCells_r16", err)
		}
	}
	if ie.half_DuplexTDD_CA_SameSCS_r16 != nil {
		if err = ie.half_DuplexTDD_CA_SameSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode half_DuplexTDD_CA_SameSCS_r16", err)
		}
	}
	if ie.scellDormancyWithinActiveTime_r16 != nil {
		if err = ie.scellDormancyWithinActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scellDormancyWithinActiveTime_r16", err)
		}
	}
	if ie.scellDormancyOutsideActiveTime_r16 != nil {
		if err = ie.scellDormancyOutsideActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scellDormancyOutsideActiveTime_r16", err)
		}
	}
	if ie.crossCarrierA_CSI_trigDiffSCS_r16 != nil {
		if err = ie.crossCarrierA_CSI_trigDiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierA_CSI_trigDiffSCS_r16", err)
		}
	}
	if ie.defaultQCL_CrossCarrierA_CSI_Trig_r16 != nil {
		if err = ie.defaultQCL_CrossCarrierA_CSI_Trig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode defaultQCL_CrossCarrierA_CSI_Trig_r16", err)
		}
	}
	if ie.interCA_NonAlignedFrame_r16 != nil {
		if err = ie.interCA_NonAlignedFrame_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interCA_NonAlignedFrame_r16", err)
		}
	}
	if ie.simul_SRS_Trans_BC_r16 != nil {
		if err = ie.simul_SRS_Trans_BC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode simul_SRS_Trans_BC_r16", err)
		}
	}
	if ie.interFreqDAPS_r16 != nil {
		if err = ie.interFreqDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqDAPS_r16", err)
		}
	}
	if ie.codebookParametersPerBC_r16 != nil {
		if err = ie.codebookParametersPerBC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode codebookParametersPerBC_r16", err)
		}
	}
	if ie.blindDetectFactor_r16 != nil {
		if err = w.WriteInteger(*ie.blindDetectFactor_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode blindDetectFactor_r16", err)
		}
	}
	if ie.pdcch_MonitoringCA_r16 != nil {
		if err = ie.pdcch_MonitoringCA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_MonitoringCA_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionCA_Mixed_r16 != nil {
		if err = ie.pdcch_BlindDetectionCA_Mixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA_Mixed_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionMCG_UE_r16 != nil {
		if err = w.WriteInteger(*ie.pdcch_BlindDetectionMCG_UE_r16, &uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionMCG_UE_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionSCG_UE_r16 != nil {
		if err = w.WriteInteger(*ie.pdcch_BlindDetectionSCG_UE_r16, &uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionSCG_UE_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionMCG_UE_Mixed_r16 != nil {
		if err = ie.pdcch_BlindDetectionMCG_UE_Mixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionMCG_UE_Mixed_r16", err)
		}
	}
	if ie.pdcch_BlindDetectionSCG_UE_Mixed_r16 != nil {
		if err = ie.pdcch_BlindDetectionSCG_UE_Mixed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionSCG_UE_Mixed_r16", err)
		}
	}
	if ie.crossCarrierSchedulingDL_DiffSCS_r16 != nil {
		if err = ie.crossCarrierSchedulingDL_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingDL_DiffSCS_r16", err)
		}
	}
	if ie.crossCarrierSchedulingDefaultQCL_r16 != nil {
		if err = ie.crossCarrierSchedulingDefaultQCL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingDefaultQCL_r16", err)
		}
	}
	if ie.crossCarrierSchedulingUL_DiffSCS_r16 != nil {
		if err = ie.crossCarrierSchedulingUL_DiffSCS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierSchedulingUL_DiffSCS_r16", err)
		}
	}
	if ie.simul_SRS_MIMO_Trans_BC_r16 != nil {
		if err = ie.simul_SRS_MIMO_Trans_BC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode simul_SRS_MIMO_Trans_BC_r16", err)
		}
	}
	if ie.codebookParametersAdditionPerBC_r16 != nil {
		if err = ie.codebookParametersAdditionPerBC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode codebookParametersAdditionPerBC_r16", err)
		}
	}
	if ie.codebookComboParametersAdditionPerBC_r16 != nil {
		if err = ie.codebookComboParametersAdditionPerBC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode codebookComboParametersAdditionPerBC_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1610) Decode(r *uper.UperReader) error {
	var err error
	var parallelTxMsgA_SRS_PUCCH_PUSCH_r16Present bool
	if parallelTxMsgA_SRS_PUCCH_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_SUL_r16Present bool
	if msgA_SUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var jointSearchSpaceSwitchAcrossCells_r16Present bool
	if jointSearchSpaceSwitchAcrossCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var half_DuplexTDD_CA_SameSCS_r16Present bool
	if half_DuplexTDD_CA_SameSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scellDormancyWithinActiveTime_r16Present bool
	if scellDormancyWithinActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scellDormancyOutsideActiveTime_r16Present bool
	if scellDormancyOutsideActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierA_CSI_trigDiffSCS_r16Present bool
	if crossCarrierA_CSI_trigDiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var defaultQCL_CrossCarrierA_CSI_Trig_r16Present bool
	if defaultQCL_CrossCarrierA_CSI_Trig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interCA_NonAlignedFrame_r16Present bool
	if interCA_NonAlignedFrame_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var simul_SRS_Trans_BC_r16Present bool
	if simul_SRS_Trans_BC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqDAPS_r16Present bool
	if interFreqDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var codebookParametersPerBC_r16Present bool
	if codebookParametersPerBC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var blindDetectFactor_r16Present bool
	if blindDetectFactor_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_MonitoringCA_r16Present bool
	if pdcch_MonitoringCA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionCA_Mixed_r16Present bool
	if pdcch_BlindDetectionCA_Mixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionMCG_UE_r16Present bool
	if pdcch_BlindDetectionMCG_UE_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionSCG_UE_r16Present bool
	if pdcch_BlindDetectionSCG_UE_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionMCG_UE_Mixed_r16Present bool
	if pdcch_BlindDetectionMCG_UE_Mixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionSCG_UE_Mixed_r16Present bool
	if pdcch_BlindDetectionSCG_UE_Mixed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingDL_DiffSCS_r16Present bool
	if crossCarrierSchedulingDL_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingDefaultQCL_r16Present bool
	if crossCarrierSchedulingDefaultQCL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierSchedulingUL_DiffSCS_r16Present bool
	if crossCarrierSchedulingUL_DiffSCS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var simul_SRS_MIMO_Trans_BC_r16Present bool
	if simul_SRS_MIMO_Trans_BC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var codebookParametersAdditionPerBC_r16Present bool
	if codebookParametersAdditionPerBC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var codebookComboParametersAdditionPerBC_r16Present bool
	if codebookComboParametersAdditionPerBC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if parallelTxMsgA_SRS_PUCCH_PUSCH_r16Present {
		ie.parallelTxMsgA_SRS_PUCCH_PUSCH_r16 = new(CA_ParametersNR_v1610_parallelTxMsgA_SRS_PUCCH_PUSCH_r16)
		if err = ie.parallelTxMsgA_SRS_PUCCH_PUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxMsgA_SRS_PUCCH_PUSCH_r16", err)
		}
	}
	if msgA_SUL_r16Present {
		ie.msgA_SUL_r16 = new(CA_ParametersNR_v1610_msgA_SUL_r16)
		if err = ie.msgA_SUL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_SUL_r16", err)
		}
	}
	if jointSearchSpaceSwitchAcrossCells_r16Present {
		ie.jointSearchSpaceSwitchAcrossCells_r16 = new(CA_ParametersNR_v1610_jointSearchSpaceSwitchAcrossCells_r16)
		if err = ie.jointSearchSpaceSwitchAcrossCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode jointSearchSpaceSwitchAcrossCells_r16", err)
		}
	}
	if half_DuplexTDD_CA_SameSCS_r16Present {
		ie.half_DuplexTDD_CA_SameSCS_r16 = new(CA_ParametersNR_v1610_half_DuplexTDD_CA_SameSCS_r16)
		if err = ie.half_DuplexTDD_CA_SameSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode half_DuplexTDD_CA_SameSCS_r16", err)
		}
	}
	if scellDormancyWithinActiveTime_r16Present {
		ie.scellDormancyWithinActiveTime_r16 = new(CA_ParametersNR_v1610_scellDormancyWithinActiveTime_r16)
		if err = ie.scellDormancyWithinActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scellDormancyWithinActiveTime_r16", err)
		}
	}
	if scellDormancyOutsideActiveTime_r16Present {
		ie.scellDormancyOutsideActiveTime_r16 = new(CA_ParametersNR_v1610_scellDormancyOutsideActiveTime_r16)
		if err = ie.scellDormancyOutsideActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scellDormancyOutsideActiveTime_r16", err)
		}
	}
	if crossCarrierA_CSI_trigDiffSCS_r16Present {
		ie.crossCarrierA_CSI_trigDiffSCS_r16 = new(CA_ParametersNR_v1610_crossCarrierA_CSI_trigDiffSCS_r16)
		if err = ie.crossCarrierA_CSI_trigDiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierA_CSI_trigDiffSCS_r16", err)
		}
	}
	if defaultQCL_CrossCarrierA_CSI_Trig_r16Present {
		ie.defaultQCL_CrossCarrierA_CSI_Trig_r16 = new(CA_ParametersNR_v1610_defaultQCL_CrossCarrierA_CSI_Trig_r16)
		if err = ie.defaultQCL_CrossCarrierA_CSI_Trig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode defaultQCL_CrossCarrierA_CSI_Trig_r16", err)
		}
	}
	if interCA_NonAlignedFrame_r16Present {
		ie.interCA_NonAlignedFrame_r16 = new(CA_ParametersNR_v1610_interCA_NonAlignedFrame_r16)
		if err = ie.interCA_NonAlignedFrame_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interCA_NonAlignedFrame_r16", err)
		}
	}
	if simul_SRS_Trans_BC_r16Present {
		ie.simul_SRS_Trans_BC_r16 = new(CA_ParametersNR_v1610_simul_SRS_Trans_BC_r16)
		if err = ie.simul_SRS_Trans_BC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode simul_SRS_Trans_BC_r16", err)
		}
	}
	if interFreqDAPS_r16Present {
		ie.interFreqDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16)
		if err = ie.interFreqDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqDAPS_r16", err)
		}
	}
	if codebookParametersPerBC_r16Present {
		ie.codebookParametersPerBC_r16 = new(CodebookParameters_v1610)
		if err = ie.codebookParametersPerBC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode codebookParametersPerBC_r16", err)
		}
	}
	if blindDetectFactor_r16Present {
		var tmp_int_blindDetectFactor_r16 int64
		if tmp_int_blindDetectFactor_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode blindDetectFactor_r16", err)
		}
		ie.blindDetectFactor_r16 = &tmp_int_blindDetectFactor_r16
	}
	if pdcch_MonitoringCA_r16Present {
		ie.pdcch_MonitoringCA_r16 = new(CA_ParametersNR_v1610_pdcch_MonitoringCA_r16)
		if err = ie.pdcch_MonitoringCA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_MonitoringCA_r16", err)
		}
	}
	if pdcch_BlindDetectionCA_Mixed_r16Present {
		ie.pdcch_BlindDetectionCA_Mixed_r16 = new(CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16)
		if err = ie.pdcch_BlindDetectionCA_Mixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA_Mixed_r16", err)
		}
	}
	if pdcch_BlindDetectionMCG_UE_r16Present {
		var tmp_int_pdcch_BlindDetectionMCG_UE_r16 int64
		if tmp_int_pdcch_BlindDetectionMCG_UE_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionMCG_UE_r16", err)
		}
		ie.pdcch_BlindDetectionMCG_UE_r16 = &tmp_int_pdcch_BlindDetectionMCG_UE_r16
	}
	if pdcch_BlindDetectionSCG_UE_r16Present {
		var tmp_int_pdcch_BlindDetectionSCG_UE_r16 int64
		if tmp_int_pdcch_BlindDetectionSCG_UE_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionSCG_UE_r16", err)
		}
		ie.pdcch_BlindDetectionSCG_UE_r16 = &tmp_int_pdcch_BlindDetectionSCG_UE_r16
	}
	if pdcch_BlindDetectionMCG_UE_Mixed_r16Present {
		ie.pdcch_BlindDetectionMCG_UE_Mixed_r16 = new(CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16)
		if err = ie.pdcch_BlindDetectionMCG_UE_Mixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionMCG_UE_Mixed_r16", err)
		}
	}
	if pdcch_BlindDetectionSCG_UE_Mixed_r16Present {
		ie.pdcch_BlindDetectionSCG_UE_Mixed_r16 = new(CA_ParametersNR_v1610_pdcch_BlindDetectionSCG_UE_Mixed_r16)
		if err = ie.pdcch_BlindDetectionSCG_UE_Mixed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionSCG_UE_Mixed_r16", err)
		}
	}
	if crossCarrierSchedulingDL_DiffSCS_r16Present {
		ie.crossCarrierSchedulingDL_DiffSCS_r16 = new(CA_ParametersNR_v1610_crossCarrierSchedulingDL_DiffSCS_r16)
		if err = ie.crossCarrierSchedulingDL_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingDL_DiffSCS_r16", err)
		}
	}
	if crossCarrierSchedulingDefaultQCL_r16Present {
		ie.crossCarrierSchedulingDefaultQCL_r16 = new(CA_ParametersNR_v1610_crossCarrierSchedulingDefaultQCL_r16)
		if err = ie.crossCarrierSchedulingDefaultQCL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingDefaultQCL_r16", err)
		}
	}
	if crossCarrierSchedulingUL_DiffSCS_r16Present {
		ie.crossCarrierSchedulingUL_DiffSCS_r16 = new(CA_ParametersNR_v1610_crossCarrierSchedulingUL_DiffSCS_r16)
		if err = ie.crossCarrierSchedulingUL_DiffSCS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierSchedulingUL_DiffSCS_r16", err)
		}
	}
	if simul_SRS_MIMO_Trans_BC_r16Present {
		ie.simul_SRS_MIMO_Trans_BC_r16 = new(CA_ParametersNR_v1610_simul_SRS_MIMO_Trans_BC_r16)
		if err = ie.simul_SRS_MIMO_Trans_BC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode simul_SRS_MIMO_Trans_BC_r16", err)
		}
	}
	if codebookParametersAdditionPerBC_r16Present {
		ie.codebookParametersAdditionPerBC_r16 = new(CodebookParametersAdditionPerBC_r16)
		if err = ie.codebookParametersAdditionPerBC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode codebookParametersAdditionPerBC_r16", err)
		}
	}
	if codebookComboParametersAdditionPerBC_r16Present {
		ie.codebookComboParametersAdditionPerBC_r16 = new(CodebookComboParametersAdditionPerBC_r16)
		if err = ie.codebookComboParametersAdditionPerBC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode codebookComboParametersAdditionPerBC_r16", err)
		}
	}
	return nil
}
