package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULInformationTransfer_CriticalExtensions_Choice_nothing uint64 = iota
	ULInformationTransfer_CriticalExtensions_Choice_ulInformationTransfer
	ULInformationTransfer_CriticalExtensions_Choice_criticalExtensionsFuture
)

type ULInformationTransfer_CriticalExtensions struct {
	Choice                   uint64
	ulInformationTransfer    *ULInformationTransfer_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *ULInformationTransfer_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransfer_CriticalExtensions_Choice_ulInformationTransfer:
		if err = ie.ulInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode ulInformationTransfer", err)
		}
	case ULInformationTransfer_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ULInformationTransfer_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransfer_CriticalExtensions_Choice_ulInformationTransfer:
		ie.ulInformationTransfer = new(ULInformationTransfer_IEs)
		if err = ie.ulInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode ulInformationTransfer", err)
		}
	case ULInformationTransfer_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
