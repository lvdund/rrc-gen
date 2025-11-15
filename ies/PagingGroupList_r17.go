package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PagingGroupList_r17 struct {
	Value []TMGI_r17 `lb:1,ub:maxNrofPageGroup_r17,madatory`
}

func (ie *PagingGroupList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*TMGI_r17]([]*TMGI_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPageGroup_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PagingGroupList_r17", err)
	}
	return nil
}

func (ie *PagingGroupList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*TMGI_r17]([]*TMGI_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPageGroup_r17}, false)
	fn := func() *TMGI_r17 {
		return new(TMGI_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PagingGroupList_r17", err)
	}
	ie.Value = []TMGI_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
