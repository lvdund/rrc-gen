package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1710 struct {
	srs_AntennaSwitchingBeyond4RX_r17 *BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17 `lb:11,ub:11,optional`
}

func (ie *BandParameters_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_AntennaSwitchingBeyond4RX_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_AntennaSwitchingBeyond4RX_r17 != nil {
		if err = ie.srs_AntennaSwitchingBeyond4RX_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_AntennaSwitchingBeyond4RX_r17", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1710) Decode(r *uper.UperReader) error {
	var err error
	var srs_AntennaSwitchingBeyond4RX_r17Present bool
	if srs_AntennaSwitchingBeyond4RX_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_AntennaSwitchingBeyond4RX_r17Present {
		ie.srs_AntennaSwitchingBeyond4RX_r17 = new(BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17)
		if err = ie.srs_AntennaSwitchingBeyond4RX_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_AntennaSwitchingBeyond4RX_r17", err)
		}
	}
	return nil
}
