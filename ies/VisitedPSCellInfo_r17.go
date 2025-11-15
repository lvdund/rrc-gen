package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VisitedPSCellInfo_r17 struct {
	visitedCellId_r17 *VisitedPSCellInfo_r17_visitedCellId_r17 `optional`
	timeSpent_r17     int64                                    `lb:0,ub:4095,madatory`
}

func (ie *VisitedPSCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.visitedCellId_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.visitedCellId_r17 != nil {
		if err = ie.visitedCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode visitedCellId_r17", err)
		}
	}
	if err = w.WriteInteger(ie.timeSpent_r17, &uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("WriteInteger timeSpent_r17", err)
	}
	return nil
}

func (ie *VisitedPSCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var visitedCellId_r17Present bool
	if visitedCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if visitedCellId_r17Present {
		ie.visitedCellId_r17 = new(VisitedPSCellInfo_r17_visitedCellId_r17)
		if err = ie.visitedCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode visitedCellId_r17", err)
		}
	}
	var tmp_int_timeSpent_r17 int64
	if tmp_int_timeSpent_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("ReadInteger timeSpent_r17", err)
	}
	ie.timeSpent_r17 = tmp_int_timeSpent_r17
	return nil
}
