package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB_Mapping struct {
	Value []SIB_TypeInfo `lb:1,ub:maxSIB,madatory`
}

func (ie *SIB_Mapping) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SIB_TypeInfo]([]*SIB_TypeInfo{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SIB_Mapping", err)
	}
	return nil
}

func (ie *SIB_Mapping) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SIB_TypeInfo]([]*SIB_TypeInfo{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	fn := func() *SIB_TypeInfo {
		return new(SIB_TypeInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SIB_Mapping", err)
	}
	ie.Value = []SIB_TypeInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
