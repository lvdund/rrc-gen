package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte32    uper.Enumerated = 0
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte100   uper.Enumerated = 1
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte200   uper.Enumerated = 2
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte400   uper.Enumerated = 3
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte600   uper.Enumerated = 4
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte800   uper.Enumerated = 5
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte1000  uper.Enumerated = 6
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte2000  uper.Enumerated = 7
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte4000  uper.Enumerated = 8
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte8000  uper.Enumerated = 9
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte9000  uper.Enumerated = 10
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte10000 uper.Enumerated = 11
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte12000 uper.Enumerated = 12
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte24000 uper.Enumerated = 13
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte48000 uper.Enumerated = 14
	SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17_Enum_byte96000 uper.Enumerated = 15
)

type SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
