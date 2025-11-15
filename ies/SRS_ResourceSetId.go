package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofSRS_ResourceSets_1,madatory`
}

func (ie *SRS_ResourceSetId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSRS_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode SRS_ResourceSetId", err)
	}
	return nil
}

func (ie *SRS_ResourceSetId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSRS_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode SRS_ResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
