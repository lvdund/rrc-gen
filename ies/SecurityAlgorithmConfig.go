package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SecurityAlgorithmConfig struct {
	cipheringAlgorithm     CipheringAlgorithm      `madatory`
	integrityProtAlgorithm *IntegrityProtAlgorithm `optional`
}

func (ie *SecurityAlgorithmConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.integrityProtAlgorithm != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.cipheringAlgorithm.Encode(w); err != nil {
		return utils.WrapError("Encode cipheringAlgorithm", err)
	}
	if ie.integrityProtAlgorithm != nil {
		if err = ie.integrityProtAlgorithm.Encode(w); err != nil {
			return utils.WrapError("Encode integrityProtAlgorithm", err)
		}
	}
	return nil
}

func (ie *SecurityAlgorithmConfig) Decode(r *uper.UperReader) error {
	var err error
	var integrityProtAlgorithmPresent bool
	if integrityProtAlgorithmPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.cipheringAlgorithm.Decode(r); err != nil {
		return utils.WrapError("Decode cipheringAlgorithm", err)
	}
	if integrityProtAlgorithmPresent {
		ie.integrityProtAlgorithm = new(IntegrityProtAlgorithm)
		if err = ie.integrityProtAlgorithm.Decode(r); err != nil {
			return utils.WrapError("Decode integrityProtAlgorithm", err)
		}
	}
	return nil
}
