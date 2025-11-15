package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_PhysCellIdRange struct {
	start      EUTRA_PhysCellId                  `madatory`
	range_cell *EUTRA_PhysCellIdRange_range_cell `optional`
}

func (ie *EUTRA_PhysCellIdRange) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.range_cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.start.Encode(w); err != nil {
		return utils.WrapError("Encode start", err)
	}
	if ie.range_cell != nil {
		if err = ie.range_cell.Encode(w); err != nil {
			return utils.WrapError("Encode range_cell", err)
		}
	}
	return nil
}

func (ie *EUTRA_PhysCellIdRange) Decode(r *uper.UperReader) error {
	var err error
	var range_cellPresent bool
	if range_cellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.start.Decode(r); err != nil {
		return utils.WrapError("Decode start", err)
	}
	if range_cellPresent {
		ie.range_cell = new(EUTRA_PhysCellIdRange_range_cell)
		if err = ie.range_cell.Decode(r); err != nil {
			return utils.WrapError("Decode range_cell", err)
		}
	}
	return nil
}
