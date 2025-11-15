package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_DMRS_Config_r16 struct {
	msgA_DMRS_AdditionalPosition_r16 *MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16 `optional`
	msgA_MaxLength_r16               *MsgA_DMRS_Config_r16_msgA_MaxLength_r16               `optional`
	msgA_PUSCH_DMRS_CDM_Group_r16    *int64                                                 `lb:0,ub:1,optional`
	msgA_PUSCH_NrofPorts_r16         *int64                                                 `lb:0,ub:1,optional`
	msgA_ScramblingID0_r16           *int64                                                 `lb:0,ub:65535,optional`
	msgA_ScramblingID1_r16           *int64                                                 `lb:0,ub:65535,optional`
}

func (ie *MsgA_DMRS_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.msgA_DMRS_AdditionalPosition_r16 != nil, ie.msgA_MaxLength_r16 != nil, ie.msgA_PUSCH_DMRS_CDM_Group_r16 != nil, ie.msgA_PUSCH_NrofPorts_r16 != nil, ie.msgA_ScramblingID0_r16 != nil, ie.msgA_ScramblingID1_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.msgA_DMRS_AdditionalPosition_r16 != nil {
		if err = ie.msgA_DMRS_AdditionalPosition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_DMRS_AdditionalPosition_r16", err)
		}
	}
	if ie.msgA_MaxLength_r16 != nil {
		if err = ie.msgA_MaxLength_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_MaxLength_r16", err)
		}
	}
	if ie.msgA_PUSCH_DMRS_CDM_Group_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_PUSCH_DMRS_CDM_Group_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode msgA_PUSCH_DMRS_CDM_Group_r16", err)
		}
	}
	if ie.msgA_PUSCH_NrofPorts_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_PUSCH_NrofPorts_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode msgA_PUSCH_NrofPorts_r16", err)
		}
	}
	if ie.msgA_ScramblingID0_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_ScramblingID0_r16, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode msgA_ScramblingID0_r16", err)
		}
	}
	if ie.msgA_ScramblingID1_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_ScramblingID1_r16, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode msgA_ScramblingID1_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_DMRS_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var msgA_DMRS_AdditionalPosition_r16Present bool
	if msgA_DMRS_AdditionalPosition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_MaxLength_r16Present bool
	if msgA_MaxLength_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PUSCH_DMRS_CDM_Group_r16Present bool
	if msgA_PUSCH_DMRS_CDM_Group_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PUSCH_NrofPorts_r16Present bool
	if msgA_PUSCH_NrofPorts_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_ScramblingID0_r16Present bool
	if msgA_ScramblingID0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_ScramblingID1_r16Present bool
	if msgA_ScramblingID1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if msgA_DMRS_AdditionalPosition_r16Present {
		ie.msgA_DMRS_AdditionalPosition_r16 = new(MsgA_DMRS_Config_r16_msgA_DMRS_AdditionalPosition_r16)
		if err = ie.msgA_DMRS_AdditionalPosition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_DMRS_AdditionalPosition_r16", err)
		}
	}
	if msgA_MaxLength_r16Present {
		ie.msgA_MaxLength_r16 = new(MsgA_DMRS_Config_r16_msgA_MaxLength_r16)
		if err = ie.msgA_MaxLength_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_MaxLength_r16", err)
		}
	}
	if msgA_PUSCH_DMRS_CDM_Group_r16Present {
		var tmp_int_msgA_PUSCH_DMRS_CDM_Group_r16 int64
		if tmp_int_msgA_PUSCH_DMRS_CDM_Group_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode msgA_PUSCH_DMRS_CDM_Group_r16", err)
		}
		ie.msgA_PUSCH_DMRS_CDM_Group_r16 = &tmp_int_msgA_PUSCH_DMRS_CDM_Group_r16
	}
	if msgA_PUSCH_NrofPorts_r16Present {
		var tmp_int_msgA_PUSCH_NrofPorts_r16 int64
		if tmp_int_msgA_PUSCH_NrofPorts_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode msgA_PUSCH_NrofPorts_r16", err)
		}
		ie.msgA_PUSCH_NrofPorts_r16 = &tmp_int_msgA_PUSCH_NrofPorts_r16
	}
	if msgA_ScramblingID0_r16Present {
		var tmp_int_msgA_ScramblingID0_r16 int64
		if tmp_int_msgA_ScramblingID0_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode msgA_ScramblingID0_r16", err)
		}
		ie.msgA_ScramblingID0_r16 = &tmp_int_msgA_ScramblingID0_r16
	}
	if msgA_ScramblingID1_r16Present {
		var tmp_int_msgA_ScramblingID1_r16 int64
		if tmp_int_msgA_ScramblingID1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode msgA_ScramblingID1_r16", err)
		}
		ie.msgA_ScramblingID1_r16 = &tmp_int_msgA_ScramblingID1_r16
	}
	return nil
}
