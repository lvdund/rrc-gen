package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UuMessageTransferSidelink_r17_CriticalExtensions_Choice_nothing uint64 = iota
	UuMessageTransferSidelink_r17_CriticalExtensions_Choice_uuMessageTransferSidelink_r17
	UuMessageTransferSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UuMessageTransferSidelink_r17_CriticalExtensions struct {
	Choice                        uint64
	uuMessageTransferSidelink_r17 *UuMessageTransferSidelink_r17_IEs
	criticalExtensionsFuture      interface{} `madatory`
}

func (ie *UuMessageTransferSidelink_r17_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UuMessageTransferSidelink_r17_CriticalExtensions_Choice_uuMessageTransferSidelink_r17:
		if err = ie.uuMessageTransferSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode uuMessageTransferSidelink_r17", err)
		}
	case UuMessageTransferSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UuMessageTransferSidelink_r17_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UuMessageTransferSidelink_r17_CriticalExtensions_Choice_uuMessageTransferSidelink_r17:
		ie.uuMessageTransferSidelink_r17 = new(UuMessageTransferSidelink_r17_IEs)
		if err = ie.uuMessageTransferSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uuMessageTransferSidelink_r17", err)
		}
	case UuMessageTransferSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
