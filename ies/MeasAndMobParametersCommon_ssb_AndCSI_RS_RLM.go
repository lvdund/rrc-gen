package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM_Enum_supported uper.Enumerated = 0
)

type MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
