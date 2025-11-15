package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_BandwidthClassNR_Enum_a         uper.Enumerated = 0
	CA_BandwidthClassNR_Enum_b         uper.Enumerated = 1
	CA_BandwidthClassNR_Enum_c         uper.Enumerated = 2
	CA_BandwidthClassNR_Enum_d         uper.Enumerated = 3
	CA_BandwidthClassNR_Enum_e         uper.Enumerated = 4
	CA_BandwidthClassNR_Enum_f         uper.Enumerated = 5
	CA_BandwidthClassNR_Enum_g         uper.Enumerated = 6
	CA_BandwidthClassNR_Enum_h         uper.Enumerated = 7
	CA_BandwidthClassNR_Enum_i         uper.Enumerated = 8
	CA_BandwidthClassNR_Enum_j         uper.Enumerated = 9
	CA_BandwidthClassNR_Enum_k         uper.Enumerated = 10
	CA_BandwidthClassNR_Enum_l         uper.Enumerated = 11
	CA_BandwidthClassNR_Enum_m         uper.Enumerated = 12
	CA_BandwidthClassNR_Enum_n         uper.Enumerated = 13
	CA_BandwidthClassNR_Enum_o         uper.Enumerated = 14
	CA_BandwidthClassNR_Enum_p         uper.Enumerated = 15
	CA_BandwidthClassNR_Enum_q         uper.Enumerated = 16
	CA_BandwidthClassNR_Enum_r2_v1730  uper.Enumerated = 17
	CA_BandwidthClassNR_Enum_r3_v1730  uper.Enumerated = 18
	CA_BandwidthClassNR_Enum_r4_v1730  uper.Enumerated = 19
	CA_BandwidthClassNR_Enum_r5_v1730  uper.Enumerated = 20
	CA_BandwidthClassNR_Enum_r6_v1730  uper.Enumerated = 21
	CA_BandwidthClassNR_Enum_r7_v1730  uper.Enumerated = 22
	CA_BandwidthClassNR_Enum_r8_v1730  uper.Enumerated = 23
	CA_BandwidthClassNR_Enum_r9_v1730  uper.Enumerated = 24
	CA_BandwidthClassNR_Enum_r10_v1730 uper.Enumerated = 25
	CA_BandwidthClassNR_Enum_r11_v1730 uper.Enumerated = 26
	CA_BandwidthClassNR_Enum_r12_v1730 uper.Enumerated = 27
)

type CA_BandwidthClassNR struct {
	Value uper.Enumerated `lb:0,ub:16,madatory`
}

func (ie *CA_BandwidthClassNR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode CA_BandwidthClassNR", err)
	}
	return nil
}

func (ie *CA_BandwidthClassNR) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode CA_BandwidthClassNR", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
