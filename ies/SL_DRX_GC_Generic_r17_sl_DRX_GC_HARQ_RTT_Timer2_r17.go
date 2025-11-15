package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_sl0    uper.Enumerated = 0
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_sl1    uper.Enumerated = 1
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_sl2    uper.Enumerated = 2
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_sl4    uper.Enumerated = 3
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_spare4 uper.Enumerated = 4
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_spare3 uper.Enumerated = 5
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_spare2 uper.Enumerated = 6
	SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17_Enum_spare1 uper.Enumerated = 7
)

type SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17", err)
	}
	return nil
}

func (ie *SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_DRX_GC_Generic_r17_sl_DRX_GC_HARQ_RTT_Timer2_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
