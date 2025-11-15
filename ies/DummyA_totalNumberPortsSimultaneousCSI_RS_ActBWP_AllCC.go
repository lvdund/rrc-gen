package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p8   uper.Enumerated = 0
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p12  uper.Enumerated = 1
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p16  uper.Enumerated = 2
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p24  uper.Enumerated = 3
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p32  uper.Enumerated = 4
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p40  uper.Enumerated = 5
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p48  uper.Enumerated = 6
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p56  uper.Enumerated = 7
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p64  uper.Enumerated = 8
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p72  uper.Enumerated = 9
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p80  uper.Enumerated = 10
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p88  uper.Enumerated = 11
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p96  uper.Enumerated = 12
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p104 uper.Enumerated = 13
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p112 uper.Enumerated = 14
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p120 uper.Enumerated = 15
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p128 uper.Enumerated = 16
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p136 uper.Enumerated = 17
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p144 uper.Enumerated = 18
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p152 uper.Enumerated = 19
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p160 uper.Enumerated = 20
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p168 uper.Enumerated = 21
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p176 uper.Enumerated = 22
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p184 uper.Enumerated = 23
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p192 uper.Enumerated = 24
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p200 uper.Enumerated = 25
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p208 uper.Enumerated = 26
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p216 uper.Enumerated = 27
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p224 uper.Enumerated = 28
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p232 uper.Enumerated = 29
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p240 uper.Enumerated = 30
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p248 uper.Enumerated = 31
	DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC_Enum_p256 uper.Enumerated = 32
)

type DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC struct {
	Value uper.Enumerated `lb:0,ub:32,madatory`
}

func (ie *DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}

func (ie *DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode DummyA_totalNumberPortsSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
