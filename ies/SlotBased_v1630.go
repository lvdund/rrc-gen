package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotBased_v1630 struct {
	tciMapping_r16          SlotBased_v1630_tciMapping_r16 `madatory`
	sequenceOffsetForRV_r16 int64                          `lb:0,ub:0,madatory`
}

func (ie *SlotBased_v1630) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.tciMapping_r16.Encode(w); err != nil {
		return utils.WrapError("Encode tciMapping_r16", err)
	}
	if err = w.WriteInteger(ie.sequenceOffsetForRV_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteInteger sequenceOffsetForRV_r16", err)
	}
	return nil
}

func (ie *SlotBased_v1630) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.tciMapping_r16.Decode(r); err != nil {
		return utils.WrapError("Decode tciMapping_r16", err)
	}
	var tmp_int_sequenceOffsetForRV_r16 int64
	if tmp_int_sequenceOffsetForRV_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadInteger sequenceOffsetForRV_r16", err)
	}
	ie.sequenceOffsetForRV_r16 = tmp_int_sequenceOffsetForRV_r16
	return nil
}
