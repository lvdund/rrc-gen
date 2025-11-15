package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl1  uper.Enumerated = 0
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl2  uper.Enumerated = 1
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl4  uper.Enumerated = 2
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl8  uper.Enumerated = 3
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl10 uper.Enumerated = 4
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl20 uper.Enumerated = 5
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl40 uper.Enumerated = 6
	RACH_ConfigGeneric_ra_ResponseWindow_Enum_sl80 uper.Enumerated = 7
)

type RACH_ConfigGeneric_ra_ResponseWindow struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_ra_ResponseWindow", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_ra_ResponseWindow", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
