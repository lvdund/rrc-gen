package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SecurityConfigSMC struct {
	securityAlgorithmConfig SecurityAlgorithmConfig `madatory`
}

func (ie *SecurityConfigSMC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.securityAlgorithmConfig.Encode(w); err != nil {
		return utils.WrapError("Encode securityAlgorithmConfig", err)
	}
	return nil
}

func (ie *SecurityConfigSMC) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.securityAlgorithmConfig.Decode(r); err != nil {
		return utils.WrapError("Decode securityAlgorithmConfig", err)
	}
	return nil
}
