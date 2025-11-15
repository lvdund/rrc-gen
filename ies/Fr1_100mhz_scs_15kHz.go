package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Fr1_100mhz_scs_15kHz_Enum_supported uper.Enumerated = 0
)

type Fr1_100mhz_scs_15kHz struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Fr1_100mhz_scs_15kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Fr1_100mhz_scs_15kHz", err)
	}
	return nil
}

func (ie *Fr1_100mhz_scs_15kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Fr1_100mhz_scs_15kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
