package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_XDD_Diff struct {
	sftd_MeasPSCell  *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell  `optional`
	sftd_MeasNR_Cell *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell `optional`
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sftd_MeasPSCell != nil, ie.sftd_MeasNR_Cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sftd_MeasPSCell != nil {
		if err = ie.sftd_MeasPSCell.Encode(w); err != nil {
			return utils.WrapError("Encode sftd_MeasPSCell", err)
		}
	}
	if ie.sftd_MeasNR_Cell != nil {
		if err = ie.sftd_MeasNR_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode sftd_MeasNR_Cell", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var sftd_MeasPSCellPresent bool
	if sftd_MeasPSCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sftd_MeasNR_CellPresent bool
	if sftd_MeasNR_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sftd_MeasPSCellPresent {
		ie.sftd_MeasPSCell = new(MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell)
		if err = ie.sftd_MeasPSCell.Decode(r); err != nil {
			return utils.WrapError("Decode sftd_MeasPSCell", err)
		}
	}
	if sftd_MeasNR_CellPresent {
		ie.sftd_MeasNR_Cell = new(MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasNR_Cell)
		if err = ie.sftd_MeasNR_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode sftd_MeasNR_Cell", err)
		}
	}
	return nil
}
