package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Q_OffsetRange_Enum_dB_24 uper.Enumerated = 0
	Q_OffsetRange_Enum_dB_22 uper.Enumerated = 1
	Q_OffsetRange_Enum_dB_20 uper.Enumerated = 2
	Q_OffsetRange_Enum_dB_18 uper.Enumerated = 3
	Q_OffsetRange_Enum_dB_16 uper.Enumerated = 4
	Q_OffsetRange_Enum_dB_14 uper.Enumerated = 5
	Q_OffsetRange_Enum_dB_12 uper.Enumerated = 6
	Q_OffsetRange_Enum_dB_10 uper.Enumerated = 7
	Q_OffsetRange_Enum_dB_8  uper.Enumerated = 8
	Q_OffsetRange_Enum_dB_6  uper.Enumerated = 9
	Q_OffsetRange_Enum_dB_5  uper.Enumerated = 10
	Q_OffsetRange_Enum_dB_4  uper.Enumerated = 11
	Q_OffsetRange_Enum_dB_3  uper.Enumerated = 12
	Q_OffsetRange_Enum_dB_2  uper.Enumerated = 13
	Q_OffsetRange_Enum_dB_1  uper.Enumerated = 14
	Q_OffsetRange_Enum_dB0   uper.Enumerated = 15
	Q_OffsetRange_Enum_dB1   uper.Enumerated = 16
	Q_OffsetRange_Enum_dB2   uper.Enumerated = 17
	Q_OffsetRange_Enum_dB3   uper.Enumerated = 18
	Q_OffsetRange_Enum_dB4   uper.Enumerated = 19
	Q_OffsetRange_Enum_dB5   uper.Enumerated = 20
	Q_OffsetRange_Enum_dB6   uper.Enumerated = 21
	Q_OffsetRange_Enum_dB8   uper.Enumerated = 22
	Q_OffsetRange_Enum_dB10  uper.Enumerated = 23
	Q_OffsetRange_Enum_dB12  uper.Enumerated = 24
	Q_OffsetRange_Enum_dB14  uper.Enumerated = 25
	Q_OffsetRange_Enum_dB16  uper.Enumerated = 26
	Q_OffsetRange_Enum_dB18  uper.Enumerated = 27
	Q_OffsetRange_Enum_dB20  uper.Enumerated = 28
	Q_OffsetRange_Enum_dB22  uper.Enumerated = 29
	Q_OffsetRange_Enum_dB24  uper.Enumerated = 30
)

type Q_OffsetRange struct {
	Value uper.Enumerated `lb:0,ub:30,madatory`
}

func (ie *Q_OffsetRange) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode Q_OffsetRange", err)
	}
	return nil
}

func (ie *Q_OffsetRange) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode Q_OffsetRange", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
