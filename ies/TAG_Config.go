package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TAG_Config struct {
	tag_ToReleaseList []TAG_Id `lb:1,ub:maxNrofTAGs,optional`
	tag_ToAddModList  []TAG    `lb:1,ub:maxNrofTAGs,optional`
}

func (ie *TAG_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.tag_ToReleaseList) > 0, len(ie.tag_ToAddModList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.tag_ToReleaseList) > 0 {
		tmp_tag_ToReleaseList := utils.NewSequence[*TAG_Id]([]*TAG_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		for _, i := range ie.tag_ToReleaseList {
			tmp_tag_ToReleaseList.Value = append(tmp_tag_ToReleaseList.Value, &i)
		}
		if err = tmp_tag_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode tag_ToReleaseList", err)
		}
	}
	if len(ie.tag_ToAddModList) > 0 {
		tmp_tag_ToAddModList := utils.NewSequence[*TAG]([]*TAG{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		for _, i := range ie.tag_ToAddModList {
			tmp_tag_ToAddModList.Value = append(tmp_tag_ToAddModList.Value, &i)
		}
		if err = tmp_tag_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode tag_ToAddModList", err)
		}
	}
	return nil
}

func (ie *TAG_Config) Decode(r *uper.UperReader) error {
	var err error
	var tag_ToReleaseListPresent bool
	if tag_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tag_ToAddModListPresent bool
	if tag_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if tag_ToReleaseListPresent {
		tmp_tag_ToReleaseList := utils.NewSequence[*TAG_Id]([]*TAG_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		fn_tag_ToReleaseList := func() *TAG_Id {
			return new(TAG_Id)
		}
		if err = tmp_tag_ToReleaseList.Decode(r, fn_tag_ToReleaseList); err != nil {
			return utils.WrapError("Decode tag_ToReleaseList", err)
		}
		ie.tag_ToReleaseList = []TAG_Id{}
		for _, i := range tmp_tag_ToReleaseList.Value {
			ie.tag_ToReleaseList = append(ie.tag_ToReleaseList, *i)
		}
	}
	if tag_ToAddModListPresent {
		tmp_tag_ToAddModList := utils.NewSequence[*TAG]([]*TAG{}, uper.Constraint{Lb: 1, Ub: maxNrofTAGs}, false)
		fn_tag_ToAddModList := func() *TAG {
			return new(TAG)
		}
		if err = tmp_tag_ToAddModList.Decode(r, fn_tag_ToAddModList); err != nil {
			return utils.WrapError("Decode tag_ToAddModList", err)
		}
		ie.tag_ToAddModList = []TAG{}
		for _, i := range tmp_tag_ToAddModList.Value {
			ie.tag_ToAddModList = append(ie.tag_ToAddModList, *i)
		}
	}
	return nil
}
