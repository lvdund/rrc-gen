package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	VisitedCellInfo_r16_visitedCellId_r16_Choice_nothing uint64 = iota
	VisitedCellInfo_r16_visitedCellId_r16_Choice_nr_CellId_r16
	VisitedCellInfo_r16_visitedCellId_r16_Choice_eutra_CellId_r16
)

type VisitedCellInfo_r16_visitedCellId_r16 struct {
	Choice           uint64
	nr_CellId_r16    *VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16
	eutra_CellId_r16 *VisitedCellInfo_r16_visitedCellId_r16_eutra_CellId_r16
}

func (ie *VisitedCellInfo_r16_visitedCellId_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VisitedCellInfo_r16_visitedCellId_r16_Choice_nr_CellId_r16:
		if err = ie.nr_CellId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode nr_CellId_r16", err)
		}
	case VisitedCellInfo_r16_visitedCellId_r16_Choice_eutra_CellId_r16:
		if err = ie.eutra_CellId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra_CellId_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *VisitedCellInfo_r16_visitedCellId_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case VisitedCellInfo_r16_visitedCellId_r16_Choice_nr_CellId_r16:
		ie.nr_CellId_r16 = new(VisitedCellInfo_r16_visitedCellId_r16_nr_CellId_r16)
		if err = ie.nr_CellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode nr_CellId_r16", err)
		}
	case VisitedCellInfo_r16_visitedCellId_r16_Choice_eutra_CellId_r16:
		ie.eutra_CellId_r16 = new(VisitedCellInfo_r16_visitedCellId_r16_eutra_CellId_r16)
		if err = ie.eutra_CellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_CellId_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
