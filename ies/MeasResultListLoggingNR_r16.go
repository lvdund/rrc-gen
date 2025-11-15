package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultListLoggingNR_r16 struct {
	Value []MeasResultLoggingNR_r16 `lb:1,ub:maxCellReport,madatory`
}

func (ie *MeasResultListLoggingNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultLoggingNR_r16]([]*MeasResultLoggingNR_r16{}, uper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListLoggingNR_r16", err)
	}
	return nil
}

func (ie *MeasResultListLoggingNR_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultLoggingNR_r16]([]*MeasResultLoggingNR_r16{}, uper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	fn := func() *MeasResultLoggingNR_r16 {
		return new(MeasResultLoggingNR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultListLoggingNR_r16", err)
	}
	ie.Value = []MeasResultLoggingNR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
