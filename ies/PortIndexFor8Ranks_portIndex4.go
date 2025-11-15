package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndexFor8Ranks_portIndex4 struct {
	rank1_4 *PortIndex4  `optional`
	rank2_4 []PortIndex4 `lb:2,ub:2,optional`
	rank3_4 []PortIndex4 `lb:3,ub:3,optional`
	rank4_4 []PortIndex4 `lb:4,ub:4,optional`
}

func (ie *PortIndexFor8Ranks_portIndex4) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rank1_4 != nil, len(ie.rank2_4) > 0, len(ie.rank3_4) > 0, len(ie.rank4_4) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rank1_4 != nil {
		if err = ie.rank1_4.Encode(w); err != nil {
			return utils.WrapError("Encode rank1_4", err)
		}
	}
	if len(ie.rank2_4) > 0 {
		tmp_rank2_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.rank2_4 {
			tmp_rank2_4.Value = append(tmp_rank2_4.Value, &i)
		}
		if err = tmp_rank2_4.Encode(w); err != nil {
			return utils.WrapError("Encode rank2_4", err)
		}
	}
	if len(ie.rank3_4) > 0 {
		tmp_rank3_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.rank3_4 {
			tmp_rank3_4.Value = append(tmp_rank3_4.Value, &i)
		}
		if err = tmp_rank3_4.Encode(w); err != nil {
			return utils.WrapError("Encode rank3_4", err)
		}
	}
	if len(ie.rank4_4) > 0 {
		tmp_rank4_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.rank4_4 {
			tmp_rank4_4.Value = append(tmp_rank4_4.Value, &i)
		}
		if err = tmp_rank4_4.Encode(w); err != nil {
			return utils.WrapError("Encode rank4_4", err)
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks_portIndex4) Decode(r *uper.UperReader) error {
	var err error
	var rank1_4Present bool
	if rank1_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank2_4Present bool
	if rank2_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank3_4Present bool
	if rank3_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank4_4Present bool
	if rank4_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rank1_4Present {
		ie.rank1_4 = new(PortIndex4)
		if err = ie.rank1_4.Decode(r); err != nil {
			return utils.WrapError("Decode rank1_4", err)
		}
	}
	if rank2_4Present {
		tmp_rank2_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_rank2_4 := func() *PortIndex4 {
			return new(PortIndex4)
		}
		if err = tmp_rank2_4.Decode(r, fn_rank2_4); err != nil {
			return utils.WrapError("Decode rank2_4", err)
		}
		ie.rank2_4 = []PortIndex4{}
		for _, i := range tmp_rank2_4.Value {
			ie.rank2_4 = append(ie.rank2_4, *i)
		}
	}
	if rank3_4Present {
		tmp_rank3_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_rank3_4 := func() *PortIndex4 {
			return new(PortIndex4)
		}
		if err = tmp_rank3_4.Decode(r, fn_rank3_4); err != nil {
			return utils.WrapError("Decode rank3_4", err)
		}
		ie.rank3_4 = []PortIndex4{}
		for _, i := range tmp_rank3_4.Value {
			ie.rank3_4 = append(ie.rank3_4, *i)
		}
	}
	if rank4_4Present {
		tmp_rank4_4 := utils.NewSequence[*PortIndex4]([]*PortIndex4{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn_rank4_4 := func() *PortIndex4 {
			return new(PortIndex4)
		}
		if err = tmp_rank4_4.Decode(r, fn_rank4_4); err != nil {
			return utils.WrapError("Decode rank4_4", err)
		}
		ie.rank4_4 = []PortIndex4{}
		for _, i := range tmp_rank4_4.Value {
			ie.rank4_4 = append(ie.rank4_4, *i)
		}
	}
	return nil
}
