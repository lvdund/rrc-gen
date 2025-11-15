package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressAndTraffic_r16 struct {
	all_Traffic_IAB_IP_Address_r16 []IAB_IP_Address_r16 `lb:1,ub:8,optional`
	f1_C_Traffic_IP_Address_r16    []IAB_IP_Address_r16 `lb:1,ub:8,optional`
	f1_U_Traffic_IP_Address_r16    []IAB_IP_Address_r16 `lb:1,ub:8,optional`
	non_F1_Traffic_IP_Address_r16  []IAB_IP_Address_r16 `lb:1,ub:8,optional`
}

func (ie *IAB_IP_AddressAndTraffic_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.all_Traffic_IAB_IP_Address_r16) > 0, len(ie.f1_C_Traffic_IP_Address_r16) > 0, len(ie.f1_U_Traffic_IP_Address_r16) > 0, len(ie.non_F1_Traffic_IP_Address_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.all_Traffic_IAB_IP_Address_r16) > 0 {
		tmp_all_Traffic_IAB_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.all_Traffic_IAB_IP_Address_r16 {
			tmp_all_Traffic_IAB_IP_Address_r16.Value = append(tmp_all_Traffic_IAB_IP_Address_r16.Value, &i)
		}
		if err = tmp_all_Traffic_IAB_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode all_Traffic_IAB_IP_Address_r16", err)
		}
	}
	if len(ie.f1_C_Traffic_IP_Address_r16) > 0 {
		tmp_f1_C_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.f1_C_Traffic_IP_Address_r16 {
			tmp_f1_C_Traffic_IP_Address_r16.Value = append(tmp_f1_C_Traffic_IP_Address_r16.Value, &i)
		}
		if err = tmp_f1_C_Traffic_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode f1_C_Traffic_IP_Address_r16", err)
		}
	}
	if len(ie.f1_U_Traffic_IP_Address_r16) > 0 {
		tmp_f1_U_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.f1_U_Traffic_IP_Address_r16 {
			tmp_f1_U_Traffic_IP_Address_r16.Value = append(tmp_f1_U_Traffic_IP_Address_r16.Value, &i)
		}
		if err = tmp_f1_U_Traffic_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode f1_U_Traffic_IP_Address_r16", err)
		}
	}
	if len(ie.non_F1_Traffic_IP_Address_r16) > 0 {
		tmp_non_F1_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.non_F1_Traffic_IP_Address_r16 {
			tmp_non_F1_Traffic_IP_Address_r16.Value = append(tmp_non_F1_Traffic_IP_Address_r16.Value, &i)
		}
		if err = tmp_non_F1_Traffic_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode non_F1_Traffic_IP_Address_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressAndTraffic_r16) Decode(r *uper.UperReader) error {
	var err error
	var all_Traffic_IAB_IP_Address_r16Present bool
	if all_Traffic_IAB_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var f1_C_Traffic_IP_Address_r16Present bool
	if f1_C_Traffic_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var f1_U_Traffic_IP_Address_r16Present bool
	if f1_U_Traffic_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var non_F1_Traffic_IP_Address_r16Present bool
	if non_F1_Traffic_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if all_Traffic_IAB_IP_Address_r16Present {
		tmp_all_Traffic_IAB_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_all_Traffic_IAB_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_all_Traffic_IAB_IP_Address_r16.Decode(r, fn_all_Traffic_IAB_IP_Address_r16); err != nil {
			return utils.WrapError("Decode all_Traffic_IAB_IP_Address_r16", err)
		}
		ie.all_Traffic_IAB_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_all_Traffic_IAB_IP_Address_r16.Value {
			ie.all_Traffic_IAB_IP_Address_r16 = append(ie.all_Traffic_IAB_IP_Address_r16, *i)
		}
	}
	if f1_C_Traffic_IP_Address_r16Present {
		tmp_f1_C_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_f1_C_Traffic_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_f1_C_Traffic_IP_Address_r16.Decode(r, fn_f1_C_Traffic_IP_Address_r16); err != nil {
			return utils.WrapError("Decode f1_C_Traffic_IP_Address_r16", err)
		}
		ie.f1_C_Traffic_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_f1_C_Traffic_IP_Address_r16.Value {
			ie.f1_C_Traffic_IP_Address_r16 = append(ie.f1_C_Traffic_IP_Address_r16, *i)
		}
	}
	if f1_U_Traffic_IP_Address_r16Present {
		tmp_f1_U_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_f1_U_Traffic_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_f1_U_Traffic_IP_Address_r16.Decode(r, fn_f1_U_Traffic_IP_Address_r16); err != nil {
			return utils.WrapError("Decode f1_U_Traffic_IP_Address_r16", err)
		}
		ie.f1_U_Traffic_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_f1_U_Traffic_IP_Address_r16.Value {
			ie.f1_U_Traffic_IP_Address_r16 = append(ie.f1_U_Traffic_IP_Address_r16, *i)
		}
	}
	if non_F1_Traffic_IP_Address_r16Present {
		tmp_non_F1_Traffic_IP_Address_r16 := utils.NewSequence[*IAB_IP_Address_r16]([]*IAB_IP_Address_r16{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_non_F1_Traffic_IP_Address_r16 := func() *IAB_IP_Address_r16 {
			return new(IAB_IP_Address_r16)
		}
		if err = tmp_non_F1_Traffic_IP_Address_r16.Decode(r, fn_non_F1_Traffic_IP_Address_r16); err != nil {
			return utils.WrapError("Decode non_F1_Traffic_IP_Address_r16", err)
		}
		ie.non_F1_Traffic_IP_Address_r16 = []IAB_IP_Address_r16{}
		for _, i := range tmp_non_F1_Traffic_IP_Address_r16.Value {
			ie.non_F1_Traffic_IP_Address_r16 = append(ie.non_F1_Traffic_IP_Address_r16, *i)
		}
	}
	return nil
}
