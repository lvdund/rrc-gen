package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n5  uper.Enumerated = 0
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n6  uper.Enumerated = 1
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n7  uper.Enumerated = 2
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n8  uper.Enumerated = 3
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n9  uper.Enumerated = 4
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n10 uper.Enumerated = 5
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n12 uper.Enumerated = 6
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n14 uper.Enumerated = 7
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n16 uper.Enumerated = 8
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n18 uper.Enumerated = 9
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n20 uper.Enumerated = 10
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n22 uper.Enumerated = 11
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n24 uper.Enumerated = 12
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n26 uper.Enumerated = 13
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n28 uper.Enumerated = 14
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n30 uper.Enumerated = 15
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n32 uper.Enumerated = 16
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n34 uper.Enumerated = 17
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n36 uper.Enumerated = 18
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n38 uper.Enumerated = 19
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n40 uper.Enumerated = 20
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n42 uper.Enumerated = 21
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n44 uper.Enumerated = 22
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n46 uper.Enumerated = 23
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n48 uper.Enumerated = 24
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n50 uper.Enumerated = 25
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n52 uper.Enumerated = 26
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n54 uper.Enumerated = 27
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n56 uper.Enumerated = 28
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n58 uper.Enumerated = 29
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n60 uper.Enumerated = 30
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n62 uper.Enumerated = 31
	DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC_Enum_n64 uper.Enumerated = 32
)

type DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC struct {
	Value uper.Enumerated `lb:0,ub:32,madatory`
}

func (ie *DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	return nil
}

func (ie *DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode DummyA_maxNumberSimultaneousCSI_RS_ActBWP_AllCC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
