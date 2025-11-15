package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PagingRecordList_v1700 struct {
	Value []PagingRecord_v1700 `lb:1,ub:maxNrofPageRec,madatory`
}

func (ie *PagingRecordList_v1700) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PagingRecord_v1700]([]*PagingRecord_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofPageRec}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PagingRecordList_v1700", err)
	}
	return nil
}

func (ie *PagingRecordList_v1700) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PagingRecord_v1700]([]*PagingRecord_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofPageRec}, false)
	fn := func() *PagingRecord_v1700 {
		return new(PagingRecord_v1700)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PagingRecordList_v1700", err)
	}
	ie.Value = []PagingRecord_v1700{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
