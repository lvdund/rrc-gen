package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2NR struct {
	ssbFrequency              *ARFCN_ValueNR    `optional`
	refFreqCSI_RS             *ARFCN_ValueNR    `optional`
	measResultServingCell     *MeasResultNR     `optional`
	measResultNeighCellListNR *MeasResultListNR `optional`
}

func (ie *MeasResult2NR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssbFrequency != nil, ie.refFreqCSI_RS != nil, ie.measResultServingCell != nil, ie.measResultNeighCellListNR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssbFrequency != nil {
		if err = ie.ssbFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode ssbFrequency", err)
		}
	}
	if ie.refFreqCSI_RS != nil {
		if err = ie.refFreqCSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode refFreqCSI_RS", err)
		}
	}
	if ie.measResultServingCell != nil {
		if err = ie.measResultServingCell.Encode(w); err != nil {
			return utils.WrapError("Encode measResultServingCell", err)
		}
	}
	if ie.measResultNeighCellListNR != nil {
		if err = ie.measResultNeighCellListNR.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCellListNR", err)
		}
	}
	return nil
}

func (ie *MeasResult2NR) Decode(r *uper.UperReader) error {
	var err error
	var ssbFrequencyPresent bool
	if ssbFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var refFreqCSI_RSPresent bool
	if refFreqCSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultServingCellPresent bool
	if measResultServingCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultNeighCellListNRPresent bool
	if measResultNeighCellListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ssbFrequencyPresent {
		ie.ssbFrequency = new(ARFCN_ValueNR)
		if err = ie.ssbFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode ssbFrequency", err)
		}
	}
	if refFreqCSI_RSPresent {
		ie.refFreqCSI_RS = new(ARFCN_ValueNR)
		if err = ie.refFreqCSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode refFreqCSI_RS", err)
		}
	}
	if measResultServingCellPresent {
		ie.measResultServingCell = new(MeasResultNR)
		if err = ie.measResultServingCell.Decode(r); err != nil {
			return utils.WrapError("Decode measResultServingCell", err)
		}
	}
	if measResultNeighCellListNRPresent {
		ie.measResultNeighCellListNR = new(MeasResultListNR)
		if err = ie.measResultNeighCellListNR.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCellListNR", err)
		}
	}
	return nil
}
