package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailReport_r16_measResultNeighCells_r16 struct {
	measResultNeighCellListNR    *MeasResultList2NR_r16    `optional`
	measResultNeighCellListEUTRA *MeasResultList2EUTRA_r16 `optional`
}

func (ie *ConnEstFailReport_r16_measResultNeighCells_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultNeighCellListNR != nil, ie.measResultNeighCellListEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResultNeighCellListNR != nil {
		if err = ie.measResultNeighCellListNR.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCellListNR", err)
		}
	}
	if ie.measResultNeighCellListEUTRA != nil {
		if err = ie.measResultNeighCellListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCellListEUTRA", err)
		}
	}
	return nil
}

func (ie *ConnEstFailReport_r16_measResultNeighCells_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResultNeighCellListNRPresent bool
	if measResultNeighCellListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultNeighCellListEUTRAPresent bool
	if measResultNeighCellListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measResultNeighCellListNRPresent {
		ie.measResultNeighCellListNR = new(MeasResultList2NR_r16)
		if err = ie.measResultNeighCellListNR.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCellListNR", err)
		}
	}
	if measResultNeighCellListEUTRAPresent {
		ie.measResultNeighCellListEUTRA = new(MeasResultList2EUTRA_r16)
		if err = ie.measResultNeighCellListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCellListEUTRA", err)
		}
	}
	return nil
}
