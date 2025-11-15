package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1710 struct {
	ul_Semi_StaticChAccessDependentConfig_r17   *SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessDependentConfig_r17   `optional`
	ul_Semi_StaticChAccessIndependentConfig_r17 *SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessIndependentConfig_r17 `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_Semi_StaticChAccessDependentConfig_r17 != nil, ie.ul_Semi_StaticChAccessIndependentConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_Semi_StaticChAccessDependentConfig_r17 != nil {
		if err = ie.ul_Semi_StaticChAccessDependentConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_Semi_StaticChAccessDependentConfig_r17", err)
		}
	}
	if ie.ul_Semi_StaticChAccessIndependentConfig_r17 != nil {
		if err = ie.ul_Semi_StaticChAccessIndependentConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_Semi_StaticChAccessIndependentConfig_r17", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1710) Decode(r *uper.UperReader) error {
	var err error
	var ul_Semi_StaticChAccessDependentConfig_r17Present bool
	if ul_Semi_StaticChAccessDependentConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_Semi_StaticChAccessIndependentConfig_r17Present bool
	if ul_Semi_StaticChAccessIndependentConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_Semi_StaticChAccessDependentConfig_r17Present {
		ie.ul_Semi_StaticChAccessDependentConfig_r17 = new(SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessDependentConfig_r17)
		if err = ie.ul_Semi_StaticChAccessDependentConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_Semi_StaticChAccessDependentConfig_r17", err)
		}
	}
	if ul_Semi_StaticChAccessIndependentConfig_r17Present {
		ie.ul_Semi_StaticChAccessIndependentConfig_r17 = new(SharedSpectrumChAccessParamsPerBand_v1710_ul_Semi_StaticChAccessIndependentConfig_r17)
		if err = ie.ul_Semi_StaticChAccessIndependentConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_Semi_StaticChAccessIndependentConfig_r17", err)
		}
	}
	return nil
}
