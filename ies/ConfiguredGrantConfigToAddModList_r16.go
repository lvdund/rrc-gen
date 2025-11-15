package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfigToAddModList_r16 struct {
	Value []ConfiguredGrantConfig `lb:1,ub:maxNrofConfiguredGrantConfig_r16,madatory`
}

func (ie *ConfiguredGrantConfigToAddModList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ConfiguredGrantConfig]([]*ConfiguredGrantConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofConfiguredGrantConfig_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfigToAddModList_r16", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfigToAddModList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ConfiguredGrantConfig]([]*ConfiguredGrantConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofConfiguredGrantConfig_r16}, false)
	fn := func() *ConfiguredGrantConfig {
		return new(ConfiguredGrantConfig)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfigToAddModList_r16", err)
	}
	ie.Value = []ConfiguredGrantConfig{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
