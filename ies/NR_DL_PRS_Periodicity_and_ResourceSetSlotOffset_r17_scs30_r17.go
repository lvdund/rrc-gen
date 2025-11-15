package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_nothing uint64 = iota
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n8_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n10_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n16_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n20_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n32_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n40_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n64_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n80_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n128_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n160_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n320_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n640_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n1280_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n2560_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n5120_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n10240_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n20480_r17
)

type NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17 struct {
	Choice     uint64
	n8_r17     int64 `lb:0,ub:7,madatory`
	n10_r17    int64 `lb:0,ub:9,madatory`
	n16_r17    int64 `lb:0,ub:15,madatory`
	n20_r17    int64 `lb:0,ub:19,madatory`
	n32_r17    int64 `lb:0,ub:31,madatory`
	n40_r17    int64 `lb:0,ub:39,madatory`
	n64_r17    int64 `lb:0,ub:63,madatory`
	n80_r17    int64 `lb:0,ub:79,madatory`
	n128_r17   int64 `lb:0,ub:127,madatory`
	n160_r17   int64 `lb:0,ub:159,madatory`
	n320_r17   int64 `lb:0,ub:319,madatory`
	n640_r17   int64 `lb:0,ub:639,madatory`
	n1280_r17  int64 `lb:0,ub:1279,madatory`
	n2560_r17  int64 `lb:0,ub:2559,madatory`
	n5120_r17  int64 `lb:0,ub:5119,madatory`
	n10240_r17 int64 `lb:0,ub:10239,madatory`
	n20480_r17 int64 `lb:0,ub:20479,madatory`
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n8_r17:
		if err = w.WriteInteger(int64(ie.n8_r17), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode n8_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n10_r17:
		if err = w.WriteInteger(int64(ie.n10_r17), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode n10_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n16_r17:
		if err = w.WriteInteger(int64(ie.n16_r17), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode n16_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n20_r17:
		if err = w.WriteInteger(int64(ie.n20_r17), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode n20_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n32_r17:
		if err = w.WriteInteger(int64(ie.n32_r17), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode n32_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n40_r17:
		if err = w.WriteInteger(int64(ie.n40_r17), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode n40_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n64_r17:
		if err = w.WriteInteger(int64(ie.n64_r17), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode n64_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n80_r17:
		if err = w.WriteInteger(int64(ie.n80_r17), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode n80_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n128_r17:
		if err = w.WriteInteger(int64(ie.n128_r17), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode n128_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n160_r17:
		if err = w.WriteInteger(int64(ie.n160_r17), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode n160_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n320_r17:
		if err = w.WriteInteger(int64(ie.n320_r17), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode n320_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n640_r17:
		if err = w.WriteInteger(int64(ie.n640_r17), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode n640_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n1280_r17:
		if err = w.WriteInteger(int64(ie.n1280_r17), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode n1280_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n2560_r17:
		if err = w.WriteInteger(int64(ie.n2560_r17), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode n2560_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n5120_r17:
		if err = w.WriteInteger(int64(ie.n5120_r17), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode n5120_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n10240_r17:
		if err = w.WriteInteger(int64(ie.n10240_r17), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode n10240_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n20480_r17:
		if err = w.WriteInteger(int64(ie.n20480_r17), &uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode n20480_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n8_r17:
		var tmp_int_n8_r17 int64
		if tmp_int_n8_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode n8_r17", err)
		}
		ie.n8_r17 = tmp_int_n8_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n10_r17:
		var tmp_int_n10_r17 int64
		if tmp_int_n10_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode n10_r17", err)
		}
		ie.n10_r17 = tmp_int_n10_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n16_r17:
		var tmp_int_n16_r17 int64
		if tmp_int_n16_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode n16_r17", err)
		}
		ie.n16_r17 = tmp_int_n16_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n20_r17:
		var tmp_int_n20_r17 int64
		if tmp_int_n20_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode n20_r17", err)
		}
		ie.n20_r17 = tmp_int_n20_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n32_r17:
		var tmp_int_n32_r17 int64
		if tmp_int_n32_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode n32_r17", err)
		}
		ie.n32_r17 = tmp_int_n32_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n40_r17:
		var tmp_int_n40_r17 int64
		if tmp_int_n40_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode n40_r17", err)
		}
		ie.n40_r17 = tmp_int_n40_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n64_r17:
		var tmp_int_n64_r17 int64
		if tmp_int_n64_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode n64_r17", err)
		}
		ie.n64_r17 = tmp_int_n64_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n80_r17:
		var tmp_int_n80_r17 int64
		if tmp_int_n80_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode n80_r17", err)
		}
		ie.n80_r17 = tmp_int_n80_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n128_r17:
		var tmp_int_n128_r17 int64
		if tmp_int_n128_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode n128_r17", err)
		}
		ie.n128_r17 = tmp_int_n128_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n160_r17:
		var tmp_int_n160_r17 int64
		if tmp_int_n160_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode n160_r17", err)
		}
		ie.n160_r17 = tmp_int_n160_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n320_r17:
		var tmp_int_n320_r17 int64
		if tmp_int_n320_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode n320_r17", err)
		}
		ie.n320_r17 = tmp_int_n320_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n640_r17:
		var tmp_int_n640_r17 int64
		if tmp_int_n640_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode n640_r17", err)
		}
		ie.n640_r17 = tmp_int_n640_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n1280_r17:
		var tmp_int_n1280_r17 int64
		if tmp_int_n1280_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode n1280_r17", err)
		}
		ie.n1280_r17 = tmp_int_n1280_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n2560_r17:
		var tmp_int_n2560_r17 int64
		if tmp_int_n2560_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode n2560_r17", err)
		}
		ie.n2560_r17 = tmp_int_n2560_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n5120_r17:
		var tmp_int_n5120_r17 int64
		if tmp_int_n5120_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode n5120_r17", err)
		}
		ie.n5120_r17 = tmp_int_n5120_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n10240_r17:
		var tmp_int_n10240_r17 int64
		if tmp_int_n10240_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode n10240_r17", err)
		}
		ie.n10240_r17 = tmp_int_n10240_r17
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17_Choice_n20480_r17:
		var tmp_int_n20480_r17 int64
		if tmp_int_n20480_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode n20480_r17", err)
		}
		ie.n20480_r17 = tmp_int_n20480_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
