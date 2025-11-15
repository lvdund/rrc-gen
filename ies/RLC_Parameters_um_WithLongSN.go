package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLC_Parameters_um_WithLongSN_Enum_supported uper.Enumerated = 0
)

type RLC_Parameters_um_WithLongSN struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RLC_Parameters_um_WithLongSN) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RLC_Parameters_um_WithLongSN", err)
	}
	return nil
}

func (ie *RLC_Parameters_um_WithLongSN) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RLC_Parameters_um_WithLongSN", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
