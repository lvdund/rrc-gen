package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts_Enum_n1 uper.Enumerated = 0
	PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts_Enum_n2 uper.Enumerated = 1
)

type PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderDisabled_maxNrofPorts", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
