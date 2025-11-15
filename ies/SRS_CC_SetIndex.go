package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_CC_SetIndex struct {
	cc_SetIndex         *int64 `lb:0,ub:3,optional`
	cc_IndexInOneCC_Set *int64 `lb:0,ub:7,optional`
}

func (ie *SRS_CC_SetIndex) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cc_SetIndex != nil, ie.cc_IndexInOneCC_Set != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cc_SetIndex != nil {
		if err = w.WriteInteger(*ie.cc_SetIndex, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode cc_SetIndex", err)
		}
	}
	if ie.cc_IndexInOneCC_Set != nil {
		if err = w.WriteInteger(*ie.cc_IndexInOneCC_Set, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Encode cc_IndexInOneCC_Set", err)
		}
	}
	return nil
}

func (ie *SRS_CC_SetIndex) Decode(r *uper.UperReader) error {
	var err error
	var cc_SetIndexPresent bool
	if cc_SetIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cc_IndexInOneCC_SetPresent bool
	if cc_IndexInOneCC_SetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cc_SetIndexPresent {
		var tmp_int_cc_SetIndex int64
		if tmp_int_cc_SetIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode cc_SetIndex", err)
		}
		ie.cc_SetIndex = &tmp_int_cc_SetIndex
	}
	if cc_IndexInOneCC_SetPresent {
		var tmp_int_cc_IndexInOneCC_Set int64
		if tmp_int_cc_IndexInOneCC_Set, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode cc_IndexInOneCC_Set", err)
		}
		ie.cc_IndexInOneCC_Set = &tmp_int_cc_IndexInOneCC_Set
	}
	return nil
}
