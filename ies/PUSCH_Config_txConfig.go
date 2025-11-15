package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Config_txConfig_Enum_codebook    uper.Enumerated = 0
	PUSCH_Config_txConfig_Enum_nonCodebook uper.Enumerated = 1
)

type PUSCH_Config_txConfig struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PUSCH_Config_txConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Config_txConfig", err)
	}
	return nil
}

func (ie *PUSCH_Config_txConfig) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Config_txConfig", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
