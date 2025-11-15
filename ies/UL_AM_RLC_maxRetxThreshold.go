package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_AM_RLC_maxRetxThreshold_Enum_t1  uper.Enumerated = 0
	UL_AM_RLC_maxRetxThreshold_Enum_t2  uper.Enumerated = 1
	UL_AM_RLC_maxRetxThreshold_Enum_t3  uper.Enumerated = 2
	UL_AM_RLC_maxRetxThreshold_Enum_t4  uper.Enumerated = 3
	UL_AM_RLC_maxRetxThreshold_Enum_t6  uper.Enumerated = 4
	UL_AM_RLC_maxRetxThreshold_Enum_t8  uper.Enumerated = 5
	UL_AM_RLC_maxRetxThreshold_Enum_t16 uper.Enumerated = 6
	UL_AM_RLC_maxRetxThreshold_Enum_t32 uper.Enumerated = 7
)

type UL_AM_RLC_maxRetxThreshold struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UL_AM_RLC_maxRetxThreshold) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UL_AM_RLC_maxRetxThreshold", err)
	}
	return nil
}

func (ie *UL_AM_RLC_maxRetxThreshold) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UL_AM_RLC_maxRetxThreshold", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
