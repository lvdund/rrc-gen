package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_Resource_Mobility_associatedSSB struct {
	ssb_Index        SSB_Index `madatory`
	isQuasiColocated bool      `madatory`
}

func (ie *CSI_RS_Resource_Mobility_associatedSSB) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssb_Index.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_Index", err)
	}
	if err = w.WriteBoolean(ie.isQuasiColocated); err != nil {
		return utils.WrapError("WriteBoolean isQuasiColocated", err)
	}
	return nil
}

func (ie *CSI_RS_Resource_Mobility_associatedSSB) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssb_Index.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_Index", err)
	}
	var tmp_bool_isQuasiColocated bool
	if tmp_bool_isQuasiColocated, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean isQuasiColocated", err)
	}
	ie.isQuasiColocated = tmp_bool_isQuasiColocated
	return nil
}
