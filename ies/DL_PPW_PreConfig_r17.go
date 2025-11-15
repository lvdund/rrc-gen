package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_PPW_PreConfig_r17 struct {
	dl_PPW_ID_r17                      DL_PPW_ID_r17                      `madatory`
	dl_PPW_PeriodicityAndStartSlot_r17 DL_PPW_PeriodicityAndStartSlot_r17 `madatory`
	length_r17                         int64                              `lb:1,ub:160,madatory`
	type_r17                           *DL_PPW_PreConfig_r17_type_r17     `optional`
	priority_r17                       *DL_PPW_PreConfig_r17_priority_r17 `optional`
}

func (ie *DL_PPW_PreConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.type_r17 != nil, ie.priority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dl_PPW_ID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode dl_PPW_ID_r17", err)
	}
	if err = ie.dl_PPW_PeriodicityAndStartSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode dl_PPW_PeriodicityAndStartSlot_r17", err)
	}
	if err = w.WriteInteger(ie.length_r17, &uper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return utils.WrapError("WriteInteger length_r17", err)
	}
	if ie.type_r17 != nil {
		if err = ie.type_r17.Encode(w); err != nil {
			return utils.WrapError("Encode type_r17", err)
		}
	}
	if ie.priority_r17 != nil {
		if err = ie.priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode priority_r17", err)
		}
	}
	return nil
}

func (ie *DL_PPW_PreConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var type_r17Present bool
	if type_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var priority_r17Present bool
	if priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dl_PPW_ID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode dl_PPW_ID_r17", err)
	}
	if err = ie.dl_PPW_PeriodicityAndStartSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode dl_PPW_PeriodicityAndStartSlot_r17", err)
	}
	var tmp_int_length_r17 int64
	if tmp_int_length_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return utils.WrapError("ReadInteger length_r17", err)
	}
	ie.length_r17 = tmp_int_length_r17
	if type_r17Present {
		ie.type_r17 = new(DL_PPW_PreConfig_r17_type_r17)
		if err = ie.type_r17.Decode(r); err != nil {
			return utils.WrapError("Decode type_r17", err)
		}
	}
	if priority_r17Present {
		ie.priority_r17 = new(DL_PPW_PreConfig_r17_priority_r17)
		if err = ie.priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode priority_r17", err)
		}
	}
	return nil
}
