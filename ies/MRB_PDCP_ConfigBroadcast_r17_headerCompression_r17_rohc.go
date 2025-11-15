package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc struct {
	maxCID_r17   int64                                                                `lb:1,ub:16,madatory`
	profiles_r17 MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc_profiles_r17 `madatory`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxCID_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger maxCID_r17", err)
	}
	if err = ie.profiles_r17.Encode(w); err != nil {
		return utils.WrapError("Encode profiles_r17", err)
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxCID_r17 int64
	if tmp_int_maxCID_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger maxCID_r17", err)
	}
	ie.maxCID_r17 = tmp_int_maxCID_r17
	if err = ie.profiles_r17.Decode(r); err != nil {
		return utils.WrapError("Decode profiles_r17", err)
	}
	return nil
}
