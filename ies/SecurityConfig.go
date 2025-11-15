package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SecurityConfig struct {
	securityAlgorithmConfig *SecurityAlgorithmConfig `optional`
	keyToUse                *SecurityConfig_keyToUse `optional`
}

func (ie *SecurityConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.securityAlgorithmConfig != nil, ie.keyToUse != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.securityAlgorithmConfig != nil {
		if err = ie.securityAlgorithmConfig.Encode(w); err != nil {
			return utils.WrapError("Encode securityAlgorithmConfig", err)
		}
	}
	if ie.keyToUse != nil {
		if err = ie.keyToUse.Encode(w); err != nil {
			return utils.WrapError("Encode keyToUse", err)
		}
	}
	return nil
}

func (ie *SecurityConfig) Decode(r *uper.UperReader) error {
	var err error
	var securityAlgorithmConfigPresent bool
	if securityAlgorithmConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var keyToUsePresent bool
	if keyToUsePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if securityAlgorithmConfigPresent {
		ie.securityAlgorithmConfig = new(SecurityAlgorithmConfig)
		if err = ie.securityAlgorithmConfig.Decode(r); err != nil {
			return utils.WrapError("Decode securityAlgorithmConfig", err)
		}
	}
	if keyToUsePresent {
		ie.keyToUse = new(SecurityConfig_keyToUse)
		if err = ie.keyToUse.Decode(r); err != nil {
			return utils.WrapError("Decode keyToUse", err)
		}
	}
	return nil
}
