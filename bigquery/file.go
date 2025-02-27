Delete









































































}

// FileConfig contains configuration options that pertain to files, typically
// text files that require interpretation to be used as a BigQuery table. A
// file may live in Google Cloud Storage (see GCSReference), or it may be
// loaded into a table via the Table.LoaderFromReader.
type FileConfig struct {
	// SourceFormat is the format of the data to be read.
	// Allowed values are: Avro, CSV, DatastoreBackup, JSON, ORC, and Parquet.  The default is CSV.
	SourceFormat DataFormat

	// Indicates if we should automatically infer the options and
	// schema for CSV and JSON sources.
	AutoDetect bool

	// MaxBadRecords is the maximum number of bad records that will be ignored
	// when reading data.
	MaxBadRecords int64

	// IgnoreUnknownValues causes values not matching the schema to be
	// tolerated. Unknown values are ignored. For CSV this ignores extra values
	// at the end of a line. For JSON this ignores named values that do not
	// match any column name. If this field is not set, records containing
	// unknown values are treated as bad records. The MaxBadRecords field can
	// be used to customize how bad records are handled.
	IgnoreUnknownValues bool

	// Schema describes the data. It is required when reading CSV or JSON data,
	// unless the data is being loaded into a table that already exists.
	Schema Schema

	// Additional options for CSV files.
	CSVOptions

	// Additional options for Parquet files.
	ParquetOptions *ParquetOptions

	// Additional options for Avro files.
	AvroOptions *AvroOptions
}

func (fc *FileConfig) populateLoadConfig(conf *bq.JobConfigurationLoad) {
	conf.SkipLeadingRows = fc.SkipLeadingRows
	conf.SourceFormat = string(fc.SourceFormat)
	conf.Autodetect = fc.AutoDetect
	conf.AllowJaggedRows = fc.AllowJaggedRows
	conf.AllowQuotedNewlines = fc.AllowQuotedNewlines
	conf.Encoding = string(fc.Encoding)
	conf.FieldDelimiter = fc.FieldDelimiter
	conf.IgnoreUnknownValues = fc.IgnoreUnknownValues
	conf.MaxBadRecords = fc.MaxBadRecords
	conf.NullMarker = fc.NullMarker
	conf.PreserveAsciiControlCharacters = fc.PreserveASCIIControlCharacters
	if fc.Schema != nil {
		conf.Schema = fc.Schema.toBQ()
	}
	if fc.ParquetOptions != nil {
		conf.ParquetOptions = &bq.ParquetOptions{
			EnumAsString:        fc.ParquetOptions.EnumAsString,
			EnableListInference: fc.ParquetOptions.EnableListInference,
		}
	}
	if fc.AvroOptions != nil {
		conf.UseAvroLogicalTypes = fc.AvroOptions.UseAvroLogicalTypes
	}
	conf.Quote = fc.quote()
}

func bqPopulateFileConfig(conf *bq.JobConfigurationLoad, fc *FileConfig) {
	fc.SourceFormat = DataFormat(conf.SourceFormat)
	fc.AutoDetect = conf.Autodetect
	fc.MaxBadRecords = conf.MaxBadRecords
	fc.IgnoreUnknownValues = conf.IgnoreUnknownValues
	fc.Schema = bqToSchema(conf.Schema)
	fc.SkipLeadingRows = conf.SkipLeadingRows
	fc.AllowJaggedRows = conf.AllowJaggedRows
	fc.AllowQuotedNewlines = conf.AllowQuotedNewlines
	fc.Encoding = Encoding(conf.Encoding)
	fc.FieldDelimiter = conf.FieldDelimiter
	fc.CSVOptions.NullMarker = conf.NullMarker
	fc.CSVOptions.PreserveASCIIControlCharacters = conf.PreserveAsciiControlCharacters
	fc.CSVOptions.setQuote(conf.Quote)
}

func (fc *FileConfig) populateExternalDataConfig(conf *bq.ExternalDataConfiguration) {
	format := fc.SourceFormat
	if format == "" {
		// Format must be explicitly set for external data sources.
		format = CSV
	}
	conf.Autodetect = fc.AutoDetect
	conf.IgnoreUnknownValues = fc.IgnoreUnknownValues
	conf.MaxBadRecords = fc.MaxBadRecords
	conf.SourceFormat = string(format)
	if fc.Schema != nil {
		conf.Schema = fc.Schema.toBQ()
	}
	if format == CSV {
		fc.CSVOptions.populateExternalDataConfig(conf)
	}
	if fc.AvroOptions != nil {
		conf.AvroOptions = &bq.AvroOptions{
			UseAvroLogicalTypes: fc.AvroOptions.UseAvroLogicalTypes,
		}
	}
	if fc.ParquetOptions != nil {
		conf.ParquetOptions = &bq.ParquetOptions{
			EnumAsString:        fc.ParquetOptions.EnumAsString,
			EnableListInference: fc.ParquetOptions.EnableListInference,
		}
	}
}

// Encoding specifies the character encoding of data to be loaded into BigQuery.
// See https://cloud.google.com/bigquery/docs/reference/v2/jobs#configuration.load.encoding
// for more details about how this is used.
type Encoding string

const (
	// UTF_8 specifies the UTF-8 encoding type.
	UTF_8 Encoding = "UTF-8"
	// ISO_8859_1 specifies the ISO-8859-1 encoding type.
	ISO_8859_1 Encoding = "ISO-8859-1"
)
