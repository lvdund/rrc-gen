package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectCLI_r16 struct {
	cli_ResourceConfig_r16 CLI_ResourceConfig_r16 `madatory`
}

func (ie *MeasObjectCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cli_ResourceConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cli_ResourceConfig_r16", err)
	}
	return nil
}

func (ie *MeasObjectCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cli_ResourceConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cli_ResourceConfig_r16", err)
	}
	return nil
}
