package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SecurityModeFailure struct {
	rrc_TransactionIdentifier RRC_TransactionIdentifier              `madatory`
	criticalExtensions        SecurityModeFailure_CriticalExtensions `madatory`
}

func (ie *SecurityModeFailure) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rrc_TransactionIdentifier.Encode(w); err != nil {
		return utils.WrapError("Encode rrc_TransactionIdentifier", err)
	}
	if err = ie.criticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode criticalExtensions", err)
	}
	return nil
}

func (ie *SecurityModeFailure) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rrc_TransactionIdentifier.Decode(r); err != nil {
		return utils.WrapError("Decode rrc_TransactionIdentifier", err)
	}
	if err = ie.criticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode criticalExtensions", err)
	}
	return nil
}
