package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultUTRA_FDD_r16_measResult_r16 struct {
	utra_FDD_RSCP_r16 *int64 `lb:-5,ub:91,optional`
	utra_FDD_EcN0_r16 *int64 `lb:0,ub:49,optional`
}

func (ie *MeasResultUTRA_FDD_r16_measResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.utra_FDD_RSCP_r16 != nil, ie.utra_FDD_EcN0_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.utra_FDD_RSCP_r16 != nil {
		if err = w.WriteInteger(*ie.utra_FDD_RSCP_r16, &uper.Constraint{Lb: -5, Ub: 91}, false); err != nil {
			return utils.WrapError("Encode utra_FDD_RSCP_r16", err)
		}
	}
	if ie.utra_FDD_EcN0_r16 != nil {
		if err = w.WriteInteger(*ie.utra_FDD_EcN0_r16, &uper.Constraint{Lb: 0, Ub: 49}, false); err != nil {
			return utils.WrapError("Encode utra_FDD_EcN0_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultUTRA_FDD_r16_measResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var utra_FDD_RSCP_r16Present bool
	if utra_FDD_RSCP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var utra_FDD_EcN0_r16Present bool
	if utra_FDD_EcN0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if utra_FDD_RSCP_r16Present {
		var tmp_int_utra_FDD_RSCP_r16 int64
		if tmp_int_utra_FDD_RSCP_r16, err = r.ReadInteger(&uper.Constraint{Lb: -5, Ub: 91}, false); err != nil {
			return utils.WrapError("Decode utra_FDD_RSCP_r16", err)
		}
		ie.utra_FDD_RSCP_r16 = &tmp_int_utra_FDD_RSCP_r16
	}
	if utra_FDD_EcN0_r16Present {
		var tmp_int_utra_FDD_EcN0_r16 int64
		if tmp_int_utra_FDD_EcN0_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 49}, false); err != nil {
			return utils.WrapError("Decode utra_FDD_EcN0_r16", err)
		}
		ie.utra_FDD_EcN0_r16 = &tmp_int_utra_FDD_EcN0_r16
	}
	return nil
}
