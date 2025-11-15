package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_ParametersXDD_Diff struct {
	rsrqMeasWidebandEUTRA *EUTRA_ParametersXDD_Diff_rsrqMeasWidebandEUTRA `optional`
}

func (ie *EUTRA_ParametersXDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rsrqMeasWidebandEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rsrqMeasWidebandEUTRA != nil {
		if err = ie.rsrqMeasWidebandEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode rsrqMeasWidebandEUTRA", err)
		}
	}
	return nil
}

func (ie *EUTRA_ParametersXDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var rsrqMeasWidebandEUTRAPresent bool
	if rsrqMeasWidebandEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if rsrqMeasWidebandEUTRAPresent {
		ie.rsrqMeasWidebandEUTRA = new(EUTRA_ParametersXDD_Diff_rsrqMeasWidebandEUTRA)
		if err = ie.rsrqMeasWidebandEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode rsrqMeasWidebandEUTRA", err)
		}
	}
	return nil
}
