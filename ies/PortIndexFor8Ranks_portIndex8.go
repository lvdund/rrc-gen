package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndexFor8Ranks_portIndex8 struct {
	rank1_8 *PortIndex8  `optional`
	rank2_8 []PortIndex8 `lb:2,ub:2,optional`
	rank3_8 []PortIndex8 `lb:3,ub:3,optional`
	rank4_8 []PortIndex8 `lb:4,ub:4,optional`
	rank5_8 []PortIndex8 `lb:5,ub:5,optional`
	rank6_8 []PortIndex8 `lb:6,ub:6,optional`
	rank7_8 []PortIndex8 `lb:7,ub:7,optional`
	rank8_8 []PortIndex8 `lb:8,ub:8,optional`
}

func (ie *PortIndexFor8Ranks_portIndex8) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rank1_8 != nil, len(ie.rank2_8) > 0, len(ie.rank3_8) > 0, len(ie.rank4_8) > 0, len(ie.rank5_8) > 0, len(ie.rank6_8) > 0, len(ie.rank7_8) > 0, len(ie.rank8_8) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rank1_8 != nil {
		if err = ie.rank1_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank1_8", err)
		}
	}
	if len(ie.rank2_8) > 0 {
		tmp_rank2_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.rank2_8 {
			tmp_rank2_8.Value = append(tmp_rank2_8.Value, &i)
		}
		if err = tmp_rank2_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank2_8", err)
		}
	}
	if len(ie.rank3_8) > 0 {
		tmp_rank3_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.rank3_8 {
			tmp_rank3_8.Value = append(tmp_rank3_8.Value, &i)
		}
		if err = tmp_rank3_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank3_8", err)
		}
	}
	if len(ie.rank4_8) > 0 {
		tmp_rank4_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.rank4_8 {
			tmp_rank4_8.Value = append(tmp_rank4_8.Value, &i)
		}
		if err = tmp_rank4_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank4_8", err)
		}
	}
	if len(ie.rank5_8) > 0 {
		tmp_rank5_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 5, Ub: 5}, false)
		for _, i := range ie.rank5_8 {
			tmp_rank5_8.Value = append(tmp_rank5_8.Value, &i)
		}
		if err = tmp_rank5_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank5_8", err)
		}
	}
	if len(ie.rank6_8) > 0 {
		tmp_rank6_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 6, Ub: 6}, false)
		for _, i := range ie.rank6_8 {
			tmp_rank6_8.Value = append(tmp_rank6_8.Value, &i)
		}
		if err = tmp_rank6_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank6_8", err)
		}
	}
	if len(ie.rank7_8) > 0 {
		tmp_rank7_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 7, Ub: 7}, false)
		for _, i := range ie.rank7_8 {
			tmp_rank7_8.Value = append(tmp_rank7_8.Value, &i)
		}
		if err = tmp_rank7_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank7_8", err)
		}
	}
	if len(ie.rank8_8) > 0 {
		tmp_rank8_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 8, Ub: 8}, false)
		for _, i := range ie.rank8_8 {
			tmp_rank8_8.Value = append(tmp_rank8_8.Value, &i)
		}
		if err = tmp_rank8_8.Encode(w); err != nil {
			return utils.WrapError("Encode rank8_8", err)
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks_portIndex8) Decode(r *uper.UperReader) error {
	var err error
	var rank1_8Present bool
	if rank1_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank2_8Present bool
	if rank2_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank3_8Present bool
	if rank3_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank4_8Present bool
	if rank4_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank5_8Present bool
	if rank5_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank6_8Present bool
	if rank6_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank7_8Present bool
	if rank7_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank8_8Present bool
	if rank8_8Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rank1_8Present {
		ie.rank1_8 = new(PortIndex8)
		if err = ie.rank1_8.Decode(r); err != nil {
			return utils.WrapError("Decode rank1_8", err)
		}
	}
	if rank2_8Present {
		tmp_rank2_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_rank2_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank2_8.Decode(r, fn_rank2_8); err != nil {
			return utils.WrapError("Decode rank2_8", err)
		}
		ie.rank2_8 = []PortIndex8{}
		for _, i := range tmp_rank2_8.Value {
			ie.rank2_8 = append(ie.rank2_8, *i)
		}
	}
	if rank3_8Present {
		tmp_rank3_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_rank3_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank3_8.Decode(r, fn_rank3_8); err != nil {
			return utils.WrapError("Decode rank3_8", err)
		}
		ie.rank3_8 = []PortIndex8{}
		for _, i := range tmp_rank3_8.Value {
			ie.rank3_8 = append(ie.rank3_8, *i)
		}
	}
	if rank4_8Present {
		tmp_rank4_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn_rank4_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank4_8.Decode(r, fn_rank4_8); err != nil {
			return utils.WrapError("Decode rank4_8", err)
		}
		ie.rank4_8 = []PortIndex8{}
		for _, i := range tmp_rank4_8.Value {
			ie.rank4_8 = append(ie.rank4_8, *i)
		}
	}
	if rank5_8Present {
		tmp_rank5_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 5, Ub: 5}, false)
		fn_rank5_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank5_8.Decode(r, fn_rank5_8); err != nil {
			return utils.WrapError("Decode rank5_8", err)
		}
		ie.rank5_8 = []PortIndex8{}
		for _, i := range tmp_rank5_8.Value {
			ie.rank5_8 = append(ie.rank5_8, *i)
		}
	}
	if rank6_8Present {
		tmp_rank6_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 6, Ub: 6}, false)
		fn_rank6_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank6_8.Decode(r, fn_rank6_8); err != nil {
			return utils.WrapError("Decode rank6_8", err)
		}
		ie.rank6_8 = []PortIndex8{}
		for _, i := range tmp_rank6_8.Value {
			ie.rank6_8 = append(ie.rank6_8, *i)
		}
	}
	if rank7_8Present {
		tmp_rank7_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 7, Ub: 7}, false)
		fn_rank7_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank7_8.Decode(r, fn_rank7_8); err != nil {
			return utils.WrapError("Decode rank7_8", err)
		}
		ie.rank7_8 = []PortIndex8{}
		for _, i := range tmp_rank7_8.Value {
			ie.rank7_8 = append(ie.rank7_8, *i)
		}
	}
	if rank8_8Present {
		tmp_rank8_8 := utils.NewSequence[*PortIndex8]([]*PortIndex8{}, uper.Constraint{Lb: 8, Ub: 8}, false)
		fn_rank8_8 := func() *PortIndex8 {
			return new(PortIndex8)
		}
		if err = tmp_rank8_8.Decode(r, fn_rank8_8); err != nil {
			return utils.WrapError("Decode rank8_8", err)
		}
		ie.rank8_8 = []PortIndex8{}
		for _, i := range tmp_rank8_8.Value {
			ie.rank8_8 = append(ie.rank8_8, *i)
		}
	}
	return nil
}
