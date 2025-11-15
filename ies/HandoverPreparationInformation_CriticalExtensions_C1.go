package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	HandoverPreparationInformation_CriticalExtensions_C1_Choice_nothing uint64 = iota
	HandoverPreparationInformation_CriticalExtensions_C1_Choice_handoverPreparationInformation
	HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare3
	HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare2
	HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare1
)

type HandoverPreparationInformation_CriticalExtensions_C1 struct {
	Choice                         uint64
	handoverPreparationInformation *HandoverPreparationInformation_IEs
	spare3                         uper.NULL `madatory`
	spare2                         uper.NULL `madatory`
	spare1                         uper.NULL `madatory`
}

func (ie *HandoverPreparationInformation_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_handoverPreparationInformation:
		if err = ie.handoverPreparationInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode handoverPreparationInformation", err)
		}
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *HandoverPreparationInformation_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_handoverPreparationInformation:
		ie.handoverPreparationInformation = new(HandoverPreparationInformation_IEs)
		if err = ie.handoverPreparationInformation.Decode(r); err != nil {
			return utils.WrapError("Decode handoverPreparationInformation", err)
		}
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case HandoverPreparationInformation_CriticalExtensions_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
