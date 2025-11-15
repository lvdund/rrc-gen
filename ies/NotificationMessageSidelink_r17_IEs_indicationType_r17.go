package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_Uu_RLF          uper.Enumerated = 0
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_HO              uper.Enumerated = 1
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_CellReselection uper.Enumerated = 2
	NotificationMessageSidelink_r17_IEs_indicationType_r17_Enum_relayUE_Uu_RRC_Failure  uper.Enumerated = 3
)

type NotificationMessageSidelink_r17_IEs_indicationType_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *NotificationMessageSidelink_r17_IEs_indicationType_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode NotificationMessageSidelink_r17_IEs_indicationType_r17", err)
	}
	return nil
}

func (ie *NotificationMessageSidelink_r17_IEs_indicationType_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode NotificationMessageSidelink_r17_IEs_indicationType_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
