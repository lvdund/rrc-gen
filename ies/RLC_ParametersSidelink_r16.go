package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_ParametersSidelink_r16 struct {
	am_WithLongSN_Sidelink_r16 *RLC_ParametersSidelink_r16_am_WithLongSN_Sidelink_r16 `optional`
	um_WithLongSN_Sidelink_r16 *RLC_ParametersSidelink_r16_um_WithLongSN_Sidelink_r16 `optional`
}

func (ie *RLC_ParametersSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.am_WithLongSN_Sidelink_r16 != nil, ie.um_WithLongSN_Sidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.am_WithLongSN_Sidelink_r16 != nil {
		if err = ie.am_WithLongSN_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode am_WithLongSN_Sidelink_r16", err)
		}
	}
	if ie.um_WithLongSN_Sidelink_r16 != nil {
		if err = ie.um_WithLongSN_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode um_WithLongSN_Sidelink_r16", err)
		}
	}
	return nil
}

func (ie *RLC_ParametersSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var am_WithLongSN_Sidelink_r16Present bool
	if am_WithLongSN_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var um_WithLongSN_Sidelink_r16Present bool
	if um_WithLongSN_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if am_WithLongSN_Sidelink_r16Present {
		ie.am_WithLongSN_Sidelink_r16 = new(RLC_ParametersSidelink_r16_am_WithLongSN_Sidelink_r16)
		if err = ie.am_WithLongSN_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode am_WithLongSN_Sidelink_r16", err)
		}
	}
	if um_WithLongSN_Sidelink_r16Present {
		ie.um_WithLongSN_Sidelink_r16 = new(RLC_ParametersSidelink_r16_um_WithLongSN_Sidelink_r16)
		if err = ie.um_WithLongSN_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode um_WithLongSN_Sidelink_r16", err)
		}
	}
	return nil
}
