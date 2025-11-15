package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersSidelink_r17 struct {
	drx_OnSidelink_r17 *MAC_ParametersSidelink_r17_drx_OnSidelink_r17 `optional`
}

func (ie *MAC_ParametersSidelink_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drx_OnSidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drx_OnSidelink_r17 != nil {
		if err = ie.drx_OnSidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode drx_OnSidelink_r17", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersSidelink_r17) Decode(r *uper.UperReader) error {
	var err error
	var drx_OnSidelink_r17Present bool
	if drx_OnSidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if drx_OnSidelink_r17Present {
		ie.drx_OnSidelink_r17 = new(MAC_ParametersSidelink_r17_drx_OnSidelink_r17)
		if err = ie.drx_OnSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode drx_OnSidelink_r17", err)
		}
	}
	return nil
}
