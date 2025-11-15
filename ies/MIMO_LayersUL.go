package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_LayersUL_Enum_oneLayer   uper.Enumerated = 0
	MIMO_LayersUL_Enum_twoLayers  uper.Enumerated = 1
	MIMO_LayersUL_Enum_fourLayers uper.Enumerated = 2
)

type MIMO_LayersUL struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MIMO_LayersUL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MIMO_LayersUL", err)
	}
	return nil
}

func (ie *MIMO_LayersUL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MIMO_LayersUL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
