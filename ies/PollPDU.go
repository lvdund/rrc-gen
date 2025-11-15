package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PollPDU_Enum_p4       uper.Enumerated = 0
	PollPDU_Enum_p8       uper.Enumerated = 1
	PollPDU_Enum_p16      uper.Enumerated = 2
	PollPDU_Enum_p32      uper.Enumerated = 3
	PollPDU_Enum_p64      uper.Enumerated = 4
	PollPDU_Enum_p128     uper.Enumerated = 5
	PollPDU_Enum_p256     uper.Enumerated = 6
	PollPDU_Enum_p512     uper.Enumerated = 7
	PollPDU_Enum_p1024    uper.Enumerated = 8
	PollPDU_Enum_p2048    uper.Enumerated = 9
	PollPDU_Enum_p4096    uper.Enumerated = 10
	PollPDU_Enum_p6144    uper.Enumerated = 11
	PollPDU_Enum_p8192    uper.Enumerated = 12
	PollPDU_Enum_p12288   uper.Enumerated = 13
	PollPDU_Enum_p16384   uper.Enumerated = 14
	PollPDU_Enum_p20480   uper.Enumerated = 15
	PollPDU_Enum_p24576   uper.Enumerated = 16
	PollPDU_Enum_p28672   uper.Enumerated = 17
	PollPDU_Enum_p32768   uper.Enumerated = 18
	PollPDU_Enum_p40960   uper.Enumerated = 19
	PollPDU_Enum_p49152   uper.Enumerated = 20
	PollPDU_Enum_p57344   uper.Enumerated = 21
	PollPDU_Enum_p65536   uper.Enumerated = 22
	PollPDU_Enum_infinity uper.Enumerated = 23
	PollPDU_Enum_spare8   uper.Enumerated = 24
	PollPDU_Enum_spare7   uper.Enumerated = 25
	PollPDU_Enum_spare6   uper.Enumerated = 26
	PollPDU_Enum_spare5   uper.Enumerated = 27
	PollPDU_Enum_spare4   uper.Enumerated = 28
	PollPDU_Enum_spare3   uper.Enumerated = 29
	PollPDU_Enum_spare2   uper.Enumerated = 30
	PollPDU_Enum_spare1   uper.Enumerated = 31
)

type PollPDU struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *PollPDU) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode PollPDU", err)
	}
	return nil
}

func (ie *PollPDU) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode PollPDU", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
