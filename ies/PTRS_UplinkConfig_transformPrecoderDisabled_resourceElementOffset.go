package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset_Enum_offset01 uper.Enumerated = 0
	PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset_Enum_offset10 uper.Enumerated = 1
	PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset_Enum_offset11 uper.Enumerated = 2
)

type PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderDisabled_resourceElementOffset", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
