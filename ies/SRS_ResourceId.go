package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceId struct {
	Value uint64 `lb:0,ub:maxNrofSRS_Resources_1,madatory`
}

func (ie *SRS_ResourceId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSRS_Resources_1}, false); err != nil {
		return utils.WrapError("Encode SRS_ResourceId", err)
	}
	return nil
}

func (ie *SRS_ResourceId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSRS_Resources_1}, false); err != nil {
		return utils.WrapError("Decode SRS_ResourceId", err)
	}
	ie.Value = uint64(v)
	return nil
}
