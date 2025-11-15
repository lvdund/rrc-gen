package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_Q_OffsetRange_Enum_dB_24 uper.Enumerated = 0
	EUTRA_Q_OffsetRange_Enum_dB_22 uper.Enumerated = 1
	EUTRA_Q_OffsetRange_Enum_dB_20 uper.Enumerated = 2
	EUTRA_Q_OffsetRange_Enum_dB_18 uper.Enumerated = 3
	EUTRA_Q_OffsetRange_Enum_dB_16 uper.Enumerated = 4
	EUTRA_Q_OffsetRange_Enum_dB_14 uper.Enumerated = 5
	EUTRA_Q_OffsetRange_Enum_dB_12 uper.Enumerated = 6
	EUTRA_Q_OffsetRange_Enum_dB_10 uper.Enumerated = 7
	EUTRA_Q_OffsetRange_Enum_dB_8  uper.Enumerated = 8
	EUTRA_Q_OffsetRange_Enum_dB_6  uper.Enumerated = 9
	EUTRA_Q_OffsetRange_Enum_dB_5  uper.Enumerated = 10
	EUTRA_Q_OffsetRange_Enum_dB_4  uper.Enumerated = 11
	EUTRA_Q_OffsetRange_Enum_dB_3  uper.Enumerated = 12
	EUTRA_Q_OffsetRange_Enum_dB_2  uper.Enumerated = 13
	EUTRA_Q_OffsetRange_Enum_dB_1  uper.Enumerated = 14
	EUTRA_Q_OffsetRange_Enum_dB0   uper.Enumerated = 15
	EUTRA_Q_OffsetRange_Enum_dB1   uper.Enumerated = 16
	EUTRA_Q_OffsetRange_Enum_dB2   uper.Enumerated = 17
	EUTRA_Q_OffsetRange_Enum_dB3   uper.Enumerated = 18
	EUTRA_Q_OffsetRange_Enum_dB4   uper.Enumerated = 19
	EUTRA_Q_OffsetRange_Enum_dB5   uper.Enumerated = 20
	EUTRA_Q_OffsetRange_Enum_dB6   uper.Enumerated = 21
	EUTRA_Q_OffsetRange_Enum_dB8   uper.Enumerated = 22
	EUTRA_Q_OffsetRange_Enum_dB10  uper.Enumerated = 23
	EUTRA_Q_OffsetRange_Enum_dB12  uper.Enumerated = 24
	EUTRA_Q_OffsetRange_Enum_dB14  uper.Enumerated = 25
	EUTRA_Q_OffsetRange_Enum_dB16  uper.Enumerated = 26
	EUTRA_Q_OffsetRange_Enum_dB18  uper.Enumerated = 27
	EUTRA_Q_OffsetRange_Enum_dB20  uper.Enumerated = 28
	EUTRA_Q_OffsetRange_Enum_dB22  uper.Enumerated = 29
	EUTRA_Q_OffsetRange_Enum_dB24  uper.Enumerated = 30
)

type EUTRA_Q_OffsetRange struct {
	Value uper.Enumerated `lb:0,ub:30,madatory`
}

func (ie *EUTRA_Q_OffsetRange) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode EUTRA_Q_OffsetRange", err)
	}
	return nil
}

func (ie *EUTRA_Q_OffsetRange) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode EUTRA_Q_OffsetRange", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
