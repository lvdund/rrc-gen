package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding_Enum_d2 uper.Enumerated = 0
)

type PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding", err)
	}
	return nil
}

func (ie *PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PTRS_UplinkConfig_transformPrecoderEnabled_timeDensityTransformPrecoding", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
