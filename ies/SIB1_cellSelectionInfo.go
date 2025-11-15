package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_cellSelectionInfo struct {
	q_RxLevMin       Q_RxLevMin  `madatory`
	q_RxLevMinOffset *int64      `lb:1,ub:8,optional`
	q_RxLevMinSUL    *Q_RxLevMin `optional`
	q_QualMin        *Q_QualMin  `optional`
	q_QualMinOffset  *int64      `lb:1,ub:8,optional`
}

func (ie *SIB1_cellSelectionInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.q_RxLevMinOffset != nil, ie.q_RxLevMinSUL != nil, ie.q_QualMin != nil, ie.q_QualMinOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.q_RxLevMin.Encode(w); err != nil {
		return utils.WrapError("Encode q_RxLevMin", err)
	}
	if ie.q_RxLevMinOffset != nil {
		if err = w.WriteInteger(*ie.q_RxLevMinOffset, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode q_RxLevMinOffset", err)
		}
	}
	if ie.q_RxLevMinSUL != nil {
		if err = ie.q_RxLevMinSUL.Encode(w); err != nil {
			return utils.WrapError("Encode q_RxLevMinSUL", err)
		}
	}
	if ie.q_QualMin != nil {
		if err = ie.q_QualMin.Encode(w); err != nil {
			return utils.WrapError("Encode q_QualMin", err)
		}
	}
	if ie.q_QualMinOffset != nil {
		if err = w.WriteInteger(*ie.q_QualMinOffset, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode q_QualMinOffset", err)
		}
	}
	return nil
}

func (ie *SIB1_cellSelectionInfo) Decode(r *uper.UperReader) error {
	var err error
	var q_RxLevMinOffsetPresent bool
	if q_RxLevMinOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_RxLevMinSULPresent bool
	if q_RxLevMinSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_QualMinPresent bool
	if q_QualMinPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var q_QualMinOffsetPresent bool
	if q_QualMinOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.q_RxLevMin.Decode(r); err != nil {
		return utils.WrapError("Decode q_RxLevMin", err)
	}
	if q_RxLevMinOffsetPresent {
		var tmp_int_q_RxLevMinOffset int64
		if tmp_int_q_RxLevMinOffset, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode q_RxLevMinOffset", err)
		}
		ie.q_RxLevMinOffset = &tmp_int_q_RxLevMinOffset
	}
	if q_RxLevMinSULPresent {
		ie.q_RxLevMinSUL = new(Q_RxLevMin)
		if err = ie.q_RxLevMinSUL.Decode(r); err != nil {
			return utils.WrapError("Decode q_RxLevMinSUL", err)
		}
	}
	if q_QualMinPresent {
		ie.q_QualMin = new(Q_QualMin)
		if err = ie.q_QualMin.Decode(r); err != nil {
			return utils.WrapError("Decode q_QualMin", err)
		}
	}
	if q_QualMinOffsetPresent {
		var tmp_int_q_QualMinOffset int64
		if tmp_int_q_QualMinOffset, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode q_QualMinOffset", err)
		}
		ie.q_QualMinOffset = &tmp_int_q_QualMinOffset
	}
	return nil
}
