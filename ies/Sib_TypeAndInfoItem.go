package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Sib_TypeAndInfoItem_Choice_nothing uint64 = iota
	Sib_TypeAndInfoItem_Choice_sib2
	Sib_TypeAndInfoItem_Choice_sib3
	Sib_TypeAndInfoItem_Choice_sib4
	Sib_TypeAndInfoItem_Choice_sib5
	Sib_TypeAndInfoItem_Choice_sib6
	Sib_TypeAndInfoItem_Choice_sib7
	Sib_TypeAndInfoItem_Choice_sib8
	Sib_TypeAndInfoItem_Choice_sib9
	Sib_TypeAndInfoItem_Choice_sib10_v1610
	Sib_TypeAndInfoItem_Choice_sib11_v1610
	Sib_TypeAndInfoItem_Choice_sib12_v1610
	Sib_TypeAndInfoItem_Choice_sib13_v1610
	Sib_TypeAndInfoItem_Choice_sib14_v1610
	Sib_TypeAndInfoItem_Choice_sib15_v1700
	Sib_TypeAndInfoItem_Choice_sib16_v1700
	Sib_TypeAndInfoItem_Choice_sib17_v1700
	Sib_TypeAndInfoItem_Choice_sib18_v1700
	Sib_TypeAndInfoItem_Choice_sib19_v1700
	Sib_TypeAndInfoItem_Choice_sib20_v1700
	Sib_TypeAndInfoItem_Choice_sib21_v1700
)

type Sib_TypeAndInfoItem struct {
	Choice      uint64
	sib2        *SIB2
	sib3        *SIB3
	sib4        *SIB4
	sib5        *SIB5
	sib6        *SIB6
	sib7        *SIB7
	sib8        *SIB8
	sib9        *SIB9
	sib10_v1610 *SIB10_r16
	sib11_v1610 *SIB11_r16
	sib12_v1610 *SIB12_r16
	sib13_v1610 *SIB13_r16
	sib14_v1610 *SIB14_r16
	sib15_v1700 *SIB15_r17
	sib16_v1700 *SIB16_r17
	sib17_v1700 *SIB17_r17
	sib18_v1700 *SIB18_r17
	sib19_v1700 *SIB19_r17
	sib20_v1700 *SIB20_r17
	sib21_v1700 *SIB21_r17
}

func (ie *Sib_TypeAndInfoItem) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Sib_TypeAndInfoItem_Choice_sib2:
		if err = ie.sib2.Encode(w); err != nil {
			err = utils.WrapError("Encode sib2", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib3:
		if err = ie.sib3.Encode(w); err != nil {
			err = utils.WrapError("Encode sib3", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib4:
		if err = ie.sib4.Encode(w); err != nil {
			err = utils.WrapError("Encode sib4", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib5:
		if err = ie.sib5.Encode(w); err != nil {
			err = utils.WrapError("Encode sib5", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib6:
		if err = ie.sib6.Encode(w); err != nil {
			err = utils.WrapError("Encode sib6", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib7:
		if err = ie.sib7.Encode(w); err != nil {
			err = utils.WrapError("Encode sib7", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib8:
		if err = ie.sib8.Encode(w); err != nil {
			err = utils.WrapError("Encode sib8", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib9:
		if err = ie.sib9.Encode(w); err != nil {
			err = utils.WrapError("Encode sib9", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib10_v1610:
		if err = ie.sib10_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode sib10_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib11_v1610:
		if err = ie.sib11_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode sib11_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib12_v1610:
		if err = ie.sib12_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode sib12_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib13_v1610:
		if err = ie.sib13_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode sib13_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib14_v1610:
		if err = ie.sib14_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode sib14_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib15_v1700:
		if err = ie.sib15_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib15_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib16_v1700:
		if err = ie.sib16_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib16_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib17_v1700:
		if err = ie.sib17_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib17_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib18_v1700:
		if err = ie.sib18_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib18_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib19_v1700:
		if err = ie.sib19_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib19_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib20_v1700:
		if err = ie.sib20_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib20_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib21_v1700:
		if err = ie.sib21_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode sib21_v1700", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *Sib_TypeAndInfoItem) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Sib_TypeAndInfoItem_Choice_sib2:
		ie.sib2 = new(SIB2)
		if err = ie.sib2.Decode(r); err != nil {
			return utils.WrapError("Decode sib2", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib3:
		ie.sib3 = new(SIB3)
		if err = ie.sib3.Decode(r); err != nil {
			return utils.WrapError("Decode sib3", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib4:
		ie.sib4 = new(SIB4)
		if err = ie.sib4.Decode(r); err != nil {
			return utils.WrapError("Decode sib4", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib5:
		ie.sib5 = new(SIB5)
		if err = ie.sib5.Decode(r); err != nil {
			return utils.WrapError("Decode sib5", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib6:
		ie.sib6 = new(SIB6)
		if err = ie.sib6.Decode(r); err != nil {
			return utils.WrapError("Decode sib6", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib7:
		ie.sib7 = new(SIB7)
		if err = ie.sib7.Decode(r); err != nil {
			return utils.WrapError("Decode sib7", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib8:
		ie.sib8 = new(SIB8)
		if err = ie.sib8.Decode(r); err != nil {
			return utils.WrapError("Decode sib8", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib9:
		ie.sib9 = new(SIB9)
		if err = ie.sib9.Decode(r); err != nil {
			return utils.WrapError("Decode sib9", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib10_v1610:
		ie.sib10_v1610 = new(SIB10_r16)
		if err = ie.sib10_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode sib10_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib11_v1610:
		ie.sib11_v1610 = new(SIB11_r16)
		if err = ie.sib11_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode sib11_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib12_v1610:
		ie.sib12_v1610 = new(SIB12_r16)
		if err = ie.sib12_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode sib12_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib13_v1610:
		ie.sib13_v1610 = new(SIB13_r16)
		if err = ie.sib13_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode sib13_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib14_v1610:
		ie.sib14_v1610 = new(SIB14_r16)
		if err = ie.sib14_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode sib14_v1610", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib15_v1700:
		ie.sib15_v1700 = new(SIB15_r17)
		if err = ie.sib15_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib15_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib16_v1700:
		ie.sib16_v1700 = new(SIB16_r17)
		if err = ie.sib16_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib16_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib17_v1700:
		ie.sib17_v1700 = new(SIB17_r17)
		if err = ie.sib17_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib17_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib18_v1700:
		ie.sib18_v1700 = new(SIB18_r17)
		if err = ie.sib18_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib18_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib19_v1700:
		ie.sib19_v1700 = new(SIB19_r17)
		if err = ie.sib19_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib19_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib20_v1700:
		ie.sib20_v1700 = new(SIB20_r17)
		if err = ie.sib20_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib20_v1700", err)
		}
	case Sib_TypeAndInfoItem_Choice_sib21_v1700:
		ie.sib21_v1700 = new(SIB21_r17)
		if err = ie.sib21_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sib21_v1700", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
