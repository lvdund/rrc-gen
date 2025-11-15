package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_freqHopping struct {
	c_SRS int64 `lb:0,ub:63,madatory`
	b_SRS int64 `lb:0,ub:3,madatory`
	b_hop int64 `lb:0,ub:3,madatory`
}

func (ie *SRS_Resource_freqHopping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.c_SRS, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger c_SRS", err)
	}
	if err = w.WriteInteger(ie.b_SRS, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger b_SRS", err)
	}
	if err = w.WriteInteger(ie.b_hop, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger b_hop", err)
	}
	return nil
}

func (ie *SRS_Resource_freqHopping) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_c_SRS int64
	if tmp_int_c_SRS, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger c_SRS", err)
	}
	ie.c_SRS = tmp_int_c_SRS
	var tmp_int_b_SRS int64
	if tmp_int_b_SRS, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger b_SRS", err)
	}
	ie.b_SRS = tmp_int_b_SRS
	var tmp_int_b_hop int64
	if tmp_int_b_hop, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger b_hop", err)
	}
	ie.b_hop = tmp_int_b_hop
	return nil
}
