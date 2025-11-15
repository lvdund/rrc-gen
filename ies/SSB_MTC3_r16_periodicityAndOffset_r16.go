package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_nothing uint64 = iota
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf5_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf10_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf20_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf40_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf80_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf160_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf320_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf640_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf1280_r16
)

type SSB_MTC3_r16_periodicityAndOffset_r16 struct {
	Choice     uint64
	sf5_r16    int64 `lb:0,ub:4,madatory`
	sf10_r16   int64 `lb:0,ub:9,madatory`
	sf20_r16   int64 `lb:0,ub:19,madatory`
	sf40_r16   int64 `lb:0,ub:39,madatory`
	sf80_r16   int64 `lb:0,ub:79,madatory`
	sf160_r16  int64 `lb:0,ub:159,madatory`
	sf320_r16  int64 `lb:0,ub:319,madatory`
	sf640_r16  int64 `lb:0,ub:639,madatory`
	sf1280_r16 int64 `lb:0,ub:1279,madatory`
}

func (ie *SSB_MTC3_r16_periodicityAndOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf5_r16:
		if err = w.WriteInteger(int64(ie.sf5_r16), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode sf5_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf10_r16:
		if err = w.WriteInteger(int64(ie.sf10_r16), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode sf10_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf20_r16:
		if err = w.WriteInteger(int64(ie.sf20_r16), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode sf20_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf40_r16:
		if err = w.WriteInteger(int64(ie.sf40_r16), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode sf40_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf80_r16:
		if err = w.WriteInteger(int64(ie.sf80_r16), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode sf80_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf160_r16:
		if err = w.WriteInteger(int64(ie.sf160_r16), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode sf160_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf320_r16:
		if err = w.WriteInteger(int64(ie.sf320_r16), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode sf320_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf640_r16:
		if err = w.WriteInteger(int64(ie.sf640_r16), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode sf640_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf1280_r16:
		if err = w.WriteInteger(int64(ie.sf1280_r16), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode sf1280_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SSB_MTC3_r16_periodicityAndOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf5_r16:
		var tmp_int_sf5_r16 int64
		if tmp_int_sf5_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode sf5_r16", err)
		}
		ie.sf5_r16 = tmp_int_sf5_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf10_r16:
		var tmp_int_sf10_r16 int64
		if tmp_int_sf10_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sf10_r16", err)
		}
		ie.sf10_r16 = tmp_int_sf10_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf20_r16:
		var tmp_int_sf20_r16 int64
		if tmp_int_sf20_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode sf20_r16", err)
		}
		ie.sf20_r16 = tmp_int_sf20_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf40_r16:
		var tmp_int_sf40_r16 int64
		if tmp_int_sf40_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode sf40_r16", err)
		}
		ie.sf40_r16 = tmp_int_sf40_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf80_r16:
		var tmp_int_sf80_r16 int64
		if tmp_int_sf80_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode sf80_r16", err)
		}
		ie.sf80_r16 = tmp_int_sf80_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf160_r16:
		var tmp_int_sf160_r16 int64
		if tmp_int_sf160_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode sf160_r16", err)
		}
		ie.sf160_r16 = tmp_int_sf160_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf320_r16:
		var tmp_int_sf320_r16 int64
		if tmp_int_sf320_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode sf320_r16", err)
		}
		ie.sf320_r16 = tmp_int_sf320_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf640_r16:
		var tmp_int_sf640_r16 int64
		if tmp_int_sf640_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode sf640_r16", err)
		}
		ie.sf640_r16 = tmp_int_sf640_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_sf1280_r16:
		var tmp_int_sf1280_r16 int64
		if tmp_int_sf1280_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode sf1280_r16", err)
		}
		ie.sf1280_r16 = tmp_int_sf1280_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
