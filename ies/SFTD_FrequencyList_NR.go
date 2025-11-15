package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SFTD_FrequencyList_NR struct {
	Value []ARFCN_ValueNR `lb:1,ub:maxCellSFTD,madatory`
}

func (ie *SFTD_FrequencyList_NR) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SFTD_FrequencyList_NR", err)
	}
	return nil
}

func (ie *SFTD_FrequencyList_NR) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	fn := func() *ARFCN_ValueNR {
		return new(ARFCN_ValueNR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SFTD_FrequencyList_NR", err)
	}
	ie.Value = []ARFCN_ValueNR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
