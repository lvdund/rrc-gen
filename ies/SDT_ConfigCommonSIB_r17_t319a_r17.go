package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms100  uper.Enumerated = 0
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms200  uper.Enumerated = 1
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms300  uper.Enumerated = 2
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms400  uper.Enumerated = 3
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms600  uper.Enumerated = 4
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms1000 uper.Enumerated = 5
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms2000 uper.Enumerated = 6
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms3000 uper.Enumerated = 7
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_ms4000 uper.Enumerated = 8
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare7 uper.Enumerated = 9
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare6 uper.Enumerated = 10
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare5 uper.Enumerated = 11
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare4 uper.Enumerated = 12
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare3 uper.Enumerated = 13
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare2 uper.Enumerated = 14
	SDT_ConfigCommonSIB_r17_t319a_r17_Enum_spare1 uper.Enumerated = 15
)

type SDT_ConfigCommonSIB_r17_t319a_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SDT_ConfigCommonSIB_r17_t319a_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SDT_ConfigCommonSIB_r17_t319a_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17_t319a_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SDT_ConfigCommonSIB_r17_t319a_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
