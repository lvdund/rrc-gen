package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ResumeCause_Enum_emergency          uper.Enumerated = 0
	ResumeCause_Enum_highPriorityAccess uper.Enumerated = 1
	ResumeCause_Enum_mt_Access          uper.Enumerated = 2
	ResumeCause_Enum_mo_Signalling      uper.Enumerated = 3
	ResumeCause_Enum_mo_Data            uper.Enumerated = 4
	ResumeCause_Enum_mo_VoiceCall       uper.Enumerated = 5
	ResumeCause_Enum_mo_VideoCall       uper.Enumerated = 6
	ResumeCause_Enum_mo_SMS             uper.Enumerated = 7
	ResumeCause_Enum_rna_Update         uper.Enumerated = 8
	ResumeCause_Enum_mps_PriorityAccess uper.Enumerated = 9
	ResumeCause_Enum_mcs_PriorityAccess uper.Enumerated = 10
	ResumeCause_Enum_spare1             uper.Enumerated = 11
	ResumeCause_Enum_spare2             uper.Enumerated = 12
	ResumeCause_Enum_spare3             uper.Enumerated = 13
	ResumeCause_Enum_spare4             uper.Enumerated = 14
	ResumeCause_Enum_spare5             uper.Enumerated = 15
)

type ResumeCause struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *ResumeCause) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode ResumeCause", err)
	}
	return nil
}

func (ie *ResumeCause) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode ResumeCause", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
