package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationSidelink struct {
	rrc_TransactionIdentifier_r16 RRC_TransactionIdentifier                     `madatory`
	criticalExtensions            RRCReconfigurationSidelink_CriticalExtensions `madatory`
}

func (ie *RRCReconfigurationSidelink) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rrc_TransactionIdentifier_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rrc_TransactionIdentifier_r16", err)
	}
	if err = ie.criticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode criticalExtensions", err)
	}
	return nil
}

func (ie *RRCReconfigurationSidelink) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rrc_TransactionIdentifier_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rrc_TransactionIdentifier_r16", err)
	}
	if err = ie.criticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode criticalExtensions", err)
	}
	return nil
}
