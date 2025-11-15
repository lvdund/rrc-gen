package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PowerControl_v1610_olpc_ParameterSet struct {
	olpc_ParameterSetDCI_0_1_r16 *int64 `lb:1,ub:2,optional`
	olpc_ParameterSetDCI_0_2_r16 *int64 `lb:1,ub:2,optional`
}

func (ie *PUSCH_PowerControl_v1610_olpc_ParameterSet) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.olpc_ParameterSetDCI_0_1_r16 != nil, ie.olpc_ParameterSetDCI_0_2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.olpc_ParameterSetDCI_0_1_r16 != nil {
		if err = w.WriteInteger(*ie.olpc_ParameterSetDCI_0_1_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode olpc_ParameterSetDCI_0_1_r16", err)
		}
	}
	if ie.olpc_ParameterSetDCI_0_2_r16 != nil {
		if err = w.WriteInteger(*ie.olpc_ParameterSetDCI_0_2_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode olpc_ParameterSetDCI_0_2_r16", err)
		}
	}
	return nil
}

func (ie *PUSCH_PowerControl_v1610_olpc_ParameterSet) Decode(r *uper.UperReader) error {
	var err error
	var olpc_ParameterSetDCI_0_1_r16Present bool
	if olpc_ParameterSetDCI_0_1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var olpc_ParameterSetDCI_0_2_r16Present bool
	if olpc_ParameterSetDCI_0_2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if olpc_ParameterSetDCI_0_1_r16Present {
		var tmp_int_olpc_ParameterSetDCI_0_1_r16 int64
		if tmp_int_olpc_ParameterSetDCI_0_1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode olpc_ParameterSetDCI_0_1_r16", err)
		}
		ie.olpc_ParameterSetDCI_0_1_r16 = &tmp_int_olpc_ParameterSetDCI_0_1_r16
	}
	if olpc_ParameterSetDCI_0_2_r16Present {
		var tmp_int_olpc_ParameterSetDCI_0_2_r16 int64
		if tmp_int_olpc_ParameterSetDCI_0_2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode olpc_ParameterSetDCI_0_2_r16", err)
		}
		ie.olpc_ParameterSetDCI_0_2_r16 = &tmp_int_olpc_ParameterSetDCI_0_2_r16
	}
	return nil
}
