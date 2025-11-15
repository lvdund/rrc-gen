package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2EUTRA struct {
	carrierFreq             ARFCN_ValueEUTRA `madatory`
	measResultServingCell   *MeasResultEUTRA `optional`
	measResultBestNeighCell *MeasResultEUTRA `optional`
}

func (ie *MeasResult2EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultServingCell != nil, ie.measResultBestNeighCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if ie.measResultServingCell != nil {
		if err = ie.measResultServingCell.Encode(w); err != nil {
			return utils.WrapError("Encode measResultServingCell", err)
		}
	}
	if ie.measResultBestNeighCell != nil {
		if err = ie.measResultBestNeighCell.Encode(w); err != nil {
			return utils.WrapError("Encode measResultBestNeighCell", err)
		}
	}
	return nil
}

func (ie *MeasResult2EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var measResultServingCellPresent bool
	if measResultServingCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultBestNeighCellPresent bool
	if measResultBestNeighCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	if measResultServingCellPresent {
		ie.measResultServingCell = new(MeasResultEUTRA)
		if err = ie.measResultServingCell.Decode(r); err != nil {
			return utils.WrapError("Decode measResultServingCell", err)
		}
	}
	if measResultBestNeighCellPresent {
		ie.measResultBestNeighCell = new(MeasResultEUTRA)
		if err = ie.measResultBestNeighCell.Decode(r); err != nil {
			return utils.WrapError("Decode measResultBestNeighCell", err)
		}
	}
	return nil
}
