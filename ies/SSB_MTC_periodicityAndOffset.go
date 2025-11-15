package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC_periodicityAndOffset_Choice_nothing uint64 = iota
	SSB_MTC_periodicityAndOffset_Choice_sf5
	SSB_MTC_periodicityAndOffset_Choice_sf10
	SSB_MTC_periodicityAndOffset_Choice_sf20
	SSB_MTC_periodicityAndOffset_Choice_sf40
	SSB_MTC_periodicityAndOffset_Choice_sf80
	SSB_MTC_periodicityAndOffset_Choice_sf160
)

type SSB_MTC_periodicityAndOffset struct {
	Choice uint64
	sf5    int64 `lb:0,ub:4,madatory`
	sf10   int64 `lb:0,ub:9,madatory`
	sf20   int64 `lb:0,ub:19,madatory`
	sf40   int64 `lb:0,ub:39,madatory`
	sf80   int64 `lb:0,ub:79,madatory`
	sf160  int64 `lb:0,ub:159,madatory`
}

func (ie *SSB_MTC_periodicityAndOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC_periodicityAndOffset_Choice_sf5:
		if err = w.WriteInteger(int64(ie.sf5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode sf5", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_sf10:
		if err = w.WriteInteger(int64(ie.sf10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode sf10", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_sf20:
		if err = w.WriteInteger(int64(ie.sf20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode sf20", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_sf40:
		if err = w.WriteInteger(int64(ie.sf40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode sf40", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_sf80:
		if err = w.WriteInteger(int64(ie.sf80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode sf80", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_sf160:
		if err = w.WriteInteger(int64(ie.sf160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode sf160", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SSB_MTC_periodicityAndOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC_periodicityAndOffset_Choice_sf5:
		var tmp_int_sf5 int64
		if tmp_int_sf5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode sf5", err)
		}
		ie.sf5 = tmp_int_sf5
	case SSB_MTC_periodicityAndOffset_Choice_sf10:
		var tmp_int_sf10 int64
		if tmp_int_sf10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sf10", err)
		}
		ie.sf10 = tmp_int_sf10
	case SSB_MTC_periodicityAndOffset_Choice_sf20:
		var tmp_int_sf20 int64
		if tmp_int_sf20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode sf20", err)
		}
		ie.sf20 = tmp_int_sf20
	case SSB_MTC_periodicityAndOffset_Choice_sf40:
		var tmp_int_sf40 int64
		if tmp_int_sf40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode sf40", err)
		}
		ie.sf40 = tmp_int_sf40
	case SSB_MTC_periodicityAndOffset_Choice_sf80:
		var tmp_int_sf80 int64
		if tmp_int_sf80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode sf80", err)
		}
		ie.sf80 = tmp_int_sf80
	case SSB_MTC_periodicityAndOffset_Choice_sf160:
		var tmp_int_sf160 int64
		if tmp_int_sf160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode sf160", err)
		}
		ie.sf160 = tmp_int_sf160
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
