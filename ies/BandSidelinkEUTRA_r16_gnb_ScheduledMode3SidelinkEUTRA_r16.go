package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16 struct {
	gnb_ScheduledMode3DelaySidelinkEUTRA_r16 BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16_gnb_ScheduledMode3DelaySidelinkEUTRA_r16 `madatory`
}

func (ie *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.gnb_ScheduledMode3DelaySidelinkEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode gnb_ScheduledMode3DelaySidelinkEUTRA_r16", err)
	}
	return nil
}

func (ie *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.gnb_ScheduledMode3DelaySidelinkEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode gnb_ScheduledMode3DelaySidelinkEUTRA_r16", err)
	}
	return nil
}
