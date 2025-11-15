package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_AM_RLC_v1610 struct {
	t_StatusProhibit_v1610 *T_StatusProhibit_v1610 `optional`
}

func (ie *DL_AM_RLC_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.t_StatusProhibit_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.t_StatusProhibit_v1610 != nil {
		if err = ie.t_StatusProhibit_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode t_StatusProhibit_v1610", err)
		}
	}
	return nil
}

func (ie *DL_AM_RLC_v1610) Decode(r *uper.UperReader) error {
	var err error
	var t_StatusProhibit_v1610Present bool
	if t_StatusProhibit_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if t_StatusProhibit_v1610Present {
		ie.t_StatusProhibit_v1610 = new(T_StatusProhibit_v1610)
		if err = ie.t_StatusProhibit_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode t_StatusProhibit_v1610", err)
		}
	}
	return nil
}
