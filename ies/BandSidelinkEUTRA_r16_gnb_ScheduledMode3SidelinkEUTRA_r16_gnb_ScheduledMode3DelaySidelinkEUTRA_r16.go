package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0       uper.Enumerated = 0
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot25  uper.Enumerated = 1
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot5   uper.Enumerated = 2
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot625 uper.Enumerated = 3
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms0dot75  uper.Enumerated = 4
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1       uper.Enumerated = 5
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1dot25  uper.Enumerated = 6
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1dot5   uper.Enumerated = 7
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms1dot75  uper.Enumerated = 8
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms2       uper.Enumerated = 9
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms2dot5   uper.Enumerated = 10
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms3       uper.Enumerated = 11
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms4       uper.Enumerated = 12
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms5       uper.Enumerated = 13
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms6       uper.Enumerated = 14
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms8       uper.Enumerated = 15
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms10      uper.Enumerated = 16
	BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16_Enum_ms20      uper.Enumerated = 17
)

type BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16 struct {
	Value uper.Enumerated `lb:0,ub:17,madatory`
}

func (ie *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Encode BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16", err)
	}
	return nil
}

func (ie *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Decode BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
