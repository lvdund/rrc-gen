package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding_Enum_enabled uper.Enumerated = 0
)

type PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding", err)
	}
	return nil
}

func (ie *PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
