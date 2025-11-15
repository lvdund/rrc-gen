package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_UM_RLC_v1700 struct {
	t_ReassemblyExt_r17 *T_ReassemblyExt_r17 `optional`
}

func (ie *DL_UM_RLC_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.t_ReassemblyExt_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.t_ReassemblyExt_r17 != nil {
		if err = ie.t_ReassemblyExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t_ReassemblyExt_r17", err)
		}
	}
	return nil
}

func (ie *DL_UM_RLC_v1700) Decode(r *uper.UperReader) error {
	var err error
	var t_ReassemblyExt_r17Present bool
	if t_ReassemblyExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if t_ReassemblyExt_r17Present {
		ie.t_ReassemblyExt_r17 = new(T_ReassemblyExt_r17)
		if err = ie.t_ReassemblyExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t_ReassemblyExt_r17", err)
		}
	}
	return nil
}
