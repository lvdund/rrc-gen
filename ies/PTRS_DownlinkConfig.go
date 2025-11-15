package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_DownlinkConfig struct {
	frequencyDensity      []int64                                    `lb:2,ub:2,e_lb:0,e_ub:0,optional`
	timeDensity           []int64                                    `lb:3,ub:3,e_lb:0,e_ub:0,optional`
	epre_Ratio            *int64                                     `lb:0,ub:3,optional`
	resourceElementOffset *PTRS_DownlinkConfig_resourceElementOffset `optional`
	maxNrofPorts_r16      *PTRS_DownlinkConfig_maxNrofPorts_r16      `optional,ext-1`
}

func (ie *PTRS_DownlinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.maxNrofPorts_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.frequencyDensity) > 0, len(ie.timeDensity) > 0, ie.epre_Ratio != nil, ie.resourceElementOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.frequencyDensity) > 0 {
		tmp_frequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.frequencyDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_frequencyDensity.Value = append(tmp_frequencyDensity.Value, &tmp_ie)
		}
		if err = tmp_frequencyDensity.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyDensity", err)
		}
	}
	if len(ie.timeDensity) > 0 {
		tmp_timeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.timeDensity {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_timeDensity.Value = append(tmp_timeDensity.Value, &tmp_ie)
		}
		if err = tmp_timeDensity.Encode(w); err != nil {
			return utils.WrapError("Encode timeDensity", err)
		}
	}
	if ie.epre_Ratio != nil {
		if err = w.WriteInteger(*ie.epre_Ratio, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode epre_Ratio", err)
		}
	}
	if ie.resourceElementOffset != nil {
		if err = ie.resourceElementOffset.Encode(w); err != nil {
			return utils.WrapError("Encode resourceElementOffset", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.maxNrofPorts_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PTRS_DownlinkConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.maxNrofPorts_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxNrofPorts_r16 optional
			if ie.maxNrofPorts_r16 != nil {
				if err = ie.maxNrofPorts_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNrofPorts_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *PTRS_DownlinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyDensityPresent bool
	if frequencyDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var timeDensityPresent bool
	if timeDensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var epre_RatioPresent bool
	if epre_RatioPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceElementOffsetPresent bool
	if resourceElementOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyDensityPresent {
		tmp_frequencyDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_frequencyDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_frequencyDensity.Decode(r, fn_frequencyDensity); err != nil {
			return utils.WrapError("Decode frequencyDensity", err)
		}
		ie.frequencyDensity = []int64{}
		for _, i := range tmp_frequencyDensity.Value {
			ie.frequencyDensity = append(ie.frequencyDensity, int64(i.Value))
		}
	}
	if timeDensityPresent {
		tmp_timeDensity := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_timeDensity := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_timeDensity.Decode(r, fn_timeDensity); err != nil {
			return utils.WrapError("Decode timeDensity", err)
		}
		ie.timeDensity = []int64{}
		for _, i := range tmp_timeDensity.Value {
			ie.timeDensity = append(ie.timeDensity, int64(i.Value))
		}
	}
	if epre_RatioPresent {
		var tmp_int_epre_Ratio int64
		if tmp_int_epre_Ratio, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode epre_Ratio", err)
		}
		ie.epre_Ratio = &tmp_int_epre_Ratio
	}
	if resourceElementOffsetPresent {
		ie.resourceElementOffset = new(PTRS_DownlinkConfig_resourceElementOffset)
		if err = ie.resourceElementOffset.Decode(r); err != nil {
			return utils.WrapError("Decode resourceElementOffset", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			maxNrofPorts_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxNrofPorts_r16 optional
			if maxNrofPorts_r16Present {
				ie.maxNrofPorts_r16 = new(PTRS_DownlinkConfig_maxNrofPorts_r16)
				if err = ie.maxNrofPorts_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNrofPorts_r16", err)
				}
			}
		}
	}
	return nil
}
