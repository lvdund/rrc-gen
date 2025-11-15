package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkEUTRA_r16 struct {
	freqBandSidelinkEUTRA_r16           FreqBandIndicatorEUTRA                                     `madatory`
	gnb_ScheduledMode3SidelinkEUTRA_r16 *BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16 `optional`
	gnb_ScheduledMode4SidelinkEUTRA_r16 *BandSidelinkEUTRA_r16_gnb_ScheduledMode4SidelinkEUTRA_r16 `optional`
}

func (ie *BandSidelinkEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gnb_ScheduledMode3SidelinkEUTRA_r16 != nil, ie.gnb_ScheduledMode4SidelinkEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.freqBandSidelinkEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode freqBandSidelinkEUTRA_r16", err)
	}
	if ie.gnb_ScheduledMode3SidelinkEUTRA_r16 != nil {
		if err = ie.gnb_ScheduledMode3SidelinkEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnb_ScheduledMode3SidelinkEUTRA_r16", err)
		}
	}
	if ie.gnb_ScheduledMode4SidelinkEUTRA_r16 != nil {
		if err = ie.gnb_ScheduledMode4SidelinkEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnb_ScheduledMode4SidelinkEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelinkEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var gnb_ScheduledMode3SidelinkEUTRA_r16Present bool
	if gnb_ScheduledMode3SidelinkEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gnb_ScheduledMode4SidelinkEUTRA_r16Present bool
	if gnb_ScheduledMode4SidelinkEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.freqBandSidelinkEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode freqBandSidelinkEUTRA_r16", err)
	}
	if gnb_ScheduledMode3SidelinkEUTRA_r16Present {
		ie.gnb_ScheduledMode3SidelinkEUTRA_r16 = new(BandSidelinkEUTRA_r16_gnb_ScheduledMode3SidelinkEUTRA_r16)
		if err = ie.gnb_ScheduledMode3SidelinkEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnb_ScheduledMode3SidelinkEUTRA_r16", err)
		}
	}
	if gnb_ScheduledMode4SidelinkEUTRA_r16Present {
		ie.gnb_ScheduledMode4SidelinkEUTRA_r16 = new(BandSidelinkEUTRA_r16_gnb_ScheduledMode4SidelinkEUTRA_r16)
		if err = ie.gnb_ScheduledMode4SidelinkEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnb_ScheduledMode4SidelinkEUTRA_r16", err)
		}
	}
	return nil
}
