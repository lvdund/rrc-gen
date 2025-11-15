package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULInformationTransferMRDC_CriticalExtensions_C1_Choice_nothing uint64 = iota
	ULInformationTransferMRDC_CriticalExtensions_C1_Choice_ulInformationTransferMRDC
	ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare3
	ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare2
	ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare1
)

type ULInformationTransferMRDC_CriticalExtensions_C1 struct {
	Choice                    uint64
	ulInformationTransferMRDC *ULInformationTransferMRDC_IEs
	spare3                    uper.NULL `madatory`
	spare2                    uper.NULL `madatory`
	spare1                    uper.NULL `madatory`
}

func (ie *ULInformationTransferMRDC_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_ulInformationTransferMRDC:
		if err = ie.ulInformationTransferMRDC.Encode(w); err != nil {
			err = utils.WrapError("Encode ulInformationTransferMRDC", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ULInformationTransferMRDC_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_ulInformationTransferMRDC:
		ie.ulInformationTransferMRDC = new(ULInformationTransferMRDC_IEs)
		if err = ie.ulInformationTransferMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode ulInformationTransferMRDC", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case ULInformationTransferMRDC_CriticalExtensions_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
