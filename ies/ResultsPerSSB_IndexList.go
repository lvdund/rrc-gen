package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_IndexList struct {
	Value []ResultsPerSSB_Index `lb:1,ub:maxNrofIndexesToReport2,madatory`
}

func (ie *ResultsPerSSB_IndexList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ResultsPerSSB_Index]([]*ResultsPerSSB_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ResultsPerSSB_IndexList", err)
	}
	return nil
}

func (ie *ResultsPerSSB_IndexList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ResultsPerSSB_Index]([]*ResultsPerSSB_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport2}, false)
	fn := func() *ResultsPerSSB_Index {
		return new(ResultsPerSSB_Index)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ResultsPerSSB_IndexList", err)
	}
	ie.Value = []ResultsPerSSB_Index{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
