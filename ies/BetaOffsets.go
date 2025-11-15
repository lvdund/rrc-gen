package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BetaOffsets struct {
	betaOffsetACK_Index1       *int64 `lb:0,ub:31,optional`
	betaOffsetACK_Index2       *int64 `lb:0,ub:31,optional`
	betaOffsetACK_Index3       *int64 `lb:0,ub:31,optional`
	betaOffsetCSI_Part1_Index1 *int64 `lb:0,ub:31,optional`
	betaOffsetCSI_Part1_Index2 *int64 `lb:0,ub:31,optional`
	betaOffsetCSI_Part2_Index1 *int64 `lb:0,ub:31,optional`
	betaOffsetCSI_Part2_Index2 *int64 `lb:0,ub:31,optional`
}

func (ie *BetaOffsets) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.betaOffsetACK_Index1 != nil, ie.betaOffsetACK_Index2 != nil, ie.betaOffsetACK_Index3 != nil, ie.betaOffsetCSI_Part1_Index1 != nil, ie.betaOffsetCSI_Part1_Index2 != nil, ie.betaOffsetCSI_Part2_Index1 != nil, ie.betaOffsetCSI_Part2_Index2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.betaOffsetACK_Index1 != nil {
		if err = w.WriteInteger(*ie.betaOffsetACK_Index1, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetACK_Index1", err)
		}
	}
	if ie.betaOffsetACK_Index2 != nil {
		if err = w.WriteInteger(*ie.betaOffsetACK_Index2, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetACK_Index2", err)
		}
	}
	if ie.betaOffsetACK_Index3 != nil {
		if err = w.WriteInteger(*ie.betaOffsetACK_Index3, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetACK_Index3", err)
		}
	}
	if ie.betaOffsetCSI_Part1_Index1 != nil {
		if err = w.WriteInteger(*ie.betaOffsetCSI_Part1_Index1, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetCSI_Part1_Index1", err)
		}
	}
	if ie.betaOffsetCSI_Part1_Index2 != nil {
		if err = w.WriteInteger(*ie.betaOffsetCSI_Part1_Index2, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetCSI_Part1_Index2", err)
		}
	}
	if ie.betaOffsetCSI_Part2_Index1 != nil {
		if err = w.WriteInteger(*ie.betaOffsetCSI_Part2_Index1, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetCSI_Part2_Index1", err)
		}
	}
	if ie.betaOffsetCSI_Part2_Index2 != nil {
		if err = w.WriteInteger(*ie.betaOffsetCSI_Part2_Index2, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode betaOffsetCSI_Part2_Index2", err)
		}
	}
	return nil
}

func (ie *BetaOffsets) Decode(r *uper.UperReader) error {
	var err error
	var betaOffsetACK_Index1Present bool
	if betaOffsetACK_Index1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var betaOffsetACK_Index2Present bool
	if betaOffsetACK_Index2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var betaOffsetACK_Index3Present bool
	if betaOffsetACK_Index3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var betaOffsetCSI_Part1_Index1Present bool
	if betaOffsetCSI_Part1_Index1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var betaOffsetCSI_Part1_Index2Present bool
	if betaOffsetCSI_Part1_Index2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var betaOffsetCSI_Part2_Index1Present bool
	if betaOffsetCSI_Part2_Index1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var betaOffsetCSI_Part2_Index2Present bool
	if betaOffsetCSI_Part2_Index2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if betaOffsetACK_Index1Present {
		var tmp_int_betaOffsetACK_Index1 int64
		if tmp_int_betaOffsetACK_Index1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetACK_Index1", err)
		}
		ie.betaOffsetACK_Index1 = &tmp_int_betaOffsetACK_Index1
	}
	if betaOffsetACK_Index2Present {
		var tmp_int_betaOffsetACK_Index2 int64
		if tmp_int_betaOffsetACK_Index2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetACK_Index2", err)
		}
		ie.betaOffsetACK_Index2 = &tmp_int_betaOffsetACK_Index2
	}
	if betaOffsetACK_Index3Present {
		var tmp_int_betaOffsetACK_Index3 int64
		if tmp_int_betaOffsetACK_Index3, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetACK_Index3", err)
		}
		ie.betaOffsetACK_Index3 = &tmp_int_betaOffsetACK_Index3
	}
	if betaOffsetCSI_Part1_Index1Present {
		var tmp_int_betaOffsetCSI_Part1_Index1 int64
		if tmp_int_betaOffsetCSI_Part1_Index1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetCSI_Part1_Index1", err)
		}
		ie.betaOffsetCSI_Part1_Index1 = &tmp_int_betaOffsetCSI_Part1_Index1
	}
	if betaOffsetCSI_Part1_Index2Present {
		var tmp_int_betaOffsetCSI_Part1_Index2 int64
		if tmp_int_betaOffsetCSI_Part1_Index2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetCSI_Part1_Index2", err)
		}
		ie.betaOffsetCSI_Part1_Index2 = &tmp_int_betaOffsetCSI_Part1_Index2
	}
	if betaOffsetCSI_Part2_Index1Present {
		var tmp_int_betaOffsetCSI_Part2_Index1 int64
		if tmp_int_betaOffsetCSI_Part2_Index1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetCSI_Part2_Index1", err)
		}
		ie.betaOffsetCSI_Part2_Index1 = &tmp_int_betaOffsetCSI_Part2_Index1
	}
	if betaOffsetCSI_Part2_Index2Present {
		var tmp_int_betaOffsetCSI_Part2_Index2 int64
		if tmp_int_betaOffsetCSI_Part2_Index2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode betaOffsetCSI_Part2_Index2", err)
		}
		ie.betaOffsetCSI_Part2_Index2 = &tmp_int_betaOffsetCSI_Part2_Index2
	}
	return nil
}
