package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p00 uper.Enumerated = 0
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p01 uper.Enumerated = 1
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p10 uper.Enumerated = 2
	PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power_Enum_p11 uper.Enumerated = 3
)

type PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderDisabled_ptrs_Power", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
