package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n1 uper.Enumerated = 0
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n2 uper.Enumerated = 1
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n3 uper.Enumerated = 2
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n4 uper.Enumerated = 3
	MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16_Enum_n7 uper.Enumerated = 4
)

type MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16", err)
	}
	return nil
}

func (ie *MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
