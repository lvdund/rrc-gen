package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_Starting_SFN_AndSubframe_r17 struct {
	starting_SFN_r17     int64 `lb:0,ub:1023,madatory`
	startingSubframe_r17 int64 `lb:0,ub:9,madatory`
}

func (ie *MUSIM_Starting_SFN_AndSubframe_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.starting_SFN_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger starting_SFN_r17", err)
	}
	if err = w.WriteInteger(ie.startingSubframe_r17, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("WriteInteger startingSubframe_r17", err)
	}
	return nil
}

func (ie *MUSIM_Starting_SFN_AndSubframe_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_starting_SFN_r17 int64
	if tmp_int_starting_SFN_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger starting_SFN_r17", err)
	}
	ie.starting_SFN_r17 = tmp_int_starting_SFN_r17
	var tmp_int_startingSubframe_r17 int64
	if tmp_int_startingSubframe_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("ReadInteger startingSubframe_r17", err)
	}
	ie.startingSubframe_r17 = tmp_int_startingSubframe_r17
	return nil
}
