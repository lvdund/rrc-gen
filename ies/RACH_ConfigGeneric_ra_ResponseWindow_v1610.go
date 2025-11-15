package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_ra_ResponseWindow_v1610_Enum_sl60  uper.Enumerated = 0
	RACH_ConfigGeneric_ra_ResponseWindow_v1610_Enum_sl160 uper.Enumerated = 1
)

type RACH_ConfigGeneric_ra_ResponseWindow_v1610 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_ra_ResponseWindow_v1610", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_ra_ResponseWindow_v1610", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
