package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellsTriggeredItem_Choice_nothing uint64 = iota
	CellsTriggeredItem_Choice_physCellId
	CellsTriggeredItem_Choice_physCellIdEUTRA
	CellsTriggeredItem_Choice_physCellIdUTRA_FDD_r16
)

type CellsTriggeredItem struct {
	Choice                 uint64
	physCellId             *PhysCellId
	physCellIdEUTRA        *EUTRA_PhysCellId
	physCellIdUTRA_FDD_r16 *PhysCellIdUTRA_FDD_r16
}

func (ie *CellsTriggeredItem) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CellsTriggeredItem_Choice_physCellId:
		if err = ie.physCellId.Encode(w); err != nil {
			err = utils.WrapError("Encode physCellId", err)
		}
	case CellsTriggeredItem_Choice_physCellIdEUTRA:
		if err = ie.physCellIdEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode physCellIdEUTRA", err)
		}
	case CellsTriggeredItem_Choice_physCellIdUTRA_FDD_r16:
		if err = ie.physCellIdUTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode physCellIdUTRA_FDD_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CellsTriggeredItem) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CellsTriggeredItem_Choice_physCellId:
		ie.physCellId = new(PhysCellId)
		if err = ie.physCellId.Decode(r); err != nil {
			return utils.WrapError("Decode physCellId", err)
		}
	case CellsTriggeredItem_Choice_physCellIdEUTRA:
		ie.physCellIdEUTRA = new(EUTRA_PhysCellId)
		if err = ie.physCellIdEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode physCellIdEUTRA", err)
		}
	case CellsTriggeredItem_Choice_physCellIdUTRA_FDD_r16:
		ie.physCellIdUTRA_FDD_r16 = new(PhysCellIdUTRA_FDD_r16)
		if err = ie.physCellIdUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode physCellIdUTRA_FDD_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
