package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasGapSharingConfig struct {
	gapSharingFR2 *MeasGapSharingScheme `optional,setuprelease`
	gapSharingFR1 *MeasGapSharingScheme `optional,ext-1,setuprelease`
	gapSharingUE  *MeasGapSharingScheme `optional,ext-1,setuprelease`
}

func (ie *MeasGapSharingConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.gapSharingFR1 != nil || ie.gapSharingUE != nil
	preambleBits := []bool{hasExtensions, ie.gapSharingFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gapSharingFR2 != nil {
		tmp_gapSharingFR2 := utils.SetupRelease[*MeasGapSharingScheme]{
			Setup: ie.gapSharingFR2,
		}
		if err = tmp_gapSharingFR2.Encode(w); err != nil {
			return utils.WrapError("Encode gapSharingFR2", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.gapSharingFR1 != nil || ie.gapSharingUE != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasGapSharingConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.gapSharingFR1 != nil, ie.gapSharingUE != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode gapSharingFR1 optional
			if ie.gapSharingFR1 != nil {
				tmp_gapSharingFR1 := utils.SetupRelease[*MeasGapSharingScheme]{
					Setup: ie.gapSharingFR1,
				}
				if err = tmp_gapSharingFR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gapSharingFR1", err)
				}
			}
			// encode gapSharingUE optional
			if ie.gapSharingUE != nil {
				tmp_gapSharingUE := utils.SetupRelease[*MeasGapSharingScheme]{
					Setup: ie.gapSharingUE,
				}
				if err = tmp_gapSharingUE.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gapSharingUE", err)
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

func (ie *MeasGapSharingConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var gapSharingFR2Present bool
	if gapSharingFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gapSharingFR2Present {
		tmp_gapSharingFR2 := utils.SetupRelease[*MeasGapSharingScheme]{}
		if err = tmp_gapSharingFR2.Decode(r); err != nil {
			return utils.WrapError("Decode gapSharingFR2", err)
		}
		ie.gapSharingFR2 = tmp_gapSharingFR2.Setup
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

			gapSharingFR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gapSharingUEPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode gapSharingFR1 optional
			if gapSharingFR1Present {
				tmp_gapSharingFR1 := utils.SetupRelease[*MeasGapSharingScheme]{}
				if err = tmp_gapSharingFR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode gapSharingFR1", err)
				}
				ie.gapSharingFR1 = tmp_gapSharingFR1.Setup
			}
			// decode gapSharingUE optional
			if gapSharingUEPresent {
				tmp_gapSharingUE := utils.SetupRelease[*MeasGapSharingScheme]{}
				if err = tmp_gapSharingUE.Decode(extReader); err != nil {
					return utils.WrapError("Decode gapSharingUE", err)
				}
				ie.gapSharingUE = tmp_gapSharingUE.Setup
			}
		}
	}
	return nil
}
