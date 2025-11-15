package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_pathlossReferenceLinking_Enum_spCell uper.Enumerated = 0
	ServingCellConfig_pathlossReferenceLinking_Enum_sCell  uper.Enumerated = 1
)

type ServingCellConfig_pathlossReferenceLinking struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ServingCellConfig_pathlossReferenceLinking) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfig_pathlossReferenceLinking", err)
	}
	return nil
}

func (ie *ServingCellConfig_pathlossReferenceLinking) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfig_pathlossReferenceLinking", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
