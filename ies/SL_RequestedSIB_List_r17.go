package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RequestedSIB_List_r17 struct {
	Value []SL_SIB_ReqInfo_r17 `lb:maxSIB_MessagePlus1_r17,ub:maxSIB_MessagePlus1_r17,madatory`
}

func (ie *SL_RequestedSIB_List_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SL_SIB_ReqInfo_r17]([]*SL_SIB_ReqInfo_r17{}, uper.Constraint{Lb: maxSIB_MessagePlus1_r17, Ub: maxSIB_MessagePlus1_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_RequestedSIB_List_r17", err)
	}
	return nil
}

func (ie *SL_RequestedSIB_List_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SL_SIB_ReqInfo_r17]([]*SL_SIB_ReqInfo_r17{}, uper.Constraint{Lb: maxSIB_MessagePlus1_r17, Ub: maxSIB_MessagePlus1_r17}, false)
	fn := func() *SL_SIB_ReqInfo_r17 {
		return new(SL_SIB_ReqInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_RequestedSIB_List_r17", err)
	}
	ie.Value = []SL_SIB_ReqInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
