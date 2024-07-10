// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_resource

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func ResourceResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"json": schema.StringAttribute{
				CustomType: jsontypes.NormalizedType{},
				Optional:   true,
			},
			"list_nested": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"json": schema.StringAttribute{
							CustomType: jsontypes.NormalizedType{},
							Optional:   true,
						},
					},
					CustomType: ListNestedType{
						ObjectType: types.ObjectType{
							AttrTypes: ListNestedValue{}.AttributeTypes(ctx),
						},
					},
				},
				Optional: true,
			},
			"single_nested": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"json": schema.StringAttribute{
						CustomType: jsontypes.NormalizedType{},
						Optional:   true,
					},
				},
				CustomType: SingleNestedType{
					ObjectType: types.ObjectType{
						AttrTypes: SingleNestedValue{}.AttributeTypes(ctx),
					},
				},
				Optional: true,
			},
		},
	}
}

type ResourceModel struct {
	Json         jsontypes.Normalized `tfsdk:"json"`
	ListNested   types.List           `tfsdk:"list_nested"`
	SingleNested SingleNestedValue    `tfsdk:"single_nested"`
}

var _ basetypes.ObjectTypable = ListNestedType{}

type ListNestedType struct {
	basetypes.ObjectType
}

func (t ListNestedType) Equal(o attr.Type) bool {
	other, ok := o.(ListNestedType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ListNestedType) String() string {
	return "ListNestedType"
}

func (t ListNestedType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	jsonAttribute, ok := attributes["json"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`json is missing from object`)

		return nil, diags
	}

	jsonVal, ok := jsonAttribute.(jsontypes.Normalized)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`json expected to be jsontypes.Normalized, was: %T`, jsonAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ListNestedValue{
		Json:  jsonVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewListNestedValueNull() ListNestedValue {
	return ListNestedValue{
		state: attr.ValueStateNull,
	}
}

func NewListNestedValueUnknown() ListNestedValue {
	return ListNestedValue{
		state: attr.ValueStateUnknown,
	}
}

func NewListNestedValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ListNestedValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ListNestedValue Attribute Value",
				"While creating a ListNestedValue value, a missing attribute value was detected. "+
					"A ListNestedValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ListNestedValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ListNestedValue Attribute Type",
				"While creating a ListNestedValue value, an invalid attribute value was detected. "+
					"A ListNestedValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ListNestedValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ListNestedValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ListNestedValue Attribute Value",
				"While creating a ListNestedValue value, an extra attribute value was detected. "+
					"A ListNestedValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ListNestedValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewListNestedValueUnknown(), diags
	}

	jsonAttribute, ok := attributes["json"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`json is missing from object`)

		return NewListNestedValueUnknown(), diags
	}

	jsonVal, ok := jsonAttribute.(jsontypes.Normalized)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`json expected to be jsontypes.Normalized, was: %T`, jsonAttribute))
	}

	if diags.HasError() {
		return NewListNestedValueUnknown(), diags
	}

	return ListNestedValue{
		Json:  jsonVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewListNestedValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ListNestedValue {
	object, diags := NewListNestedValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewListNestedValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ListNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewListNestedValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewListNestedValueUnknown(), nil
	}

	if in.IsNull() {
		return NewListNestedValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewListNestedValueMust(ListNestedValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ListNestedType) ValueType(ctx context.Context) attr.Value {
	return ListNestedValue{}
}

var _ basetypes.ObjectValuable = ListNestedValue{}

type ListNestedValue struct {
	Json  jsontypes.Normalized `tfsdk:"json"`
	state attr.ValueState
}

func (v ListNestedValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["json"] = jsontypes.NormalizedType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.Json.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["json"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ListNestedValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ListNestedValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ListNestedValue) String() string {
	return "ListNestedValue"
}

func (v ListNestedValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"json": jsontypes.NormalizedType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"json": v.Json,
		})

	return objVal, diags
}

func (v ListNestedValue) Equal(o attr.Value) bool {
	other, ok := o.(ListNestedValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Json.Equal(other.Json) {
		return false
	}

	return true
}

func (v ListNestedValue) Type(ctx context.Context) attr.Type {
	return ListNestedType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ListNestedValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"json": jsontypes.NormalizedType{},
	}
}

var _ basetypes.ObjectTypable = SingleNestedType{}

type SingleNestedType struct {
	basetypes.ObjectType
}

func (t SingleNestedType) Equal(o attr.Type) bool {
	other, ok := o.(SingleNestedType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t SingleNestedType) String() string {
	return "SingleNestedType"
}

func (t SingleNestedType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	jsonAttribute, ok := attributes["json"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`json is missing from object`)

		return nil, diags
	}

	jsonVal, ok := jsonAttribute.(jsontypes.Normalized)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`json expected to be jsontypes.Normalized, was: %T`, jsonAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return SingleNestedValue{
		Json:  jsonVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewSingleNestedValueNull() SingleNestedValue {
	return SingleNestedValue{
		state: attr.ValueStateNull,
	}
}

func NewSingleNestedValueUnknown() SingleNestedValue {
	return SingleNestedValue{
		state: attr.ValueStateUnknown,
	}
}

func NewSingleNestedValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (SingleNestedValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing SingleNestedValue Attribute Value",
				"While creating a SingleNestedValue value, a missing attribute value was detected. "+
					"A SingleNestedValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SingleNestedValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid SingleNestedValue Attribute Type",
				"While creating a SingleNestedValue value, an invalid attribute value was detected. "+
					"A SingleNestedValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SingleNestedValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("SingleNestedValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra SingleNestedValue Attribute Value",
				"While creating a SingleNestedValue value, an extra attribute value was detected. "+
					"A SingleNestedValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra SingleNestedValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewSingleNestedValueUnknown(), diags
	}

	jsonAttribute, ok := attributes["json"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`json is missing from object`)

		return NewSingleNestedValueUnknown(), diags
	}

	jsonVal, ok := jsonAttribute.(jsontypes.Normalized)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`json expected to be jsontypes.Normalized, was: %T`, jsonAttribute))
	}

	if diags.HasError() {
		return NewSingleNestedValueUnknown(), diags
	}

	return SingleNestedValue{
		Json:  jsonVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewSingleNestedValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) SingleNestedValue {
	object, diags := NewSingleNestedValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewSingleNestedValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t SingleNestedType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewSingleNestedValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewSingleNestedValueUnknown(), nil
	}

	if in.IsNull() {
		return NewSingleNestedValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewSingleNestedValueMust(SingleNestedValue{}.AttributeTypes(ctx), attributes), nil
}

func (t SingleNestedType) ValueType(ctx context.Context) attr.Value {
	return SingleNestedValue{}
}

var _ basetypes.ObjectValuable = SingleNestedValue{}

type SingleNestedValue struct {
	Json  jsontypes.Normalized `tfsdk:"json"`
	state attr.ValueState
}

func (v SingleNestedValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["json"] = jsontypes.NormalizedType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.Json.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["json"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v SingleNestedValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v SingleNestedValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v SingleNestedValue) String() string {
	return "SingleNestedValue"
}

func (v SingleNestedValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"json": jsontypes.NormalizedType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"json": v.Json,
		})

	return objVal, diags
}

func (v SingleNestedValue) Equal(o attr.Value) bool {
	other, ok := o.(SingleNestedValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Json.Equal(other.Json) {
		return false
	}

	return true
}

func (v SingleNestedValue) Type(ctx context.Context) attr.Type {
	return SingleNestedType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v SingleNestedValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"json": jsontypes.NormalizedType{},
	}
}
