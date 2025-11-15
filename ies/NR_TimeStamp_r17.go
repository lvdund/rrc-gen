package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_TimeStamp_r17 struct {
	nr_SFN_r17  int64                        `lb:0,ub:1023,madatory`
	nr_Slot_r17 NR_TimeStamp_r17_nr_Slot_r17 `lb:0,ub:9,madatory`
}

func (ie *NR_TimeStamp_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.nr_SFN_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger nr_SFN_r17", err)
	}
	if err = ie.nr_Slot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nr_Slot_r17", err)
	}
	return nil
}

func (ie *NR_TimeStamp_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_nr_SFN_r17 int64
	if tmp_int_nr_SFN_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger nr_SFN_r17", err)
	}
	ie.nr_SFN_r17 = tmp_int_nr_SFN_r17
	if err = ie.nr_Slot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nr_Slot_r17", err)
	}
	return nil
}
