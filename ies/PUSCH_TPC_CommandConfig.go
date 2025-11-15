package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TPC_CommandConfig struct {
	tpc_Index    *int64         `lb:1,ub:15,optional`
	tpc_IndexSUL *int64         `lb:1,ub:15,optional`
	targetCell   *ServCellIndex `optional`
}

func (ie *PUSCH_TPC_CommandConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.tpc_Index != nil, ie.tpc_IndexSUL != nil, ie.targetCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tpc_Index != nil {
		if err = w.WriteInteger(*ie.tpc_Index, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode tpc_Index", err)
		}
	}
	if ie.tpc_IndexSUL != nil {
		if err = w.WriteInteger(*ie.tpc_IndexSUL, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode tpc_IndexSUL", err)
		}
	}
	if ie.targetCell != nil {
		if err = ie.targetCell.Encode(w); err != nil {
			return utils.WrapError("Encode targetCell", err)
		}
	}
	return nil
}

func (ie *PUSCH_TPC_CommandConfig) Decode(r *uper.UperReader) error {
	var err error
	var tpc_IndexPresent bool
	if tpc_IndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_IndexSULPresent bool
	if tpc_IndexSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var targetCellPresent bool
	if targetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if tpc_IndexPresent {
		var tmp_int_tpc_Index int64
		if tmp_int_tpc_Index, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode tpc_Index", err)
		}
		ie.tpc_Index = &tmp_int_tpc_Index
	}
	if tpc_IndexSULPresent {
		var tmp_int_tpc_IndexSUL int64
		if tmp_int_tpc_IndexSUL, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode tpc_IndexSUL", err)
		}
		ie.tpc_IndexSUL = &tmp_int_tpc_IndexSUL
	}
	if targetCellPresent {
		ie.targetCell = new(ServCellIndex)
		if err = ie.targetCell.Decode(r); err != nil {
			return utils.WrapError("Decode targetCell", err)
		}
	}
	return nil
}
