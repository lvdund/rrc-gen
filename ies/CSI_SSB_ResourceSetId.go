package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_SSB_ResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofCSI_SSB_ResourceSets_1,madatory`
}

func (ie *CSI_SSB_ResourceSetId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofCSI_SSB_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode CSI_SSB_ResourceSetId", err)
	}
	return nil
}

func (ie *CSI_SSB_ResourceSetId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCSI_SSB_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode CSI_SSB_ResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
