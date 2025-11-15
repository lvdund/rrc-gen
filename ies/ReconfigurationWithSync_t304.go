package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReconfigurationWithSync_t304_Enum_ms50    uper.Enumerated = 0
	ReconfigurationWithSync_t304_Enum_ms100   uper.Enumerated = 1
	ReconfigurationWithSync_t304_Enum_ms150   uper.Enumerated = 2
	ReconfigurationWithSync_t304_Enum_ms200   uper.Enumerated = 3
	ReconfigurationWithSync_t304_Enum_ms500   uper.Enumerated = 4
	ReconfigurationWithSync_t304_Enum_ms1000  uper.Enumerated = 5
	ReconfigurationWithSync_t304_Enum_ms2000  uper.Enumerated = 6
	ReconfigurationWithSync_t304_Enum_ms10000 uper.Enumerated = 7
)

type ReconfigurationWithSync_t304 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ReconfigurationWithSync_t304) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ReconfigurationWithSync_t304", err)
	}
	return nil
}

func (ie *ReconfigurationWithSync_t304) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ReconfigurationWithSync_t304", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
