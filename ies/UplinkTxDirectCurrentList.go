package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentList struct {
	Value []UplinkTxDirectCurrentCell `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *UplinkTxDirectCurrentList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UplinkTxDirectCurrentCell]([]*UplinkTxDirectCurrentCell{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UplinkTxDirectCurrentList", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UplinkTxDirectCurrentCell]([]*UplinkTxDirectCurrentCell{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *UplinkTxDirectCurrentCell {
		return new(UplinkTxDirectCurrentCell)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UplinkTxDirectCurrentList", err)
	}
	ie.Value = []UplinkTxDirectCurrentCell{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
