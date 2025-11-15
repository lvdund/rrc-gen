package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULInformationTransfer struct {
	criticalExtensions ULInformationTransfer_CriticalExtensions `madatory`
}

func (ie *ULInformationTransfer) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.criticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode criticalExtensions", err)
	}
	return nil
}

func (ie *ULInformationTransfer) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.criticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode criticalExtensions", err)
	}
	return nil
}
