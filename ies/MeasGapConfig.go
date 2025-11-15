package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasGapConfig struct {
	gapFR2                               *GapConfig                            `optional,setuprelease`
	gapFR1                               *GapConfig                            `optional,ext-1,setuprelease`
	gapUE                                *GapConfig                            `optional,ext-1,setuprelease`
	gapToAddModList_r17                  []GapConfig_r17                       `lb:1,ub:maxNrofGapId_r17,optional,ext-2`
	gapToReleaseList_r17                 []MeasGapId_r17                       `lb:1,ub:maxNrofGapId_r17,optional,ext-2`
	posMeasGapPreConfigToAddModList_r17  *PosMeasGapPreConfigToAddModList_r17  `optional,ext-2`
	posMeasGapPreConfigToReleaseList_r17 *PosMeasGapPreConfigToReleaseList_r17 `optional,ext-2`
}

func (ie *MeasGapConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.gapFR1 != nil || ie.gapUE != nil || len(ie.gapToAddModList_r17) > 0 || len(ie.gapToReleaseList_r17) > 0 || ie.posMeasGapPreConfigToAddModList_r17 != nil || ie.posMeasGapPreConfigToReleaseList_r17 != nil
	preambleBits := []bool{hasExtensions, ie.gapFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gapFR2 != nil {
		tmp_gapFR2 := utils.SetupRelease[*GapConfig]{
			Setup: ie.gapFR2,
		}
		if err = tmp_gapFR2.Encode(w); err != nil {
			return utils.WrapError("Encode gapFR2", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.gapFR1 != nil || ie.gapUE != nil, len(ie.gapToAddModList_r17) > 0 || len(ie.gapToReleaseList_r17) > 0 || ie.posMeasGapPreConfigToAddModList_r17 != nil || ie.posMeasGapPreConfigToReleaseList_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasGapConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.gapFR1 != nil, ie.gapUE != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode gapFR1 optional
			if ie.gapFR1 != nil {
				tmp_gapFR1 := utils.SetupRelease[*GapConfig]{
					Setup: ie.gapFR1,
				}
				if err = tmp_gapFR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gapFR1", err)
				}
			}
			// encode gapUE optional
			if ie.gapUE != nil {
				tmp_gapUE := utils.SetupRelease[*GapConfig]{
					Setup: ie.gapUE,
				}
				if err = tmp_gapUE.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gapUE", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.gapToAddModList_r17) > 0, len(ie.gapToReleaseList_r17) > 0, ie.posMeasGapPreConfigToAddModList_r17 != nil, ie.posMeasGapPreConfigToReleaseList_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode gapToAddModList_r17 optional
			if len(ie.gapToAddModList_r17) > 0 {
				tmp_gapToAddModList_r17 := utils.NewSequence[*GapConfig_r17]([]*GapConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				for _, i := range ie.gapToAddModList_r17 {
					tmp_gapToAddModList_r17.Value = append(tmp_gapToAddModList_r17.Value, &i)
				}
				if err = tmp_gapToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gapToAddModList_r17", err)
				}
			}
			// encode gapToReleaseList_r17 optional
			if len(ie.gapToReleaseList_r17) > 0 {
				tmp_gapToReleaseList_r17 := utils.NewSequence[*MeasGapId_r17]([]*MeasGapId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				for _, i := range ie.gapToReleaseList_r17 {
					tmp_gapToReleaseList_r17.Value = append(tmp_gapToReleaseList_r17.Value, &i)
				}
				if err = tmp_gapToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gapToReleaseList_r17", err)
				}
			}
			// encode posMeasGapPreConfigToAddModList_r17 optional
			if ie.posMeasGapPreConfigToAddModList_r17 != nil {
				if err = ie.posMeasGapPreConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode posMeasGapPreConfigToAddModList_r17", err)
				}
			}
			// encode posMeasGapPreConfigToReleaseList_r17 optional
			if ie.posMeasGapPreConfigToReleaseList_r17 != nil {
				if err = ie.posMeasGapPreConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode posMeasGapPreConfigToReleaseList_r17", err)
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

func (ie *MeasGapConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var gapFR2Present bool
	if gapFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gapFR2Present {
		tmp_gapFR2 := utils.SetupRelease[*GapConfig]{}
		if err = tmp_gapFR2.Decode(r); err != nil {
			return utils.WrapError("Decode gapFR2", err)
		}
		ie.gapFR2 = tmp_gapFR2.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			gapFR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gapUEPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode gapFR1 optional
			if gapFR1Present {
				tmp_gapFR1 := utils.SetupRelease[*GapConfig]{}
				if err = tmp_gapFR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode gapFR1", err)
				}
				ie.gapFR1 = tmp_gapFR1.Setup
			}
			// decode gapUE optional
			if gapUEPresent {
				tmp_gapUE := utils.SetupRelease[*GapConfig]{}
				if err = tmp_gapUE.Decode(extReader); err != nil {
					return utils.WrapError("Decode gapUE", err)
				}
				ie.gapUE = tmp_gapUE.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			gapToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gapToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			posMeasGapPreConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			posMeasGapPreConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode gapToAddModList_r17 optional
			if gapToAddModList_r17Present {
				tmp_gapToAddModList_r17 := utils.NewSequence[*GapConfig_r17]([]*GapConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				fn_gapToAddModList_r17 := func() *GapConfig_r17 {
					return new(GapConfig_r17)
				}
				if err = tmp_gapToAddModList_r17.Decode(extReader, fn_gapToAddModList_r17); err != nil {
					return utils.WrapError("Decode gapToAddModList_r17", err)
				}
				ie.gapToAddModList_r17 = []GapConfig_r17{}
				for _, i := range tmp_gapToAddModList_r17.Value {
					ie.gapToAddModList_r17 = append(ie.gapToAddModList_r17, *i)
				}
			}
			// decode gapToReleaseList_r17 optional
			if gapToReleaseList_r17Present {
				tmp_gapToReleaseList_r17 := utils.NewSequence[*MeasGapId_r17]([]*MeasGapId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false)
				fn_gapToReleaseList_r17 := func() *MeasGapId_r17 {
					return new(MeasGapId_r17)
				}
				if err = tmp_gapToReleaseList_r17.Decode(extReader, fn_gapToReleaseList_r17); err != nil {
					return utils.WrapError("Decode gapToReleaseList_r17", err)
				}
				ie.gapToReleaseList_r17 = []MeasGapId_r17{}
				for _, i := range tmp_gapToReleaseList_r17.Value {
					ie.gapToReleaseList_r17 = append(ie.gapToReleaseList_r17, *i)
				}
			}
			// decode posMeasGapPreConfigToAddModList_r17 optional
			if posMeasGapPreConfigToAddModList_r17Present {
				ie.posMeasGapPreConfigToAddModList_r17 = new(PosMeasGapPreConfigToAddModList_r17)
				if err = ie.posMeasGapPreConfigToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode posMeasGapPreConfigToAddModList_r17", err)
				}
			}
			// decode posMeasGapPreConfigToReleaseList_r17 optional
			if posMeasGapPreConfigToReleaseList_r17Present {
				ie.posMeasGapPreConfigToReleaseList_r17 = new(PosMeasGapPreConfigToReleaseList_r17)
				if err = ie.posMeasGapPreConfigToReleaseList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode posMeasGapPreConfigToReleaseList_r17", err)
				}
			}
		}
	}
	return nil
}
