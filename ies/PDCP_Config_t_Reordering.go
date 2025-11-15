package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_t_Reordering_Enum_ms0     uper.Enumerated = 0
	PDCP_Config_t_Reordering_Enum_ms1     uper.Enumerated = 1
	PDCP_Config_t_Reordering_Enum_ms2     uper.Enumerated = 2
	PDCP_Config_t_Reordering_Enum_ms4     uper.Enumerated = 3
	PDCP_Config_t_Reordering_Enum_ms5     uper.Enumerated = 4
	PDCP_Config_t_Reordering_Enum_ms8     uper.Enumerated = 5
	PDCP_Config_t_Reordering_Enum_ms10    uper.Enumerated = 6
	PDCP_Config_t_Reordering_Enum_ms15    uper.Enumerated = 7
	PDCP_Config_t_Reordering_Enum_ms20    uper.Enumerated = 8
	PDCP_Config_t_Reordering_Enum_ms30    uper.Enumerated = 9
	PDCP_Config_t_Reordering_Enum_ms40    uper.Enumerated = 10
	PDCP_Config_t_Reordering_Enum_ms50    uper.Enumerated = 11
	PDCP_Config_t_Reordering_Enum_ms60    uper.Enumerated = 12
	PDCP_Config_t_Reordering_Enum_ms80    uper.Enumerated = 13
	PDCP_Config_t_Reordering_Enum_ms100   uper.Enumerated = 14
	PDCP_Config_t_Reordering_Enum_ms120   uper.Enumerated = 15
	PDCP_Config_t_Reordering_Enum_ms140   uper.Enumerated = 16
	PDCP_Config_t_Reordering_Enum_ms160   uper.Enumerated = 17
	PDCP_Config_t_Reordering_Enum_ms180   uper.Enumerated = 18
	PDCP_Config_t_Reordering_Enum_ms200   uper.Enumerated = 19
	PDCP_Config_t_Reordering_Enum_ms220   uper.Enumerated = 20
	PDCP_Config_t_Reordering_Enum_ms240   uper.Enumerated = 21
	PDCP_Config_t_Reordering_Enum_ms260   uper.Enumerated = 22
	PDCP_Config_t_Reordering_Enum_ms280   uper.Enumerated = 23
	PDCP_Config_t_Reordering_Enum_ms300   uper.Enumerated = 24
	PDCP_Config_t_Reordering_Enum_ms500   uper.Enumerated = 25
	PDCP_Config_t_Reordering_Enum_ms750   uper.Enumerated = 26
	PDCP_Config_t_Reordering_Enum_ms1000  uper.Enumerated = 27
	PDCP_Config_t_Reordering_Enum_ms1250  uper.Enumerated = 28
	PDCP_Config_t_Reordering_Enum_ms1500  uper.Enumerated = 29
	PDCP_Config_t_Reordering_Enum_ms1750  uper.Enumerated = 30
	PDCP_Config_t_Reordering_Enum_ms2000  uper.Enumerated = 31
	PDCP_Config_t_Reordering_Enum_ms2250  uper.Enumerated = 32
	PDCP_Config_t_Reordering_Enum_ms2500  uper.Enumerated = 33
	PDCP_Config_t_Reordering_Enum_ms2750  uper.Enumerated = 34
	PDCP_Config_t_Reordering_Enum_ms3000  uper.Enumerated = 35
	PDCP_Config_t_Reordering_Enum_spare28 uper.Enumerated = 36
	PDCP_Config_t_Reordering_Enum_spare27 uper.Enumerated = 37
	PDCP_Config_t_Reordering_Enum_spare26 uper.Enumerated = 38
	PDCP_Config_t_Reordering_Enum_spare25 uper.Enumerated = 39
	PDCP_Config_t_Reordering_Enum_spare24 uper.Enumerated = 40
	PDCP_Config_t_Reordering_Enum_spare23 uper.Enumerated = 41
	PDCP_Config_t_Reordering_Enum_spare22 uper.Enumerated = 42
	PDCP_Config_t_Reordering_Enum_spare21 uper.Enumerated = 43
	PDCP_Config_t_Reordering_Enum_spare20 uper.Enumerated = 44
	PDCP_Config_t_Reordering_Enum_spare19 uper.Enumerated = 45
	PDCP_Config_t_Reordering_Enum_spare18 uper.Enumerated = 46
	PDCP_Config_t_Reordering_Enum_spare17 uper.Enumerated = 47
	PDCP_Config_t_Reordering_Enum_spare16 uper.Enumerated = 48
	PDCP_Config_t_Reordering_Enum_spare15 uper.Enumerated = 49
	PDCP_Config_t_Reordering_Enum_spare14 uper.Enumerated = 50
	PDCP_Config_t_Reordering_Enum_spare13 uper.Enumerated = 51
	PDCP_Config_t_Reordering_Enum_spare12 uper.Enumerated = 52
	PDCP_Config_t_Reordering_Enum_spare11 uper.Enumerated = 53
	PDCP_Config_t_Reordering_Enum_spare10 uper.Enumerated = 54
	PDCP_Config_t_Reordering_Enum_spare09 uper.Enumerated = 55
	PDCP_Config_t_Reordering_Enum_spare08 uper.Enumerated = 56
	PDCP_Config_t_Reordering_Enum_spare07 uper.Enumerated = 57
	PDCP_Config_t_Reordering_Enum_spare06 uper.Enumerated = 58
	PDCP_Config_t_Reordering_Enum_spare05 uper.Enumerated = 59
	PDCP_Config_t_Reordering_Enum_spare04 uper.Enumerated = 60
	PDCP_Config_t_Reordering_Enum_spare03 uper.Enumerated = 61
	PDCP_Config_t_Reordering_Enum_spare02 uper.Enumerated = 62
	PDCP_Config_t_Reordering_Enum_spare01 uper.Enumerated = 63
)

type PDCP_Config_t_Reordering struct {
	Value uper.Enumerated `lb:0,ub:63,madatory`
}

func (ie *PDCP_Config_t_Reordering) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_t_Reordering", err)
	}
	return nil
}

func (ie *PDCP_Config_t_Reordering) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_t_Reordering", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
