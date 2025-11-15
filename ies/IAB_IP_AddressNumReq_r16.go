package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressNumReq_r16 struct {
	all_Traffic_NumReq_r16    *int64 `lb:1,ub:8,optional`
	f1_C_Traffic_NumReq_r16   *int64 `lb:1,ub:8,optional`
	f1_U_Traffic_NumReq_r16   *int64 `lb:1,ub:8,optional`
	non_F1_Traffic_NumReq_r16 *int64 `lb:1,ub:8,optional`
}

func (ie *IAB_IP_AddressNumReq_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.all_Traffic_NumReq_r16 != nil, ie.f1_C_Traffic_NumReq_r16 != nil, ie.f1_U_Traffic_NumReq_r16 != nil, ie.non_F1_Traffic_NumReq_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.all_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.all_Traffic_NumReq_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode all_Traffic_NumReq_r16", err)
		}
	}
	if ie.f1_C_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.f1_C_Traffic_NumReq_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode f1_C_Traffic_NumReq_r16", err)
		}
	}
	if ie.f1_U_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.f1_U_Traffic_NumReq_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode f1_U_Traffic_NumReq_r16", err)
		}
	}
	if ie.non_F1_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.non_F1_Traffic_NumReq_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode non_F1_Traffic_NumReq_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressNumReq_r16) Decode(r *uper.UperReader) error {
	var err error
	var all_Traffic_NumReq_r16Present bool
	if all_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var f1_C_Traffic_NumReq_r16Present bool
	if f1_C_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var f1_U_Traffic_NumReq_r16Present bool
	if f1_U_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var non_F1_Traffic_NumReq_r16Present bool
	if non_F1_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if all_Traffic_NumReq_r16Present {
		var tmp_int_all_Traffic_NumReq_r16 int64
		if tmp_int_all_Traffic_NumReq_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode all_Traffic_NumReq_r16", err)
		}
		ie.all_Traffic_NumReq_r16 = &tmp_int_all_Traffic_NumReq_r16
	}
	if f1_C_Traffic_NumReq_r16Present {
		var tmp_int_f1_C_Traffic_NumReq_r16 int64
		if tmp_int_f1_C_Traffic_NumReq_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode f1_C_Traffic_NumReq_r16", err)
		}
		ie.f1_C_Traffic_NumReq_r16 = &tmp_int_f1_C_Traffic_NumReq_r16
	}
	if f1_U_Traffic_NumReq_r16Present {
		var tmp_int_f1_U_Traffic_NumReq_r16 int64
		if tmp_int_f1_U_Traffic_NumReq_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode f1_U_Traffic_NumReq_r16", err)
		}
		ie.f1_U_Traffic_NumReq_r16 = &tmp_int_f1_U_Traffic_NumReq_r16
	}
	if non_F1_Traffic_NumReq_r16Present {
		var tmp_int_non_F1_Traffic_NumReq_r16 int64
		if tmp_int_non_F1_Traffic_NumReq_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode non_F1_Traffic_NumReq_r16", err)
		}
		ie.non_F1_Traffic_NumReq_r16 = &tmp_int_non_F1_Traffic_NumReq_r16
	}
	return nil
}
