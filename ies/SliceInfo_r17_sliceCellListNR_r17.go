package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SliceInfo_r17_sliceCellListNR_r17_Choice_nothing uint64 = iota
	SliceInfo_r17_sliceCellListNR_r17_Choice_sliceAllowedCellListNR_r17
	SliceInfo_r17_sliceCellListNR_r17_Choice_sliceExcludedCellListNR_r17
)

type SliceInfo_r17_sliceCellListNR_r17 struct {
	Choice                      uint64
	sliceAllowedCellListNR_r17  *SliceCellListNR_r17
	sliceExcludedCellListNR_r17 *SliceCellListNR_r17
}

func (ie *SliceInfo_r17_sliceCellListNR_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SliceInfo_r17_sliceCellListNR_r17_Choice_sliceAllowedCellListNR_r17:
		if err = ie.sliceAllowedCellListNR_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode sliceAllowedCellListNR_r17", err)
		}
	case SliceInfo_r17_sliceCellListNR_r17_Choice_sliceExcludedCellListNR_r17:
		if err = ie.sliceExcludedCellListNR_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode sliceExcludedCellListNR_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SliceInfo_r17_sliceCellListNR_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SliceInfo_r17_sliceCellListNR_r17_Choice_sliceAllowedCellListNR_r17:
		ie.sliceAllowedCellListNR_r17 = new(SliceCellListNR_r17)
		if err = ie.sliceAllowedCellListNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sliceAllowedCellListNR_r17", err)
		}
	case SliceInfo_r17_sliceCellListNR_r17_Choice_sliceExcludedCellListNR_r17:
		ie.sliceExcludedCellListNR_r17 = new(SliceCellListNR_r17)
		if err = ie.sliceExcludedCellListNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sliceExcludedCellListNR_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
