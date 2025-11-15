package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_FRX_Diff struct {
	simultaneousRxDataSSB_DiffNumerology *MeasAndMobParametersMRDC_FRX_Diff_simultaneousRxDataSSB_DiffNumerology `optional`
}

func (ie *MeasAndMobParametersMRDC_FRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.simultaneousRxDataSSB_DiffNumerology != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.simultaneousRxDataSSB_DiffNumerology != nil {
		if err = ie.simultaneousRxDataSSB_DiffNumerology.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxDataSSB_DiffNumerology", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_FRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var simultaneousRxDataSSB_DiffNumerologyPresent bool
	if simultaneousRxDataSSB_DiffNumerologyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if simultaneousRxDataSSB_DiffNumerologyPresent {
		ie.simultaneousRxDataSSB_DiffNumerology = new(MeasAndMobParametersMRDC_FRX_Diff_simultaneousRxDataSSB_DiffNumerology)
		if err = ie.simultaneousRxDataSSB_DiffNumerology.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxDataSSB_DiffNumerology", err)
		}
	}
	return nil
}
