package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_FreqNeighCellInfo struct {
	physCellId           EUTRA_PhysCellId    `madatory`
	dummy                EUTRA_Q_OffsetRange `madatory`
	q_RxLevMinOffsetCell *int64              `lb:1,ub:8,optional`
	q_QualMinOffsetCell  *int64              `lb:1,ub:8,optional`
}

func (ie *EUTRA_FreqNeighCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.q_RxLevMinOffsetCell != nil, ie.q_QualMinOffsetCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.physCellId.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId", err)
	}
	if err = ie.dummy.Encode(w); err != nil {
		return utils.WrapError("Encode dummy", err)
	}
	if ie.q_RxLevMinOffsetCell != nil {
		if err = w.WriteInteger(*ie.q_RxLevMinOffsetCell, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode q_RxLevMinOffsetCell", err)
		}
	}
	if ie.q_QualMinOffsetCell != nil {
		if err = w.WriteInteger(*ie.q_QualMinOffsetCell, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode q_QualMinOffsetCell", err)
		}
	}
	return nil
}

func (ie *EUTRA_FreqNeighCellInfo) Decode(r *uper.UperReader) error {
	var err error
	var q_RxLevMinOffsetCellPresent bool
	if q_RxLevMinOffsetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_QualMinOffsetCellPresent bool
	if q_QualMinOffsetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.physCellId.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId", err)
	}
	if err = ie.dummy.Decode(r); err != nil {
		return utils.WrapError("Decode dummy", err)
	}
	if q_RxLevMinOffsetCellPresent {
		var tmp_int_q_RxLevMinOffsetCell int64
		if tmp_int_q_RxLevMinOffsetCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode q_RxLevMinOffsetCell", err)
		}
		ie.q_RxLevMinOffsetCell = &tmp_int_q_RxLevMinOffsetCell
	}
	if q_QualMinOffsetCellPresent {
		var tmp_int_q_QualMinOffsetCell int64
		if tmp_int_q_QualMinOffsetCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode q_QualMinOffsetCell", err)
		}
		ie.q_QualMinOffsetCell = &tmp_int_q_QualMinOffsetCell
	}
	return nil
}
