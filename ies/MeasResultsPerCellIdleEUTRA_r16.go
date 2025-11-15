package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleEUTRA_r16 struct {
	eutra_PhysCellId_r16    EUTRA_PhysCellId                                         `madatory`
	measIdleResultEUTRA_r16 *MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16 `optional`
}

func (ie *MeasResultsPerCellIdleEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measIdleResultEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.eutra_PhysCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode eutra_PhysCellId_r16", err)
	}
	if ie.measIdleResultEUTRA_r16 != nil {
		if err = ie.measIdleResultEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measIdleResultEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var measIdleResultEUTRA_r16Present bool
	if measIdleResultEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.eutra_PhysCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode eutra_PhysCellId_r16", err)
	}
	if measIdleResultEUTRA_r16Present {
		ie.measIdleResultEUTRA_r16 = new(MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16)
		if err = ie.measIdleResultEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measIdleResultEUTRA_r16", err)
		}
	}
	return nil
}
