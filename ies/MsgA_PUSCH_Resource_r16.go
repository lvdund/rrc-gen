package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_PUSCH_Resource_r16 struct {
	msgA_MCS_r16                         int64                                                       `lb:0,ub:15,madatory`
	nrofSlotsMsgA_PUSCH_r16              int64                                                       `lb:1,ub:4,madatory`
	nrofMsgA_PO_PerSlot_r16              MsgA_PUSCH_Resource_r16_nrofMsgA_PO_PerSlot_r16             `madatory`
	msgA_PUSCH_TimeDomainOffset_r16      int64                                                       `lb:1,ub:32,madatory`
	msgA_PUSCH_TimeDomainAllocation_r16  *int64                                                      `lb:1,ub:maxNrofUL_Allocations,optional`
	startSymbolAndLengthMsgA_PO_r16      *int64                                                      `lb:0,ub:127,optional`
	mappingTypeMsgA_PUSCH_r16            *MsgA_PUSCH_Resource_r16_mappingTypeMsgA_PUSCH_r16          `optional`
	guardPeriodMsgA_PUSCH_r16            *int64                                                      `lb:0,ub:3,optional`
	guardBandMsgA_PUSCH_r16              int64                                                       `lb:0,ub:1,madatory`
	frequencyStartMsgA_PUSCH_r16         int64                                                       `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	nrofPRBs_PerMsgA_PO_r16              int64                                                       `lb:1,ub:32,madatory`
	nrofMsgA_PO_FDM_r16                  MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16                 `madatory`
	msgA_IntraSlotFrequencyHopping_r16   *MsgA_PUSCH_Resource_r16_msgA_IntraSlotFrequencyHopping_r16 `optional`
	msgA_HoppingBits_r16                 *uper.BitString                                             `lb:2,ub:2,optional`
	msgA_DMRS_Config_r16                 MsgA_DMRS_Config_r16                                        `madatory`
	nrofDMRS_Sequences_r16               int64                                                       `lb:1,ub:2,madatory`
	msgA_Alpha_r16                       *MsgA_PUSCH_Resource_r16_msgA_Alpha_r16                     `optional`
	interlaceIndexFirstPO_MsgA_PUSCH_r16 *int64                                                      `lb:1,ub:10,optional`
	nrofInterlacesPerMsgA_PO_r16         *int64                                                      `lb:1,ub:10,optional`
}

func (ie *MsgA_PUSCH_Resource_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.msgA_PUSCH_TimeDomainAllocation_r16 != nil, ie.startSymbolAndLengthMsgA_PO_r16 != nil, ie.mappingTypeMsgA_PUSCH_r16 != nil, ie.guardPeriodMsgA_PUSCH_r16 != nil, ie.msgA_IntraSlotFrequencyHopping_r16 != nil, ie.msgA_HoppingBits_r16 != nil, ie.msgA_Alpha_r16 != nil, ie.interlaceIndexFirstPO_MsgA_PUSCH_r16 != nil, ie.nrofInterlacesPerMsgA_PO_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.msgA_MCS_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger msgA_MCS_r16", err)
	}
	if err = w.WriteInteger(ie.nrofSlotsMsgA_PUSCH_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger nrofSlotsMsgA_PUSCH_r16", err)
	}
	if err = ie.nrofMsgA_PO_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode nrofMsgA_PO_PerSlot_r16", err)
	}
	if err = w.WriteInteger(ie.msgA_PUSCH_TimeDomainOffset_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger msgA_PUSCH_TimeDomainOffset_r16", err)
	}
	if ie.msgA_PUSCH_TimeDomainAllocation_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_PUSCH_TimeDomainAllocation_r16, &uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
			return utils.WrapError("Encode msgA_PUSCH_TimeDomainAllocation_r16", err)
		}
	}
	if ie.startSymbolAndLengthMsgA_PO_r16 != nil {
		if err = w.WriteInteger(*ie.startSymbolAndLengthMsgA_PO_r16, &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Encode startSymbolAndLengthMsgA_PO_r16", err)
		}
	}
	if ie.mappingTypeMsgA_PUSCH_r16 != nil {
		if err = ie.mappingTypeMsgA_PUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mappingTypeMsgA_PUSCH_r16", err)
		}
	}
	if ie.guardPeriodMsgA_PUSCH_r16 != nil {
		if err = w.WriteInteger(*ie.guardPeriodMsgA_PUSCH_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode guardPeriodMsgA_PUSCH_r16", err)
		}
	}
	if err = w.WriteInteger(ie.guardBandMsgA_PUSCH_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteInteger guardBandMsgA_PUSCH_r16", err)
	}
	if err = w.WriteInteger(ie.frequencyStartMsgA_PUSCH_r16, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger frequencyStartMsgA_PUSCH_r16", err)
	}
	if err = w.WriteInteger(ie.nrofPRBs_PerMsgA_PO_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger nrofPRBs_PerMsgA_PO_r16", err)
	}
	if err = ie.nrofMsgA_PO_FDM_r16.Encode(w); err != nil {
		return utils.WrapError("Encode nrofMsgA_PO_FDM_r16", err)
	}
	if ie.msgA_IntraSlotFrequencyHopping_r16 != nil {
		if err = ie.msgA_IntraSlotFrequencyHopping_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_IntraSlotFrequencyHopping_r16", err)
		}
	}
	if ie.msgA_HoppingBits_r16 != nil {
		if err = w.WriteBitString(ie.msgA_HoppingBits_r16.Bytes, uint(ie.msgA_HoppingBits_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode msgA_HoppingBits_r16", err)
		}
	}
	if err = ie.msgA_DMRS_Config_r16.Encode(w); err != nil {
		return utils.WrapError("Encode msgA_DMRS_Config_r16", err)
	}
	if err = w.WriteInteger(ie.nrofDMRS_Sequences_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger nrofDMRS_Sequences_r16", err)
	}
	if ie.msgA_Alpha_r16 != nil {
		if err = ie.msgA_Alpha_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_Alpha_r16", err)
		}
	}
	if ie.interlaceIndexFirstPO_MsgA_PUSCH_r16 != nil {
		if err = w.WriteInteger(*ie.interlaceIndexFirstPO_MsgA_PUSCH_r16, &uper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode interlaceIndexFirstPO_MsgA_PUSCH_r16", err)
		}
	}
	if ie.nrofInterlacesPerMsgA_PO_r16 != nil {
		if err = w.WriteInteger(*ie.nrofInterlacesPerMsgA_PO_r16, &uper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode nrofInterlacesPerMsgA_PO_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_PUSCH_Resource_r16) Decode(r *uper.UperReader) error {
	var err error
	var msgA_PUSCH_TimeDomainAllocation_r16Present bool
	if msgA_PUSCH_TimeDomainAllocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var startSymbolAndLengthMsgA_PO_r16Present bool
	if startSymbolAndLengthMsgA_PO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mappingTypeMsgA_PUSCH_r16Present bool
	if mappingTypeMsgA_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var guardPeriodMsgA_PUSCH_r16Present bool
	if guardPeriodMsgA_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_IntraSlotFrequencyHopping_r16Present bool
	if msgA_IntraSlotFrequencyHopping_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_HoppingBits_r16Present bool
	if msgA_HoppingBits_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_Alpha_r16Present bool
	if msgA_Alpha_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interlaceIndexFirstPO_MsgA_PUSCH_r16Present bool
	if interlaceIndexFirstPO_MsgA_PUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofInterlacesPerMsgA_PO_r16Present bool
	if nrofInterlacesPerMsgA_PO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_msgA_MCS_r16 int64
	if tmp_int_msgA_MCS_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger msgA_MCS_r16", err)
	}
	ie.msgA_MCS_r16 = tmp_int_msgA_MCS_r16
	var tmp_int_nrofSlotsMsgA_PUSCH_r16 int64
	if tmp_int_nrofSlotsMsgA_PUSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger nrofSlotsMsgA_PUSCH_r16", err)
	}
	ie.nrofSlotsMsgA_PUSCH_r16 = tmp_int_nrofSlotsMsgA_PUSCH_r16
	if err = ie.nrofMsgA_PO_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode nrofMsgA_PO_PerSlot_r16", err)
	}
	var tmp_int_msgA_PUSCH_TimeDomainOffset_r16 int64
	if tmp_int_msgA_PUSCH_TimeDomainOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger msgA_PUSCH_TimeDomainOffset_r16", err)
	}
	ie.msgA_PUSCH_TimeDomainOffset_r16 = tmp_int_msgA_PUSCH_TimeDomainOffset_r16
	if msgA_PUSCH_TimeDomainAllocation_r16Present {
		var tmp_int_msgA_PUSCH_TimeDomainAllocation_r16 int64
		if tmp_int_msgA_PUSCH_TimeDomainAllocation_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false); err != nil {
			return utils.WrapError("Decode msgA_PUSCH_TimeDomainAllocation_r16", err)
		}
		ie.msgA_PUSCH_TimeDomainAllocation_r16 = &tmp_int_msgA_PUSCH_TimeDomainAllocation_r16
	}
	if startSymbolAndLengthMsgA_PO_r16Present {
		var tmp_int_startSymbolAndLengthMsgA_PO_r16 int64
		if tmp_int_startSymbolAndLengthMsgA_PO_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode startSymbolAndLengthMsgA_PO_r16", err)
		}
		ie.startSymbolAndLengthMsgA_PO_r16 = &tmp_int_startSymbolAndLengthMsgA_PO_r16
	}
	if mappingTypeMsgA_PUSCH_r16Present {
		ie.mappingTypeMsgA_PUSCH_r16 = new(MsgA_PUSCH_Resource_r16_mappingTypeMsgA_PUSCH_r16)
		if err = ie.mappingTypeMsgA_PUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mappingTypeMsgA_PUSCH_r16", err)
		}
	}
	if guardPeriodMsgA_PUSCH_r16Present {
		var tmp_int_guardPeriodMsgA_PUSCH_r16 int64
		if tmp_int_guardPeriodMsgA_PUSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode guardPeriodMsgA_PUSCH_r16", err)
		}
		ie.guardPeriodMsgA_PUSCH_r16 = &tmp_int_guardPeriodMsgA_PUSCH_r16
	}
	var tmp_int_guardBandMsgA_PUSCH_r16 int64
	if tmp_int_guardBandMsgA_PUSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadInteger guardBandMsgA_PUSCH_r16", err)
	}
	ie.guardBandMsgA_PUSCH_r16 = tmp_int_guardBandMsgA_PUSCH_r16
	var tmp_int_frequencyStartMsgA_PUSCH_r16 int64
	if tmp_int_frequencyStartMsgA_PUSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger frequencyStartMsgA_PUSCH_r16", err)
	}
	ie.frequencyStartMsgA_PUSCH_r16 = tmp_int_frequencyStartMsgA_PUSCH_r16
	var tmp_int_nrofPRBs_PerMsgA_PO_r16 int64
	if tmp_int_nrofPRBs_PerMsgA_PO_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger nrofPRBs_PerMsgA_PO_r16", err)
	}
	ie.nrofPRBs_PerMsgA_PO_r16 = tmp_int_nrofPRBs_PerMsgA_PO_r16
	if err = ie.nrofMsgA_PO_FDM_r16.Decode(r); err != nil {
		return utils.WrapError("Decode nrofMsgA_PO_FDM_r16", err)
	}
	if msgA_IntraSlotFrequencyHopping_r16Present {
		ie.msgA_IntraSlotFrequencyHopping_r16 = new(MsgA_PUSCH_Resource_r16_msgA_IntraSlotFrequencyHopping_r16)
		if err = ie.msgA_IntraSlotFrequencyHopping_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_IntraSlotFrequencyHopping_r16", err)
		}
	}
	if msgA_HoppingBits_r16Present {
		var tmp_bs_msgA_HoppingBits_r16 []byte
		var n_msgA_HoppingBits_r16 uint
		if tmp_bs_msgA_HoppingBits_r16, n_msgA_HoppingBits_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode msgA_HoppingBits_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_msgA_HoppingBits_r16,
			NumBits: uint64(n_msgA_HoppingBits_r16),
		}
		ie.msgA_HoppingBits_r16 = &tmp_bitstring
	}
	if err = ie.msgA_DMRS_Config_r16.Decode(r); err != nil {
		return utils.WrapError("Decode msgA_DMRS_Config_r16", err)
	}
	var tmp_int_nrofDMRS_Sequences_r16 int64
	if tmp_int_nrofDMRS_Sequences_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger nrofDMRS_Sequences_r16", err)
	}
	ie.nrofDMRS_Sequences_r16 = tmp_int_nrofDMRS_Sequences_r16
	if msgA_Alpha_r16Present {
		ie.msgA_Alpha_r16 = new(MsgA_PUSCH_Resource_r16_msgA_Alpha_r16)
		if err = ie.msgA_Alpha_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_Alpha_r16", err)
		}
	}
	if interlaceIndexFirstPO_MsgA_PUSCH_r16Present {
		var tmp_int_interlaceIndexFirstPO_MsgA_PUSCH_r16 int64
		if tmp_int_interlaceIndexFirstPO_MsgA_PUSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode interlaceIndexFirstPO_MsgA_PUSCH_r16", err)
		}
		ie.interlaceIndexFirstPO_MsgA_PUSCH_r16 = &tmp_int_interlaceIndexFirstPO_MsgA_PUSCH_r16
	}
	if nrofInterlacesPerMsgA_PO_r16Present {
		var tmp_int_nrofInterlacesPerMsgA_PO_r16 int64
		if tmp_int_nrofInterlacesPerMsgA_PO_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode nrofInterlacesPerMsgA_PO_r16", err)
		}
		ie.nrofInterlacesPerMsgA_PO_r16 = &tmp_int_nrofInterlacesPerMsgA_PO_r16
	}
	return nil
}
