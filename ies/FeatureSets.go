package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSets struct {
	featureSetsDownlink            []FeatureSetDownlink            `lb:1,ub:maxDownlinkFeatureSets,optional`
	featureSetsDownlinkPerCC       []FeatureSetDownlinkPerCC       `lb:1,ub:maxPerCC_FeatureSets,optional`
	featureSetsUplink              []FeatureSetUplink              `lb:1,ub:maxUplinkFeatureSets,optional`
	featureSetsUplinkPerCC         []FeatureSetUplinkPerCC         `lb:1,ub:maxPerCC_FeatureSets,optional`
	featureSetsDownlink_v1540      []FeatureSetDownlink_v1540      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-1`
	featureSetsUplink_v1540        []FeatureSetUplink_v1540        `lb:1,ub:maxUplinkFeatureSets,optional,ext-1`
	featureSetsUplinkPerCC_v1540   []FeatureSetUplinkPerCC_v1540   `lb:1,ub:maxPerCC_FeatureSets,optional,ext-1`
	featureSetsDownlink_v15a0      []FeatureSetDownlink_v15a0      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-2`
	featureSetsDownlink_v1610      []FeatureSetDownlink_v1610      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-3`
	featureSetsUplink_v1610        []FeatureSetUplink_v1610        `lb:1,ub:maxUplinkFeatureSets,optional,ext-3`
	featureSetDownlinkPerCC_v1620  []FeatureSetDownlinkPerCC_v1620 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-3`
	featureSetsUplink_v1630        []FeatureSetUplink_v1630        `lb:1,ub:maxUplinkFeatureSets,optional,ext-4`
	featureSetsUplink_v1640        []FeatureSetUplink_v1640        `lb:1,ub:maxUplinkFeatureSets,optional,ext-5`
	featureSetsDownlink_v1700      []FeatureSetDownlink_v1700      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-6`
	featureSetsDownlinkPerCC_v1700 []FeatureSetDownlinkPerCC_v1700 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-6`
	featureSetsUplink_v1710        []FeatureSetUplink_v1710        `lb:1,ub:maxUplinkFeatureSets,optional,ext-6`
	featureSetsUplinkPerCC_v1700   []FeatureSetUplinkPerCC_v1700   `lb:1,ub:maxPerCC_FeatureSets,optional,ext-6`
	featureSetsDownlink_v1720      []FeatureSetDownlink_v1720      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-7`
	featureSetsDownlinkPerCC_v1720 []FeatureSetDownlinkPerCC_v1720 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-7`
	featureSetsUplink_v1720        []FeatureSetUplink_v1720        `lb:1,ub:maxUplinkFeatureSets,optional,ext-7`
	featureSetsDownlink_v1730      []FeatureSetDownlink_v1730      `lb:1,ub:maxDownlinkFeatureSets,optional,ext-8`
	featureSetsDownlinkPerCC_v1730 []FeatureSetDownlinkPerCC_v1730 `lb:1,ub:maxPerCC_FeatureSets,optional,ext-8`
}

func (ie *FeatureSets) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.featureSetsDownlink_v1540) > 0 || len(ie.featureSetsUplink_v1540) > 0 || len(ie.featureSetsUplinkPerCC_v1540) > 0 || len(ie.featureSetsDownlink_v15a0) > 0 || len(ie.featureSetsDownlink_v1610) > 0 || len(ie.featureSetsUplink_v1610) > 0 || len(ie.featureSetDownlinkPerCC_v1620) > 0 || len(ie.featureSetsUplink_v1630) > 0 || len(ie.featureSetsUplink_v1640) > 0 || len(ie.featureSetsDownlink_v1700) > 0 || len(ie.featureSetsDownlinkPerCC_v1700) > 0 || len(ie.featureSetsUplink_v1710) > 0 || len(ie.featureSetsUplinkPerCC_v1700) > 0 || len(ie.featureSetsDownlink_v1720) > 0 || len(ie.featureSetsDownlinkPerCC_v1720) > 0 || len(ie.featureSetsUplink_v1720) > 0 || len(ie.featureSetsDownlink_v1730) > 0 || len(ie.featureSetsDownlinkPerCC_v1730) > 0
	preambleBits := []bool{hasExtensions, len(ie.featureSetsDownlink) > 0, len(ie.featureSetsDownlinkPerCC) > 0, len(ie.featureSetsUplink) > 0, len(ie.featureSetsUplinkPerCC) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.featureSetsDownlink) > 0 {
		tmp_featureSetsDownlink := utils.NewSequence[*FeatureSetDownlink]([]*FeatureSetDownlink{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
		for _, i := range ie.featureSetsDownlink {
			tmp_featureSetsDownlink.Value = append(tmp_featureSetsDownlink.Value, &i)
		}
		if err = tmp_featureSetsDownlink.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetsDownlink", err)
		}
	}
	if len(ie.featureSetsDownlinkPerCC) > 0 {
		tmp_featureSetsDownlinkPerCC := utils.NewSequence[*FeatureSetDownlinkPerCC]([]*FeatureSetDownlinkPerCC{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		for _, i := range ie.featureSetsDownlinkPerCC {
			tmp_featureSetsDownlinkPerCC.Value = append(tmp_featureSetsDownlinkPerCC.Value, &i)
		}
		if err = tmp_featureSetsDownlinkPerCC.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetsDownlinkPerCC", err)
		}
	}
	if len(ie.featureSetsUplink) > 0 {
		tmp_featureSetsUplink := utils.NewSequence[*FeatureSetUplink]([]*FeatureSetUplink{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
		for _, i := range ie.featureSetsUplink {
			tmp_featureSetsUplink.Value = append(tmp_featureSetsUplink.Value, &i)
		}
		if err = tmp_featureSetsUplink.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetsUplink", err)
		}
	}
	if len(ie.featureSetsUplinkPerCC) > 0 {
		tmp_featureSetsUplinkPerCC := utils.NewSequence[*FeatureSetUplinkPerCC]([]*FeatureSetUplinkPerCC{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		for _, i := range ie.featureSetsUplinkPerCC {
			tmp_featureSetsUplinkPerCC.Value = append(tmp_featureSetsUplinkPerCC.Value, &i)
		}
		if err = tmp_featureSetsUplinkPerCC.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetsUplinkPerCC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 8 bits for 8 extension groups
		extBitmap := []bool{len(ie.featureSetsDownlink_v1540) > 0 || len(ie.featureSetsUplink_v1540) > 0 || len(ie.featureSetsUplinkPerCC_v1540) > 0, len(ie.featureSetsDownlink_v15a0) > 0, len(ie.featureSetsDownlink_v1610) > 0 || len(ie.featureSetsUplink_v1610) > 0 || len(ie.featureSetDownlinkPerCC_v1620) > 0, len(ie.featureSetsUplink_v1630) > 0, len(ie.featureSetsUplink_v1640) > 0, len(ie.featureSetsDownlink_v1700) > 0 || len(ie.featureSetsDownlinkPerCC_v1700) > 0 || len(ie.featureSetsUplink_v1710) > 0 || len(ie.featureSetsUplinkPerCC_v1700) > 0, len(ie.featureSetsDownlink_v1720) > 0 || len(ie.featureSetsDownlinkPerCC_v1720) > 0 || len(ie.featureSetsUplink_v1720) > 0, len(ie.featureSetsDownlink_v1730) > 0 || len(ie.featureSetsDownlinkPerCC_v1730) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FeatureSets", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.featureSetsDownlink_v1540) > 0, len(ie.featureSetsUplink_v1540) > 0, len(ie.featureSetsUplinkPerCC_v1540) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsDownlink_v1540 optional
			if len(ie.featureSetsDownlink_v1540) > 0 {
				tmp_featureSetsDownlink_v1540 := utils.NewSequence[*FeatureSetDownlink_v1540]([]*FeatureSetDownlink_v1540{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.featureSetsDownlink_v1540 {
					tmp_featureSetsDownlink_v1540.Value = append(tmp_featureSetsDownlink_v1540.Value, &i)
				}
				if err = tmp_featureSetsDownlink_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlink_v1540", err)
				}
			}
			// encode featureSetsUplink_v1540 optional
			if len(ie.featureSetsUplink_v1540) > 0 {
				tmp_featureSetsUplink_v1540 := utils.NewSequence[*FeatureSetUplink_v1540]([]*FeatureSetUplink_v1540{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.featureSetsUplink_v1540 {
					tmp_featureSetsUplink_v1540.Value = append(tmp_featureSetsUplink_v1540.Value, &i)
				}
				if err = tmp_featureSetsUplink_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplink_v1540", err)
				}
			}
			// encode featureSetsUplinkPerCC_v1540 optional
			if len(ie.featureSetsUplinkPerCC_v1540) > 0 {
				tmp_featureSetsUplinkPerCC_v1540 := utils.NewSequence[*FeatureSetUplinkPerCC_v1540]([]*FeatureSetUplinkPerCC_v1540{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.featureSetsUplinkPerCC_v1540 {
					tmp_featureSetsUplinkPerCC_v1540.Value = append(tmp_featureSetsUplinkPerCC_v1540.Value, &i)
				}
				if err = tmp_featureSetsUplinkPerCC_v1540.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplinkPerCC_v1540", err)
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
			optionals_ext_2 := []bool{len(ie.featureSetsDownlink_v15a0) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsDownlink_v15a0 optional
			if len(ie.featureSetsDownlink_v15a0) > 0 {
				tmp_featureSetsDownlink_v15a0 := utils.NewSequence[*FeatureSetDownlink_v15a0]([]*FeatureSetDownlink_v15a0{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.featureSetsDownlink_v15a0 {
					tmp_featureSetsDownlink_v15a0.Value = append(tmp_featureSetsDownlink_v15a0.Value, &i)
				}
				if err = tmp_featureSetsDownlink_v15a0.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlink_v15a0", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{len(ie.featureSetsDownlink_v1610) > 0, len(ie.featureSetsUplink_v1610) > 0, len(ie.featureSetDownlinkPerCC_v1620) > 0}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsDownlink_v1610 optional
			if len(ie.featureSetsDownlink_v1610) > 0 {
				tmp_featureSetsDownlink_v1610 := utils.NewSequence[*FeatureSetDownlink_v1610]([]*FeatureSetDownlink_v1610{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.featureSetsDownlink_v1610 {
					tmp_featureSetsDownlink_v1610.Value = append(tmp_featureSetsDownlink_v1610.Value, &i)
				}
				if err = tmp_featureSetsDownlink_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlink_v1610", err)
				}
			}
			// encode featureSetsUplink_v1610 optional
			if len(ie.featureSetsUplink_v1610) > 0 {
				tmp_featureSetsUplink_v1610 := utils.NewSequence[*FeatureSetUplink_v1610]([]*FeatureSetUplink_v1610{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.featureSetsUplink_v1610 {
					tmp_featureSetsUplink_v1610.Value = append(tmp_featureSetsUplink_v1610.Value, &i)
				}
				if err = tmp_featureSetsUplink_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplink_v1610", err)
				}
			}
			// encode featureSetDownlinkPerCC_v1620 optional
			if len(ie.featureSetDownlinkPerCC_v1620) > 0 {
				tmp_featureSetDownlinkPerCC_v1620 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1620]([]*FeatureSetDownlinkPerCC_v1620{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.featureSetDownlinkPerCC_v1620 {
					tmp_featureSetDownlinkPerCC_v1620.Value = append(tmp_featureSetDownlinkPerCC_v1620.Value, &i)
				}
				if err = tmp_featureSetDownlinkPerCC_v1620.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetDownlinkPerCC_v1620", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{len(ie.featureSetsUplink_v1630) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsUplink_v1630 optional
			if len(ie.featureSetsUplink_v1630) > 0 {
				tmp_featureSetsUplink_v1630 := utils.NewSequence[*FeatureSetUplink_v1630]([]*FeatureSetUplink_v1630{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.featureSetsUplink_v1630 {
					tmp_featureSetsUplink_v1630.Value = append(tmp_featureSetsUplink_v1630.Value, &i)
				}
				if err = tmp_featureSetsUplink_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplink_v1630", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{len(ie.featureSetsUplink_v1640) > 0}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsUplink_v1640 optional
			if len(ie.featureSetsUplink_v1640) > 0 {
				tmp_featureSetsUplink_v1640 := utils.NewSequence[*FeatureSetUplink_v1640]([]*FeatureSetUplink_v1640{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.featureSetsUplink_v1640 {
					tmp_featureSetsUplink_v1640.Value = append(tmp_featureSetsUplink_v1640.Value, &i)
				}
				if err = tmp_featureSetsUplink_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplink_v1640", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{len(ie.featureSetsDownlink_v1700) > 0, len(ie.featureSetsDownlinkPerCC_v1700) > 0, len(ie.featureSetsUplink_v1710) > 0, len(ie.featureSetsUplinkPerCC_v1700) > 0}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsDownlink_v1700 optional
			if len(ie.featureSetsDownlink_v1700) > 0 {
				tmp_featureSetsDownlink_v1700 := utils.NewSequence[*FeatureSetDownlink_v1700]([]*FeatureSetDownlink_v1700{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.featureSetsDownlink_v1700 {
					tmp_featureSetsDownlink_v1700.Value = append(tmp_featureSetsDownlink_v1700.Value, &i)
				}
				if err = tmp_featureSetsDownlink_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlink_v1700", err)
				}
			}
			// encode featureSetsDownlinkPerCC_v1700 optional
			if len(ie.featureSetsDownlinkPerCC_v1700) > 0 {
				tmp_featureSetsDownlinkPerCC_v1700 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1700]([]*FeatureSetDownlinkPerCC_v1700{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.featureSetsDownlinkPerCC_v1700 {
					tmp_featureSetsDownlinkPerCC_v1700.Value = append(tmp_featureSetsDownlinkPerCC_v1700.Value, &i)
				}
				if err = tmp_featureSetsDownlinkPerCC_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlinkPerCC_v1700", err)
				}
			}
			// encode featureSetsUplink_v1710 optional
			if len(ie.featureSetsUplink_v1710) > 0 {
				tmp_featureSetsUplink_v1710 := utils.NewSequence[*FeatureSetUplink_v1710]([]*FeatureSetUplink_v1710{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.featureSetsUplink_v1710 {
					tmp_featureSetsUplink_v1710.Value = append(tmp_featureSetsUplink_v1710.Value, &i)
				}
				if err = tmp_featureSetsUplink_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplink_v1710", err)
				}
			}
			// encode featureSetsUplinkPerCC_v1700 optional
			if len(ie.featureSetsUplinkPerCC_v1700) > 0 {
				tmp_featureSetsUplinkPerCC_v1700 := utils.NewSequence[*FeatureSetUplinkPerCC_v1700]([]*FeatureSetUplinkPerCC_v1700{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.featureSetsUplinkPerCC_v1700 {
					tmp_featureSetsUplinkPerCC_v1700.Value = append(tmp_featureSetsUplinkPerCC_v1700.Value, &i)
				}
				if err = tmp_featureSetsUplinkPerCC_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplinkPerCC_v1700", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{len(ie.featureSetsDownlink_v1720) > 0, len(ie.featureSetsDownlinkPerCC_v1720) > 0, len(ie.featureSetsUplink_v1720) > 0}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsDownlink_v1720 optional
			if len(ie.featureSetsDownlink_v1720) > 0 {
				tmp_featureSetsDownlink_v1720 := utils.NewSequence[*FeatureSetDownlink_v1720]([]*FeatureSetDownlink_v1720{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.featureSetsDownlink_v1720 {
					tmp_featureSetsDownlink_v1720.Value = append(tmp_featureSetsDownlink_v1720.Value, &i)
				}
				if err = tmp_featureSetsDownlink_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlink_v1720", err)
				}
			}
			// encode featureSetsDownlinkPerCC_v1720 optional
			if len(ie.featureSetsDownlinkPerCC_v1720) > 0 {
				tmp_featureSetsDownlinkPerCC_v1720 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1720]([]*FeatureSetDownlinkPerCC_v1720{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.featureSetsDownlinkPerCC_v1720 {
					tmp_featureSetsDownlinkPerCC_v1720.Value = append(tmp_featureSetsDownlinkPerCC_v1720.Value, &i)
				}
				if err = tmp_featureSetsDownlinkPerCC_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlinkPerCC_v1720", err)
				}
			}
			// encode featureSetsUplink_v1720 optional
			if len(ie.featureSetsUplink_v1720) > 0 {
				tmp_featureSetsUplink_v1720 := utils.NewSequence[*FeatureSetUplink_v1720]([]*FeatureSetUplink_v1720{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				for _, i := range ie.featureSetsUplink_v1720 {
					tmp_featureSetsUplink_v1720.Value = append(tmp_featureSetsUplink_v1720.Value, &i)
				}
				if err = tmp_featureSetsUplink_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsUplink_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{len(ie.featureSetsDownlink_v1730) > 0, len(ie.featureSetsDownlinkPerCC_v1730) > 0}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featureSetsDownlink_v1730 optional
			if len(ie.featureSetsDownlink_v1730) > 0 {
				tmp_featureSetsDownlink_v1730 := utils.NewSequence[*FeatureSetDownlink_v1730]([]*FeatureSetDownlink_v1730{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				for _, i := range ie.featureSetsDownlink_v1730 {
					tmp_featureSetsDownlink_v1730.Value = append(tmp_featureSetsDownlink_v1730.Value, &i)
				}
				if err = tmp_featureSetsDownlink_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlink_v1730", err)
				}
			}
			// encode featureSetsDownlinkPerCC_v1730 optional
			if len(ie.featureSetsDownlinkPerCC_v1730) > 0 {
				tmp_featureSetsDownlinkPerCC_v1730 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1730]([]*FeatureSetDownlinkPerCC_v1730{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				for _, i := range ie.featureSetsDownlinkPerCC_v1730 {
					tmp_featureSetsDownlinkPerCC_v1730.Value = append(tmp_featureSetsDownlinkPerCC_v1730.Value, &i)
				}
				if err = tmp_featureSetsDownlinkPerCC_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureSetsDownlinkPerCC_v1730", err)
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

func (ie *FeatureSets) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetsDownlinkPresent bool
	if featureSetsDownlinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetsDownlinkPerCCPresent bool
	if featureSetsDownlinkPerCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetsUplinkPresent bool
	if featureSetsUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetsUplinkPerCCPresent bool
	if featureSetsUplinkPerCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if featureSetsDownlinkPresent {
		tmp_featureSetsDownlink := utils.NewSequence[*FeatureSetDownlink]([]*FeatureSetDownlink{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
		fn_featureSetsDownlink := func() *FeatureSetDownlink {
			return new(FeatureSetDownlink)
		}
		if err = tmp_featureSetsDownlink.Decode(r, fn_featureSetsDownlink); err != nil {
			return utils.WrapError("Decode featureSetsDownlink", err)
		}
		ie.featureSetsDownlink = []FeatureSetDownlink{}
		for _, i := range tmp_featureSetsDownlink.Value {
			ie.featureSetsDownlink = append(ie.featureSetsDownlink, *i)
		}
	}
	if featureSetsDownlinkPerCCPresent {
		tmp_featureSetsDownlinkPerCC := utils.NewSequence[*FeatureSetDownlinkPerCC]([]*FeatureSetDownlinkPerCC{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		fn_featureSetsDownlinkPerCC := func() *FeatureSetDownlinkPerCC {
			return new(FeatureSetDownlinkPerCC)
		}
		if err = tmp_featureSetsDownlinkPerCC.Decode(r, fn_featureSetsDownlinkPerCC); err != nil {
			return utils.WrapError("Decode featureSetsDownlinkPerCC", err)
		}
		ie.featureSetsDownlinkPerCC = []FeatureSetDownlinkPerCC{}
		for _, i := range tmp_featureSetsDownlinkPerCC.Value {
			ie.featureSetsDownlinkPerCC = append(ie.featureSetsDownlinkPerCC, *i)
		}
	}
	if featureSetsUplinkPresent {
		tmp_featureSetsUplink := utils.NewSequence[*FeatureSetUplink]([]*FeatureSetUplink{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
		fn_featureSetsUplink := func() *FeatureSetUplink {
			return new(FeatureSetUplink)
		}
		if err = tmp_featureSetsUplink.Decode(r, fn_featureSetsUplink); err != nil {
			return utils.WrapError("Decode featureSetsUplink", err)
		}
		ie.featureSetsUplink = []FeatureSetUplink{}
		for _, i := range tmp_featureSetsUplink.Value {
			ie.featureSetsUplink = append(ie.featureSetsUplink, *i)
		}
	}
	if featureSetsUplinkPerCCPresent {
		tmp_featureSetsUplinkPerCC := utils.NewSequence[*FeatureSetUplinkPerCC]([]*FeatureSetUplinkPerCC{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
		fn_featureSetsUplinkPerCC := func() *FeatureSetUplinkPerCC {
			return new(FeatureSetUplinkPerCC)
		}
		if err = tmp_featureSetsUplinkPerCC.Decode(r, fn_featureSetsUplinkPerCC); err != nil {
			return utils.WrapError("Decode featureSetsUplinkPerCC", err)
		}
		ie.featureSetsUplinkPerCC = []FeatureSetUplinkPerCC{}
		for _, i := range tmp_featureSetsUplinkPerCC.Value {
			ie.featureSetsUplinkPerCC = append(ie.featureSetsUplinkPerCC, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 8 bits for 8 extension groups
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

			featureSetsDownlink_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsUplink_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsUplinkPerCC_v1540Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsDownlink_v1540 optional
			if featureSetsDownlink_v1540Present {
				tmp_featureSetsDownlink_v1540 := utils.NewSequence[*FeatureSetDownlink_v1540]([]*FeatureSetDownlink_v1540{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_featureSetsDownlink_v1540 := func() *FeatureSetDownlink_v1540 {
					return new(FeatureSetDownlink_v1540)
				}
				if err = tmp_featureSetsDownlink_v1540.Decode(extReader, fn_featureSetsDownlink_v1540); err != nil {
					return utils.WrapError("Decode featureSetsDownlink_v1540", err)
				}
				ie.featureSetsDownlink_v1540 = []FeatureSetDownlink_v1540{}
				for _, i := range tmp_featureSetsDownlink_v1540.Value {
					ie.featureSetsDownlink_v1540 = append(ie.featureSetsDownlink_v1540, *i)
				}
			}
			// decode featureSetsUplink_v1540 optional
			if featureSetsUplink_v1540Present {
				tmp_featureSetsUplink_v1540 := utils.NewSequence[*FeatureSetUplink_v1540]([]*FeatureSetUplink_v1540{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_featureSetsUplink_v1540 := func() *FeatureSetUplink_v1540 {
					return new(FeatureSetUplink_v1540)
				}
				if err = tmp_featureSetsUplink_v1540.Decode(extReader, fn_featureSetsUplink_v1540); err != nil {
					return utils.WrapError("Decode featureSetsUplink_v1540", err)
				}
				ie.featureSetsUplink_v1540 = []FeatureSetUplink_v1540{}
				for _, i := range tmp_featureSetsUplink_v1540.Value {
					ie.featureSetsUplink_v1540 = append(ie.featureSetsUplink_v1540, *i)
				}
			}
			// decode featureSetsUplinkPerCC_v1540 optional
			if featureSetsUplinkPerCC_v1540Present {
				tmp_featureSetsUplinkPerCC_v1540 := utils.NewSequence[*FeatureSetUplinkPerCC_v1540]([]*FeatureSetUplinkPerCC_v1540{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_featureSetsUplinkPerCC_v1540 := func() *FeatureSetUplinkPerCC_v1540 {
					return new(FeatureSetUplinkPerCC_v1540)
				}
				if err = tmp_featureSetsUplinkPerCC_v1540.Decode(extReader, fn_featureSetsUplinkPerCC_v1540); err != nil {
					return utils.WrapError("Decode featureSetsUplinkPerCC_v1540", err)
				}
				ie.featureSetsUplinkPerCC_v1540 = []FeatureSetUplinkPerCC_v1540{}
				for _, i := range tmp_featureSetsUplinkPerCC_v1540.Value {
					ie.featureSetsUplinkPerCC_v1540 = append(ie.featureSetsUplinkPerCC_v1540, *i)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsDownlink_v15a0Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsDownlink_v15a0 optional
			if featureSetsDownlink_v15a0Present {
				tmp_featureSetsDownlink_v15a0 := utils.NewSequence[*FeatureSetDownlink_v15a0]([]*FeatureSetDownlink_v15a0{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_featureSetsDownlink_v15a0 := func() *FeatureSetDownlink_v15a0 {
					return new(FeatureSetDownlink_v15a0)
				}
				if err = tmp_featureSetsDownlink_v15a0.Decode(extReader, fn_featureSetsDownlink_v15a0); err != nil {
					return utils.WrapError("Decode featureSetsDownlink_v15a0", err)
				}
				ie.featureSetsDownlink_v15a0 = []FeatureSetDownlink_v15a0{}
				for _, i := range tmp_featureSetsDownlink_v15a0.Value {
					ie.featureSetsDownlink_v15a0 = append(ie.featureSetsDownlink_v15a0, *i)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsDownlink_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsUplink_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetDownlinkPerCC_v1620Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsDownlink_v1610 optional
			if featureSetsDownlink_v1610Present {
				tmp_featureSetsDownlink_v1610 := utils.NewSequence[*FeatureSetDownlink_v1610]([]*FeatureSetDownlink_v1610{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_featureSetsDownlink_v1610 := func() *FeatureSetDownlink_v1610 {
					return new(FeatureSetDownlink_v1610)
				}
				if err = tmp_featureSetsDownlink_v1610.Decode(extReader, fn_featureSetsDownlink_v1610); err != nil {
					return utils.WrapError("Decode featureSetsDownlink_v1610", err)
				}
				ie.featureSetsDownlink_v1610 = []FeatureSetDownlink_v1610{}
				for _, i := range tmp_featureSetsDownlink_v1610.Value {
					ie.featureSetsDownlink_v1610 = append(ie.featureSetsDownlink_v1610, *i)
				}
			}
			// decode featureSetsUplink_v1610 optional
			if featureSetsUplink_v1610Present {
				tmp_featureSetsUplink_v1610 := utils.NewSequence[*FeatureSetUplink_v1610]([]*FeatureSetUplink_v1610{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_featureSetsUplink_v1610 := func() *FeatureSetUplink_v1610 {
					return new(FeatureSetUplink_v1610)
				}
				if err = tmp_featureSetsUplink_v1610.Decode(extReader, fn_featureSetsUplink_v1610); err != nil {
					return utils.WrapError("Decode featureSetsUplink_v1610", err)
				}
				ie.featureSetsUplink_v1610 = []FeatureSetUplink_v1610{}
				for _, i := range tmp_featureSetsUplink_v1610.Value {
					ie.featureSetsUplink_v1610 = append(ie.featureSetsUplink_v1610, *i)
				}
			}
			// decode featureSetDownlinkPerCC_v1620 optional
			if featureSetDownlinkPerCC_v1620Present {
				tmp_featureSetDownlinkPerCC_v1620 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1620]([]*FeatureSetDownlinkPerCC_v1620{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_featureSetDownlinkPerCC_v1620 := func() *FeatureSetDownlinkPerCC_v1620 {
					return new(FeatureSetDownlinkPerCC_v1620)
				}
				if err = tmp_featureSetDownlinkPerCC_v1620.Decode(extReader, fn_featureSetDownlinkPerCC_v1620); err != nil {
					return utils.WrapError("Decode featureSetDownlinkPerCC_v1620", err)
				}
				ie.featureSetDownlinkPerCC_v1620 = []FeatureSetDownlinkPerCC_v1620{}
				for _, i := range tmp_featureSetDownlinkPerCC_v1620.Value {
					ie.featureSetDownlinkPerCC_v1620 = append(ie.featureSetDownlinkPerCC_v1620, *i)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsUplink_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsUplink_v1630 optional
			if featureSetsUplink_v1630Present {
				tmp_featureSetsUplink_v1630 := utils.NewSequence[*FeatureSetUplink_v1630]([]*FeatureSetUplink_v1630{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_featureSetsUplink_v1630 := func() *FeatureSetUplink_v1630 {
					return new(FeatureSetUplink_v1630)
				}
				if err = tmp_featureSetsUplink_v1630.Decode(extReader, fn_featureSetsUplink_v1630); err != nil {
					return utils.WrapError("Decode featureSetsUplink_v1630", err)
				}
				ie.featureSetsUplink_v1630 = []FeatureSetUplink_v1630{}
				for _, i := range tmp_featureSetsUplink_v1630.Value {
					ie.featureSetsUplink_v1630 = append(ie.featureSetsUplink_v1630, *i)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsUplink_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsUplink_v1640 optional
			if featureSetsUplink_v1640Present {
				tmp_featureSetsUplink_v1640 := utils.NewSequence[*FeatureSetUplink_v1640]([]*FeatureSetUplink_v1640{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_featureSetsUplink_v1640 := func() *FeatureSetUplink_v1640 {
					return new(FeatureSetUplink_v1640)
				}
				if err = tmp_featureSetsUplink_v1640.Decode(extReader, fn_featureSetsUplink_v1640); err != nil {
					return utils.WrapError("Decode featureSetsUplink_v1640", err)
				}
				ie.featureSetsUplink_v1640 = []FeatureSetUplink_v1640{}
				for _, i := range tmp_featureSetsUplink_v1640.Value {
					ie.featureSetsUplink_v1640 = append(ie.featureSetsUplink_v1640, *i)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsDownlink_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsDownlinkPerCC_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsUplink_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsUplinkPerCC_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsDownlink_v1700 optional
			if featureSetsDownlink_v1700Present {
				tmp_featureSetsDownlink_v1700 := utils.NewSequence[*FeatureSetDownlink_v1700]([]*FeatureSetDownlink_v1700{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_featureSetsDownlink_v1700 := func() *FeatureSetDownlink_v1700 {
					return new(FeatureSetDownlink_v1700)
				}
				if err = tmp_featureSetsDownlink_v1700.Decode(extReader, fn_featureSetsDownlink_v1700); err != nil {
					return utils.WrapError("Decode featureSetsDownlink_v1700", err)
				}
				ie.featureSetsDownlink_v1700 = []FeatureSetDownlink_v1700{}
				for _, i := range tmp_featureSetsDownlink_v1700.Value {
					ie.featureSetsDownlink_v1700 = append(ie.featureSetsDownlink_v1700, *i)
				}
			}
			// decode featureSetsDownlinkPerCC_v1700 optional
			if featureSetsDownlinkPerCC_v1700Present {
				tmp_featureSetsDownlinkPerCC_v1700 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1700]([]*FeatureSetDownlinkPerCC_v1700{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_featureSetsDownlinkPerCC_v1700 := func() *FeatureSetDownlinkPerCC_v1700 {
					return new(FeatureSetDownlinkPerCC_v1700)
				}
				if err = tmp_featureSetsDownlinkPerCC_v1700.Decode(extReader, fn_featureSetsDownlinkPerCC_v1700); err != nil {
					return utils.WrapError("Decode featureSetsDownlinkPerCC_v1700", err)
				}
				ie.featureSetsDownlinkPerCC_v1700 = []FeatureSetDownlinkPerCC_v1700{}
				for _, i := range tmp_featureSetsDownlinkPerCC_v1700.Value {
					ie.featureSetsDownlinkPerCC_v1700 = append(ie.featureSetsDownlinkPerCC_v1700, *i)
				}
			}
			// decode featureSetsUplink_v1710 optional
			if featureSetsUplink_v1710Present {
				tmp_featureSetsUplink_v1710 := utils.NewSequence[*FeatureSetUplink_v1710]([]*FeatureSetUplink_v1710{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_featureSetsUplink_v1710 := func() *FeatureSetUplink_v1710 {
					return new(FeatureSetUplink_v1710)
				}
				if err = tmp_featureSetsUplink_v1710.Decode(extReader, fn_featureSetsUplink_v1710); err != nil {
					return utils.WrapError("Decode featureSetsUplink_v1710", err)
				}
				ie.featureSetsUplink_v1710 = []FeatureSetUplink_v1710{}
				for _, i := range tmp_featureSetsUplink_v1710.Value {
					ie.featureSetsUplink_v1710 = append(ie.featureSetsUplink_v1710, *i)
				}
			}
			// decode featureSetsUplinkPerCC_v1700 optional
			if featureSetsUplinkPerCC_v1700Present {
				tmp_featureSetsUplinkPerCC_v1700 := utils.NewSequence[*FeatureSetUplinkPerCC_v1700]([]*FeatureSetUplinkPerCC_v1700{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_featureSetsUplinkPerCC_v1700 := func() *FeatureSetUplinkPerCC_v1700 {
					return new(FeatureSetUplinkPerCC_v1700)
				}
				if err = tmp_featureSetsUplinkPerCC_v1700.Decode(extReader, fn_featureSetsUplinkPerCC_v1700); err != nil {
					return utils.WrapError("Decode featureSetsUplinkPerCC_v1700", err)
				}
				ie.featureSetsUplinkPerCC_v1700 = []FeatureSetUplinkPerCC_v1700{}
				for _, i := range tmp_featureSetsUplinkPerCC_v1700.Value {
					ie.featureSetsUplinkPerCC_v1700 = append(ie.featureSetsUplinkPerCC_v1700, *i)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsDownlink_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsDownlinkPerCC_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsUplink_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsDownlink_v1720 optional
			if featureSetsDownlink_v1720Present {
				tmp_featureSetsDownlink_v1720 := utils.NewSequence[*FeatureSetDownlink_v1720]([]*FeatureSetDownlink_v1720{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_featureSetsDownlink_v1720 := func() *FeatureSetDownlink_v1720 {
					return new(FeatureSetDownlink_v1720)
				}
				if err = tmp_featureSetsDownlink_v1720.Decode(extReader, fn_featureSetsDownlink_v1720); err != nil {
					return utils.WrapError("Decode featureSetsDownlink_v1720", err)
				}
				ie.featureSetsDownlink_v1720 = []FeatureSetDownlink_v1720{}
				for _, i := range tmp_featureSetsDownlink_v1720.Value {
					ie.featureSetsDownlink_v1720 = append(ie.featureSetsDownlink_v1720, *i)
				}
			}
			// decode featureSetsDownlinkPerCC_v1720 optional
			if featureSetsDownlinkPerCC_v1720Present {
				tmp_featureSetsDownlinkPerCC_v1720 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1720]([]*FeatureSetDownlinkPerCC_v1720{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_featureSetsDownlinkPerCC_v1720 := func() *FeatureSetDownlinkPerCC_v1720 {
					return new(FeatureSetDownlinkPerCC_v1720)
				}
				if err = tmp_featureSetsDownlinkPerCC_v1720.Decode(extReader, fn_featureSetsDownlinkPerCC_v1720); err != nil {
					return utils.WrapError("Decode featureSetsDownlinkPerCC_v1720", err)
				}
				ie.featureSetsDownlinkPerCC_v1720 = []FeatureSetDownlinkPerCC_v1720{}
				for _, i := range tmp_featureSetsDownlinkPerCC_v1720.Value {
					ie.featureSetsDownlinkPerCC_v1720 = append(ie.featureSetsDownlinkPerCC_v1720, *i)
				}
			}
			// decode featureSetsUplink_v1720 optional
			if featureSetsUplink_v1720Present {
				tmp_featureSetsUplink_v1720 := utils.NewSequence[*FeatureSetUplink_v1720]([]*FeatureSetUplink_v1720{}, uper.Constraint{Lb: 1, Ub: maxUplinkFeatureSets}, false)
				fn_featureSetsUplink_v1720 := func() *FeatureSetUplink_v1720 {
					return new(FeatureSetUplink_v1720)
				}
				if err = tmp_featureSetsUplink_v1720.Decode(extReader, fn_featureSetsUplink_v1720); err != nil {
					return utils.WrapError("Decode featureSetsUplink_v1720", err)
				}
				ie.featureSetsUplink_v1720 = []FeatureSetUplink_v1720{}
				for _, i := range tmp_featureSetsUplink_v1720.Value {
					ie.featureSetsUplink_v1720 = append(ie.featureSetsUplink_v1720, *i)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featureSetsDownlink_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureSetsDownlinkPerCC_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featureSetsDownlink_v1730 optional
			if featureSetsDownlink_v1730Present {
				tmp_featureSetsDownlink_v1730 := utils.NewSequence[*FeatureSetDownlink_v1730]([]*FeatureSetDownlink_v1730{}, uper.Constraint{Lb: 1, Ub: maxDownlinkFeatureSets}, false)
				fn_featureSetsDownlink_v1730 := func() *FeatureSetDownlink_v1730 {
					return new(FeatureSetDownlink_v1730)
				}
				if err = tmp_featureSetsDownlink_v1730.Decode(extReader, fn_featureSetsDownlink_v1730); err != nil {
					return utils.WrapError("Decode featureSetsDownlink_v1730", err)
				}
				ie.featureSetsDownlink_v1730 = []FeatureSetDownlink_v1730{}
				for _, i := range tmp_featureSetsDownlink_v1730.Value {
					ie.featureSetsDownlink_v1730 = append(ie.featureSetsDownlink_v1730, *i)
				}
			}
			// decode featureSetsDownlinkPerCC_v1730 optional
			if featureSetsDownlinkPerCC_v1730Present {
				tmp_featureSetsDownlinkPerCC_v1730 := utils.NewSequence[*FeatureSetDownlinkPerCC_v1730]([]*FeatureSetDownlinkPerCC_v1730{}, uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false)
				fn_featureSetsDownlinkPerCC_v1730 := func() *FeatureSetDownlinkPerCC_v1730 {
					return new(FeatureSetDownlinkPerCC_v1730)
				}
				if err = tmp_featureSetsDownlinkPerCC_v1730.Decode(extReader, fn_featureSetsDownlinkPerCC_v1730); err != nil {
					return utils.WrapError("Decode featureSetsDownlinkPerCC_v1730", err)
				}
				ie.featureSetsDownlinkPerCC_v1730 = []FeatureSetDownlinkPerCC_v1730{}
				for _, i := range tmp_featureSetsDownlinkPerCC_v1730.Value {
					ie.featureSetsDownlinkPerCC_v1730 = append(ie.featureSetsDownlinkPerCC_v1730, *i)
				}
			}
		}
	}
	return nil
}
