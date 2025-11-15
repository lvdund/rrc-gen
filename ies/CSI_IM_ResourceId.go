package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_ResourceId struct {
	Value uint64 `lb:0,ub:maxNrofCSI_IM_Resources_1,madatory`
}

func (ie *CSI_IM_ResourceId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofCSI_IM_Resources_1}, false); err != nil {
		return utils.WrapError("Encode CSI_IM_ResourceId", err)
	}
	return nil
}

func (ie *CSI_IM_ResourceId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCSI_IM_Resources_1}, false); err != nil {
		return utils.WrapError("Decode CSI_IM_ResourceId", err)
	}
	ie.Value = uint64(v)
	return nil
}
