package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServMO struct {
	servCellId              ServCellIndex `madatory`
	measResultServingCell   MeasResultNR  `madatory`
	measResultBestNeighCell *MeasResultNR `optional`
}

func (ie *MeasResultServMO) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultBestNeighCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servCellId.Encode(w); err != nil {
		return utils.WrapError("Encode servCellId", err)
	}
	if err = ie.measResultServingCell.Encode(w); err != nil {
		return utils.WrapError("Encode measResultServingCell", err)
	}
	if ie.measResultBestNeighCell != nil {
		if err = ie.measResultBestNeighCell.Encode(w); err != nil {
			return utils.WrapError("Encode measResultBestNeighCell", err)
		}
	}
	return nil
}

func (ie *MeasResultServMO) Decode(r *uper.UperReader) error {
	var err error
	var measResultBestNeighCellPresent bool
	if measResultBestNeighCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servCellId.Decode(r); err != nil {
		return utils.WrapError("Decode servCellId", err)
	}
	if err = ie.measResultServingCell.Decode(r); err != nil {
		return utils.WrapError("Decode measResultServingCell", err)
	}
	if measResultBestNeighCellPresent {
		ie.measResultBestNeighCell = new(MeasResultNR)
		if err = ie.measResultBestNeighCell.Decode(r); err != nil {
			return utils.WrapError("Decode measResultBestNeighCell", err)
		}
	}
	return nil
}
