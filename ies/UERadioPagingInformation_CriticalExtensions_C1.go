package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UERadioPagingInformation_CriticalExtensions_C1_Choice_nothing uint64 = iota
	UERadioPagingInformation_CriticalExtensions_C1_Choice_ueRadioPagingInformation
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare7
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare6
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare5
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare4
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare3
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare2
	UERadioPagingInformation_CriticalExtensions_C1_Choice_spare1
)

type UERadioPagingInformation_CriticalExtensions_C1 struct {
	Choice                   uint64
	ueRadioPagingInformation *UERadioPagingInformation_IEs
	spare7                   uper.NULL `madatory`
	spare6                   uper.NULL `madatory`
	spare5                   uper.NULL `madatory`
	spare4                   uper.NULL `madatory`
	spare3                   uper.NULL `madatory`
	spare2                   uper.NULL `madatory`
	spare1                   uper.NULL `madatory`
}

func (ie *UERadioPagingInformation_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_ueRadioPagingInformation:
		if err = ie.ueRadioPagingInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ueRadioPagingInformation", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare7:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare7", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare6:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare6", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare5:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare5", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare4:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare4", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UERadioPagingInformation_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_ueRadioPagingInformation:
		ie.ueRadioPagingInformation = new(UERadioPagingInformation_IEs)
		if err = ie.ueRadioPagingInformation.Decode(r); err != nil {
			return utils.WrapError("Decode ueRadioPagingInformation", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare7:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare7", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare6:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare6", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare5:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare5", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare4:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare4", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case UERadioPagingInformation_CriticalExtensions_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
