package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookVariantsList_r16 struct {
	Value []SupportedCSI_RS_Resource `lb:1,ub:maxNrofCSI_RS_ResourcesAlt_r16,madatory`
}

func (ie *CodebookVariantsList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesAlt_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CodebookVariantsList_r16", err)
	}
	return nil
}

func (ie *CodebookVariantsList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesAlt_r16}, false)
	fn := func() *SupportedCSI_RS_Resource {
		return new(SupportedCSI_RS_Resource)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CodebookVariantsList_r16", err)
	}
	ie.Value = []SupportedCSI_RS_Resource{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
