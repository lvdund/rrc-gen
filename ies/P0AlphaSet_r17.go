package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0AlphaSet_r17 struct {
	p0_r17              *int64                             `lb:-16,ub:15,optional`
	alpha_r17           *Alpha                             `optional`
	closedLoopIndex_r17 P0AlphaSet_r17_closedLoopIndex_r17 `madatory`
}

func (ie *P0AlphaSet_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.p0_r17 != nil, ie.alpha_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.p0_r17 != nil {
		if err = w.WriteInteger(*ie.p0_r17, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode p0_r17", err)
		}
	}
	if ie.alpha_r17 != nil {
		if err = ie.alpha_r17.Encode(w); err != nil {
			return utils.WrapError("Encode alpha_r17", err)
		}
	}
	if err = ie.closedLoopIndex_r17.Encode(w); err != nil {
		return utils.WrapError("Encode closedLoopIndex_r17", err)
	}
	return nil
}

func (ie *P0AlphaSet_r17) Decode(r *uper.UperReader) error {
	var err error
	var p0_r17Present bool
	if p0_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var alpha_r17Present bool
	if alpha_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if p0_r17Present {
		var tmp_int_p0_r17 int64
		if tmp_int_p0_r17, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode p0_r17", err)
		}
		ie.p0_r17 = &tmp_int_p0_r17
	}
	if alpha_r17Present {
		ie.alpha_r17 = new(Alpha)
		if err = ie.alpha_r17.Decode(r); err != nil {
			return utils.WrapError("Decode alpha_r17", err)
		}
	}
	if err = ie.closedLoopIndex_r17.Decode(r); err != nil {
		return utils.WrapError("Decode closedLoopIndex_r17", err)
	}
	return nil
}
