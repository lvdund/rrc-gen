package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_precoderGranularity_Enum_sameAsREG_bundle uper.Enumerated = 0
	ControlResourceSet_precoderGranularity_Enum_allContiguousRBs uper.Enumerated = 1
)

type ControlResourceSet_precoderGranularity struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ControlResourceSet_precoderGranularity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSet_precoderGranularity", err)
	}
	return nil
}

func (ie *ControlResourceSet_precoderGranularity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSet_precoderGranularity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
