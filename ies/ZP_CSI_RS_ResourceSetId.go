package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ZP_CSI_RS_ResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofZP_CSI_RS_ResourceSets_1,madatory`
}

func (ie *ZP_CSI_RS_ResourceSetId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofZP_CSI_RS_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode ZP_CSI_RS_ResourceSetId", err)
	}
	return nil
}

func (ie *ZP_CSI_RS_ResourceSetId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofZP_CSI_RS_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode ZP_CSI_RS_ResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
