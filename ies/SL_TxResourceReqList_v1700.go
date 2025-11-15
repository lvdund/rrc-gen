package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqList_v1700 struct {
	Value []SL_TxResourceReq_v1700 `lb:1,ub:maxNrofSL_Dest_r16,madatory`
}

func (ie *SL_TxResourceReqList_v1700) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_TxResourceReq_v1700]([]*SL_TxResourceReq_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_TxResourceReqList_v1700", err)
	}
	return nil
}

func (ie *SL_TxResourceReqList_v1700) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_TxResourceReq_v1700]([]*SL_TxResourceReq_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
	fn := func() *SL_TxResourceReq_v1700 {
		return new(SL_TxResourceReq_v1700)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_TxResourceReqList_v1700", err)
	}
	ie.Value = []SL_TxResourceReq_v1700{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
