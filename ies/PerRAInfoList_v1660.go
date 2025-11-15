package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRAInfoList_v1660 struct {
	Value []PerRACSI_RSInfo_v1660 `lb:1,ub:200,madatory`
}

func (ie *PerRAInfoList_v1660) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PerRACSI_RSInfo_v1660]([]*PerRACSI_RSInfo_v1660{}, uper.Constraint{Lb: 1, Ub: 200}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PerRAInfoList_v1660", err)
	}
	return nil
}

func (ie *PerRAInfoList_v1660) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PerRACSI_RSInfo_v1660]([]*PerRACSI_RSInfo_v1660{}, uper.Constraint{Lb: 1, Ub: 200}, false)
	fn := func() *PerRACSI_RSInfo_v1660 {
		return new(PerRACSI_RSInfo_v1660)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PerRAInfoList_v1660", err)
	}
	ie.Value = []PerRACSI_RSInfo_v1660{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
