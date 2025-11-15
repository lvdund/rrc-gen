package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandUTRA_FDD_r16_Enum_bandI      uper.Enumerated = 0
	SupportedBandUTRA_FDD_r16_Enum_bandII     uper.Enumerated = 1
	SupportedBandUTRA_FDD_r16_Enum_bandIII    uper.Enumerated = 2
	SupportedBandUTRA_FDD_r16_Enum_bandIV     uper.Enumerated = 3
	SupportedBandUTRA_FDD_r16_Enum_bandV      uper.Enumerated = 4
	SupportedBandUTRA_FDD_r16_Enum_bandVI     uper.Enumerated = 5
	SupportedBandUTRA_FDD_r16_Enum_bandVII    uper.Enumerated = 6
	SupportedBandUTRA_FDD_r16_Enum_bandVIII   uper.Enumerated = 7
	SupportedBandUTRA_FDD_r16_Enum_bandIX     uper.Enumerated = 8
	SupportedBandUTRA_FDD_r16_Enum_bandX      uper.Enumerated = 9
	SupportedBandUTRA_FDD_r16_Enum_bandXI     uper.Enumerated = 10
	SupportedBandUTRA_FDD_r16_Enum_bandXII    uper.Enumerated = 11
	SupportedBandUTRA_FDD_r16_Enum_bandXIII   uper.Enumerated = 12
	SupportedBandUTRA_FDD_r16_Enum_bandXIV    uper.Enumerated = 13
	SupportedBandUTRA_FDD_r16_Enum_bandXV     uper.Enumerated = 14
	SupportedBandUTRA_FDD_r16_Enum_bandXVI    uper.Enumerated = 15
	SupportedBandUTRA_FDD_r16_Enum_bandXVII   uper.Enumerated = 16
	SupportedBandUTRA_FDD_r16_Enum_bandXVIII  uper.Enumerated = 17
	SupportedBandUTRA_FDD_r16_Enum_bandXIX    uper.Enumerated = 18
	SupportedBandUTRA_FDD_r16_Enum_bandXX     uper.Enumerated = 19
	SupportedBandUTRA_FDD_r16_Enum_bandXXI    uper.Enumerated = 20
	SupportedBandUTRA_FDD_r16_Enum_bandXXII   uper.Enumerated = 21
	SupportedBandUTRA_FDD_r16_Enum_bandXXIII  uper.Enumerated = 22
	SupportedBandUTRA_FDD_r16_Enum_bandXXIV   uper.Enumerated = 23
	SupportedBandUTRA_FDD_r16_Enum_bandXXV    uper.Enumerated = 24
	SupportedBandUTRA_FDD_r16_Enum_bandXXVI   uper.Enumerated = 25
	SupportedBandUTRA_FDD_r16_Enum_bandXXVII  uper.Enumerated = 26
	SupportedBandUTRA_FDD_r16_Enum_bandXXVIII uper.Enumerated = 27
	SupportedBandUTRA_FDD_r16_Enum_bandXXIX   uper.Enumerated = 28
	SupportedBandUTRA_FDD_r16_Enum_bandXXX    uper.Enumerated = 29
	SupportedBandUTRA_FDD_r16_Enum_bandXXXI   uper.Enumerated = 30
	SupportedBandUTRA_FDD_r16_Enum_bandXXXII  uper.Enumerated = 31
)

type SupportedBandUTRA_FDD_r16 struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SupportedBandUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SupportedBandUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *SupportedBandUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SupportedBandUTRA_FDD_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
