package hkmg

type Option interface {
	apply(options *Options)
}

func (f FieldNameFormat) apply(options *Options) {
	options.fieldNameFc = f
}

func WithFieldNameFormat(format FieldNameFormat) Option {
	return format
}

func (f TableNameFormat) apply(options *Options) {
	options.tableNameFc = f
}

func WithTableNameFormat(format TableNameFormat) Option {
	return format
}
