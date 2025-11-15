package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PagingRecordList struct {
	Value []PagingRecord `lb:1,ub:maxNrofPageRec,madatory`
}

func (ie *PagingRecordList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PagingRecord]([]*PagingRecord{}, uper.Constraint{Lb: 1, Ub: maxNrofPageRec}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PagingRecordList", err)
	}
	return nil
}

func (ie *PagingRecordList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PagingRecord]([]*PagingRecord{}, uper.Constraint{Lb: 1, Ub: maxNrofPageRec}, false)
	fn := func() *PagingRecord {
		return new(PagingRecord)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PagingRecordList", err)
	}
	ie.Value = []PagingRecord{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
