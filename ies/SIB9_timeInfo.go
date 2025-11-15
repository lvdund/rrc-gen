package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB9_timeInfo struct {
	timeInfoUTC        int64           `lb:0,ub:549755813887,madatory`
	dayLightSavingTime *uper.BitString `lb:2,ub:2,optional`
	leapSeconds        *int64          `lb:-127,ub:128,optional`
	localTimeOffset    *int64          `lb:-63,ub:64,optional`
}

func (ie *SIB9_timeInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dayLightSavingTime != nil, ie.leapSeconds != nil, ie.localTimeOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.timeInfoUTC, &uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
		return utils.WrapError("WriteInteger timeInfoUTC", err)
	}
	if ie.dayLightSavingTime != nil {
		if err = w.WriteBitString(ie.dayLightSavingTime.Bytes, uint(ie.dayLightSavingTime.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode dayLightSavingTime", err)
		}
	}
	if ie.leapSeconds != nil {
		if err = w.WriteInteger(*ie.leapSeconds, &uper.Constraint{Lb: -127, Ub: 128}, false); err != nil {
			return utils.WrapError("Encode leapSeconds", err)
		}
	}
	if ie.localTimeOffset != nil {
		if err = w.WriteInteger(*ie.localTimeOffset, &uper.Constraint{Lb: -63, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode localTimeOffset", err)
		}
	}
	return nil
}

func (ie *SIB9_timeInfo) Decode(r *uper.UperReader) error {
	var err error
	var dayLightSavingTimePresent bool
	if dayLightSavingTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var leapSecondsPresent bool
	if leapSecondsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var localTimeOffsetPresent bool
	if localTimeOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_timeInfoUTC int64
	if tmp_int_timeInfoUTC, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
		return utils.WrapError("ReadInteger timeInfoUTC", err)
	}
	ie.timeInfoUTC = tmp_int_timeInfoUTC
	if dayLightSavingTimePresent {
		var tmp_bs_dayLightSavingTime []byte
		var n_dayLightSavingTime uint
		if tmp_bs_dayLightSavingTime, n_dayLightSavingTime, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode dayLightSavingTime", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_dayLightSavingTime,
			NumBits: uint64(n_dayLightSavingTime),
		}
		ie.dayLightSavingTime = &tmp_bitstring
	}
	if leapSecondsPresent {
		var tmp_int_leapSeconds int64
		if tmp_int_leapSeconds, err = r.ReadInteger(&uper.Constraint{Lb: -127, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode leapSeconds", err)
		}
		ie.leapSeconds = &tmp_int_leapSeconds
	}
	if localTimeOffsetPresent {
		var tmp_int_localTimeOffset int64
		if tmp_int_localTimeOffset, err = r.ReadInteger(&uper.Constraint{Lb: -63, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode localTimeOffset", err)
		}
		ie.localTimeOffset = &tmp_int_localTimeOffset
	}
	return nil
}
