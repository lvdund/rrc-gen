package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB_Mapping_v1700 struct {
	Value []SIB_TypeInfo_v1700 `lb:1,ub:maxSIB,madatory`
}

func (ie *SIB_Mapping_v1700) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SIB_TypeInfo_v1700]([]*SIB_TypeInfo_v1700{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SIB_Mapping_v1700", err)
	}
	return nil
}

func (ie *SIB_Mapping_v1700) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SIB_TypeInfo_v1700]([]*SIB_TypeInfo_v1700{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	fn := func() *SIB_TypeInfo_v1700 {
		return new(SIB_TypeInfo_v1700)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SIB_Mapping_v1700", err)
	}
	ie.Value = []SIB_TypeInfo_v1700{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
