package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NZP_CSI_RS_ResourceId struct {
	Value uint64 `lb:0,ub:maxNrofNZP_CSI_RS_Resources_1,madatory`
}

func (ie *NZP_CSI_RS_ResourceId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofNZP_CSI_RS_Resources_1}, false); err != nil {
		return utils.WrapError("Encode NZP_CSI_RS_ResourceId", err)
	}
	return nil
}

func (ie *NZP_CSI_RS_ResourceId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofNZP_CSI_RS_Resources_1}, false); err != nil {
		return utils.WrapError("Decode NZP_CSI_RS_ResourceId", err)
	}
	ie.Value = uint64(v)
	return nil
}
