package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureCombinationPreambles_r17 struct {
	featureCombination_r17                       FeatureCombination_r17                                `madatory`
	startPreambleForThisPartition_r17            int64                                                 `lb:0,ub:63,madatory`
	numberOfPreamblesPerSSB_ForThisPartition_r17 int64                                                 `lb:1,ub:64,madatory`
	ssb_SharedRO_MaskIndex_r17                   *int64                                                `lb:1,ub:15,optional`
	groupBconfigured_r17                         *FeatureCombinationPreambles_r17_groupBconfigured_r17 `lb:1,ub:64,optional`
	separateMsgA_PUSCH_Config_r17                *MsgA_PUSCH_Config_r16                                `optional`
	msgA_RSRP_Threshold_r17                      *RSRP_Range                                           `optional`
	rsrp_ThresholdSSB_r17                        *RSRP_Range                                           `optional`
	deltaPreamble_r17                            *int64                                                `lb:-1,ub:6,optional`
}

func (ie *FeatureCombinationPreambles_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_SharedRO_MaskIndex_r17 != nil, ie.groupBconfigured_r17 != nil, ie.separateMsgA_PUSCH_Config_r17 != nil, ie.msgA_RSRP_Threshold_r17 != nil, ie.rsrp_ThresholdSSB_r17 != nil, ie.deltaPreamble_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.featureCombination_r17.Encode(w); err != nil {
		return utils.WrapError("Encode featureCombination_r17", err)
	}
	if err = w.WriteInteger(ie.startPreambleForThisPartition_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger startPreambleForThisPartition_r17", err)
	}
	if err = w.WriteInteger(ie.numberOfPreamblesPerSSB_ForThisPartition_r17, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfPreamblesPerSSB_ForThisPartition_r17", err)
	}
	if ie.ssb_SharedRO_MaskIndex_r17 != nil {
		if err = w.WriteInteger(*ie.ssb_SharedRO_MaskIndex_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode ssb_SharedRO_MaskIndex_r17", err)
		}
	}
	if ie.groupBconfigured_r17 != nil {
		if err = ie.groupBconfigured_r17.Encode(w); err != nil {
			return utils.WrapError("Encode groupBconfigured_r17", err)
		}
	}
	if ie.separateMsgA_PUSCH_Config_r17 != nil {
		if err = ie.separateMsgA_PUSCH_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode separateMsgA_PUSCH_Config_r17", err)
		}
	}
	if ie.msgA_RSRP_Threshold_r17 != nil {
		if err = ie.msgA_RSRP_Threshold_r17.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_RSRP_Threshold_r17", err)
		}
	}
	if ie.rsrp_ThresholdSSB_r17 != nil {
		if err = ie.rsrp_ThresholdSSB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_ThresholdSSB_r17", err)
		}
	}
	if ie.deltaPreamble_r17 != nil {
		if err = w.WriteInteger(*ie.deltaPreamble_r17, &uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode deltaPreamble_r17", err)
		}
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17) Decode(r *uper.UperReader) error {
	var err error
	var ssb_SharedRO_MaskIndex_r17Present bool
	if ssb_SharedRO_MaskIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var groupBconfigured_r17Present bool
	if groupBconfigured_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var separateMsgA_PUSCH_Config_r17Present bool
	if separateMsgA_PUSCH_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_RSRP_Threshold_r17Present bool
	if msgA_RSRP_Threshold_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrp_ThresholdSSB_r17Present bool
	if rsrp_ThresholdSSB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaPreamble_r17Present bool
	if deltaPreamble_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.featureCombination_r17.Decode(r); err != nil {
		return utils.WrapError("Decode featureCombination_r17", err)
	}
	var tmp_int_startPreambleForThisPartition_r17 int64
	if tmp_int_startPreambleForThisPartition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger startPreambleForThisPartition_r17", err)
	}
	ie.startPreambleForThisPartition_r17 = tmp_int_startPreambleForThisPartition_r17
	var tmp_int_numberOfPreamblesPerSSB_ForThisPartition_r17 int64
	if tmp_int_numberOfPreamblesPerSSB_ForThisPartition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfPreamblesPerSSB_ForThisPartition_r17", err)
	}
	ie.numberOfPreamblesPerSSB_ForThisPartition_r17 = tmp_int_numberOfPreamblesPerSSB_ForThisPartition_r17
	if ssb_SharedRO_MaskIndex_r17Present {
		var tmp_int_ssb_SharedRO_MaskIndex_r17 int64
		if tmp_int_ssb_SharedRO_MaskIndex_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode ssb_SharedRO_MaskIndex_r17", err)
		}
		ie.ssb_SharedRO_MaskIndex_r17 = &tmp_int_ssb_SharedRO_MaskIndex_r17
	}
	if groupBconfigured_r17Present {
		ie.groupBconfigured_r17 = new(FeatureCombinationPreambles_r17_groupBconfigured_r17)
		if err = ie.groupBconfigured_r17.Decode(r); err != nil {
			return utils.WrapError("Decode groupBconfigured_r17", err)
		}
	}
	if separateMsgA_PUSCH_Config_r17Present {
		ie.separateMsgA_PUSCH_Config_r17 = new(MsgA_PUSCH_Config_r16)
		if err = ie.separateMsgA_PUSCH_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode separateMsgA_PUSCH_Config_r17", err)
		}
	}
	if msgA_RSRP_Threshold_r17Present {
		ie.msgA_RSRP_Threshold_r17 = new(RSRP_Range)
		if err = ie.msgA_RSRP_Threshold_r17.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_RSRP_Threshold_r17", err)
		}
	}
	if rsrp_ThresholdSSB_r17Present {
		ie.rsrp_ThresholdSSB_r17 = new(RSRP_Range)
		if err = ie.rsrp_ThresholdSSB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_ThresholdSSB_r17", err)
		}
	}
	if deltaPreamble_r17Present {
		var tmp_int_deltaPreamble_r17 int64
		if tmp_int_deltaPreamble_r17, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode deltaPreamble_r17", err)
		}
		ie.deltaPreamble_r17 = &tmp_int_deltaPreamble_r17
	}
	return nil
}
