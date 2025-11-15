package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressPrefixReq_r16 struct {
	all_Traffic_PrefixReq_r16    *IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16    `optional`
	f1_C_Traffic_PrefixReq_r16   *IAB_IP_AddressPrefixReq_r16_f1_C_Traffic_PrefixReq_r16   `optional`
	f1_U_Traffic_PrefixReq_r16   *IAB_IP_AddressPrefixReq_r16_f1_U_Traffic_PrefixReq_r16   `optional`
	non_F1_Traffic_PrefixReq_r16 *IAB_IP_AddressPrefixReq_r16_non_F1_Traffic_PrefixReq_r16 `optional`
}

func (ie *IAB_IP_AddressPrefixReq_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.all_Traffic_PrefixReq_r16 != nil, ie.f1_C_Traffic_PrefixReq_r16 != nil, ie.f1_U_Traffic_PrefixReq_r16 != nil, ie.non_F1_Traffic_PrefixReq_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.all_Traffic_PrefixReq_r16 != nil {
		if err = ie.all_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode all_Traffic_PrefixReq_r16", err)
		}
	}
	if ie.f1_C_Traffic_PrefixReq_r16 != nil {
		if err = ie.f1_C_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode f1_C_Traffic_PrefixReq_r16", err)
		}
	}
	if ie.f1_U_Traffic_PrefixReq_r16 != nil {
		if err = ie.f1_U_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode f1_U_Traffic_PrefixReq_r16", err)
		}
	}
	if ie.non_F1_Traffic_PrefixReq_r16 != nil {
		if err = ie.non_F1_Traffic_PrefixReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode non_F1_Traffic_PrefixReq_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressPrefixReq_r16) Decode(r *uper.UperReader) error {
	var err error
	var all_Traffic_PrefixReq_r16Present bool
	if all_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var f1_C_Traffic_PrefixReq_r16Present bool
	if f1_C_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var f1_U_Traffic_PrefixReq_r16Present bool
	if f1_U_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var non_F1_Traffic_PrefixReq_r16Present bool
	if non_F1_Traffic_PrefixReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if all_Traffic_PrefixReq_r16Present {
		ie.all_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_all_Traffic_PrefixReq_r16)
		if err = ie.all_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode all_Traffic_PrefixReq_r16", err)
		}
	}
	if f1_C_Traffic_PrefixReq_r16Present {
		ie.f1_C_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_f1_C_Traffic_PrefixReq_r16)
		if err = ie.f1_C_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode f1_C_Traffic_PrefixReq_r16", err)
		}
	}
	if f1_U_Traffic_PrefixReq_r16Present {
		ie.f1_U_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_f1_U_Traffic_PrefixReq_r16)
		if err = ie.f1_U_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode f1_U_Traffic_PrefixReq_r16", err)
		}
	}
	if non_F1_Traffic_PrefixReq_r16Present {
		ie.non_F1_Traffic_PrefixReq_r16 = new(IAB_IP_AddressPrefixReq_r16_non_F1_Traffic_PrefixReq_r16)
		if err = ie.non_F1_Traffic_PrefixReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode non_F1_Traffic_PrefixReq_r16", err)
		}
	}
	return nil
}
