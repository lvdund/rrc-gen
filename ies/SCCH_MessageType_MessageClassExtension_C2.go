package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_MessageClassExtension_C2_Choice_nothing uint64 = iota
	SCCH_MessageType_MessageClassExtension_C2_Choice_notificationMessageSidelink_r17
	SCCH_MessageType_MessageClassExtension_C2_Choice_ueAssistanceInformationSidelink_r17
	SCCH_MessageType_MessageClassExtension_C2_Choice_spare6
	SCCH_MessageType_MessageClassExtension_C2_Choice_spare5
	SCCH_MessageType_MessageClassExtension_C2_Choice_spare4
	SCCH_MessageType_MessageClassExtension_C2_Choice_spare3
	SCCH_MessageType_MessageClassExtension_C2_Choice_spare2
	SCCH_MessageType_MessageClassExtension_C2_Choice_spare1
)

type SCCH_MessageType_MessageClassExtension_C2 struct {
	Choice                              uint64
	notificationMessageSidelink_r17     *NotificationMessageSidelink_r17
	ueAssistanceInformationSidelink_r17 *UEAssistanceInformationSidelink_r17
	spare6                              uper.NULL `madatory`
	spare5                              uper.NULL `madatory`
	spare4                              uper.NULL `madatory`
	spare3                              uper.NULL `madatory`
	spare2                              uper.NULL `madatory`
	spare1                              uper.NULL `madatory`
}

func (ie *SCCH_MessageType_MessageClassExtension_C2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_MessageClassExtension_C2_Choice_notificationMessageSidelink_r17:
		if err = ie.notificationMessageSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode notificationMessageSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_ueAssistanceInformationSidelink_r17:
		if err = ie.ueAssistanceInformationSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode ueAssistanceInformationSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare6:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare6", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare5:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare5", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare4:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare4", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType_MessageClassExtension_C2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_MessageClassExtension_C2_Choice_notificationMessageSidelink_r17:
		ie.notificationMessageSidelink_r17 = new(NotificationMessageSidelink_r17)
		if err = ie.notificationMessageSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode notificationMessageSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_ueAssistanceInformationSidelink_r17:
		ie.ueAssistanceInformationSidelink_r17 = new(UEAssistanceInformationSidelink_r17)
		if err = ie.ueAssistanceInformationSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ueAssistanceInformationSidelink_r17", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare6:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare6", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare5:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare5", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare4:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare4", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case SCCH_MessageType_MessageClassExtension_C2_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
