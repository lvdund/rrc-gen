package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_Parameters_r17 struct {
	maxMRB_Add_r17 *int64 `lb:1,ub:16,optional`
}

func (ie *MBS_Parameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxMRB_Add_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxMRB_Add_r17 != nil {
		if err = w.WriteInteger(*ie.maxMRB_Add_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode maxMRB_Add_r17", err)
		}
	}
	return nil
}

func (ie *MBS_Parameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var maxMRB_Add_r17Present bool
	if maxMRB_Add_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxMRB_Add_r17Present {
		var tmp_int_maxMRB_Add_r17 int64
		if tmp_int_maxMRB_Add_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode maxMRB_Add_r17", err)
		}
		ie.maxMRB_Add_r17 = &tmp_int_maxMRB_Add_r17
	}
	return nil
}
