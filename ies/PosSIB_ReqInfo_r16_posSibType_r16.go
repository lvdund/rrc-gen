package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_1        uper.Enumerated = 0
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_2        uper.Enumerated = 1
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_3        uper.Enumerated = 2
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_4        uper.Enumerated = 3
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_5        uper.Enumerated = 4
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_6        uper.Enumerated = 5
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_7        uper.Enumerated = 6
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_8        uper.Enumerated = 7
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_1        uper.Enumerated = 8
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_2        uper.Enumerated = 9
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_3        uper.Enumerated = 10
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_4        uper.Enumerated = 11
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_5        uper.Enumerated = 12
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_6        uper.Enumerated = 13
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_7        uper.Enumerated = 14
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_8        uper.Enumerated = 15
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_9        uper.Enumerated = 16
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_10       uper.Enumerated = 17
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_11       uper.Enumerated = 18
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_12       uper.Enumerated = 19
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_13       uper.Enumerated = 20
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_14       uper.Enumerated = 21
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_15       uper.Enumerated = 22
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_16       uper.Enumerated = 23
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_17       uper.Enumerated = 24
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_18       uper.Enumerated = 25
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_19       uper.Enumerated = 26
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_20       uper.Enumerated = 27
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_21       uper.Enumerated = 28
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_22       uper.Enumerated = 29
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_23       uper.Enumerated = 30
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType3_1        uper.Enumerated = 31
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType4_1        uper.Enumerated = 32
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType5_1        uper.Enumerated = 33
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_1        uper.Enumerated = 34
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_2        uper.Enumerated = 35
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_3        uper.Enumerated = 36
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_9_v1710  uper.Enumerated = 37
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType1_10_v1710 uper.Enumerated = 38
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_24_v1710 uper.Enumerated = 39
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType2_25_v1710 uper.Enumerated = 40
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_4_v1710  uper.Enumerated = 41
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_5_v1710  uper.Enumerated = 42
	PosSIB_ReqInfo_r16_posSibType_r16_Enum_posSibType6_6_v1710  uper.Enumerated = 43
)

type PosSIB_ReqInfo_r16_posSibType_r16 struct {
	Value uper.Enumerated `lb:0,ub:36,madatory`
}

func (ie *PosSIB_ReqInfo_r16_posSibType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 36}, false); err != nil {
		return utils.WrapError("Encode PosSIB_ReqInfo_r16_posSibType_r16", err)
	}
	return nil
}

func (ie *PosSIB_ReqInfo_r16_posSibType_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 36}, false); err != nil {
		return utils.WrapError("Decode PosSIB_ReqInfo_r16_posSibType_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
