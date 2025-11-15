package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SFTD_FrequencyList_EUTRA struct {
	Value []ARFCN_ValueEUTRA `lb:1,ub:maxCellSFTD,madatory`
}

func (ie *SFTD_FrequencyList_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SFTD_FrequencyList_EUTRA", err)
	}
	return nil
}

func (ie *SFTD_FrequencyList_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	fn := func() *ARFCN_ValueEUTRA {
		return new(ARFCN_ValueEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SFTD_FrequencyList_EUTRA", err)
	}
	ie.Value = []ARFCN_ValueEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
