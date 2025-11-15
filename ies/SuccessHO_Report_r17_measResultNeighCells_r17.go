package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17_measResultNeighCells_r17 struct {
	measResultListNR_r17    *MeasResultList2NR_r16    `optional`
	measResultListEUTRA_r17 *MeasResultList2EUTRA_r16 `optional`
}

func (ie *SuccessHO_Report_r17_measResultNeighCells_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultListNR_r17 != nil, ie.measResultListEUTRA_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResultListNR_r17 != nil {
		if err = ie.measResultListNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measResultListNR_r17", err)
		}
	}
	if ie.measResultListEUTRA_r17 != nil {
		if err = ie.measResultListEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measResultListEUTRA_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17_measResultNeighCells_r17) Decode(r *uper.UperReader) error {
	var err error
	var measResultListNR_r17Present bool
	if measResultListNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultListEUTRA_r17Present bool
	if measResultListEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measResultListNR_r17Present {
		ie.measResultListNR_r17 = new(MeasResultList2NR_r16)
		if err = ie.measResultListNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListNR_r17", err)
		}
	}
	if measResultListEUTRA_r17Present {
		ie.measResultListEUTRA_r17 = new(MeasResultList2EUTRA_r16)
		if err = ie.measResultListEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListEUTRA_r17", err)
		}
	}
	return nil
}
