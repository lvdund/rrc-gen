package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSIB_TypeAndInfo_r16_Choice_nothing uint64 = iota
	PosSIB_TypeAndInfo_r16_Choice_posSib1_1_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_2_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_3_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_4_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_5_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_6_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_7_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_8_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_1_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_2_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_3_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_4_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_5_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_6_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_7_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_8_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_9_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_10_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_11_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_12_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_13_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_14_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_15_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_16_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_17_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_18_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_19_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_20_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_21_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_22_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib2_23_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib3_1_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib4_1_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib5_1_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib6_1_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib6_2_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib6_3_r16
	PosSIB_TypeAndInfo_r16_Choice_posSib1_9_v1700
	PosSIB_TypeAndInfo_r16_Choice_posSib1_10_v1700
	PosSIB_TypeAndInfo_r16_Choice_posSib2_24_v1700
	PosSIB_TypeAndInfo_r16_Choice_posSib2_25_v1700
	PosSIB_TypeAndInfo_r16_Choice_posSib6_4_v1700
	PosSIB_TypeAndInfo_r16_Choice_posSib6_5_v1700
	PosSIB_TypeAndInfo_r16_Choice_posSib6_6_v1700
)

type PosSIB_TypeAndInfo_r16 struct {
	Choice           uint64
	posSib1_1_r16    *SIBpos_r16
	posSib1_2_r16    *SIBpos_r16
	posSib1_3_r16    *SIBpos_r16
	posSib1_4_r16    *SIBpos_r16
	posSib1_5_r16    *SIBpos_r16
	posSib1_6_r16    *SIBpos_r16
	posSib1_7_r16    *SIBpos_r16
	posSib1_8_r16    *SIBpos_r16
	posSib2_1_r16    *SIBpos_r16
	posSib2_2_r16    *SIBpos_r16
	posSib2_3_r16    *SIBpos_r16
	posSib2_4_r16    *SIBpos_r16
	posSib2_5_r16    *SIBpos_r16
	posSib2_6_r16    *SIBpos_r16
	posSib2_7_r16    *SIBpos_r16
	posSib2_8_r16    *SIBpos_r16
	posSib2_9_r16    *SIBpos_r16
	posSib2_10_r16   *SIBpos_r16
	posSib2_11_r16   *SIBpos_r16
	posSib2_12_r16   *SIBpos_r16
	posSib2_13_r16   *SIBpos_r16
	posSib2_14_r16   *SIBpos_r16
	posSib2_15_r16   *SIBpos_r16
	posSib2_16_r16   *SIBpos_r16
	posSib2_17_r16   *SIBpos_r16
	posSib2_18_r16   *SIBpos_r16
	posSib2_19_r16   *SIBpos_r16
	posSib2_20_r16   *SIBpos_r16
	posSib2_21_r16   *SIBpos_r16
	posSib2_22_r16   *SIBpos_r16
	posSib2_23_r16   *SIBpos_r16
	posSib3_1_r16    *SIBpos_r16
	posSib4_1_r16    *SIBpos_r16
	posSib5_1_r16    *SIBpos_r16
	posSib6_1_r16    *SIBpos_r16
	posSib6_2_r16    *SIBpos_r16
	posSib6_3_r16    *SIBpos_r16
	posSib1_9_v1700  *SIBpos_r16
	posSib1_10_v1700 *SIBpos_r16
	posSib2_24_v1700 *SIBpos_r16
	posSib2_25_v1700 *SIBpos_r16
	posSib6_4_v1700  *SIBpos_r16
	posSib6_5_v1700  *SIBpos_r16
	posSib6_6_v1700  *SIBpos_r16
}

func (ie *PosSIB_TypeAndInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 44, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_1_r16:
		if err = ie.posSib1_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_2_r16:
		if err = ie.posSib1_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_3_r16:
		if err = ie.posSib1_3_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_4_r16:
		if err = ie.posSib1_4_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_5_r16:
		if err = ie.posSib1_5_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_6_r16:
		if err = ie.posSib1_6_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_7_r16:
		if err = ie.posSib1_7_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_8_r16:
		if err = ie.posSib1_8_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_1_r16:
		if err = ie.posSib2_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_2_r16:
		if err = ie.posSib2_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_3_r16:
		if err = ie.posSib2_3_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_4_r16:
		if err = ie.posSib2_4_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_5_r16:
		if err = ie.posSib2_5_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_6_r16:
		if err = ie.posSib2_6_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_7_r16:
		if err = ie.posSib2_7_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_8_r16:
		if err = ie.posSib2_8_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_9_r16:
		if err = ie.posSib2_9_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_9_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_10_r16:
		if err = ie.posSib2_10_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_10_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_11_r16:
		if err = ie.posSib2_11_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_11_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_12_r16:
		if err = ie.posSib2_12_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_12_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_13_r16:
		if err = ie.posSib2_13_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_13_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_14_r16:
		if err = ie.posSib2_14_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_14_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_15_r16:
		if err = ie.posSib2_15_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_15_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_16_r16:
		if err = ie.posSib2_16_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_16_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_17_r16:
		if err = ie.posSib2_17_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_17_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_18_r16:
		if err = ie.posSib2_18_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_18_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_19_r16:
		if err = ie.posSib2_19_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_19_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_20_r16:
		if err = ie.posSib2_20_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_20_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_21_r16:
		if err = ie.posSib2_21_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_21_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_22_r16:
		if err = ie.posSib2_22_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_22_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_23_r16:
		if err = ie.posSib2_23_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_23_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib3_1_r16:
		if err = ie.posSib3_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib3_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib4_1_r16:
		if err = ie.posSib4_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib4_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib5_1_r16:
		if err = ie.posSib5_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib5_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_1_r16:
		if err = ie.posSib6_1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib6_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_2_r16:
		if err = ie.posSib6_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib6_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_3_r16:
		if err = ie.posSib6_3_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib6_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_9_v1700:
		if err = ie.posSib1_9_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_9_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_10_v1700:
		if err = ie.posSib1_10_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib1_10_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_24_v1700:
		if err = ie.posSib2_24_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_24_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_25_v1700:
		if err = ie.posSib2_25_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib2_25_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_4_v1700:
		if err = ie.posSib6_4_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib6_4_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_5_v1700:
		if err = ie.posSib6_5_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib6_5_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_6_v1700:
		if err = ie.posSib6_6_v1700.Encode(w); err != nil {
			err = utils.WrapError("Encode posSib6_6_v1700", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PosSIB_TypeAndInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(44, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_1_r16:
		ie.posSib1_1_r16 = new(SIBpos_r16)
		if err = ie.posSib1_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_2_r16:
		ie.posSib1_2_r16 = new(SIBpos_r16)
		if err = ie.posSib1_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_3_r16:
		ie.posSib1_3_r16 = new(SIBpos_r16)
		if err = ie.posSib1_3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_4_r16:
		ie.posSib1_4_r16 = new(SIBpos_r16)
		if err = ie.posSib1_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_5_r16:
		ie.posSib1_5_r16 = new(SIBpos_r16)
		if err = ie.posSib1_5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_6_r16:
		ie.posSib1_6_r16 = new(SIBpos_r16)
		if err = ie.posSib1_6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_7_r16:
		ie.posSib1_7_r16 = new(SIBpos_r16)
		if err = ie.posSib1_7_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_8_r16:
		ie.posSib1_8_r16 = new(SIBpos_r16)
		if err = ie.posSib1_8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_1_r16:
		ie.posSib2_1_r16 = new(SIBpos_r16)
		if err = ie.posSib2_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_2_r16:
		ie.posSib2_2_r16 = new(SIBpos_r16)
		if err = ie.posSib2_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_3_r16:
		ie.posSib2_3_r16 = new(SIBpos_r16)
		if err = ie.posSib2_3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_4_r16:
		ie.posSib2_4_r16 = new(SIBpos_r16)
		if err = ie.posSib2_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_4_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_5_r16:
		ie.posSib2_5_r16 = new(SIBpos_r16)
		if err = ie.posSib2_5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_5_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_6_r16:
		ie.posSib2_6_r16 = new(SIBpos_r16)
		if err = ie.posSib2_6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_6_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_7_r16:
		ie.posSib2_7_r16 = new(SIBpos_r16)
		if err = ie.posSib2_7_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_7_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_8_r16:
		ie.posSib2_8_r16 = new(SIBpos_r16)
		if err = ie.posSib2_8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_8_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_9_r16:
		ie.posSib2_9_r16 = new(SIBpos_r16)
		if err = ie.posSib2_9_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_9_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_10_r16:
		ie.posSib2_10_r16 = new(SIBpos_r16)
		if err = ie.posSib2_10_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_10_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_11_r16:
		ie.posSib2_11_r16 = new(SIBpos_r16)
		if err = ie.posSib2_11_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_11_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_12_r16:
		ie.posSib2_12_r16 = new(SIBpos_r16)
		if err = ie.posSib2_12_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_12_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_13_r16:
		ie.posSib2_13_r16 = new(SIBpos_r16)
		if err = ie.posSib2_13_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_13_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_14_r16:
		ie.posSib2_14_r16 = new(SIBpos_r16)
		if err = ie.posSib2_14_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_14_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_15_r16:
		ie.posSib2_15_r16 = new(SIBpos_r16)
		if err = ie.posSib2_15_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_15_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_16_r16:
		ie.posSib2_16_r16 = new(SIBpos_r16)
		if err = ie.posSib2_16_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_16_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_17_r16:
		ie.posSib2_17_r16 = new(SIBpos_r16)
		if err = ie.posSib2_17_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_17_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_18_r16:
		ie.posSib2_18_r16 = new(SIBpos_r16)
		if err = ie.posSib2_18_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_18_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_19_r16:
		ie.posSib2_19_r16 = new(SIBpos_r16)
		if err = ie.posSib2_19_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_19_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_20_r16:
		ie.posSib2_20_r16 = new(SIBpos_r16)
		if err = ie.posSib2_20_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_20_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_21_r16:
		ie.posSib2_21_r16 = new(SIBpos_r16)
		if err = ie.posSib2_21_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_21_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_22_r16:
		ie.posSib2_22_r16 = new(SIBpos_r16)
		if err = ie.posSib2_22_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_22_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_23_r16:
		ie.posSib2_23_r16 = new(SIBpos_r16)
		if err = ie.posSib2_23_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_23_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib3_1_r16:
		ie.posSib3_1_r16 = new(SIBpos_r16)
		if err = ie.posSib3_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib3_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib4_1_r16:
		ie.posSib4_1_r16 = new(SIBpos_r16)
		if err = ie.posSib4_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib4_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib5_1_r16:
		ie.posSib5_1_r16 = new(SIBpos_r16)
		if err = ie.posSib5_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib5_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_1_r16:
		ie.posSib6_1_r16 = new(SIBpos_r16)
		if err = ie.posSib6_1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib6_1_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_2_r16:
		ie.posSib6_2_r16 = new(SIBpos_r16)
		if err = ie.posSib6_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib6_2_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_3_r16:
		ie.posSib6_3_r16 = new(SIBpos_r16)
		if err = ie.posSib6_3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSib6_3_r16", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_9_v1700:
		ie.posSib1_9_v1700 = new(SIBpos_r16)
		if err = ie.posSib1_9_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_9_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib1_10_v1700:
		ie.posSib1_10_v1700 = new(SIBpos_r16)
		if err = ie.posSib1_10_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib1_10_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_24_v1700:
		ie.posSib2_24_v1700 = new(SIBpos_r16)
		if err = ie.posSib2_24_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_24_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib2_25_v1700:
		ie.posSib2_25_v1700 = new(SIBpos_r16)
		if err = ie.posSib2_25_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib2_25_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_4_v1700:
		ie.posSib6_4_v1700 = new(SIBpos_r16)
		if err = ie.posSib6_4_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib6_4_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_5_v1700:
		ie.posSib6_5_v1700 = new(SIBpos_r16)
		if err = ie.posSib6_5_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib6_5_v1700", err)
		}
	case PosSIB_TypeAndInfo_r16_Choice_posSib6_6_v1700:
		ie.posSib6_6_v1700 = new(SIBpos_r16)
		if err = ie.posSib6_6_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode posSib6_6_v1700", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
