package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_nothing uint64 = iota
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms20_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms40_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms80_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms160_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms320_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms640_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms1280_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms2560_r17
	MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms5120_r17
)

type MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17 struct {
	Choice     uint64
	ms20_r17   int64 `lb:0,ub:19,madatory`
	ms40_r17   int64 `lb:0,ub:39,madatory`
	ms80_r17   int64 `lb:0,ub:79,madatory`
	ms160_r17  int64 `lb:0,ub:159,madatory`
	ms320_r17  int64 `lb:0,ub:319,madatory`
	ms640_r17  int64 `lb:0,ub:639,madatory`
	ms1280_r17 int64 `lb:0,ub:1279,madatory`
	ms2560_r17 int64 `lb:0,ub:2559,madatory`
	ms5120_r17 int64 `lb:0,ub:5119,madatory`
}

func (ie *MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms20_r17:
		if err = w.WriteInteger(int64(ie.ms20_r17), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode ms20_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms40_r17:
		if err = w.WriteInteger(int64(ie.ms40_r17), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode ms40_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms80_r17:
		if err = w.WriteInteger(int64(ie.ms80_r17), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode ms80_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms160_r17:
		if err = w.WriteInteger(int64(ie.ms160_r17), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode ms160_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms320_r17:
		if err = w.WriteInteger(int64(ie.ms320_r17), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode ms320_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms640_r17:
		if err = w.WriteInteger(int64(ie.ms640_r17), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode ms640_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms1280_r17:
		if err = w.WriteInteger(int64(ie.ms1280_r17), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode ms1280_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms2560_r17:
		if err = w.WriteInteger(int64(ie.ms2560_r17), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode ms2560_r17", err)
		}
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms5120_r17:
		if err = w.WriteInteger(int64(ie.ms5120_r17), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode ms5120_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms20_r17:
		var tmp_int_ms20_r17 int64
		if tmp_int_ms20_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode ms20_r17", err)
		}
		ie.ms20_r17 = tmp_int_ms20_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms40_r17:
		var tmp_int_ms40_r17 int64
		if tmp_int_ms40_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode ms40_r17", err)
		}
		ie.ms40_r17 = tmp_int_ms40_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms80_r17:
		var tmp_int_ms80_r17 int64
		if tmp_int_ms80_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode ms80_r17", err)
		}
		ie.ms80_r17 = tmp_int_ms80_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms160_r17:
		var tmp_int_ms160_r17 int64
		if tmp_int_ms160_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode ms160_r17", err)
		}
		ie.ms160_r17 = tmp_int_ms160_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms320_r17:
		var tmp_int_ms320_r17 int64
		if tmp_int_ms320_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode ms320_r17", err)
		}
		ie.ms320_r17 = tmp_int_ms320_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms640_r17:
		var tmp_int_ms640_r17 int64
		if tmp_int_ms640_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode ms640_r17", err)
		}
		ie.ms640_r17 = tmp_int_ms640_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms1280_r17:
		var tmp_int_ms1280_r17 int64
		if tmp_int_ms1280_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode ms1280_r17", err)
		}
		ie.ms1280_r17 = tmp_int_ms1280_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms2560_r17:
		var tmp_int_ms2560_r17 int64
		if tmp_int_ms2560_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode ms2560_r17", err)
		}
		ie.ms2560_r17 = tmp_int_ms2560_r17
	case MUSIM_GapInfo_r17_musim_GapRepetitionAndOffset_r17_Choice_ms5120_r17:
		var tmp_int_ms5120_r17 int64
		if tmp_int_ms5120_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode ms5120_r17", err)
		}
		ie.ms5120_r17 = tmp_int_ms5120_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
