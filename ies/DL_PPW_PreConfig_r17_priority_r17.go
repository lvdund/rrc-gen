package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PreConfig_r17_priority_r17_Enum_st1 uper.Enumerated = 0
	DL_PPW_PreConfig_r17_priority_r17_Enum_st2 uper.Enumerated = 1
	DL_PPW_PreConfig_r17_priority_r17_Enum_st3 uper.Enumerated = 2
)

type DL_PPW_PreConfig_r17_priority_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DL_PPW_PreConfig_r17_priority_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DL_PPW_PreConfig_r17_priority_r17", err)
	}
	return nil
}

func (ie *DL_PPW_PreConfig_r17_priority_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DL_PPW_PreConfig_r17_priority_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
