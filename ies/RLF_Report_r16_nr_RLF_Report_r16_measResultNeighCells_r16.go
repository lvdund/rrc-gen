package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16 struct {
	measResultListNR_r16    *MeasResultList2NR_r16    `optional`
	measResultListEUTRA_r16 *MeasResultList2EUTRA_r16 `optional`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultListNR_r16 != nil, ie.measResultListEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResultListNR_r16 != nil {
		if err = ie.measResultListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultListNR_r16", err)
		}
	}
	if ie.measResultListEUTRA_r16 != nil {
		if err = ie.measResultListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultListEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResultListNR_r16Present bool
	if measResultListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultListEUTRA_r16Present bool
	if measResultListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measResultListNR_r16Present {
		ie.measResultListNR_r16 = new(MeasResultList2NR_r16)
		if err = ie.measResultListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListNR_r16", err)
		}
	}
	if measResultListEUTRA_r16Present {
		ie.measResultListEUTRA_r16 = new(MeasResultList2EUTRA_r16)
		if err = ie.measResultListEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultListEUTRA_r16", err)
		}
	}
	return nil
}
