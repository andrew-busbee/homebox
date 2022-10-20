// Code generated by ent, DO NOT EDIT.

package item

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImportRef holds the string denoting the import_ref field in the database.
	FieldImportRef = "import_ref"
	// FieldNotes holds the string denoting the notes field in the database.
	FieldNotes = "notes"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldInsured holds the string denoting the insured field in the database.
	FieldInsured = "insured"
	// FieldSerialNumber holds the string denoting the serial_number field in the database.
	FieldSerialNumber = "serial_number"
	// FieldModelNumber holds the string denoting the model_number field in the database.
	FieldModelNumber = "model_number"
	// FieldManufacturer holds the string denoting the manufacturer field in the database.
	FieldManufacturer = "manufacturer"
	// FieldLifetimeWarranty holds the string denoting the lifetime_warranty field in the database.
	FieldLifetimeWarranty = "lifetime_warranty"
	// FieldWarrantyExpires holds the string denoting the warranty_expires field in the database.
	FieldWarrantyExpires = "warranty_expires"
	// FieldWarrantyDetails holds the string denoting the warranty_details field in the database.
	FieldWarrantyDetails = "warranty_details"
	// FieldPurchaseTime holds the string denoting the purchase_time field in the database.
	FieldPurchaseTime = "purchase_time"
	// FieldPurchaseFrom holds the string denoting the purchase_from field in the database.
	FieldPurchaseFrom = "purchase_from"
	// FieldPurchasePrice holds the string denoting the purchase_price field in the database.
	FieldPurchasePrice = "purchase_price"
	// FieldSoldTime holds the string denoting the sold_time field in the database.
	FieldSoldTime = "sold_time"
	// FieldSoldTo holds the string denoting the sold_to field in the database.
	FieldSoldTo = "sold_to"
	// FieldSoldPrice holds the string denoting the sold_price field in the database.
	FieldSoldPrice = "sold_price"
	// FieldSoldNotes holds the string denoting the sold_notes field in the database.
	FieldSoldNotes = "sold_notes"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeLabel holds the string denoting the label edge name in mutations.
	EdgeLabel = "label"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeFields holds the string denoting the fields edge name in mutations.
	EdgeFields = "fields"
	// EdgeAttachments holds the string denoting the attachments edge name in mutations.
	EdgeAttachments = "attachments"
	// Table holds the table name of the item in the database.
	Table = "items"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "items"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_items"
	// LabelTable is the table that holds the label relation/edge. The primary key declared below.
	LabelTable = "label_items"
	// LabelInverseTable is the table name for the Label entity.
	// It exists in this package in order to avoid circular dependency with the "label" package.
	LabelInverseTable = "labels"
	// LocationTable is the table that holds the location relation/edge.
	LocationTable = "items"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "location_items"
	// FieldsTable is the table that holds the fields relation/edge.
	FieldsTable = "item_fields"
	// FieldsInverseTable is the table name for the ItemField entity.
	// It exists in this package in order to avoid circular dependency with the "itemfield" package.
	FieldsInverseTable = "item_fields"
	// FieldsColumn is the table column denoting the fields relation/edge.
	FieldsColumn = "item_fields"
	// AttachmentsTable is the table that holds the attachments relation/edge.
	AttachmentsTable = "attachments"
	// AttachmentsInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentsInverseTable = "attachments"
	// AttachmentsColumn is the table column denoting the attachments relation/edge.
	AttachmentsColumn = "item_attachments"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldImportRef,
	FieldNotes,
	FieldQuantity,
	FieldInsured,
	FieldSerialNumber,
	FieldModelNumber,
	FieldManufacturer,
	FieldLifetimeWarranty,
	FieldWarrantyExpires,
	FieldWarrantyDetails,
	FieldPurchaseTime,
	FieldPurchaseFrom,
	FieldPurchasePrice,
	FieldSoldTime,
	FieldSoldTo,
	FieldSoldPrice,
	FieldSoldNotes,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_items",
	"location_items",
}

var (
	// LabelPrimaryKey and LabelColumn2 are the table columns denoting the
	// primary key for the label relation (M2M).
	LabelPrimaryKey = []string{"label_id", "item_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// ImportRefValidator is a validator for the "import_ref" field. It is called by the builders before save.
	ImportRefValidator func(string) error
	// NotesValidator is a validator for the "notes" field. It is called by the builders before save.
	NotesValidator func(string) error
	// DefaultQuantity holds the default value on creation for the "quantity" field.
	DefaultQuantity int
	// DefaultInsured holds the default value on creation for the "insured" field.
	DefaultInsured bool
	// SerialNumberValidator is a validator for the "serial_number" field. It is called by the builders before save.
	SerialNumberValidator func(string) error
	// ModelNumberValidator is a validator for the "model_number" field. It is called by the builders before save.
	ModelNumberValidator func(string) error
	// ManufacturerValidator is a validator for the "manufacturer" field. It is called by the builders before save.
	ManufacturerValidator func(string) error
	// DefaultLifetimeWarranty holds the default value on creation for the "lifetime_warranty" field.
	DefaultLifetimeWarranty bool
	// WarrantyDetailsValidator is a validator for the "warranty_details" field. It is called by the builders before save.
	WarrantyDetailsValidator func(string) error
	// DefaultPurchasePrice holds the default value on creation for the "purchase_price" field.
	DefaultPurchasePrice float64
	// DefaultSoldPrice holds the default value on creation for the "sold_price" field.
	DefaultSoldPrice float64
	// SoldNotesValidator is a validator for the "sold_notes" field. It is called by the builders before save.
	SoldNotesValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)