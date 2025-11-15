package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PerRAInfo_r16_Choice_nothing uint64 = iota
	PerRAInfo_r16_Choice_perRASSBInfoList_r16
	PerRAInfo_r16_Choice_perRACSI_RSInfoList_r16
)

type PerRAInfo_r16 struct {
	Choice                  uint64
	perRASSBInfoList_r16    *PerRASSBInfo_r16
	perRACSI_RSInfoList_r16 *PerRACSI_RSInfo_r16
}

func (ie *PerRAInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PerRAInfo_r16_Choice_perRASSBInfoList_r16:
		if err = ie.perRASSBInfoList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode perRASSBInfoList_r16", err)
		}
	case PerRAInfo_r16_Choice_perRACSI_RSInfoList_r16:
		if err = ie.perRACSI_RSInfoList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode perRACSI_RSInfoList_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PerRAInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PerRAInfo_r16_Choice_perRASSBInfoList_r16:
		ie.perRASSBInfoList_r16 = new(PerRASSBInfo_r16)
		if err = ie.perRASSBInfoList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode perRASSBInfoList_r16", err)
		}
	case PerRAInfo_r16_Choice_perRACSI_RSInfoList_r16:
		ie.perRACSI_RSInfoList_r16 = new(PerRACSI_RSInfo_r16)
		if err = ie.perRACSI_RSInfoList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode perRACSI_RSInfoList_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
