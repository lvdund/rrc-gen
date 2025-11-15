package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PollByte_Enum_kB1      uper.Enumerated = 0
	PollByte_Enum_kB2      uper.Enumerated = 1
	PollByte_Enum_kB5      uper.Enumerated = 2
	PollByte_Enum_kB8      uper.Enumerated = 3
	PollByte_Enum_kB10     uper.Enumerated = 4
	PollByte_Enum_kB15     uper.Enumerated = 5
	PollByte_Enum_kB25     uper.Enumerated = 6
	PollByte_Enum_kB50     uper.Enumerated = 7
	PollByte_Enum_kB75     uper.Enumerated = 8
	PollByte_Enum_kB100    uper.Enumerated = 9
	PollByte_Enum_kB125    uper.Enumerated = 10
	PollByte_Enum_kB250    uper.Enumerated = 11
	PollByte_Enum_kB375    uper.Enumerated = 12
	PollByte_Enum_kB500    uper.Enumerated = 13
	PollByte_Enum_kB750    uper.Enumerated = 14
	PollByte_Enum_kB1000   uper.Enumerated = 15
	PollByte_Enum_kB1250   uper.Enumerated = 16
	PollByte_Enum_kB1500   uper.Enumerated = 17
	PollByte_Enum_kB2000   uper.Enumerated = 18
	PollByte_Enum_kB3000   uper.Enumerated = 19
	PollByte_Enum_kB4000   uper.Enumerated = 20
	PollByte_Enum_kB4500   uper.Enumerated = 21
	PollByte_Enum_kB5000   uper.Enumerated = 22
	PollByte_Enum_kB5500   uper.Enumerated = 23
	PollByte_Enum_kB6000   uper.Enumerated = 24
	PollByte_Enum_kB6500   uper.Enumerated = 25
	PollByte_Enum_kB7000   uper.Enumerated = 26
	PollByte_Enum_kB7500   uper.Enumerated = 27
	PollByte_Enum_mB8      uper.Enumerated = 28
	PollByte_Enum_mB9      uper.Enumerated = 29
	PollByte_Enum_mB10     uper.Enumerated = 30
	PollByte_Enum_mB11     uper.Enumerated = 31
	PollByte_Enum_mB12     uper.Enumerated = 32
	PollByte_Enum_mB13     uper.Enumerated = 33
	PollByte_Enum_mB14     uper.Enumerated = 34
	PollByte_Enum_mB15     uper.Enumerated = 35
	PollByte_Enum_mB16     uper.Enumerated = 36
	PollByte_Enum_mB17     uper.Enumerated = 37
	PollByte_Enum_mB18     uper.Enumerated = 38
	PollByte_Enum_mB20     uper.Enumerated = 39
	PollByte_Enum_mB25     uper.Enumerated = 40
	PollByte_Enum_mB30     uper.Enumerated = 41
	PollByte_Enum_mB40     uper.Enumerated = 42
	PollByte_Enum_infinity uper.Enumerated = 43
	PollByte_Enum_spare20  uper.Enumerated = 44
	PollByte_Enum_spare19  uper.Enumerated = 45
	PollByte_Enum_spare18  uper.Enumerated = 46
	PollByte_Enum_spare17  uper.Enumerated = 47
	PollByte_Enum_spare16  uper.Enumerated = 48
	PollByte_Enum_spare15  uper.Enumerated = 49
	PollByte_Enum_spare14  uper.Enumerated = 50
	PollByte_Enum_spare13  uper.Enumerated = 51
	PollByte_Enum_spare12  uper.Enumerated = 52
	PollByte_Enum_spare11  uper.Enumerated = 53
	PollByte_Enum_spare10  uper.Enumerated = 54
	PollByte_Enum_spare9   uper.Enumerated = 55
	PollByte_Enum_spare8   uper.Enumerated = 56
	PollByte_Enum_spare7   uper.Enumerated = 57
	PollByte_Enum_spare6   uper.Enumerated = 58
	PollByte_Enum_spare5   uper.Enumerated = 59
	PollByte_Enum_spare4   uper.Enumerated = 60
	PollByte_Enum_spare3   uper.Enumerated = 61
	PollByte_Enum_spare2   uper.Enumerated = 62
	PollByte_Enum_spare1   uper.Enumerated = 63
)

type PollByte struct {
	Value uper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *PollByte) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode PollByte", err)
	}
	return nil
}

func (ie *PollByte) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode PollByte", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
