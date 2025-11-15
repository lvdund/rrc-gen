package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl240  uper.Enumerated = 0
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl320  uper.Enumerated = 1
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl640  uper.Enumerated = 2
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl960  uper.Enumerated = 3
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl1280 uper.Enumerated = 4
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl1920 uper.Enumerated = 5
	RACH_ConfigGeneric_ra_ResponseWindow_v1700_Enum_sl2560 uper.Enumerated = 6
)

type RACH_ConfigGeneric_ra_ResponseWindow_v1700 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_ra_ResponseWindow_v1700", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_ra_ResponseWindow_v1700) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_ra_ResponseWindow_v1700", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
