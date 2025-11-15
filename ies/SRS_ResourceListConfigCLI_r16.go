package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceListConfigCLI_r16 struct {
	Value []SRS_ResourceConfigCLI_r16 `lb:1,ub:maxNrofCLI_SRS_Resources_r16,madatory`
}

func (ie *SRS_ResourceListConfigCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SRS_ResourceConfigCLI_r16]([]*SRS_ResourceConfigCLI_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCLI_SRS_Resources_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SRS_ResourceListConfigCLI_r16", err)
	}
	return nil
}

func (ie *SRS_ResourceListConfigCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SRS_ResourceConfigCLI_r16]([]*SRS_ResourceConfigCLI_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCLI_SRS_Resources_r16}, false)
	fn := func() *SRS_ResourceConfigCLI_r16 {
		return new(SRS_ResourceConfigCLI_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SRS_ResourceListConfigCLI_r16", err)
	}
	ie.Value = []SRS_ResourceConfigCLI_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
