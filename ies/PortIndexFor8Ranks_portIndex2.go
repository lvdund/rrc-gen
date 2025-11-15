package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndexFor8Ranks_portIndex2 struct {
	rank1_2 *PortIndex2  `optional`
	rank2_2 []PortIndex2 `lb:2,ub:2,optional`
}

func (ie *PortIndexFor8Ranks_portIndex2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rank1_2 != nil, len(ie.rank2_2) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rank1_2 != nil {
		if err = ie.rank1_2.Encode(w); err != nil {
			return utils.WrapError("Encode rank1_2", err)
		}
	}
	if len(ie.rank2_2) > 0 {
		tmp_rank2_2 := utils.NewSequence[*PortIndex2]([]*PortIndex2{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.rank2_2 {
			tmp_rank2_2.Value = append(tmp_rank2_2.Value, &i)
		}
		if err = tmp_rank2_2.Encode(w); err != nil {
			return utils.WrapError("Encode rank2_2", err)
		}
	}
	return nil
}

func (ie *PortIndexFor8Ranks_portIndex2) Decode(r *uper.UperReader) error {
	var err error
	var rank1_2Present bool
	if rank1_2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rank2_2Present bool
	if rank2_2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rank1_2Present {
		ie.rank1_2 = new(PortIndex2)
		if err = ie.rank1_2.Decode(r); err != nil {
			return utils.WrapError("Decode rank1_2", err)
		}
	}
	if rank2_2Present {
		tmp_rank2_2 := utils.NewSequence[*PortIndex2]([]*PortIndex2{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_rank2_2 := func() *PortIndex2 {
			return new(PortIndex2)
		}
		if err = tmp_rank2_2.Decode(r, fn_rank2_2); err != nil {
			return utils.WrapError("Decode rank2_2", err)
		}
		ie.rank2_2 = []PortIndex2{}
		for _, i := range tmp_rank2_2.Value {
			ie.rank2_2 = append(ie.rank2_2, *i)
		}
	}
	return nil
}
