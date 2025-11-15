package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_PollRetransmit_Enum_ms5       uper.Enumerated = 0
	T_PollRetransmit_Enum_ms10      uper.Enumerated = 1
	T_PollRetransmit_Enum_ms15      uper.Enumerated = 2
	T_PollRetransmit_Enum_ms20      uper.Enumerated = 3
	T_PollRetransmit_Enum_ms25      uper.Enumerated = 4
	T_PollRetransmit_Enum_ms30      uper.Enumerated = 5
	T_PollRetransmit_Enum_ms35      uper.Enumerated = 6
	T_PollRetransmit_Enum_ms40      uper.Enumerated = 7
	T_PollRetransmit_Enum_ms45      uper.Enumerated = 8
	T_PollRetransmit_Enum_ms50      uper.Enumerated = 9
	T_PollRetransmit_Enum_ms55      uper.Enumerated = 10
	T_PollRetransmit_Enum_ms60      uper.Enumerated = 11
	T_PollRetransmit_Enum_ms65      uper.Enumerated = 12
	T_PollRetransmit_Enum_ms70      uper.Enumerated = 13
	T_PollRetransmit_Enum_ms75      uper.Enumerated = 14
	T_PollRetransmit_Enum_ms80      uper.Enumerated = 15
	T_PollRetransmit_Enum_ms85      uper.Enumerated = 16
	T_PollRetransmit_Enum_ms90      uper.Enumerated = 17
	T_PollRetransmit_Enum_ms95      uper.Enumerated = 18
	T_PollRetransmit_Enum_ms100     uper.Enumerated = 19
	T_PollRetransmit_Enum_ms105     uper.Enumerated = 20
	T_PollRetransmit_Enum_ms110     uper.Enumerated = 21
	T_PollRetransmit_Enum_ms115     uper.Enumerated = 22
	T_PollRetransmit_Enum_ms120     uper.Enumerated = 23
	T_PollRetransmit_Enum_ms125     uper.Enumerated = 24
	T_PollRetransmit_Enum_ms130     uper.Enumerated = 25
	T_PollRetransmit_Enum_ms135     uper.Enumerated = 26
	T_PollRetransmit_Enum_ms140     uper.Enumerated = 27
	T_PollRetransmit_Enum_ms145     uper.Enumerated = 28
	T_PollRetransmit_Enum_ms150     uper.Enumerated = 29
	T_PollRetransmit_Enum_ms155     uper.Enumerated = 30
	T_PollRetransmit_Enum_ms160     uper.Enumerated = 31
	T_PollRetransmit_Enum_ms165     uper.Enumerated = 32
	T_PollRetransmit_Enum_ms170     uper.Enumerated = 33
	T_PollRetransmit_Enum_ms175     uper.Enumerated = 34
	T_PollRetransmit_Enum_ms180     uper.Enumerated = 35
	T_PollRetransmit_Enum_ms185     uper.Enumerated = 36
	T_PollRetransmit_Enum_ms190     uper.Enumerated = 37
	T_PollRetransmit_Enum_ms195     uper.Enumerated = 38
	T_PollRetransmit_Enum_ms200     uper.Enumerated = 39
	T_PollRetransmit_Enum_ms205     uper.Enumerated = 40
	T_PollRetransmit_Enum_ms210     uper.Enumerated = 41
	T_PollRetransmit_Enum_ms215     uper.Enumerated = 42
	T_PollRetransmit_Enum_ms220     uper.Enumerated = 43
	T_PollRetransmit_Enum_ms225     uper.Enumerated = 44
	T_PollRetransmit_Enum_ms230     uper.Enumerated = 45
	T_PollRetransmit_Enum_ms235     uper.Enumerated = 46
	T_PollRetransmit_Enum_ms240     uper.Enumerated = 47
	T_PollRetransmit_Enum_ms245     uper.Enumerated = 48
	T_PollRetransmit_Enum_ms250     uper.Enumerated = 49
	T_PollRetransmit_Enum_ms300     uper.Enumerated = 50
	T_PollRetransmit_Enum_ms350     uper.Enumerated = 51
	T_PollRetransmit_Enum_ms400     uper.Enumerated = 52
	T_PollRetransmit_Enum_ms450     uper.Enumerated = 53
	T_PollRetransmit_Enum_ms500     uper.Enumerated = 54
	T_PollRetransmit_Enum_ms800     uper.Enumerated = 55
	T_PollRetransmit_Enum_ms1000    uper.Enumerated = 56
	T_PollRetransmit_Enum_ms2000    uper.Enumerated = 57
	T_PollRetransmit_Enum_ms4000    uper.Enumerated = 58
	T_PollRetransmit_Enum_ms1_v1610 uper.Enumerated = 59
	T_PollRetransmit_Enum_ms2_v1610 uper.Enumerated = 60
	T_PollRetransmit_Enum_ms3_v1610 uper.Enumerated = 61
	T_PollRetransmit_Enum_ms4_v1610 uper.Enumerated = 62
	T_PollRetransmit_Enum_spare1    uper.Enumerated = 63
)

type T_PollRetransmit struct {
	Value uper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *T_PollRetransmit) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode T_PollRetransmit", err)
	}
	return nil
}

func (ie *T_PollRetransmit) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode T_PollRetransmit", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
