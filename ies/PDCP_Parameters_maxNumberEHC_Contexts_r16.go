package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs2     uper.Enumerated = 0
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs4     uper.Enumerated = 1
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs8     uper.Enumerated = 2
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs16    uper.Enumerated = 3
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs32    uper.Enumerated = 4
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs64    uper.Enumerated = 5
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs128   uper.Enumerated = 6
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs256   uper.Enumerated = 7
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs512   uper.Enumerated = 8
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs1024  uper.Enumerated = 9
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs2048  uper.Enumerated = 10
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs4096  uper.Enumerated = 11
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs8192  uper.Enumerated = 12
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs16384 uper.Enumerated = 13
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs32768 uper.Enumerated = 14
	PDCP_Parameters_maxNumberEHC_Contexts_r16_Enum_cs65536 uper.Enumerated = 15
)

type PDCP_Parameters_maxNumberEHC_Contexts_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PDCP_Parameters_maxNumberEHC_Contexts_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PDCP_Parameters_maxNumberEHC_Contexts_r16", err)
	}
	return nil
}

func (ie *PDCP_Parameters_maxNumberEHC_Contexts_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PDCP_Parameters_maxNumberEHC_Contexts_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
