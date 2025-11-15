package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MsgA_PUSCH_Config_r16 struct {
	msgA_PUSCH_ResourceGroupA_r16 *MsgA_PUSCH_Resource_r16                          `optional`
	msgA_PUSCH_ResourceGroupB_r16 *MsgA_PUSCH_Resource_r16                          `optional`
	msgA_TransformPrecoder_r16    *MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16 `optional`
	msgA_DataScramblingIndex_r16  *int64                                            `lb:0,ub:1023,optional`
	msgA_DeltaPreamble_r16        *int64                                            `lb:-1,ub:6,optional`
}

func (ie *MsgA_PUSCH_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.msgA_PUSCH_ResourceGroupA_r16 != nil, ie.msgA_PUSCH_ResourceGroupB_r16 != nil, ie.msgA_TransformPrecoder_r16 != nil, ie.msgA_DataScramblingIndex_r16 != nil, ie.msgA_DeltaPreamble_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.msgA_PUSCH_ResourceGroupA_r16 != nil {
		if err = ie.msgA_PUSCH_ResourceGroupA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_PUSCH_ResourceGroupA_r16", err)
		}
	}
	if ie.msgA_PUSCH_ResourceGroupB_r16 != nil {
		if err = ie.msgA_PUSCH_ResourceGroupB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_PUSCH_ResourceGroupB_r16", err)
		}
	}
	if ie.msgA_TransformPrecoder_r16 != nil {
		if err = ie.msgA_TransformPrecoder_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_TransformPrecoder_r16", err)
		}
	}
	if ie.msgA_DataScramblingIndex_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_DataScramblingIndex_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode msgA_DataScramblingIndex_r16", err)
		}
	}
	if ie.msgA_DeltaPreamble_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_DeltaPreamble_r16, &uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode msgA_DeltaPreamble_r16", err)
		}
	}
	return nil
}

func (ie *MsgA_PUSCH_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var msgA_PUSCH_ResourceGroupA_r16Present bool
	if msgA_PUSCH_ResourceGroupA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PUSCH_ResourceGroupB_r16Present bool
	if msgA_PUSCH_ResourceGroupB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_TransformPrecoder_r16Present bool
	if msgA_TransformPrecoder_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_DataScramblingIndex_r16Present bool
	if msgA_DataScramblingIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_DeltaPreamble_r16Present bool
	if msgA_DeltaPreamble_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if msgA_PUSCH_ResourceGroupA_r16Present {
		ie.msgA_PUSCH_ResourceGroupA_r16 = new(MsgA_PUSCH_Resource_r16)
		if err = ie.msgA_PUSCH_ResourceGroupA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_PUSCH_ResourceGroupA_r16", err)
		}
	}
	if msgA_PUSCH_ResourceGroupB_r16Present {
		ie.msgA_PUSCH_ResourceGroupB_r16 = new(MsgA_PUSCH_Resource_r16)
		if err = ie.msgA_PUSCH_ResourceGroupB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_PUSCH_ResourceGroupB_r16", err)
		}
	}
	if msgA_TransformPrecoder_r16Present {
		ie.msgA_TransformPrecoder_r16 = new(MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16)
		if err = ie.msgA_TransformPrecoder_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_TransformPrecoder_r16", err)
		}
	}
	if msgA_DataScramblingIndex_r16Present {
		var tmp_int_msgA_DataScramblingIndex_r16 int64
		if tmp_int_msgA_DataScramblingIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode msgA_DataScramblingIndex_r16", err)
		}
		ie.msgA_DataScramblingIndex_r16 = &tmp_int_msgA_DataScramblingIndex_r16
	}
	if msgA_DeltaPreamble_r16Present {
		var tmp_int_msgA_DeltaPreamble_r16 int64
		if tmp_int_msgA_DeltaPreamble_r16, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode msgA_DeltaPreamble_r16", err)
		}
		ie.msgA_DeltaPreamble_r16 = &tmp_int_msgA_DeltaPreamble_r16
	}
	return nil
}
