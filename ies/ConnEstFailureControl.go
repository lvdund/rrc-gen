package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailureControl struct {
	connEstFailCount          ConnEstFailureControl_connEstFailCount          `madatory`
	connEstFailOffsetValidity ConnEstFailureControl_connEstFailOffsetValidity `madatory`
	connEstFailOffset         *int64                                          `lb:0,ub:15,optional`
}

func (ie *ConnEstFailureControl) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.connEstFailOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.connEstFailCount.Encode(w); err != nil {
		return utils.WrapError("Encode connEstFailCount", err)
	}
	if err = ie.connEstFailOffsetValidity.Encode(w); err != nil {
		return utils.WrapError("Encode connEstFailOffsetValidity", err)
	}
	if ie.connEstFailOffset != nil {
		if err = w.WriteInteger(*ie.connEstFailOffset, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode connEstFailOffset", err)
		}
	}
	return nil
}

func (ie *ConnEstFailureControl) Decode(r *uper.UperReader) error {
	var err error
	var connEstFailOffsetPresent bool
	if connEstFailOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.connEstFailCount.Decode(r); err != nil {
		return utils.WrapError("Decode connEstFailCount", err)
	}
	if err = ie.connEstFailOffsetValidity.Decode(r); err != nil {
		return utils.WrapError("Decode connEstFailOffsetValidity", err)
	}
	if connEstFailOffsetPresent {
		var tmp_int_connEstFailOffset int64
		if tmp_int_connEstFailOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode connEstFailOffset", err)
		}
		ie.connEstFailOffset = &tmp_int_connEstFailOffset
	}
	return nil
}
