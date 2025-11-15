package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EstablishmentCause_Enum_emergency          uper.Enumerated = 0
	EstablishmentCause_Enum_highPriorityAccess uper.Enumerated = 1
	EstablishmentCause_Enum_mt_Access          uper.Enumerated = 2
	EstablishmentCause_Enum_mo_Signalling      uper.Enumerated = 3
	EstablishmentCause_Enum_mo_Data            uper.Enumerated = 4
	EstablishmentCause_Enum_mo_VoiceCall       uper.Enumerated = 5
	EstablishmentCause_Enum_mo_VideoCall       uper.Enumerated = 6
	EstablishmentCause_Enum_mo_SMS             uper.Enumerated = 7
	EstablishmentCause_Enum_mps_PriorityAccess uper.Enumerated = 8
	EstablishmentCause_Enum_mcs_PriorityAccess uper.Enumerated = 9
	EstablishmentCause_Enum_spare6             uper.Enumerated = 10
	EstablishmentCause_Enum_spare5             uper.Enumerated = 11
	EstablishmentCause_Enum_spare4             uper.Enumerated = 12
	EstablishmentCause_Enum_spare3             uper.Enumerated = 13
	EstablishmentCause_Enum_spare2             uper.Enumerated = 14
	EstablishmentCause_Enum_spare1             uper.Enumerated = 15
)

type EstablishmentCause struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *EstablishmentCause) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode EstablishmentCause", err)
	}
	return nil
}

func (ie *EstablishmentCause) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode EstablishmentCause", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
