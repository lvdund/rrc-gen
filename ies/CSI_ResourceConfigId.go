package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ResourceConfigId struct {
	Value uint64 `lb:0,ub:maxNrofCSI_ResourceConfigurations_1,madatory`
}

func (ie *CSI_ResourceConfigId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofCSI_ResourceConfigurations_1}, false); err != nil {
		return utils.WrapError("Encode CSI_ResourceConfigId", err)
	}
	return nil
}

func (ie *CSI_ResourceConfigId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCSI_ResourceConfigurations_1}, false); err != nil {
		return utils.WrapError("Decode CSI_ResourceConfigId", err)
	}
	ie.Value = uint64(v)
	return nil
}
