package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerCSI_RS_IndexList struct {
	Value []ResultsPerCSI_RS_Index `lb:1,ub:maxNrofIndexesToReport2,madatory`
}

func (ie *ResultsPerCSI_RS_IndexList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ResultsPerCSI_RS_Index]([]*ResultsPerCSI_RS_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ResultsPerCSI_RS_IndexList", err)
	}
	return nil
}

func (ie *ResultsPerCSI_RS_IndexList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ResultsPerCSI_RS_Index]([]*ResultsPerCSI_RS_Index{}, uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport2}, false)
	fn := func() *ResultsPerCSI_RS_Index {
		return new(ResultsPerCSI_RS_Index)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ResultsPerCSI_RS_IndexList", err)
	}
	ie.Value = []ResultsPerCSI_RS_Index{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
