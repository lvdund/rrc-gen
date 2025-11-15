package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_FreqNeighCellList struct {
	Value []EUTRA_FreqNeighCellInfo `lb:1,ub:maxCellEUTRA,madatory`
}

func (ie *EUTRA_FreqNeighCellList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_FreqNeighCellInfo]([]*EUTRA_FreqNeighCellInfo{}, uper.Constraint{Lb: 1, Ub: maxCellEUTRA}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_FreqNeighCellList", err)
	}
	return nil
}

func (ie *EUTRA_FreqNeighCellList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_FreqNeighCellInfo]([]*EUTRA_FreqNeighCellInfo{}, uper.Constraint{Lb: 1, Ub: maxCellEUTRA}, false)
	fn := func() *EUTRA_FreqNeighCellInfo {
		return new(EUTRA_FreqNeighCellInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_FreqNeighCellList", err)
	}
	ie.Value = []EUTRA_FreqNeighCellInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
